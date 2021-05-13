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
	"github.com/SkycoinProject/skycoin/src/cipher"
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
	CoinName = "test5"

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "2a2e3a50c0f69c06ea483b837ee645d70c2f5882ddf0a0a66bce11f497fd27ce633048c1f6fed1052c3ff733b207558c00c1fd56f5b39bd83962611dbecf4dc801"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "9Sju3QKJVJnLtMn237s2sxQJdkkbcuJFMf"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "0374f540a4d3d4ec843d2f603d6627e5ea36ba41a0c8d7bbae41005cc357035b07"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = "5a9a5f022577bd19a735d7258e2a247648fa5b1a2a46094f9b8e15a95669a251"

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1426562704
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume []cipher.SHA256 = ConvertSHA256HexStrings([]string{
		"59a4c1142cad064776d3816e3177bb23bb9ebcc4bdb7ec67f008410736db13ec",
		"12ab9620096f010a6334e78ecfe06e18f3a91dd517df81c3bf4bda24b1dff8e0",
		"cdf593e718d3f82f438c11fc5923b19c63f26f490d149ff560404561f516d284",
		"63768a97846c64d17eee84ca1a7c5d02267372fe194350e04c69f467957cbb82",
		"fea66184ab3c5f32e765a9369b3175d90dc7383c838d06cf4fa2d9048e0a9961",
		"b50b6533a9b596a080e45de49161210339d18d7bcce378f0f8fbf186c3459f75",
		"46711638474498aecd6c784dad63917e71cf5321b302b8f5c64fdf6572414be0",
		"7118eef062c810f4642918f7f55284e296b2c10086a818561de6cb29fc73513c",
		"e81e82e79187d6e96faa8cec2d135f9ec234b2e4abef3ea1fd424224ba6b15b5",
		"49e25e5620df1e45af02198725da9055f488b259eecc774edb5e487143525bd5",
		"e396c632777ef4f6c9ed392efca65717d63a5aab33bb81ac311fd2db5278e6f3",
		"4b380c96ebb1284e6f23b4af23b9786e424a7f8b9e8be231ace9168c79f959b7",
		"41a4d834ae1624f183dbbf01c8e05ffaa80ee7e4bee0fda85f11df427966b7bd",
		"4436106d0517addb3830a5272d9c4ce8400c43a444f2b6a5607fa07acb90775e",
		"a4bcf31c7358d794a2f91aba85e20f3f4f85e7f905d63a8cfc2763ec00fad336",
		"622ed012be4ebadd8a326025180d98b5c3f0036b4549c0f994ef88e17933612c",
		"e84252ea728fa47268c48376d07af715d4947f570a164f2df8442ae039449455",
		"21a572daf2d7df79eb290b4ad2f3c0f1412bcb72e788f4f222382a3ebec5bd80",
		"17ea9f6e54bffa0edc2772817d561f0579df401c203032fc07f8de69bdf7a24f",
		"2ac2f6d9e697aca549454916af291d6b17ebdcf30aae45ae9d4e979dec13ce35",
		"7220100413a8f1e3707cbf34a00adb42638d9027f853b0e669a28ae30d761ba6",
		"8043cdf93fc2a055d09e4a678c677b671bf775c175d5053285d70a0b93c139a3",
		"520b144f41dd79ae7da14228aba059fdc3929e43afecc46dc4050e4fafaeec70",
		"4704accdc4d45ffc0f7cb26eb01158598fe213ad925050207d04ecfda04a152d",
		"b19bc356f60f2cf6aff459dcc2d5bf7cdfa6f55bdbee890decb8a33bef927307",
		"599a68fe07610da1b1c0c8058356facb9d353397c364810c5e47080207dafe64",
		"c28d5a7b5a9feeb6e6748a43bb9ccfadea2359be326ce4ef6ddef2e27677b00b",
		"a561f0cf5f0acbf8f55d00f9c5fea51ff4b72e61c3a9acb122e6226bd4bfc206",
		"fad95ffd77c407080a77f3ee8becc0c91d548d10ffb3b5c119def349d9bdd70f",
		"3ff685084ed343d36c22814442b4e7b188ab6c92e8feb060e2f623074b2db8b2",
		"12ed3d689f2604d2e28f4a63aa56d835a8f1d2d205af5eaf1c26e08681b178d5",
		"3cdff4669bc136379c205f1bf43fba37f9c1d43965b38057c53379ce90d69974",
		"da738f5145b6a9f7ac4d2d73b584dee0a475cca6165b4f29b9bf38635278dc44",
		"694eb393c1bf5e22f06139af0c2cf08059abe6bba1973ac2ed8319c0c88949ce",
		"e512863711757f751b4621022d4f39676d5749ca57ca637e07eb3fabe93b0043",
		"a7f483c9402977e1770d1a65f330ef2e3f4f88023e80575be48c7b969c10c5b9",
		"ad67056d5510f0c30ac54ebccc7982ab7fdc321a7289ae2d4b02f494ddb0f96a",
		"fb13d1fd5ecd349adb81295dbf8b75773819e4f4880f824cb3f2c861b3b4853c",
		"d426a76432b4fddc7b2fd9868b6d55bf2ac709c56c82d2178d46d4de5e86f974",
		"a744183da929fa44a382aef2a7f9b001a771bbd5406ae9ce62933522a9ad7ae7",
		"7adb7087f96623cb821972e399a5687890070686a4de8407543e93e6b7766c9d",
		"5b60eae87888d7fcbcdcc9fa734a393df5c8d74e9ef83a83ded1e78ea88ddf24",
		"d57c7c508df0598b990fce1f62f66a870711859758183922d1f8544e3a2d5bd5",
		"89308c5235f548878e184d60923d4667338bb1a40007a34a2893793a5a8a5981",
		"5e8f0aab78b2809f8746fa3efd20dff394decb4fdc674d4c6f529f86717363f9",
		"f37949c20e13b227bba32f6f965dd950b3dcb591e3c0875a26605f3b2a92b4a3",
		"7a57aa11298336d455ba1ee87ce0f38133cb94c2e39a1e9d739b78d32909e1e8",
		"a484fcafa44e2185e75b3be35e44c50d11420eb84ffacc0911f06ec01620dec2",
		"763dd6b875a4c223fe9eb5a7b4bd8eda3cf999f6038e4d4628a54d59d0d57969",
		"ef2b25721d157d5af0643ffbcaee35f7b974c1498ba6d1b1e615fad158f23413",
		"cf2341eb96497317d1c8ed62613083eab2ecd1b61026eca84d3cfb7e204a51cc",
		"da0c494b860dc4b8333ab16d6618bd40954fe15796f50bf069d3a8119e6a57d5",
		"1c0c006265dda38278324878703dc86dc900dfa86d47ce391e340d273ff2a914",
		"748a61ec05b81fe091ea0aac48e36bba4af67f3fd6238e0cb440db03a888d241",
		"64d5be6ebc7f1fcee35460c9737343fe0adf5bbaafb650a29d33c2c175a61c29",
		"c52f3069ec1ec9c2af10f40fdb785b53f6be12396c70d046aa8c33fd87da9c27",
		"dc6f1a171b9e82c0e2c25b624ed969660209a2977f8b063d45f0d9ea52536420",
		"3290c6a95cd16bd78c3fc912371bb9d26794ef9e97d1b9d9b024bc657d8da7b2",
		"28acf3276da4eba02d751a2d87481b4ad0dcd407eff6a3eb3b922661a9862d9e",
		"46e5cc3a29e553e83876d7ce7b65b55fa2b38dc77ec06ca59df40a46521af92d",
		"e60650afc4b2bbb88f9540beb9dede351bea52dba7accf5f92049d68b017c35d",
		"139446fbf2032e2672965426304d5490b0d1b0aa8ee598974c698bc1889d138f",
		"b029c5e8df0982e571b014ca66f82238a19527428c00ede9e4810c6904d9a5d5",
		"b1e8532d7f85683806c68cae179d0e6f3f893f4b01974983fcfadcb0b39369e4",
		"7ccca1fc0bc4ad213184412c4d2685681513b99607892c7acde9b876baaa72ce",
		"0b0f9180ab64459a4c6e2eee5ba4df7fa24b6b6830b2085855a3c6ec0950b55d",
		"6c65a802cb33dfa80b38c6cf507eba481b0c26e83f18434f4e15b8205da839c7",
		"8b32d3ed2e2e0a03806d5492086fe508c20ab13b9633c57119414215311c2a94",
		"6498dee798ad7b39b4aa76f9504f0a0d92ccced2d1e8bc55c1428fa9483b1eb5",
		"a02ea6bdc250c4ceb027b6d3d64a31e898087db1ce11362764f305677f390c61",
		"04982e6ee8f31555ccef91da66ceffd24183e367c7fe6e20411ab9e7209bd527",
		"77fa0e76e5867630e117d27ab20aaee844781731a41dfa48f48a56ce9e90996b",
		"25eabaa843f0b57f69676236b905dacb0b4001416110a4c8f7a8f2b03a27d188",
		"6149ec182c1285c708815fcc0dfac54cdf8741abe00c0971faa8b76b4b6385ba",
		"639459ddd83386cad7aaffa8ce862683f179ebb22bd45989623feb4c63c593b8",
		"1117a1caf9bc396a4e98afbe8772d45554ca98fe5653d0bfce69bc6971e16ee2",
		"8b7713b47e1e4587d0cd434fcefe8d014e50af60b35b9fbd14ed5ce58467c6a8",
		"501c765f6e07f3d027ec29d406d15dee95a08bbfa8b8caeb17f1674b6671b516",
		"f576635d1abe5e2f7ba6fe728efaf5ff80d0c02b8d6395094db5b9d7f43736c7",
		"8dd42747afe08974385bd6e637e5b252831a280b974fa6e840f3371c04048b1a",
		"69e8277363f9aad6c555cdc6ac5da7c15d9bc028a9719b1d8966ec0994b1799a",
		"6aa1eba80bf204400db507ab626da076d36aef2316f7aed019b0bbc0386c9b28",
		"3cf3762dda4bd6f6261b3b6438e7e1e44d5d135a0d5aea4f1e56f7aa4e98eb9a",
		"b57d8f461775c952c78761817e603edfdae5ed358e667a6f6d9a8fbb216c8a33",
		"f922b85cb5290df720906e53fe6357091bf2d9036e5a7b29b1480d39d25a89bd",
		"33dfd0e9df5beb5566b9eae84fa46fd81eaf128f6db62e7012adfe5e13fb103f",
		"3ae043d63347d551f840dcc5d97e90a5db97d7d8aa00586f20d0711548f46a71",
		"1df8d35045bcc65c4a4761ffd4f261fcf1abc1833c9f7e84d14b5cc84ff75c34",
		"293cbe5c93299a5b1c5b80f5c5d8ff43b81587b93f29a243d4be6b708c29343d",
		"c11c648e3c522150f013cc663dcf18e8900b5d155a2a36a2df430fd8cf5ddf93",
		"7ffb58ab81c3daf1d8b2cd4e2806449e1f0c99fc15bd2720747827fd96f8bc94",
		"d2ad88ba5bba31d62df0438e437e8b3d49757f48a99b52d4a7b6e82d72d19b93",
		"9dd2a39f52a040b9e80ab84028a6654f533a307139d222bc9f65774387f300b3",
		"27e1d1cdfcbbd6ad88e0caa6991dd5fea06807d78c19df80b231c77f9d1d40b8",
		"c9576e0fedd3c342ce203e44cd0627bc4ae19e52a7b24111272c9abc33a613dc",
		"0c35db2b7c392ae23bfe1f3c22f11e44a1d9d43ef412fc4b127fc30601cc11c9",
		"2cddc43df87ce5de9d505ef74bae8a99c30e43397dccea194e20feb5c52d6642",
		"2814a376c8465e71a3580c78d6e1cd3875cdf7817404e7a2939adf1f715dadd7",
		"950b977727b1949f35e544591da66acfec3d9bb940c5ea28a508805e39066034",
		"ff5644c51ddde68b4c32b25308740f5e865fc150b9e7f330b2832a3351f089ef",
	})

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
		DataDirectory:       "$HOME/.test5",

		UnconfirmedBurnFactor:          100,
		UnconfirmedMaxTransactionSize:  134217728,
		UnconfirmedMaxDropletPrecision: 3,
		CreateBlockBurnFactor:          100,
		CreateBlockMaxTransactionSize:  134217728,
		CreateBlockMaxDropletPrecision: 3,
		MaxBlockTransactionsSize:       134217728,

		DisplayName:           "Skycoin",
		Ticker:                "SKY",
		CoinHoursName:         "Coin Hours",
		CoinHoursNameSingular: "Coin Hour",
		CoinHoursTicker:       "SCH",
		ExplorerURL:           "https://explorer.skycoin.com",
		VersionURL:            "https://version.skycoin.com/skycoin/version.txt",
		Bip44Coin:             8000,
	})

	parseFlags = true
)

func ConvertSHA256HexStrings(a []string) []cipher.SHA256 {
	rv := make([]cipher.SHA256, len(a))

	for i, _ := range a {
		rv[i], _ = cipher.SHA256FromHex(a[i])
	}

	return rv
}

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
