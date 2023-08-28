package boot_config

import (
	"github.com/mimis-s/golang_tools/app"
	"github.com/mimis-s/golang_tools/dfs"
)

/*
	解析boot_config.yaml文件
*/

type BootFlags struct {
	IP              string `flag:"ip" default:"192.168.1.1" desc:"服务器地址"`
	Port            string `flag:"port" default:"8888" desc:"服务器tcp端口"`
	WebPort         string `flag:"web_port" default:"8889" desc:"服务器http端口"`
	RpcExposePort   string `flag:"rpc_expose_port" default:"8890" desc:"集群使用的服务之间的rpc监听端口"`
	RpcListenPort   string `flag:"rpc_listen_port" default:"9999" desc:"本地服务之间的rpc监听端口, 非集群情况下等于exposer_port"`
	FileServiceIP   string `flag:"file_service_ip" default:"localhost" desc:"文件服务ip"`
	FileServicePort string `flag:"file_service_port" default:"8999" desc:"文件服务端口"`
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

var BootConfigData = new(BootConfig)
var BootFlagsData = new(app.GlobalCmdFlag)
var CustomBootFlagsData = new(BootFlags)
