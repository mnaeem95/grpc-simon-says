/* Copyright 2015 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
==============================================================================*/

// Code generated by protoc-gen-go.
// source: simonsays.proto
// DO NOT EDIT!

/*
Package simonsays is a generated protocol buffer package.

It is generated from these files:
	simonsays.proto

It has these top-level messages:
	Request
	Response
*/
package simonsays

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Color int32

const (
	Color_RED    Color = 0
	Color_GREEN  Color = 1
	Color_YELLOW Color = 2
	Color_BLUE   Color = 3
)

var Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "YELLOW",
	3: "BLUE",
}
var Color_value = map[string]int32{
	"RED":    0,
	"GREEN":  1,
	"YELLOW": 2,
	"BLUE":   3,
}

func (x Color) String() string {
	return proto.EnumName(Color_name, int32(x))
}
func (Color) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response_State int32

const (
	Response_BEGIN      Response_State = 0
	Response_START_TURN Response_State = 1
	Response_STOP_TURN  Response_State = 2
	Response_WIN        Response_State = 3
	Response_LOSE       Response_State = 4
)

var Response_State_name = map[int32]string{
	0: "BEGIN",
	1: "START_TURN",
	2: "STOP_TURN",
	3: "WIN",
	4: "LOSE",
}
var Response_State_value = map[string]int32{
	"BEGIN":      0,
	"START_TURN": 1,
	"STOP_TURN":  2,
	"WIN":        3,
	"LOSE":       4,
}

func (x Response_State) String() string {
	return proto.EnumName(Response_State_name, int32(x))
}
func (Response_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Request struct {
	// Types that are valid to be assigned to Event:
	//	*Request_Join
	//	*Request_Press
	Event isRequest_Event `protobuf_oneof:"event"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isRequest_Event interface {
	isRequest_Event()
}

type Request_Join struct {
	Join *Request_Player `protobuf:"bytes,1,opt,name=join,oneof"`
}
type Request_Press struct {
	Press Color `protobuf:"varint,2,opt,name=press,enum=simonsays.Color,oneof"`
}

func (*Request_Join) isRequest_Event()  {}
func (*Request_Press) isRequest_Event() {}

func (m *Request) GetEvent() isRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Request) GetJoin() *Request_Player {
	if x, ok := m.GetEvent().(*Request_Join); ok {
		return x.Join
	}
	return nil
}

func (m *Request) GetPress() Color {
	if x, ok := m.GetEvent().(*Request_Press); ok {
		return x.Press
	}
	return Color_RED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Request) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Request_OneofMarshaler, _Request_OneofUnmarshaler, _Request_OneofSizer, []interface{}{
		(*Request_Join)(nil),
		(*Request_Press)(nil),
	}
}

func _Request_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Request)
	// event
	switch x := m.Event.(type) {
	case *Request_Join:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Join); err != nil {
			return err
		}
	case *Request_Press:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Press))
	case nil:
	default:
		return fmt.Errorf("Request.Event has unexpected type %T", x)
	}
	return nil
}

func _Request_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Request)
	switch tag {
	case 1: // event.join
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request_Player)
		err := b.DecodeMessage(msg)
		m.Event = &Request_Join{msg}
		return true, err
	case 2: // event.press
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Event = &Request_Press{Color(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Request_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Request)
	// event
	switch x := m.Event.(type) {
	case *Request_Join:
		s := proto.Size(x.Join)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Request_Press:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Press))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A Player of the Simon says game.
type Request_Player struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Request_Player) Reset()                    { *m = Request_Player{} }
func (m *Request_Player) String() string            { return proto.CompactTextString(m) }
func (*Request_Player) ProtoMessage()               {}
func (*Request_Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Response struct {
	// Types that are valid to be assigned to Event:
	//	*Response_Turn
	//	*Response_Lightup
	Event isResponse_Event `protobuf_oneof:"event"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isResponse_Event interface {
	isResponse_Event()
}

type Response_Turn struct {
	Turn Response_State `protobuf:"varint,1,opt,name=turn,enum=simonsays.Response_State,oneof"`
}
type Response_Lightup struct {
	Lightup Color `protobuf:"varint,2,opt,name=lightup,enum=simonsays.Color,oneof"`
}

func (*Response_Turn) isResponse_Event()    {}
func (*Response_Lightup) isResponse_Event() {}

func (m *Response) GetEvent() isResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Response) GetTurn() Response_State {
	if x, ok := m.GetEvent().(*Response_Turn); ok {
		return x.Turn
	}
	return Response_BEGIN
}

func (m *Response) GetLightup() Color {
	if x, ok := m.GetEvent().(*Response_Lightup); ok {
		return x.Lightup
	}
	return Color_RED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Response) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Response_OneofMarshaler, _Response_OneofUnmarshaler, _Response_OneofSizer, []interface{}{
		(*Response_Turn)(nil),
		(*Response_Lightup)(nil),
	}
}

