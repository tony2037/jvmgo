package control

import "jvmgo/instructions/base"
import "jvmgo/rtda"

// see: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-6.html#jvms-6.5.goto
// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
