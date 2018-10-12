// Contains initialization code for the mbile library.

package smartshare

import (
	"os"
	"runtime"

	"github.com/smartshareproject/smartshare/log"
)

func init() {
	// Initialize the logger
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))

	// Initialize the goroutine count
	runtime.GOMAXPROCS(runtime.NumCPU())
}
