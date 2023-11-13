// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pat42smith/gotest"
)

// Check that Exitn and Logn handle non-nil errors correctly.
// This requires a separate process.

var test_source = `package main

import (
	"errors"
	"log"
	"os"
	"github.com/pat42smith/or"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("logged: ")

	switch os.Args[1] {
	case "exit":
		switch os.Args[2] {
		case "0":
			or.Exit0(errors.New("zero"))
		case "1":
			or.Exit1(1, errors.New("one"))
		case "2":
			or.Exit2(1, 2, errors.New("two"))
		case "3":
			or.Exit3(1, 2, 3, errors.New("three"))
		case "4":
			or.Exit4(1, 2, 3, 4, errors.New("four"))
		case "5":
			or.Exit5(1, 2, 3, 4, 5, errors.New("five"))
		}
	case "log":
		switch os.Args[2] {
		case "0":
			or.Log0(errors.New("zero"))
		case "1":
			or.Log1(1, errors.New("one"))
		case "2":
			or.Log2(1, 2, errors.New("two"))
		case "3":
			or.Log3(1, 2, 3, errors.New("three"))
		case "4":
			or.Log4(1, 2, 3, 4, errors.New("four"))
		case "5":
			or.Log5(1, 2, 3, 4, 5, errors.New("five"))
		}
	}
}`

func checkExitLog(t *testing.T, program, which, n, expect string) {
	t.Helper()
	cmd := gotest.Command(program, which, n)
	cmd.WantStderr(expect + "\n")
	cmd.WantCode(1)
	cmd.Run(t, "")
}

func TestExitLog(t *testing.T) {
	tmp := t.TempDir()
	testgo := filepath.Join(tmp, "test.go")
	testexe := filepath.Join(tmp, "test.exe")
	if e := os.WriteFile(testgo, []byte(test_source), 0444); e != nil {
		t.Fatal(e)
	}
	gotest.Command("go", "build", "-o", testexe, testgo).Run(t, "")

	checkExitLog(t, testexe, "exit", "0", "zero")
	checkExitLog(t, testexe, "exit", "1", "one")
	checkExitLog(t, testexe, "exit", "2", "two")
	checkExitLog(t, testexe, "exit", "3", "three")
	checkExitLog(t, testexe, "exit", "4", "four")
	checkExitLog(t, testexe, "exit", "5", "five")

	checkExitLog(t, testexe, "log", "0", "logged: zero")
	checkExitLog(t, testexe, "log", "1", "logged: one")
	checkExitLog(t, testexe, "log", "2", "logged: two")
	checkExitLog(t, testexe, "log", "3", "logged: three")
	checkExitLog(t, testexe, "log", "4", "logged: four")
	checkExitLog(t, testexe, "log", "5", "logged: five")
}
