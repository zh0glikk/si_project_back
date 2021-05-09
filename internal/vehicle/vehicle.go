package vehicle

type Location struct {
	X int
	Y int
}

type Vehicle struct {
	Location Location
	Fullness int
	Capacity int
}


