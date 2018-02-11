# Go implementation of the Blockchain - Proof of Work

This is a basic implementation of the blockchain written in Go. This
was done for the purposes of learning Go, learning how the internals of
blockchain technology works as well as holding a tech share
presentation for the company I work at. Because of these reasons, this
implementation is not fully secured and has a long way to go to be a
full and usable implementation.

## Building
You can build this app like you would any other go application.

You can resolve dependencies using [glide](https://glide.sh/)
```bash
$ glide install
$ go build -i -o ./build/blockchain-go main.go
```

## Running
You can run multiple nodes (or services) simply by running the built
executable with the server argument and a different port for each
service. If no port is specified, it will use 8080.

```bash
$ ./build/blockchain-go server --port <port>
```

#### Commands
You can run below commands to get a working network set up

First of all, you must run the service 2 or more times:
```
$ ./build/blockchain-go server --port 8080
$ ./build/blockchain-go server --port 8081
```

You can then register each node with all other nodes.
```
curl -sSL -X POST -d '["localhost:8081"]' http://localhost:8080/nodes/register
curl -sSL -X POST -d '["localhost:8080"]' http://localhost:8081/nodes/register
```

You can create new transactions.
```
$ curl -sSL -X POST -d '{"sender":"Hosh","recipient":"Blockchain Foundation","amount":1.5}' http://localhost:8080/transactions/new
```

You can then mine blocks.
```
$ curl -sSL -X POST http://localhost:8080/mine
```

You can retrieve the full chain by called the `GET /chain` endpoint.
```
$ curl -sSL -X GET http://localhost:8080/chain
```

You can gain consensus between nodes by calling the `POST /consensus`
endpoint
```
$ curl -sSL -X POST http://localhost:8081/consensus
```

## Caveats / limitations
As you can see there's a lot of manual work being done (node
registration, reaching consensus) and additionally.

This is obviously not very manageable especially not for non-technical
users of the cryptocurrency. As mentioned before this is not a full
blockchain implementation and amongst others the following is missing:

* Automatic node discovery
  Of course it does not make sense to manually go to everyone and tell
  your node where to find each individual node. This will most certainly
  lead to divergence of the chain.
* Unconfirmed transactions sharing.
  Everyone must be aware of all unconfirmed transaction
* It needs to be able automatically tell all nodes that a new block has
  been found. This can then get all other nodes to stop their current
  mining process (if they are mining) so they can try and mine the next
  block instead. Additionally, it prevents a transaction from being
  mined twice as other nodes can discard those transactions from their
  list of transactions to add to the next block.
* Automatic consensus
  This is similar to the previous point, but differs in that, when a new
  node joins the network, it must be able to find the correct chain to
  prevent an accidental chain split (or even start a whole new chain on
  its own).
* Security
  No thought has gone into the security of transactions and blocks
  besides ensuring that the chain cannot be manipulated, however, a lot
  more security is required. For example we need to be able to verify
  a transaction is actually from the sender (done through public and
  and private keys). In general the idea of a wallet doesn't exist in
  this implementation and therefore everything related to wallets such
  does as preventing double spend not either.

There's probably other things I'm missing that would need to be
implemented for this to be a viable cryptocurrency, however, as this
was a learning exercise to learn the core mechanics of what makes a
blockchain implementation tick, I will likely not be implementing those.
I will however, accept pull requests if you would want to implement
these. Additionally, if I have gotten anything wrong (and I hope I
haven't), please feel free to either open an issue or a PR to correct
me.
