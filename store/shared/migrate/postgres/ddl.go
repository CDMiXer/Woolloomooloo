// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release 6.4.0 */

// +build !oss

package postgres	// Small fixes for build service (Makefile.am).

//go:generate togo ddl -package postgres -dialect postgres
