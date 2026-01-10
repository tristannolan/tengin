package tengin

import "strings"

func Box(width, height int, bg Color) *Canvas {
	c := NewCanvas(width, height)
	for y := range c.Tiles {
		for x := range c.Tiles[y] {
			tile := NewTile("", NewStyle().SetBg(bg))
			c.SetTile(x, y, tile)
		}
	}
	return c
}

func Text(str string) *Canvas {
	c := NewCanvas(len(str), 1)
	i := 0
	for char := range strings.SplitSeq(str, "") {
		tile := NewTile(char, NewStyle())
		c.SetTile(i, 0, tile)
		i++
	}
	return c
}

func Paragraph(width int, str string) *Canvas {
	var lines []string

	for p := range strings.SplitSeq(str, "\n") {
		// Preserve blank lines
		if len(p) == 0 {
			lines = append(lines, "")
			continue
		}

		lastIndex := 0

		for {
			if lastIndex+width >= len(p) {
				lines = append(lines,
					strings.TrimSpace(string(p[lastIndex:])),
				)
				break
			}

			i := lastIndex + width

			// Go back to last space
			for i > lastIndex && p[i] != ' ' {
				i--
			}

			// No space found, force a wrap
			if i == lastIndex {
				i += width
			}

			lines = append(lines,
				strings.TrimSpace(string(p[lastIndex:i])),
			)

			if i < len(p) && p[i] == ' ' {
				i++
			}
			lastIndex = i
		}
	}

	c := NewCanvas(width, len(lines))
	for i, line := range lines {
		chars := strings.Split(line, "")
		for j, char := range chars {
			tile := NewTile(char, NewStyle())
			c.SetTile(j, i, tile)
		}
	}

	return c
}
