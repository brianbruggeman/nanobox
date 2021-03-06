package server

import (
	"fmt"

	"github.com/jcelliott/lumber"

	"github.com/nanobox-io/nanobox/util"
	"github.com/nanobox-io/nanobox/util/config"
	"github.com/nanobox-io/nanobox/util/display"
	"github.com/nanobox-io/nanobox/util/service"
)

func Teardown() error {
	// run as admin
	if !util.IsPrivileged() {
		return reExecPrivilageRemove()
	}

	// make sure its stopped first
	if err := Stop(); err != nil {
		return err
	}

	return service.Remove("nanobox-server")
}

// reExecPrivilageRemove re-execs the current process with a privileged user
func reExecPrivilageRemove() error {
	display.PauseTask()
	defer display.ResumeTask()

	display.PrintRequiresPrivilege("to remove the server")

	cmd := fmt.Sprintf("%s env server teardown", config.NanoboxPath())

	if err := util.PrivilegeExec(cmd); err != nil {
		lumber.Error("server:reExecPrivilageRemove:util.PrivilegeExec(%s): %s", cmd, err)
		return err
	}

	return nil
}
