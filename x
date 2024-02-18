#!/usr/bin/env bash

go generate internal/interface/akuma/embed.go
go build cmd/akuma/main.go
