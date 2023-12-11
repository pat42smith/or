// Copyright 2023 Patrick Smith
// Use of this source code is subject to the MIT-style license in the LICENSE file.

// Package or provides some functions to simplify error handling in Go.
//
// For example, instead of
//
//	f, err := os.Open(filename)
//	if err != nil {
//	        panic(err)
//	}
//
// we can write
//
//	f := or.Panic1(os.Open(filename))
//
// The main benefit of using these functions is the elimination of the error
// variable, and so the possibility of interference with nearby code handling
// errors from other function calls. In addition, the code becomes a bit
// shorter, though opinions will vary as to how much this matters.
//
// The functions here are named in the format Xxxxn, where Xxxx describes what
// the function does in case of non-nil error, and n is the number of values
// expected before the trailing error. If the error value is nil, the function
// returns those n values.
//
// All of these functions guarantee that they will not return normally when
// presented with a non-nil error.
//
// It would be preferable to write the parameter lists for the Fataln and Handlen
// functions in the opposite order to that given here. For example,
//
// func Fatal1[T1 any](f Fataler) func(t1 T1, e error) T1
//
// However, current versions of Go (1.21.5 as of this writing) do not allow this
// function to be used without explicit instantiation, as documented in
// [this proposal](https://github.com/golang/go/issues/64197#issuecomment-1814245087).
// If that proposal is implemented, then the functions here may be changed to
// the preferred order.
//
// It is technically possible to add a set of Returnn functions that would simulate
// a return from the containing function to return when the error is not nil.
// But this would use panic(), and seems to be a misuse. Also, it would require
// explicit deferred code in the containing function. For these reasons, the Returnn
// functions have been omitted.
package or
