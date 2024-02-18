#!/usr/bin/env bash

 rm -rf internal/interface/akuma/assets/
 mkdir ./internal/interface/akuma/assets
cp -r static web internal/interface/akuma/assets
go build cmd/akuma/main.go
