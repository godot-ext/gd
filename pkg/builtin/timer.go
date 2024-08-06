package builtin

import (
	. "github.com/godot-go/godot-go/pkg/constant"
	"github.com/godot-go/godot-go/pkg/log"
	"go.uber.org/zap"
)

// NewCallable, index: 2
func NewCallableWithGoString(object Object, methodGoStr string) Callable {
	method := NewStringNameWithUtf8Chars(methodGoStr)
	return NewCallableWithObjectStringName(object, method)
}
func DelayCallTimer(object Object, methodGoStr string, timer Timer) {
	callable := NewCallableWithGoString(object, methodGoStr)
	defer callable.Destroy()
	err := timer.Connect_StrExt("timeout", callable, OBJECT_CONNECT_FLAGS_CONNECT_ONE_SHOT)
	if err != OK {
		log.Panic("message timer connect failure", zap.Any("error", err))
	}
}

func DelayCall(object Node, methodGoStr string, time_sec float64) {
	sceneTreeTimer := object.GetTree().CreateTimer(time_sec, true, false, false).TypedPtr()
	callable := NewCallableWithGoString(object, methodGoStr)
	defer callable.Destroy()
	err := sceneTreeTimer.Connect_StrExt("timeout", callable, OBJECT_CONNECT_FLAGS_CONNECT_ONE_SHOT)
	if err != OK {
		log.Panic("message timer connect failure", zap.Any("error", err))
	}
}
