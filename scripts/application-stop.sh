#!/bin/bash
echo "stopping go application"
ps -ef | grep main.go | grep -v grep | awk '{print $2}' | xargs kill
