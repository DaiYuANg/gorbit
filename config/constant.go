package config

type Mode = int

const (
	Dev  Mode = iota
	Prod Mode = iota
	Test Mode = iota
)
