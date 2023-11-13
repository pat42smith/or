// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

import (
	"log"
)

// Log0, if e is not nil, calls log.Fatal(e) to report the error and terminate the process.
func Log0(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Log1 is like Log0, but returns t1 when e is nil.
func Log1[T1 any](t1 T1, e error) T1 {
	Log0(e)
	return t1
}

// Log2 is like Log0, but returns (t1, t2) when e is nil.
func Log2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	Log0(e)
	return t1, t2
}

// Log3 is like Log0, but returns (t1, t2, t3) when e is nil.
func Log3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, e error) (T1, T2, T3) {
	Log0(e)
	return t1, t2, t3
}

// Log4 is like Log0, but returns (t1, t2, t3, t4) when e is nil.
func Log4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, e error) (T1, T2, T3, T4) {
	Log0(e)
	return t1, t2, t3, t4
}

// Log5 is like Log0, but returns (t1, t2, t3, t4, t5) when e is nil.
func Log5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, e error) (T1, T2, T3, T4, T5) {
	Log0(e)
	return t1, t2, t3, t4, t5
}
