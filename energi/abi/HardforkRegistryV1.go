// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	ethereum "energi.world/core/gen3"
	"energi.world/core/gen3/accounts/abi"
	"energi.world/core/gen3/accounts/abi/bind"
	"energi.world/core/gen3/common"
	"energi.world/core/gen3/core/types"
	"energi.world/core/gen3/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HardforkRegistryV1ABI is the input ABI used to generate the binding from.
const HardforkRegistryV1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_HF_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_HF_finalization_period\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"Hardfork\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"HF_signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enumerate\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"}],\"name\":\"getByBlockNo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"}],\"name\":\"getByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIGovernedContract\",\"name\":\"_oldImpl\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"block_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sw_features\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"block_no\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"v1storage\",\"outputs\":[{\"internalType\":\"contractStorageHardforkRegistryV1\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HardforkRegistryV1Bin is the compiled bytecode used for deploying new contracts.
const HardforkRegistryV1Bin = `608060405234801561001057600080fd5b5060405161194f38038061194f8339818101604052606081101561003357600080fd5b5080516020820151604092830151600080546001600160a01b0319166001600160a01b03851617905592519192909161006b906100c1565b604051809103906000f080158015610087573d6000803e3d6000fd5b50600380546001600160a01b039283166001600160a01b0319918216179091556002805494909216931692909217909155600455506100ce565b6108308061111f83390190565b611042806100dd6000396000f3fe6080604052600436106100b05760003560e01c80634cc8221511610069578063ce5494bb1161004e578063ce5494bb14610284578063ec556889146102c4578063ff9f78b3146102d9576100b0565b80634cc82215146102305780638bc237f11461025a576100b0565b80631658312e1161009a5780631658312e146101955780631894d2ea146101dd5780632d0593051461021b576100b0565b8062f55d9d14610117578063072a982314610159575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561012357600080fd5b506101576004803603602081101561013a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661033e565b005b34801561016557600080fd5b506101576004803603608081101561017c57600080fd5b50803590602081013590604081013590606001356103e6565b3480156101a157600080fd5b506101bf600480360360208110156101b857600080fd5b503561087f565b60408051938452602084019290925282820152519081900360600190f35b3480156101e957600080fd5b506101f26108b6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561022757600080fd5b506101f26108d2565b34801561023c57600080fd5b506101576004803603602081101561025357600080fd5b50356108ee565b34801561026657600080fd5b506101bf6004803603602081101561027d57600080fd5b5035610b19565b34801561029057600080fd5b50610157600480360360208110156102a757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bef565b3480156102d057600080fd5b506101f2610c7d565b3480156102e557600080fd5b506102ee610c99565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561032a578181015183820152602001610312565b505050509050019250505060405180910390f35b60005473ffffffffffffffffffffffffffffffffffffffff1633146103c457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6103cd81610df1565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6001541561045557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001805560025473ffffffffffffffffffffffffffffffffffffffff1661047a610e80565b73ffffffffffffffffffffffffffffffffffffffff16146104fc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f496e76616c69642068617264666f726b207369676e65722063616c6c65720000604482015290519081900360640190fd5b8261056857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f48617264666f726b206e616d652063616e6e6f7420626520656d707479000000604482015290519081900360640190fd5b600354604080517f7f1a5c6200000000000000000000000000000000000000000000000000000000815260048101869052905160009273ffffffffffffffffffffffffffffffffffffffff1691637f1a5c62916024808301926020929190829003018186803b1580156105da57600080fd5b505afa1580156105ee573d6000803e3d6000fd5b505050506040513d602081101561060457600080fd5b50519050801561072657848114610666576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180610f4b6028913960400191505060405180910390fd5b43600454860110156106c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610fbd6027913960400191505060405180910390fd5b438510156107215782610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610fe4602a913960400191505060405180910390fd5b61077f565b4385101561077f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610f736026913960400191505060405180910390fd5b600354604080517fbad612cd00000000000000000000000000000000000000000000000000000000815260048101889052602481018690526044810187905260648101859052905173ffffffffffffffffffffffffffffffffffffffff9092169163bad612cd9160848082019260009290919082900301818387803b15801561080757600080fd5b505af115801561081b573d6000803e3d6000fd5b50505084158015915061082d57508215155b1561087357604080518681526020810185905280820186905290517f9e3eb3a1090f7e2eb48f596218f9322ec1584fad2673784a5cbd5f9e452f18b39181900360600190a15b50506000600155505050565b600354600090819081906108a99073ffffffffffffffffffffffffffffffffffffffff1685610eaf565b9196909550909350915050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b6001541561095d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001805560025473ffffffffffffffffffffffffffffffffffffffff16610982610e80565b73ffffffffffffffffffffffffffffffffffffffff1614610a0457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f496e76616c69642068617264666f726b207369676e65722063616c6c65720000604482015290519081900360640190fd5b600354600090610a2a9073ffffffffffffffffffffffffffffffffffffffff1683610eaf565b509150508015610a85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610f996024913960400191505060405180910390fd5b600354604080517fd906415700000000000000000000000000000000000000000000000000000000815260048101859052905173ffffffffffffffffffffffffffffffffffffffff9092169163d90641579160248082019260009290919082900301818387803b158015610af857600080fd5b505af1158015610b0c573d6000803e3d6000fd5b5050600060015550505050565b600354604080517f7f1a5c620000000000000000000000000000000000000000000000000000000081526004810184905290516000928392839273ffffffffffffffffffffffffffffffffffffffff90921691637f1a5c6291602480820192602092909190829003018186803b158015610b9257600080fd5b505afa158015610ba6573d6000803e3d6000fd5b505050506040513d6020811015610bbc57600080fd5b5051600354909350610be49073ffffffffffffffffffffffffffffffffffffffff1684610eaf565b949690955092505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610c7557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610c7a815b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600354604080517f723922e7000000000000000000000000000000000000000000000000000000008152905160609273ffffffffffffffffffffffffffffffffffffffff169163723922e7916004808301926000929190829003018186803b158015610d0457600080fd5b505afa158015610d18573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526020811015610d5f57600080fd5b8101908080516040519392919084640100000000821115610d7f57600080fd5b908301906020820185811115610d9457600080fd5b8251866020820283011164010000000082111715610db157600080fd5b82525081516020918201928201910280838360005b83811015610dde578181015183820152602001610dc6565b5050505090500160405250505090505b90565b600354604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b158015610e6557600080fd5b505af1158015610e79573d6000803e3d6000fd5b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16331415610ea8575032610dee565b5033610dee565b60008060008473ffffffffffffffffffffffffffffffffffffffff1663fce183be856040518263ffffffff1660e01b81526004018082815260200191505060606040518083038186803b158015610f0557600080fd5b505afa158015610f19573d6000803e3d6000fd5b505050506040513d6060811015610f2f57600080fd5b5080516020820151604090920151909791965094509250505056fe4475706c69636174652068617264666f726b206e616d657320617265206e6f7420616c6c6f77656448617264666f726b2063616e6e6f74206265206372656174656420696e20746865207061737446696e616c697a65642068617264666f726b2063616e6e6f742062652064656c6574656448617264666f726b2066696e616c697a6174696f6e20696e74657276616c20657863656564656448462066696e616c697a6174696f6e20626c6f636b20686173682063616e6e6f7420626520656d707479a265627a7a72315820e2d78648acd1c0eaf460724edb81b4230c7ab855137f09a7b6c7cf98d76ce75264736f6c634300051000326080604052600080546001600160a01b0319163317905561080b806100256000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063bad612cd1161005b578063bad612cd14610151578063d906415714610180578063e2d2cf711461019d578063fce183be146101ba57610088565b806313af40351461008d57806341c0e1b5146100c2578063723922e7146100ca5780637f1a5c6214610122575b600080fd5b6100c0600480360360208110156100a357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101f5565b005b6100c06102c2565b6100d261034b565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561010e5781810151838201526020016100f6565b505050509050019250505060405180910390f35b61013f6004803603602081101561013857600080fd5b50356103a3565b60408051918252519081900360200190f35b6100c06004803603608081101561016757600080fd5b50803590602081013590604081013590606001356103b5565b6100c06004803603602081101561019657600080fd5b5035610550565b61013f600480360360208110156101b357600080fd5b5035610778565b6101d7600480360360208110156101d057600080fd5b5035610796565b60408051938452602084019290925282820152519081900360600190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461027b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff16331461034857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b33ff5b6060600180548060200260200160405190810160405280929190818152602001828054801561039957602002820191906000526020600020905b815481526020019060010190808311610385575b5050505050905090565b60036020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff16331461043b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000848152600260205260409020600101548490156104bb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f68617264666f726b206368616e676573206e6f74206564697461626c65000000604482015290519081900360640190fd5b600085815260026020526040902080541580156104d757508315155b15610511576001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6018690555b841561051f57600181018590555b821561052d57600281018390555b83156105485783815560008481526003602052604090208690555b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105d657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f4e6f74206f776e65722100000000000000000000000000000000000000000000604482015290519081900360640190fd5b60008181526002602052604090206001015481901561065657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f68617264666f726b206368616e676573206e6f74206564697461626c65000000604482015290519081900360640190fd5b61065e6107b6565b5060008281526002602081815260408084208151606081018352815480825260018301805483870152838701805484870152918852600386529387208790558887529490935284905583905590829055905b6001548110156107725783600182815481106106c857fe5b9060005260206000200154141561076a57805b6001547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01811015610743576001816001018154811061071757fe5b90600052602060002001546001828154811061072f57fe5b6000918252602090912001556001016106db565b600180548061074e57fe5b6001900381819060005260206000200160009055905550610772565b6001016106b0565b50505050565b6001818154811061078557fe5b600091825260209091200154905081565b600260208190526000918252604090912080546001820154919092015483565b60408051606081018252600080825260208201819052918101919091529056fea265627a7a72315820c6cd3b549bbed63c6b2e654036ec1d2ee48711189356933130320472c233c07d64736f6c63430005100032`

