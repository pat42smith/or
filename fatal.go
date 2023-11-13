// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

// Type Fataler has the methods we need from a *testing.T (or .B or .F).
type Fataler interface {
	Helper()
	Fatal(...any)
}

// Fatal0(e)(f), when e is not nil, aborts the current test case by calling f.Fatal(e).
//
// If the call to f.Fatal0(e) returns, Fatal0(e)(f) will panic. This is not possible
// in the typical case that f is a *testing.T, but may be possible if f is from
// some other implementation of Fataler.
func Fatal0(e error) func(f Fataler) {
	if e == nil {
		return func(Fataler) {}
	}
	return func(f Fataler) {
		f.Helper()
		f.Fatal(e)
		panic("Fatal returned")
	}
}

// Fatal1(t1, e)(f) is like Fatal0(e)(f), but returns t1 when e is nil.
func Fatal1[T1 any](t1 T1, e error) func(f Fataler) T1 {
	if e == nil {
		return func(Fataler) T1 {
			return t1
		}
	}
	return func(f Fataler) T1 {
		f.Helper()
		f.Fatal(e)
		panic("Fatal returned")
	}
}

// Fatal2(t1, t2, e)(f) is like Fatal0(e)(f), but returns (t1, t2) when e is nil.
func Fatal2[T1, T2 any](t1 T1, t2 T2, e error) func(f Fataler) (T1, T2) {
	if e == nil {
		return func(Fataler) (T1, T2) {
			return t1, t2
		}
	}
	return func(f Fataler) (T1, T2) {
		f.Helper()
		f.Fatal(e)
		panic("Fatal returned")
	}
}

// Fatal3(t1, t2, t3, e)(f) is like Fatal0(e)(f), but returns (t1, t2, t3) when e is nil.
func Fatal3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, e error) func(f Fataler) (T1, T2, T3) {
	if e == nil {
		return func(Fataler) (T1, T2, T3) {
			return t1, t2, t3
		}
	}
	return func(f Fataler) (T1, T2, T3) {
		f.Helper()
		f.Fatal(e)
		panic("Fatal returned")
	}
}

// Fatal4(t1, t2, t3, t4, e)(f) is like Fatal0(e)(f), but returns (t1, t2, t3, t4) when e is nil.
func Fatal4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, e error) func(f Fataler) (T1, T2, T3, T4) {
	if e == nil {
		return func(Fataler) (T1, T2, T3, T4) {
			return t1, t2, t3, t4
		}
	}
	return func(f Fataler) (T1, T2, T3, T4) {
		f.Helper()
		f.Fatal(e)
		panic("Fatal returned")
	}
}

// Fatal5(t1, t2, t3, t4, t5, e)(f) is like Fatal0(e)(f), but returns (t1, t2, t3, t4), t5 when e is nil.
func Fatal5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, e error) func(f Fataler) (T1, T2, T3, T4, T5) {
	if e == nil {
		return func(Fataler) (T1, T2, T3, T4, T5) {
			return t1, t2, t3, t4, t5
		}
	}
	return func(f Fataler) (T1, T2, T3, T4, T5) {
		f.Helper()
		f.Fatal(e)
		panic("Fatal returned")
	}
}
