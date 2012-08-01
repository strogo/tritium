// Code generated by protoc-gen-go from "rewrite_rule.proto"
// DO NOT EDIT!

package proto

import proto1 "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

type RewriteRule_RuleDirection int32

const (
	RewriteRule_BIDIRECTIONAL     RewriteRule_RuleDirection = 0
	RewriteRule_PROXY_TO_UPSTREAM RewriteRule_RuleDirection = 1
	RewriteRule_UPSTREAM_TO_PROXY RewriteRule_RuleDirection = 2
)

var RewriteRule_RuleDirection_name = map[int32]string{
	0: "BIDIRECTIONAL",
	1: "PROXY_TO_UPSTREAM",
	2: "UPSTREAM_TO_PROXY",
}
var RewriteRule_RuleDirection_value = map[string]int32{
	"BIDIRECTIONAL":     0,
	"PROXY_TO_UPSTREAM": 1,
	"UPSTREAM_TO_PROXY": 2,
}

func (x RewriteRule_RuleDirection) Enum() *RewriteRule_RuleDirection {
	p := new(RewriteRule_RuleDirection)
	*p = x
	return p
}
func (x RewriteRule_RuleDirection) String() string {
	return proto1.EnumName(RewriteRule_RuleDirection_name, int32(x))
}

type RewriteRule struct {
	Proxy            *string                    `protobuf:"bytes,1,opt,name=proxy" json:"proxy,omitempty"`
	Upstream         *string                    `protobuf:"bytes,2,opt,name=upstream" json:"upstream,omitempty"`
	Direction        *RewriteRule_RuleDirection `protobuf:"varint,3,opt,name=direction,enum=proto.RewriteRule_RuleDirection,def=0" json:"direction,omitempty"`
	CookieDomain     *string                    `protobuf:"bytes,4,opt,name=cookie_domain" json:"cookie_domain,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (this *RewriteRule) Reset()         { *this = RewriteRule{} }
func (this *RewriteRule) String() string { return proto1.CompactTextString(this) }
func (*RewriteRule) ProtoMessage()       {}

const Default_RewriteRule_Direction RewriteRule_RuleDirection = RewriteRule_BIDIRECTIONAL

func (this *RewriteRule) GetProxy() string {
	if this != nil && this.Proxy != nil {
		return *this.Proxy
	}
	return ""
}

func (this *RewriteRule) GetUpstream() string {
	if this != nil && this.Upstream != nil {
		return *this.Upstream
	}
	return ""
}

func (this *RewriteRule) GetDirection() RewriteRule_RuleDirection {
	if this != nil && this.Direction != nil {
		return *this.Direction
	}
	return Default_RewriteRule_Direction
}

func (this *RewriteRule) GetCookieDomain() string {
	if this != nil && this.CookieDomain != nil {
		return *this.CookieDomain
	}
	return ""
}

func init() {
	proto1.RegisterEnum("proto.RewriteRule_RuleDirection", RewriteRule_RuleDirection_name, RewriteRule_RuleDirection_value)
}
