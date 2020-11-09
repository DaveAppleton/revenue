package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/DaveAppleton/etherUtils"
	"github.com/DaveAppleton/memorykeys"
	"github.com/DaveAppleton/revenue/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	baseClient  *backends.SimulatedBackend
	srtContract *contracts.StakedRevenueToken
	srtAddress  common.Address
	userMap     map[string]common.Address
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

func stats() {
	stats, err := srtContract.Stats(nil)
	chkErr(err)
	fmt.Println(
		"numTokens  :", stats.NumTokens,
		"Pot  :", etherUtils.EtherToStr(stats.Pot),
		"PotPerToken  :", etherUtils.EtherToStr(stats.PotPerToken),
	)
}

func balance(addr *common.Address) {
	bal, err := baseClient.BalanceAt(context.Background(), *addr, nil)
	chkErr(err)
	balData, err := srtContract.Balance(nil, *addr)
	chkErr(err)
	fmt.Println(
		"bal:", etherUtils.EtherToStr(bal),
		",tokens:", etherUtils.EtherToStr(balData.Tokens),
		",accumulated:", etherUtils.EtherToStr(balData.AccumulatedReward),
		",lastPotPerToken:", etherUtils.EtherToStr(balData.LastPotPerToken),
	)
}

func contractBalance() {
	bal, err := baseClient.BalanceAt(context.Background(), srtAddress, nil)
	chkErr(err)
	fmt.Println(
		"Contract balance",
		etherUtils.EtherToStr(bal),
	)
}

func fundUsers(data [][]string, client *backends.SimulatedBackend) error {
	fmt.Println("funding accounts")
	userMap = make(map[string]common.Address)
	for _, line := range data {
		if _, ok := userMap[line[0]]; !ok && line[0] != "banker" {
			fmt.Println(line[0])
			addr, err := memorykeys.GetAddress(line[0])
			if err != nil {
				return err
			}
			userMap[line[0]] = *addr
			tx, err := sendEthereum("banker", line[0], etherUtils.OneEther(), 25000, client)
			checkTxErr(tx, err)
		}
	}
	return nil
}

func main() {
	var tx *types.Transaction
	var ok bool
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
	input := pflag.String("input", "", "input file")
	pflag.Parse()
	if len(*input) == 0 {
		chkErr(errors.New("no input file specified"))
	}
	bankerTx, _ := memorykeys.GetTransactor("banker")

	data, err := readCSV(*input)
	chkErr(err)

	err = fundUsers(data, client)
	chkErr(err)

	srtAddress, tx, srtContract, err = contracts.DeployStakedRevenueToken(bankerTx, client, big.NewInt(1))
	checkTxErr(tx, err)
	fmt.Println("Contract deployed at ", srtAddress.Hex(), "in tx ", tx.Hash().Hex())

	for pos, line := range data {
		Xtx, err := memorykeys.GetTransactor(line[0])
		chkErr(err)
		fmt.Println(pos+1, line)
		switch line[1] {
		case "profit":
			bankerTx.Value, ok = etherUtils.StrToEther(line[2])
			if !ok {
				chkErr(fmt.Errorf("error in number at line %d (%s)", pos, line[2]))
			}
			chkErr(err)
			tx, err = srtContract.Receive(bankerTx)
			checkTxErr(tx, err)
			stats, err := srtContract.Stats(nil)
			chkErr(err)
			fmt.Println(
				"numTokens  :", stats.NumTokens,
				"Pot  :", etherUtils.EtherToStr(stats.Pot),
				"PotPerToken  :", etherUtils.EtherToStr(stats.PotPerToken),
			)

		case "deposit":
			amount, ok := new(big.Int).SetString(line[2], 10)
			if !ok {
				chkErr(fmt.Errorf("error in number at line %d (%s)", pos, line[2]))
			}
			tx, err = srtContract.DepositTokens(Xtx, amount)
			checkTxErr(tx, err)
			stats, err := srtContract.Stats(nil)
			chkErr(err)
			fmt.Println(
				"numTokens  :", stats.NumTokens,
				"Pot  :", etherUtils.EtherToStr(stats.Pot),
				"PotPerToken  :", etherUtils.EtherToStr(stats.PotPerToken),
			)
		case "wdtoken":
			addr := userMap[line[0]]
			fmt.Print("before ")
			balance(&addr)
			tx, err = srtContract.WithdrawTokens(Xtx)
			checkTxErr(tx, err)
			fmt.Print("after  ")
			balance(&addr)
		case "wdprofit":
			addr := userMap[line[0]]
			fmt.Print("before ")
			balance(&addr)
			tx, err = srtContract.WithDrawRewards(Xtx)
			checkTxErr(tx, err)
			fmt.Print("after  ")
			balance(&addr)
		case "pending":
			addr := userMap[line[0]]
			bbb, err := srtContract.PendingReward(nil, addr)
			chkErr(err)
			fmt.Println("pending", etherUtils.EtherToStr(bbb))
		case "balance":
			// bal, err := client.BalanceAt(context.Background(), srtAddress, nil)
			// chkErr(err)
			// fmt.Println("Balance at all received = ", etherUtils.EtherToStr(bal))
			addr := userMap[line[0]]
			balance(&addr)
		case "stats":
			stats()
		case "contractbal":
			contractBalance()
		default:
			chkErr(errors.New("illegal operation :" + line[1]))
		}

	}

}
