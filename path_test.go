package path

import "testing"
import "path/filepath"

//TODO(as): complete tests

func TestPath(t *testing.T) {
	filepath.IsAbs("/")
	have := Path{`/windows/system32/`, `drivers/etc/`}.Look("hosts")
	input := "hosts"
	want := Path{`/windows/system32/`, `drivers/etc/hosts`}
	if have.Look(input) != want {
		t.Logf("\n\t\thave: %q\n\t\twant: %q\n", have, want)
	}
}
