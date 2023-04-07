// Copyright 2023 Google LLC
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
	"testing"
)

type intError struct{ x int }
type stringError struct{ x string }
type boolError struct{ b bool }

func (i intError) Error() string    { return fmt.Sprint(i.x) }
func (s stringError) Error() string { return s.x }
func (b boolError) Error() string   { return fmt.Sprint(b.b) }

func TestJoinErrorsSupportsErrorsAs(t *testing.T) {
	one := intError{1}
	a := stringError{"a"}
	joined := JoinErrors(one, a)

	var i intError
	if !errors.As(joined, &i) {
		t.Error("JoinErrors(one, a) is not an intError.")
	}

	var s stringError
	if !errors.As(joined, &s) {
		t.Error("JoinErrors(one, a) is not a stringError.")
	}

	var b boolError
	if errors.As(joined, &b) {
		t.Error("JoinErrors(one, a) is a boolError.")
	}
}

func TestJoinErrorsSupportsErrorsIs(t *testing.T) {
	a := fmt.Errorf("a")
	b := fmt.Errorf("b")
	c := fmt.Errorf("c")
	ab := JoinErrors(a, b)
	if !errors.Is(ab, a) {
		t.Error("JoinErrors(a, b) is not a.")
	}
	if !errors.Is(ab, b) {
		t.Error("JoinErrors(a, b) is not b.")
	}
	if errors.Is(ab, c) {
		t.Error("JoinErrors(a, b) is c.")
	}
}
