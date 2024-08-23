#!/bin/bash

rm main
go build -o main cmd/ddos/main.go
./main