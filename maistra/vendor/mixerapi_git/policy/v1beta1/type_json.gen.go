// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1beta1/type.proto

// Describes the rules used to configure Mixer's policy and telemetry features.

package v1beta1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for Value
func (this *Value) MarshalJSON() ([]byte, error) {
	str, err := TypeMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for Value
func (this *Value) UnmarshalJSON(b []byte) error {
	return TypeUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TypeMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TypeUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
