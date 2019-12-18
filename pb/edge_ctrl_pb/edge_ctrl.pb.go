/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edge_ctrl.proto

package edge_ctrl_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContentType int32

const (
	ContentType_Zero                    ContentType = 0
	ContentType_ServerHelloType         ContentType = 20000
	ContentType_ClientHelloType         ContentType = 20001
	ContentType_ErrorType               ContentType = 20002
	ContentType_SessionAddedType        ContentType = 20100
	ContentType_SessionUpdatedType      ContentType = 20101
	ContentType_SessionRemovedType      ContentType = 20102
	ContentType_ApiSessionAddedType     ContentType = 20200
	ContentType_ApiSessionUpdatedType   ContentType = 20201
	ContentType_ApiSessionRemovedType   ContentType = 20202
	ContentType_ApiSessionHeartbeatType ContentType = 20203
	ContentType_EnrollType              ContentType = 20300
	ContentType_EnrollCertsType         ContentType = 20301
)

var ContentType_name = map[int32]string{
	0:     "Zero",
	20000: "ServerHelloType",
	20001: "ClientHelloType",
	20002: "ErrorType",
	20100: "SessionAddedType",
	20101: "SessionUpdatedType",
	20102: "SessionRemovedType",
	20200: "ApiSessionAddedType",
	20201: "ApiSessionUpdatedType",
	20202: "ApiSessionRemovedType",
	20203: "ApiSessionHeartbeatType",
	20300: "EnrollType",
	20301: "EnrollCertsType",
}

var ContentType_value = map[string]int32{
	"Zero":                    0,
	"ServerHelloType":         20000,
	"ClientHelloType":         20001,
	"ErrorType":               20002,
	"SessionAddedType":        20100,
	"SessionUpdatedType":      20101,
	"SessionRemovedType":      20102,
	"ApiSessionAddedType":     20200,
	"ApiSessionUpdatedType":   20201,
	"ApiSessionRemovedType":   20202,
	"ApiSessionHeartbeatType": 20203,
	"EnrollType":              20300,
	"EnrollCertsType":         20301,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}

func (ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{0}
}

type ServerHello struct {
	Version              string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServerHello) Reset()         { *m = ServerHello{} }
func (m *ServerHello) String() string { return proto.CompactTextString(m) }
func (*ServerHello) ProtoMessage()    {}
func (*ServerHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{0}
}

func (m *ServerHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerHello.Unmarshal(m, b)
}
func (m *ServerHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerHello.Marshal(b, m, deterministic)
}
func (m *ServerHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerHello.Merge(m, src)
}
func (m *ServerHello) XXX_Size() int {
	return xxx_messageInfo_ServerHello.Size(m)
}
func (m *ServerHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerHello.DiscardUnknown(m)
}

var xxx_messageInfo_ServerHello proto.InternalMessageInfo

func (m *ServerHello) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerHello) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type ClientHello struct {
	Version              string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Hostname             string            `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Protocols            []string          `protobuf:"bytes,3,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClientHello) Reset()         { *m = ClientHello{} }
func (m *ClientHello) String() string { return proto.CompactTextString(m) }
func (*ClientHello) ProtoMessage()    {}
func (*ClientHello) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{1}
}

func (m *ClientHello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientHello.Unmarshal(m, b)
}
func (m *ClientHello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientHello.Marshal(b, m, deterministic)
}
func (m *ClientHello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientHello.Merge(m, src)
}
func (m *ClientHello) XXX_Size() int {
	return xxx_messageInfo_ClientHello.Size(m)
}
func (m *ClientHello) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientHello.DiscardUnknown(m)
}

var xxx_messageInfo_ClientHello proto.InternalMessageInfo

