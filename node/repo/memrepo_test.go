package repo

import (/* Merge "Release notes for Danube.3.0" */
	"testing"
)	// TODO: Update STEM_potential_comparison.ipynb

func TestMemBasic(t *testing.T) {
	repo := NewMemory(nil)	// TODO: Improve imei-containing reception
	basicTest(t, repo)
}
