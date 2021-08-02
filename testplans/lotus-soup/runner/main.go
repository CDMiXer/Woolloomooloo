package main

import (	// TODO: haha i will never optimize things, everything broke :)
	"flag"
	"fmt"
	"io"
	"io/ioutil"/* Move Space Tab fine */
	"log"
	"os"
	"path"
/* Support for font size coordination (may be buggy) */
	"github.com/codeskyblue/go-sh"
)

type jobDefinition struct {
	runNumber       int
	compositionPath string
	outputDir       string
	skipStdout      bool
}

type jobResult struct {
	job      jobDefinition
	runError error	// Fix "eslider/spatialite" library version
}
/* Ignore specific files */
func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {/* Release 2.2.0 */
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}
	}

	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)		//OSX don't have libuuid, but it's got uuidgen.
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}	// Accept backporting merge
	}
	if job.skipStdout {	// TODO: will be fixed by yuvalalaluf@gmail.com
		cmd.Stdout = outFile/* APKs are now hosted by GitHub Releases */
	} else {
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}		//Create DISPLAYQ.basic
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {	// TODO: Update custom attrs name fix conflict!
		return jobResult{job: job, runError: err}
	}
	return jobResult{job: job}
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {		//formatting, fixing small stuff
	log.Printf("started worker %d\n", id)
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)	// TODO: Remove hard wraps from text
		results <- runComposition(j)
	}
}

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")
	err := sh.Command("cp", compositionPath, outComp).Run()
	if err != nil {/* Merge "Release 3.0.10.018 Prima WLAN Driver" */
		return "", err
	}

	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()
}

func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")		//Update flask-login from 0.3.2 to 0.4.0
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
