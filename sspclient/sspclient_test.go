package sspclient

import "github.com/smartshareproject/smartshare"

// Verify that Client implements the smartshare interfaces.
var (
	_ = smartshare.ChainReader(&Client{})
	_ = smartshare.TransactionReader(&Client{})
	_ = smartshare.ChainStateReader(&Client{})
	_ = smartshare.ChainSyncReader(&Client{})
	_ = smartshare.ContractCaller(&Client{})
	_ = smartshare.GasEstimator(&Client{})
	_ = smartshare.GasPricer(&Client{})
	_ = smartshare.LogFilterer(&Client{})
	_ = smartshare.PendingStateReader(&Client{})
	// _ = smartshare.PendingStateEventer(&Client{})
	_ = smartshare.PendingContractCaller(&Client{})
)
