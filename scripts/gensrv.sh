#!/bin/sh

goSwaggerVersion="v0.29.0"
targetSpecPath="./api-spec/swagger.json"

gobin=${GOPATH}/bin
goSwaggerBin=${gobin}/go-swagger-${goSwaggerVersion}

if [ ! -f ${goSwaggerBin} ]; then
  echo "install go-swagger..."
  go install github.com/go-swagger/go-swagger/cmd/swagger@${goSwaggerVersion}
  mv ${gobin}/swagger ${goSwaggerBin}
fi

httpPath="internal/server/http"
modelsPath="${httpPath}/models"
operationsPath="${httpPath}/operations"

# Clear prev
rm -rf ${modelsPath}
rm -rf ${operationsPath}

# Generate
${goSwaggerBin} generate server  -f ${targetSpecPath} \
        --model-package=${modelsPath} --server-package=${httpPath} --exclude-main
