package boot_config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"github.com/mimis-s/golang_tools/dfs"
	"gopkg.in/yaml.v3"
)

/*
	解析boot_config.yaml文件
*/

type BootFlags struct {
	BootConfigPath string `flag:"boot_config_path|.|yaml文件路径"`
	Port           string `flag:"port|8888|tcp端口"`
	RpcExposePort  string `flag:"rpc_expose_port|8890|集群使用的服务之间的rpc监听端口"`
	RpcListenPort  string `flag:"rpc_listen_port|9999|本地服务之间的rpc监听端口, 非集群情况下等于exposer_port"`
}

type DBConfig struct {
	Addr     string `yaml:"addr"`
	DBName   string `yaml:"db_name"` // 数据库名字
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type DBManageConfig struct {
	Tag    string      `yaml:"tag"`
	Master DBConfig    `yaml:"master"` // 主数据库
	Slaves []*DBConfig `yaml:"slaves"` // 从数据库
}

type BootConfig struct {
	Server string `yaml:"server"`
	Etcd   struct {
		Timeout      int      `yaml:"time_out"`      // 超时时间(单位秒)
		EtcdBasePath string   `yaml:"etc_base_path"` // etcd中存储的路径
		Addrs        []string `yaml:"addrs"`         // etcd地址
	} `yaml:"etcd"`

	DB []DBManageConfig `yaml:"db"`

	Redis struct {
		Addr     string `yaml:"addr"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
	} `yaml:"redis"`

	MQ struct {
		Durable bool   `yaml:"durable"` // 数据持久的
		Url     string `yaml:"url"`
	} `yaml:"mq"`

	Log struct {
		Path string `yaml:"path"`
	} `yaml:"log"`

	DFS dfs.Config `yaml:"dfs"`

	IsLocal bool `yaml:"is_local"` // 所有服务一个本地进程启动

	DataBaseShard bool `yaml:"database_shard"` // 数据库是否分表
	IsTest        bool `yaml:"is_test"`        // 是否是测试环境
}

type ConfigOptions struct {
	BootConfigFile *BootConfig
	CommandFlags   *BootFlags
}

// 解析启服参数
func ParseBootConfigOptions() *ConfigOptions {

	bootConfig := &ConfigOptions{
		BootConfigFile: &BootConfig{},
		CommandFlags:   &BootFlags{},
	}

	// 解析命令行参数
	parseBootFlagStruct(bootConfig.CommandFlags)

	yamlFile, err := ioutil.ReadFile(bootConfig.CommandFlags.BootConfigPath)
	if err != nil {
		errStr := fmt.Sprintf("read yaml path[%v] is err:%v", bootConfig.CommandFlags.BootConfigPath, err)
		panic(errStr)
	}

	err = yaml.Unmarshal(yamlFile, bootConfig.BootConfigFile)
	if err != nil {
		errStr := fmt.Sprintf("unmarshal yaml[%v] is err:%v", yamlFile, err)
		panic(errStr)
	}
	return bootConfig
}

// 原始数据
type flagOriginDesc struct {
	flag         string // 参数
	value        reflect.Value
	fieldStruct  reflect.StructField
	flagValuePtr interface{}
}

// 反射赋值
func setValue(field reflect.Value, strData string) {
	switch field.Kind() {
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(strData)
		if err != nil {
			errStr := fmt.Sprintf("parse bool value[%v] is err:%v", strData, err)
			panic(errStr)
		}
		field.SetBool(boolValue)
	case reflect.String:
		field.SetString(strData)
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16, reflect.Int8:
		if strData == "" {
			field.SetInt(0)
		} else {
			intValue, err := strconv.ParseInt(strData, 10, field.Type().Bits())
			if err != nil {
				errStr := fmt.Sprintf("parse int value[%v] is err:%v", strData, err)
				panic(errStr)
			}
			field.SetInt(intValue)
		}
	case reflect.Float32, reflect.Float64:
		if strData == "" {
			field.SetFloat(0)
		} else {
			floatValue, err := strconv.ParseFloat(strData, field.Type().Bits())
			if err != nil {
				errStr := fmt.Sprintf("parse float value[%v] is err:%v", strData, err)
				panic(errStr)
			}
			field.SetFloat(floatValue)
		}
	default:
		errStr := fmt.Sprintf("parse value type[%v] not support", field.Kind().String())
		panic(errStr)
	}
}

// 初始化原始数据结构体
func initParseFlagOriginDesc(filedValue reflect.Value, filedStruct reflect.StructField, tagName string) *flagOriginDesc {

	flagValue, find := filedStruct.Tag.Lookup(tagName)
	if find {

		commonds := strings.Split(flagValue, "|")
		if len(commonds) != 3 {
			// 格式必须是 参数|默认值|解释 (可以为空)
			errStr := fmt.Sprintf("cmd arg[%v] tag:%v value must split 2 '|'", filedStruct.Name, tagName)
			panic(errStr)
		}

		cmd := commonds[0]
		defaultValue := commonds[1]
		desc := commonds[2]

		fOriginDesc := &flagOriginDesc{
			flag:        cmd,
			value:       filedValue,
			fieldStruct: filedStruct,
		}

		switch filedStruct.Type.Kind() {
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(defaultValue)
			if err != nil {
				errStr := fmt.Sprintf("cmd arg[%v] tag:%v default value[%v] is err:%v", filedStruct.Name, tagName, defaultValue, err)
				panic(errStr)
			}
			fOriginDesc.flagValuePtr = flag.CommandLine.Bool(cmd, boolValue, desc)
			fOriginDesc.value.SetBool(*fOriginDesc.flagValuePtr.(*bool))
		case reflect.Int, reflect.Int64:
			intValue, err := strconv.ParseInt(defaultValue, 10, 64)
			if err != nil {
				errStr := fmt.Sprintf("cmd arg[%v] tag:%v default value[%v] is err:%v", filedStruct.Name, tagName, defaultValue, err)
				panic(errStr)
			}
			fOriginDesc.flagValuePtr = flag.CommandLine.Int64(cmd, intValue, desc)
			fOriginDesc.value.SetInt(*fOriginDesc.flagValuePtr.(*int64))
		case reflect.String:
			fOriginDesc.flagValuePtr = flag.CommandLine.String(cmd, defaultValue, desc)
			fOriginDesc.value.SetString(*fOriginDesc.flagValuePtr.(*string))
		case reflect.Slice:
			fOriginDesc.flagValuePtr = flag.CommandLine.String(cmd, defaultValue, desc)
			strValues := strings.Split(*fOriginDesc.flagValuePtr.(*string), ",")
			if len(strValues) == 1 && strValues[0] == "" {
				strValues = []string{}
			}
			fOriginDesc.value.Set(reflect.MakeSlice(fOriginDesc.value.Type(), len(strValues), len(strValues)))
			for i := 0; i < len(strValues); i++ {
				setValue(fOriginDesc.value.Index(i), strValues[i])
			}

		case reflect.Float32, reflect.Float64:
			floatValue, err := strconv.ParseFloat(defaultValue, 64)
			if err != nil {
				errStr := fmt.Sprintf("cmd arg[%v] tag:%v default value[%v] is err:%v", filedStruct.Name, tagName, defaultValue, err)
				panic(errStr)
			}
			fOriginDesc.flagValuePtr = flag.CommandLine.Float64(cmd, floatValue, desc)
			fOriginDesc.value.SetFloat(*fOriginDesc.flagValuePtr.(*float64))
		default:
			errStr := fmt.Sprintf("cmd arg[%v] tag:%v type[%v] not support", filedStruct.Name, tagName, filedStruct.Type.Name())
			panic(errStr)
		}
		return fOriginDesc
	}
	return nil
}

func parseBootFlagStruct(data *BootFlags) {

	flag.Parse()

	// 首先初始化原始结构体
	tagName := "flag"
	rValue := reflect.ValueOf(data).Elem() // 值反射
	rType := reflect.TypeOf(data).Elem()   // 类型反射

	for i := 0; i < rValue.NumField(); i++ {
		filedValue := rValue.Field(i)
		filedStruct := rType.Field(i)
		// 初始化参数,为参数赋值
		initParseFlagOriginDesc(filedValue, filedStruct, tagName)
	}
}
