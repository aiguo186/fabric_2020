// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/events.proto

package peer // import "github.com/aiguo186/fabric/protos/peer"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import common "github.com/aiguo186/fabric/protos/common"

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

// FilteredBlock is a minimal set of information about a block
type FilteredBlock struct {
	ChannelId            string                 `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Number               uint64                 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	FilteredTransactions []*FilteredTransaction `protobuf:"bytes,4,rep,name=filtered_transactions,json=filteredTransactions,proto3" json:"filtered_transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FilteredBlock) Reset()         { *m = FilteredBlock{} }
func (m *FilteredBlock) String() string { return proto.CompactTextString(m) }
func (*FilteredBlock) ProtoMessage()    {}
func (*FilteredBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_8af932975aef5a3c, []int{0}
}
func (m *FilteredBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilteredBlock.Unmarshal(m, b)
}
func (m *FilteredBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilteredBlock.Marshal(b, m, deterministic)
}
func (dst *FilteredBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilteredBlock.Merge(dst, src)
}
func (m *FilteredBlock) XXX_Size() int {
	return xxx_messageInfo_FilteredBlock.Size(m)
}
func (m *FilteredBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_FilteredBlock.DiscardUnknown(m)
}

var xxx_messageInfo_FilteredBlock proto.InternalMessageInfo

func (m *FilteredBlock) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *FilteredBlock) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *FilteredBlock) GetFilteredTransactions() []*FilteredTransaction {
	if m != nil {
		return m.FilteredTransactions
	}
	return nil
}

// FilteredTransaction is a minimal set of information about a transaction
// within a block
type FilteredTransaction struct {
	Txid             string            `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Type             common.HeaderType `protobuf:"varint,2,opt,name=type,proto3,enum=common.HeaderType" json:"type,omitempty"`
	TxValidationCode TxValidationCode  `protobuf:"varint,3,opt,name=tx_validation_code,json=txValidationCode,proto3,enum=protos.TxValidationCode" json:"tx_validation_code,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*FilteredTransaction_TransactionActions
	Data                 isFilteredTransaction_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FilteredTransaction) Reset()         { *m = FilteredTransaction{} }
func (m *FilteredTransaction) String() string { return proto.CompactTextString(m) }
func (*FilteredTransaction) ProtoMessage()    {}
func (*FilteredTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_8af932975aef5a3c, []int{1}
}
func (m *FilteredTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilteredTransaction.Unmarshal(m, b)
}
func (m *FilteredTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilteredTransaction.Marshal(b, m, deterministic)
}
func (dst *FilteredTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilteredTransaction.Merge(dst, src)
}
func (m *FilteredTransaction) XXX_Size() int {
	return xxx_messageInfo_FilteredTransaction.Size(m)
}
func (m *FilteredTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_FilteredTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_FilteredTransaction proto.InternalMessageInfo

func (m *FilteredTransaction) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *FilteredTransaction) GetType() common.HeaderType {
	if m != nil {
		return m.Type
	}
	return common.HeaderType_MESSAGE
}

func (m *FilteredTransaction) GetTxValidationCode() TxValidationCode {
	if m != nil {
		return m.TxValidationCode
	}
	return TxValidationCode_VALID
}

type isFilteredTransaction_Data interface {
	isFilteredTransaction_Data()
}

type FilteredTransaction_TransactionActions struct {
	TransactionActions *FilteredTransactionActions `protobuf:"bytes,4,opt,name=transaction_actions,json=transactionActions,proto3,oneof"`
}

func (*FilteredTransaction_TransactionActions) isFilteredTransaction_Data() {}

func (m *FilteredTransaction) GetData() isFilteredTransaction_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FilteredTransaction) GetTransactionActions() *FilteredTransactionActions {
	if x, ok := m.GetData().(*FilteredTransaction_TransactionActions); ok {
		return x.TransactionActions
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FilteredTransaction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FilteredTransaction_OneofMarshaler, _FilteredTransaction_OneofUnmarshaler, _FilteredTransaction_OneofSizer, []interface{}{
		(*FilteredTransaction_TransactionActions)(nil),
	}
}

func _FilteredTransaction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FilteredTransaction)
	// Data
	switch x := m.Data.(type) {
	case *FilteredTransaction_TransactionActions:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransactionActions); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FilteredTransaction.Data has unexpected type %T", x)
	}
	return nil
}

