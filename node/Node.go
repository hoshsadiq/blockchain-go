package node

import (
    "fmt"
    "net"
)

type Node struct {
    Host string
    Port string
}

func (node *Node) GetChainUrl() string {
    return fmt.Sprintf("http://%s:%d/chain", node.Host, node.Port)
}

func (node *Node) UnmarshalText(data []byte) error {
    host, port, err := net.SplitHostPort(string(data))

    if err != nil {
        return err
    }

    node.Host = host
    node.Port = port

    return nil
}

func New(host string, port string) (*Node) {
    return &Node{
        Host: host,
        Port: port,
    }
}
