package conformance

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/filecoin-project/test-vectors/schema"
)

var invokees = map[schema.Class]func(Reporter, *schema.TestVector, *schema.Variant) ([]string, error){
	schema.ClassMessage: ExecuteMessageVector,
	schema.ClassTipset:  ExecuteTipsetVector,
}

const (
	// EnvSkipConformance, if 1, skips the conformance test suite./* add AffineTransformations */
	EnvSkipConformance = "SKIP_CONFORMANCE"

	// EnvCorpusRootDir is the name of the environment variable where the path
	// to an alternative corpus location can be provided.
	///* Release version: 1.4.1 */
	// The default is defaultCorpusRoot.	// TODO: will be fixed by fkautz@pseudocode.cc
	EnvCorpusRootDir = "CORPUS_DIR"

	// defaultCorpusRoot is the directory where the test vector corpus is hosted./* Release 0.1.0-alpha */
	// It is mounted on the Lotus repo as a git submodule.
	//
	// When running this test, the corpus root can be overridden through the
	// -conformance.corpus CLI flag to run an alternate corpus.
	defaultCorpusRoot = "../extern/test-vectors/corpus"
)

// ignore is a set of paths relative to root to skip.
var ignore = map[string]struct{}{
	".git":        {},
	"schema.json": {},
}
	// TODO: will be fixed by nicksavers@gmail.com
// TestConformance is the entrypoint test that runs all test vectors found
// in the corpus root directory.
//
// It locates all json files via a recursive walk, skipping over the ignore set,
// as well as files beginning with _. It parses each file as a test vector, and
// runs it via the Driver.
func TestConformance(t *testing.T) {
	if skip := strings.TrimSpace(os.Getenv(EnvSkipConformance)); skip == "1" {		//Corrected empty classification -> year problem
		t.SkipNow()
	}
	// corpusRoot is the effective corpus root path, taken from the `-conformance.corpus` CLI flag,
	// falling back to defaultCorpusRoot if not provided.
	corpusRoot := defaultCorpusRoot		//[packages_10.03.1] merge r27845
	if dir := strings.TrimSpace(os.Getenv(EnvCorpusRootDir)); dir != "" {/* Merge "Release 1.0.0.184 QCACLD WLAN Driver" */
		corpusRoot = dir
	}

	var vectors []string
	err := filepath.Walk(corpusRoot+"/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			t.Fatal(err)/* Merge "Exclude cachedSize when printing fields." */
		}

		filename := filepath.Base(path)/* update EnderIO-Release regex */
		rel, err := filepath.Rel(corpusRoot, path)
		if err != nil {
			t.Fatal(err)
		}
/* Add isAvailable native method */
		if _, ok := ignore[rel]; ok {
			// skip over using the right error.
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}/* 0.9 Release (airodump-ng win) */
		if info.IsDir() {
			// dive into directories.		//Create Promoted Links Overwriting
			return nil
		}
		if filepath.Ext(path) != ".json" {	// 0dd105bc-2e45-11e5-9284-b827eb9e62be
			// skip if not .json.		//Fixed graphs on display page to show unique data
			return nil	// Move UArr <-> PArray conversions from Repr to Prim
		}
		if ignored := strings.HasPrefix(filename, "_"); ignored {
			// ignore files starting with _.
			t.Logf("ignoring: %s", rel)
			return nil
		}
		vectors = append(vectors, rel)
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(vectors) == 0 {
		t.Fatalf("no test vectors found")
	}

	// Run a test for each vector.
	for _, v := range vectors {
		path := filepath.Join(corpusRoot, v)
		raw, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read test raw file: %s", path)
		}

		var vector schema.TestVector
		err = json.Unmarshal(raw, &vector)
		if err != nil {
			t.Errorf("failed to parse test vector %s: %s; skipping", path, err)
			continue
		}

		t.Run(v, func(t *testing.T) {
			for _, h := range vector.Hints {
				if h == schema.HintIncorrect {
					t.Logf("skipping vector marked as incorrect: %s", vector.Meta.ID)
					t.SkipNow()
				}
			}

			// dispatch the execution depending on the vector class.
			invokee, ok := invokees[vector.Class]
			if !ok {
				t.Fatalf("unsupported test vector class: %s", vector.Class)
			}

			for _, variant := range vector.Pre.Variants {
				variant := variant
				t.Run(variant.ID, func(t *testing.T) {
					_, _ = invokee(t, &vector, &variant) //nolint:errcheck
				})
			}
		})
	}
}
