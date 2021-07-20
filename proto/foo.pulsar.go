package examples

import (
	errors "errors"
	fmt "fmt"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	bits "math/bits"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var (
	// Interface guards to verify each message implements proto message interface
	_ protoreflect.Message = &Bar{}
	_ protoreflect.Message = &Hello{}
	_ protoreflect.Message = &B{}
)

<<<<<<< HEAD
=======
type HelloN int32

const (
	Hello_foonum HelloN = 0
	Hello_barnum HelloN = 1
)

// Enum value maps for HelloN.
var (
	HelloN_name = map[int32]string{
		0: "foonum",
		1: "barnum",
	}
	HelloN_value = map[string]int32{
		"foonum": 0,
		"barnum": 1,
	}
)

func (x HelloN) Enum() *HelloN {
	p := new(HelloN)
	*p = x
	return p
}

func (x HelloN) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloN) Descriptor() protoreflect.EnumDescriptor {
	return file_foo_proto_enumTypes[0].Descriptor()
}

func (HelloN) Type() protoreflect.EnumType {
	return &file_foo_proto_enumTypes[0]
}

func (x HelloN) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HelloN.Descriptor instead.
func (HelloN) EnumDescriptor() ([]byte, []int) {
	return file_foo_proto_rawDescGZIP(), []int{1, 0}
}

>>>>>>> ty/pulsar-msg_interface
type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qux []string `protobuf:"bytes,1,rep,name=qux,proto3" json:"qux,omitempty"`
	Baz string   `protobuf:"bytes,2,opt,name=baz,proto3" json:"baz,omitempty"`
	// Types that are assignable to ONEOF:
	//	*Bar_ONEOF_B
	//	*Bar_ONEOF_STRING
	ONEOF isBar_ONEOF `protobuf_oneof:"ONEOF"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Bar) GetQux() []string {
	if x != nil {
		return x.Qux
	}
	return nil
}

func (x *Bar) GetBaz() string {
	if x != nil {
		return x.Baz
	}
	return ""
}

func (m *Bar) GetONEOF() isBar_ONEOF {
	if m != nil {
		return m.ONEOF
	}
	return nil
}

func (x *Bar) GetONEOF_B() *B {
	if x, ok := x.GetONEOF().(*Bar_ONEOF_B); ok {
		return x.ONEOF_B
	}
	return nil
}

func (x *Bar) GetONEOF_STRING() string {
	if x, ok := x.GetONEOF().(*Bar_ONEOF_STRING); ok {
		return x.ONEOF_STRING
	}
	return ""
}

type isBar_ONEOF interface {
	isBar_ONEOF()
}

type Bar_ONEOF_B struct {
	ONEOF_B *B `protobuf:"bytes,20,opt,name=ONEOF_B,json=ONEOFB,proto3,oneof"`
}

type Bar_ONEOF_STRING struct {
	ONEOF_STRING string `protobuf:"bytes,21,opt,name=ONEOF_STRING,json=ONEOFSTRING,proto3,oneof"`
}

func (*Bar_ONEOF_B) isBar_ONEOF() {}

func (*Bar_ONEOF_STRING) isBar_ONEOF() {}

// returns the fast methods for the message
func (x Bar) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(*Bar)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("size error: Bar does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*Bar)
			if !ok {
<<<<<<< HEAD
				return protoiface.UnmarshalOutput{}, errors.New("marshal error: Bar does not implement the protoreflect.Message interface")
=======
				return protoiface.UnmarshalOutput{}, errors.New("Bar does not implement the protoreflect.Message interface")
>>>>>>> ty/pulsar-msg_interface
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x Bar) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x Bar) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x Bar) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x Bar) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x Bar) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x Bar) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x Bar) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
}

type _Bar_1_list struct {
	list []string
}

var _ protoreflect.List = (*_Bar_1_list)(nil)

func (x *_Bar_1_list) Len() int {
	return len(x.list)
}

func (x *_Bar_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString(x.list[i])
}

func (x *_Bar_1_list) Set(i int, value protoreflect.Value) {
	unwrapped := value.String()
	concreteValue := unwrapped
	x.list[i] = concreteValue
}

func (x *_Bar_1_list) Append(value protoreflect.Value) {
	unwrapped := value.String()
	concreteValue := unwrapped
	x.list = append(x.list, concreteValue)
}

func (x *_Bar_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Bar at list field Qux as it is not of Message kind"))
}

func (x *_Bar_1_list) Truncate(n int) {
	x.list = x.list[:n]
}

func (x *_Bar_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Bar_1_list) IsValid() bool {
	return x.list == nil || len(x.list) == 0
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *Bar) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
<<<<<<< HEAD
	switch descriptor.Name() {
	case "baz":
		return protoreflect.ValueOfString(x.Baz)
=======
	switch descriptor.Number() {
	case 0:
	case 1:
		return protoreflect.ValueOfString(x.Baz)
	case 2:
		return protoreflect.ValueOfMessage(x.ONEOF_B)
	case 3:
		return protoreflect.ValueOfString(x.ONEOF_STRING)
>>>>>>> ty/pulsar-msg_interface
	default:
		panic(fmt.Errorf("message cosmos.proto.Bar does not contain field %s", descriptor.Name()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x Bar) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x Bar) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x Bar) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x Bar) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x Bar) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x Bar) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x Bar) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x Bar) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	World    string `protobuf:"bytes,2,opt,name=world,proto3" json:"world,omitempty"`
	Universe bool   `protobuf:"varint,3,opt,name=universe,proto3" json:"universe,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Hello) GetWorld() string {
	if x != nil {
		return x.World
	}
	return ""
}

