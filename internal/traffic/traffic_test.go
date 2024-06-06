package traffic

import "fmt"

func Example_lightState() {
	var state LightState = LightAmber
	fmt.Println(state)
	state = 0
	fmt.Println(state)

	// Output:
	// 3 AMBER
	// 0 RED
}

func Example_lane() {

	l := Lane[Car]{
		queue: []Car{},
	}

	for i := 1; i < 4; i++ {
		c := Car{
			ID: uint(i),
		}
		l.Enqueue(c)
	}
	fmt.Println(l)

	l.Dequeue()
	fmt.Println(l)

	// Output:
}
