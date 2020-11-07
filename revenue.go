package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/memorykeys"
	"github.com/DaveAppleton/revenue/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	baseClient *backends.SimulatedBackend
)

func getClient() (client *backends.SimulatedBackend, err error) {
	if baseClient != nil {
		return baseClient, nil
	}
	gaslimit := viper.GetUint64("gaslimit")
	funds, _ := new(big.Int).SetString("1000000000000000000000000000", 10)
	banker, _ := memorykeys.GetAddress("banker")
	baseClient = backends.NewSimulatedBackend(core.GenesisAlloc{
		*banker: {Balance: funds},
	}, gaslimit)

	return baseClient, nil
}

func initViper() {
	viper.SetConfigFile("./config.json")
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("./config.json")
		log.Fatal("config.json", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("viper config changed", e.Name)
	})
}

func readCSV(input string) (data [][]string, err error) {
	sf, err := os.Open(input)
	if err != nil {
		fmt.Println(err, "[", input, "]")
		log.Fatal(err, input)
	}
	data, err = csv.NewReader(sf).ReadAll()
	return
}

func chkErr(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func checkTxErrB(tx *types.Transaction, err error, expected uint64) {
	chkErr(err)
	baseClient.Commit()
	rctp, err := baseClient.TransactionReceipt(context.Background(), tx.Hash())
	chkErr(err)
	if rctp.Status != expected {
		chkErr(fmt.Errorf("transaction returns %d", rctp.Status))
	}
	fmt.Println("Gas Used ", rctp.GasUsed)
}

func checkTxErr(tx *types.Transaction, err error) {
	checkTxErrB(tx, err, types.ReceiptStatusSuccessful)
}

func checkTxErrF(tx *types.Transaction, err error) {
	checkTxErrB(tx, err, types.ReceiptStatusFailed)
}

func main() {
	initViper()
	logName := viper.GetString("log")
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})
	client, err := getClient()
	chkErr(err)

	bankerTx, _ := memorykeys.GetTransactor("banker")
	txA, _ := memorykeys.GetTransactor("userA")
	txB, _ := memorykeys.GetTransactor("userB")
	txC, _ := memorykeys.GetTransactor("userC")
	tx, err := sendEthereum("banker", "userA", etherUtils.OneEther(), 25000, client)
	checkTxErr(tx, err)
	tx, err = sendEthereum("banker", "userB", etherUtils.OneEther(), 25000, client)
	checkTxErr(tx, err)
	tx, err = sendEthereum("banker", "userC", etherUtils.OneEther(), 25000, client)
	checkTxErr(tx, err)
	// privA, _ := memorykeys.GetPrivateKey("userA")
	// privB, _ := memorykeys.GetPrivateKey("userB")
	// privC, _ := memorykeys.GetPrivateKey("userC")
	// userD, _ := memorykeys.GetAddress("userD")

	srtAddress, tx, srtContract, err := contracts.DeployStakedRevenueToken(bankerTx, client)
	checkTxErr(tx, err)
	fmt.Println("Contract deployed at ", srtAddress.Hex(), "in tx ", tx.Hash().Hex())

	for pos, Xtx := range []*bind.TransactOpts{bankerTx, txA, txB, txC, bankerTx} {
		users := []string{"banker", "userA", "userB", "userC", "banker"}
		fmt.Println(users[pos])

		stats, err := srtContract.Stats(nil)
		chkErr(err)
		fmt.Println(
			etherUtils.EtherToStr(stats.NumTokens),
			etherUtils.EtherToStr(stats.Pot),
			etherUtils.EtherToStr(stats.PotPerToken),
		)
		addr, err := memorykeys.GetAddress(users[pos])
		chkErr(err)
		balData, err := srtContract.Balance(nil, *addr)
		chkErr(err)
		fmt.Println(
			etherUtils.EtherToStr(balData.Tokens),
			etherUtils.EtherToStr(balData.AccumulatedReward),
			etherUtils.EtherToStr(balData.LastPotPerToken),
		)

		tx, err = srtContract.DepositTokens(Xtx, etherUtils.OneEther())
		checkTxErr(tx, err)
		bankerTx.Value = etherUtils.OneEther()
		tx, err = srtContract.Receive(bankerTx)
		checkTxErr(tx, err)
	}
	bal, err := client.BalanceAt(context.Background(), srtAddress, nil)
	chkErr(err)
	fmt.Println("Balance at all received = ", etherUtils.EtherToStr(bal))
	stats, err := srtContract.Stats(nil)
	chkErr(err)
	fmt.Println(
		etherUtils.EtherToStr(stats.NumTokens),
		etherUtils.EtherToStr(stats.Pot),
		etherUtils.EtherToStr(stats.PotPerToken),
	)
	for pos, _ := range []*bind.TransactOpts{txC, txB, txA, bankerTx} {
		users := []string{"userC", "userB", "userA", "banker"}
		fmt.Println(users[pos])
		addr, err := memorykeys.GetAddress(users[pos])
		chkErr(err)
		balData, err := srtContract.Balance(nil, *addr)
		chkErr(err)
		fmt.Println(
			etherUtils.EtherToStr(balData.Tokens),
			etherUtils.EtherToStr(balData.AccumulatedReward),
			etherUtils.EtherToStr(balData.LastPotPerToken),
		)
		// Xtx.Value = nil
		// tx, err = srtContract.WithDrawRewards(Xtx)
		// checkTxErr(tx, err)
	}
}
