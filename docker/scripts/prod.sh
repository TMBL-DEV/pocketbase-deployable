#!/bin/bash
go mod tidy
go mod vendor
go build -o pocketbase

./pocketbase migrate
./pocketbase serve --http=0.0.0.0:8090