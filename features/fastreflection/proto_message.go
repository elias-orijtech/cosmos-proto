package fastreflection

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-proto/features/fastreflection/copied"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	protoreflectPkg = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	protoifacePkg   = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoiface")
	protoimplPkg    = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoimpl")

	fmtPkg = protogen.GoImportPath("fmt")
)

func GenProtoMessage(f *protogen.File, g *protogen.GeneratedFile, message *protogen.Message) {
	gen := newGenerator(f, g, message)
	gen.generateExtraTypes()
	gen.generateReflectionType()
	gen.genDescriptor()
	gen.genType()
	gen.genNew()
	gen.genInterface()
	gen.genRange()
	gen.genHas()
	gen.genClear()
	gen.genGet()
	gen.genSet()
	gen.gentMutable()
	gen.genNewField()
	gen.genWhichOneof()
	gen.genGetUnknown()
	gen.genSetUnknown()
	gen.genIsValid()
	gen.genProtoMethods()
}

func newGenerator(f *protogen.File, g *protogen.GeneratedFile, message *protogen.Message) *generator {
	return &generator{
		GeneratedFile: g,
		file:          f,
		message:       message,
		typeName:      fmt.Sprintf("fastReflection_%s", message.GoIdent.GoName),
		err:           nil,
	}
}

type generator struct {
	*protogen.GeneratedFile
	file    *protogen.File
	message *protogen.Message

	typeName string
	err      error
}

// generateExtraTypes generates the protoreflect.List and protoreflect.Map types required.
func (g *generator) generateExtraTypes() {
	for _, field := range g.message.Fields {
		switch {
		case field.Desc.IsMap():
			g.generateMapType(field)
		case field.Desc.IsList():
			g.generateListType(field)
		}
	}
}

// generateMapType generates the fast reflection protoreflect.Map type
// related to the provided protogen.Field.
func (g *generator) generateMapType(field *protogen.Field) {
	(&mapGen{
		GeneratedFile: g.GeneratedFile,
		field:         field,
	}).generate()
}

// generateListType generates the fast reflection protoreflect.List type
// related to the provided protogen.Field.
func (g *generator) generateListType(field *protogen.Field) {
	(&listGen{
		GeneratedFile: g.GeneratedFile,
		field:         field,
	}).generate()
}

func (g *generator) generateReflectionType() {
	// gen interface assertion
	g.P("var _ ", protoreflectPkg.Ident("Message"), " = (*", g.typeName, ")(nil)")
	g.P()
	// gen type
	g.P("type ", g.typeName, " ", g.message.GoIdent.GoName)
	// gen msg implementation
	g.P("func (x *", g.message.GoIdent.GoName, ") ProtoReflect() ", protoreflectPkg.Ident("Message"), "{")
	g.P("return (*", g.typeName, ")(x)")
	g.P("}")
	g.P()

	// gen slowreflection
	f := copied.NewFileInfo(g.file)
	idx := func() int {
		var id int
		var found bool
		for mInfo, index := range f.AllMessagesByPtr {
			if mInfo.Message.Desc.FullName() == g.message.Desc.FullName() {
				id = index
				found = true
			}
		}
		if !found {
			panic("not found")
		}
		return id
	}()
	typesVar := copied.MessageTypesVarName(f)

	// ProtoReflect method.
	g.P("func (x *", g.message.GoIdent, ") slowProtoReflect() ", protoreflectPkg.Ident("Message"), " {")
	g.P("mi := &", typesVar, "[", idx, "]")
	g.P("if ", protoimplPkg.Ident("UnsafeEnabled"), " && x != nil {")
	g.P("ms := ", protoimplPkg.Ident("X"), ".MessageStateOf(", protoimplPkg.Ident("Pointer"), "(x))")
	g.P("if ms.LoadMessageInfo() == nil {")
	g.P("ms.StoreMessageInfo(mi)")
	g.P("}")
	g.P("return ms")
	g.P("}")
	g.P("return mi.MessageOf(x)")
	g.P("}")
	g.P()
}

func (g *generator) genDescriptor() {
	g.P("// Descriptor returns message descriptor, which contains only the protobuf")
	g.P("// type information for the message.")
	g.P("func (x *", g.typeName, ") Descriptor() ", protoreflectPkg.Ident("MessageDescriptor"), " {")
	slowReflectionFallBack(g.GeneratedFile, g.message, true, "Descriptor")
	g.P("}")
	g.P()
}

func (g *generator) genType() {
	g.P("// Type returns the message type, which encapsulates both Go and protobuf")
	g.P("// type information. If the Go type information is not needed,")
	g.P("// it is recommended that the message descriptor be used instead.")
	g.P("func (x *", g.typeName, ") Type() ", protoreflectPkg.Ident("MessageType"), " {")
	slowReflectionFallBack(g.GeneratedFile, g.message, true, "Type")
	g.P("}")
	g.P()
}

func (g *generator) genNew() {
	g.P("// New returns a newly allocated and mutable empty message.")
	g.P("func (x *", g.typeName, ") New() ", protoreflectPkg.Ident("Message"), " {")
	g.P("return new(", g.typeName, ")")
	g.P("}")
	g.P()
}

func (g *generator) genInterface() {
	g.P("// Interface unwraps the message reflection interface and")
	g.P("// returns the underlying ProtoMessage interface.")
	g.P("func (x *", g.typeName, ") Interface() ", protoreflectPkg.Ident("ProtoMessage"), " {")
	g.P("return (*", g.message.GoIdent, ")(x)")
	g.P("}")
	g.P()
}

