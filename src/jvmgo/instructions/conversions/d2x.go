package conversions

import "jvmgo/instructions/base"
import "jvmgo/rtda"

/*
The value on the top of the operand stack must be of type double. It is popped from the operand stack and undergoes value set conversion (§2.8.3) resulting in value'. Then value' is converted to a float result using IEEE 754 round to nearest mode. The result is pushed onto the operand stack.
*/
// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.d2f
// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
