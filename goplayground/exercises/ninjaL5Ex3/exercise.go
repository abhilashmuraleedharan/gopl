package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println("truck t:", t.doors, t.color, t.fourWheel)
	fmt.Println("sedan s:", s.doors, s.color, s.luxury)
}
