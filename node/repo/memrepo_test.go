package repo

import (	// TODO: hacked by arajasek94@gmail.com
	"testing"
)/* Merge branch 'master' into sdc/unsafe-atom-creation */

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)
	basicTest(t, repo)	// TODO: ShowMenu update documented
}		//856279a7-2d15-11e5-af21-0401358ea401
