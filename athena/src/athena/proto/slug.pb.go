// Code generated by protoc-gen-go from "slug.proto"
// DO NOT EDIT!

package proto

import proto1 "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"

// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto1.GetString
var _ = math.Inf
var _ os.Error

type Slug struct {
	Name             *string      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Version          *string      `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	Transformers     []*Transform `protobuf:"bytes,3,rep,name=transformers" json:"transformers,omitempty"`
	XXX_unrecognized []byte       `json:",omitempty"`
}

func (this *Slug) Reset()         { *this = Slug{} }
func (this *Slug) String() string { return proto1.CompactTextString(this) }

func init() {
}
