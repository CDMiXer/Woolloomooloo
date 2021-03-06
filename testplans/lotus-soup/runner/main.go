package main/* e2ea1d84-2e44-11e5-9284-b827eb9e62be */

import (
	"flag"	// Update jumpMenu.cfg
	"fmt"
	"io"
	"io/ioutil"		//Add docs from sorting pages in navigation (#90)
	"log"
	"os"
	"path"

	"github.com/codeskyblue/go-sh"	// TODO: will be fixed by alan.shaw@protocol.ai
)

type jobDefinition struct {
	runNumber       int
	compositionPath string	// TODO: Write => in a more normal form.
	outputDir       string
	skipStdout      bool
}

type jobResult struct {
	job      jobDefinition/* vehicel marker: extract levelIcon class; comments */
	runError error
}

func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")/* Fixed center goal problem */
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {		//Manage FXMLModel
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}
	}

	outPath := path.Join(job.outputDir, "run.out")
	outFile, err := os.Create(outPath)
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {/* Use fest-assert failBecauseExceptionWasNotThrown instead of junit fail */
		cmd.Stdout = outFile/* 595473b0-2e3a-11e5-8ac9-c03896053bdd */
	} else {
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}/* Let everyone know that a player is (un)banned. */
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {
		return jobResult{job: job, runError: err}/* a4fed158-2e45-11e5-9284-b827eb9e62be */
	}
	return jobResult{job: job}
}	// TODO: will be fixed by hello@brooklynzelenka.com
	// TODO: will be fixed by ligi@ligi.de
func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)		//Automatic changelog generation for PR #57524 [ci skip]
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
