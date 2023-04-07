// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"errors"
	"fmt"
)

// CatchPanics recovers from panic() calls that occur during encoding,
// decoding, and RPC execution.
func CatchPanics(r interface{}) error {
	if r == nil {
		return nil
	}
	err, ok := r.(error)
	if !ok {
		panic(r)
	}
	if errors.As(err, &encoderError{}) || errors.As(err, &decoderError{}) {
		return err
	}
	panic(r)
}

// JoinErrors is a backport of [errors.Join][1].
//
// TODO(mwhittaker): This is obsoleted by errors.Join. If we drop support for
// go versions before go 1.20, we can remove this function.
//
// [1]: https://pkg.go.dev/errors#Join
func JoinErrors(err1, err2 error) error {
	return joinedErrors{err1, err2}
}

// joinedErrors is the combination of two errors.
type joinedErrors struct {
	err1 error
	err2 error
}

// Error implements the error interface.
func (j joinedErrors) Error() string {
	return fmt.Sprintf("%v: %v", j.err1, j.err2)
}

// Is makes joinedErrors compatible with errors.Is.
func (j joinedErrors) Is(err error) bool {
	return errors.Is(j.err1, err)
}

// As makes joinedErrors compatible with errors.Ass.
func (j joinedErrors) As(target any) bool {
	return errors.As(j.err1, target)
}

// Unwrap makes joinedErrors compatible with errors.Is, errors.As, and
// errors.Unwrap.
func (j joinedErrors) Unwrap() error {
	return j.err2
}
