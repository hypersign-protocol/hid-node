#!/usr/bin/env bash
set -eux pipefail

# generated-ts
generated_ts_dir=./generated-ts

# Remove generated files director to update changes
rm -rf $generated_ts_dir

# Make a generated files directory
mkdir $generated_ts_dir

cd proto

# Get the path of the cosmos-sdk repo from go/pkg/mod
proto_dirs=$(find ./hypersign -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for proto_dir in $proto_dirs; do
  proto_files=$(find "${proto_dir}" -maxdepth 1 -name '*.proto')
  for f in $proto_files; do
    buf generate --template buf.gen.ts.yaml "$f"
  done
done

cd ..

echo "TS files are generated at $generated_ts_dir"