// Copyright 2022, Offchain Labs, Inc.
// For license information, see https://github.com/fog/blob/master/LICENSE

package main

import "time"

func main() {
	println("What time is it??")
	println(time.Now().Nanosecond())
	println(time.Now().Nanosecond())
	println(time.Now().Nanosecond())
}
