#!/bin/bash

server=`netstat -nlpa | grep ':10101' | wc -l`

if [[ $server -eq 0 ]]; then
    ./api.bin >> ../logs/apis.log &
fi

echo "Started \n";