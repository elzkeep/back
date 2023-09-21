package color

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

type Color int

const (
	Empty Color = -1
	None  Color = -1
	River Color = iota - 2
	Red
	Yellow
	Brown
	Black
	Blue
	Green
	Gray
)

func (p Color) ToString() string {
	colors := []string{"River", "Red", "Yellow", "Brown", "Black", "Blue", "Green", "Gray"}

	pos := int(p)
	ret := "      "
	if pos > 0 {
		ret = colors[pos]
	}

	return fmt.Sprintf("%-6s", ret)
}

func (p Color) ToShortString() string {
	colors := []string{"Ri", "Re", "Ye", "Br", "Bl", "Bu", "Gr", "Gy"}

	pos := int(p)
	str := colors[pos]

	if pos == 1 {
		return fmt.Sprintf("%v", aurora.BgRed(aurora.Black(str)))
	} else if pos == 2 {
		return fmt.Sprintf("%v", aurora.BgYellow(aurora.Black(str)))
	} else if pos == 3 {
		return fmt.Sprintf("%v", aurora.BgBrightMagenta(aurora.Black(str)))
	} else if pos == 4 {
		return fmt.Sprintf("%v", aurora.BgBrightBlack(aurora.Black(str)))
	} else if pos == 5 {
		return fmt.Sprintf("%v", aurora.BgBlue(aurora.Black(str)))
	} else if pos == 6 {
		return fmt.Sprintf("%v", aurora.BgGreen(aurora.Black(str)))
	} else if pos == 7 {
		return fmt.Sprintf("%v", aurora.White(aurora.Black(str)))
	} else {
		return str
	}
}

func (p Color) ToStringBackground(value string) string {
	pos := int(p)

	str := fmt.Sprintf("%-6s", value)

	if pos == 1 {
		return fmt.Sprintf("%v", aurora.BgRed(aurora.Black(str)))
	} else if pos == 2 {
		return fmt.Sprintf("%v", aurora.BgYellow(aurora.Black(str)))
	} else if pos == 3 {
		return fmt.Sprintf("%v", aurora.BgBrightMagenta(aurora.Black(str)))
	} else if pos == 4 {
		return fmt.Sprintf("%v", aurora.BgBrightCyan(aurora.Black(str)))
	} else if pos == 5 {
		return fmt.Sprintf("%v", aurora.BgBlue(aurora.Black(str)))
	} else if pos == 6 {
		return fmt.Sprintf("%v", aurora.BgGreen(aurora.Black(str)))
	} else if pos == 7 {
		return fmt.Sprintf("%v", aurora.BgWhite(aurora.Black(str)))
	} else {
		return str
	}
}
