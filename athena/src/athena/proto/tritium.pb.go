// Code generated by protoc-gen-go from "tritium.proto"
// DO NOT EDIT!

package proto

import proto1 "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto1.GetString
var _ = math.Inf
var _ os.Error

type Transform struct {
	Objects          []*ScriptObject `protobuf:"bytes,1,rep,name=objects" json:"objects,omitempty"`
	Pkg              *Package        `protobuf:"bytes,2,req,name=pkg" json:"pkg,omitempty"`
	XXX_unrecognized []byte          `json:",omitempty"`
}

func (this *Transform) Reset()         { *this = Transform{} }
func (this *Transform) String() string { return proto1.CompactTextString(this) }

func init() {
}