func (x *Hello) GetUniverse() bool {
	if x != nil {
		return x.Universe
	}
	return false
}

// returns the fast methods for the message
func (x Hello) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(*Hello)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("size error: Hello does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*Hello)
			if !ok {
<<<<<<< HEAD
				return protoiface.UnmarshalOutput{}, errors.New("marshal error: Hello does not implement the protoreflect.Message interface")
=======
				return protoiface.UnmarshalOutput{}, errors.New("Hello does not implement the protoreflect.Message interface")
>>>>>>> ty/pulsar-msg_interface
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x Hello) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x Hello) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x Hello) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x Hello) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x Hello) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x Hello) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x Hello) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *Hello) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
<<<<<<< HEAD
	switch descriptor.Name() {
	case "world":
		return protoreflect.ValueOfString(x.World)
	case "universe":
=======
	switch descriptor.Number() {
	case 0:
		return protoreflect.ValueOfString(x.World)
	case 1:
>>>>>>> ty/pulsar-msg_interface
		return protoreflect.ValueOfBool(x.Universe)
	default:
		panic(fmt.Errorf("message cosmos.proto.Hello does not contain field %s", descriptor.Name()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x Hello) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x Hello) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x Hello) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x Hello) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x Hello) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x Hello) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x Hello) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x Hello) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

type B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *B) Reset() {
	*x = B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*B) ProtoMessage() {}

func (x *B) ProtoReflect() protoreflect.Message {
	mi := &file_foo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *B) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

// returns the fast methods for the message
func (x B) GetMethods() *protoiface.Methods {
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             0,
		Size: func(input protoiface.SizeInput) protoiface.SizeOutput {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: struct{}{},
				Size:              x.Size(),
			}
		},
		Marshal: func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
			v, ok := input.Message.(B)
			if !ok {
				return protoiface.MarshalOutput{}, errors.New("B does not implement the protoreflect.Message interface")
			}

			bz, err := v.Marshal()
			if err != nil {
				return protoiface.MarshalOutput{}, err
			}
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Buf:               bz,
			}, nil
		},
		Unmarshal: func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
			v, ok := input.Message.(*B)
			if !ok {
				return protoiface.UnmarshalOutput{}, errors.New("B does not implement the protoreflect.Message interface")
			}

			if len(input.Buf) < 1 {
				return protoiface.UnmarshalOutput{}, errors.New("unmarshal input did not contain any bytes to unmarshal")
			}
			err := v.Unmarshal(input.Buf)
			if err != nil {
				return protoiface.UnmarshalOutput{}, err
			}
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: struct{}{},
				Flags:             0,
			}, nil
		},
		Merge:            nil,
		CheckInitialized: nil,
	}
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x B) Descriptor() protoreflect.MessageDescriptor {
	return x.ProtoReflect().Descriptor()
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x B) Type() protoreflect.MessageType {
	return x.ProtoReflect().Type()
}

// New returns a newly allocated and mutable empty message.
func (x B) New() protoreflect.Message {
	return x.ProtoReflect().New()
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x B) Interface() protoreflect.ProtoMessage {
	return x.ProtoReflect().Interface()
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x B) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	x.ProtoReflect().Range(f)
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x B) Has(descriptor protoreflect.FieldDescriptor) bool {
	return x.ProtoReflect().Has(descriptor)
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x B) Clear(descriptor protoreflect.FieldDescriptor) {
	x.ProtoReflect().Clear(descriptor)
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *B) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.Number() {
	case 0:
		return protoreflect.ValueOfString(x.X)
	default:
		panic(fmt.Errorf("message cosmos.proto.B does not contain field %s", descriptor.Name()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x B) Set(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) {
	x.ProtoReflect().Set(descriptor, value)
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x B) Mutable(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().Mutable(descriptor)
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x B) NewField(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	return x.ProtoReflect().NewField(descriptor)
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x B) WhichOneof(descriptor protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	return x.ProtoReflect().WhichOneof(descriptor)
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x B) GetUnknown() protoreflect.RawFields {
	return x.ProtoReflect().GetUnknown()
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x B) SetUnknown(fields protoreflect.RawFields) {
	x.ProtoReflect().SetUnknown(fields)
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x B) IsValid() bool {
	return x.ProtoReflect().IsValid()
}

// ProtoMethods returns optional fast-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x B) ProtoMethods() *protoiface.Methods {
	return x.GetMethods()
}

var File_foo_proto protoreflect.FileDescriptor

var file_foo_proto_rawDesc = []byte{
	0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x42, 0x61,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x75, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x71, 0x75, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x62, 0x61, 0x7a, 0x12, 0x2a, 0x0a, 0x07, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x42,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x48, 0x00, 0x52, 0x06, 0x4f, 0x4e, 0x45, 0x4f, 0x46,
	0x42, 0x12, 0x23, 0x0a, 0x0c, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x4e, 0x45, 0x4f, 0x46,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x42, 0x07, 0x0a, 0x05, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x22,
	0x56, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x22, 0x1b, 0x0a, 0x01, 0x6e, 0x12,
	0x0a, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x6e, 0x75, 0x6d, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x62,
	0x61, 0x72, 0x6e, 0x75, 0x6d, 0x10, 0x01, 0x22, 0x11, 0x0a, 0x01, 0x42, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_proto_rawDescOnce sync.Once
	file_foo_proto_rawDescData = file_foo_proto_rawDesc
)

func file_foo_proto_rawDescGZIP() []byte {
	file_foo_proto_rawDescOnce.Do(func() {
		file_foo_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_proto_rawDescData)
	})
	return file_foo_proto_rawDescData
}

