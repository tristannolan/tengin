package tengin

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type Pattern struct {
	Canvas        *Canvas
	Base          [][]string
	Args          map[string][]Position
	Styles        [][]string
	ArgsPhrases   map[string]string
	stylesPhrases map[string]string
	width         int
	height        int
}

func LoadPattern(path string, defaultStyle *Style, styles map[string]*Style) (Pattern, error) {
	p := Pattern{
		Args:        map[string][]Position{},
		ArgsPhrases: map[string]string{},
	}

	const (
		headerPhase uint8 = iota
		argsCollectionPhase
		basePhase
		argsPhase
		stylePhase
	)

	phase := headerPhase

	widthMarker := "-"
	argsMarker := "="

	lineNum := 0
	argsLineNum := 0
	styleLineNum := 0

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

Scan:
	for scanner.Scan() {
		line := scanner.Text()

		lineNum++

		if len(line) == 0 {
			continue
		}

		trimmed := strings.TrimSpace(line)
		switch trimmed {
		case "BASE":
			phase = basePhase
			continue
		case "ARGS":
			phase = argsPhase
			continue
		case "STYLE":
			phase = stylePhase
			p.Styles = make([][]string, len(p.Base))
			for i := range p.Base {
				p.Styles[i] = make([]string, len(p.Base[i]))
			}
			continue
		case "END":
			break Scan
		}

		switch phase {
		case headerPhase:
			p.width = getPatternWidth(line, widthMarker)
			if p.width == 0 {
				return p, errors.New("Pattern width not defined. Check header.")
			}

			phase = argsCollectionPhase

		case argsCollectionPhase:
			argSplit := strings.Split(trimmed, argsMarker)
			if len(argSplit) != 2 {
				continue
			}

			key := argSplit[0]
			value := argSplit[1]

			if len(key) == 1 && len(value) > 0 {
				p.ArgsPhrases[key] = value
			}

		case basePhase:
			chars := strings.Split(line, "")
			baseChars := make([]string, p.width)

			for x := range baseChars {
				if x > len(chars)-1 {
					break
				}

				char := chars[x]
				if char == " " {
					char = ""
				}
				baseChars[x] = char
			}

			p.Base = append(p.Base, baseChars)

		case argsPhase:
			if argsLineNum >= len(p.Base) {
				break
			}

			chars := strings.Split(line, "")

			for x := range p.Base[argsLineNum] {
				if x > len(chars)-1 {
					break
				}

				argsChar := chars[x]
				baseChar := p.Base[argsLineNum][x]

				if argsChar == baseChar || argsChar == "" {
					continue
				}

				if _, ok := p.ArgsPhrases[argsChar]; !ok {
					continue
				}

				p.Args[argsChar] = append(p.Args[argsChar], NewPosition(x, argsLineNum))
			}

			argsLineNum++

		case stylePhase:
			if styleLineNum >= len(p.Base) {
				break
			}

			styleChars := strings.Split(line, "")

			for x := range p.Styles[styleLineNum] {
				if x > len(styleChars)-1 {
					break
				}

				styleChar := styleChars[x]
				baseChar := p.Base[styleLineNum][x]

				if styleChar == baseChar {
					styleChar = ""
				}

				p.Styles[styleLineNum][x] = styleChar
			}

			styleLineNum++

		default:
			return p, errors.New("Overran phase while loading pattern.")
		}
	}

	p.Canvas = NewCanvas(p.width, len(p.Base[0]))

	for y := range p.Base {
		for x, char := range p.Base[y] {
			style := defaultStyle
			if p.Styles[y][x] != "" {
				if _, ok := styles[p.Styles[y][x]]; ok {
					style = styles[p.Styles[y][x]]
				}
			}

			p.Canvas.SetTile(x, y, NewTile(char, style))
		}
	}

	return p, nil
}

func getPatternWidth(line, widthMarker string) int {
	width := 0
	for char := range strings.SplitSeq(line, "") {
		if char != widthMarker {
			break
		}
		width++
	}
	return width
}
