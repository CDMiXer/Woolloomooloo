package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {		//Added a makefile for building the PDF.
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
