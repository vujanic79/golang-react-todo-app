#!/bin/bash
export BUILD_NUMBER=$1
docker-compose -f ../../../docker-compose.yaml build