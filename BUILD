load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

sh_binary(
    name = "shell-history-client",
    srcs = ["shell-history-client.sh"],
    data = [
        "//:shell_history_client",
    ],
    deps = [
        ":util",
        "//dependencies/bash-preexec",
    ],
)

sh_library(
    name = "util",
    srcs = ["util.sh"],
)

# go mod tidy && bazel run //:gazelle -- update-repos -from_file=go.mod && bazel run //:gazelle
gazelle(name = "gazelle")

go_library(
    name = "shell_history_client_lib",
    srcs = ["client_main.go"],
    importpath = "shell_history_client",
    visibility = ["//visibility:private"],
    deps = [
        "//cmd",
        "//data",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_binary(
    name = "shell_history_client",
    embed = [":shell_history_client_lib"],
    visibility = ["//visibility:public"],
)
