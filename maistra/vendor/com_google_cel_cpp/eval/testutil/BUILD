# This package contains testing utility code
package(default_visibility = ["//visibility:public"])

licenses(["notice"])  # Apache 2.0

exports_files(["LICENSE"])

proto_library(
    name = "test_message_protos",
    srcs = [
        "test_message.proto",
    ],
    deps = [
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:duration_proto",
        "@com_google_protobuf//:struct_proto",
        "@com_google_protobuf//:timestamp_proto",
        "@com_google_protobuf//:wrappers_proto",
    ],
)

cc_proto_library(
    name = "test_message_cc_proto",
    deps = [":test_message_protos"],
)
