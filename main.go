package main

import (
	"context"
	"fmt"
	ethNodeCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
	"net"
)

func node(opts ...libp2p.Option) {
	fmt.Printf("")
	hostAddr, err := net.ResolveTCPAddr("tcp", "localhost:60006")
	if err != nil {
		panic(err)
	}
	var multiAddresses []ma.Multiaddr
	hostAddrMA, err := manet.FromNetAddr(hostAddr)
	if err != nil {
		panic(err)
	}

	key := "bbf358ab08ab29d70b6b20845e4aa417124bb8051ecdbaf4f822bba18f28f7fb"
	privKey, err := ethNodeCrypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	nodeKey := crypto.PrivKey((*crypto.Secp256k1PrivateKey)(privKey))
	host, err := libp2p.New(
		context.Background(),
		append(opts,
			libp2p.ListenAddrs(append(multiAddresses, hostAddrMA)...),
			libp2p.Identity(nodeKey),
		)...,
	)
	if err != nil {
		panic(err)
	}
	hostInfo, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", host.ID().Pretty()))
	for _, addr := range host.Addrs() {
		fullAddr := addr.Encapsulate(hostInfo)
		fmt.Println("Listening on", fullAddr)
	}

}

func main() {
	node()
}
