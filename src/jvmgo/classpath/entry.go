package classpath

import "os"
import "strings"

// :(linux/unix) or ;(windows)
// see: https://golang.org/pkg/os/
const pathListSeparator = string(os.PathListSeparator)

// Define entry interface, and there will be four different implementation: DirEntry, ZipEntry, CompositeEntry and WildcardEntry
type Entry interface {
	// className: fully/qualified/ClassName.class
	// []byte are byte codes
	readClass(className string) ([]byte, Entry, error)
	// just like toString()
	String() string
}

// Given a path, determine which situation it is, and as the result has different implementations
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
