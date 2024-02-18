#!/usr/bin/env bash

go generate internal/interface/akuma/embed.go
go build -o akuma cmd/akuma/main.go
