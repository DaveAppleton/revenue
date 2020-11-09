// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220975785e5cf27386b5bf8f783699435533ba08ea925c0b64e46c3f1b1850815a364736f6c63430007040033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StakedRevenueTokenABI is the input ABI used to generate the binding from.
const StakedRevenueTokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unitToken\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastPotPerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"depositTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_potPerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withDrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// StakedRevenueTokenFuncSigs maps the 4-byte function signature to its string representation.
var StakedRevenueTokenFuncSigs = map[string]string{
	"e3d670d7": "balance(address)",
	"dd49756e": "depositTokens(uint256)",
	"06fdde03": "name()",
	"f40f0f52": "pendingReward(address)",
	"d80528ae": "stats()",
	"696c81b1": "withDrawRewards()",
	"8d8f2adb": "withdrawTokens()",
}

// StakedRevenueTokenBin is the compiled bytecode used for deploying new contracts.
var StakedRevenueTokenBin = "0x60c0604052601560808190527f5374616b6564526576656e7565436f6e7472616374000000000000000000000060a090815261003e9160059190610078565b5034801561004b57600080fd5b506040516108d23803806108d28339818101604052602081101561006e57600080fd5b5051600455610119565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826100ae57600085556100f4565b82601f106100c757805160ff19168380011785556100f4565b828001600101855582156100f4579182015b828111156100f45782518255916020019190600101906100d9565b50610100929150610104565b5090565b5b808211156101005760008155600101610105565b6107aa806101286000396000f3fe6080604052600436106100745760003560e01c8063d80528ae1161004e578063d80528ae146101b4578063dd49756e146101e7578063e3d670d714610211578063f40f0f5214610244576100fb565b806306fdde0314610100578063696c81b11461018a5780638d8f2adb1461019f576100fb565b366100fb576002546100869034610289565b600255600154156100e85760006100a86003543461028990919063ffffffff16565b905060006003819055506100df6100d66001546100d0600454856102ec90919063ffffffff16565b90610345565b60005490610289565b600055506100f9565b6003546100f59034610289565b6003555b005b600080fd5b34801561010c57600080fd5b50610115610387565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014f578181015183820152602001610137565b50505050905090810190601f16801561017c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019657600080fd5b506100f9610415565b3480156101ab57600080fd5b506100f9610421565b3480156101c057600080fd5b506101c961042b565b60408051938452602084019290925282820152519081900360600190f35b3480156101f357600080fd5b506100f96004803603602081101561020a57600080fd5b5035610439565b34801561021d57600080fd5b506101c96004803603602081101561023457600080fd5b50356001600160a01b031661047a565b34801561025057600080fd5b506102776004803603602081101561026757600080fd5b50356001600160a01b03166104cf565b60408051918252519081900360200190f35b6000828201838110156102e3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b90505b92915050565b6000826102fb575060006102e6565b8282028284828161030857fe5b04146102e35760405162461bcd60e51b81526004018080602001828103825260218152602001806107546021913960400191505060405180910390fd5b60006102e383836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610517565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561040d5780601f106103e25761010080835404028352916020019161040d565b820191906000526020600020905b8154815290600101906020018083116103f057829003601f168201915b505050505081565b61041f60006105b9565b565b61041f60016105b9565b600054600154600254909192565b610441610646565b60015461044e9082610289565b600155336000908152600660205260409020805461046c9083610289565b815560005460019091015550565b6000806000610487610732565b505050506001600160a01b0316600090815260066020908152604091829020825160608101845281548082526001830154938201849052600290920154930183905292909190565b6001600160a01b038116600090815260066020526040812060045460018201548354849261050f9290916100d0916105079190610692565b8554906102ec565b949350505050565b600081836105a35760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610568578181015183820152602001610550565b50505050905090810190601f1680156105955780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816105af57fe5b0495945050505050565b6105c1610646565b33600090815260066020526040812060028101805492905590821561060c576105ee3383600001546106d4565b33600090815260066020526040812081815560018101829055600201555b600080546001840155604051339183156108fc02918491818181858888f19350505050158015610640573d6000803e3d6000fd5b50505050565b3360009081526006602052604081206004546001820154835492939261067492916100d09161050791610692565b60028301549091506106869082610289565b82600201819055505050565b60006102e383836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506106d8565b5050565b6000818484111561072a5760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610568578181015183820152602001610550565b505050900390565b6040518060600160405280600081526020016000815260200160008152509056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a26469706673582212209af34a78623f9a0c0a95edf67015dd11ae94221f97395f7fdd8d5330938b4d3164736f6c63430007040033"

// DeployStakedRevenueToken deploys a new Ethereum contract, binding an instance of StakedRevenueToken to it.
func DeployStakedRevenueToken(auth *bind.TransactOpts, backend bind.ContractBackend, unitToken *big.Int) (common.Address, *types.Transaction, *StakedRevenueToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StakedRevenueTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakedRevenueTokenBin), backend, unitToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakedRevenueToken{StakedRevenueTokenCaller: StakedRevenueTokenCaller{contract: contract}, StakedRevenueTokenTransactor: StakedRevenueTokenTransactor{contract: contract}, StakedRevenueTokenFilterer: StakedRevenueTokenFilterer{contract: contract}}, nil
}

