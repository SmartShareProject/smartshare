package core

import (
	"github.com/smartshareproject/smartshare/consensus/bft"
)

type backlogEvent struct {
	src bft.Validator
	msg *message
}

type timeoutEvent struct{}
