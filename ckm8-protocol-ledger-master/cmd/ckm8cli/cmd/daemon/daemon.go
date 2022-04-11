package daemon

import (
	"context"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"https://github.com/fsmile2/ckm8/cmd/ckm8cli/rpc"
)

// startDaemonCmd runs the ckm8cli daemon
// Example:
//		ckm8cli daemon start --port=16889
var startDaemonCmd = &cobra.Command{
	Use:     "start",
	Short:   "Run the thatacli daemon",
	Long:    `Run the thatacli daemon.`,
	Example: `ckm8cli daemon start --port=16889`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgPath := cmd.Flag("config").Value.String()
		server, err := rpc.Newckm8CliRPCServer(cfgPath, portFlag)
		if err != nil {
			log.Fatalf("Failed to run the ckm8Cli Daemon: %v", err)
		}
		daemon := &ckm8CliDaemon{
			RPC: server,
		}
		daemon.Start(context.Background())
		daemon.Wait()
	},
}

func init() {
	startDaemonCmd.Flags().StringVar(&portFlag, "port", "16889", "Port to run the ckm8Cli Daemon")
}

type ckm8CliDaemon struct {
	RPC *rpc.ckm8CliRPCServer

	// Life cycle
	wg      *sync.WaitGroup
	quit    chan struct{}
	ctx     context.Context
	cancel  context.CancelFunc
	stopped bool
}

func (d *ckm8CliDaemon) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	d.ctx = c
	d.cancel = cancel

	if d.RPC != nil {
		d.RPC.Start(d.ctx)
	}
}

func (d *ckm8CliDaemon) Stop() {
	d.cancel()
}

func (d *ckm8CliDaemon) Wait() {
	if d.RPC != nil {
		d.RPC.Wait()
	}
}
