#!/bin/bash

command="run"

while getopts 'cs:' flag; do
  case "${flag}" in
  c) command=${OPTARG} ;;
  *) break ;;
  esac
done

case ${command} in
run) go run . ;;
build) goscript build ;;
install) go install ;;
*) go run one.go ;;
esac
