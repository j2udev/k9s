// Code generated by protoc-gen-go. DO NOT EDIT.
// source: opencensus/proto/trace/v1/trace_config.proto

package v1 // import "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Global configuration of the trace service.
type TraceConfig struct {
	// The global default sampler used to make decisions on span sampling.
	//
	// Types that are valid to be assigned to Sampler:
	//	*TraceConfig_ProbabilitySampler
	//	*TraceConfig_ConstantSampler
	//	*TraceConfig_RateLimitingSampler
	Sampler isTraceConfig_Sampler `protobuf_oneof:"sampler"`
	// The global default max number of attributes per span.
	MaxNumberOfAttributes int64 `protobuf:"varint,4,opt,name=max_number_of_attributes,json=maxNumberOfAttributes,proto3" json:"max_number_of_attributes,omitempty"`
	// The global default max number of annotation events per span.
	MaxNumberOfAnnotations int64 `protobuf:"varint,5,opt,name=max_number_of_annotations,json=maxNumberOfAnnotations,proto3" json:"max_number_of_annotations,omitempty"`
	// The global default max number of message events per span.
	MaxNumberOfMessageEvents int64 `protobuf:"varint,6,opt,name=max_number_of_message_events,json=maxNumberOfMessageEvents,proto3" json:"max_number_of_message_events,omitempty"`
	// The global default max number of link entries per span.
	MaxNumberOfLinks     int64    `protobuf:"varint,7,opt,name=max_number_of_links,json=maxNumberOfLinks,proto3" json:"max_number_of_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_config_2f6d4a0dc67515b2, []int{0}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceConfig.Unmarshal(m, b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
}
func (dst *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(dst, src)
}
func (m *TraceConfig) XXX_Size() int {
	return xxx_messageInfo_TraceConfig.Size(m)
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

type isTraceConfig_Sampler interface {
	isTraceConfig_Sampler()
}

type TraceConfig_ProbabilitySampler struct {
	ProbabilitySampler *ProbabilitySampler `protobuf:"bytes,1,opt,name=probability_sampler,json=probabilitySampler,proto3,oneof"`
}

type TraceConfig_ConstantSampler struct {
	ConstantSampler *ConstantSampler `protobuf:"bytes,2,opt,name=constant_sampler,json=constantSampler,proto3,oneof"`
}

type TraceConfig_RateLimitingSampler struct {
	RateLimitingSampler *RateLimitingSampler `protobuf:"bytes,3,opt,name=rate_limiting_sampler,json=rateLimitingSampler,proto3,oneof"`
}

func (*TraceConfig_ProbabilitySampler) isTraceConfig_Sampler() {}

func (*TraceConfig_ConstantSampler) isTraceConfig_Sampler() {}

func (*TraceConfig_RateLimitingSampler) isTraceConfig_Sampler() {}

func (m *TraceConfig) GetSampler() isTraceConfig_Sampler {
	if m != nil {
		return m.Sampler
	}
	return nil
}

func (m *TraceConfig) GetProbabilitySampler() *ProbabilitySampler {
	if x, ok := m.GetSampler().(*TraceConfig_ProbabilitySampler); ok {
		return x.ProbabilitySampler
	}
	return nil
}

func (m *TraceConfig) GetConstantSampler() *ConstantSampler {
	if x, ok := m.GetSampler().(*TraceConfig_ConstantSampler); ok {
		return x.ConstantSampler
	}
	return nil
}

func (m *TraceConfig) GetRateLimitingSampler() *RateLimitingSampler {
	if x, ok := m.GetSampler().(*TraceConfig_RateLimitingSampler); ok {
		return x.RateLimitingSampler
	}
	return nil
}

func (m *TraceConfig) GetMaxNumberOfAttributes() int64 {
	if m != nil {
		return m.MaxNumberOfAttributes
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfAnnotations() int64 {
	if m != nil {
		return m.MaxNumberOfAnnotations
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfMessageEvents() int64 {
	if m != nil {
		return m.MaxNumberOfMessageEvents
	}
	return 0
}

func (m *TraceConfig) GetMaxNumberOfLinks() int64 {
	if m != nil {
		return m.MaxNumberOfLinks
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TraceConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TraceConfig_OneofMarshaler, _TraceConfig_OneofUnmarshaler, _TraceConfig_OneofSizer, []interface{}{
		(*TraceConfig_ProbabilitySampler)(nil),
		(*TraceConfig_ConstantSampler)(nil),
		(*TraceConfig_RateLimitingSampler)(nil),
	}
}

func _TraceConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TraceConfig)
	// sampler
	switch x := m.Sampler.(type) {
	case *TraceConfig_ProbabilitySampler:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProbabilitySampler); err != nil {
			return err
		}
	case *TraceConfig_ConstantSampler:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConstantSampler); err != nil {
			return err
		}
	case *TraceConfig_RateLimitingSampler:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RateLimitingSampler); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TraceConfig.Sampler has unexpected type %T", x)
	}
	return nil
}

func _TraceConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TraceConfig)
	switch tag {
	case 1: // sampler.probability_sampler
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProbabilitySampler)
		err := b.DecodeMessage(msg)
		m.Sampler = &TraceConfig_ProbabilitySampler{msg}
		return true, err
	case 2: // sampler.constant_sampler
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConstantSampler)
		err := b.DecodeMessage(msg)
		m.Sampler = &TraceConfig_ConstantSampler{msg}
		return true, err
	case 3: // sampler.rate_limiting_sampler
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RateLimitingSampler)
		err := b.DecodeMessage(msg)
		m.Sampler = &TraceConfig_RateLimitingSampler{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TraceConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TraceConfig)
	// sampler
	switch x := m.Sampler.(type) {
	case *TraceConfig_ProbabilitySampler:
		s := proto.Size(x.ProbabilitySampler)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TraceConfig_ConstantSampler:
		s := proto.Size(x.ConstantSampler)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TraceConfig_RateLimitingSampler:
		s := proto.Size(x.RateLimitingSampler)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Sampler that tries to uniformly sample traces with a given probability.
// The probability of sampling a trace is equal to that of the specified probability.
type ProbabilitySampler struct {
	// The desired probability of sampling. Must be within [0.0, 1.0].
	SamplingProbability  float64  `protobuf:"fixed64,1,opt,name=samplingProbability,proto3" json:"samplingProbability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbabilitySampler) Reset()         { *m = ProbabilitySampler{} }
func (m *ProbabilitySampler) String() string { return proto.CompactTextString(m) }
func (*ProbabilitySampler) ProtoMessage()    {}
func (*ProbabilitySampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_config_2f6d4a0dc67515b2, []int{1}
}
func (m *ProbabilitySampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbabilitySampler.Unmarshal(m, b)
}
func (m *ProbabilitySampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbabilitySampler.Marshal(b, m, deterministic)
}
func (dst *ProbabilitySampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbabilitySampler.Merge(dst, src)
}
func (m *ProbabilitySampler) XXX_Size() int {
	return xxx_messageInfo_ProbabilitySampler.Size(m)
}
func (m *ProbabilitySampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbabilitySampler.DiscardUnknown(m)
}

var xxx_messageInfo_ProbabilitySampler proto.InternalMessageInfo

func (m *ProbabilitySampler) GetSamplingProbability() float64 {
	if m != nil {
		return m.SamplingProbability
	}
	return 0
}

// Sampler that makes a constant decision (either always "yes" or always "no")
// on span sampling.
type ConstantSampler struct {
	// Whether spans should be always sampled, or never sampled.
	Decision             bool     `protobuf:"varint,1,opt,name=decision,proto3" json:"decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConstantSampler) Reset()         { *m = ConstantSampler{} }
func (m *ConstantSampler) String() string { return proto.CompactTextString(m) }
func (*ConstantSampler) ProtoMessage()    {}
func (*ConstantSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_config_2f6d4a0dc67515b2, []int{2}
}
func (m *ConstantSampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConstantSampler.Unmarshal(m, b)
}
func (m *ConstantSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConstantSampler.Marshal(b, m, deterministic)
}
func (dst *ConstantSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConstantSampler.Merge(dst, src)
}
func (m *ConstantSampler) XXX_Size() int {
	return xxx_messageInfo_ConstantSampler.Size(m)
}
func (m *ConstantSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_ConstantSampler.DiscardUnknown(m)
}

var xxx_messageInfo_ConstantSampler proto.InternalMessageInfo

func (m *ConstantSampler) GetDecision() bool {
	if m != nil {
		return m.Decision
	}
	return false
}

// Sampler that tries to sample with a rate per time window.
type RateLimitingSampler struct {
	// Rate per second.
	Qps                  int64    `protobuf:"varint,1,opt,name=qps,proto3" json:"qps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitingSampler) Reset()         { *m = RateLimitingSampler{} }
func (m *RateLimitingSampler) String() string { return proto.CompactTextString(m) }
func (*RateLimitingSampler) ProtoMessage()    {}
func (*RateLimitingSampler) Descriptor() ([]byte, []int) {
	return fileDescriptor_trace_config_2f6d4a0dc67515b2, []int{3}
}
func (m *RateLimitingSampler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitingSampler.Unmarshal(m, b)
}
func (m *RateLimitingSampler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitingSampler.Marshal(b, m, deterministic)
}
func (dst *RateLimitingSampler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitingSampler.Merge(dst, src)
}
func (m *RateLimitingSampler) XXX_Size() int {
	return xxx_messageInfo_RateLimitingSampler.Size(m)
}
func (m *RateLimitingSampler) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitingSampler.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitingSampler proto.InternalMessageInfo

func (m *RateLimitingSampler) GetQps() int64 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func init() {
	proto.RegisterType((*TraceConfig)(nil), "opencensus.proto.trace.v1.TraceConfig")
	proto.RegisterType((*ProbabilitySampler)(nil), "opencensus.proto.trace.v1.ProbabilitySampler")
	proto.RegisterType((*ConstantSampler)(nil), "opencensus.proto.trace.v1.ConstantSampler")
	proto.RegisterType((*RateLimitingSampler)(nil), "opencensus.proto.trace.v1.RateLimitingSampler")
}

func init() {
	proto.RegisterFile("opencensus/proto/trace/v1/trace_config.proto", fileDescriptor_trace_config_2f6d4a0dc67515b2)
}

var fileDescriptor_trace_config_2f6d4a0dc67515b2 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x6f, 0xd4, 0x30,
	0x10, 0x85, 0x09, 0x5b, 0xda, 0x32, 0x3d, 0x74, 0xe5, 0xa8, 0x28, 0x45, 0x3d, 0x54, 0x7b, 0xa1,
	0x42, 0x24, 0xa1, 0x70, 0x40, 0x5c, 0x90, 0xd8, 0x0a, 0xc4, 0xa1, 0xc0, 0x2a, 0x20, 0x21, 0x71,
	0x09, 0x4e, 0xea, 0x0d, 0x16, 0xf1, 0x38, 0xd8, 0x93, 0xa8, 0xfc, 0x17, 0x7e, 0x6c, 0x15, 0x67,
	0x95, 0x64, 0xbb, 0xed, 0xde, 0xec, 0xf7, 0xde, 0xf7, 0x2c, 0x79, 0x6c, 0x78, 0xa1, 0x2b, 0x81,
	0xb9, 0x40, 0x5b, 0xdb, 0xb8, 0x32, 0x9a, 0x74, 0x4c, 0x86, 0xe7, 0x22, 0x6e, 0xce, 0xbb, 0x45,
	0x9a, 0x6b, 0x5c, 0xca, 0x22, 0x72, 0x1e, 0x3b, 0x1e, 0xd2, 0x9d, 0x12, 0xb9, 0x50, 0xd4, 0x9c,
	0xcf, 0xfe, 0xef, 0xc0, 0xc1, 0xf7, 0x76, 0x73, 0xe1, 0x00, 0xf6, 0x0b, 0xfc, 0xca, 0xe8, 0x8c,
	0x67, 0xb2, 0x94, 0xf4, 0x2f, 0xb5, 0x5c, 0x55, 0xa5, 0x30, 0x81, 0x77, 0xea, 0x9d, 0x1d, 0xbc,
	0x0a, 0xa3, 0x7b, 0x8b, 0xa2, 0xc5, 0x40, 0x7d, 0xeb, 0xa0, 0x4f, 0x0f, 0x12, 0x56, 0x6d, 0xa8,
	0xec, 0x07, 0x4c, 0x73, 0x8d, 0x96, 0x38, 0x52, 0x5f, 0xff, 0xd0, 0xd5, 0x3f, 0xdf, 0x52, 0x7f,
	0xb1, 0x42, 0x86, 0xee, 0xc3, 0x7c, 0x5d, 0x62, 0x57, 0x70, 0x64, 0x38, 0x89, 0xb4, 0x94, 0x4a,
	0x92, 0xc4, 0xa2, 0x6f, 0x9f, 0xb8, 0xf6, 0x68, 0x4b, 0x7b, 0xc2, 0x49, 0x5c, 0xae, 0xb0, 0xe1,
	0x04, 0xdf, 0x6c, 0xca, 0xec, 0x0d, 0x04, 0x8a, 0x5f, 0xa7, 0x58, 0xab, 0x4c, 0x98, 0x54, 0x2f,
	0x53, 0x4e, 0x64, 0x64, 0x56, 0x93, 0xb0, 0xc1, 0xce, 0xa9, 0x77, 0x36, 0x49, 0x8e, 0x14, 0xbf,
	0xfe, 0xe2, 0xec, 0xaf, 0xcb, 0xf7, 0xbd, 0xc9, 0xde, 0xc2, 0xf1, 0x2d, 0x10, 0x51, 0x13, 0x27,
	0xa9, 0xd1, 0x06, 0x8f, 0x1c, 0xf9, 0x64, 0x4c, 0x0e, 0x2e, 0x7b, 0x07, 0x27, 0xeb, 0xa8, 0x12,
	0xd6, 0xf2, 0x42, 0xa4, 0xa2, 0x11, 0x48, 0x36, 0xd8, 0x75, 0x74, 0x30, 0xa2, 0x3f, 0x77, 0x81,
	0x0f, 0xce, 0x67, 0x21, 0xf8, 0xeb, 0x7c, 0x29, 0xf1, 0x8f, 0x0d, 0xf6, 0x1c, 0x36, 0x1d, 0x61,
	0x97, 0xad, 0x3e, 0x7f, 0x0c, 0x7b, 0xab, 0xab, 0x9b, 0x7d, 0x04, 0xb6, 0x39, 0x58, 0xf6, 0x12,
	0x7c, 0x17, 0x90, 0x58, 0x8c, 0x5c, 0xf7, 0x48, 0xbc, 0xe4, 0x2e, 0x6b, 0x16, 0xc2, 0xe1, 0xad,
	0x09, 0xb2, 0xa7, 0xb0, 0x7f, 0x25, 0x72, 0x69, 0xa5, 0x46, 0x47, 0xee, 0x27, 0xfd, 0x7e, 0xf6,
	0x0c, 0xfc, 0x3b, 0x46, 0xc2, 0xa6, 0x30, 0xf9, 0x5b, 0x59, 0x97, 0x9e, 0x24, 0xed, 0x72, 0xde,
	0xc0, 0x89, 0xd4, 0xf7, 0x0f, 0x76, 0x3e, 0x1d, 0xbd, 0xed, 0x45, 0x6b, 0x2d, 0xbc, 0x9f, 0xf3,
	0x42, 0xd2, 0xef, 0x3a, 0x8b, 0x72, 0xad, 0xe2, 0x8e, 0x0a, 0x25, 0x5a, 0x32, 0xb5, 0x12, 0xd8,
	0xdd, 0x78, 0x3c, 0x14, 0x86, 0xdd, 0xef, 0x2a, 0x04, 0x86, 0xc5, 0xf0, 0xc9, 0xb2, 0x5d, 0x27,
	0xbf, 0xbe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x98, 0xa1, 0x9e, 0x88, 0x03, 0x00, 0x00,
}