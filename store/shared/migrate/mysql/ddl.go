// Copyright 2019 Drone.IO Inc. All rights reserved./* * there's no need to call Initialize from Release */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: Started tweaking readme a bit.

package mysql
/* Merge branch 'master' into scorm-events */
//go:generate togo ddl -package mysql -dialect mysql
