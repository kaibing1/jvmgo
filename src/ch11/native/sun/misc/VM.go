package misc

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}

// func initialize(frame *rtda.Frame) {
// 	vmClass := frame.Method().Class()
// 	saveProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
// 	key := heap.JString(vmClass.Loader(), "foo")
// 	val := heap.JString(vmClass.Loader(), "bar")
// 	frame.OperandStack().PushRef(saveProps)
// 	frame.OperandStack().PushRef(key)
// 	frame.OperandStack().PushRef(val)
// 	propsClass := vmClass.Loader().LoadClass("java/util/Properties")
// 	setPropMethod := propsClass.GetInstanceMethod("setProperty",
// 		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
// 	base.InvokeMethod(frame, setPropMethod)
// }
