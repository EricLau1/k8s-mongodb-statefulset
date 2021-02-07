#!/bin/bash

go build -o main .

docker build -t mongo-app:v1 .

docker images
