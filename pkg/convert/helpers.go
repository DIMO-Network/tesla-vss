package convert

func windowStateToIsOpen(state string) float64 {
	if state == "PartiallyOpen" || state == "Opened" {
		return 1
	}
	return 0
}
