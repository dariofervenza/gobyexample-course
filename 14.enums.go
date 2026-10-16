package main

import "fmt"

// Enums := Enumerated types

// its a fixed type with fixed number of values, each with a distinct value
// go doesnt have a distinct language feature for them but they can be implemented using existing language idioms

// our enum ServerState has inside an int
type ServerState int

// define the possible vlaues of the enum with const

// the keyword iota generates succesive constant values --> 0, 1, 2 and so on...
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// we implement the fmt.Stringer (https://pkg.go.dev/fmt#Stringer) interface
// so values of ServerState can be printed as strings

func (ss ServerState) String() string {
	return stateName[ss]
}

// that can get cumbersome if there are many possible values. In such cases the stringer tool can be used in conjunction with go:generate to automate the process. See this post for a longer explanation.
// see https://pkg.go.dev/golang.org/x/tools/cmd/stringer
// and https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unkown state: %s", s))
	}
}

func main() {
	// if we have a type int we cannot pass it to transition
	// the Stringer interface helps us with this mismatch
	// and provides some degree of compite-time type safety for enums
	ns := transition(StateIdle)
	fmt.Println("Initial: ", ns)

	ns2 := transition(ns)
	fmt.Println("\nAfter change: ", ns2)

}
