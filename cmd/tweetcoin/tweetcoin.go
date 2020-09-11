/*
skycoin daemon
*/
package main

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

import (
	"flag"
	_ "net/http/pprof"
	"os"

	"github.com/SkycoinProject/skycoin/src/fiber"
	"github.com/SkycoinProject/skycoin/src/readable"
	"github.com/SkycoinProject/skycoin/src/skycoin"
	"github.com/SkycoinProject/skycoin/src/util/logging"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.27.0"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// CoinName name of coin
	CoinName = "tweetcoin"

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "e77f34b3a22ee7c2f63cf4860dd0db79cbe1b1fc1203fe840240d918cc5e220f3a0a559ba2b9520da1142727ce8920b5eace6913abc18f7e74f93fa71366780501"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "pZ9RPocV6CGoQYbwe36wyChSVtKmXfoQW2"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "022c67fe688692a73f98386bca1cf1d75e72212692397ae0f3425c46ac66773784"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = "83dc765c9f0dc5187176293de3c4fd0cb269b6a2fcece9978d989cbf44c107a7"

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1426562704
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 1000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	}

	nodeConfig = skycoin.NewNodeConfig(ConfigMode, fiber.NodeConfig{
		CoinName:            CoinName,
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://127.0.0.1/peers.txt",
		Port:                6000,
		WebInterfacePort:    6420,
		DataDirectory:       "$HOME/.tweetcoin",

		UnconfirmedBurnFactor:          10,
		UnconfirmedMaxTransactionSize:  32768,
		UnconfirmedMaxDropletPrecision: 3,
		CreateBlockBurnFactor:          10,
		CreateBlockMaxTransactionSize:  32768,
		CreateBlockMaxDropletPrecision: 3,
		MaxBlockTransactionsSize:       32768,

		DisplayName:           "TweetCoin",
		Ticker:                "TWT",
		CoinHoursName:         "Coin Hours",
		CoinHoursNameSingular: "Coin Hour",
		CoinHoursTicker:       "TCH",
		ExplorerURL:           "https://explorer.skycoin.com",
		VersionURL:            "https://version.skycoin.com/skycoin/version.txt",
		Bip44Coin:             8000,
	})

	parseFlags = true
)

func init() {
	nodeConfig.RegisterFlags()
}

func main() {
	if parseFlags {
		flag.Parse()
	}

	// create a new fiber coin instance
	coin := skycoin.NewCoin(skycoin.Config{
		Node: nodeConfig,
		Build: readable.BuildInfo{
			Version: Version,
			Commit:  Commit,
			Branch:  Branch,
		},
	}, logger)

	// parse config values
	if err := coin.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// run fiber coin node
	if err := coin.Run(); err != nil {
		os.Exit(1)
	}
}
