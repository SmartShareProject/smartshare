package ssp

import (
	"math/big"
	"os"
	"os/user"
	"time"

	"github.com/smartshareproject/smartshare/common"
	"github.com/smartshareproject/smartshare/common/hexutil"
	"github.com/smartshareproject/smartshare/consensus/bft"
	"github.com/smartshareproject/smartshare/core"
	"github.com/smartshareproject/smartshare/ssp/downloader"
	"github.com/smartshareproject/smartshare/ssp/gasprice"
	"github.com/smartshareproject/smartshare/params"
)

// DefaultConfig contains default settings for use on the Ethereum main net.
var DefaultConfig = Config{
	SyncMode: downloader.FastSync,
	//NetworkId:     1,
	NetworkId:     2017,
	LightPeers:    100,
	DatabaseCache: 768,
	TrieCache:     256,
	TrieTimeout:   5 * time.Minute,
	GasPrice:      big.NewInt(18 * params.GuoSanBin),

	TxPool: core.DefaultTxPoolConfig,
	GPO: gasprice.Config{
		Blocks:     20,
		Percentile: 60,
	},
	BFT: *bft.DefaultConfig,
}

func init() {
	home := os.Getenv("HOME")
	if home == "" {
		if user, err := user.Current(); err == nil {
			home = user.HomeDir
		}
	}
}

//go:generate gencodec -type Config -field-override configMarshaling -formats toml -out gen_config.go

type Config struct {
	// The genesis block, which is inserted if the database is empty.
	// If nil, the Ethereum main net block is used.
	Genesis *core.Genesis `toml:",omitempty"`

	// Protocol options
	NetworkId uint64 // Network ID to use for selecting peers to connect to
	SyncMode  downloader.SyncMode
	NoPruning bool

	// Light client options
	LightServ  int `toml:",omitempty"` // Maximum percentage of time allowed for serving LES requests
	LightPeers int `toml:",omitempty"` // Maximum number of LES client peers

	// Database options
	SkipBcVersionCheck bool `toml:"-"`
	DatabaseHandles    int  `toml:"-"`
	DatabaseCache      int
	TrieCache          int
	TrieTimeout        time.Duration

	// Mining-related options
	Etherbase    common.Address `toml:",omitempty"`
	MinerThreads int            `toml:",omitempty"`
	ExtraData    []byte         `toml:",omitempty"`
	GasPrice     *big.Int

	// Transaction pool options
	TxPool core.TxPoolConfig

	// Gas Price Oracle options
	GPO gasprice.Config

	// Enables tracking of SHA3 preimages in the VM
	EnablePreimageRecording bool

	// BFT options
	BFT bft.Config

	// Miscellaneous options
	DocRoot string `toml:"-"`
}

type configMarshaling struct {
	ExtraData hexutil.Bytes
}
