package secret

var testVersion = 1

func Handshake(code uint) []string {
	decoded := make([]string, 0, 4)
	//1 = wink
	//10 = double blink
	//100 = close your eyes
	//1000 = jump

	var wink uint = 1
	var dblWink uint = 1 << 1
	var clsEyes uint = 1 << 2
	var jump uint = 1 << 3
	var reverse uint = 1 << 4

	if code&wink > 0 {
		decoded = append(decoded, "wink")
	}

	if code&dblWink > 0 {
		decoded = append(decoded, "double blink")
	}

	if code&clsEyes > 0 {
		decoded = append(decoded, "close your eyes")
	}

	if code&jump > 0 {
		decoded = append(decoded, "jump")
	}

	if code&reverse > 0 {
		for left, right := 0, len(decoded)-1; left < right; left, right = left+1, right-1 {
			decoded[left], decoded[right] = decoded[right], decoded[left]
		}
	}
	return decoded
}
