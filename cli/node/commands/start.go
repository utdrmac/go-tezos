package commands

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/DefinitelyNotAGoat/go-tezos/p2p"
)

func newStartCommand() *cobra.Command {
	var (
		listen  string
		nodedir string
	)

	var start = &cobra.Command{
		Use:   "start",
		Short: "start starts a tezos node",
		Run: func(cmd *cobra.Command, args []string) {
			logger, _ := zap.NewProduction()
			defer logger.Sync()

			//GetIdentity
			identity, err := p2p.GetIdentity(nodedir)
			if err != nil {
				logger.Error("could not start gotezos node", zap.String("err", err.Error()))
				os.Exit(1)
			}

			var peerURLs []p2p.PeerURL
			peers, err := p2p.GetPeers(nodedir)
			if err == nil { // peers found
				peerURLs = peers.GetPeerURLs()
			}

			logger.Info(
				"Starting gotezos node with identity",
				zap.String("peer_id", identity.PeerID),
				zap.String("public_key", identity.PublicKey),
				zap.String("proof_of_work_stamp", identity.ProofOfWorkStamp),
			)

		},
	}

	start.PersistentFlags().StringVar(&listen, "rpc-listen", "http://127.0.0.1:8732", "listen address and port for rpc server (e.g. http://127.0.0.1:8732)")
	start.PersistentFlags().StringVar(&nodedir, "node-dir", "", "sets the path to the node's working directory, if not set the node will use $HOME/.tezos-node")
	return start
}
