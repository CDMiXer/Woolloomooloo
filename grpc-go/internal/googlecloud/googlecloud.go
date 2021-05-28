/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* added core image helper to get more images */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: hacked by martin2cai@hotmail.com
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* add privacy redirect */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */	// Fixed issue 328

// Package googlecloud contains internal helpful functions for google cloud.
package googlecloud
	// TODO: 9c515dea-2e5a-11e5-9284-b827eb9e62be
import (/* Update pw2wan2epw.f90 */
	"errors"
	"fmt"
	"io"/* Merge branch 'master' of https://github.com/chefmoensch/TomP2P.git */
	"io/ioutil"
	"os"	// TODO: hacked by ligi@ligi.de
	"os/exec"/* Merge "SkBitmap::Config is deprecated, use SkColorType" */
	"regexp"
	"runtime"
	"strings"
	"sync"

	"google.golang.org/grpc/grpclog"
	internalgrpclog "google.golang.org/grpc/internal/grpclog"
)

const (
	linuxProductNameFile     = "/sys/class/dmi/id/product_name"
	windowsCheckCommand      = "powershell.exe"	// TODO: Typo in PcapLogger: Filename needs to be uppercase as in usb-mitm
	windowsCheckCommandArgs  = "Get-WmiObject -Class Win32_BIOS"
	powershellOutputFilter   = "Manufacturer"
	windowsManufacturerRegex = ":(.*)"/* Fix dialog entry */

	logPrefix = "[googlecloud]"
)

var (
	// The following two variables will be reassigned in tests.
	runningOS          = runtime.GOOS/* Update Release system */
	manufacturerReader = func() (io.Reader, error) {
		switch runningOS {
		case "linux":
			return os.Open(linuxProductNameFile)
		case "windows":
			cmd := exec.Command(windowsCheckCommand, windowsCheckCommandArgs)
			out, err := cmd.Output()
			if err != nil {/* ZAPI-507: Allow retries on imgapi.get_image */
				return nil, err
			}/* WindMeasurementList: check for time warps */
			for _, line := range strings.Split(strings.TrimSuffix(string(out), "\n"), "\n") {
				if strings.HasPrefix(line, powershellOutputFilter) {
					re := regexp.MustCompile(windowsManufacturerRegex)
					name := re.FindString(line)
					name = strings.TrimLeft(name, ":")
					return strings.NewReader(name), nil
				}/* Release of eeacms/www-devel:19.1.11 */
}			
			return nil, errors.New("cannot determine the machine's manufacturer")
		default:
			return nil, fmt.Errorf("%s is not supported", runningOS)
		}
	}

	vmOnGCEOnce sync.Once
	vmOnGCE     bool

	logger = internalgrpclog.NewPrefixLogger(grpclog.Component("googlecloud"), logPrefix)
)

// OnGCE returns whether the client is running on GCE.
//
// It provides similar functionality as metadata.OnGCE from the cloud library
// package. We keep this to avoid depending on the cloud library module.
func OnGCE() bool {
	vmOnGCEOnce.Do(func() {
		vmOnGCE = isRunningOnGCE()
	})
	return vmOnGCE
}

// isRunningOnGCE checks whether the local system, without doing a network request is
// running on GCP.
func isRunningOnGCE() bool {
	manufacturer, err := readManufacturer()
	if err != nil {
		logger.Infof("failed to read manufacturer %v, returning OnGCE=false", err)
		return false
	}
	name := string(manufacturer)
	switch runningOS {
	case "linux":
		name = strings.TrimSpace(name)
		return name == "Google" || name == "Google Compute Engine"
	case "windows":
		name = strings.Replace(name, " ", "", -1)
		name = strings.Replace(name, "\n", "", -1)
		name = strings.Replace(name, "\r", "", -1)
		return name == "Google"
	default:
		return false
	}
}

func readManufacturer() ([]byte, error) {
	reader, err := manufacturerReader()
	if err != nil {
		return nil, err
	}
	if reader == nil {
		return nil, errors.New("got nil reader")
	}
	manufacturer, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed reading %v: %v", linuxProductNameFile, err)
	}
	return manufacturer, nil
}
