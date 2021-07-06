package references

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arRef := stack.PopRef()
	if arRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arRef.ArrayLength()
	stack.PushInt(arrLen)
}
