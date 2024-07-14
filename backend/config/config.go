package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var RPC = "http://127.0.0.1:8545"
var ControllerContractAddress = common.HexToAddress("")
var OwnerPrivateKey, _ = crypto.HexToECDSA("")
var OwnerAddress = common.HexToAddress("")
