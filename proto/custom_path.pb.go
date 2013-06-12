// Code generated by protoc-gen-go.
// source: custom_path.proto
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto1.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CustomPath struct {
	Origin           *string `protobuf:"bytes,1,req,name=origin" json:"origin,omitempty"`
	Target           *string `protobuf:"bytes,2,req,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CustomPath) Reset()         { *m = CustomPath{} }
func (m *CustomPath) String() string { return proto1.CompactTextString(m) }
func (*CustomPath) ProtoMessage()    {}

func (m *CustomPath) GetOrigin() string {
	if m != nil && m.Origin != nil {
		return *m.Origin
	}
	return ""
}

func (m *CustomPath) GetTarget() string {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return ""
}

func init() {
}
