#!/usr/bin/env bash
set -o errexit

FILES=(
  ../desumasu-converter/test/*
)

i=0
for file in "${FILES[@]}"; do
  echo "Processing $file with -j option"
  cp "$file" "./tests/$i-input.text"
  ../desumasu-converter/dist/cli.js -j <"$file" >"./tests/$i-j-output.text"
  diff --color "./tests/$i-j-output.text" <(go run ./cmd/main.go -j <"./tests/$i-input.text")
  echo "Processing $file with -k option"
  ../desumasu-converter/dist/cli.js -k <"$file" >"./tests/$i-k-output.text"
  diff --color "./tests/$i-k-output.text" <(go run ./cmd/main.go -k <"$file")
  echo "Processing $file with -j -n option"
  ../desumasu-converter/dist/cli.js -j -n <"$file" >"./tests/$i-jn-output.text"
  diff --color "./tests/$i-jn-output.text" <(go run ./cmd/main.go -j -n <"$file")
  ((i++)) || true
done
