package stringutil

// NOT逻辑门
func Not(input byte) byte {
	switch input {
	case '1':
		return '0'
	case '0':
		return '1'
	default:
		return 'e'
	}
}

// AND逻辑门
func And(inputFirst, inputSecond byte) byte {
	z := string(inputFirst) + string(inputSecond)
	switch z {
	case "11":
		return '1'
	case "01", "10", "00":
		return '0'
	default:
		return 'e'
	}
}

// OR逻辑门
func Or(inputFirst, inputSecond byte) byte {
	z := string(inputFirst) + string(inputSecond)
	switch z {
	case "11", "01", "10":
		return '1'
	case "00":
		return '0'
	default:
		return 'e'
	}
}

// XOR逻辑门 或 一位加法器
func XOr(inputFirst, inputSecond byte) byte {
	and := And(inputFirst, inputSecond)
	notAnd := Not(and)
	or := Or(inputFirst, inputSecond)
	output := And(notAnd, or)
	return output
}
