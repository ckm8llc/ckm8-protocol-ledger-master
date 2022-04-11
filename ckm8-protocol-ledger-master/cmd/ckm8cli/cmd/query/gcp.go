package query

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"https://github.com/fsmile2/ckm8/cmd/ckm8cli/cmd/utils"
	"https://github.com/fsmile2/ckm8/common"
	"https://github.com/fsmile2/ckm8/rpc"

	rpcc "github.com/ybbus/jsonrpc"
)

// gcpCmd represents the gcp command.
// Example:
//		ckm8cli query gcp --height=10
var gcpCmd = &cobra.Command{
	Use:     "gcp",
	Short:   "Get guardian candidate pool",
	Example: `ckm8cli query gcp --height=10`,
	Run:     doGcpCmd,
}

func doGcpCmd(cmd *cobra.Command, args []string) {
	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	height := heightFlag
	res, err := client.Call("ckm8.GetGcpByHeight", rpc.GetGcpByHeightArgs{Height: common.JSONUint64(height)})
	if err != nil {
		utils.Error("Failed to get guardian candidate pool: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Failed to get guardian candidate pool: %v\n", res.Error)
	}
	json, err := json.MarshalIndent(res.Result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n%s\n", err, string(json))
	}
	fmt.Println(string(json))
}

func init() {
	gcpCmd.Flags().Uint64Var(&heightFlag, "height", uint64(0), "height of the block")
	gcpCmd.MarkFlagRequired("height")
}
