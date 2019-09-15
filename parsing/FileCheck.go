package parsing

func isComment(char uint8) bool {
	if char == '#' {
		return true
	}
	return false
}

func iterateComment(f string, i int, len int) int {
	infinity := true
	for infinity {
		if isDigit(f[i]) {
			break
		}
		i++
		if i >= len {
			break
		}
	}
	return i
}

func isBlank(char uint8) bool {
	if char == ' ' {
		return true
	}
	return false
}

func iterateBlank(f string, i int, len int) int {
	infinity := true
	for infinity {
		if isBlank(f[i]) == false {
			break
		}
		i++
		if i >= len {
			break
		}
	}
	return i
}

func isDigit(char uint8) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func iterateDigit(f string, i int, len int) int {
	infinity := true
	for infinity {
		if isDigit(f[i]) == false {
			break
		}
		i++
		if i >= len {
			break
		}
	}
	return i
}

func isNewLine(char uint8) bool {
	if char == 10 {
		return true
	}
	return false
}

func FileIsValid(file []byte) bool {
	f := string(file)
	len := len(f)
	if len < 21 {
		return false
	}
	i := 0
	// // iterate past comment
	if isComment(f[i]) {
		i = iterateComment(f, i, len)
	}
	if i >= len {
		return false
	}
	// sizes
	if isDigit(f[i]) == false {
		return false
	}
	size := int(f[i]) - 48
	i++
	if i >= len {
		return false
	}
	if isDigit(f[i]) == true {
		size = size*10 + (int(f[i]) - 48)
		i++
		if i >= len {
			return false
		}
	}
	if isNewLine(f[i]) == false {
		return false
	}
	i++
	if i >= len {
		return false
	}
	// Board
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			i = iterateBlank(f, i, len)
			if i >= len {
				return false
			}
			if isDigit(f[i]) == false {
				return false
			}
			i = iterateDigit(f, i, len)
			if i >= len {
				return false
			}
			i = iterateBlank(f, i, len)
			if i >= len {
				return false
			}
		}
		if isComment(f[i]) == true {
			i = iterateComment(f, i, len)
		}
		if isNewLine(f[i]) == false {
			return false
		}
		i++
	}
	// EOF
	if i != len {
		return false
	}
	return true
}
