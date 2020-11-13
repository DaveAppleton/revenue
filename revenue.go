package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

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
	baseClient      *backends.SimulatedBackend
	stakingContract *contracts.XStaking
	stakingAddress  common.Address
	userMap         map[string]common.Address
	lotoAddress     common.Address
	lotoToken       *contracts.StakingToken
	wethContract    *contracts.Weth
	wethAddress     common.Address
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
		//fmt.Printf("%+v\n", rctp)
		chkErr(fmt.Errorf("transaction returns %d", rctp.Status))

	}
	head, err := baseClient.HeaderByHash(context.Background(), rctp.BlockHash)
	chkErr(err)
	fmt.Println("Gas Used ", rctp.GasUsed, "timestamp", head.Time)
}

func checkTxErr(tx *types.Transaction, err error) {
	checkTxErrB(tx, err, types.ReceiptStatusSuccessful)
}

func checkTxErrF(tx *types.Transaction, err error) {
	checkTxErrB(tx, err, types.ReceiptStatusFailed)
}

// func stats() {
// 	stats, err := stakingContract.Stats(nil)
// 	chkErr(err)
// 	fmt.Println(
// 		"numTokens  :", stats.NumTokens,
// 		"Pot  :", etherUtils.EtherToStr(stats.Pot),
// 		"PotPerToken  :", etherUtils.EtherToStr(stats.PotPerToken),
// 	)
// }

func balance(addr *common.Address) {
	bal, err := wethContract.BalanceOf(nil, *addr)
	chkErr(err)
	earned, err := stakingContract.Earned(nil, *addr)
	fmt.Println(
		"bal:", etherUtils.EtherToStr(bal),
		"earned:", etherUtils.EtherToStr(earned),
	)
}

func stats() {
	rpt, err := stakingContract.RewardPerToken(nil)
	chkErr(err)
	bal, err := wethContract.BalanceOf(nil, stakingAddress)
	chkErr(err)

	lrta, err := stakingContract.LastTimeRewardApplicable(nil)
	chkErr(err)

	pf, err := stakingContract.PeriodFinish(nil)
	chkErr(err)

	rd, err := stakingContract.RewardsDuration(nil)
	chkErr(err)

	// rt, err := stakingContract.RewardsToken(nil)
	// chkErr(err)

	fmt.Println(
		"Rewards Per Token", etherUtils.EtherToStr(rpt),
		"WETH balance", etherUtils.EtherToStr(bal),
		"LastTimeRewardAppliccable ", lrta,
		"periodFinish", pf,
		"rewardsDuration", rd,
		//"rewardsToken", rt.Hex(),
	)
}

func contractBalance() {
	wbal, err := wethContract.BalanceOf(nil, stakingAddress)
	chkErr(err)
	sbal, err := lotoToken.BalanceOf(nil, stakingAddress)
	chkErr(err)

	fmt.Println(
		"Contract WETH", etherUtils.EtherToStr(wbal),
		"Contract LOTO", etherUtils.EtherToStr(sbal),
	)
}

func fundUsers(data [][]string, client *backends.SimulatedBackend) error {
	fmt.Println("funding accounts")
	bigAmount := new(big.Int).Mul(etherUtils.OneEther(), big.NewInt(100))
	userMap = make(map[string]common.Address)
	for _, line := range data {
		if _, ok := userMap[line[0]]; !ok && line[0] != "banker" {
			fmt.Println(line[0])
			addr, err := memorykeys.GetAddress(line[0])
			if err != nil {
				return err
			}
			bankerTx, _ := memorykeys.GetTransactor("banker")
			userMap[line[0]] = *addr
			tx, err := sendEthereum("banker", line[0], etherUtils.OneEther(), 25000, client)
			checkTxErr(tx, err)
			tx, err = lotoToken.Transfer(bankerTx, *addr, bigAmount)
			checkTxErr(tx, err)
		}
	}
	bad, err := memorykeys.GetAddress("banker")
	chkErr(err)
	userMap["banker"] = *bad
	return nil
}

func identify(addr common.Address) string {
	for key, val := range userMap {
		if addr == val {
			return key
		}
	}
	return addr.Hex()
}

// func endPeriod(client *backends.SimulatedBackend) {
// 	end, err := stakingContract.PeriodFinish(nil)
// 	chkErr(err)
// 	hdr, err := client.HeaderByNumber(context.Background(), nil)
// 	chkErr(err)
// 	jump := end.Uint64() - hdr.Time
// 	err = client.AdjustTime(time.Duration(jump) * time.Second)
// 	chkErr(err)
// 	client.Commit()
// 	endHdr, err := client.HeaderByNumber(context.Background(), nil)
// 	chkErr(err)

// 	fmt.Println("start", hdr.Time, "periodEnd", end.Uint64(), "jump", jump, "final", endHdr.Time)

// }

func endPeriod(client *backends.SimulatedBackend, qualifier string) {
	period, err := stakingContract.RewardsDuration(nil)
	chkErr(err)
	end, err := stakingContract.PeriodFinish(nil)
	chkErr(err)
	hdr, err := client.HeaderByNumber(context.Background(), nil)
	chkErr(err)
	jump := period.Uint64()
	if end.Uint64() > hdr.Time {
		jump = end.Uint64() - hdr.Time
	}
	switch qualifier {
	case "long":
		jump = period.Uint64() + jump
	case "short":
		jump -= 1000
	case "day":
		jump = 24 * 60 * 60
	default:
	}
	fmt.Println(jump)
	err = client.AdjustTime(time.Duration(jump) * time.Second)
	chkErr(err)
	client.Commit()
	endHdr, err := client.HeaderByNumber(context.Background(), nil)
	chkErr(err)

	fmt.Println("start", hdr.Time, "periodEnd", end.Uint64(), "jump", jump, "final", endHdr.Time)

}

