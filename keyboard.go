package telegram

func MakeKeyboard(rows int, buttons ...KeyboardButton) [][]KeyboardButton {
	cols := len(buttons) / rows

	if len(buttons)%rows > 0 {
		cols = cols + 1
	}

	keyboard := make([][]KeyboardButton, rows)
	for i := 0; i < rows; i++ {
		keyboard[i] = make([]KeyboardButton, cols)

		for j := 0; j < cols && i*cols+j < len(buttons); j++ {
			keyboard[i][j] = buttons[i*cols+j]
		}
	}

	return keyboard
}
