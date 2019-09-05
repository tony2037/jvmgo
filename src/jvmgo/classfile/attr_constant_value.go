package classfile

// The ConstantValue attribute is a fixed-length attribute in the attributes table of a field_info structure (§4.5). A ConstantValue attribute represents the value of a constant expression (JLS §15.28)
/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length; // must be 2
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