func (g *generator) genRange() {
	g.P("// Range iterates over every populated field in an undefined order,")
	g.P("// calling f for each field descriptor and value encountered.")
	g.P("// Range returns immediately if f returns false.")
	g.P("// While iterating, mutating operations may only be performed")
	g.P("// on the current field descriptor.")
	g.P("func (x *", g.typeName, ") Range(f func(", protoreflectPkg.Ident("FieldDescriptor"), ", ", protoreflectPkg.Ident("Value"), ") bool) {")
	slowReflectionFallBack(g.GeneratedFile, g.message, false, "Range", "f")
	g.P("}")
	g.P()
}

func (g *generator) genHas() {
	(&hasGen{
		GeneratedFile: g.GeneratedFile,
		typeName:      g.typeName,
		message:       g.message,
	}).generate()
}

func (g *generator) genClear() {
	(&clearGen{
		GeneratedFile: g.GeneratedFile,
		typeName:      g.typeName,
		message:       g.message,
	}).generate()

	g.P()
}

func (g *generator) genSet() {
	(&setGen{
		GeneratedFile: g.GeneratedFile,
		typeName:      g.typeName,
		message:       g.message,
	}).generate()
}

func (g *generator) gentMutable() {
	g.P("// Mutable returns a mutable reference to a composite type.")
	g.P("//")
	g.P("// If the field is unpopulated, it may allocate a composite value.")
	g.P("// For a field belonging to a oneof, it implicitly clears any other field")
	g.P("// that may be currently set within the same oneof.")
	g.P("// For extension fields, it implicitly stores the provided ExtensionType")
	g.P("// if not already stored.")
	g.P("// It panics if the field does not contain a composite type.")
	g.P("//")
	g.P("// Mutable is a mutating operation and unsafe for concurrent use.")
	g.P("func (x *", g.typeName, ") Mutable(descriptor ", protoreflectPkg.Ident("FieldDescriptor"), ") ", protoreflectPkg.Ident("Value"), " {")
	slowReflectionFallBack(g.GeneratedFile, g.message, true, "Mutable", "descriptor")
	g.P("}")
	g.P()
}

func (g *generator) genNewField() {
	(&newFieldGen{
		GeneratedFile: g.GeneratedFile,
		typeName:      g.typeName,
		message:       g.message,
	}).generate()
	g.P()
}

func (g *generator) genWhichOneof() {
	(&whichOneofGen{
		GeneratedFile: g.GeneratedFile,
		typeName:      g.typeName,
		message:       g.message,
	}).generate()
}

func (g *generator) genGetUnknown() {
	g.P("// GetUnknown retrieves the entire list of unknown fields.")
	g.P("// The caller may only mutate the contents of the RawFields")
	g.P("// if the mutated bytes are stored back into the message with SetUnknown.")
	g.P("func (x *", g.typeName, ") GetUnknown() ", protoreflectPkg.Ident("RawFields"), " {")
	slowReflectionFallBack(g.GeneratedFile, g.message, true, "GetUnknown")
	g.P("}")
	g.P()
}

func (g *generator) genSetUnknown() {
	g.P("// SetUnknown stores an entire list of unknown fields.")
	g.P("// The raw fields must be syntactically valid according to the wire format.")
	g.P("// An implementation may panic if this is not the case.")
	g.P("// Once stored, the caller must not mutate the content of the RawFields.")
	g.P("// An empty RawFields may be passed to clear the fields.")
	g.P("//")
	g.P("// SetUnknown is a mutating operation and unsafe for concurrent use.")
	g.P("func (x *", g.typeName, ") SetUnknown(fields ", protoreflectPkg.Ident("RawFields"), ") {")
	slowReflectionFallBack(g.GeneratedFile, g.message, false, "SetUnknown", "fields")
	g.P("}")
	g.P()
}

func (g *generator) genIsValid() {
	g.P("// IsValid reports whether the message is valid.")
	g.P("//")
	g.P("// An invalid message is an empty, read-only value.")
	g.P("//")
	g.P("// An invalid message often corresponds to a nil pointer of the concrete")
	g.P("// message type, but the details are implementation dependent.")
	g.P("// Validity is not part of the protobuf data model, and may not")
	g.P("// be preserved in marshaling or other operations.")

	g.P("func (x *", g.typeName, ") IsValid() bool {")
	slowReflectionFallBack(g.GeneratedFile, g.message, true, "IsValid")
	g.P("}")
	g.P()
}

func (g *generator) genProtoMethods() {
	g.P("// ProtoMethods returns optional fast-path implementations of various operations.")
	g.P("// This method may return nil.")
	g.P("//")
	g.P("// The returned methods type is identical to")
	g.P(`// "google.golang.org/protobuf/runtime/protoiface".Methods.`)
	g.P("// Consult the protoiface package documentation for details.")
	g.P("func (x *", g.typeName, ") ProtoMethods() *", protoifacePkg.Ident("Methods"), " {")
	slowReflectionFallBack(g.GeneratedFile, g.message, true, "ProtoMethods")
	g.P("}")
}

// slowReflectionFallBack can be used to fallback on slow reflection methods
func slowReflectionFallBack(g *protogen.GeneratedFile, msg *protogen.Message, returns bool, method string, args ...string) {
	switch returns {
	case true:
		g.P("return (*", msg.GoIdent.GoName, ")(x).slowProtoReflect().", method, "(", strings.Join(args, ","), ")")
	case false:
		g.P("(*", msg.GoIdent.GoName, ")(x).slowProtoReflect().", method, "(", strings.Join(args, ","), ")")
	}
}
