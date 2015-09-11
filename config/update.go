// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package config

import (
	"fmt"
	"path/filepath"
	// "io/ioutil"
	// "net/http"
	"os"
	// "strings"
	// "time"

	// semver "github.com/coreos/go-semver/semver"

	// "github.com/pagodabox/pagodabox-cli/commands"
	// "github.com/pagodabox/pagodabox-cli/util"
	"github.com/pagodabox/nanobox-golang-stylish"
)

// init
func init() {

	// check for a ~/.nanobox/.update file and create one if it's not found
	updatefile := filepath.Clean(Root + "/.update")
	if fi, _ := os.Stat(updatefile); fi == nil {
		fmt.Printf(stylish.Bullet("Creating %s directory", updatefile))
		if _, err := os.Create(updatefile); err != nil {
			panic(err)
		}
	}
}

//
// // update
// func update() {
//
// 	//
// 	fmt.Println("Checking for updates...")
//
// 	//
// 	release, err := getRelease()
// 	if err != nil {
// 		util.LogFatal("[update] getRelease() failed", err)
// 	}
//
// 	if less := config.Version.LessThan(*release); !less {
// 		fmt.Println("No updates available at this time.")
// 		touchUpdate()
// 		return
// 	}
//
// 	cmd := commands.Commands["update"]
// 	cmd.Run("", nil)
// 	os.Exit(0)
// }
//
// // update
// func checkUpdate() {
//
// 	config.Console.Debug("[update] Current version %v", config.Version.String())
//
// 	//
// 	fi, err := os.Stat(config.UpdateFile)
// 	if err != nil {
// 		util.LogFatal("[update] os.Stat() failed", err)
// 	}
//
// 	// check when the last time our .update file was modified
// 	lastUpdate := time.Since(fi.ModTime()).Hours()
//
// 	//
// 	config.Console.Info("[update] last updated %f hours ago", lastUpdate)
//
// 	// set the time (in hours) that we should wait before we check again for updates
// 	updateAfter := 168.0 // 7 days
//
// 	// if the last update was longer ago than our wait time, start the update process
// 	if lastUpdate >= updateAfter {
//
// 		//
// 		fmt.Println("Checking for updates...")
//
// 		//
// 		release, err := getRelease()
// 		if err != nil {
// 			util.LogFatal("[update] getRelease() failed", err)
// 		}
//
// 		// check to see if the current version of the CLI is less than the release
// 		// version. If so, check to see what type of update is needed.
// 		if less := config.Version.LessThan(*release); less {
//
// 			//
// 			config.Console.Info("[update] update required...")
//
// 			util.CPrint(`
// There is a newer version of the CLI available:
// [yellow]Current version: %v[reset]
// [green]Available version: %v[reset]
// 			`, config.Version.String(), release.String())
//
// 			// major update
// 			if config.Version.Major < release.Major {
//
// 				//
// 				config.Console.Info("[update] major update detected")
//
// 				util.CPrint(`
// This is a [red]major[reset] update (required). Major updates contain fixes for
// breaking issues and bugs, or updates needed for the CLI to continue working with
// other Pagodab Box services. [red]The CLI will not function without this update![reset]`)
//
// 				runUpdate()
//
// 				// minor update
// 			} else if config.Version.Minor < release.Minor {
//
// 				//
// 				config.Console.Info("[update] minor update detected")
//
// 				util.CPrint(`
// This is a [yellow]minor[reset] update (not required). Minor updates contain
// fixes to small issues or bugs that may effect workflow. You can continue to use
// the CLI, however its recommneded you update.`)
//
// 				runUpdate()
//
// 				// patch
// 			} else if config.Version.Patch < release.Patch {
//
// 				//
// 				config.Console.Info("[update] patch update detected")
//
// 				util.CPrint(`
// This is a [green]patch[reset] (not required). Patches contain fixes to output,
// help text, or improvements to performance that may enhance the quality of the
// CLI. You can continue to use the CLI or update with [green]'pagoda update'[reset].`)
// 			}
// 		}
//
// 		touchUpdate()
// 	}
// }
//
// // runUpdate
// func runUpdate() {
//
// 	switch util.Prompt("Would you like to update the CLI now (y/N)? ") {
//
// 	// if yes just return and run the update command
// 	case "Yes", "yes", "Y", "y":
// 		return
//
// 		// otherwise exit
// 	default:
// 		util.CPrint("You can manually update at any time with [green]'pagoda update'[reset].")
// 		os.Exit(0)
// 	}
//
// 	// ...run the update
// 	cmd := commands.Commands["update"]
// 	cmd.Run("", nil)
// }
//
// // touchUpdate updates the modification time of the .update file
// func touchUpdate() {
// 	if err := os.Chtimes(config.UpdateFile, time.Now(), time.Now()); err != nil {
// 		util.LogFatal("[commands.update] os.Chtimes() failed", err)
// 	}
//
// 	os.Exit(0)
// }
//
// //
// func getRelease() (*semver.Version, error) {
//
// 	// create a request to get the most recent version from github
// 	req, err := http.NewRequest("GET", "https://api.github.com/repos/pagodabox/pagodabox-cli/contents/version", nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// tell github we want raw format
// 	req.Header.Set("Accept", "application/vnd.github.V3.raw")
//
// 	//
// 	config.Console.Debug("[update] request version from github")
//
// 	// pull the version from github
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// read response body, which should be our version string
// 	b, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	defer res.Body.Close()
//
// 	//
// 	config.Console.Debug("[update] response from github (version): %q", string(b))
//
// 	// trim the version of all newlines and returns
// 	version := strings.TrimSpace(string(b))
//
// 	//
// 	config.Console.Debug("[update] trimmed version: %q", version)
//
// 	// create a release version using the version we pulled from github
// 	return semver.NewVersion(version)
// }