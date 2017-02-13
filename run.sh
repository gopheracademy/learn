#!/bin/bash

set -e

buffalo db migrate up
go get ./...
go run main.go
