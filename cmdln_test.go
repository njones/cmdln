package cmdln

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func ExampleSplit() {
	// here is a somewhat complex command line to execute
	str := `echo -e 'Starting LS\n===========' && ls -la && echo -e "===========\nI'm Done."`

	// split the string, so we can run this on
	// the command line...
	cmmd, args := Split(str)

	// you should be checking errors, even though
	// I don't here. run it!
	o, _ := exec.Command(cmmd, args...).Output()

	// print out the current directory listing
	// with a header and footer
	fmt.Println(string(o))
}

func TestCommand1(t *testing.T) {
	want1, want2 := "echo", []string{"-e", `"My name is earl\n"`}

	have1, have2 := Split(`echo -e "My name is Earl\n"`)
	if want1 != have1 && !reflect.DeepEqual(want2, have2) {
		t.Errorf("\nwant1: %q\nhave1: %q", want1, have1)
		t.Errorf("\nwant2: %q\nhave2: %q", want2, have2)
	}
}

func TestCommand2(t *testing.T) {
	want1, want2 := "echo", []string{"-e", `"My name is earl\n"`}

	have1, have2 := Splitf(`echo %s "%s"`, "-e", "My name is Earl\n")
	if want1 != have1 && !reflect.DeepEqual(want2, have2) {
		t.Errorf("\nwant1: %q\nhave1: %q", want1, have1)
		t.Errorf("\nwant2: %q\nhave2: %q", want2, have2)
	}
}

func TestSplit(t *testing.T) {
	base := []string{
		"a b c",
		`a "b c"`,
		`a 'b c'`,
		`a 'b c' "d e" f`,
		`a "b'c"`,
		`a 'b" c'`,
		`a "b\" \"c"`,
		`a "b c \\"`,
	}
	want := [][]string{
		{"a", "b", "c"},
		{"a", `"b c"`},
		{"a", `'b c'`},
		{"a", `'b c'`, `"d e"`, "f"},
		{"a", `"b'c"`},
		{"a", `'b" c'`},
		{"a", `"b\" \"c"`},
		{"a", `"b c \\"`},
	}

	for i, s := range base {
		cmdln := &line{}
		have := strings.FieldsFunc(s, cmdln.SplitFunc)
		if !reflect.DeepEqual(want[i], have) {
			t.Errorf("\nwant: %q\nhave: %q", want[i], have)
		}
	}
}