func _FilteredTransaction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FilteredTransaction)
	switch tag {
	case 4: // Data.transaction_actions
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FilteredTransactionActions)
		err := b.DecodeMessage(msg)
		m.Data = &FilteredTransaction_TransactionActions{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FilteredTransaction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FilteredTransaction)
	// Data
	switch x := m.Data.(type) {
	case *FilteredTransaction_TransactionActions:
		s := proto.Size(x.TransactionActions)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// FilteredTransactionActions is a wrapper for array of TransactionAction
// message from regular block
type FilteredTransactionActions struct {
	ChaincodeActions     []*FilteredChaincodeAction `protobuf:"bytes,1,rep,name=chaincode_actions,json=chaincodeActions,proto3" json:"chaincode_actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FilteredTransactionActions) Reset()         { *m = FilteredTransactionActions{} }
func (m *FilteredTransactionActions) String() string { return proto.CompactTextString(m) }
func (*FilteredTransactionActions) ProtoMessage()    {}
func (*FilteredTransactionActions) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_8af932975aef5a3c, []int{2}
}
func (m *FilteredTransactionActions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilteredTransactionActions.Unmarshal(m, b)
}
func (m *FilteredTransactionActions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilteredTransactionActions.Marshal(b, m, deterministic)
}
func (dst *FilteredTransactionActions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilteredTransactionActions.Merge(dst, src)
}
func (m *FilteredTransactionActions) XXX_Size() int {
	return xxx_messageInfo_FilteredTransactionActions.Size(m)
}
func (m *FilteredTransactionActions) XXX_DiscardUnknown() {
	xxx_messageInfo_FilteredTransactionActions.DiscardUnknown(m)
}

var xxx_messageInfo_FilteredTransactionActions proto.InternalMessageInfo

func (m *FilteredTransactionActions) GetChaincodeActions() []*FilteredChaincodeAction {
	if m != nil {
		return m.ChaincodeActions
	}
	return nil
}

// FilteredChaincodeAction is a minimal set of information about an action
// within a transaction
type FilteredChaincodeAction struct {
	ChaincodeEvent       *ChaincodeEvent `protobuf:"bytes,1,opt,name=chaincode_event,json=chaincodeEvent,proto3" json:"chaincode_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FilteredChaincodeAction) Reset()         { *m = FilteredChaincodeAction{} }
func (m *FilteredChaincodeAction) String() string { return proto.CompactTextString(m) }
func (*FilteredChaincodeAction) ProtoMessage()    {}
func (*FilteredChaincodeAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_8af932975aef5a3c, []int{3}
}
func (m *FilteredChaincodeAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilteredChaincodeAction.Unmarshal(m, b)
}
func (m *FilteredChaincodeAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilteredChaincodeAction.Marshal(b, m, deterministic)
}
func (dst *FilteredChaincodeAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilteredChaincodeAction.Merge(dst, src)
}
func (m *FilteredChaincodeAction) XXX_Size() int {
	return xxx_messageInfo_FilteredChaincodeAction.Size(m)
}
func (m *FilteredChaincodeAction) XXX_DiscardUnknown() {
	xxx_messageInfo_FilteredChaincodeAction.DiscardUnknown(m)
}

var xxx_messageInfo_FilteredChaincodeAction proto.InternalMessageInfo

func (m *FilteredChaincodeAction) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

// DeliverResponse
type DeliverResponse struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverResponse_Status
	//	*DeliverResponse_Block
	//	*DeliverResponse_FilteredBlock
	Type                 isDeliverResponse_Type `protobuf_oneof:"Type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DeliverResponse) Reset()         { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()    {}
func (*DeliverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_8af932975aef5a3c, []int{4}
}
func (m *DeliverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverResponse.Unmarshal(m, b)
}
func (m *DeliverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverResponse.Marshal(b, m, deterministic)
}
func (dst *DeliverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverResponse.Merge(dst, src)
}
func (m *DeliverResponse) XXX_Size() int {
	return xxx_messageInfo_DeliverResponse.Size(m)
}
func (m *DeliverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverResponse proto.InternalMessageInfo

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Status struct {
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status,oneof"`
}

type DeliverResponse_Block struct {
	Block *common.Block `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

type DeliverResponse_FilteredBlock struct {
	FilteredBlock *FilteredBlock `protobuf:"bytes,3,opt,name=filtered_block,json=filteredBlock,proto3,oneof"`
}

func (*DeliverResponse_Status) isDeliverResponse_Type() {}

func (*DeliverResponse_Block) isDeliverResponse_Type() {}

func (*DeliverResponse_FilteredBlock) isDeliverResponse_Type() {}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverResponse) GetStatus() common.Status {
	if x, ok := m.GetType().(*DeliverResponse_Status); ok {
		return x.Status
	}
	return common.Status_UNKNOWN
}

func (m *DeliverResponse) GetBlock() *common.Block {
	if x, ok := m.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

func (m *DeliverResponse) GetFilteredBlock() *FilteredBlock {
	if x, ok := m.GetType().(*DeliverResponse_FilteredBlock); ok {
		return x.FilteredBlock
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeliverResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DeliverResponse_OneofMarshaler, _DeliverResponse_OneofUnmarshaler, _DeliverResponse_OneofSizer, []interface{}{
		(*DeliverResponse_Status)(nil),
		(*DeliverResponse_Block)(nil),
		(*DeliverResponse_FilteredBlock)(nil),
	}
}

func _DeliverResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Status:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Status))
	case *DeliverResponse_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *DeliverResponse_FilteredBlock:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FilteredBlock); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeliverResponse.Type has unexpected type %T", x)
	}
	return nil
}

func _DeliverResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeliverResponse)
	switch tag {
	case 1: // Type.status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &DeliverResponse_Status{common.Status(x)}
		return true, err
	case 2: // Type.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.Block)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverResponse_Block{msg}
		return true, err
	case 3: // Type.filtered_block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FilteredBlock)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverResponse_FilteredBlock{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DeliverResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Status:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Status))
	case *DeliverResponse_Block:
		s := proto.Size(x.Block)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DeliverResponse_FilteredBlock:
		s := proto.Size(x.FilteredBlock)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FilteredBlock)(nil), "protos.FilteredBlock")
	proto.RegisterType((*FilteredTransaction)(nil), "protos.FilteredTransaction")
	proto.RegisterType((*FilteredTransactionActions)(nil), "protos.FilteredTransactionActions")
	proto.RegisterType((*FilteredChaincodeAction)(nil), "protos.FilteredChaincodeAction")
	proto.RegisterType((*DeliverResponse)(nil), "protos.DeliverResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeliverClient is the client API for Deliver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeliverClient interface {
	// deliver first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of block replies is received
	Deliver(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverClient, error)
	// deliver first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of **filtered** block replies is received
	DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverFilteredClient, error)
}

type deliverClient struct {
	cc *grpc.ClientConn
}

func NewDeliverClient(cc *grpc.ClientConn) DeliverClient {
	return &deliverClient{cc}
}

func (c *deliverClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Deliver_serviceDesc.Streams[0], "/protos.Deliver/Deliver", opts...)
	if err != nil {
		return nil, err
	}
	x := &deliverDeliverClient{stream}
	return x, nil
}

type Deliver_DeliverClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type deliverDeliverClient struct {
	grpc.ClientStream
}

func (x *deliverDeliverClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deliverDeliverClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deliverClient) DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverFilteredClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Deliver_serviceDesc.Streams[1], "/protos.Deliver/DeliverFiltered", opts...)
	if err != nil {
		return nil, err
	}
	x := &deliverDeliverFilteredClient{stream}
	return x, nil
}

type Deliver_DeliverFilteredClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type deliverDeliverFilteredClient struct {
	grpc.ClientStream
}

func (x *deliverDeliverFilteredClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deliverDeliverFilteredClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DeliverServer is the server API for Deliver service.
type DeliverServer interface {
	// deliver first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of block replies is received
	Deliver(Deliver_DeliverServer) error
	// deliver first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of **filtered** block replies is received
	DeliverFiltered(Deliver_DeliverFilteredServer) error
}

func RegisterDeliverServer(s *grpc.Server, srv DeliverServer) {
	s.RegisterService(&_Deliver_serviceDesc, srv)
}

func _Deliver_Deliver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliverServer).Deliver(&deliverDeliverServer{stream})
}

type Deliver_DeliverServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type deliverDeliverServer struct {
	grpc.ServerStream
}

func (x *deliverDeliverServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deliverDeliverServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Deliver_DeliverFiltered_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliverServer).DeliverFiltered(&deliverDeliverFilteredServer{stream})
}

type Deliver_DeliverFilteredServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type deliverDeliverFilteredServer struct {
	grpc.ServerStream
}

func (x *deliverDeliverFilteredServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deliverDeliverFilteredServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Deliver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Deliver",
	HandlerType: (*DeliverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deliver",
			Handler:       _Deliver_Deliver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeliverFiltered",
			Handler:       _Deliver_DeliverFiltered_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer/events.proto",
}

func init() { proto.RegisterFile("peer/events.proto", fileDescriptor_events_8af932975aef5a3c) }

var fileDescriptor_events_8af932975aef5a3c = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x69, 0x08, 0xea, 0x44, 0x49, 0xdb, 0x2d, 0x4d, 0xa3, 0x20, 0xd4, 0xc8, 0x12, 0xc8,
	0x5c, 0x62, 0x64, 0x6e, 0x1c, 0x40, 0xa4, 0x3f, 0x0a, 0x12, 0x87, 0x6a, 0x09, 0x1c, 0x7a, 0xc0,
	0x5a, 0xdb, 0x13, 0xc7, 0xd4, 0xf1, 0x5a, 0xbb, 0x9b, 0x28, 0x79, 0x04, 0xde, 0x80, 0x67, 0xe0,
	0x09, 0x39, 0x22, 0xaf, 0xbd, 0x49, 0x9a, 0x52, 0x24, 0x4e, 0xf6, 0xce, 0x7c, 0x3f, 0x33, 0xb3,
	0x63, 0xc3, 0x51, 0x8e, 0x28, 0x5c, 0x5c, 0x60, 0xa6, 0xe4, 0x20, 0x17, 0x5c, 0x71, 0xd2, 0xd0,
	0x0f, 0xd9, 0x3b, 0x0e, 0xf9, 0x6c, 0xc6, 0x33, 0xb7, 0x7c, 0x94, 0xc9, 0xde, 0x59, 0xcc, 0x79,
	0x9c, 0xa2, 0xab, 0x4f, 0xc1, 0x7c, 0xe2, 0xaa, 0x64, 0x86, 0x52, 0xb1, 0x59, 0x5e, 0x01, 0x7a,
	0x5a, 0x30, 0x9c, 0xb2, 0x24, 0x0b, 0x79, 0x84, 0xbe, 0x96, 0xae, 0x72, 0x1d, 0x9d, 0x53, 0x82,
	0x65, 0x92, 0x85, 0x2a, 0x31, 0xa2, 0xf6, 0x4f, 0x0b, 0x5a, 0x57, 0x49, 0xaa, 0x50, 0x60, 0x34,
	0x4c, 0x79, 0x78, 0x4b, 0x9e, 0x03, 0x84, 0x53, 0x96, 0x65, 0x98, 0xfa, 0x49, 0xd4, 0xb5, 0xfa,
	0x96, 0xb3, 0x4f, 0xf7, 0xab, 0xc8, 0xc7, 0x88, 0x74, 0xa0, 0x91, 0xcd, 0x67, 0x01, 0x8a, 0xee,
	0xa3, 0xbe, 0xe5, 0xd4, 0x69, 0x75, 0x22, 0xd7, 0x70, 0x32, 0xa9, 0x74, 0xfc, 0x2d, 0x1b, 0xd9,
	0xad, 0xf7, 0xf7, 0x9c, 0xa6, 0xf7, 0xac, 0xf4, 0x93, 0x03, 0x63, 0x36, 0xde, 0x60, 0xe8, 0xd3,
	0xc9, 0xfd, 0xa0, 0xb4, 0x7f, 0x5b, 0x70, 0xfc, 0x17, 0x34, 0x21, 0x50, 0x57, 0xcb, 0x75, 0x69,
	0xfa, 0x9d, 0xbc, 0x84, 0xba, 0x5a, 0xe5, 0xa8, 0x6b, 0x6a, 0x7b, 0x64, 0x50, 0x0d, 0x6e, 0x84,
	0x2c, 0x42, 0x31, 0x5e, 0xe5, 0x48, 0x75, 0x9e, 0x5c, 0x01, 0x51, 0x4b, 0x7f, 0xc1, 0xd2, 0x24,
	0x62, 0x85, 0x98, 0x5f, 0x0c, 0xaa, 0xbb, 0xa7, 0x59, 0x5d, 0x53, 0xe2, 0x78, 0xf9, 0x75, 0x0d,
	0x38, 0xe7, 0x11, 0xd2, 0x43, 0xb5, 0x13, 0x21, 0x5f, 0xe0, 0x78, 0xab, 0x49, 0x7f, 0xd3, 0xab,
	0xe5, 0x34, 0x3d, 0xfb, 0x1f, 0xbd, 0x7e, 0x28, 0x91, 0xa3, 0x1a, 0x25, 0xea, 0x5e, 0x74, 0xd8,
	0x80, 0xfa, 0x05, 0x53, 0xcc, 0xfe, 0x0e, 0xbd, 0x87, 0xb9, 0xe4, 0x13, 0x1c, 0x6d, 0x2e, 0xd9,
	0x58, 0x5b, 0x7a, 0xcc, 0x67, 0xbb, 0xd6, 0xe7, 0x06, 0x58, 0x92, 0xe9, 0x61, 0x78, 0x37, 0x20,
	0xed, 0x1b, 0x38, 0x7d, 0x00, 0x4c, 0xde, 0xc3, 0xc1, 0xce, 0x36, 0xe9, 0xa1, 0x37, 0xbd, 0x8e,
	0xb1, 0x59, 0x33, 0x2e, 0x8b, 0x2c, 0x6d, 0x87, 0x77, 0xce, 0xf6, 0x2f, 0x0b, 0x0e, 0x2e, 0x30,
	0x4d, 0x16, 0x28, 0x28, 0xca, 0x9c, 0x67, 0x12, 0x89, 0x03, 0x0d, 0xa9, 0x98, 0x9a, 0x4b, 0xad,
	0xd5, 0xf6, 0xda, 0xe6, 0xb2, 0x3e, 0xeb, 0xe8, 0xa8, 0x46, 0xab, 0x3c, 0x79, 0x01, 0x8f, 0x83,
	0x62, 0x25, 0xf5, 0xad, 0x36, 0xbd, 0x96, 0x01, 0xea, 0x3d, 0x1d, 0xd5, 0x68, 0x99, 0x25, 0xef,
	0xa0, 0xbd, 0xde, 0xbc, 0x12, 0xbf, 0xa7, 0xf1, 0x27, 0xbb, 0xb3, 0x30, 0xbc, 0xd6, 0x64, 0x3b,
	0x50, 0x0c, 0xbd, 0xd8, 0x10, 0xef, 0x87, 0x05, 0x4f, 0xaa, 0x62, 0xc9, 0xdb, 0xcd, 0xeb, 0xa1,
	0xb1, 0xbd, 0xcc, 0x16, 0x98, 0xf2, 0x1c, 0x7b, 0xa7, 0x46, 0x78, 0xa7, 0x35, 0xbb, 0xe6, 0x58,
	0xaf, 0x2d, 0x32, 0x5c, 0xf7, 0x6c, 0x8c, 0xff, 0x5b, 0x63, 0xf8, 0x0d, 0x6c, 0x2e, 0xe2, 0xc1,
	0x74, 0x95, 0xa3, 0x48, 0x31, 0x8a, 0x51, 0x0c, 0x26, 0x2c, 0x10, 0x49, 0x68, 0x68, 0xc5, 0xe7,
	0x3c, 0x6c, 0xe9, 0x29, 0xcb, 0x6b, 0x16, 0xde, 0xb2, 0x18, 0x6f, 0x5e, 0xc5, 0x89, 0x9a, 0xce,
	0x83, 0xc2, 0xcb, 0xdd, 0x62, 0xba, 0x25, 0xb3, 0xfc, 0x6f, 0x48, 0xb7, 0x60, 0x06, 0xe5, 0x8f,
	0xe6, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x19, 0xd1, 0xa9, 0x84, 0x04, 0x00, 0x00,
}
