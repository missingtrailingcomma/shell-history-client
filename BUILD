load("@bazel_gazelle//:def.bzl", "gazelle")
load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

gazelle(name = "gazelle")

buildifier(
    name = "buildifier",
)

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

go_library(
    name = "shell_history_client_lib",
    srcs = ["client_main.go"],
    importpath = "shell_history_client",
    visibility = ["//visibility:private"],
    deps = [
        "//cmd",
        "//data",
        "//proto:command",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_binary(
    name = "shell_history_client",
    embed = [":shell_history_client_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "shell-history-client_lib",
    srcs = ["client_main.go"],
    importpath = "github.com/missingtrailingcomma/shell-history-client",
    visibility = ["//visibility:private"],
    deps = [
        "//proto",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)