// DeployHardforkRegistryV1 deploys a new Ethereum contract, binding an instance of HardforkRegistryV1 to it.
func DeployHardforkRegistryV1(auth *bind.TransactOpts, backend bind.ContractBackend, _proxy common.Address, _HF_signer common.Address, _HF_finalization_period *big.Int) (common.Address, *types.Transaction, *HardforkRegistryV1, error) {
	parsed, err := abi.JSON(strings.NewReader(HardforkRegistryV1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HardforkRegistryV1Bin), backend, _proxy, _HF_signer, _HF_finalization_period)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HardforkRegistryV1{HardforkRegistryV1Caller: HardforkRegistryV1Caller{contract: contract}, HardforkRegistryV1Transactor: HardforkRegistryV1Transactor{contract: contract}, HardforkRegistryV1Filterer: HardforkRegistryV1Filterer{contract: contract}}, nil
}

// HardforkRegistryV1Bin is the compiled bytecode of contract after deployment.
const HardforkRegistryV1RuntimeBin = `6080604052600436106100b05760003560e01c80634cc8221511610069578063ce5494bb1161004e578063ce5494bb14610284578063ec556889146102c4578063ff9f78b3146102d9576100b0565b80634cc82215146102305780638bc237f11461025a576100b0565b80631658312e1161009a5780631658312e146101955780631894d2ea146101dd5780632d0593051461021b576100b0565b8062f55d9d14610117578063072a982314610159575b604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f4e6f7420737570706f7274656400000000000000000000000000000000000000604482015290519081900360640190fd5b34801561012357600080fd5b506101576004803603602081101561013a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661033e565b005b34801561016557600080fd5b506101576004803603608081101561017c57600080fd5b50803590602081013590604081013590606001356103e6565b3480156101a157600080fd5b506101bf600480360360208110156101b857600080fd5b503561087f565b60408051938452602084019290925282820152519081900360600190f35b3480156101e957600080fd5b506101f26108b6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561022757600080fd5b506101f26108d2565b34801561023c57600080fd5b506101576004803603602081101561025357600080fd5b50356108ee565b34801561026657600080fd5b506101bf6004803603602081101561027d57600080fd5b5035610b19565b34801561029057600080fd5b50610157600480360360208110156102a757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bef565b3480156102d057600080fd5b506101f2610c7d565b3480156102e557600080fd5b506102ee610c99565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561032a578181015183820152602001610312565b505050509050019250505060405180910390f35b60005473ffffffffffffffffffffffffffffffffffffffff1633146103c457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6103cd81610df1565b8073ffffffffffffffffffffffffffffffffffffffff16ff5b6001541561045557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001805560025473ffffffffffffffffffffffffffffffffffffffff1661047a610e80565b73ffffffffffffffffffffffffffffffffffffffff16146104fc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f496e76616c69642068617264666f726b207369676e65722063616c6c65720000604482015290519081900360640190fd5b8261056857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f48617264666f726b206e616d652063616e6e6f7420626520656d707479000000604482015290519081900360640190fd5b600354604080517f7f1a5c6200000000000000000000000000000000000000000000000000000000815260048101869052905160009273ffffffffffffffffffffffffffffffffffffffff1691637f1a5c62916024808301926020929190829003018186803b1580156105da57600080fd5b505afa1580156105ee573d6000803e3d6000fd5b505050506040513d602081101561060457600080fd5b50519050801561072657848114610666576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180610f4b6028913960400191505060405180910390fd5b43600454860110156106c3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610fbd6027913960400191505060405180910390fd5b438510156107215782610721576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610fe4602a913960400191505060405180910390fd5b61077f565b4385101561077f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610f736026913960400191505060405180910390fd5b600354604080517fbad612cd00000000000000000000000000000000000000000000000000000000815260048101889052602481018690526044810187905260648101859052905173ffffffffffffffffffffffffffffffffffffffff9092169163bad612cd9160848082019260009290919082900301818387803b15801561080757600080fd5b505af115801561081b573d6000803e3d6000fd5b50505084158015915061082d57508215155b1561087357604080518681526020810185905280820186905290517f9e3eb3a1090f7e2eb48f596218f9322ec1584fad2673784a5cbd5f9e452f18b39181900360600190a15b50506000600155505050565b600354600090819081906108a99073ffffffffffffffffffffffffffffffffffffffff1685610eaf565b9196909550909350915050565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60035473ffffffffffffffffffffffffffffffffffffffff1681565b6001541561095d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f5265656e74727900000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6001805560025473ffffffffffffffffffffffffffffffffffffffff16610982610e80565b73ffffffffffffffffffffffffffffffffffffffff1614610a0457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f496e76616c69642068617264666f726b207369676e65722063616c6c65720000604482015290519081900360640190fd5b600354600090610a2a9073ffffffffffffffffffffffffffffffffffffffff1683610eaf565b509150508015610a85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610f996024913960400191505060405180910390fd5b600354604080517fd906415700000000000000000000000000000000000000000000000000000000815260048101859052905173ffffffffffffffffffffffffffffffffffffffff9092169163d90641579160248082019260009290919082900301818387803b158015610af857600080fd5b505af1158015610b0c573d6000803e3d6000fd5b5050600060015550505050565b600354604080517f7f1a5c620000000000000000000000000000000000000000000000000000000081526004810184905290516000928392839273ffffffffffffffffffffffffffffffffffffffff90921691637f1a5c6291602480820192602092909190829003018186803b158015610b9257600080fd5b505afa158015610ba6573d6000803e3d6000fd5b505050506040513d6020811015610bbc57600080fd5b5051600354909350610be49073ffffffffffffffffffffffffffffffffffffffff1684610eaf565b949690955092505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610c7557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f742070726f78790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b610c7a815b50565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b600354604080517f723922e7000000000000000000000000000000000000000000000000000000008152905160609273ffffffffffffffffffffffffffffffffffffffff169163723922e7916004808301926000929190829003018186803b158015610d0457600080fd5b505afa158015610d18573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526020811015610d5f57600080fd5b8101908080516040519392919084640100000000821115610d7f57600080fd5b908301906020820185811115610d9457600080fd5b8251866020820283011164010000000082111715610db157600080fd5b82525081516020918201928201910280838360005b83811015610dde578181015183820152602001610dc6565b5050505090500160405250505090505b90565b600354604080517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152915191909216916313af403591602480830192600092919082900301818387803b158015610e6557600080fd5b505af1158015610e79573d6000803e3d6000fd5b5050505050565b6000805473ffffffffffffffffffffffffffffffffffffffff16331415610ea8575032610dee565b5033610dee565b60008060008473ffffffffffffffffffffffffffffffffffffffff1663fce183be856040518263ffffffff1660e01b81526004018082815260200191505060606040518083038186803b158015610f0557600080fd5b505afa158015610f19573d6000803e3d6000fd5b505050506040513d6060811015610f2f57600080fd5b5080516020820151604090920151909791965094509250505056fe4475706c69636174652068617264666f726b206e616d657320617265206e6f7420616c6c6f77656448617264666f726b2063616e6e6f74206265206372656174656420696e20746865207061737446696e616c697a65642068617264666f726b2063616e6e6f742062652064656c6574656448617264666f726b2066696e616c697a6174696f6e20696e74657276616c20657863656564656448462066696e616c697a6174696f6e20626c6f636b20686173682063616e6e6f7420626520656d707479a265627a7a72315820e2d78648acd1c0eaf460724edb81b4230c7ab855137f09a7b6c7cf98d76ce75264736f6c63430005100032`

