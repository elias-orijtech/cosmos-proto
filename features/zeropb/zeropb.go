package zeropb

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-proto/features/protoc"
	"github.com/cosmos/cosmos-proto/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	errorsPackage = protogen.GoImportPath("errors")
	mathPackage   = protogen.GoImportPath("math")
	binaryPackage = protogen.GoImportPath("encoding/binary")

	runtimePackage = protogen.GoImportPath("github.com/cosmos/cosmos-proto/runtime/zeropb")
)

func init() {
	generator.RegisterFeature("zeropb", func(gen *generator.GeneratedFile, _ *protogen.Plugin) generator.FeatureGenerator {
		return zeropbFeature{
			gen: gen,
		}
	})
}

type zeropbFeature struct {
	gen *generator.GeneratedFile
}

func (g zeropbFeature) GenerateFile(file *protogen.File, _ *protogen.Plugin) bool {
	for _, m := range file.Messages {
		g.generateMessage(file, m)
	}
	return true // only do this once
}

func (g zeropbFeature) GenerateHelpers() {}

func (g zeropbFeature) generateMessage(f *protogen.File, m *protogen.Message) {
	g.generateMarshal(m)
	g.generateUnmarshal(m)
}

func (g zeropbFeature) generateMarshal(m *protogen.Message) {
	g.gen.P("func (x *", m.GoIdent, ") MarshalZeroPB(buf []byte) (n int, err error) {")
	g.gen.P("    defer func() {")
	g.gen.P("        if e := recover(); e != nil {")
	g.gen.P("            err = ", errorsPackage.Ident("New"), `("buffer overflow")`)
	g.gen.P("        }")
	g.gen.P("    }()")
	g.gen.P("    b := ", runtimePackage.Ident("NewBuffer"), "(buf)")
	g.gen.P("    x.marshalZeroPB(b, b.Alloc(", structSize(m), "))")
	g.gen.P("    return int(b.Allocated()), nil")
	g.gen.P("}")
	g.gen.P()
	g.gen.P("func (x *", m.GoIdent, ") marshalZeroPB(b *", runtimePackage.Ident("Buffer"), ", buf ", runtimePackage.Ident("Allocation"), ") {")
	offset := 0
	for _, f := range m.Fields {
		g.generateMarshalField(f, offset)
		offset += fieldSize(f)
	}
	g.gen.P("}")
}

const (
	sliceSize         = 2 * 2
	segmentHeaderSize = 1 + 1 + 2
)

func structSize(m *protogen.Message) int {
	n := 0
	for _, f := range m.Fields {
		d := f.Desc
		switch {
		case d.IsList(), d.IsMap():
			n += sliceSize
		default:
			n += fieldSize(f)
		}
	}
	return n
}

func fieldSize(f *protogen.Field) int {
	d := f.Desc
	switch {
	case d.IsList(), d.IsMap():
		return 4
	}
	return fieldElemSize(f)
}

func fieldElemSize(f *protogen.Field) int {
	d := f.Desc
	switch d.Kind() {
	case protoreflect.FloatKind:
		return 4
	case protoreflect.DoubleKind:
		return 8
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.EnumKind:
		return 4
	case protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind:
		return 8
	case protoreflect.BoolKind:
		return 4
	case protoreflect.StringKind, protoreflect.BytesKind:
		return sliceSize
	case protoreflect.MessageKind:
		return structSize(f.Message)
	}
	return 0
}

