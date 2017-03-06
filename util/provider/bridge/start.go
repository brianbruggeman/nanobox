package bridge

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/nanobox-io/nanobox/util/config"
	"github.com/nanobox-io/nanobox/util"
	"github.com/nanobox-io/nanobox/util/service"
)

func CreateService() error {
	cmd := []string{filepath.Join(config.BinDir(), BridgeClient), "--config", ConfigFile()}

	if runtime.GOOS == "windows" {
		cmd = []string{fmt.Sprintf(`"%s\%s.exe"`, config.BinDir(), BridgeClient), "--config", fmt.Sprintf(`"%s"`, ConfigFile())}
	}

	return service.Create("nanobox-vpn", cmd)
}

func StartService() error {
	fn := func() error {
		return service.Start("nanobox-vpn")
	}
	return util.Retry(fn, 2, 10*time.Second)
}