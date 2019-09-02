# Java virtual machine implementation based on golang
[ztex]


## Search class
```
To activate a JAVA program, Before calling **main()**, the jvm has to load the relative classes, e.g **java.lang.Object**.

As the resulte, a jvm has to be able to search classes. The question is: where should jvm go to search?

The standard (see [java loader](https://en.wikipedia.org/wiki/Java_Classloader), also [classpath, java](https://en.wikipedia.org/wiki/Classpath_(Java))) does not specify where should JVM go to search class. Different JVMs are allowed to search based on different implementation.

Here, classpath can be seperate into the 3 parts:
* bootstrap classpath
	* default: `jre/lib`, Java standard library (most in `rt.jar`) is here.
* extension classpath
	* `jre/lib/ext`, the classes using Java extension mechanism (https://docs.oracle.com/javase/tutorial/ext/index.html) are here
* user classpath
	* The third-party classes
	* `CLASSPATH` enviroment variable
	* `-classpath/-cp` option, e.g `java -cp path/to/classes;lib/a.jar;lib/b.jar`
```
* `classpath/` is used to implement:
	* Entry interface
	* DirEntry: absolute path, read `.class`
	* ZipEntry: read `ZIP` or `JAR`, which the classes being extracted from
	* CompositeEntry: traverse all
	* WildcardEntry: `*`, all classes, but do not dig into child directories

## Parse `.class`
- [ ] what's gonna do
```
Parse `.class`
reference: https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html
```
- [ ] quotation
```
A class file consists of a stream of 8-bit bytes. All 16-bit, 32-bit, and 64-bit quantities are constructed by reading in two, four, and eight consecutive 8-bit bytes, respectively. Multibyte data items are always stored in big-endian order, where the high bytes come first. In the Java SE platform, this format is supported by interfaces java.io.DataInput and java.io.DataOutput and classes such as java.io.DataInputStream and java.io.DataOutputStream.

This chapter defines its own set of data types representing class file data: The types **u1, u2, and u4 represent an unsigned one-, two-, or four-byte quantity**, respectively. In the Java SE platform, these types may be read by methods such as readUnsignedByte, readUnsignedShort, and readInt of the interface java.io.DataInput.
```
- [ ] The ClassFile Structure
```
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
```
