// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_experimental.proto

package Ydb_Experimental

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb"
	Ydb_Issue "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Issue"
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Operations"
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

type UploadRowsRequest struct {
	Table                string                          `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Rows                 *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,3,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UploadRowsRequest) Reset()         { *m = UploadRowsRequest{} }
func (m *UploadRowsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRowsRequest) ProtoMessage()    {}
func (*UploadRowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{0}
}

func (m *UploadRowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsRequest.Unmarshal(m, b)
}
func (m *UploadRowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsRequest.Marshal(b, m, deterministic)
}
func (m *UploadRowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsRequest.Merge(m, src)
}
func (m *UploadRowsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRowsRequest.Size(m)
}
func (m *UploadRowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsRequest proto.InternalMessageInfo

func (m *UploadRowsRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UploadRowsRequest) GetRows() *Ydb.TypedValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *UploadRowsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type UploadRowsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadRowsResponse) Reset()         { *m = UploadRowsResponse{} }
func (m *UploadRowsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResponse) ProtoMessage()    {}
func (*UploadRowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{1}
}

func (m *UploadRowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResponse.Unmarshal(m, b)
}
func (m *UploadRowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResponse.Marshal(b, m, deterministic)
}
func (m *UploadRowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResponse.Merge(m, src)
}
func (m *UploadRowsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResponse.Size(m)
}
func (m *UploadRowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResponse proto.InternalMessageInfo

func (m *UploadRowsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type UploadRowsResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRowsResult) Reset()         { *m = UploadRowsResult{} }
func (m *UploadRowsResult) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResult) ProtoMessage()    {}
func (*UploadRowsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{2}
}

func (m *UploadRowsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResult.Unmarshal(m, b)
}
func (m *UploadRowsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResult.Marshal(b, m, deterministic)
}
func (m *UploadRowsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResult.Merge(m, src)
}
func (m *UploadRowsResult) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResult.Size(m)
}
func (m *UploadRowsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResult.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResult proto.InternalMessageInfo

type ExecuteStreamQueryRequest struct {
	YqlText              string                     `protobuf:"bytes,1,opt,name=yql_text,json=yqlText,proto3" json:"yql_text,omitempty"`
	Parameters           map[string]*Ydb.TypedValue `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExecuteStreamQueryRequest) Reset()         { *m = ExecuteStreamQueryRequest{} }
func (m *ExecuteStreamQueryRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryRequest) ProtoMessage()    {}
func (*ExecuteStreamQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{3}
}

