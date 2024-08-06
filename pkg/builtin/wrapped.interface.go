package builtin

import (
	. "github.com/godot-go/godot-go/pkg/ffi"
)

type GDExtensionClass interface {
	Wrapped
}

type HasDestructor interface {
	Destroy()
}

// Base for all engine classes, to contain the pointer to the engine instance.
type Wrapped interface {
	HasDestructor
	GetGodotObjectOwner() *GodotObject
	SetGodotObjectOwner(owner *GodotObject)
	GetClassName(realInstance interface{}) string
	GetParentClassName(realInstance interface{}) string
	AsGDExtensionObjectPtr() GDExtensionObjectPtr
	AsGDExtensionConstObjectPtr() GDExtensionConstObjectPtr
	AsGDExtensionTypePtr() GDExtensionTypePtr
	AsGDExtensionConstTypePtr() GDExtensionConstTypePtr
	GetSignals() []string
	RegisterClassDB()
}

func GetClassName(in Wrapped) string {
	return in.GetClassName(in)
}
func GetParentClassName(in Wrapped) string {
	return in.GetParentClassName(in)
}
