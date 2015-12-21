// Code generated by protoc-gen-go.
// source: openchain/consensus/obcpbft/messages.proto
// DO NOT EDIT!

/*
Package obcpbft is a generated protocol buffer package.

It is generated from these files:
	openchain/consensus/obcpbft/messages.proto

It has these top-level messages:
	Message
	Request
	PrePrepare
	Prepare
	Commit
	Checkpoint
	ViewChange
	NewView
	SieveMessage
	Execute
	Verify
	Decision
*/
package obcpbft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	// Types that are valid to be assigned to Payload:
	//	*Message_Request
	//	*Message_PrePrepare
	//	*Message_Prepare
	//	*Message_Commit
	//	*Message_Checkpoint
	//	*Message_ViewChange
	//	*Message_NewView
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,oneof"`
}
type Message_PrePrepare struct {
	PrePrepare *PrePrepare `protobuf:"bytes,2,opt,name=pre_prepare,oneof"`
}
type Message_Prepare struct {
	Prepare *Prepare `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type Message_Commit struct {
	Commit *Commit `protobuf:"bytes,4,opt,name=commit,oneof"`
}
type Message_Checkpoint struct {
	Checkpoint *Checkpoint `protobuf:"bytes,5,opt,name=checkpoint,oneof"`
}
type Message_ViewChange struct {
	ViewChange *ViewChange `protobuf:"bytes,6,opt,name=view_change,oneof"`
}
type Message_NewView struct {
	NewView *NewView `protobuf:"bytes,7,opt,name=new_view,oneof"`
}

func (*Message_Request) isMessage_Payload()    {}
func (*Message_PrePrepare) isMessage_Payload() {}
func (*Message_Prepare) isMessage_Payload()    {}
func (*Message_Commit) isMessage_Payload()     {}
func (*Message_Checkpoint) isMessage_Payload() {}
func (*Message_ViewChange) isMessage_Payload() {}
func (*Message_NewView) isMessage_Payload()    {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequest() *Request {
	if x, ok := m.GetPayload().(*Message_Request); ok {
		return x.Request
	}
	return nil
}

func (m *Message) GetPrePrepare() *PrePrepare {
	if x, ok := m.GetPayload().(*Message_PrePrepare); ok {
		return x.PrePrepare
	}
	return nil
}

func (m *Message) GetPrepare() *Prepare {
	if x, ok := m.GetPayload().(*Message_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Message) GetCommit() *Commit {
	if x, ok := m.GetPayload().(*Message_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Message) GetCheckpoint() *Checkpoint {
	if x, ok := m.GetPayload().(*Message_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (m *Message) GetViewChange() *ViewChange {
	if x, ok := m.GetPayload().(*Message_ViewChange); ok {
		return x.ViewChange
	}
	return nil
}

func (m *Message) GetNewView() *NewView {
	if x, ok := m.GetPayload().(*Message_NewView); ok {
		return x.NewView
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, []interface{}{
		(*Message_Request)(nil),
		(*Message_PrePrepare)(nil),
		(*Message_Prepare)(nil),
		(*Message_Commit)(nil),
		(*Message_Checkpoint)(nil),
		(*Message_ViewChange)(nil),
		(*Message_NewView)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *Message_PrePrepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrePrepare); err != nil {
			return err
		}
	case *Message_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Message_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Message_Checkpoint:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Checkpoint); err != nil {
			return err
		}
	case *Message_ViewChange:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ViewChange); err != nil {
			return err
		}
	case *Message_NewView:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewView); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // payload.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Request{msg}
		return true, err
	case 2: // payload.pre_prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrePrepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_PrePrepare{msg}
		return true, err
	case 3: // payload.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Prepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Prepare{msg}
		return true, err
	case 4: // payload.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Commit)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Commit{msg}
		return true, err
	case 5: // payload.checkpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Checkpoint)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Checkpoint{msg}
		return true, err
	case 6: // payload.view_change
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ViewChange)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_ViewChange{msg}
		return true, err
	case 7: // payload.new_view
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NewView)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_NewView{msg}
		return true, err
	default:
		return false, nil
	}
}

type Request struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	ReplicaId uint64                     `protobuf:"varint,3,opt,name=replica_id" json:"replica_id,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PrePrepare struct {
	View           uint64   `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64   `protobuf:"varint,2,opt,name=sequence_number" json:"sequence_number,omitempty"`
	RequestDigest  string   `protobuf:"bytes,3,opt,name=request_digest" json:"request_digest,omitempty"`
	Request        *Request `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	ReplicaId      uint64   `protobuf:"varint,5,opt,name=replica_id" json:"replica_id,omitempty"`
}

func (m *PrePrepare) Reset()         { *m = PrePrepare{} }
func (m *PrePrepare) String() string { return proto.CompactTextString(m) }
func (*PrePrepare) ProtoMessage()    {}

