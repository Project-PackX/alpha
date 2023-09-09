package enums

var Statuses = newStatusEnum()

func newStatusEnum() *statusEnum {
	return &statusEnum{
		Dispatch:  "Posted for dispatch",     // Sender to locker
		Transit:   "In transit to warehouse", // Locker to warehouse
		Warehouse: "In the warehouse",        // Warehouse
		Delivery:  "In delivery",             // Warehouse to locker
		Delivered: "Delivered",               // Locker to receiver
		Cancelled: "Cancelled",               // Cancelled
	}
}

type statusEnum struct {
	Dispatch  string
	Transit   string
	Warehouse string
	Delivery  string
	Delivered string
	Cancelled string
}
