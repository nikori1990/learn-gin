package enum

type Flag int

const (
	FALSE Flag = iota
	TRUE
)

func (s Flag) Name() string {
	if s < FALSE || s > TRUE {
		return "Unknown"
	}
	return [...]string{"FALSE", "TRUE"}[s]
}

func (s Flag) Original() int {
	return int(s)
}

func (s Flag) String() string {
	return s.Name()
}
