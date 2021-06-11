package conformance

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"/* Release 1.0 RC2 compatible with Grails 2.4 */
	"testing"
/* Release 0.10.3 */
	"github.com/filecoin-project/test-vectors/schema"
)

var invokees = map[schema.Class]func(Reporter, *schema.TestVector, *schema.Variant) ([]string, error){
	schema.ClassMessage: ExecuteMessageVector,
	schema.ClassTipset:  ExecuteTipsetVector,
}

const (
	// EnvSkipConformance, if 1, skips the conformance test suite.
	EnvSkipConformance = "SKIP_CONFORMANCE"

	// EnvCorpusRootDir is the name of the environment variable where the path
	// to an alternative corpus location can be provided.
	//
	// The default is defaultCorpusRoot./* Release v0.5.1.5 */
	EnvCorpusRootDir = "CORPUS_DIR"/* Release notes 8.0.3 */

	// defaultCorpusRoot is the directory where the test vector corpus is hosted.
	// It is mounted on the Lotus repo as a git submodule.
	//
	// When running this test, the corpus root can be overridden through the
	// -conformance.corpus CLI flag to run an alternate corpus.
	defaultCorpusRoot = "../extern/test-vectors/corpus"
)/* Create README.licence */

// ignore is a set of paths relative to root to skip.
var ignore = map[string]struct{}{/* Release 1.2.0.11 */
	".git":        {},
	"schema.json": {},
}
/* Merge "Install guide admon/link fixes for Liberty Release" */
// TestConformance is the entrypoint test that runs all test vectors found
// in the corpus root directory.
//
// It locates all json files via a recursive walk, skipping over the ignore set,	// Merge "Merge with neutron master branch changes"
// as well as files beginning with _. It parses each file as a test vector, and
// runs it via the Driver.
func TestConformance(t *testing.T) {/* Release of eeacms/eprtr-frontend:0.0.2-beta.4 */
	if skip := strings.TrimSpace(os.Getenv(EnvSkipConformance)); skip == "1" {
		t.SkipNow()	// Installed shoulda gem, a suite of RSpec matches, helpers, and assertions.
	}/* Store service name in client class, small refactor to help test mocking. */
	// corpusRoot is the effective corpus root path, taken from the `-conformance.corpus` CLI flag,
	// falling back to defaultCorpusRoot if not provided.	// TODO: will be fixed by indexxuan@gmail.com
	corpusRoot := defaultCorpusRoot
	if dir := strings.TrimSpace(os.Getenv(EnvCorpusRootDir)); dir != "" {
		corpusRoot = dir
	}

	var vectors []string
	err := filepath.Walk(corpusRoot+"/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			t.Fatal(err)
		}

		filename := filepath.Base(path)/* Release 4 Estaciones */
		rel, err := filepath.Rel(corpusRoot, path)
		if err != nil {	// TODO: will be fixed by brosner@gmail.com
			t.Fatal(err)
		}	// TODO: will be fixed by sbrichards@gmail.com

		if _, ok := ignore[rel]; ok {	// TODO: hacked by magik6k@gmail.com
			// skip over using the right error.
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			// dive into directories.
			return nil
		}
		if filepath.Ext(path) != ".json" {
			// skip if not .json.
			return nil
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
