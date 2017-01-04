// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	ElevationQuery
*/
package message

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

type ElevationQuery struct {
	QueryList []*ElevationQuery_LatLonAlt `protobuf:"bytes,1,rep,name=query_list,json=queryList" json:"query_list,omitempty"`
}

func (m *ElevationQuery) Reset()                    { *m = ElevationQuery{} }
func (m *ElevationQuery) String() string            { return proto.CompactTextString(m) }
func (*ElevationQuery) ProtoMessage()               {}
func (*ElevationQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ElevationQuery) GetQueryList() []*ElevationQuery_LatLonAlt {
	if m != nil {
		return m.QueryList
	}
	return nil
}

type ElevationQuery_LatLonAlt struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
	Altitude  float64 `protobuf:"fixed64,3,opt,name=altitude" json:"altitude,omitempty"`
}

func (m *ElevationQuery_LatLonAlt) Reset()                    { *m = ElevationQuery_LatLonAlt{} }
func (m *ElevationQuery_LatLonAlt) String() string            { return proto.CompactTextString(m) }
func (*ElevationQuery_LatLonAlt) ProtoMessage()               {}
func (*ElevationQuery_LatLonAlt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ElevationQuery_LatLonAlt) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ElevationQuery_LatLonAlt) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *ElevationQuery_LatLonAlt) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func init() {
	proto.RegisterType((*ElevationQuery)(nil), "message.ElevationQuery")
	proto.RegisterType((*ElevationQuery_LatLonAlt)(nil), "message.ElevationQuery.LatLonAlt")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xb6, 0x32,
	0x72, 0xf1, 0xb9, 0xe6, 0xa4, 0x96, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x05, 0x96, 0xa6, 0x16, 0x55,
	0x0a, 0x39, 0x70, 0x71, 0x15, 0x82, 0x18, 0xf1, 0x39, 0x99, 0xc5, 0x25, 0x12, 0x8c, 0x0a, 0xcc,
	0x1a, 0xdc, 0x46, 0x8a, 0x7a, 0x30, 0xfd, 0xa8, 0x8a, 0xf5, 0x7c, 0x12, 0x4b, 0x7c, 0xf2, 0xf3,
	0x1c, 0x73, 0x4a, 0x82, 0x38, 0xc1, 0x9a, 0x7c, 0x32, 0x8b, 0x4b, 0xa4, 0x12, 0xb9, 0x38, 0xe1,
	0xe2, 0x42, 0x52, 0x5c, 0x1c, 0x39, 0x89, 0x25, 0x99, 0x25, 0xa5, 0x29, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x8c, 0x41, 0x70, 0xbe, 0x90, 0x0c, 0x17, 0x67, 0x4e, 0x7e, 0x5e, 0x3a, 0x44, 0x92,
	0x09, 0x2c, 0x89, 0x10, 0x00, 0xe9, 0x4c, 0xcc, 0x81, 0xea, 0x64, 0x86, 0xe8, 0x84, 0xf1, 0x93,
	0xd8, 0xc0, 0xfe, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xeb, 0x34, 0x3f, 0xd8, 0x00,
	0x00, 0x00,
}