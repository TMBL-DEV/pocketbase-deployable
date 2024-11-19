#!/bin/bash
go build -o pocketbase

./pocketbase migrate
./pocketbase serve --http=0.0.0.0:8090