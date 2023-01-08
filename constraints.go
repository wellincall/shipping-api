package main

import "fmt"

func WeightConstraint(s *ShipmentInput) (float64, string) {
	switch {
	case 0.0 < s.Weight && s.Weight < 10.0:
		return 100, ""
	case 10.0 <= s.Weight && s.Weight < 25.0:
		return 300, ""
	case 25.0 <= s.Weight && s.Weight < 50.0:
		return 500, ""
	case 50.0 <= s.Weight && s.Weight < 1000.0:
		return 2000, ""
	case s.Weight < 0.0:
		return 0, "Weight cannot be negative"
	case s.Weight > 1000.0:
		return 0, "Cannot ship items with more than 1000Kg"
	default:
		return 0, "No weight constraint possible for this case"
	}
}

func RegionConstraint(s *ShipmentInput) (float64, string) {
	euMembers := getEuMembers()
	senderInEu, validSender := euMembers[s.Sender]
	receiverInEu, validReceiver := euMembers[s.Receiver]

	if !validSender || !validReceiver {
		return 0, fmt.Sprintf("%v or %v are invalid country codes", s.Sender, s.Receiver)
	}

	if s.Sender == s.Receiver {
		return 1, ""
	}
	if senderInEu && receiverInEu {
		return 1.5, ""
	}
	if !senderInEu || !receiverInEu {
		return 2.5, ""
	}
	return 0, "No region constraint possible for this case"
}
