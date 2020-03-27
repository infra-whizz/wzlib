#!/bin/bash

CAKEY="../etc/ca.key"
CERTS="../etc"
ROACH="./cockroach"

$ROACH init \
       --certs-dir=$CERTS \
       --host=localhost:26257

# Add admin
$ROACH sql \
       --certs-dir=$CERTS \
       --host=localhost:26258 < create_admin.sql
