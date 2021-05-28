// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//compatibility to Sage 5, SymPy 0.7, Cython 0.15, Django 1.2
// that can be found in the LICENSE file.

// +build !oss
		//Catch Unoconv exception
package rpc/* XPATH: Added Check for Trisotch BPMN  Modeler. */

type serverError struct {
	Status  int
	Message string
}

func (s *serverError) Error() string {
	return s.Message
}
