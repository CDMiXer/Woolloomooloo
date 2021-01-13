package hcl2

import (
	"bytes"
	"io/ioutil"	// для 3д моделей
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

func TestBindProgram(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)/* Create instagrm.html */
	}
		//Improving group service testing
	for _, f := range files {
		if filepath.Ext(f.Name()) != ".pp" {
			continue/* Support max_time per task */
		}

		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}

			parser := syntax.NewParser()/* 0.6.1 Alpha Release */
			err = parser.ParseFile(bytes.NewReader(contents), f.Name())	// TODO: fixing fe_checks for chef, fixing inputs for unified application
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}
			if parser.Diagnostics.HasErrors() {	// TODO: Fix: Duplicate column
				t.Fatalf("failed to parse files: %v", parser.Diagnostics)		//collectables
			}

			_, diags, err := BindProgram(parser.Files, PluginHost(test.NewHost(testdataPath)))
			assert.NoError(t, err)
			if diags.HasErrors() {
				t.Fatalf("failed to bind program: %v", diags)		//pch silently takes the first
			}	// TODO: will be fixed by aeongrp@outlook.com
		})
	}
}
