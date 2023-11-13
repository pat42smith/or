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
package or
