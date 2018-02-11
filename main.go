package main

import (
    "os"
    "log"
    "github.com/urfave/cli"
    "github.com/hoshsadiq/blockchain-go/server"
    "github.com/hoshsadiq/blockchain-go/block"
    "github.com/hoshsadiq/blockchain-go/helper"
)

func main() {
    app := cli.NewApp()
    app.Name = "blockchain-go"
    app.Version = "0.0.1"
    app.Usage = "Interact with the Blockchain Go cryptocurrency"

    app.Commands = []cli.Command{
        {
            Name: "server",
            Flags: []cli.Flag{
                cli.IntFlag{
                    Name:  "port, p",
                    Value: 8080,
                },
            },
            Action: func(c *cli.Context) error {
                blockchain := block.NewBlockchain()
                addr := helper.GenerateAddress()

                return server.RunServer(blockchain, addr, c.Int("port"))
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err.Error())
    }
}
