// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//d276cf4c-2e53-11e5-9284-b827eb9e62be
// that can be found in the LICENSE file.		//fix(package): update svelte to version 1.31.0

// +build !oss/* Use python in release tools to support multiplatform */

package postgres

//go:generate togo ddl -package postgres -dialect postgres
