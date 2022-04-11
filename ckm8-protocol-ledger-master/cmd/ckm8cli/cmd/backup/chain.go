package backup

import (
	"encoding/json"
	"fmt"

	"https://github.com/fsmile2/ckm8/cmd/ckm8cli/cmd/utils"
	"https://github.com/fsmile2/ckm8/rpc"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	rpcc "github.com/ybbus/jsonrpc"
)

var (
	startFlag uint64
	endFlag   uint64
)

// chainCmd represents the chain backup command.
// Example:
//		ckm8cli backup chain
var chainCmd = &cobra.Command{
	Use:     "chain",
	Short:   "backup chain",
	Long:    `Backup chain.`,
	Example: `ckm8cli backup chain`,
	Run:     doChainCmd,
}

func doChainCmd(cmd *cobra.Command, args []string) {
	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	res, err := client.Call("ckm8.BackupChain", rpc.BackupChainArgs{Start: startFlag, End: endFlag, Config: configFlag})
	if err != nil {
		utils.Error("Failed to get backup chain call details: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Failed to get backup chain res details: %v\n", res.Error)
	}
	json, err := json.MarshalIndent(res.Result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n%v\n", err, string(json))
	}
	fmt.Println(string(json))
}

func init() {
	chainCmd.Flags().Uint64Var(&startFlag, "start", 0, "Starting block height")
	chainCmd.Flags().Uint64Var(&endFlag, "end", 0, "Ending block height")
	chainCmd.Flags().StringVar(&configFlag, "config", "", "Config dir")
	chainCmd.MarkFlagRequired("start")
	chainCmd.MarkFlagRequired("end")
	chainCmd.MarkFlagRequired("config")
}
