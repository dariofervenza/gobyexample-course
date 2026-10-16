package main

import (
	"fmt"
	"time"
)

func main() {

	// for loop represent the only iteration in go: for := for + while + loop

	i := 1

	// while equivalent
	for i <= 3 {
		fmt.Println("This is i:", i)
		i = i + 1
		i += 1
	}

	// classic for
	for d := 0; d < 5; d++ {
		fmt.Println("This is d:", d)
	}

	// loop equivalent

	for {
		break
	}

	// in range equivalent

	for r := range 6 {
		if r%2 == 0 {
			continue
		}
		fmt.Println("r is:", r)
	}

	// if else

	if 8%2 == 0 || 5-3 == 2 {
		fmt.Println("Gay")
	} else if 2-1 == -3 {
		fmt.Println("No gay")
	} else {
		fmt.Println("sUPER GAY")
	}

	// declare variable -> vlaid in the if else scope

	if if_var := 30; if_var < 20 {
		return
	} else if if_var < 40 {
		fmt.Println("Is less than 40")
	}

	// switch

	s_var := 50
	switch s_var {
	case 1:
		fmt.Println("switch 1")
	case 2, 50:
		fmt.Println("swithc 2 or 50")
	default:
		fmt.Println("switch rest")
	}

	switch time.Now().Weekday() {
	case time.Tuesday, time.Thursday:
		fmt.Println("Working at the office")
	default:
		fmt.Println("rest of the days")
	}

	// can be used as if, without indicating the variable
	t_n := time.Now()
	switch {
	case t_n.Hour() < 12:
		fmt.Println("Before noon")
	}

	// TYPE SWITCH
	whichIsMyType := func(val any) {
		// to get the type you could use also use reflect.TypeOf(val)
		switch t := val.(type) {
		case bool:
			fmt.Println("Im a boolean")
		case int:
			fmt.Println("im an integer")
		default:
			// notice the printf and also
			// t will hold the value not the type
			// if you want the type you need the %T
			// also without the \n printf does not add line jumps
			// NOTICE that %T with val would print the type
			fmt.Printf("I dont know the typeof: %T\n", t)
		}
	}
	whichIsMyType("hey")
	whichIsMyType(true)

}
