#!/bin/bash

CAKEY="../etc/ca.key"
CERTS="../etc"
ROACH="./cockroach"

$ROACH node status --certs-dir=$CERTS

$ROACH quit --certs-dir=$CERTS --host=localhost:26257
$ROACH quit --certs-dir=$CERTS --host=localhost:26258
$ROACH quit --certs-dir=$CERTS --host=localhost:26259