// StakedRevenueToken is an auto generated Go binding around an Ethereum contract.
type StakedRevenueToken struct {
	StakedRevenueTokenCaller     // Read-only binding to the contract
	StakedRevenueTokenTransactor // Write-only binding to the contract
	StakedRevenueTokenFilterer   // Log filterer for contract events
}

// StakedRevenueTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedRevenueTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedRevenueTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedRevenueTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedRevenueTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedRevenueTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedRevenueTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedRevenueTokenSession struct {
	Contract     *StakedRevenueToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakedRevenueTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedRevenueTokenCallerSession struct {
	Contract *StakedRevenueTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StakedRevenueTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedRevenueTokenTransactorSession struct {
	Contract     *StakedRevenueTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StakedRevenueTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedRevenueTokenRaw struct {
	Contract *StakedRevenueToken // Generic contract binding to access the raw methods on
}

// StakedRevenueTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedRevenueTokenCallerRaw struct {
	Contract *StakedRevenueTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StakedRevenueTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedRevenueTokenTransactorRaw struct {
	Contract *StakedRevenueTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedRevenueToken creates a new instance of StakedRevenueToken, bound to a specific deployed contract.
func NewStakedRevenueToken(address common.Address, backend bind.ContractBackend) (*StakedRevenueToken, error) {
	contract, err := bindStakedRevenueToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedRevenueToken{StakedRevenueTokenCaller: StakedRevenueTokenCaller{contract: contract}, StakedRevenueTokenTransactor: StakedRevenueTokenTransactor{contract: contract}, StakedRevenueTokenFilterer: StakedRevenueTokenFilterer{contract: contract}}, nil
}

// NewStakedRevenueTokenCaller creates a new read-only instance of StakedRevenueToken, bound to a specific deployed contract.
func NewStakedRevenueTokenCaller(address common.Address, caller bind.ContractCaller) (*StakedRevenueTokenCaller, error) {
	contract, err := bindStakedRevenueToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedRevenueTokenCaller{contract: contract}, nil
}

// NewStakedRevenueTokenTransactor creates a new write-only instance of StakedRevenueToken, bound to a specific deployed contract.
func NewStakedRevenueTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedRevenueTokenTransactor, error) {
	contract, err := bindStakedRevenueToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedRevenueTokenTransactor{contract: contract}, nil
}

// NewStakedRevenueTokenFilterer creates a new log filterer instance of StakedRevenueToken, bound to a specific deployed contract.
func NewStakedRevenueTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedRevenueTokenFilterer, error) {
	contract, err := bindStakedRevenueToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedRevenueTokenFilterer{contract: contract}, nil
}

// bindStakedRevenueToken binds a generic wrapper to an already deployed contract.
func bindStakedRevenueToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakedRevenueTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedRevenueToken *StakedRevenueTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedRevenueToken.Contract.StakedRevenueTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedRevenueToken *StakedRevenueTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.StakedRevenueTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedRevenueToken *StakedRevenueTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.StakedRevenueTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedRevenueToken *StakedRevenueTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedRevenueToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedRevenueToken *StakedRevenueTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedRevenueToken *StakedRevenueTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address account) view returns(uint256 tokens, uint256 lastPotPerToken, uint256 accumulatedReward)
func (_StakedRevenueToken *StakedRevenueTokenCaller) Balance(opts *bind.CallOpts, account common.Address) (struct {
	Tokens            *big.Int
	LastPotPerToken   *big.Int
	AccumulatedReward *big.Int
}, error) {
	var out []interface{}
	err := _StakedRevenueToken.contract.Call(opts, &out, "balance", account)

	outstruct := new(struct {
		Tokens            *big.Int
		LastPotPerToken   *big.Int
		AccumulatedReward *big.Int
	})

	outstruct.Tokens = out[0].(*big.Int)
	outstruct.LastPotPerToken = out[1].(*big.Int)
	outstruct.AccumulatedReward = out[2].(*big.Int)

	return *outstruct, err

}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address account) view returns(uint256 tokens, uint256 lastPotPerToken, uint256 accumulatedReward)
func (_StakedRevenueToken *StakedRevenueTokenSession) Balance(account common.Address) (struct {
	Tokens            *big.Int
	LastPotPerToken   *big.Int
	AccumulatedReward *big.Int
}, error) {
	return _StakedRevenueToken.Contract.Balance(&_StakedRevenueToken.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0xe3d670d7.
//
// Solidity: function balance(address account) view returns(uint256 tokens, uint256 lastPotPerToken, uint256 accumulatedReward)
func (_StakedRevenueToken *StakedRevenueTokenCallerSession) Balance(account common.Address) (struct {
	Tokens            *big.Int
	LastPotPerToken   *big.Int
	AccumulatedReward *big.Int
}, error) {
	return _StakedRevenueToken.Contract.Balance(&_StakedRevenueToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedRevenueToken *StakedRevenueTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedRevenueToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedRevenueToken *StakedRevenueTokenSession) Name() (string, error) {
	return _StakedRevenueToken.Contract.Name(&_StakedRevenueToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedRevenueToken *StakedRevenueTokenCallerSession) Name() (string, error) {
	return _StakedRevenueToken.Contract.Name(&_StakedRevenueToken.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address account) view returns(uint256)
func (_StakedRevenueToken *StakedRevenueTokenCaller) PendingReward(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedRevenueToken.contract.Call(opts, &out, "pendingReward", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address account) view returns(uint256)
func (_StakedRevenueToken *StakedRevenueTokenSession) PendingReward(account common.Address) (*big.Int, error) {
	return _StakedRevenueToken.Contract.PendingReward(&_StakedRevenueToken.CallOpts, account)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address account) view returns(uint256)
func (_StakedRevenueToken *StakedRevenueTokenCallerSession) PendingReward(account common.Address) (*big.Int, error) {
	return _StakedRevenueToken.Contract.PendingReward(&_StakedRevenueToken.CallOpts, account)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns(uint256 _potPerToken, uint256 _numTokens, uint256 _pot)
func (_StakedRevenueToken *StakedRevenueTokenCaller) Stats(opts *bind.CallOpts) (struct {
	PotPerToken *big.Int
	NumTokens   *big.Int
	Pot         *big.Int
}, error) {
	var out []interface{}
	err := _StakedRevenueToken.contract.Call(opts, &out, "stats")

	outstruct := new(struct {
		PotPerToken *big.Int
		NumTokens   *big.Int
		Pot         *big.Int
	})

	outstruct.PotPerToken = out[0].(*big.Int)
	outstruct.NumTokens = out[1].(*big.Int)
	outstruct.Pot = out[2].(*big.Int)

	return *outstruct, err

}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns(uint256 _potPerToken, uint256 _numTokens, uint256 _pot)
func (_StakedRevenueToken *StakedRevenueTokenSession) Stats() (struct {
	PotPerToken *big.Int
	NumTokens   *big.Int
	Pot         *big.Int
}, error) {
	return _StakedRevenueToken.Contract.Stats(&_StakedRevenueToken.CallOpts)
}

// Stats is a free data retrieval call binding the contract method 0xd80528ae.
//
// Solidity: function stats() view returns(uint256 _potPerToken, uint256 _numTokens, uint256 _pot)
func (_StakedRevenueToken *StakedRevenueTokenCallerSession) Stats() (struct {
	PotPerToken *big.Int
	NumTokens   *big.Int
	Pot         *big.Int
}, error) {
	return _StakedRevenueToken.Contract.Stats(&_StakedRevenueToken.CallOpts)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 tokens) returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactor) DepositTokens(opts *bind.TransactOpts, tokens *big.Int) (*types.Transaction, error) {
	return _StakedRevenueToken.contract.Transact(opts, "depositTokens", tokens)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 tokens) returns()
func (_StakedRevenueToken *StakedRevenueTokenSession) DepositTokens(tokens *big.Int) (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.DepositTokens(&_StakedRevenueToken.TransactOpts, tokens)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 tokens) returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactorSession) DepositTokens(tokens *big.Int) (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.DepositTokens(&_StakedRevenueToken.TransactOpts, tokens)
}

// WithDrawRewards is a paid mutator transaction binding the contract method 0x696c81b1.
//
// Solidity: function withDrawRewards() returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactor) WithDrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedRevenueToken.contract.Transact(opts, "withDrawRewards")
}

// WithDrawRewards is a paid mutator transaction binding the contract method 0x696c81b1.
//
// Solidity: function withDrawRewards() returns()
func (_StakedRevenueToken *StakedRevenueTokenSession) WithDrawRewards() (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.WithDrawRewards(&_StakedRevenueToken.TransactOpts)
}

// WithDrawRewards is a paid mutator transaction binding the contract method 0x696c81b1.
//
// Solidity: function withDrawRewards() returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactorSession) WithDrawRewards() (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.WithDrawRewards(&_StakedRevenueToken.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactor) WithdrawTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedRevenueToken.contract.Transact(opts, "withdrawTokens")
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_StakedRevenueToken *StakedRevenueTokenSession) WithdrawTokens() (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.WithdrawTokens(&_StakedRevenueToken.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactorSession) WithdrawTokens() (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.WithdrawTokens(&_StakedRevenueToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedRevenueToken.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakedRevenueToken *StakedRevenueTokenSession) Receive() (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.Receive(&_StakedRevenueToken.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakedRevenueToken *StakedRevenueTokenTransactorSession) Receive() (*types.Transaction, error) {
	return _StakedRevenueToken.Contract.Receive(&_StakedRevenueToken.TransactOpts)
}
