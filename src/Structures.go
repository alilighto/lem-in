package src

type antcolonie struct {
	StartRoom string
	EndRoom   string
	Rooms     map[string][]string
	Ants      int
}