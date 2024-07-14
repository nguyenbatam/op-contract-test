package models

import "github.com/ethereum/go-ethereum/common"

type Document struct {
	Id          uint           `json:"id"`
	Owner       common.Address `json:"owner"`
	Data        string         `json:"data"`
	Hash        common.Hash    `json:"hash"`
	RiskScore   uint           `json:"risk_score"`
	Date        string         `json:"date"`
	Timestamp   int64          `json:"timestamp"`
	Proof       []common.Hash  `json:"proof"`
	Root        common.Hash    `json:"root"`
	ProofStatus bool           `json:"proof_status"`
}