func _Response_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Response)
	// event
	switch x := m.Event.(type) {
	case *Response_Turn:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Turn))
	case *Response_Lightup:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Lightup))
	case nil:
	default:
		return fmt.Errorf("Response.Event has unexpected type %T", x)
	}
	return nil
}

func _Response_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Response)
	switch tag {
	case 1: // event.turn
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Event = &Response_Turn{Response_State(x)}
		return true, err
	case 2: // event.lightup
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Event = &Response_Lightup{Color(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Response_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Response)
	// event
	switch x := m.Event.(type) {
	case *Response_Turn:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Turn))
	case *Response_Lightup:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Lightup))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Request)(nil), "simonsays.Request")
	proto.RegisterType((*Request_Player)(nil), "simonsays.Request.Player")
	proto.RegisterType((*Response)(nil), "simonsays.Response")
	proto.RegisterEnum("simonsays.Color", Color_name, Color_value)
	proto.RegisterEnum("simonsays.Response_State", Response_State_name, Response_State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SimonSays service

type SimonSaysClient interface {
	//
	// Game process is the following:
	//
	// A Join Request should be sent to the game. This tells it
	// to join a game (or start a new one if one isn't already waiting on a game).
	//
	// The Response stream will send through a BEGIN Response.State to let you know that
	// the Game has been started.
	//
	// When the player recieves a START_TURN response, the server can take your input for the turn.
	//
	// When the player recieves a STOP_TURN response, the serer is no longer taking input for the turn.
	//
	// A WIN state says you won the game. A LOSE state means that you got an input wrong, and have lost.
	//
	// To send input, send a Request with an event type of Color.
	//
	// When you recieve a Response of type Color, then light up that colour.
	Game(ctx context.Context, opts ...grpc.CallOption) (SimonSays_GameClient, error)
}

type simonSaysClient struct {
	cc *grpc.ClientConn
}

func NewSimonSaysClient(cc *grpc.ClientConn) SimonSaysClient {
	return &simonSaysClient{cc}
}

func (c *simonSaysClient) Game(ctx context.Context, opts ...grpc.CallOption) (SimonSays_GameClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SimonSays_serviceDesc.Streams[0], c.cc, "/simonsays.SimonSays/Game", opts...)
	if err != nil {
		return nil, err
	}
	x := &simonSaysGameClient{stream}
	return x, nil
}

type SimonSays_GameClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type simonSaysGameClient struct {
	grpc.ClientStream
}

func (x *simonSaysGameClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simonSaysGameClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SimonSays service

type SimonSaysServer interface {
	//
	// Game process is the following:
	//
	// A Join Request should be sent to the game. This tells it
	// to join a game (or start a new one if one isn't already waiting on a game).
	//
	// The Response stream will send through a BEGIN Response.State to let you know that
	// the Game has been started.
	//
	// When the player recieves a START_TURN response, the server can take your input for the turn.
	//
	// When the player recieves a STOP_TURN response, the serer is no longer taking input for the turn.
	//
	// A WIN state says you won the game. A LOSE state means that you got an input wrong, and have lost.
	//
	// To send input, send a Request with an event type of Color.
	//
	// When you recieve a Response of type Color, then light up that colour.
	Game(SimonSays_GameServer) error
}

func RegisterSimonSaysServer(s *grpc.Server, srv SimonSaysServer) {
	s.RegisterService(&_SimonSays_serviceDesc, srv)
}

func _SimonSays_Game_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimonSaysServer).Game(&simonSaysGameServer{stream})
}

type SimonSays_GameServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type simonSaysGameServer struct {
	grpc.ServerStream
}

