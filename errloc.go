package errloc

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
)

var (
	// RecordSep delineates error records
	RecordSep = "\n"
	// UnitSep is used to make the error more readable
	UnitSep = "\n\t"
)

// Location is the name:line of a file. Ideally returned by Here(). In usage
// it'll give you the file:line of the invocation of Here() to be passed as part
// of the error.
type Location string

// Here returns the file:line at the point of invocation. This is purely sugar.
func Here() Location {
	return here(2)
}

// here returns the file:line at the point of invocation
func here(depth int) Location {
	var l Location
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		path := filepath.Base(file)
		l = Location(path + ":" + strconv.Itoa(line))
	}
	return l
}

// New creates a new Error of our own liking. The e is assumed
// to be the error message.
func New(e string) error {
	return fmt.Errorf("%s%s%s%s", here(2), UnitSep, e, RecordSep)
}

// AddLoc adds a location value to an existing error
func AddLoc(e error) error {
	return fmt.Errorf("%s: %w", here(2), e)
}
