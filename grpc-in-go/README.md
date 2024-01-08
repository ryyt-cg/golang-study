# gRPC in Go

## Goal

| Task # | Description                           | URL                                            |
|--------|---------------------------------------|------------------------------------------------|
| 1      | Converting multi-repos into mono-repo | https://github.com/industriousparadigm/go-grpc |
| 2      | Implementing order-payment service    |                                                |



## Implementation
* create grpc-in-go folder
* add WORKSPACE and BUILD.bazel

Note: adding bazel extension is optional.  They will treat the same in Bazel configuration.  For consistency, I'd 
like WORKSPACE with no extension and BUILD.bazel with extension throughout the project and packages.

<!-- TOC -->
* [gRPC in Go](#grpc-in-go)
  * [Goal](#goal)
  * [Implementation](#implementation)
  * [Create Protobuf Domain Package](#create-protobuf-domain-package)
    * [Add Order Protobuf](#add-order-protobuf)
  * [Configure Gazelle for go mod](#configure-gazelle-for-go-mod)
  * [Order service implementation](#order-service-implementation)
<!-- TOC -->

## Create Protobuf Domain Package
* create pb-domain
* add BUILD or BUILD.bazel empty file

### Add Order Protobuf
```protobuf
syntax = "proto3";

message Item {
  string name = 1;
}

message CreateOrderRequest {
  int64 user_id = 1;
  repeated Item items = 2;
  float total_price = 3;
}

message CreateOrderResponse {
  int64 order_id = 1;
}

service Order {
  rpc Create(CreateOrderRequest) returns (CreateOrderResponse){}
}
```

* Add proto dependency rules into WORKSPACE
* Add loading Proto Library into package to compile .proto files
* Define proto_library or libraries. 
* Add loading go proto library to compile proto library go Go
* Define go_proto_library + 1 proto library dependency

WORKSPACE
```build
## proto
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()
```

BUILD.bazel
```build
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "order_proto",
    srcs = ["order.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "order_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain",
    proto = ":order_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "pb-domain",
    embed = [":order_go_proto"],
    importpath = "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain",
    visibility = ["//visibility:public"],
)
```
After run bazel build //..., you would see some bazel-* directories created.
```shell
bazel build //...
```

## Configure Gazelle for go mod
* Add this to BUILD.bazel in root project

BUILD.bazel
```build
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix gitlab.com/aionx/bazel-demo-projects/beginning-bazel
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```

WORKSPACE
```build
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6ab6b57e48c09523e93050f13698f708428cfd5e619252e369d377af6597707",
    urls = [
        "https://mirror.bazel.build/github.cmd/bazelbuild/rules_go/releases/download/v0.43.0/rules_go-v0.43.0.zip",
        "https://github.cmd/bazelbuild/rules_go/releases/download/v0.43.0/rules_go-v0.43.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
    urls = [
        "https://mirror.bazel.build/github.cmd/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
        "https://github.cmd/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()
go_rules_dependencies()
go_register_toolchains(version = "1.21.3")

gazelle_dependencies()
```

After adding this code, you can run Gazelle with Bazel.
```shell
bazel run //:gazelle
```

If you get the error of deps.bzl not found because your go.mod has no dependencies, you may create it yourself and 
your build will succeed.

deps.bzl
```build
load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    pass
```

## Order service implementation
