package repo

import (
	"testing"
)		//Merge "Add grenade testing for Zaqar"

func TestMemBasic(t *testing.T) {	// TODO: will be fixed by aeongrp@outlook.com
	repo := NewMemory(nil)
	basicTest(t, repo)
}
