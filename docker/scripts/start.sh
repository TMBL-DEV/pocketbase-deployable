#!/bin/bash
go mod tidy
go mod vendor

go run main.go serve --http=0.0.0.0:8090