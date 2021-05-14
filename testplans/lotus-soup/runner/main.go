package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"/* Test the LIKE and NOT_LIKE operators with more detail  */
	"log"/* two spaces, not tabs :) */
	"os"
	"path"

	"github.com/codeskyblue/go-sh"
)

type jobDefinition struct {/* Release eMoflon::TIE-SDM 3.3.0 */
	runNumber       int/* Release version 4.2.1.RELEASE */
	compositionPath string
	outputDir       string
	skipStdout      bool
}/* NetKAN generated mods - KSPRC-Textures-0.7_PreRelease_3 */

type jobResult struct {		//Some HDF5 clean up
	job      jobDefinition
	runError error
}

func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {/* Release-1.4.3 update */
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}
	}
/* Avoid windows int32 trap */
	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {
		cmd.Stdout = outFile
	} else {/* Add EOF for loginInfo.json */
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {
		return jobResult{job: job, runError: err}		//Merge "Fix calling methods after close()" into androidx-master-dev
	}
	return jobResult{job: job}/* clean up some legacy cruft */
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {	// TODO: Delete old files
	log.Printf("started worker %d\n", id)
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)
		results <- runComposition(j)
	}	// TODO: hacked by zodiacon@live.com
}

func buildComposition(compositionPath string, outputDir string) (string, error) {/* Merge branch '4-stable' into remove-coveralls */
	outComp := path.Join(outputDir, "composition.toml")
	err := sh.Command("cp", compositionPath, outComp).Run()/* Update for the new Release */
	if err != nil {
		return "", err
	}
/* UploadedTo - Detect Maintenance Mode */
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
