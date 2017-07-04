# GopherJS bindings for Improbable's gRPC-Web implementation

[![GoDoc](https://godoc.org/github.com/johanbrandhorst/grpcweb?status.svg)](https://godoc.org/github.com/johanbrandhorst/grpcweb)

This package provides GopherJS bindings for [Improbable's gRPC-web implementation](https://github.com/improbable-eng/grpc-web/).

The API is still experimental, and is not currently intended for general use outside
of via the [GopherJS protoc compiler plugin](https://github.com/johanbrandhorst/protoc-gen-gopherjs). See the `protoc-gen-gopherjs`
README for more information on generating the interface.

## Contributions
Contributions are very welcome, please submit issues and PRs for review.

## An example
See [the example repo](https://github.com/johanbrandhorst/grpcweb-example) for an
example use of this package together with `protoc-gen-gopherjs`.