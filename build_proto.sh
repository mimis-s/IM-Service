#!/bin/bash

set -e

project_path="$(cd "$(dirname "$0")" && pwd)"
root_client_proto="${project_path}/src/proto"
out_path="${project_path}/src/common/commonproto"
services_root="${project_path}/src/services"
mkdir -p $out_path

echo "[INFO] ==> compile common proto out_path:"$out_path

mkdir -p $out_path/github.com/mimis-s/IM-Service/src/common
ln -sf $out_path/ $out_path/github.com/mimis-s/IM-Service/src/common/commonproto

errCode="0"

echo_red() {
    str=$1
    echo -e "\033[31m$str\033[0m"
}

echo_yellow() {
    str=$1
    echo -e "\033[33m$str\033[0m"
}

gen_client_proto_go() {
    proto=$1
    # 要生成golang代码结构
    protoc -I$root_client_proto/ --gogofaster_out=$out_path $root_client_proto/$proto

    # 上一个命令执行退出状态不等于0, 则说明出错了
    if [ $? -ne 0 ]; then
        # errCode="1"
        echo_red "[ERROR] ==> compile $proto golang not ok."
        exit 1
    else
        echo "[INFO] ==> compile $proto golang ok."
    fi
}

gen_service_proto() {
    path=$1
    log_path=$(basename "$(dirname $path)")/$(basename $path)
    cd $path || (echo_red "cd $path error" && exit 1)
    protoc -I$root_client_proto -I. --gogofaster_out=. --rpcx_out=. *.proto
    if [ $? != 0 ]; then
        errCode="1"
        echo_red "[ERROR] ==> compile $log_path not ok"
        exit 1
    else
        echo "[INFO] ==> compile $log_path ok."
    fi
    #   cd -
}

echo "[INFO] ==> start compile common proto."
all_client_protos=$(find $root_client_proto -name "*.proto" -type f)
for client_proto in $all_client_protos; do
    baseName=$(basename $client_proto)
    gen_client_proto_go $baseName
done

echo "[INFO] ==> compile all common proto finish."

echo "[INFO] ==> start compile RPC services proto."
all_services_api_path=$(find ${services_root} -name "api_*" -type d)
for api_path in $all_services_api_path; do
    gen_service_proto $api_path
done
echo "[INFO] ==> compile all services proto finish."

rm -rf $out_path/github.com/mimis-s/IM-Service/src/common/commonproto
rm -rf $out_path/github.com
rm -rf $out_path/commonproto

if [ $errCode != "0" ]; then
    echo_red "[ERROR] ==> compile has error, please check output."
fi
