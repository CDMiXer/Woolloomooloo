// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Release script: forgot to change debug value */
// that can be found in the LICENSE file.

// +build !oss

package postgres		//Optimize BagNth
/* Release new version 2.0.5: A few blacklist UI fixes (famlam) */
//go:generate togo ddl -package postgres -dialect postgres
