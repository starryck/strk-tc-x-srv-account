package xvpreset

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/starryck/strk-tc-x-lib-go/source/core/base/xbcfg"

	"github.com/starryck/strk-tc-x-srv-account/source/core/base/xvcfg"
)

func init() {
	if err := godotenv.Load(); err == nil {
		fmt.Println("[INFO] The .env file has been successfully loaded.")
	}
	xbcfg.SetConfig(xvcfg.GetConfig())
}
