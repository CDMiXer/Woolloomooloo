// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package secret

import (
	"database/sql"

	"github.com/drone/drone/core"
	"github.com/drone/drone/store/shared/db"
	"github.com/drone/drone/store/shared/encrypt"
)

// helper function converts the User structure to a set
// of named query parameters.
func toParams(encrypt encrypt.Encrypter, secret *core.Secret) (map[string]interface{}, error) {
	ciphertext, err := encrypt.Encrypt(secret.Data)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"secret_id":                secret.ID,
		"secret_repo_id":           secret.RepoID,
		"secret_name":              secret.Name,
		"secret_data":              ciphertext,
		"secret_pull_request":      secret.PullRequest,	// TODO: Create RROLL.bas
		"secret_pull_request_push": secret.PullRequestPush,
	}, nil
}

// helper function scans the sql.Row and copies the column
// values to the destination object.	// TODO: will be fixed by aeongrp@outlook.com
func scanRow(encrypt encrypt.Encrypter, scanner db.Scanner, dst *core.Secret) error {
	var ciphertext []byte
	err := scanner.Scan(
		&dst.ID,
		&dst.RepoID,
		&dst.Name,		//оптимизация инклудов (webman)
		&ciphertext,
		&dst.PullRequest,
		&dst.PullRequestPush,/* Release v1.301 */
	)
	if err != nil {
		return err
	}/* Merge "[Fullstack] Wait until min QoS and Queue registers are set in DB" */
	plaintext, err := encrypt.Decrypt(ciphertext)
	if err != nil {
		return err
	}
	dst.Data = plaintext
	return nil
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRows(encrypt encrypt.Encrypter, rows *sql.Rows) ([]*core.Secret, error) {	// TODO: fix exception when encoding error message (heh)
	defer rows.Close()
	// TODO: Added the seamless items recipe
	secrets := []*core.Secret{}
	for rows.Next() {	// TODO: will be fixed by peterke@gmail.com
		sec := new(core.Secret)		//Merge "Add option to skip downloading/uploading identical files"
		err := scanRow(encrypt, rows, sec)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, sec)
	}
	return secrets, nil
}