func (m *PrePrepare) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type Prepare struct {
	View           uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number" json:"sequence_number,omitempty"`
	RequestDigest  string `protobuf:"bytes,3,opt,name=request_digest" json:"request_digest,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,4,opt,name=replica_id" json:"replica_id,omitempty"`
}

func (m *Prepare) Reset()         { *m = Prepare{} }
func (m *Prepare) String() string { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()    {}

type Commit struct {
	View           uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number" json:"sequence_number,omitempty"`
	RequestDigest  string `protobuf:"bytes,3,opt,name=request_digest" json:"request_digest,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,4,opt,name=replica_id" json:"replica_id,omitempty"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}

type Checkpoint struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number" json:"sequence_number,omitempty"`
	BlockHash      []byte `protobuf:"bytes,2,opt,name=block_hash,proto3" json:"block_hash,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,3,opt,name=replica_id" json:"replica_id,omitempty"`
	BlockNumber    uint64 `protobuf:"varint,4,opt,name=block_number" json:"block_number,omitempty"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}

type ViewChange struct {
	View      uint64           `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	H         uint64           `protobuf:"varint,2,opt,name=h" json:"h,omitempty"`
	Cset      []*ViewChange_C  `protobuf:"bytes,3,rep,name=cset" json:"cset,omitempty"`
	Pset      []*ViewChange_PQ `protobuf:"bytes,4,rep,name=pset" json:"pset,omitempty"`
	Qset      []*ViewChange_PQ `protobuf:"bytes,5,rep,name=qset" json:"qset,omitempty"`
	ReplicaId uint64           `protobuf:"varint,6,opt,name=replica_id" json:"replica_id,omitempty"`
	Signature []byte           `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ViewChange) Reset()         { *m = ViewChange{} }
func (m *ViewChange) String() string { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()    {}

func (m *ViewChange) GetCset() []*ViewChange_C {
	if m != nil {
		return m.Cset
	}
	return nil
}

func (m *ViewChange) GetPset() []*ViewChange_PQ {
	if m != nil {
		return m.Pset
	}
	return nil
}

func (m *ViewChange) GetQset() []*ViewChange_PQ {
	if m != nil {
		return m.Qset
	}
	return nil
}

type ViewChange_C struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number" json:"sequence_number,omitempty"`
	Digest         []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *ViewChange_C) Reset()         { *m = ViewChange_C{} }
func (m *ViewChange_C) String() string { return proto.CompactTextString(m) }
func (*ViewChange_C) ProtoMessage()    {}

type ViewChange_PQ struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number" json:"sequence_number,omitempty"`
	Digest         string `protobuf:"bytes,2,opt,name=digest" json:"digest,omitempty"`
	View           uint64 `protobuf:"varint,3,opt,name=view" json:"view,omitempty"`
}

func (m *ViewChange_PQ) Reset()         { *m = ViewChange_PQ{} }
func (m *ViewChange_PQ) String() string { return proto.CompactTextString(m) }
func (*ViewChange_PQ) ProtoMessage()    {}

type NewView struct {
	View      uint64            `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Vset      []*ViewChange     `protobuf:"bytes,2,rep,name=vset" json:"vset,omitempty"`
	Xset      map[uint64]string `protobuf:"bytes,3,rep,name=xset" json:"xset,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReplicaId uint64            `protobuf:"varint,4,opt,name=replica_id" json:"replica_id,omitempty"`
}

func (m *NewView) Reset()         { *m = NewView{} }
func (m *NewView) String() string { return proto.CompactTextString(m) }
func (*NewView) ProtoMessage()    {}

func (m *NewView) GetVset() []*ViewChange {
	if m != nil {
		return m.Vset
	}
	return nil
}

func (m *NewView) GetXset() map[uint64]string {
	if m != nil {
		return m.Xset
	}
	return nil
}

type SieveMessage struct {
	// Types that are valid to be assigned to Payload:
	//	*SieveMessage_Request
	//	*SieveMessage_Execute
	//	*SieveMessage_Verify
	//	*SieveMessage_Decision
	//	*SieveMessage_PbftMessage
	Payload isSieveMessage_Payload `protobuf_oneof:"payload"`
}

func (m *SieveMessage) Reset()         { *m = SieveMessage{} }
func (m *SieveMessage) String() string { return proto.CompactTextString(m) }
func (*SieveMessage) ProtoMessage()    {}

type isSieveMessage_Payload interface {
	isSieveMessage_Payload()
}

type SieveMessage_Request struct {
	Request []byte `protobuf:"bytes,1,opt,name=request,proto3,oneof"`
}
type SieveMessage_Execute struct {
	Execute *Execute `protobuf:"bytes,2,opt,name=execute,oneof"`
}
type SieveMessage_Verify struct {
	Verify *Verify `protobuf:"bytes,3,opt,name=verify,oneof"`
}
type SieveMessage_Decision struct {
	Decision *Decision `protobuf:"bytes,4,opt,name=decision,oneof"`
}
type SieveMessage_PbftMessage struct {
	PbftMessage []byte `protobuf:"bytes,5,opt,name=pbft_message,proto3,oneof"`
}

