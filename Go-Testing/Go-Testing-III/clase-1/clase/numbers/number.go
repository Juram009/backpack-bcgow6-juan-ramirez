package numbers

func Number(a int) error {
	return nil
}

func TypeNumber(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a > 0:
		return "positive"
	default:
		return "zero"
	}
}
