// Code generated by protoc-gen-go.
// source: tritium.proto
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto1.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Transform struct {
	Objects          []*ScriptObject `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty"`
	Pkg              *Package        `protobuf:"bytes,2,req,name=pkg" json:"pkg,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Transform) Reset()         { *m = Transform{} }
func (m *Transform) String() string { return proto1.CompactTextString(m) }
func (*Transform) ProtoMessage()    {}

func (m *Transform) GetObjects() []*ScriptObject {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *Transform) GetPkg() *Package {
	if m != nil {
		return m.Pkg
	}
	return nil
}

func init() {
}
