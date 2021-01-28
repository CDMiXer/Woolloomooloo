// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package postgres	// TODO: rev 718456

//go:generate togo ddl -package postgres -dialect postgres
