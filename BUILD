load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

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

gazelle(name = "gazelle")

go_library(
    name = "shell_history_client_lib",
    srcs = ["client_main.go"],
    importpath = "shell_history_client",
    visibility = ["//visibility:private"],
    deps = [
        "//cmd",
        "//data",
    ],
)

go_binary(
    name = "shell_history_client",
    embed = [":shell_history_client_lib"],
    visibility = ["//visibility:public"],
)
