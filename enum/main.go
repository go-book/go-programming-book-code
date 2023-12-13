package main

import (
	"fmt"
)

type Status int

const (
	minStatus Status = iota // 0
	Success                 // 1
	Failed                  // 2
	Forbidden               // 3
	_                       // 4
	_                       // 5
	maxStatus               // 6
)

const (
	StatusSuccess   = iota // 0
	StatusFailed           // 1
	StatusForbidden        // 2
	_                      // 3
	_                      // 4
)

// String 实现fmt.Stringer接口
func (s Status) String() string {
	switch s {
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	case Forbidden:
		return "Forbidden"
	default:
		return "Unknown"
	}
}

func main() {
	for status := minStatus; status <= maxStatus; status++ {
		fmt.Printf("%d->%v\n", status, status)
	}
}
