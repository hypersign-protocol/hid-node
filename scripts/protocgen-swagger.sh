#!/bin/bash

set -eux pipefail

echo "Generating OpenAPI Docs"

mkdir -p ./tmp-swagger
cd proto

# Find all proto files
proto_dirs=$(find ./hypersign -type f -path '*/client/*' -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

echo $proto_dirs

for dir in $proto_dirs; do
  query_file=$(find "${dir}" -maxdepth 1 -name 'query.proto')
  if [[ -n "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml "$query_file"
  fi
done

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
cd ..
swagger-combine ./app/client/docs/config.json \
-o ./app/client/docs/static/swagger.yaml -f yaml \
--continueOnConflictingPaths true \
--includeDefinitions true

rm -rf ./tmp-swagger