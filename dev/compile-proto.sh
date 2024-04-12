#!/usr/bin/env bash

# Configs

PROTO_DIR="./pkg/proto"

# API (model, service)
API_DIR="${PROTO_DIR}/api"
UNITY_API_FILEPATH="${API_DIR}/order.api.proto"
HEAD_API_DIR="${API_DIR}/head.proto"

# Common
INCLUDES_DIR="${PROTO_DIR}/includes"
INCLUDES="-I${INCLUDES_DIR} -I${PROTO_DIR}"
PACKAGE_DIR="pkg/"
GO_PACKAGE="pkg/api"
SWAGGER_DIR="${PROTO_DIR}/swagger"

OS="unknown: $OSTYPE"
case "$OSTYPE" in
darwin*) OS="darwin" ;;
linux*) OS="linux" ;;
esac

export PATH="${PROTO_DIR}/bin/${OS}:$PATH"

# Generate unity for internals (model + grpc service)
echo "" >${UNITY_API_FILEPATH}
cat ${HEAD_API_DIR} >> ${UNITY_API_FILEPATH}

for file in $(grep -rl "//Union//" ${API_DIR}); do
  cat ${file} | sed -n -e '/\/\/Union\/\//,$p' >> ${UNITY_API_FILEPATH}
done

# Generate internal pb
protoc ${INCLUDES} --go_out=plugins=grpc:${PACKAGE_DIR} --go_opt=paths=source_relative ${UNITY_API_FILEPATH}

## Remove unities
if [ -f $UNITY_API_FILEPATH ]; then
   rm -rf $UNITY_API_FILEPATH
   echo "$UNITY_API_FILEPATH is removed"
fi
