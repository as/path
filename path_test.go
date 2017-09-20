package path

import "testing"
import "path/filepath"
import "os"

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

var (
	have, want string
)

func assert(t *testing.T, part string) {
	if have != want {
		msg := "\n\t\twant: %q\n\t\thave: %q\n"
		if part != "" {
			msg = part + ": " + msg
		}
		t.Logf(msg, want, have)
		t.Fail()
	}
}

func test(t *testing.T, nm, h, w string) {
	have, want = h, w
	assert(t, nm)
}

func TestNew(t *testing.T) {
	// cd /etc && acme.exe
	// click on a '.' in a New window
	test(t, "TA01", New(`C:\windows\`).Abs(), `C:\windows`)
	test(t, "TA02", New(`C:\windows\`).Look(`.`).Abs(), `C:\windows`)
	test(t, "TA03", New(`C:\windows\`).Look(`..`).Abs(), `C:\`)
	test(t, "TA04", New(`C:\windows\`).Look(`..`).Abs(), `C:\`)
	test(t, "TA05", New(`C:\windows\`).Look(`C:\`).Abs(), `C:\`)

	test(t, "TA06", New(`C:\windows\`).Look(`system32\drivers\etc`).Abs(), `C:\windows\system32\drivers\etc`)
	test(t, "TA07", New(`C:\windows\`).Look(`system32\drivers\etc\..`).Abs(), `C:\windows\system32\drivers`)
	test(t, "TA08", New(`C:\windows\`).Look(`system32\drivers\etc\..\..`).Abs(), `C:\windows\system32`)
	test(t, "TA09", New(`C:\windows\`).Look(`system32\drivers\etc\..\..\..`).Abs(), `C:\windows`)
	test(t, "TA0A", New(`C:\windows\`).Look(`system32\drivers\etc\..\..\..\..`).Abs(), `C:\`)

	test(t, "TX01", New(`C:\windows\`).Look(`system32\drivers\etc`).Name(), `system32\drivers\etc`)
	test(t, "TX02", New(`C:\windows\`).Look(`system32\drivers\etc\..`).Name(), `system32\drivers`)
	test(t, "TX03", New(`C:\windows\`).Look(`system32\drivers\etc\..\..`).Name(), `system32`)
	test(t, "TX04", New(`C:\windows\`).Look(`system32\drivers\etc\..\..\..`).Name(), `windows`)
	test(t, "TX05", New(`C:\windows\`).Look(`system32\drivers\etc\..\..\..\..`).Name(), `C:\`)

	test(t, "TY01", New(`C:\windows\`).Look(`system32\`).Look(`drivers\`).Look(`etc`).Name(), `system32\drivers\etc`)
	test(t, "TY02", New(`C:\windows\`).Look(`system32\`).Look(`drivers\`).Look(`etc`).Look(`..`).Name(), `system32\drivers`)
	test(t, "TY03", New(`C:\windows\`).Look(`system32\`).Look(`drivers\`).Look(`etc`).Look(`..`).Look(`..`).Name(), `system32`)
	test(t, "TY04", New(`C:\windows\`).Look(`system32\`).Look(`drivers\`).Look(`etc`).Look(`..`).Look(`..`).Look(`..`).Name(), `windows`)
	test(t, "TY05", New(`C:\windows\`).Look(`system32\`).Look(`drivers\`).Look(`etc`).Look(`..`).Look(`..`).Look(`..`).Look(`..`).Name(), `C:\`)
}
func TestNewDot(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	// acme.exe /etc
	wd := New(`.`)
	have, want = wd.Abs(), cwd
	assert(t, "abs")
}

func TestWorkflow(t *testing.T) {
	wd := New(`C:\windows\system32\drivers`)
	hosts := wd.Look(`etc\hosts`)
	if !hosts.Exists() {
		t.Logf("host file unreachable at %q\n", hosts.Abs())
	}
	if hosts.IsDir() {
		t.Logf("hosts file is a directory %q\n", hosts.Abs())
	}
	if hosts.Name() != `etc\hosts` {
		t.Logf("hosts file display name is not etc\\hosts, have %q\n", hosts.Name())
	}
}
