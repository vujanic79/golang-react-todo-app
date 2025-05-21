#!/bin/bash

export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin

go install github.com/jstemmer/go-junit-report/v2@latest

go test -v server/... | go-junit-report > report.xml