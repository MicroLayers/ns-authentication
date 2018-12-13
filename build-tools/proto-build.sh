#!/usr/bin/env bash

### Compiles the proto messages into go objects

protoc_min_version="3.5.0" # tested on 3.5.0
protoc_gen_go_min_version="1.2.0"

current_directory="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

source "${current_directory}/common.sh"


function checkProtoc {
  checkBinary protoc "${protoc_min_version}"
  checkBinary protoc-gen-go "${protoc_gen_go_min_version}"

  protoc_version=$(protoc --version | awk '{print $2}')

  echo "protoc: $protoc_version"
}

checkProtoc

cd "${current_directory}/.."
echo ""
echo "Compiling proto messages"
protoc --go_out=. messages/*.proto
