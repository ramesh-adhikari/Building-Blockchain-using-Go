#!/bin/bash

# Define the number of shards and nodes per shard
num_shards=1
nodes_per_shard=1
start_port=3000
BLOCKCHAIN_SERVER_PORT_RANGE_START=5001
BLOCKCHAIN_SERVER_PORT_RANGE_END=5003
WALLET_SERVER_ONE_PORT=8000
WALLET_SERVER_ONE_GATEWAT=http://127.0.0.1:5001
WALLET_SERVER_TWO_PORT=8001
WALLET_SERVER_TWO_GATEWAT=http://127.0.0.1:5002


# Loop through each blockchain servers to start server
for ((b_server = "$BLOCKCHAIN_SERVER_PORT_RANGE_START"; b_server <= "$BLOCKCHAIN_SERVER_PORT_RANGE_END"; b_server++)); do
        go run blockchain/server/main.go blockchain/server/server.go -port "$b_server" &
done

# Wait for a brief moment to ensure servers have started
sleep 25

# start wlalet server 1 and 2
go run wallet/server/main.go wallet/server/server.go -port "$WALLET_SERVER_ONE_PORT" -gateway "$WALLET_SERVER_ONE_GATEWAT" &
sleep 10
go run wallet/server/main.go wallet/server/server.go -port "$WALLET_SERVER_TWO_PORT" -gateway "$WALLET_SERVER_TWO_GATEWAT" &

# Wait for all background processes to finish
wait

#### To find the process ID from port on windows ###
# netstat -ano | findstr :<PORT>
#### To kill the task by process ID####
# taskkill //PID <PORT>

# for mac
# npx kill-port 3000