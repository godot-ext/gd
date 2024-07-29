package common

import (
	"log"
	"strings"

	"github.com/godot-go/godot-go/cmd/extensionapiparser"
	"github.com/iancoleman/strcase"
)

func GoHasStrTypeInParams(args []extensionapiparser.Argument) bool {
	for _, a := range args {
		if a.IsStringType() {
			return true
		}
	}
	return false
}
func GoIsStringType(arg extensionapiparser.Argument) bool {
	return arg.IsStringType()
}
func GoArgumentNameExt(t extensionapiparser.Argument) string {
	if t.Type == "String" || t.Type == "StringName" {
		return "str_" + goArgumentName(t.Name)
	} else {
		return goArgumentName(t.Name)
	}
}

func GoArgumentTypeExt(t extensionapiparser.Argument) string {
	if t.Type == "String" || t.Type == "StringName" {
		return "string"
	} else {
		return goArgumentType(t.Type)
	}
}

func goArgumentName(t string) string {
	switch t {
	case "string":
		return "strValue"
	case "internal":
		return "internalMode"
	case "type":
		return "typeName"
	case "range":
		return "valueRange"
	case "default":
		return "defaultName"
	case "interface":
		return "interfaceName"
	case "map":
		return "resourceMap"
	case "var":
		return "varName"
	case "func":
		return "callbackFunc"
	default:
		return t
	}
}

func goArgumentType(t string) string {
	if strings.HasPrefix(t, "enum::") {
		t = t[6:]
	}

	if strings.HasPrefix(t, "const ") {
		t = t[6:]
	}

	if strings.HasPrefix(t, "bitfield") {
		t = t[8:]
	}

	var (
		indirection  int
		isTypedArray bool
	)

	if strings.HasPrefix(t, "typedarray::") {
		t = t[12:]
		isTypedArray = true
	}

	if strings.HasSuffix(t, "**") {
		indirection = 2
		t = strings.TrimSpace(t[:len(t)-2])
	}

	if strings.HasSuffix(t, "*") {
		indirection = 1
		t = strings.TrimSpace(t[:len(t)-1])
	}

	switch t {
	case "void", "":
		if isTypedArray {
			log.Panic("unexpected type array")
		}

		switch indirection {
		case 0:
			return ""
		case 1:
			return "unsafe.Pointer"
		case 2:
			return "*unsafe.Pointer"
		default:
			panic("unexepected pointer indirection")
		}
	case "Vector2i", "Vector3i", "Vector4i", "Rect2i":
	case "float", "real_t":
		t = "float32"
	case "double":
		t = "float64"
	case "int8":
		t = "int8"
	case "int16":
		t = "int16"
	case "int32":
		t = "int32"
	case "int", "int64":
		t = "int64"
	case "uint8", "uint8_t":
		t = "uint8"
	case "uint16", "uint16_t":
		t = "uint16"
	case "uint32", "uint32_t":
		t = "uint32"
	case "uint64", "uint64_t":
		t = "uint64"
	case "bool":
		t = "bool"
	case "String":
		t = "String"
	case "Nil":
		t = "Variant"
	default:
		t = strcase.ToCamel(t)
	}

	// if isTypedArray {
	// 	return "[]" + strings.Repeat("*", indirection) + t
	// } else {
	// 	return strings.Repeat("*", indirection) + t
	// }
	return strings.Repeat("*", indirection) + t
}
