#!/bin/bash

echo "Stopping exist Process..."
webdict_server_pid=`ps -ef | grep webdict_server | grep -v "grep" | awk '{print $2}'`
for id in $webdict_server_pid
do
    kill -9 $id
    echo "Killed $id"
done
echo "Stopped exist Process"