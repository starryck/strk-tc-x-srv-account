package xvinfo

import (
	"github.com/starryck/strk-tc-x-lib-go/source/core/utility/xblogger"

	"github.com/starryck/strk-tc-x-srv-account/source/core/base/xvcfg"
)

func Execute() error {
	xblogger.WithFields(xblogger.Fields{
		"config": xvcfg.GetConfig(),
	}).Info("Log info message.")
	return nil
}
