package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetLong(slotId, stack.PopLong())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
