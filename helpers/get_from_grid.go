package helpers

func GetFromGrid(x int, y int, lines []string) rune {
	if y >= 0 && x >= 0 && y < len(lines) && x < len(lines[y]) {
		return rune(lines[y][x])
	}

	return 0
}
