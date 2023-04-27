package enum

type Status int

const (
	DISABLE Status = iota
	ENABLE
)

func (s Status) Name() string {
	if s < DISABLE || s > ENABLE {
		return "Unknown"
	}
	return [...]string{"DISABLE", "ENABLE"}[s]
}

func (s Status) Original() int {
	return int(s)
}

func (s Status) String() string {
	return s.Name()
}

//func Values() []Status {
//	return []Status{DISABLE, ENABLE}
//}
//
//func ValueOf(name string) (Status, error) {
//	switch name {
//	case "Disable":
//		return DISABLE, nil
//	case "Enable":
//		return ENABLE, nil
//	default:
//		return 0, fmt.Errorf("invalid status name : %s", name)
//	}
//}
