// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

import (
	"slices"
	"testing"

	"github.com/pat42smith/gotest"
)

// Check that the correct values are returned for a nil error.

func data1() (int, error) {
	return -17, nil
}

func check1(i int) bool {
	return i == -17
}

func data2() (int, string, error) {
	return 99, "IC", nil
}

func check2(i int, s string) bool {
	return i == 99 && s == "IC"
}

func data3() (uint, string, byte, error) {
	return 1023, "whatever", 255, nil
}

func check3(u uint, s string, b byte) bool {
	return u == 1023 && s == "whatever" && b == 255
}

func data4() ([]int, string, rune, uint, error) {
	return []int{3, -5, 9}, "sum", '=', 7, nil
}

func check4(sl []int, s string, r rune, u uint) bool {
	return slices.Equal(sl, []int{3, -5, 9}) && s == "sum" && r == '=' && u == 7
}

func data5() (string, byte, func() int, int, int16, error) {
	return "something", 37, func() int { return 88 }, -1728, 1728, nil
}

func check5(s string, b byte, f func() int, i int, i16 int16) bool {
	return s == "something" && b == 37 && f() == 88 && i == -1728 && i16 == 1728
}

func TestNilExit(t *testing.T) {
	Exit0(nil)
	gotest.Require(t, check1(Exit1(data1())))
	gotest.Require(t, check2(Exit2(data2())))
	gotest.Require(t, check3(Exit3(data3())))
	gotest.Require(t, check4(Exit4(data4())))
	gotest.Require(t, check5(Exit5(data5())))
}

func TestNilLog(t *testing.T) {
	Log0(nil)
	gotest.Require(t, check1(Log1(data1())))
	gotest.Require(t, check2(Log2(data2())))
	gotest.Require(t, check3(Log3(data3())))
	gotest.Require(t, check4(Log4(data4())))
	gotest.Require(t, check5(Log5(data5())))
}

func TestNilPanic(t *testing.T) {
	Panic0(nil)
	gotest.Require(t, check1(Panic1(data1())))
	gotest.Require(t, check2(Panic2(data2())))
	gotest.Require(t, check3(Panic3(data3())))
	gotest.Require(t, check4(Panic4(data4())))
	gotest.Require(t, check5(Panic5(data5())))
}

func TestNilFatal(t *testing.T) {
	var sr gotest.StubReporter

	Fatal0(nil)(&sr)
	sr.Expect(t, false, false, "")

	gotest.Require(t, check1(Fatal1(data1())(&sr)))
	sr.Expect(t, false, false, "")

	gotest.Require(t, check2(Fatal2(data2())(&sr)))
	sr.Expect(t, false, false, "")

	gotest.Require(t, check3(Fatal3(data3())(&sr)))
	sr.Expect(t, false, false, "")

	gotest.Require(t, check4(Fatal4(data4())(&sr)))
	sr.Expect(t, false, false, "")

	gotest.Require(t, check5(Fatal5(data5())(&sr)))
	sr.Expect(t, false, false, "")
}

func TestNilHandle(t *testing.T) {
	Handle0(nil)(nil)
	gotest.Require(t, check1(Handle1(data1())(nil)))
	gotest.Require(t, check2(Handle2(data2())(nil)))
	gotest.Require(t, check3(Handle3(data3())(nil)))
	gotest.Require(t, check4(Handle4(data4())(nil)))
	gotest.Require(t, check5(Handle5(data5())(nil)))
}
