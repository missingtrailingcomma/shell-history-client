load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cmd",
    srcs = [
        "create.go",
        "web_portal.go",
    ],
    data = ["//web_portal"],
    importpath = "shell_history_client/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "//data",
        "//proto:command",
        "@org_golang_google_protobuf//encoding/prototext",
        "@org_golang_google_protobuf//proto",
    ],
)
