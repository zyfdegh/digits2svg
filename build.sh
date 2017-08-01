#!/bin/sh

go get github.com/ajstarks/svgo

export GOARCH=arm
export GOARM=7

go build -o bin/digits2svg

docker build -t 192.168.10.60:5000/digits2svg:dev .
docker push 192.168.10.60:5000/digits2svg:dev
