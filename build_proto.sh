#!/bin/sh

project_path="$( cd "$( dirname "$0"  )" && pwd  )"
root_client_proto="$( cd "$( dirname "$0"  )" && pwd  )"/src/proto
root_server_proto="$( cd "$( dirname "$0"  )" && pwd  )"/src/proto_server
out_path="$( cd $root_client_proto && cd .. && pwd)"/common/commonproto
services_root="$( cd $root_client_proto && cd .. && pwd)"/services
mkdir -p $out_path
echo "[INFO] ==> compile common proto out_path:"$out_path

errCode="0"

# 创造一个目录软链接，将生成的文件链接到生成目录，然后删除软链接
mkdir -p $out_path/im-service/src/common/
ln -sf $out_path/ $out_path/im-service/src/common/commonproto

echo_red() {
    str=$1
    echo -e "\033[31m$str\033[0m"
}

echo_yellow() {
    str=$1
    echo -e "\033[33m$str\033[0m"
}

gen_client_proto() {
   path=$1
   proto=$2
#   ret=`protoc -I$root/ --gogofaster_out=$out_path $root/$proto 2>&1`
   protoc -I$path/ --gogofaster_out=$out_path $path/$proto
   if [ $? -ne 0 ]; then
       errCode="1"
       echo_red "[ERROR] ==> compile $proto not ok."
       exit 1
   else
       echo "[INFO] ==> compile $proto ok."
   fi
}

gen_server_proto() {
   path_server=$1
   path_client=$2
   proto=$3
#   ret=`protoc -I$root/ --gogofaster_out=$out_path $root/$proto 2>&1`
   protoc -I$path_server/ -I$path_client/ --gogofaster_out=$out_path $path_server/$proto
   if [ $? -ne 0 ]; then
       errCode="1"
       echo_red "[ERROR] ==> compile $proto not ok."
       exit 1
   else
       echo "[INFO] ==> compile $proto ok."
   fi
}

gen_service_proto() {
   path=$1
   log_path=$(basename "$(dirname $path)")/$(basename $path)
   cd $path || (echo_red "cd $path error" && exit 1)
   protoc -I$root_client_proto -I$root_server_proto -I. --gogofaster_out=. --rpcx_out=. *.proto
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

all_client_protos=`find $project_path/src/proto -name "*.proto" -type f`
for client_proto in $all_client_protos; do
#   echo "[INFO] ==> compile path:"$api_path
   baseName=`basename $client_proto`
   gen_client_proto $root_client_proto $baseName
done

all_server_protos=`find $project_path/src/proto_server -name "*.proto" -type f`
#gen_server_proto $root_server_proto $root_client_proto server.proto
for server_proto in $all_server_protos; do
#   echo "[INFO] ==> compile path:"$api_path
   baseName=`basename $server_proto`
   gen_server_proto $root_server_proto $root_client_proto $baseName
done

echo "[INFO] ==> compile all common proto finish."

rm -rf $out_path/im-service/src/common/commonproto
rm -rf $out_path/im-service
rm -rf $out_path/commonproto

echo
echo "[INFO] ==> start compile services proto."
all_services_api_path=`find $project_path/src/services -name "api_*" -type d`
for api_path in $all_services_api_path; do
#   echo "[INFO] ==> compile path:"$api_path
   gen_service_proto $api_path
done
echo "[INFO] ==> compile all services proto finish."

if [ $errCode != "0" ]; then
    echo_red "[ERROR] ==> compile has error, please check output."
fi