// HardforkRegistryV1 is an auto generated Go binding around an Ethereum contract.
type HardforkRegistryV1 struct {
	HardforkRegistryV1Caller     // Read-only binding to the contract
	HardforkRegistryV1Transactor // Write-only binding to the contract
	HardforkRegistryV1Filterer   // Log filterer for contract events
}

// HardforkRegistryV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type HardforkRegistryV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HardforkRegistryV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type HardforkRegistryV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HardforkRegistryV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HardforkRegistryV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HardforkRegistryV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HardforkRegistryV1Session struct {
	Contract     *HardforkRegistryV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HardforkRegistryV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HardforkRegistryV1CallerSession struct {
	Contract *HardforkRegistryV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HardforkRegistryV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HardforkRegistryV1TransactorSession struct {
	Contract     *HardforkRegistryV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HardforkRegistryV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type HardforkRegistryV1Raw struct {
	Contract *HardforkRegistryV1 // Generic contract binding to access the raw methods on
}

// HardforkRegistryV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HardforkRegistryV1CallerRaw struct {
	Contract *HardforkRegistryV1Caller // Generic read-only contract binding to access the raw methods on
}

// HardforkRegistryV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HardforkRegistryV1TransactorRaw struct {
	Contract *HardforkRegistryV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewHardforkRegistryV1 creates a new instance of HardforkRegistryV1, bound to a specific deployed contract.
func NewHardforkRegistryV1(address common.Address, backend bind.ContractBackend) (*HardforkRegistryV1, error) {
	contract, err := bindHardforkRegistryV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HardforkRegistryV1{HardforkRegistryV1Caller: HardforkRegistryV1Caller{contract: contract}, HardforkRegistryV1Transactor: HardforkRegistryV1Transactor{contract: contract}, HardforkRegistryV1Filterer: HardforkRegistryV1Filterer{contract: contract}}, nil
}

// NewHardforkRegistryV1Caller creates a new read-only instance of HardforkRegistryV1, bound to a specific deployed contract.
func NewHardforkRegistryV1Caller(address common.Address, caller bind.ContractCaller) (*HardforkRegistryV1Caller, error) {
	contract, err := bindHardforkRegistryV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HardforkRegistryV1Caller{contract: contract}, nil
}

// NewHardforkRegistryV1Transactor creates a new write-only instance of HardforkRegistryV1, bound to a specific deployed contract.
func NewHardforkRegistryV1Transactor(address common.Address, transactor bind.ContractTransactor) (*HardforkRegistryV1Transactor, error) {
	contract, err := bindHardforkRegistryV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HardforkRegistryV1Transactor{contract: contract}, nil
}

// NewHardforkRegistryV1Filterer creates a new log filterer instance of HardforkRegistryV1, bound to a specific deployed contract.
func NewHardforkRegistryV1Filterer(address common.Address, filterer bind.ContractFilterer) (*HardforkRegistryV1Filterer, error) {
	contract, err := bindHardforkRegistryV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HardforkRegistryV1Filterer{contract: contract}, nil
}

// bindHardforkRegistryV1 binds a generic wrapper to an already deployed contract.
func bindHardforkRegistryV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HardforkRegistryV1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HardforkRegistryV1 *HardforkRegistryV1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HardforkRegistryV1.Contract.HardforkRegistryV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HardforkRegistryV1 *HardforkRegistryV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.HardforkRegistryV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HardforkRegistryV1 *HardforkRegistryV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.HardforkRegistryV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HardforkRegistryV1 *HardforkRegistryV1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HardforkRegistryV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HardforkRegistryV1 *HardforkRegistryV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HardforkRegistryV1 *HardforkRegistryV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.contract.Transact(opts, method, params...)
}

