package classfile

// see: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.5
/*
Exceptions_attribute {
    u2 attribute_name_index; // To a constant_pool entry that must be the CONSTANT_Utf8_info representing "Exceptions"
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