func (g zeropbFeature) generateMarshalField(f *protogen.Field, offset int) {
	d := f.Desc
	switch {
	case d.IsList():
		g.gen.P("buf_", d.Index(), " := b.AllocRel(len(x.", f.GoName, ")*", fieldElemSize(f), " + ", segmentHeaderSize, ", buf, ", offset, ", uint16(len(x.", f.GoName, ")))")
		g.gen.P("{")
		g.gen.P("    buf := buf_", d.Index())
		// Write a segment header.
		g.gen.P("    buf.Buf[0] = byte(len(x.", f.GoName, "))")
		g.gen.P("    buf.Buf[1] = byte(len(x.", f.GoName, "))")
		g.gen.P(binaryPackage.Ident("LittleEndian"), ".PutUint16(buf.Buf[2:], 0)")
		g.gen.P("    for i, e := range x.", f.GoName, " {")
		g.generateMarshalPrimitive(f, "e", fmt.Sprintf("uint16(i)*%d+4", fieldElemSize(f)))
		g.gen.P("    }")
		g.gen.P("}")
	case d.IsMap():
		sz := fieldSize(f.Message.Fields[0]) + fieldSize(f.Message.Fields[1])
		g.gen.P("buf_", d.Index(), " := b.AllocRel(len(x.", f.GoName, ")*", sz, ", buf, ", offset, ", uint16(len(x."+f.GoName+")))")
		g.gen.P("{")
		g.gen.P("    var n uint16")
		g.gen.P("    buf := buf_", d.Index())
		g.gen.P("    for k, v := range x.", f.GoName, " {")
		g.generateMarshalPrimitive(f.Message.Fields[0], "k", "n")
		g.gen.P("    n += ", fieldSize(f.Message.Fields[0]))
		g.generateMarshalPrimitive(f.Message.Fields[1], "v", "n")
		g.gen.P("    n += ", fieldSize(f.Message.Fields[1]))
		g.gen.P("    }")
		g.gen.P("}")
	case d.ContainingOneof() != nil:
		g.gen.P("// TODO: field ", f.GoName)
		return
	default:
		g.generateMarshalPrimitive(f, "x."+f.GoName, strconv.Itoa(offset))
	}
}

func (g zeropbFeature) generateMarshalPrimitive(f *protogen.Field, name, offset string) {
	d := f.Desc
	switch d.Kind() {
	case protoreflect.FloatKind:
		g.gen.P("binary.LittleEndian.PutUint32(buf.Buf[", offset, ":], ", mathPackage.Ident("Float32bits"), "(", name, "))")
	case protoreflect.DoubleKind:
		g.gen.P("binary.LittleEndian.PutUint64(buf.Buf[", offset, ":], ", mathPackage.Ident("Float64bits"), "(", name, "))")
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind, protoreflect.EnumKind:
		g.gen.P("binary.LittleEndian.PutUint32(buf.Buf[", offset, ":], uint32(", name, "))")
	case protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind:
		g.gen.P("binary.LittleEndian.PutUint64(buf.Buf[", offset, ":], uint64(", name, "))")
	case protoreflect.BoolKind:
		g.gen.P("bool_", d.Index(), " := uint32(0)")
		g.gen.P("if ", name, " {")
		g.gen.P("    bool_", d.Index(), " = 1")
		g.gen.P("}")
		g.gen.P("binary.LittleEndian.PutUint32(buf.Buf[", offset, ":], bool_", d.Index(), ")")
	case protoreflect.StringKind, protoreflect.BytesKind:
		g.gen.P("buf_", d.Index(), " := b.AllocRel(len(", name, "), buf, ", offset, ", uint16(len(", name, ")))")
		g.gen.P("copy(buf_", d.Index(), ".Buf, ", name, ")")
	case protoreflect.MessageKind:
		g.gen.P(name, ".marshalZeroPB(b, buf.Slice(", offset, "))")
	default:
		g.gen.P("// TODO: field ", name)
		g.gen.P("_ = ", name)
	}
}

func (g zeropbFeature) generateUnmarshal(m *protogen.Message) {
	g.gen.P("func (x *", m.GoIdent, ") UnmarshalZeroPB(buf []byte) (err error) {")
	g.gen.P("    defer func() {")
	g.gen.P("        if e := recover(); e != nil {")
	g.gen.P("            err = ", errorsPackage.Ident("New"), `("buffer underflow")`)
	g.gen.P("        }")
	g.gen.P("    }()")
	g.gen.P("    x.unmarshalZeroPB(buf, 0)")
	g.gen.P("    return nil")
	g.gen.P("}")
	g.gen.P()
	g.gen.P("func (x *", m.GoIdent, ") unmarshalZeroPB(buf []byte, n uint16) {")
	offset := 0
	for _, f := range m.Fields {
		g.generateUnmarshalField(f, offset)
		offset += fieldSize(f)
	}
	g.gen.P("}")
}

