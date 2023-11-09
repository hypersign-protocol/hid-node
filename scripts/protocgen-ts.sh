#!/usr/bin/env bash
set -eux pipefail

# Install ts-proto plugin
npm i ts-proto

# generated-ts
generated_ts_dir=./generated-ts

# Remove generated files director to update changes
rm -rf $generated_ts_dir

# Make a generated files directory
mkdir $generated_ts_dir

# Get the path of the cosmos-sdk repo from go/pkg/mod
cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk)
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate protobuf bind
  protoc \
  --plugin ./node_modules/.bin/protoc-gen-ts_proto \
  --ts_proto_out=$generated_ts_dir \
  --ts_proto_opt=useOptionals=all \
  --ts_proto_opt=initializeFieldsAsUndefined=false \
  --ts_proto_opt=unrecognizedEnum=false \
  --ts_proto_opt=useJsonName=true \
  --ts_proto_opt=esModuleInterop=true \
  --ts_proto_opt=stringEnums=true \
  --ts_proto_opt=outputExtensions=true \
  -I "proto" \
  -I "third_party/proto" \
  -I "$cosmos_sdk_dir/third_party/proto" \
  -I "$cosmos_sdk_dir/proto" \
  $(find "${dir}" -name '*.proto')

done

rm -rf ./node_modules package-lock.json

echo "TS files are generated at $generated_ts_dir"