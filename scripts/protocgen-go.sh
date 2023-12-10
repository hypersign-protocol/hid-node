#!/bin/bash

set -eux pipefail

# Get protoc-gen-gocosmos

echo "Generating gogo proto code"
cd proto

# Find all proto files
proto_dirs=$(find ./hypersign -type f -path '*/client/*' -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

echo $proto_dirs

for proto_dir in $proto_dirs; do
  proto_files=$(find "${proto_dir}" -maxdepth 1 -name '*.proto')
  for f in $proto_files; do
    if grep -q "option go_package" "$f"; then
      buf generate --template buf.gen.gogo.yaml "$f"
    fi
  done
done

# move proto files to the right places
cd ..
ls

cp -r github.com/hypersign-protocol/hid-node/* ./
rm -rf github.com

go mod tidy
