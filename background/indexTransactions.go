package background

import (
	"block-banter/database"
	"block-banter/models"
	"block-banter/repository"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strings"
)

const erc20ABI = `[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

type ERCC20TransferEvent struct {
	From   common.Address
	To     common.Address
	Value  *big.Int
	TxHash common.Hash
}

func IndexTransactions() {

	client, err := ethclient.Dial("wss://ethereum-mainnet.core.chainstack.com/" + os.Getenv("CHAINSTACK_API_KEY"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum MainNet: %v", err)
	}
	defer client.Close()
	log.Println("Connected to Ethereum MainNet")

	contractAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // usdc
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to filter logs: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		log.Fatalf("Failed to parse ERC20 ABI: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:

			if vLog.Topics[0].Hex() != "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
				continue
			}

			var transferEvent ERCC20TransferEvent
			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatalf("Failed to unpack log data: %v", err)
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
			transferEvent.TxHash = vLog.TxHash

			if err != nil {
				log.Fatalf("Failed to marshal transfer event: %v", err)
			}

			repo := repository.NewTransferEventRepository(database.DB)

			event := &repository.TransferEvent{
				From:   transferEvent.From.String(),
				To:     transferEvent.To.String(),
				Value:  models.BigInt{Int: transferEvent.Value},
				TxHash: transferEvent.TxHash.String(),
			}

			err = repo.Create(event)
			if err != nil {
				log.Fatalf("failed to create event: %v", err)
			}

		}
	}
}
