package enums

var Sizes = newSizeEnum()

func newSizeEnum() *sizeEnum {
	return &sizeEnum{
		Small:  "Small",
		Medium: "Medium",
		Large:  "Large",
	}
}

type sizeEnum struct {
	Small  string
	Medium string
	Large  string
}
