#!/bin/bash

SOURCE="$0"
while [ -h "$SOURCE"  ]; do # resolve $SOURCE until the file is no longer a symlink
    DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /*  ]] && SOURCE="$DIR/$SOURCE" # if $SOURCE was a relative symlink, we need to resolve it relative to the path where the symlink file was located
done
DIR="$( cd -P "$( dirname "$SOURCE"  )" && pwd  )"

thrift -strict -out "$DIR/../thrift_go"  --gen go "$DIR/../conf/UserSer.thrift"
thrift -strict -out "$DIR/../thrift_android"  --gen java:android "$DIR/../conf/UserSer.thrift"
#thrift -strict -out "$DIR/../thrift_js"  --gen js "$DIR/../api/UserSer.thrift"

pgsql_map  -in "$DIR/../conf/db.json" -o "$DIR/../pkg/models/type.go"