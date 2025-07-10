package framework

type Mode = int

const (
	Dev  Mode = iota
	Prod Mode = iota
	Test Mode = iota
)
