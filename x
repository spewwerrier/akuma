#!/usr/bin/env bash

mkdir ./internal/interface/akuma/assets
cp -r static web internal/interface/akuma/assets
go build cmd/akuma/main.go
