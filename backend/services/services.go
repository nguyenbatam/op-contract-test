package services

import (
	"context"
	"crypto/sha256"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nguyenbatam/op-contract-test/backend/config"
	"github.com/nguyenbatam/op-contract-test/backend/contracts"
	"github.com/nguyenbatam/op-contract-test/backend/db"
	"github.com/nguyenbatam/op-contract-test/backend/libs"
	"github.com/nguyenbatam/op-contract-test/backend/models"
	"github.com/nguyenbatam/op-contract-test/backend/serializers"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Service struct {
	DB    *db.MemoryDB
	index uint
	mu    sync.Mutex
}

func NewService() *Service {
	s := &Service{
		DB:    &db.MemoryDB{},
		index: 0,
	}
	go func() {
		for {
			s.JobUpdateProofDocumentWaiting()
			time.Sleep(1 * time.Hour)
		}
	}()
	return s
}

func (s *Service) UserUploadSample(ctx context.Context, info *serializers.DocumentUpload) *models.Document {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.index++
	riskScore := uint(1 + rand.Intn(4))
	newDocument := models.Document{
		Id:          s.index,
		Owner:       common.HexToAddress(info.Owner),
		Data:        info.Data,
		Hash:        sha256.Sum256([]byte(strconv.FormatUint(uint64(riskScore), 10) + info.Data)),
		RiskScore:   riskScore,
		Timestamp:   time.Now().Unix(),
		ProofStatus: false,
	}
	s.DB.SaveDocument(&newDocument)
	return &newDocument
}

func (s *Service) GetDocumentSample(ctx context.Context, id uint) *models.Document {
	return s.DB.GetDocument(id)
}

func (s *Service) JobUpdateProofDocumentWaiting() {

	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	date := now.Format("2006-01-02 15:04:05")

	listWaiting := s.DB.FindAllDocumentWaiting()

	length := len(listWaiting)
	n := int(math.Ceil(math.Log2(float64(length))))
	powerOfTwo := int(math.Pow(2, float64(n)))
	leaves := make([]common.Hash, powerOfTwo)
	for index, document := range listWaiting {
		leaves[index] = document.Hash
	}
	root := libs.CalculateMerkleRoot(leaves)

	err := s.SubmitAMerkleRoot(date, root)
	if err != nil {
		return
	}
	for index, document := range listWaiting {
		document.Root = root
		document.Proof = libs.GenerateMerkleProof(leaves, index)
		document.ProofStatus = true
		document.Date = date
		s.DB.SaveDocument(document)
	}
}

func (s *Service) SubmitAMerkleRoot(date string, root common.Hash) error {
	client, err := ethclient.Dial(config.RPC)
	if err != nil {
		return err
	}
	defer client.Close()

	nonce, err := client.NonceAt(context.Background(), config.OwnerAddress, nil)
	if err != nil {
		return err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(config.OwnerPrivateKey, chainID)
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(0)
	auth.GasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	controllerContract, err := contracts.NewController(config.ControllerContractAddress, client)
	tx, err := controllerContract.InitializeMerkleRoot(auth, date, root)
	if err != nil {
		return err
	}
	for {
		time.Sleep(2 * time.Second)
		txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil && err != ethereum.NotFound {
			return err
		}
		_ = txReceipt
		break
	}
	return nil
}