// HFSigner is a free data retrieval call binding the contract method 0x1894d2ea.
//
// Solidity: function HF_signer() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1Caller) HFSigner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HardforkRegistryV1.contract.Call(opts, out, "HF_signer")
	return *ret0, err
}

// HFSigner is a free data retrieval call binding the contract method 0x1894d2ea.
//
// Solidity: function HF_signer() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1Session) HFSigner() (common.Address, error) {
	return _HardforkRegistryV1.Contract.HFSigner(&_HardforkRegistryV1.CallOpts)
}

// HFSigner is a free data retrieval call binding the contract method 0x1894d2ea.
//
// Solidity: function HF_signer() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1CallerSession) HFSigner() (common.Address, error) {
	return _HardforkRegistryV1.Contract.HFSigner(&_HardforkRegistryV1.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(uint256[])
func (_HardforkRegistryV1 *HardforkRegistryV1Caller) Enumerate(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _HardforkRegistryV1.contract.Call(opts, out, "enumerate")
	return *ret0, err
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(uint256[])
func (_HardforkRegistryV1 *HardforkRegistryV1Session) Enumerate() ([]*big.Int, error) {
	return _HardforkRegistryV1.Contract.Enumerate(&_HardforkRegistryV1.CallOpts)
}

// Enumerate is a free data retrieval call binding the contract method 0xff9f78b3.
//
// Solidity: function enumerate() constant returns(uint256[])
func (_HardforkRegistryV1 *HardforkRegistryV1CallerSession) Enumerate() ([]*big.Int, error) {
	return _HardforkRegistryV1.Contract.Enumerate(&_HardforkRegistryV1.CallOpts)
}

// GetByBlockNo is a free data retrieval call binding the contract method 0x1658312e.
//
// Solidity: function getByBlockNo(uint256 block_no) constant returns(bytes32 name, bytes32 block_hash, uint256 sw_features)
func (_HardforkRegistryV1 *HardforkRegistryV1Caller) GetByBlockNo(opts *bind.CallOpts, block_no *big.Int) (struct {
	Name       [32]byte
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	ret := new(struct {
		Name       [32]byte
		BlockHash  [32]byte
		SwFeatures *big.Int
	})
	out := ret
	err := _HardforkRegistryV1.contract.Call(opts, out, "getByBlockNo", block_no)
	return *ret, err
}

// GetByBlockNo is a free data retrieval call binding the contract method 0x1658312e.
//
// Solidity: function getByBlockNo(uint256 block_no) constant returns(bytes32 name, bytes32 block_hash, uint256 sw_features)
func (_HardforkRegistryV1 *HardforkRegistryV1Session) GetByBlockNo(block_no *big.Int) (struct {
	Name       [32]byte
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	return _HardforkRegistryV1.Contract.GetByBlockNo(&_HardforkRegistryV1.CallOpts, block_no)
}

// GetByBlockNo is a free data retrieval call binding the contract method 0x1658312e.
//
// Solidity: function getByBlockNo(uint256 block_no) constant returns(bytes32 name, bytes32 block_hash, uint256 sw_features)
func (_HardforkRegistryV1 *HardforkRegistryV1CallerSession) GetByBlockNo(block_no *big.Int) (struct {
	Name       [32]byte
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	return _HardforkRegistryV1.Contract.GetByBlockNo(&_HardforkRegistryV1.CallOpts, block_no)
}

// GetByName is a free data retrieval call binding the contract method 0x8bc237f1.
//
// Solidity: function getByName(bytes32 name) constant returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
func (_HardforkRegistryV1 *HardforkRegistryV1Caller) GetByName(opts *bind.CallOpts, name [32]byte) (struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	ret := new(struct {
		BlockNo    *big.Int
		BlockHash  [32]byte
		SwFeatures *big.Int
	})
	out := ret
	err := _HardforkRegistryV1.contract.Call(opts, out, "getByName", name)
	return *ret, err
}

// GetByName is a free data retrieval call binding the contract method 0x8bc237f1.
//
// Solidity: function getByName(bytes32 name) constant returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
func (_HardforkRegistryV1 *HardforkRegistryV1Session) GetByName(name [32]byte) (struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	return _HardforkRegistryV1.Contract.GetByName(&_HardforkRegistryV1.CallOpts, name)
}

// GetByName is a free data retrieval call binding the contract method 0x8bc237f1.
//
// Solidity: function getByName(bytes32 name) constant returns(uint256 block_no, bytes32 block_hash, uint256 sw_features)
func (_HardforkRegistryV1 *HardforkRegistryV1CallerSession) GetByName(name [32]byte) (struct {
	BlockNo    *big.Int
	BlockHash  [32]byte
	SwFeatures *big.Int
}, error) {
	return _HardforkRegistryV1.Contract.GetByName(&_HardforkRegistryV1.CallOpts, name)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1Caller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HardforkRegistryV1.contract.Call(opts, out, "proxy")
	return *ret0, err
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1Session) Proxy() (common.Address, error) {
	return _HardforkRegistryV1.Contract.Proxy(&_HardforkRegistryV1.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1CallerSession) Proxy() (common.Address, error) {
	return _HardforkRegistryV1.Contract.Proxy(&_HardforkRegistryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1Caller) V1storage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HardforkRegistryV1.contract.Call(opts, out, "v1storage")
	return *ret0, err
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1Session) V1storage() (common.Address, error) {
	return _HardforkRegistryV1.Contract.V1storage(&_HardforkRegistryV1.CallOpts)
}

// V1storage is a free data retrieval call binding the contract method 0x2d059305.
//
// Solidity: function v1storage() constant returns(address)
func (_HardforkRegistryV1 *HardforkRegistryV1CallerSession) V1storage() (common.Address, error) {
	return _HardforkRegistryV1.Contract.V1storage(&_HardforkRegistryV1.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Transactor) Destroy(opts *bind.TransactOpts, _newImpl common.Address) (*types.Transaction, error) {
	return _HardforkRegistryV1.contract.Transact(opts, "destroy", _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Session) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Destroy(&_HardforkRegistryV1.TransactOpts, _newImpl)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _newImpl) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1TransactorSession) Destroy(_newImpl common.Address) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Destroy(&_HardforkRegistryV1.TransactOpts, _newImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Transactor) Migrate(opts *bind.TransactOpts, _oldImpl common.Address) (*types.Transaction, error) {
	return _HardforkRegistryV1.contract.Transact(opts, "migrate", _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Session) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Migrate(&_HardforkRegistryV1.TransactOpts, _oldImpl)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address _oldImpl) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1TransactorSession) Migrate(_oldImpl common.Address) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Migrate(&_HardforkRegistryV1.TransactOpts, _oldImpl)
}

// Propose is a paid mutator transaction binding the contract method 0x072a9823.
//
// Solidity: function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Transactor) Propose(opts *bind.TransactOpts, block_no *big.Int, name [32]byte, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _HardforkRegistryV1.contract.Transact(opts, "propose", block_no, name, block_hash, sw_features)
}

// Propose is a paid mutator transaction binding the contract method 0x072a9823.
//
// Solidity: function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Session) Propose(block_no *big.Int, name [32]byte, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Propose(&_HardforkRegistryV1.TransactOpts, block_no, name, block_hash, sw_features)
}

// Propose is a paid mutator transaction binding the contract method 0x072a9823.
//
// Solidity: function propose(uint256 block_no, bytes32 name, bytes32 block_hash, uint256 sw_features) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1TransactorSession) Propose(block_no *big.Int, name [32]byte, block_hash [32]byte, sw_features *big.Int) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Propose(&_HardforkRegistryV1.TransactOpts, block_no, name, block_hash, sw_features)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 block_no) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Transactor) Remove(opts *bind.TransactOpts, block_no *big.Int) (*types.Transaction, error) {
	return _HardforkRegistryV1.contract.Transact(opts, "remove", block_no)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 block_no) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1Session) Remove(block_no *big.Int) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Remove(&_HardforkRegistryV1.TransactOpts, block_no)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 block_no) returns()
