package resources

type Building int

const (
	None Building = iota
	D
	TP
	TE
	SH
	SA
	WHITE_D
	WHITE_TOWER
	WHITE_TP
	WHITE_TE
	WHITE_SH
	WHITE_SA
	WHITE_MT
	CITY
)

func (p Building) ToString() string {
	names := []string{"", "D", "TP", "TE", "SH", "SA", "WHITE_D", "WHITE_TOWER", "WHITE_TP", "WHITE_TE", "WHITE_SH", "WHITE_SA", "WHITE_MT", "CITY"}

	pos := int(p)

	if pos == -1 {
		return ""
	}

	return names[pos]
}

func (p Building) Power() int {
	powers := []int{0, 1, 2, 2, 3, 3, 1, 2, 2, 2, 3, 3, 4, 0}
	pos := int(p)

	return powers[pos]
}
