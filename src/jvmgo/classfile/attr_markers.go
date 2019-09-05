package classfile

// see: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.15
// The presence of a Deprecated attribute does not alter the semantics of a class or interface.
/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

// see: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.8
/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
