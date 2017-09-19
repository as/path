// Package path simulates a current working directory in a text
// editing environment. It provides a way to keep track of a
// relative display path that represents a larger absolute path.
//
// The primary operation is 'Look', which inspects a third junction
// to the managed path. The goal of this package is to seperate
// the file system specific nature of relative/absolute path and
// provide a straightforward structure that a text editor can consume.
//
// Please be aware that none of the methods for path.Path have
// pointer references. Their state will not change, and they
// return a new path if necessary.
package path

import "path/filepath"
import "log"
import "os"

// NewPath creates a starting path
func NewPath(path string) (t Path) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	wd = Clean(wd)
	if !filepath.IsAbs(path) {
		t.base = wd
		t.disp = path
	} else {
		t.base = DirOf(Clean(path))
		t.disp = FileOf(Clean(path))
	}
	return t
}

// Path consists of a base path and a display path. The
// display path is what you see in a text editor. The base
// path holds the prefix to that display path so that combining
// them creates a valid absolute path to the file.
//
type Path struct {
	base string
	disp string
}

// Name returns the display name of the path
func (t Path) Name() string {
	log.Printf("name is: %s\n", t.base)
	return t.disp
}

// Blank returns a copy of Path without a display name set
// The path points to the base path
func (t Path) Blank() Path{
	t.disp = ""
	return t
}

func (t Path) Exists() bool{
	return Exists(t.Abs())
}

// Base returns the base path
func (t Path) Base() string{
	log.Printf("base is: %s\n", t.base)
	return t.base
}

// IsDir returns true if the path is a directory
func (t Path) IsDir() bool {
	return IsDir(t.Abs())
}

// Path returns an absolute path of the current state.
func (t Path) Abs() string {
	if filepath.IsAbs(t.disp) {
	log.Printf("abs is: %s\n", t.disp)
		return t.disp
	}
	log.Printf("abs is NOT NOT NOT: %s\n", filepath.Join(t.base, t.disp))
	return filepath.Join(t.base, t.disp)
}

// Look returns a new state. An absolute path returns a state with a new base
// and display path set to that path. A relative path adds on to the existing
// display path unless the path consists of enough double-dots to erase the
// display path. In that case, the state of both base and display path is
// set to the join of the base path and the double-dot path.
func (t Path) Look(dir string) Path {
	if filepath.IsAbs(dir) {
		dir = Clean(dir)
		return Path{DirOf(Clean(dir)), dir}
	}
	t.disp = filepath.Join(t.disp, dir)
	if s := filepath.Join(t.base, t.disp); len(s) < len(t.base) {
		t.base = DirOf(Clean(dir))
		t.disp = s
	}
	log.Printf("%s\n", Path{Clean(t.base), Clean(t.disp)})
	return Path{Clean(t.base), Clean(t.disp)}
}

/*
func main() {
	fmt.Println(NewPath("."))
	fmt.Println(Path{`/windows/system32/`, `drivers/etc/`}.Look("../hosts").Look("..").Look("/"))
}
*/
