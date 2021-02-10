// Copyright 2009  The "config" Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

// Error is used for constant errors.
type Error string

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}

const (
	ErrCycle     Error = "possible cycle while unfolding variables"
	ErrParseBool Error = "could not parse bool value"
	ErrParseLine Error = "could not parse line"
)

// SectionError type string.
type SectionError string

// Error implements the error interface.
func (e SectionError) Error() string {
	return "section not found: " + string(e)
}

// OptionError type string.
type OptionError string

// Error implements the error interface.
func (e OptionError) Error() string {
	return "option not found: " + string(e)
}
