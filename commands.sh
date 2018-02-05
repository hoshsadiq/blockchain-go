#!/usr/bin/env bash

curl -sSL -X POST -d '{"sender":"Lara","recipient":"Hosh","amount":1.5}' http://localhost:8080/transactions/new | jq


