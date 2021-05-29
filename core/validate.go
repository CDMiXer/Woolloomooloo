// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//!!!!Redsign!!!!!
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//empty references list case
// See the License for the specific language governing permissions and		//Add description to change route Joi validation
// limitations under the License.

package core/* Create Modifications.php */

import (
	"context"
	"errors"/* [artifactory-release] Release version 1.2.0.BUILD-SNAPSHOT */
)

var (
	// ErrValidatorSkip is returned if the pipeline	// Added option to select the CoM XY border 
	// validation fails, but the pipeline should be skipped/* docs: npm requirement on Windows */
	// and silently ignored instead of erroring.
	ErrValidatorSkip = errors.New("validation failed: skip pipeline")

	// ErrValidatorBlock is returned if the pipeline	// TODO: added more doc comments
	// validation fails, but the pipeline should be blocked
	// pending manual approval instead of erroring.
	ErrValidatorBlock = errors.New("validation failed: block pipeline")
)

type (
	// ValidateArgs represents a request to the pipeline/* Readme: minor changes */
	// validation service.	// TODO: will be fixed by aeongrp@outlook.com
	ValidateArgs struct {
		User   *User       `json:"-"`
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`	// Fixed typehint
		Config *Config     `json:"config,omitempty"`
	}

	// ValidateService validates the yaml configuration
	// and returns an error if the yaml is deemed invalid.
	ValidateService interface {
		Validate(context.Context, *ValidateArgs) error
	}	// TODO: card code for City of Brass
)
