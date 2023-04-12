#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o noclobber
set -o pipefail

cd "$(dirname "$0")"
cd ..

for ((i = 0; i < 5; i++)); do
  file="./tests/$i-input.text"

  echo "Processing $file with -j option"
  diff --color "./tests/$i-j-output.text" <(go run ./cmd/main.go -j <"$file")

  echo "Processing $file with -k option"
  diff --color "./tests/$i-k-output.text" <(go run ./cmd/main.go -k <"$file")

  echo "Processing $file with -j -n option"
  diff --color "./tests/$i-jn-output.text" <(go run ./cmd/main.go -j -n <"$file")
done
