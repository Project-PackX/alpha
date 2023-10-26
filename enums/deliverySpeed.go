package enums

var DeliverySpeeds = newDeliverySpeedEnum()

func newDeliverySpeedEnum() *deliverySpeedsEnum {
	return &deliverySpeedsEnum{
		SameDay:    "SameDay",
		UltraRapid: "UltraRapid",
		Rapid:      "Rapid",
		Standard:   "Standard",
	}
}

type deliverySpeedsEnum struct {
	SameDay    string
	UltraRapid string
	Rapid      string
	Standard   string
}
