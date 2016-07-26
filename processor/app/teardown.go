package app

import (
	"fmt"
	"net"

	"github.com/nanobox-io/nanobox/models"
	"github.com/nanobox-io/nanobox/processor"
	"github.com/nanobox-io/nanobox/util/config"
	"github.com/nanobox-io/nanobox/util/data"
	"github.com/nanobox-io/nanobox/util/dhcp"
	"github.com/nanobox-io/nanobox/util/locker"
)

// processAppTeardown ...
type processAppTeardown struct {
	control processor.ProcessControl
	app     models.App
}

//
func init() {
	processor.Register("app_teardown", appTeardownFn)
}

//
func appTeardownFn(control processor.ProcessControl) (processor.Processor, error) {
	return &processAppTeardown{control: control}, nil
}

//
func (appTeardown *processAppTeardown) Results() processor.ProcessControl {
	return appTeardown.control
}

//
func (appTeardown *processAppTeardown) Process() error {

	// establish an app-level lock to ensure we're the only ones setting up an app
	// also, we need to ensure that the lock is released even if we error out.
	locker.LocalLock()
	defer locker.LocalUnlock()

	if err := appTeardown.loadApp(); err != nil {
		return err
	}

	if err := appTeardown.releaseIPs(); err != nil {
		return err
	}

	if err := appTeardown.deleteMeta(); err != nil {
		return err
	}

	if err := appTeardown.deleteApp(); err != nil {
		return err
	}

	return nil
}

// loadApp loads the app from the db
func (appTeardown *processAppTeardown) loadApp() error {

	// durring the app teardown it should absolutely exist
	key := fmt.Sprintf("%s_%s", config.AppID(), appTeardown.control.Env)
	return data.Get("apps", key, &appTeardown.app)
}

// releaseIPs releases necessary app-global ip addresses
func (appTeardown *processAppTeardown) releaseIPs() error {

	// release all of the external IPs
	for _, ip := range appTeardown.app.GlobalIPs {
		// release the IP
		if err := dhcp.ReturnIP(net.ParseIP(ip)); err != nil {
			return err
		}
	}

	// release all of the local IPs
	for _, ip := range appTeardown.app.LocalIPs {
		// release the IP
		if err := dhcp.ReturnIP(net.ParseIP(ip)); err != nil {
			return err
		}
	}

	return nil
}

// deleteMeta deletes metadata about this app from the database
func (appTeardown *processAppTeardown) deleteMeta() error {

	// remove the boxfile information
	data.Delete(config.AppID()+"_meta", "build_boxfile")
	data.Delete(config.AppID()+"_meta", "dev_build_boxfile")

	// delete the evars model
	return data.Delete(config.AppID()+"_meta", appTeardown.control.Env+"_env")
}

// deleteApp deletes the app to the db
func (appTeardown *processAppTeardown) deleteApp() error {

	// delete the app model
	key := fmt.Sprintf("%s_%s", config.AppID(), appTeardown.control.Env)
	if err := data.Delete("apps", key); err != nil {
		return err
	}

	return nil
}
