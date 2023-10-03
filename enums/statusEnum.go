package enums

var Statuses = newStatusEnum()

func newStatusEnum() *statusEnum {
	// The name of an enum should match the value (rarely in lowercase, if thats the use case)
	// Good job otherwise ;)
	return &statusEnum{
		Dispatch:  "Dispatch",     // Sender to locker
		Transit:   "Transit",      // Locker to warehouse
		Warehouse: "In Warehouse", // Warehouse
		Delivery:  "In Delivery",  // Warehouse to locker
		Delivered: "Delivered",    // Locker to receiver
		Canceled:  "Canceled",     // Canceled
	}
}

type statusEnum struct {
	Dispatch  string
	Transit   string
	Warehouse string
	Delivery  string
	Delivered string
	Canceled  string
}