func (m *ClientHello) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClientHello) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ClientHello) GetProtocols() []string {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *ClientHello) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Cause                string   `protobuf:"bytes,3,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

type Dns struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Protocol             string   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dns) Reset()         { *m = Dns{} }
func (m *Dns) String() string { return proto.CompactTextString(m) }
func (*Dns) ProtoMessage()    {}
func (*Dns) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{3}
}

func (m *Dns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dns.Unmarshal(m, b)
}
func (m *Dns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dns.Marshal(b, m, deterministic)
}
func (m *Dns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dns.Merge(m, src)
}
func (m *Dns) XXX_Size() int {
	return xxx_messageInfo_Dns.Size(m)
}
func (m *Dns) XXX_DiscardUnknown() {
	xxx_messageInfo_Dns.DiscardUnknown(m)
}

var xxx_messageInfo_Dns proto.InternalMessageInfo

func (m *Dns) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Dns) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Dns) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type Service struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SupportedProtocol    string   `protobuf:"bytes,4,opt,name=supportedProtocol,proto3" json:"supportedProtocol,omitempty"`
	EgressRouter         string   `protobuf:"bytes,5,opt,name=egressRouter,proto3" json:"egressRouter,omitempty"`
	EndpointAddress      string   `protobuf:"bytes,6,opt,name=endpointAddress,proto3" json:"endpointAddress,omitempty"`
	Dns                  *Dns     `protobuf:"bytes,8,opt,name=dns,proto3" json:"dns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{4}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetSupportedProtocol() string {
	if m != nil {
		return m.SupportedProtocol
	}
	return ""
}

func (m *Service) GetEgressRouter() string {
	if m != nil {
		return m.EgressRouter
	}
	return ""
}

func (m *Service) GetEndpointAddress() string {
	if m != nil {
		return m.EndpointAddress
	}
	return ""
}

func (m *Service) GetDns() *Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

type Session struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Hosting              bool     `protobuf:"varint,2,opt,name=hosting,proto3" json:"hosting,omitempty"`
	CertFingerprints     []string `protobuf:"bytes,3,rep,name=certFingerprints,proto3" json:"certFingerprints,omitempty"`
	Urls                 []string `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	Service              *Service `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	SessionToken         string   `protobuf:"bytes,6,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{5}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Session) GetHosting() bool {
	if m != nil {
		return m.Hosting
	}
	return false
}

func (m *Session) GetCertFingerprints() []string {
	if m != nil {
		return m.CertFingerprints
	}
	return nil
}

func (m *Session) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *Session) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Session) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type ApiSession struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CertFingerprints     []string `protobuf:"bytes,2,rep,name=certFingerprints,proto3" json:"certFingerprints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiSession) Reset()         { *m = ApiSession{} }
func (m *ApiSession) String() string { return proto.CompactTextString(m) }
func (*ApiSession) ProtoMessage()    {}
func (*ApiSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{6}
}

func (m *ApiSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSession.Unmarshal(m, b)
}
func (m *ApiSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSession.Marshal(b, m, deterministic)
}
func (m *ApiSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSession.Merge(m, src)
}
func (m *ApiSession) XXX_Size() int {
	return xxx_messageInfo_ApiSession.Size(m)
}
func (m *ApiSession) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSession.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSession proto.InternalMessageInfo

func (m *ApiSession) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ApiSession) GetCertFingerprints() []string {
	if m != nil {
		return m.CertFingerprints
	}
	return nil
}

