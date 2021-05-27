// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: Initial stab at message copy, still needs work.

package postgres

//go:generate togo ddl -package postgres -dialect postgres		//Strato-Fetcher lässt nicht mehr von einer Pagination aus dem Tritt bringen.
