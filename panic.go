// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

// Panic0 calls panic(e) if e is non-nil.
func Panic0(e error) {
	if e != nil {
		panic(e)
	}
}

// Panic1 calls panic(e) if e is non-nil. Otherwise, it returns t1.
func Panic1[T1 any](t1 T1, e error) T1 {
	Panic0(e)
	return t1
}

// Panic2 calls panic(e) if e is non-nil. Otherwise, it returns (t1, t2).
func Panic2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	Panic0(e)
	return t1, t2
}

// Panic3 calls panic(e) if e is non-nil. Otherwise, it returns (t1, t2, t3).
func Panic3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, e error) (T1, T2, T3) {
	Panic0(e)
	return t1, t2, t3
}

// Panic4 calls panic(e) if e is non-nil. Otherwise, it returns (t1, t2, t3, t4).
func Panic4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, e error) (T1, T2, T3, T4) {
	Panic0(e)
	return t1, t2, t3, t4
}

// Panic5 calls panic(e) if e is non-nil. Otherwise, it returns (t1, t2, t3, t4, t5).
func Panic5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, e error) (T1, T2, T3, T4, T5) {
	Panic0(e)
	return t1, t2, t3, t4, t5
}
