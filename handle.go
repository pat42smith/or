// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

// Handler is the type of a function to handle errors.
type Handler func(error)

// Handle0(e)(h), when e is not nil, calls h(e).
//
// If the call to h(e) returns, Handle0(e)(h) will panic.
// This guarantees that Handle0(e)(h) will not return unless e is nil.
func Handle0(e error) func(h Handler) {
	if e == nil {
		return func(Handler) {}
	}
	return func(h Handler) {
		h(e)
		panic("Error handler returned")
	}
}

// Handle1(t1, e)(h) is like Handle0(e)(h), but returns t1 when e is nil.
func Handle1[T1 any](t1 T1, e error) func(h Handler) T1 {
	if e == nil {
		return func(Handler) T1 {
			return t1
		}
	}
	return func(h Handler) T1 {
		h(e)
		panic("Error handler returned")
	}
}

// Handle2(t1, t2, e)(h) is like Handle0(e)(h), but returns (t1, t2) when e is nil.
func Handle2[T1, T2 any](t1 T1, t2 T2, e error) func(h Handler) (T1, T2) {
	if e == nil {
		return func(Handler) (T1, T2) {
			return t1, t2
		}
	}
	return func(h Handler) (T1, T2) {
		h(e)
		panic("Error handler returned")
	}
}

// Handle3(t1, t2, t3, e)(h) is like Handle0(e)(h), but returns (t1, t2, t3) when e is nil.
func Handle3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, e error) func(h Handler) (T1, T2, T3) {
	if e == nil {
		return func(Handler) (T1, T2, T3) {
			return t1, t2, t3
		}
	}
	return func(h Handler) (T1, T2, T3) {
		h(e)
		panic("Error handler returned")
	}
}

// Handle4(t1, t2, t3, t4, e)(h) is like Handle0(e)(h), but returns (t1, t2, t3, t4) when e is nil.
func Handle4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, e error) func(h Handler) (T1, T2, T3, T4) {
	if e == nil {
		return func(Handler) (T1, T2, T3, T4) {
			return t1, t2, t3, t4
		}
	}
	return func(h Handler) (T1, T2, T3, T4) {
		h(e)
		panic("Error handler returned")
	}
}

// Handle5(t1, t2, t3, t4, t5, e)(h) is like Handle0(e)(h), but returns (t1, t2, t3, t4), t5 when e is nil.
func Handle5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, e error) func(h Handler) (T1, T2, T3, T4, T5) {
	if e == nil {
		return func(Handler) (T1, T2, T3, T4, T5) {
			return t1, t2, t3, t4, t5
		}
	}
	return func(h Handler) (T1, T2, T3, T4, T5) {
		h(e)
		panic("Error handler returned")
	}
}
