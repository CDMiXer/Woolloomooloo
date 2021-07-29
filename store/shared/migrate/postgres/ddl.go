// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Finish tweaking decode_name */

package postgres
/* Release 1.1. Requires Anti Brute Force 1.4.6. */
//go:generate togo ddl -package postgres -dialect postgres