func (m *ExecuteStreamQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryRequest.Merge(m, src)
}
func (m *ExecuteStreamQueryRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Size(m)
}
func (m *ExecuteStreamQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryRequest proto.InternalMessageInfo

func (m *ExecuteStreamQueryRequest) GetYqlText() string {
	if m != nil {
		return m.YqlText
	}
	return ""
}

func (m *ExecuteStreamQueryRequest) GetParameters() map[string]*Ydb.TypedValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ExecuteStreamQueryResponse struct {
	Status               Ydb.StatusIds_StatusCode  `protobuf:"varint,1,opt,name=status,proto3,enum=Ydb.StatusIds_StatusCode" json:"status,omitempty"`
	Issues               []*Ydb_Issue.IssueMessage `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	Result               *ExecuteStreamQueryResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExecuteStreamQueryResponse) Reset()         { *m = ExecuteStreamQueryResponse{} }
func (m *ExecuteStreamQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResponse) ProtoMessage()    {}
func (*ExecuteStreamQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{4}
}

func (m *ExecuteStreamQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResponse.Merge(m, src)
}
func (m *ExecuteStreamQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Size(m)
}
func (m *ExecuteStreamQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResponse proto.InternalMessageInfo

func (m *ExecuteStreamQueryResponse) GetStatus() Ydb.StatusIds_StatusCode {
	if m != nil {
		return m.Status
	}
	return Ydb.StatusIds_STATUS_CODE_UNSPECIFIED
}

func (m *ExecuteStreamQueryResponse) GetIssues() []*Ydb_Issue.IssueMessage {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ExecuteStreamQueryResponse) GetResult() *ExecuteStreamQueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type ExecuteStreamQueryResult struct {
	ResultSet            *Ydb.ResultSet `protobuf:"bytes,1,opt,name=result_set,json=resultSet,proto3" json:"result_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecuteStreamQueryResult) Reset()         { *m = ExecuteStreamQueryResult{} }
func (m *ExecuteStreamQueryResult) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResult) ProtoMessage()    {}
func (*ExecuteStreamQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{5}
}

func (m *ExecuteStreamQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResult.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResult.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResult.Merge(m, src)
}
func (m *ExecuteStreamQueryResult) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResult.Size(m)
}
func (m *ExecuteStreamQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResult proto.InternalMessageInfo

func (m *ExecuteStreamQueryResult) GetResultSet() *Ydb.ResultSet {
	if m != nil {
		return m.ResultSet
	}
	return nil
}

func init() {
	proto.RegisterType((*UploadRowsRequest)(nil), "Ydb.Experimental.UploadRowsRequest")
	proto.RegisterType((*UploadRowsResponse)(nil), "Ydb.Experimental.UploadRowsResponse")
	proto.RegisterType((*UploadRowsResult)(nil), "Ydb.Experimental.UploadRowsResult")
	proto.RegisterType((*ExecuteStreamQueryRequest)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest")
	proto.RegisterMapType((map[string]*Ydb.TypedValue)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry")
	proto.RegisterType((*ExecuteStreamQueryResponse)(nil), "Ydb.Experimental.ExecuteStreamQueryResponse")
	proto.RegisterType((*ExecuteStreamQueryResult)(nil), "Ydb.Experimental.ExecuteStreamQueryResult")
}

func init() { proto.RegisterFile("ydb_experimental.proto", fileDescriptor_ac21a693e2c386a5) }

var fileDescriptor_ac21a693e2c386a5 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0x5b, 0xb7, 0xda, 0x5b, 0xd8, 0xd6, 0x41, 0xb4, 0xad, 0x0f, 0x2e, 0x11, 0xa1, 0x88,
	0x26, 0x5a, 0x05, 0x45, 0xdf, 0x2a, 0x7d, 0xa8, 0xb0, 0x5a, 0xa7, 0xab, 0x20, 0x3e, 0x84, 0x49,
	0x73, 0x91, 0xd0, 0x24, 0x93, 0xce, 0x4c, 0xdc, 0xcc, 0xaf, 0xf8, 0x3d, 0xfe, 0x89, 0x3f, 0xe1,
	0xa3, 0x64, 0x32, 0xed, 0x66, 0x17, 0x5b, 0x7c, 0x29, 0xb7, 0xf7, 0x9e, 0x73, 0x72, 0xe6, 0xcc,
	0x1d, 0xb8, 0xab, 0xa3, 0x30, 0xc0, 0x32, 0x47, 0x11, 0xa7, 0x98, 0x29, 0x96, 0x78, 0xb9, 0xe0,
	0x8a, 0x93, 0xfe, 0xd7, 0x28, 0xf4, 0x66, 0x8d, 0xfe, 0xe8, 0xd9, 0x3a, 0x5e, 0xc7, 0xa9, 0xf0,
	0xf3, 0x22, 0x4c, 0xe2, 0x95, 0xcf, 0xf2, 0xd8, 0x37, 0x50, 0xe9, 0x57, 0x12, 0xb1, 0x94, 0x05,
	0x06, 0x29, 0x4a, 0xc9, 0xbe, 0x63, 0xad, 0x31, 0x7a, 0x72, 0x90, 0xc1, 0x73, 0x14, 0x4c, 0xc5,
	0x3c, 0xb3, 0x68, 0xff, 0x20, 0x5a, 0x2a, 0xa6, 0x0a, 0x19, 0xac, 0x78, 0x84, 0xd2, 0x12, 0xc6,
	0x07, 0x09, 0x3f, 0x58, 0x52, 0x58, 0x23, 0xee, 0x4f, 0x07, 0x6e, 0x7f, 0xce, 0x13, 0xce, 0x22,
	0xca, 0x2f, 0x24, 0xc5, 0x4d, 0x81, 0x52, 0x91, 0x3b, 0x70, 0xac, 0x58, 0x98, 0xe0, 0xc0, 0x39,
	0x75, 0xc6, 0x1d, 0x5a, 0xff, 0x21, 0x0f, 0xe1, 0x86, 0xe0, 0x17, 0x72, 0x70, 0x74, 0xea, 0x8c,
	0xbb, 0x93, 0x9e, 0x57, 0xe5, 0x70, 0xae, 0x73, 0x8c, 0xbe, 0x54, 0x82, 0xd4, 0x0c, 0xc9, 0x7b,
	0xe8, 0xef, 0xec, 0x07, 0x39, 0x13, 0x2c, 0x95, 0x83, 0x96, 0x21, 0x3c, 0x30, 0x84, 0x8f, 0xdb,
	0xa1, 0xbc, 0x2c, 0x17, 0x06, 0x46, 0x7b, 0xfc, 0x6a, 0xc3, 0x3d, 0x03, 0xd2, 0xf4, 0x26, 0x73,
	0x9e, 0x49, 0x24, 0xaf, 0xa0, 0xb3, 0x03, 0x1a, 0x83, 0xdd, 0xc9, 0x70, 0xaf, 0x34, 0xbd, 0xc4,
	0xba, 0x04, 0xfa, 0x57, 0xe4, 0x8a, 0x44, 0xb9, 0xbf, 0x1d, 0x18, 0xce, 0x4a, 0x5c, 0x15, 0x0a,
	0x97, 0x4a, 0x20, 0x4b, 0x3f, 0x15, 0x28, 0xf4, 0x36, 0x87, 0x21, 0xdc, 0xd2, 0x9b, 0x24, 0x50,
	0x58, 0x2a, 0x1b, 0xc5, 0x4d, 0xbd, 0x49, 0xce, 0xb1, 0x54, 0xe4, 0x1b, 0x80, 0x39, 0x1d, 0x2a,
	0x14, 0x55, 0x24, 0xad, 0x71, 0x77, 0xf2, 0xd6, 0xbb, 0xbe, 0x1a, 0xde, 0x5e, 0x6d, 0x6f, 0xb1,
	0x63, 0xcf, 0x32, 0x25, 0x34, 0x6d, 0xc8, 0x8d, 0x3e, 0x40, 0xef, 0xda, 0x98, 0xf4, 0xa1, 0xb5,
	0x46, 0x6d, 0x5d, 0x54, 0x25, 0x79, 0x04, 0xc7, 0xe6, 0x26, 0xf7, 0xdd, 0x47, 0x3d, 0x7d, 0x73,
	0xf4, 0xda, 0x71, 0x7f, 0x39, 0x30, 0xfa, 0x97, 0x13, 0x9b, 0xe8, 0x73, 0x68, 0xd7, 0x4b, 0x64,
	0xe4, 0x4f, 0x6c, 0x9c, 0x4b, 0xd3, 0x9a, 0x47, 0xd2, 0x56, 0xef, 0x78, 0x84, 0xd4, 0x02, 0x89,
	0x0f, 0x6d, 0xb3, 0xd7, 0xdb, 0xa3, 0xdf, 0x33, 0x94, 0x79, 0xd5, 0xaa, 0x7f, 0xcf, 0xea, 0x7d,
	0xa7, 0x16, 0x46, 0xa6, 0xd0, 0x16, 0x26, 0x72, 0xbb, 0x0d, 0x8f, 0xff, 0x2f, 0xab, 0x8a, 0x41,
	0x2d, 0xd3, 0x9d, 0xc3, 0x60, 0x1f, 0x86, 0x3c, 0x05, 0xa8, 0x51, 0x81, 0x44, 0x65, 0xd7, 0xe2,
	0xc4, 0x7c, 0xa3, 0x06, 0x2c, 0x51, 0xd1, 0x8e, 0xd8, 0x96, 0xd3, 0x97, 0x70, 0x7f, 0xc5, 0x53,
	0x4f, 0xb3, 0x2c, 0xc2, 0xd2, 0xd3, 0x51, 0xe8, 0x35, 0x5f, 0xfa, 0x94, 0x34, 0x8d, 0x2d, 0xcc,
	0xd3, 0xf9, 0xe3, 0x38, 0x61, 0xdb, 0x3c, 0x9a, 0x17, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0x81, 0x69, 0xb0, 0x1b, 0x04, 0x00, 0x00,
}

const ()

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *UploadRowsRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}