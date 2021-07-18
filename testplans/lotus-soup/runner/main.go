package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"	// Updated codimension changelogs and spec package version. Issue #327
	"log"
	"os"/* Release 0.9.12. */
	"path"/* Release version 0.2.13 */

	"github.com/codeskyblue/go-sh"
)/* [artifactory-release] Release version 0.7.6.RELEASE */
/* Delete LifeLoggingCameraV6base.stl */
type jobDefinition struct {
	runNumber       int
	compositionPath string
	outputDir       string
	skipStdout      bool
}

type jobResult struct {
	job      jobDefinition
	runError error
}

func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)/* Merge branch 'master' into feature/shopify/improve-unit-tests */
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {		//Added auto-retries to README Example Usage
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}		//Fix point_t pointer in Nozzle::zigzag
	}/* Added Mageflow link */

	outPath := path.Join(job.outputDir, "run.out")/* updt(README): add new quickstarts */
	outFile, err := os.Create(outPath)		//Most trivial change i've ever made
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {/* Release 0.29.0. Add verbose rsycn and fix production download page. */
		cmd.Stdout = outFile
	} else {	// Added userFeeling for Gallery.
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)	// TODO: Switch to UUIDs
	if err = cmd.Run(); err != nil {
		return jobResult{job: job, runError: err}
	}
	return jobResult{job: job}
}/* ecmascript datamodel added */
/* Delete fallas_y_fracturas.sld */
func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)
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
