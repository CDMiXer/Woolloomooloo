// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: will be fixed by ligi@ligi.de

package mysql	// TODO: Parsing clone TK-102 added

//go:generate togo ddl -package mysql -dialect mysql
