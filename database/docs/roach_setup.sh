#!/bin/bash

CAKEY="../etc/ca.key"
CERTS="../etc"
ROACH="./cockroach"

# Generate key
$ROACH cert create-ca \
       --certs-dir=$CERTS \
       --ca-key=$CAKEY

# Create node key
$ROACH cert create-node \
       localhost \
       $(hostname) \
       --certs-dir=$CERTS \
       --ca-key=$CAKEY

# Create root key
$ROACH cert create-client \
       root \
       --certs-dir=$CERTS \
       --ca-key=$CAKEY
