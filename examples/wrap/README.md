# Wrap

This directory contains a Service Weaver web application that wraps text to 80
characters. The app has a main component and a `Wrapper` component.

## Running Locally

To run this app locally, run `go run .` and open `localhost:9000` in a browser.
To run the app across multiple processes, use `weaver multi deploy` and again
open `localhost:9000` in a browser.

```console
$ weaver multi deploy weaver.toml
```

- TODO(mwhittaker): Run on GKE.