func (g zeropbFeature) generateUnmarshalField(f *protogen.Field, offset int) {
	d := f.Desc
	switch {
	case d.IsList():
		g.gen.P("n_", d.Index(), ", len_", d.Index(), " := ", runtimePackage.Ident("ReadSlice"), "(buf, n+", offset, ")")
		typ, pointer := protoc.FieldGoType(g.gen, f)
		if pointer {
			typ = "*" + typ
		}
		g.gen.P("x.", f.GoName, " = make(", typ, ", len_", d.Index(), ")")
		g.gen.P("{")
		g.gen.P("    for i := range x.", f.GoName, "{")
		// Skip segment header.
		g.generateUnmarshalPrimitive(f, "x."+f.GoName+"[i]", fmt.Sprintf("n_%d+4+uint16(i)*%d", d.Index(), fieldElemSize(f)))
		g.gen.P("    }")
		g.gen.P("}")
	case d.IsMap():
		g.gen.P("n_", d.Index(), ", len_", d.Index(), " := ", runtimePackage.Ident("ReadSlice"), "(buf, n+", offset, ")")
		typ, _ := protoc.FieldGoType(g.gen, f)
		g.gen.P("x.", f.GoName, " = make(", typ, ", len_", d.Index(), ")")
		keyType, _ := protoc.FieldGoType(g.gen, f.Message.Fields[0])
		valType, _ := protoc.FieldGoType(g.gen, f.Message.Fields[1])
		g.gen.P("{")
		g.gen.P("    n := n_", d.Index())
		g.gen.P("    for i := uint16(0); i < len_", d.Index(), "; i++ {")
		g.gen.P("        var k ", keyType)
		g.gen.P("        var v ", valType)
		g.generateUnmarshalPrimitive(f.Message.Fields[0], "k", "n")
		g.gen.P("        n += ", fieldSize(f.Message.Fields[0]))
		g.generateUnmarshalPrimitive(f.Message.Fields[1], "v", "n")
		g.gen.P("        n += ", fieldSize(f.Message.Fields[1]))
		g.gen.P("        x.", f.GoName, "[k] = v")
		g.gen.P("    }")
		g.gen.P("}")
	case d.ContainingOneof() != nil:
		g.gen.P("// TODO: field ", f.GoName)
	default:
		g.generateUnmarshalPrimitive(f, "x."+f.GoName, fmt.Sprintf("n+%d", offset))
	}
}

func (g zeropbFeature) generateUnmarshalPrimitive(f *protogen.Field, name, offset string) {
	switch d := f.Desc; d.Kind() {
	case protoreflect.FloatKind:
		g.gen.P(name, " = float32(", mathPackage.Ident("Float32frombits"), "(", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[", offset, ":])))")
	case protoreflect.DoubleKind:
		g.gen.P(name, " = float64(", mathPackage.Ident("Float64frombits"), "(", binaryPackage.Ident("LittleEndian"), ".Uint64(buf[", offset, ":])))")
	case protoreflect.Sfixed32Kind, protoreflect.Int32Kind, protoreflect.Sint32Kind:
		g.gen.P(name, " = int32(", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[", offset, ":]))")
	case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
		g.gen.P(name, " = ", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[", offset, ":])")
	case protoreflect.Sfixed64Kind, protoreflect.Int64Kind, protoreflect.Sint64Kind:
		g.gen.P(name, " = int64(", binaryPackage.Ident("LittleEndian"), ".Uint64(buf[", offset, ":]))")
	case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		g.gen.P(name, " = ", binaryPackage.Ident("LittleEndian"), ".Uint64(buf[", offset, ":])")
	case protoreflect.EnumKind:
		typ := g.gen.QualifiedGoIdent(f.Enum.GoIdent)
		g.gen.P(name, " = ", typ, "(", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[", offset, ":]))")
	case protoreflect.BoolKind:
		g.gen.P("bool_", d.Index(), " := ", binaryPackage.Ident("LittleEndian"), ".Uint32(buf[", offset, ":])")
		g.gen.P(name, " = false")
		g.gen.P("if bool_", d.Index(), " != 0 {")
		g.gen.P("    ", name, " = true")
		g.gen.P("}")
	case protoreflect.StringKind:
		g.gen.P("n_", d.Index(), ", len_", d.Index(), " := ", runtimePackage.Ident("ReadSlice"), "(buf, ", offset, ")")
		g.gen.P(name, " = string(buf[n_", d.Index(), ":n_", d.Index(), "+len_", d.Index(), "])")
	case protoreflect.BytesKind:
		g.gen.P("n_", d.Index(), ", len_", d.Index(), " := ", runtimePackage.Ident("ReadSlice"), "(buf, ", offset, ")")
		g.gen.P(name, " = append([]byte{}, buf[n_", d.Index(), ":n_", d.Index(), "+len_", d.Index(), "]...)")
	case protoreflect.MessageKind:
		typ := g.gen.QualifiedGoIdent(f.Message.GoIdent)
		g.gen.P(name, " = new(", typ, ")")
		g.gen.P(name, ".unmarshalZeroPB(buf, ", offset, ")")
	default:
		g.gen.P("// TODO: field ", name)
	}
}
