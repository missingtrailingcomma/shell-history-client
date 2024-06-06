load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/missingtrailingcomma/shell-history-client
gazelle(name = "gazelle")

go_library(
    name = "shell-history-client_lib",
    srcs = ["client.go"],
    importpath = "github.com/missingtrailingcomma/shell-history-client",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "shell-history-client",
    embed = [":shell-history-client_lib"],
    visibility = ["//visibility:public"],
)
