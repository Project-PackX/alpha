package enums

var AccessLevel = newAccessLevelEnum()

func newAccessLevelEnum() *accessLevelEnum {
	return &accessLevelEnum{
		Normal:  1,
		Courier: 2,
		Admin:   3,
	}
}

type accessLevelEnum struct {
	Normal  uint
	Courier uint
	Admin   uint
}
