package convert

import (
	"log"
	"os"
)

var (
	errLog *log.Logger
	stdLog *log.Logger
)

func init() {
	errLog = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
}
