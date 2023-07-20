#!/bin/bash

set -e
project_path="$(cd "$(dirname "$0")" && pwd)"

echo "[INFO] ==> start compile service wire."
all_wires=$(find $project_path/src/services -name "wire.go" -type f)
for w in $all_wires; do
    wire $w
    echo "[INFO] ==> compile path:"$api_path
done
