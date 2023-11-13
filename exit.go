// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

package or

import (
	"fmt"
	"os"
)

// Exit0, if e is not nil, prints e to standard error and then calls os.Exit(1).
func Exit0(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

// Exit1 is like Exit0, but returns t1 when e is nil.
func Exit1[T1 any](t1 T1, e error) T1 {
	Exit0(e)
	return t1
}

// Exit2 is like Exit0, but returns (t1, t2) when e is nil.
func Exit2[T1, T2 any](t1 T1, t2 T2, e error) (T1, T2) {
	Exit0(e)
	return t1, t2
}

// Exit3 is like Exit0, but returns (t1, t2, t3) when e is nil.
func Exit3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, e error) (T1, T2, T3) {
	Exit0(e)
	return t1, t2, t3
}

// Exit4 is like Exit0, but returns (t1, t2, t3, t4) when e is nil.
func Exit4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, e error) (T1, T2, T3, T4) {
	Exit0(e)
	return t1, t2, t3, t4
}

// Exit5 is like Exit0, but returns (t1, t2, t3, t4, t5) when e is nil.
func Exit5[T1, T2, T3, T4, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, e error) (T1, T2, T3, T4, T5) {
	Exit0(e)
	return t1, t2, t3, t4, t5
}
