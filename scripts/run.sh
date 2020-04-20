#!/usr/bin/env bash

pid=0

# SIGTERM-handler
term_handler() {
  if [ $pid -ne 0 ]; then
    kill -SIGTERM "$children"
    ## Waiting children to gracefully die
    sleep 20
    ## Killing now actual the process launcher! Bye bye
    kill -SIGTERM "$pid"
    wait "$pid"
  fi
  exit 143; # 128 + 15 -- SIGTERM
}

# Setting up SIGTERM trap handler
trap 'kill ${!}; term_handler' SIGTERM

# Run child application
"$@" &

# Fetching the pid
pid="$!"

# Fetching grand-child application, typically the webserver (scala/node)
children=""
child_pids() {
  children=$(ps --ppid "$pid" | awk 'NR>1 { printf $1 }')
}

while [ -z "$children"  ]; do
  sleep 1
  child_pids
done

# Wait forever, but allow logs to flow, even after SIGTERM
while true
do
  tail -f /dev/null & wait ${!}
done