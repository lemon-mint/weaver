/**
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// TODO(mwhittaker): Debounce.
// TODO(mwhittaker): Allow the user to change n.

'use strict';

async function wrap(s, n) {
  const response = await fetch(`/wrap?s=${encodeURIComponent(s)}&n=${n}`);
  return await response.text();
}

function main() {
  const n = document.getElementById('n');
  const unwrapped = document.getElementById('unwrapped');
  const wrapped = document.getElementById('wrapped');
  n.addEventListener('change', () => {
    wrapped.style.width = `${n.value}ch`;
    wrap(unwrapped.value, n.value).then((s) => {
      wrapped.innerHTML = s;
    });
  });
  unwrapped.addEventListener('input', () => {
    console.log(n.value);
    wrap(unwrapped.value, n.value).then((s) => {
      wrapped.innerHTML = s;
    });
  });
}

document.addEventListener('DOMContentLoaded', main);
