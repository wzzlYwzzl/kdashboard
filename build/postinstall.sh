#!/bin/bash

# This script will run after npm install

../node_modules/.bin/bower install --allow-root

# The pwd cmd will return a path where the program which calls the script. 
# Not the path of the postinstall.sh.
GOPATH=`pwd`/.tools/go 
go get github.com/tools/godep