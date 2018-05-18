#!/bin/bash

killall -9 simple-proxy
nohup $SIMPLE_PROXY/simple-proxy >> $SIMPLE_PROXY/log/simple-proxy.log 2>&1 &