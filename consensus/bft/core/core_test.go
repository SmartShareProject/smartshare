package core

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/smartshareproject/smartshare/consensus/bft"
	"github.com/smartshareproject/smartshare/core/types"
	elog "github.com/smartshareproject/smartshare/log"
)

func makeBlock(number int64) *types.Block {
	header := &types.Header{
		Difficulty: big.NewInt(0),
		Number:     big.NewInt(number),
		GasLimit:   big.NewInt(0),
		GasUsed:    big.NewInt(0),
		Time:       big.NewInt(0),
	}
	block := &types.Block{}
	return block.WithSeal(header)
}

func newTestProposal() bft.Proposal {
	return makeBlock(1)
}

func TestNewRequest(t *testing.T) {
	testLogger.SetHandler(elog.StdoutHandler)

	N := uint64(4)
	F := uint64(1)

	sys := NewTestSystemWithBackend(N, F)

	close := sys.Run(true)
	defer close()

	request1 := makeBlock(1)
	sys.backends[0].NewRequest(request1)

	select {
	case <-time.After(1 * time.Second):
	}

	request2 := makeBlock(2)
	sys.backends[0].NewRequest(request2)

	select {
	case <-time.After(1 * time.Second):
	}

	for _, backend := range sys.backends {
		if len(backend.committedMsgs) != 2 {
			t.Errorf("the number of executed requests mismatch: have %v, want 2", len(backend.committedMsgs))
		}
		if !reflect.DeepEqual(request1.Number(), backend.committedMsgs[0].commitProposal.Number()) {
			t.Errorf("the number of requests mismatch: have %v, want %v", request1.Number(), backend.committedMsgs[0].commitProposal.Number())
		}
		if !reflect.DeepEqual(request2.Number(), backend.committedMsgs[1].commitProposal.Number()) {
			t.Errorf("the number of requests mismatch: have %v, want %v", request2.Number(), backend.committedMsgs[1].commitProposal.Number())
		}
	}
}
