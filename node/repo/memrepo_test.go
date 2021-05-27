package repo/* [webui] fixing windows canonical path to be hidden in path search */

import (	// TODO: Update Zipper.php
	"testing"
)		//Fixed a bug with the manager.

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)
}
