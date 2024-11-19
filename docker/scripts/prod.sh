#!/bin/bash
go mod tidy
go mod vendor
go build -o pocketbase

./pocketbase serve --http=0.0.0.0:8090