// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

import (
	"errors"
	"testing"

	"github.com/pat42smith/gotest"
)

// Check that non-nil errors are handled correctly.
// Exit and Log need to be tested in a separate process,
// so they are not included here.

var disaster = errors.New("disaster")

func jabber0() error {
	return disaster
}

func jabber1() (string, error) {
	return "'Twas", disaster
}

func jabber2() (string, string, error) {
	return "brillig,", "and", disaster
}

func jabber3() (string, string, string, error) {
	return "the", "slithy", "toves", disaster
}

func jabber4() (string, string, string, string, error) {
	return "Did", "gyre", "and", "gimble", disaster
}

func jabber5() (string, string, string, string, string, error) {
	return "in", "the", "wabe;", "All", "mimsy", disaster
}

func checkPanic(t *testing.T, f func()) {
	t.Helper()
	p := gotest.MustPanic(t, f)
	gotest.Expect(t, disaster, p.(error))
}

func TestErrorPanic(t *testing.T) {
	checkPanic(t, func() {
		Panic0(jabber0())
	})
	checkPanic(t, func() {
		Panic1(jabber1())
	})
	checkPanic(t, func() {
		Panic2(jabber2())
	})
	checkPanic(t, func() {
		Panic3(jabber3())
	})
	checkPanic(t, func() {
		Panic4(jabber4())
	})
	checkPanic(t, func() {
		Panic5(jabber5())
	})
}

func checkFatal(t *testing.T, f func(*gotest.StubReporter)) {
	t.Helper()
	var sr gotest.StubReporter
	// StubReporter.Fatal actually returns, so Fataln should panic.
	gotest.MustPanic(t, func() {
		f(&sr)
	})
	sr.Expect(t, true, true, "disaster\n")
}

func TestErrorFatal(t *testing.T) {
	checkFatal(t, func(sr *gotest.StubReporter) {
		Fatal0(jabber0())(sr)
	})
	checkFatal(t, func(sr *gotest.StubReporter) {
		Fatal1(jabber1())(sr)
	})
	checkFatal(t, func(sr *gotest.StubReporter) {
		Fatal2(jabber2())(sr)
	})
	checkFatal(t, func(sr *gotest.StubReporter) {
		Fatal3(jabber3())(sr)
	})
	checkFatal(t, func(sr *gotest.StubReporter) {
		Fatal4(jabber4())(sr)
	})
	checkFatal(t, func(sr *gotest.StubReporter) {
		Fatal5(jabber5())(sr)
	})
}

func checkHandle(t *testing.T, f func(Handler)) {
	var saved error
	p := gotest.MustPanic(t, func() {
		f(func(e error) {
			saved = e
			panic(e)
		})
	})
	gotest.Require(t, saved == disaster && p.(error) == disaster)

	saved = nil
	p = gotest.MustPanic(t, func() {
		f(func(e error) {
			saved = e
		})
	})
	gotest.Require(t, saved == disaster && p.(string) == "Error handler returned")
}

func TestErrorHandle(t *testing.T) {
	checkHandle(t, func(h Handler) {
		Handle0(jabber0())(h)
	})
	checkHandle(t, func(h Handler) {
		Handle1(jabber1())(h)
	})
	checkHandle(t, func(h Handler) {
		Handle2(jabber2())(h)
	})
	checkHandle(t, func(h Handler) {
		Handle3(jabber3())(h)
	})
	checkHandle(t, func(h Handler) {
		Handle4(jabber4())(h)
	})
	checkHandle(t, func(h Handler) {
		Handle5(jabber5())(h)
	})
}
