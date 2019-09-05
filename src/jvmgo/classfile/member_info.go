package classfile

// [field_info] see: https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.5
// [methond_info] see: https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.6
/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
// For the sake of convenience, implement struct MebmberInfo for the both

type MemberInfo struct {
	cp              ConstantPool // see https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4
	accessFlags     uint16 // related to different access permission, e.g private, public, static ...
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// read field or method table
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
