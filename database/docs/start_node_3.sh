#!/bin/bash

CERTS="../etc"
ROACH="./cockroach"

$ROACH start --certs-dir=$CERTS \
      --store=../node3 \
      --listen-addr=localhost:26259 \
      --http-addr=localhost:8082 \
      --join=localhost:26257,localhost:26258,localhost:26259
      #--background
