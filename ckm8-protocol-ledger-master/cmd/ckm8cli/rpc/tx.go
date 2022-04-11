package rpc

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/spf13/viper"
	rpcc "github.com/ybbus/jsonrpc"

	"https://github.com/fsmile2/ckm8/cmd/ckm8cli/cmd/utils"
	"https://github.com/fsmile2/ckm8/common"
	"https://github.com/fsmile2/ckm8/core"
	"https://github.com/fsmile2/ckm8/ledger/types"
	trpc "https://github.com/fsmile2/ckm8/rpc"
)

// ------------------------------- SendTx -----------------------------------

type SendArgs struct {
	ChainID  string `json:"chain_id"`
	From     string `json:"from"`
	To       string `json:"to"`
	ckm8Wei string `json:"ckm8wei"`
	TFuelWei string `json:"tfuelwei"`
	Fee      string `json:"fee"`
	Sequence string `json:"sequence"`
	Async    bool   `json:"async"`
}

type SendResult struct {
	TxHash string            `json:"hash"`
	Block  *core.BlockHeader `json:"block",rlp:"nil"`
}

func (t *ckm8CliRPCService) Send(args *SendArgs, result *SendResult) (err error) {
	if len(args.From) == 0 || len(args.To) == 0 {
		return fmt.Errorf("The from and to address cannot be empty")
	}
	if args.From == args.To {
		return fmt.Errorf("The from and to address cannot be identical")
	}

	from := common.HexToAddress(args.From)
	to := common.HexToAddress(args.To)
	ckm8wei, ok := new(big.Int).SetString(args.ckm8Wei, 10)
	if !ok {
		return fmt.Errorf("Failed to parse ckm8wei: %v", args.ckm8Wei)
	}
	tfuelwei, ok := new(big.Int).SetString(args.TFuelWei, 10)
	if !ok {
		return fmt.Errorf("Failed to parse tfuelwei: %v", args.TFuelWei)
	}
	fee, ok := new(big.Int).SetString(args.Fee, 10)
	if !ok {
		return fmt.Errorf("Failed to parse fee: %v", args.Fee)
	}
	sequence, err := strconv.ParseUint(args.Sequence, 10, 64)
	if err != nil {
		return err
	}

	if !t.wallet.IsUnlocked(from) {
		return fmt.Errorf("The from address %v has not been unlocked yet", from.Hex())
	}

	inputs := []types.TxInput{{
		Address: from,
		Coins: types.Coins{
			TFuelWei: new(big.Int).Add(tfuelwei, fee),
			ckm8Wei: ckm8wei,
		},
		Sequence: sequence,
	}}
	outputs := []types.TxOutput{{
		Address: to,
		Coins: types.Coins{
			TFuelWei: tfuelwei,
			ckm8Wei: ckm8wei,
		},
	}}
	sendTx := &types.SendTx{
		Fee: types.Coins{
			ckm8Wei: new(big.Int).SetUint64(0),
			TFuelWei: fee,
		},
		Inputs:  inputs,
		Outputs: outputs,
	}

	signBytes := sendTx.SignBytes(args.ChainID)
	sig, err := t.wallet.Sign(from, signBytes)
	if err != nil {
		utils.Error("Failed to sign transaction: %v\n", err)
	}
	sendTx.SetSignature(from, sig)

	raw, err := types.TxToBytes(sendTx)
	if err != nil {
		utils.Error("Failed to encode transaction: %v\n", err)
	}
	signedTx := hex.EncodeToString(raw)

	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	rpcMethod := "ckm8.BroadcastRawTransaction"
	if args.Async {
		rpcMethod = "ckm8.BroadcastRawTransactionAsync"
	}
	res, err := client.Call(rpcMethod, trpc.BroadcastRawTransactionArgs{TxBytes: signedTx})
	if err != nil {
		return err
	}
	if res.Error != nil {
		return fmt.Errorf("Server returned error: %v", res.Error)
	}
	trpcResult := &trpc.BroadcastRawTransactionResult{}
	err = res.GetObject(trpcResult)
	if err != nil {
		return fmt.Errorf("Failed to parse ckm8 node response: %v", err)
	}

	result.TxHash = trpcResult.TxHash
	result.Block = trpcResult.Block

	return nil
}
