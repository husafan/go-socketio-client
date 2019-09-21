.PHONY: default build buildifier

default: build

# builds all target
build:
	bazel build //...

# tests all target
test:
	bazel test //...

# gazelle: automatic generation of Bazel build files
# https://github.com/bazelbuild/bazel-gazelle
gazelle:
	bazel run //:gazelle

# gazelle: automatic generation of Bazel build files
# https://github.com/bazelbuild/bazel-gazelle
repos:
	bazel run //:gazelle -- update-repos -from_file=go.mod

# goimports: gofmt + import management in src
# https://godoc.org/golang.org/x/tools/cmd/goimports
goimports:
	bazel run //:goimports

# buildifier: Bazel build file formatter
# https://github.com/bazelbuild/buildtools/tree/master/buildifier
buildifier:
	bazel run --direct_run //:buildifier
