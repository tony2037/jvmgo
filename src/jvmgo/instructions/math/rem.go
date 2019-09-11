package math

import "math"
import "jvmgo/instructions/base"
import "jvmgo/rtda"

// see: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.drem
// Remainder double
/*
The result of a drem instruction is not the same as that of the so-called remainder operation defined by IEEE 754. The IEEE 754 "remainder" operation computes the remainder from a rounding division, not a truncating division, and so its behavior is not analogous to that of the usual integer remainder operator. Instead, the Java Virtual Machine defines drem to behave in a manner analogous to that of the Java Virtual Machine integer remainder instructions (irem and lrem); this may be compared with the C library function fmod.
*/
type DREM struct{ base.NoOperandsInstruction }

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2) // todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct{ base.NoOperandsInstruction }

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder int
type IREM struct{ base.NoOperandsInstruction }

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushInt(result)
}

// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
