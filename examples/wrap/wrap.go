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

package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/ServiceWeaver/weaver"
)

// The Wrapper component interface.
type Wrapper interface {
	Wrap(context.Context, string, int) (string, error)
}

// The Wrapper component implementation.
type wrapper struct {
	weaver.Implements[Wrapper]
}

// Wrap wraps the provided text to n characters.
func (wrapper) Wrap(_ context.Context, s string, n int) (string, error) {
	return wrap(s, n), nil
}

// wrap wraps the provided text to n characters.
func wrap(s string, n int) string {
	var b strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		fmt.Fprintln(&b, wrapline(scanner.Text(), n))
	}
	return b.String()
}

// wrapline wraps the provided line to n characters.
func wrapline(line string, n int) string {
	var b strings.Builder
	length := 0
	parts := split(line)
	for i, part := range parts {
		if length == 0 || length+len(part) <= n {
			b.WriteString(part)
			length += len(part)
		} else if unicode.IsSpace(firstRune(part)) {
			if i != len(parts)-1 {
				b.WriteString("\n")
				length = 0
			}
		} else {
			b.WriteString("\n")
			b.WriteString(part)
			length = len(part)
		}
	}
	return b.String()
}

// split splits s into abutting substrings of wholly whitespace or wholly
// non-whitespace characters. For example, split splits " foo \t bar " into
// " ", "foo", " \t ",  "bar", and " ".
func split(s string) []string {
	var parts []string
	var part []rune
	for _, r := range s {
		if len(part) == 0 || unicode.IsSpace(r) == unicode.IsSpace(part[0]) {
			part = append(part, r)
		} else {
			parts = append(parts, string(part))
			part = []rune{r}
		}
	}
	if len(part) > 0 {
		parts = append(parts, string(part))
	}
	return parts
}

// firstRune returns the first rune of s, or panics if s is empty.
func firstRune(s string) rune {
	for _, r := range s {
		return r
	}
	panic(fmt.Errorf("firstRune: empty string"))
}
