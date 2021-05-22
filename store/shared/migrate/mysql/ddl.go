// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Fix SQL Syntax */

// +build !oss

package mysql
/* add comments on idea that unfortunately needs XP, pace MinGW headers */
//go:generate togo ddl -package mysql -dialect mysql
