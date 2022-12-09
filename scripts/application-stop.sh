#!/bin/bash
# echo "stopping go application"
# ps -ef | grep goapp | grep -v grep | awk '{print $2}' | sudo xargs kill
killall app
exit 0