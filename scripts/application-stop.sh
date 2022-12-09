#!/bin/bash
# echo "stopping go application"
# ps -ef | grep goapp | grep -v grep | awk '{print $2}' | sudo xargs kill
# kill goapp
# exit 0

pid=`pgrep goapp`
if [ -n "$pid" ]
then
	kill $pid
fi