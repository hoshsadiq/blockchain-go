#!/usr/bin/env bash

go build -i -o ./build/tescoin main.go

./build/tescoin server --port 8080
./build/tescoin server --port 8081
./build/tescoin server --port 8082

curl -sSL -X POST -d '["localhost:8081","localhost:8082"]' http://localhost:8080/nodes/register
curl -sSL -X POST -d '["localhost:8080","localhost:8082"]' http://localhost:8081/nodes/register
curl -sSL -X POST -d '["localhost:8080","localhost:8081"]' http://localhost:8082/nodes/register

curl -sSL -X GET http://localhost:8080/nodes | jq
curl -sSL -X GET http://localhost:8081/nodes | jq
curl -sSL -X GET http://localhost:8082/nodes | jq

curl -sSL -X POST -d '{"sender":"Hosh","recipient":"David","amount":1.5}' http://localhost:8080/transactions/new | jq
curl -sSL -X POST -d '{"sender":"Hosh","recipient":"Tescoin Foundation","amount":1.5}' http://localhost:8080/transactions/new | jq

curl -sSL -X POST http://localhost:8080/mine | jq
curl -sSL -X GET http://localhost:8080/chain | jq


curl -sSL -X POST http://localhost:8081/consensus | jq
curl -sSL -X POST http://localhost:8082/consensus | jq