func allBalances() {
	fmt.Println("User Balances")
	for name, addr := range userMap {
		lotoBal, err := lotoToken.BalanceOf(nil, addr)
		chkErr(err)
		ethBal, err := wethContract.BalanceOf(nil, addr)
		chkErr(err)
		fmt.Printf("%s : loto : %s, weth : %s\n", name, etherUtils.EtherToStr(lotoBal), etherUtils.EtherToStr(ethBal))
	}
}

func main() {
	var tx *types.Transaction
	// var ok bool
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
		pflag.Usage()
		chkErr(errors.New("no input file specified"))
	}
	bankerTx, _ := memorykeys.GetTransactor("banker")
	bankerAddress, _ := memorykeys.GetAddress("banker")

	data, err := readCSV(*input)
	chkErr(err)

	wethAddress, tx, wethContract, err = contracts.DeployWeth(bankerTx, client)
	checkTxErr(tx, err)
	lotoAddress, tx, lotoToken, err = contracts.DeployStakingToken(bankerTx, client)
	checkTxErr(tx, err)
	stakingAddress, tx, stakingContract, err = contracts.DeployXStaking(bankerTx, client, *bankerAddress, *bankerAddress, wethAddress, lotoAddress)
	checkTxErr(tx, err)
	fmt.Println("Contract deployed at ", stakingAddress.Hex(), "in tx ", tx.Hash().Hex())

	err = fundUsers(data, client)
	chkErr(err)
	userMap["staking"] = stakingAddress
	rda, err := stakingContract.RewardsDistribution(nil)
	chkErr(err)

	fmt.Println("RewardsDistribution :", identify(rda))

	for pos, line := range data {
		Xtx, err := memorykeys.GetTransactor(line[0])
		chkErr(err)
		fmt.Println(pos+1, line)
		switch line[1] {
		case "profit":
			end, err := stakingContract.PeriodFinish(nil)
			chkErr(err)
			hdr, err := client.HeaderByNumber(context.Background(), nil)
			chkErr(err)
			if hdr.Time > end.Uint64() {
				fmt.Println("+++++++ PERIOD FINISHED ++++++ ")
			}

			reward, ok := etherUtils.StrToEther(line[2])
			if !ok {
				chkErr(fmt.Errorf("error in number at line %d (%s)", pos, line[2]))
			}
			bankerTx.Value = nil
			chkErr(err)
			tx, err = wethContract.Transfer(bankerTx, stakingAddress, reward)
			checkTxErr(tx, err)
			stats()
			bankerTx.GasLimit = 1000000
			tx, err = stakingContract.NotifyRewardAmount(bankerTx, reward)
			checkTxErr(tx, err)
			rr, err := stakingContract.RewardRate(nil)
			chkErr(err)
			fmt.Println("rewardRate", etherUtils.EtherToStr(rr))

		case "deposit":
			amount, ok := etherUtils.StrToEther(line[2])
			if !ok {
				chkErr(fmt.Errorf("error in number at line %d (%s)", pos, line[2]))
			}
			addr := userMap[line[0]]
			bal, err := lotoToken.BalanceOf(nil, addr)
			chkErr(err)
			fmt.Println(line[0], "has", etherUtils.EtherToStr(bal), "LOTO")
			tx, err = lotoToken.Approve(Xtx, stakingAddress, amount)
			checkTxErr(tx, err)
			tx, err = stakingContract.Stake(Xtx, amount)
			checkTxErr(tx, err)
		case "wdprofit":
			addr := userMap[line[0]]
			fmt.Print("before ")
			balance(&addr)
			tx, err = stakingContract.GetReward(Xtx)
			checkTxErr(tx, err)
			fmt.Print("after  ")
			balance(&addr)
		case "wdtoken":
			addr := userMap[line[0]]
			bal, err := lotoToken.BalanceOf(nil, addr)
			chkErr(err)
			fmt.Println(line[0], "has", etherUtils.EtherToStr(bal), "LOTO")
			amount, ok := etherUtils.StrToEther(line[2])
			if !ok {
				chkErr(fmt.Errorf("Bad String %s", line[2]))
			}
			fmt.Print("before ")
			balance(&addr)
			tx, err = stakingContract.Withdraw(Xtx, amount)
			checkTxErr(tx, err)
			fmt.Print("after  ")
			balance(&addr)
		case "pending":
			addr := userMap[line[0]]
			bbb, err := stakingContract.Earned(nil, addr)
			chkErr(err)
			fmt.Println("pending", etherUtils.EtherToStr(bbb))
		case "balance":
			addr := userMap[line[0]]
			balance(&addr)
		case "contractbal":
			contractBalance()
		case "stats":
			stats()
		case "endperiod":
			endPeriod(client, "end")
		case "long":
			endPeriod(client, "long")
		case "short":
			endPeriod(client, "short")
		case "day":
			endPeriod(client, "day")
		case "balances":
			allBalances()
		case "exit":
			tx, err = stakingContract.Exit(Xtx)
			checkTxErr(tx, err)
		default:
			chkErr(errors.New("illegal operation :" + line[1]))
		}

	}

}
