// Code generated by protoc-gen-go from "function.proto"
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type Function struct {
	Name             *string              `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description      *string              `protobuf:"bytes,11,opt,name=description" json:"description,omitempty"`
	ScopeTypeId      *int32               `protobuf:"varint,2,opt,name=scope_type_id" json:"scope_type_id,omitempty"`
	ScopeType        *string              `protobuf:"bytes,8,opt,name=scope_type" json:"scope_type,omitempty"`
	ReturnTypeId     *int32               `protobuf:"varint,3,opt,name=return_type_id" json:"return_type_id,omitempty"`
	ReturnType       *string              `protobuf:"bytes,9,opt,name=return_type" json:"return_type,omitempty"`
	OpensTypeId      *int32               `protobuf:"varint,4,opt,name=opens_type_id" json:"opens_type_id,omitempty"`
	OpensType        *string              `protobuf:"bytes,10,opt,name=opens_type" json:"opens_type,omitempty"`
	BuiltIn          *bool                `protobuf:"varint,5,opt,name=built_in" json:"built_in,omitempty"`
	Args             []*Function_Argument `protobuf:"bytes,6,rep,name=args" json:"args,omitempty"`
	Instruction      *Instruction         `protobuf:"bytes,7,opt,name=instruction" json:"instruction,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (this *Function) Reset()         { *this = Function{} }
func (this *Function) String() string { return proto1.CompactTextString(this) }
func (*Function) ProtoMessage()       {}

func (this *Function) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func (this *Function) GetDescription() string {
	if this != nil && this.Description != nil {
		return *this.Description
	}
	return ""
}

func (this *Function) GetScopeTypeId() int32 {
	if this != nil && this.ScopeTypeId != nil {
		return *this.ScopeTypeId
	}
	return 0
}

func (this *Function) GetScopeType() string {
	if this != nil && this.ScopeType != nil {
		return *this.ScopeType
	}
	return ""
}

func (this *Function) GetReturnTypeId() int32 {
	if this != nil && this.ReturnTypeId != nil {
		return *this.ReturnTypeId
	}
	return 0
}

func (this *Function) GetReturnType() string {
	if this != nil && this.ReturnType != nil {
		return *this.ReturnType
	}
	return ""
}

func (this *Function) GetOpensTypeId() int32 {
	if this != nil && this.OpensTypeId != nil {
		return *this.OpensTypeId
	}
	return 0
}

func (this *Function) GetOpensType() string {
	if this != nil && this.OpensType != nil {
		return *this.OpensType
	}
	return ""
}

func (this *Function) GetBuiltIn() bool {
	if this != nil && this.BuiltIn != nil {
		return *this.BuiltIn
	}
	return false
}

func (this *Function) GetInstruction() *Instruction {
	if this != nil {
		return this.Instruction
	}
	return nil
}

type Function_Argument struct {
	TypeId           *int32  `protobuf:"varint,1,opt,name=type_id" json:"type_id,omitempty"`
	TypeString       *string `protobuf:"bytes,2,opt,name=type_string" json:"type_string,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Function_Argument) Reset()         { *this = Function_Argument{} }
func (this *Function_Argument) String() string { return proto1.CompactTextString(this) }
func (*Function_Argument) ProtoMessage()       {}

func (this *Function_Argument) GetTypeId() int32 {
	if this != nil && this.TypeId != nil {
		return *this.TypeId
	}
	return 0
}

func (this *Function_Argument) GetTypeString() string {
	if this != nil && this.TypeString != nil {
		return *this.TypeString
	}
	return ""
}

func (this *Function_Argument) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

type FunctionArray struct {
	Functions        []*Function `protobuf:"bytes,1,rep,name=functions" json:"functions,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (this *FunctionArray) Reset()         { *this = FunctionArray{} }
func (this *FunctionArray) String() string { return proto1.CompactTextString(this) }
func (*FunctionArray) ProtoMessage()       {}

func init() {
}