func (*SieveMessage_Request) isSieveMessage_Payload()     {}
func (*SieveMessage_Execute) isSieveMessage_Payload()     {}
func (*SieveMessage_Verify) isSieveMessage_Payload()      {}
func (*SieveMessage_Decision) isSieveMessage_Payload()    {}
func (*SieveMessage_PbftMessage) isSieveMessage_Payload() {}

func (m *SieveMessage) GetPayload() isSieveMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SieveMessage) GetRequest() []byte {
	if x, ok := m.GetPayload().(*SieveMessage_Request); ok {
		return x.Request
	}
	return nil
}

func (m *SieveMessage) GetExecute() *Execute {
	if x, ok := m.GetPayload().(*SieveMessage_Execute); ok {
		return x.Execute
	}
	return nil
}

func (m *SieveMessage) GetVerify() *Verify {
	if x, ok := m.GetPayload().(*SieveMessage_Verify); ok {
		return x.Verify
	}
	return nil
}

func (m *SieveMessage) GetDecision() *Decision {
	if x, ok := m.GetPayload().(*SieveMessage_Decision); ok {
		return x.Decision
	}
	return nil
}

func (m *SieveMessage) GetPbftMessage() []byte {
	if x, ok := m.GetPayload().(*SieveMessage_PbftMessage); ok {
		return x.PbftMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SieveMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _SieveMessage_OneofMarshaler, _SieveMessage_OneofUnmarshaler, []interface{}{
		(*SieveMessage_Request)(nil),
		(*SieveMessage_Execute)(nil),
		(*SieveMessage_Verify)(nil),
		(*SieveMessage_Decision)(nil),
		(*SieveMessage_PbftMessage)(nil),
	}
}

func _SieveMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SieveMessage)
	// payload
	switch x := m.Payload.(type) {
	case *SieveMessage_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Request)
	case *SieveMessage_Execute:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Execute); err != nil {
			return err
		}
	case *SieveMessage_Verify:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Verify); err != nil {
			return err
		}
	case *SieveMessage_Decision:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Decision); err != nil {
			return err
		}
	case *SieveMessage_PbftMessage:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.PbftMessage)
	case nil:
	default:
		return fmt.Errorf("SieveMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _SieveMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SieveMessage)
	switch tag {
	case 1: // payload.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &SieveMessage_Request{x}
		return true, err
	case 2: // payload.execute
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Execute)
		err := b.DecodeMessage(msg)
		m.Payload = &SieveMessage_Execute{msg}
		return true, err
	case 3: // payload.verify
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Verify)
		err := b.DecodeMessage(msg)
		m.Payload = &SieveMessage_Verify{msg}
		return true, err
	case 4: // payload.decision
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Decision)
		err := b.DecodeMessage(msg)
		m.Payload = &SieveMessage_Decision{msg}
		return true, err
	case 5: // payload.pbft_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &SieveMessage_PbftMessage{x}
		return true, err
	default:
		return false, nil
	}
}

type Execute struct {
	View        uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	BlockNumber uint64 `protobuf:"varint,2,opt,name=block_number" json:"block_number,omitempty"`
	Request     []byte `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	ReplicaId   uint64 `protobuf:"varint,4,opt,name=replica_id" json:"replica_id,omitempty"`
}

func (m *Execute) Reset()         { *m = Execute{} }
func (m *Execute) String() string { return proto.CompactTextString(m) }
func (*Execute) ProtoMessage()    {}

type Verify struct {
	BlockNumber   uint64 `protobuf:"varint,2,opt,name=block_number" json:"block_number,omitempty"`
	RequestDigest string `protobuf:"bytes,3,opt,name=request_digest" json:"request_digest,omitempty"`
	ResultDigest  []byte `protobuf:"bytes,4,opt,name=result_digest,proto3" json:"result_digest,omitempty"`
	ReplicaId     uint64 `protobuf:"varint,5,opt,name=replica_id" json:"replica_id,omitempty"`
	Signature     []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Verify) Reset()         { *m = Verify{} }
func (m *Verify) String() string { return proto.CompactTextString(m) }
func (*Verify) ProtoMessage()    {}

type Decision struct {
	BlockNumber   uint64    `protobuf:"varint,2,opt,name=block_number" json:"block_number,omitempty"`
	RequestDigest string    `protobuf:"bytes,3,opt,name=request_digest" json:"request_digest,omitempty"`
	Dset          []*Verify `protobuf:"bytes,4,rep,name=dset" json:"dset,omitempty"`
}

func (m *Decision) Reset()         { *m = Decision{} }
func (m *Decision) String() string { return proto.CompactTextString(m) }
func (*Decision) ProtoMessage()    {}

func (m *Decision) GetDset() []*Verify {
	if m != nil {
		return m.Dset
	}
	return nil
}
