#!/bin/bash

CERTS="../etc"
ROACH="./cockroach"

$ROACH start --certs-dir=$CERTS \
      --store=../node2 \
      --listen-addr=localhost:26258 \
      --http-addr=localhost:8081 \
      --join=localhost:26257,localhost:26258,localhost:26259
      #--background
