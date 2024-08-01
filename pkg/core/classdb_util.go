package core

import (
	"reflect"
	"strings"

	. "github.com/godot-go/godot-go/pkg/builtin"
)

func getMethods(t reflect.Type) []string {
	var methods []string
	if t.Kind() != reflect.Ptr {
		t = reflect.PointerTo(t)
	}

	for i := 0; i < t.NumMethod(); i++ {
		methods = append(methods, t.Method(i).Name)
	}
	return methods
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func autoRegisterFunc2ClassDB[T GDClass](t GDClass) {
	// only register the methods that are defined in the current class
	ptrType := reflect.TypeOf((*T)(nil))
	classType := ptrType.Elem()
	if classType.Kind() == reflect.Ptr {
		classType = classType.Elem()
	}

	var embeddedTypes []reflect.Type
	for i := 0; i < classType.NumField(); i++ {
		field := classType.Field(i)
		if field.Anonymous {
			embeddedTypes = append(embeddedTypes, field.Type)
		}
	}

	// TODO(tanjp) use the correct way to get the mothods by reflection
	curMethods := getMethods(classType)
	var embeddedMethods []string
	for _, t := range embeddedTypes {
		msd := getMethods(t)
		for _, m := range msd {
			embeddedMethods = append(embeddedMethods, m)
		}
	}
	for _, iMethod := range curMethods {
		if !contains(embeddedMethods, iMethod) || iMethod == "V_ready" {
			methodName := iMethod

			if strings.HasPrefix(methodName, "V_") {
				virtualName := convertVirtualMethodName(methodName)
				println("Registering virtual method ", methodName, "=>", virtualName)
				ClassDBBindMethodVirtual(t, methodName, virtualName, nil, nil)
			} else {
				cFuncName := convertMethodName(methodName)
				println("Registering method", methodName, "=>", cFuncName)
				ClassDBBindMethod(t, methodName, cFuncName, nil, nil)
			}
		}
	}

}
func convertVirtualMethodName(methodName string) string {
	return methodName[1:]
}

func convertMethodName(methodName string) string {
	var result []rune
	for i, char := range methodName {
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, char+'a'-'A')
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
