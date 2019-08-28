package errmsg

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
)

// RecordSep delineates error records
const RecordSep = "\036" // byte(30) || "\x1e" ... is the ascii Record Separator (RS) character

// Location is the name:line of a file. Ideally returned by Here(). In usage
// it'll give you the file:line of the invocation of Here() to be passed as part
// of the error.
type Location string

// Here returns the file:line at the point of invocation
func Here() Location {
	var l Location
	_, file, line, ok := runtime.Caller(1)
	if ok {
		path := filepath.Base(file)
		l = Location(path + ":" + strconv.Itoa(line))
	}
	return l
}

// here returns the file:line at the point of invocation
func here() Location {
	var l Location
	_, file, line, ok := runtime.Caller(2)
	if ok {
		path := filepath.Base(file)
		l = Location(path + ":" + strconv.Itoa(line))
	}
	return l
}

// New creates a new Error of our own liking. The `string` args are assumed
// to be the error message. The `error`/`Error` arg is assumed to be a Prev.
// The `Location` arg is assumed to be the Location. The `Kind` arg is the
// Kind of the error.
func New(msg string) error {
	return fmt.Errorf("%s [%s]%s", msg, here(), RecordSep)
}
