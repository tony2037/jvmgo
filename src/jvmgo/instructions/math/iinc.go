package math

import "jvmgo/instructions/base"
import "jvmgo/rtda"

// see: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.iinc
// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
