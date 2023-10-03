package enums

var AccessLevel = newAccessLevelEnum()

func newAccessLevelEnum() *accessLevelEnum {
	return &accessLevelEnum{
		Normal: 1,
		Admin:  2,
	}
}

type accessLevelEnum struct {
	Normal uint
	Admin  uint
}
