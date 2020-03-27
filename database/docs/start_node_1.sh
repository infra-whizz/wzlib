#!/bin/bash

CERTS="../etc"
ROACH="./cockroach"

$ROACH start --certs-dir=$CERTS \
      --store=../node1 \
      --listen-addr=localhost:26257 \
      --http-addr=localhost:8080 \
      --join=localhost:26257,localhost:26258,localhost:26259
      #--background
