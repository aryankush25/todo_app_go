# Todo API

This is a simple Todo API implemented in Go using the Gin framework.

## Understanding go.mod and go.sum

The `go.mod` file is used to manage project dependencies. It defines the module’s module path, which is its identity, and its dependency requirements, which are the other modules needed for a successful build.

The `go.sum` file is used to ensure that future downloads of these modules retrieve the same bits as the first download, to ensure the modules your project depends on do not change unexpectedly, whether for malicious, accidental, or other reasons.

## Running the Project

Before running the project, ensure that all dependencies are in sync:

```sh
go mod tidy
```

To run the project, use the following command:

```sh
go run .
```
