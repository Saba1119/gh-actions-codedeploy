#!/bin/bash
# echo "stopping go application"
# ps -ef | grep goapp | grep -v grep | awk '{print $2}' | sudo xargs kill
killall goapp
exit 0