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