func (_HardforkRegistryV1 *HardforkRegistryV1TransactorSession) Remove(block_no *big.Int) (*types.Transaction, error) {
	return _HardforkRegistryV1.Contract.Remove(&_HardforkRegistryV1.TransactOpts, block_no)
}

// HardforkRegistryV1HardforkIterator is returned from FilterHardfork and is used to iterate over the raw logs and unpacked data for Hardfork events raised by the HardforkRegistryV1 contract.
type HardforkRegistryV1HardforkIterator struct {
	Event *HardforkRegistryV1Hardfork // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HardforkRegistryV1HardforkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HardforkRegistryV1Hardfork)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HardforkRegistryV1Hardfork)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HardforkRegistryV1HardforkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HardforkRegistryV1HardforkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HardforkRegistryV1Hardfork represents a Hardfork event raised by the HardforkRegistryV1 contract.
type HardforkRegistryV1Hardfork struct {
	BlockNo   *big.Int
	BlockHash [32]byte
	Name      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHardfork is a free log retrieval operation binding the contract event 0x9e3eb3a1090f7e2eb48f596218f9322ec1584fad2673784a5cbd5f9e452f18b3.
//
// Solidity: event Hardfork(uint256 block_no, bytes32 block_hash, bytes32 name)
func (_HardforkRegistryV1 *HardforkRegistryV1Filterer) FilterHardfork(opts *bind.FilterOpts) (*HardforkRegistryV1HardforkIterator, error) {

	logs, sub, err := _HardforkRegistryV1.contract.FilterLogs(opts, "Hardfork")
	if err != nil {
		return nil, err
	}
	return &HardforkRegistryV1HardforkIterator{contract: _HardforkRegistryV1.contract, event: "Hardfork", logs: logs, sub: sub}, nil
}

// WatchHardfork is a free log subscription operation binding the contract event 0x9e3eb3a1090f7e2eb48f596218f9322ec1584fad2673784a5cbd5f9e452f18b3.
//
// Solidity: event Hardfork(uint256 block_no, bytes32 block_hash, bytes32 name)
func (_HardforkRegistryV1 *HardforkRegistryV1Filterer) WatchHardfork(opts *bind.WatchOpts, sink chan<- *HardforkRegistryV1Hardfork) (event.Subscription, error) {

	logs, sub, err := _HardforkRegistryV1.contract.WatchLogs(opts, "Hardfork")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HardforkRegistryV1Hardfork)
				if err := _HardforkRegistryV1.contract.UnpackLog(event, "Hardfork", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
