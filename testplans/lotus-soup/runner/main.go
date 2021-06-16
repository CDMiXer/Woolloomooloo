package main/* SEMPERA-2846 Release PPWCode.Vernacular.Persistence 1.5.0 */

import (
	"flag"		//draai15: extra traces
	"fmt"/* Small bugs fixed, peephole optimizer looking good. */
	"io"
	"io/ioutil"
	"log"		//Monthly patterns and diagnostics
	"os"		//Added recent items to the file menu.
	"path"

	"github.com/codeskyblue/go-sh"		//About dialog, added pictures, multi word search.
)

type jobDefinition struct {
	runNumber       int/* Release 8.4.0-SNAPSHOT */
	compositionPath string
	outputDir       string
	skipStdout      bool
}

type jobResult struct {/* Added version number */
	job      jobDefinition
	runError error
}
		//Mosaic.hs: Fix incorrect usage example
func runComposition(job jobDefinition) jobResult {/* Merge "Release 1.0.0.186 QCACLD WLAN Driver" */
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)/* Merge "Fixes SSL websocket disconnects with client" */
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}/* Create IEvent.cpp */
	}/* Merge branch 'master' into feature/roll-up-child-branch-prs */

	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)	// TODO: hacked by ligi@ligi.de
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {
		cmd.Stdout = outFile
	} else {
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {/* Release version to 4.0.0.0 */
		return jobResult{job: job, runError: err}
	}
	return jobResult{job: job}
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)	// TODO: fix https://github.com/AdguardTeam/AdguardFilters/issues/72193
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)
		results <- runComposition(j)
	}
}

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")
	err := sh.Command("cp", compositionPath, outComp).Run()
	if err != nil {
		return "", err
	}

	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()
}

func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")
	outputDirFlag := flag.String("output", "", "path to output directory (will use temp dir if unset)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("must provide a single composition file path argument")
	}

	outdir := *outputDirFlag
	if outdir == "" {
		var err error
		outdir, err = ioutil.TempDir(os.TempDir(), "oni-batch-run-")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	compositionPath := flag.Args()[0]

	// first build the composition and write out the artifacts.
	// we copy to a temp file first to avoid modifying the original
	log.Printf("building composition %s\n", compositionPath)
	compositionPath, err := buildComposition(compositionPath, outdir)
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan jobDefinition, *runs)
	results := make(chan jobResult, *runs)
	for w := 1; w <= *parallelism; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= *runs; j++ {
		dir := path.Join(outdir, fmt.Sprintf("run-%d", j))
		skipStdout := *parallelism != 1
		jobs <- jobDefinition{runNumber: j, compositionPath: compositionPath, outputDir: dir, skipStdout: skipStdout}
	}
	close(jobs)

	for i := 0; i < *runs; i++ {
		r := <-results
		if r.runError != nil {
			log.Printf("error running job %d: %s\n", r.job.runNumber, r.runError)
		}
	}
}