type ApiSessionAdded struct {
	IsFullState          bool          `protobuf:"varint,1,opt,name=isFullState,proto3" json:"isFullState,omitempty"`
	ApiSessions          []*ApiSession `protobuf:"bytes,2,rep,name=apiSessions,proto3" json:"apiSessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ApiSessionAdded) Reset()         { *m = ApiSessionAdded{} }
func (m *ApiSessionAdded) String() string { return proto.CompactTextString(m) }
func (*ApiSessionAdded) ProtoMessage()    {}
func (*ApiSessionAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{7}
}

func (m *ApiSessionAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSessionAdded.Unmarshal(m, b)
}
func (m *ApiSessionAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSessionAdded.Marshal(b, m, deterministic)
}
func (m *ApiSessionAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSessionAdded.Merge(m, src)
}
func (m *ApiSessionAdded) XXX_Size() int {
	return xxx_messageInfo_ApiSessionAdded.Size(m)
}
func (m *ApiSessionAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSessionAdded.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSessionAdded proto.InternalMessageInfo

func (m *ApiSessionAdded) GetIsFullState() bool {
	if m != nil {
		return m.IsFullState
	}
	return false
}

func (m *ApiSessionAdded) GetApiSessions() []*ApiSession {
	if m != nil {
		return m.ApiSessions
	}
	return nil
}

type ApiSessionUpdated struct {
	ApiSessions          []*ApiSession `protobuf:"bytes,1,rep,name=apiSessions,proto3" json:"apiSessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ApiSessionUpdated) Reset()         { *m = ApiSessionUpdated{} }
func (m *ApiSessionUpdated) String() string { return proto.CompactTextString(m) }
func (*ApiSessionUpdated) ProtoMessage()    {}
func (*ApiSessionUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{8}
}

func (m *ApiSessionUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSessionUpdated.Unmarshal(m, b)
}
func (m *ApiSessionUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSessionUpdated.Marshal(b, m, deterministic)
}
func (m *ApiSessionUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSessionUpdated.Merge(m, src)
}
func (m *ApiSessionUpdated) XXX_Size() int {
	return xxx_messageInfo_ApiSessionUpdated.Size(m)
}
func (m *ApiSessionUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSessionUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSessionUpdated proto.InternalMessageInfo

func (m *ApiSessionUpdated) GetApiSessions() []*ApiSession {
	if m != nil {
		return m.ApiSessions
	}
	return nil
}

type ApiSessionRemoved struct {
	Tokens               []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiSessionRemoved) Reset()         { *m = ApiSessionRemoved{} }
func (m *ApiSessionRemoved) String() string { return proto.CompactTextString(m) }
func (*ApiSessionRemoved) ProtoMessage()    {}
func (*ApiSessionRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{9}
}

func (m *ApiSessionRemoved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSessionRemoved.Unmarshal(m, b)
}
func (m *ApiSessionRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSessionRemoved.Marshal(b, m, deterministic)
}
func (m *ApiSessionRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSessionRemoved.Merge(m, src)
}
func (m *ApiSessionRemoved) XXX_Size() int {
	return xxx_messageInfo_ApiSessionRemoved.Size(m)
}
func (m *ApiSessionRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSessionRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSessionRemoved proto.InternalMessageInfo

func (m *ApiSessionRemoved) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type ApiSessionHeartbeat struct {
	Tokens               []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiSessionHeartbeat) Reset()         { *m = ApiSessionHeartbeat{} }
func (m *ApiSessionHeartbeat) String() string { return proto.CompactTextString(m) }
func (*ApiSessionHeartbeat) ProtoMessage()    {}
func (*ApiSessionHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{10}
}

func (m *ApiSessionHeartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSessionHeartbeat.Unmarshal(m, b)
}
func (m *ApiSessionHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSessionHeartbeat.Marshal(b, m, deterministic)
}
func (m *ApiSessionHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSessionHeartbeat.Merge(m, src)
}
func (m *ApiSessionHeartbeat) XXX_Size() int {
	return xxx_messageInfo_ApiSessionHeartbeat.Size(m)
}
func (m *ApiSessionHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSessionHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSessionHeartbeat proto.InternalMessageInfo

func (m *ApiSessionHeartbeat) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type SessionAdded struct {
	IsFullState          bool       `protobuf:"varint,1,opt,name=isFullState,proto3" json:"isFullState,omitempty"`
	Sessions             []*Session `protobuf:"bytes,2,rep,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SessionAdded) Reset()         { *m = SessionAdded{} }
func (m *SessionAdded) String() string { return proto.CompactTextString(m) }
func (*SessionAdded) ProtoMessage()    {}
func (*SessionAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{11}
}

func (m *SessionAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionAdded.Unmarshal(m, b)
}
func (m *SessionAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionAdded.Marshal(b, m, deterministic)
}
func (m *SessionAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionAdded.Merge(m, src)
}
func (m *SessionAdded) XXX_Size() int {
	return xxx_messageInfo_SessionAdded.Size(m)
}
func (m *SessionAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionAdded.DiscardUnknown(m)
}

var xxx_messageInfo_SessionAdded proto.InternalMessageInfo

func (m *SessionAdded) GetIsFullState() bool {
	if m != nil {
		return m.IsFullState
	}
	return false
}

func (m *SessionAdded) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

type SessionRemoved struct {
	Tokens               []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRemoved) Reset()         { *m = SessionRemoved{} }
func (m *SessionRemoved) String() string { return proto.CompactTextString(m) }
func (*SessionRemoved) ProtoMessage()    {}
func (*SessionRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{12}
}

func (m *SessionRemoved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRemoved.Unmarshal(m, b)
}
func (m *SessionRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRemoved.Marshal(b, m, deterministic)
}
func (m *SessionRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRemoved.Merge(m, src)
}
func (m *SessionRemoved) XXX_Size() int {
	return xxx_messageInfo_SessionRemoved.Size(m)
}
func (m *SessionRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRemoved proto.InternalMessageInfo

func (m *SessionRemoved) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type SessionUpdated struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Urls                 []string `protobuf:"bytes,2,rep,name=urls,proto3" json:"urls,omitempty"`
	Service              *Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionUpdated) Reset()         { *m = SessionUpdated{} }
func (m *SessionUpdated) String() string { return proto.CompactTextString(m) }
func (*SessionUpdated) ProtoMessage()    {}
func (*SessionUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_23f46a161f139ee1, []int{13}
}

func (m *SessionUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionUpdated.Unmarshal(m, b)
}
func (m *SessionUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionUpdated.Marshal(b, m, deterministic)
}
func (m *SessionUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionUpdated.Merge(m, src)
}
func (m *SessionUpdated) XXX_Size() int {
	return xxx_messageInfo_SessionUpdated.Size(m)
}
func (m *SessionUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_SessionUpdated proto.InternalMessageInfo

func (m *SessionUpdated) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SessionUpdated) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *SessionUpdated) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterEnum("edge_ctrl.pb.ContentType", ContentType_name, ContentType_value)
	proto.RegisterType((*ServerHello)(nil), "edge_ctrl.pb.ServerHello")
	proto.RegisterMapType((map[string]string)(nil), "edge_ctrl.pb.ServerHello.DataEntry")
	proto.RegisterType((*ClientHello)(nil), "edge_ctrl.pb.ClientHello")
	proto.RegisterMapType((map[string]string)(nil), "edge_ctrl.pb.ClientHello.DataEntry")
	proto.RegisterType((*Error)(nil), "edge_ctrl.pb.Error")
	proto.RegisterType((*Dns)(nil), "edge_ctrl.pb.Dns")
	proto.RegisterType((*Service)(nil), "edge_ctrl.pb.Service")
	proto.RegisterType((*Session)(nil), "edge_ctrl.pb.Session")
	proto.RegisterType((*ApiSession)(nil), "edge_ctrl.pb.ApiSession")
	proto.RegisterType((*ApiSessionAdded)(nil), "edge_ctrl.pb.ApiSessionAdded")
	proto.RegisterType((*ApiSessionUpdated)(nil), "edge_ctrl.pb.ApiSessionUpdated")
	proto.RegisterType((*ApiSessionRemoved)(nil), "edge_ctrl.pb.ApiSessionRemoved")
	proto.RegisterType((*ApiSessionHeartbeat)(nil), "edge_ctrl.pb.ApiSessionHeartbeat")
	proto.RegisterType((*SessionAdded)(nil), "edge_ctrl.pb.SessionAdded")
	proto.RegisterType((*SessionRemoved)(nil), "edge_ctrl.pb.SessionRemoved")
	proto.RegisterType((*SessionUpdated)(nil), "edge_ctrl.pb.SessionUpdated")
}

func init() { proto.RegisterFile("edge_ctrl.proto", fileDescriptor_23f46a161f139ee1) }

var fileDescriptor_23f46a161f139ee1 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x3d, 0xce, 0xa5, 0x4d, 0x76, 0xaa, 0x26, 0x9d, 0x73, 0xda, 0xe3, 0xd3, 0x03, 0x52, 0xe5,
	0xbe, 0x44, 0x05, 0x82, 0x28, 0x0f, 0x45, 0x7d, 0xab, 0x7a, 0x51, 0x25, 0x24, 0x2e, 0x6e, 0x79,
	0xe1, 0x05, 0xb9, 0xf1, 0x28, 0x58, 0x75, 0x3d, 0xd6, 0xcc, 0xa4, 0x52, 0xdf, 0x81, 0x7f, 0x40,
	0x82, 0x0a, 0xf8, 0x11, 0x7e, 0x00, 0x78, 0xe1, 0x27, 0xb8, 0xfc, 0x04, 0x7b, 0xb6, 0x2f, 0xb1,
	0x93, 0x00, 0xad, 0xc4, 0xdb, 0xec, 0xb5, 0xb6, 0x97, 0xf7, 0x5a, 0xde, 0x1e, 0x68, 0x73, 0x7f,
	0xc0, 0x9f, 0xf4, 0xb5, 0x0c, 0x7b, 0xb1, 0x14, 0x5a, 0xb0, 0xb9, 0x02, 0x70, 0xe4, 0xbc, 0xb4,
	0xa0, 0x75, 0xc0, 0xe5, 0x29, 0x97, 0xfb, 0x3c, 0x0c, 0x05, 0xb3, 0x61, 0x16, 0xcf, 0x2a, 0x10,
	0x91, 0x6d, 0xad, 0x58, 0xdd, 0xa6, 0x9b, 0x95, 0x6c, 0x03, 0x6a, 0xbe, 0xa7, 0x3d, 0xbb, 0xb2,
	0x52, 0xed, 0xb6, 0xd6, 0x57, 0x7b, 0x45, 0x99, 0x5e, 0x41, 0xa2, 0xb7, 0x83, 0x5d, 0xbb, 0x91,
	0x96, 0x67, 0x2e, 0x3d, 0xb0, 0xbc, 0x01, 0xcd, 0x1c, 0x62, 0x1d, 0xa8, 0x1e, 0xf3, 0xb3, 0x54,
	0xdb, 0x1c, 0xd9, 0x3f, 0x50, 0x3f, 0xf5, 0xc2, 0x21, 0x47, 0x61, 0x83, 0x25, 0xc5, 0x66, 0xe5,
	0x8e, 0xe5, 0x7c, 0xc6, 0xd9, 0xb6, 0xc3, 0x80, 0x47, 0xfa, 0x77, 0xb3, 0x2d, 0x43, 0xe3, 0xa9,
	0x50, 0x3a, 0xf2, 0x4e, 0x32, 0x99, 0xbc, 0x66, 0x57, 0xa0, 0x49, 0xc6, 0xfb, 0x22, 0x54, 0x76,
	0x15, 0x87, 0x6f, 0xba, 0x23, 0x20, 0x77, 0x55, 0x9b, 0xe6, 0xaa, 0xf0, 0xf2, 0x3f, 0xe7, 0xea,
	0x2e, 0xd4, 0x77, 0xa5, 0x14, 0x92, 0x31, 0xa8, 0xf5, 0x85, 0xcf, 0xd3, 0xa7, 0xe8, 0x6c, 0x2c,
	0x9e, 0x70, 0xa5, 0xbc, 0x41, 0xf6, 0x60, 0x56, 0x1a, 0xc1, 0xbe, 0x37, 0x54, 0x1c, 0x2d, 0x90,
	0x20, 0x15, 0xce, 0x43, 0xa8, 0xee, 0x44, 0xaa, 0xe4, 0xdf, 0x1a, 0xf3, 0x8f, 0xaf, 0x89, 0x85,
	0xd4, 0xa4, 0x57, 0x77, 0xe9, 0x6c, 0xfa, 0xb3, 0x08, 0x52, 0xbd, 0xbc, 0x76, 0x3e, 0x59, 0x30,
	0x6b, 0x3e, 0x67, 0xd0, 0xe7, 0x6c, 0x1e, 0x2a, 0x81, 0x9f, 0x2a, 0xe2, 0xc9, 0x68, 0x15, 0x32,
	0xa6, 0x33, 0xbb, 0x0e, 0x0b, 0x6a, 0x18, 0x1b, 0x59, 0xee, 0x3f, 0xc8, 0x44, 0x6b, 0xd4, 0x30,
	0x49, 0x30, 0x07, 0xe6, 0xf8, 0x40, 0xa2, 0x27, 0x57, 0x0c, 0x35, 0x97, 0x76, 0x9d, 0x1a, 0x4b,
	0x18, 0xeb, 0xe2, 0xd2, 0x46, 0x7e, 0x2c, 0x82, 0x48, 0x6f, 0xf9, 0xbe, 0x21, 0xec, 0x19, 0x6a,
	0x1b, 0x87, 0xd9, 0x2a, 0x54, 0xfd, 0x48, 0xd9, 0x0d, 0x64, 0x5b, 0xeb, 0x0b, 0xe5, 0x8f, 0x87,
	0xb9, 0xb8, 0x86, 0x4d, 0x0d, 0x29, 0x5a, 0x14, 0x4c, 0x51, 0x8b, 0x63, 0x9e, 0x2d, 0x50, 0x52,
	0x98, 0xd4, 0x4d, 0x5c, 0x41, 0x34, 0x20, 0x67, 0x0d, 0x37, 0x2b, 0xd9, 0x1a, 0x74, 0xfa, 0x5c,
	0xea, 0x3d, 0x3c, 0x73, 0x19, 0x4b, 0x7c, 0x75, 0xb6, 0x43, 0x13, 0xb8, 0x09, 0x67, 0x28, 0x71,
	0xc7, 0x6a, 0xc4, 0xd3, 0x99, 0xdd, 0x84, 0x59, 0x95, 0x64, 0x49, 0x4e, 0x5b, 0xeb, 0x8b, 0x93,
	0xff, 0x0d, 0x92, 0x6e, 0xd6, 0x65, 0xf2, 0x51, 0xc9, 0xac, 0x87, 0x34, 0x67, 0x62, 0xbc, 0x84,
	0x39, 0xf7, 0x00, 0xb6, 0xe2, 0xe0, 0xd7, 0x96, 0xa6, 0x0d, 0x5e, 0x99, 0x3e, 0xb8, 0x23, 0xa0,
	0x3d, 0xd2, 0xc3, 0x68, 0xb9, 0xcf, 0x56, 0xa0, 0x15, 0xa8, 0xbd, 0x61, 0x18, 0x1e, 0x68, 0x4f,
	0x27, 0x3b, 0xd5, 0x70, 0x8b, 0x10, 0xdb, 0x84, 0x96, 0x97, 0x3f, 0xa4, 0xd2, 0x5b, 0xc1, 0x2e,
	0xbb, 0x1b, 0xa9, 0xba, 0xc5, 0x66, 0xe7, 0x3e, 0x2c, 0x8c, 0xa8, 0x47, 0x31, 0xfe, 0x4f, 0xf8,
	0xca, 0x31, 0x41, 0xeb, 0x32, 0x82, 0xd7, 0x8a, 0x82, 0x2e, 0x3f, 0x11, 0xa7, 0x28, 0xb8, 0x04,
	0x33, 0x94, 0x45, 0xa2, 0xd5, 0x74, 0xd3, 0xca, 0xb9, 0x01, 0x7f, 0x8f, 0x9a, 0xf7, 0xb9, 0x27,
	0xf5, 0x11, 0xf7, 0xf4, 0x4f, 0xdb, 0xfb, 0x30, 0x77, 0xc9, 0x68, 0x6e, 0x41, 0x43, 0x95, 0x73,
	0x99, 0xf8, 0xea, 0xc9, 0xa0, 0x79, 0x9b, 0xd3, 0x85, 0xf9, 0x0b, 0x4e, 0x7f, 0x9c, 0x77, 0x66,
	0xc1, 0x4d, 0x5f, 0x80, 0x6c, 0x1b, 0x2b, 0xd3, 0xb7, 0xb1, 0x7a, 0x91, 0x6d, 0x5c, 0x7b, 0x5f,
	0xc1, 0x1b, 0x58, 0x44, 0x1a, 0x6f, 0xc1, 0xc3, 0xb3, 0x98, 0xb3, 0x06, 0xd4, 0x1e, 0x73, 0x29,
	0x3a, 0x7f, 0xb1, 0x45, 0x68, 0x17, 0xee, 0x7c, 0x43, 0x76, 0xde, 0xbc, 0xb2, 0x0c, 0x5c, 0xb8,
	0x34, 0x09, 0x7e, 0x8b, 0x70, 0x1b, 0x9a, 0x74, 0xe7, 0x11, 0xf0, 0x0e, 0x81, 0x25, 0xe8, 0x14,
	0x43, 0x25, 0xfc, 0xd9, 0x6b, 0x0b, 0xff, 0x44, 0x56, 0x76, 0x47, 0xcc, 0xf3, 0x12, 0x93, 0x26,
	0x44, 0xcc, 0x0b, 0x64, 0xfe, 0x2b, 0x7e, 0xcf, 0x91, 0xdc, 0x17, 0xa4, 0xfe, 0x87, 0xc5, 0x89,
	0x45, 0x23, 0xf2, 0xeb, 0x38, 0x59, 0x14, 0xfd, 0x86, 0xe4, 0x55, 0xf8, 0x77, 0xca, 0x92, 0x10,
	0xfd, 0x1d, 0xe9, 0x0e, 0xc0, 0x6e, 0x24, 0x45, 0x18, 0x12, 0xf2, 0xe1, 0x9c, 0x9c, 0x27, 0xc8,
	0x36, 0xfe, 0x5e, 0x8a, 0xe0, 0x8f, 0xe7, 0xd6, 0xd1, 0x0c, 0xdd, 0xab, 0xb7, 0x7f, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xa4, 0x2c, 0xb1, 0x89, 0x87, 0x07, 0x00, 0x00,
}
