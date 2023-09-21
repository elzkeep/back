package color

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

type Color int

const (
	None Color = iota
	River
	Red
	Yellow
	Brown
	Black
	Blue
	Green
	Gray
)

func (p Color) ToString() string {
	colors := []string{"", "River", "Red", "Yellow", "Brown", "Black", "Blue", "Green", "Gray"}

	pos := int(p)
	ret := "      "
	if pos > 0 {
		ret = colors[pos]
	}

	return fmt.Sprintf("%-6s", ret)
}

func (p Color) ToShortString() string {
	colors := []string{"", "Ri", "Re", "Ye", "Br", "Bl", "Bu", "Gr", "Gy"}

	pos := int(p)
	str := colors[pos]

	if p == Red {
		return fmt.Sprintf("%v", aurora.BgRed(aurora.Black(str)))
	} else if p == Yellow {
		return fmt.Sprintf("%v", aurora.BgYellow(aurora.Black(str)))
	} else if p == Brown {
		return fmt.Sprintf("%v", aurora.BgBrightMagenta(aurora.Black(str)))
	} else if p == Black {
		return fmt.Sprintf("%v", aurora.BgBrightBlack(aurora.Black(str)))
	} else if p == Blue {
		return fmt.Sprintf("%v", aurora.BgBlue(aurora.Black(str)))
	} else if p == Green {
		return fmt.Sprintf("%v", aurora.BgGreen(aurora.Black(str)))
	} else if p == Gray {
		return fmt.Sprintf("%v", aurora.White(aurora.Black(str)))
	} else {
		return str
	}
}

func (p Color) ToStringBackground(value string) string {
	str := fmt.Sprintf("%-6s", value)

	if p == Red {
		return fmt.Sprintf("%v", aurora.BgRed(aurora.Black(str)))
	} else if p == Yellow {
		return fmt.Sprintf("%v", aurora.BgYellow(aurora.Black(str)))
	} else if p == Brown {
		return fmt.Sprintf("%v", aurora.BgBrightMagenta(aurora.Black(str)))
	} else if p == Black {
		return fmt.Sprintf("%v", aurora.BgBrightCyan(aurora.Black(str)))
	} else if p == Blue {
		return fmt.Sprintf("%v", aurora.BgBlue(aurora.Black(str)))
	} else if p == Green {
		return fmt.Sprintf("%v", aurora.BgGreen(aurora.Black(str)))
	} else if p == Gray {
		return fmt.Sprintf("%v", aurora.BgWhite(aurora.Black(str)))
	} else {
		return str
	}
}
