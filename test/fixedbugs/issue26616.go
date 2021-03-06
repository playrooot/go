// errorcheck

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

var x int = three() // ERROR "1 variable but 3 values"

func f() {
	var _ int = three() // ERROR "1 variable but 3 values"
	var a int = three() // ERROR "1 variable but 3 values"
	a = three()         // ERROR "1 variable but 3 values"
	b := three()        // ERROR "1 variable but 3 values"

	_, _ = a, b
}

func three() (int, int, int)
