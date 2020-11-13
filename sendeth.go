package main

import (
	"context"
	"math/big"

	"github.com/DaveAppleton/memorykeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func sendEthereum(sender string, recipientStr string, amount *big.Int, gasLimit uint64, client *backends.SimulatedBackend) (*types.Transaction, error) {

	senderAddress, _ := memorykeys.GetAddress(sender)
	senderPrivateKey, _ := memorykeys.GetPrivateKey(sender)
	recipient, _ := memorykeys.GetAddress(recipientStr)

	nonce, err := client.PendingNonceAt(context.TODO(), *senderAddress)
	//gasPrice := big.NewInt(gasPrice)
	//fmt.Println("Nonce : ", nonce)
	s := types.HomesteadSigner{}

	data := common.FromHex("0x")
	gasPrice := big.NewInt(1000000000)
	t := types.NewTransaction(nonce, *recipient, amount, gasLimit, gasPrice, data)
	nt, err := types.SignTx(t, s, senderPrivateKey)
	if err != nil {
		return nil, err
	}
	err = client.SendTransaction(context.TODO(), nt)
	return nt, err
}
