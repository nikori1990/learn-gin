package enum

type PermissionType int

const (
	PAGE PermissionType = iota + 1
	BUTTON
)

func (pt PermissionType) Name() string {
	if pt < PAGE || pt > BUTTON {
		return "Unknown"
	}
	return [...]string{"PAGE", "BUTTON"}[pt-1]
}

func (pt PermissionType) Original() int {
	return int(pt)
}

func (pt PermissionType) String() string {
	return pt.Name()
}

func PermissionTypeValues() []PermissionType {
	return []PermissionType{PAGE, BUTTON}
}

//func Values() []PermissionType {
//	return []PermissionType{PAGE, BUTTON}
//}
//
//func ValueOf(name string) (PermissionType, error) {
//	switch name {
//	case "PAGE":
//		return PAGE, nil
//	case "BUTTON":
//		return BUTTON, nil
//	default:
//		return 0, fmt.Errorf("invalid PermissionType name : %s", name)
//	}
//}