var file_foo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_foo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_foo_proto_goTypes = []interface{}{
	(HelloN)(0),   // 0: cosmos.proto.Hello.n
	(*Bar)(nil),   // 1: cosmos.proto.Bar
	(*Hello)(nil), // 2: cosmos.proto.Hello
	(*B)(nil),     // 3: cosmos.proto.B
}
var file_foo_proto_depIdxs = []int32{
	3, // 0: cosmos.proto.Bar.ONEOF_B:type_name -> cosmos.proto.B
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_foo_proto_init() }
func file_foo_proto_init() {
	if File_foo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_foo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*B); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_foo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Bar_ONEOF_B)(nil),
		(*Bar_ONEOF_STRING)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foo_proto_goTypes,
		DependencyIndexes: file_foo_proto_depIdxs,
		EnumInfos:         file_foo_proto_enumTypes,
		MessageInfos:      file_foo_proto_msgTypes,
	}.Build()
	File_foo_proto = out.File
	file_foo_proto_rawDesc = nil
	file_foo_proto_goTypes = nil
	file_foo_proto_depIdxs = nil
}
func (m *Bar) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bar) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.ONEOF.(interface {
		MarshalTo([]byte) (int, error)
		Size() int
	}); ok {
		{
			size := vtmsg.Size()
			i -= size
			if _, err := vtmsg.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Baz) > 0 {
		i -= len(m.Baz)
		copy(dAtA[i:], m.Baz)
		i = encodeVarint(dAtA, i, uint64(len(m.Baz)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Qux) > 0 {
		for iNdEx := len(m.Qux) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Qux[iNdEx])
			copy(dAtA[i:], m.Qux[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.Qux[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Bar_ONEOF_B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bar_ONEOF_B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ONEOF_B != nil {
		size, err := m.ONEOF_B.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	return len(dAtA) - i, nil
}
func (m *Bar_ONEOF_STRING) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bar_ONEOF_STRING) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.ONEOF_STRING)
	copy(dAtA[i:], m.ONEOF_STRING)
	i = encodeVarint(dAtA, i, uint64(len(m.ONEOF_STRING)))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0xaa
	return len(dAtA) - i, nil
}
func (m *Hello) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hello) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Hello) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Universe {
		i--
		if m.Universe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.World) > 0 {
		i -= len(m.World)
		copy(dAtA[i:], m.World)
		i = encodeVarint(dAtA, i, uint64(len(m.World)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *B) Marshal() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *B) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *B) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.X) > 0 {
		i -= len(m.X)
		copy(dAtA[i:], m.X)
		i = encodeVarint(dAtA, i, uint64(len(m.X)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Bar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Qux) > 0 {
		for _, s := range m.Qux {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	l = len(m.Baz)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if vtmsg, ok := m.ONEOF.(interface{ Size() int }); ok {
		n += vtmsg.Size()
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func (m *Bar_ONEOF_B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ONEOF_B != nil {
		l = m.ONEOF_B.Size()
		n += 2 + l + sov(uint64(l))
	}
	return n
}
func (m *Bar_ONEOF_STRING) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ONEOF_STRING)
	n += 2 + l + sov(uint64(l))
	return n
}
func (m *Hello) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.World)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.Universe {
		n += 2
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func (m *B) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.X)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.unknownFields != nil {
		n += len(m.unknownFields)
	}
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Bar) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Qux", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Qux = append(m.Qux, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Baz", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Baz = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ONEOF_B", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if oneof, ok := m.ONEOF.(*Bar_ONEOF_B); ok {
				if err := oneof.ONEOF_B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
			} else {
				v := &B{}
				if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				}
				m.ONEOF = &Bar_ONEOF_B{v}
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ONEOF_STRING", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ONEOF = &Bar_ONEOF_STRING{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Hello) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Hello: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hello: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field World", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.World = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Universe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Universe = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *B) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: B: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: B: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.X = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skip(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflow
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflow
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroup = fmt.Errorf("proto: unexpected end of group")
)
