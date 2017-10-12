#!/bin/bash

protoc -I../../../bazel-envoy/external/googleapis -I. --include_imports --include_source_info --descriptor_set_out=gen/kv.pb protos/kv.proto