func (x *simonSaysGameServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simonSaysGameServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SimonSays_serviceDesc = grpc.ServiceDesc{
	ServiceName: "simonsays.SimonSays",
	HandlerType: (*SimonSaysServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Game",
			Handler:       _SimonSays_Game_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "simonsays.proto",
}

func init() { proto.RegisterFile("simonsays.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4e, 0xab, 0x30,
	0x18, 0xc7, 0x29, 0x83, 0x31, 0xbe, 0x93, 0xc3, 0x69, 0x7a, 0x6e, 0x76, 0x76, 0xb5, 0x70, 0xc5,
	0x31, 0x06, 0x75, 0xc6, 0x07, 0x10, 0x47, 0xe6, 0x12, 0xb2, 0x2d, 0x85, 0x65, 0xf1, 0xca, 0xa0,
	0x6b, 0x26, 0x86, 0x51, 0xa4, 0x9d, 0x71, 0x4f, 0xe0, 0x2b, 0xf9, 0x78, 0xa6, 0xa0, 0x9b, 0x51,
	0xe3, 0x65, 0x9b, 0xdf, 0xff, 0xff, 0xfd, 0xda, 0x0f, 0xfe, 0x88, 0x6c, 0xcd, 0x0b, 0x91, 0x6e,
	0x85, 0x5f, 0x56, 0x5c, 0x72, 0x62, 0xef, 0x2e, 0xdc, 0x67, 0x04, 0x16, 0x65, 0x0f, 0x1b, 0x26,
	0x24, 0x39, 0x02, 0xe3, 0x9e, 0x67, 0x45, 0x17, 0xf5, 0x91, 0xf7, 0x6b, 0xf0, 0xcf, 0xdf, 0xc7,
	0xde, 0x08, 0x7f, 0x96, 0xa7, 0x5b, 0x56, 0x5d, 0x6a, 0xb4, 0x06, 0x89, 0x07, 0x66, 0x59, 0x31,
	0x21, 0xba, 0x7a, 0x1f, 0x79, 0xce, 0x00, 0x7f, 0x48, 0x5c, 0xf0, 0x9c, 0x2b, 0xb0, 0x01, 0x7a,
	0x5d, 0x68, 0x37, 0x59, 0xe2, 0x80, 0x9e, 0x2d, 0xeb, 0x11, 0x36, 0xd5, 0xb3, 0x65, 0x60, 0x81,
	0xc9, 0x1e, 0x59, 0x21, 0xdd, 0x17, 0x04, 0x1d, 0xca, 0x44, 0xc9, 0x0b, 0xc1, 0x94, 0x8a, 0xdc,
	0x54, 0x8d, 0x8a, 0xf3, 0x49, 0xa5, 0x41, 0xfc, 0x58, 0xa6, 0x92, 0x29, 0x15, 0x05, 0x92, 0x43,
	0xb0, 0xf2, 0x6c, 0x75, 0x27, 0x37, 0xe5, 0x0f, 0x32, 0xef, 0x88, 0x3b, 0x04, 0xb3, 0x8e, 0x13,
	0x1b, 0xcc, 0x20, 0x1c, 0x8d, 0x27, 0x58, 0x23, 0x0e, 0x40, 0x9c, 0x9c, 0xd3, 0xe4, 0x3a, 0x99,
	0xd3, 0x09, 0x46, 0xe4, 0x37, 0xd8, 0x71, 0x32, 0x9d, 0x35, 0x47, 0x9d, 0x58, 0xd0, 0x5a, 0x8c,
	0x27, 0xb8, 0x45, 0x3a, 0x60, 0x44, 0xd3, 0x38, 0xc4, 0xc6, 0x4e, 0xfd, 0xe0, 0x04, 0xcc, 0x7a,
	0x84, 0x82, 0x68, 0x38, 0xc4, 0x9a, 0xea, 0x1d, 0xd1, 0x30, 0x54, 0x3d, 0x00, 0xed, 0xab, 0x30,
	0x8a, 0xa6, 0x0b, 0xac, 0xab, 0x6c, 0x10, 0xcd, 0x43, 0xdc, 0x1a, 0x04, 0x60, 0xc7, 0xca, 0x2f,
	0x4e, 0xb7, 0x82, 0x9c, 0x81, 0x31, 0x4a, 0xd7, 0x8c, 0x90, 0xaf, 0x5f, 0xde, 0xfb, 0xfb, 0xcd,
	0xdb, 0x5d, 0xcd, 0x43, 0xc7, 0x28, 0xf8, 0x0f, 0xbd, 0x8c, 0xfb, 0xab, 0xaa, 0xbc, 0xf5, 0xd9,
	0x53, 0xba, 0x2e, 0x73, 0x26, 0xf6, 0x70, 0xb0, 0xef, 0x9f, 0xa1, 0x9b, 0x76, 0xbd, 0xf8, 0xd3,
	0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe1, 0xba, 0x9a, 0x0b, 0x02, 0x00, 0x00,
}
