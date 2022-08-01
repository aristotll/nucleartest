#!/bin/sh
#trap "ps -ef | grep nsq | grep -v grep | awk '{print $2}' | xargs kill" SIGINT

# 神坑：= 前后不能有空格，不然会识别为 command
pid="$$"

trap "kill $pid" SIGINT
echo "PID: " $pid

/opt/homebrew/Cellar/nsq/1.2.1/bin/nsqlookupd &
/opt/homebrew/Cellar/nsq/1.2.1/bin/nsqd --lookupd-tcp-address=127.0.0.1:4160 -data-path=/opt/homebrew/var/nsq &
/opt/homebrew/Cellar/nsq/1.2.1/bin/nsqadmin --lookupd-http-address=127.0.0.1:4161 &

wait