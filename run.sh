#!/bin/bash

set -e

buffalo db migrate up
go run main.go
