package resources

type Building int

const (
	None Building = -1
	D    Building = iota - 1
	TP
	TE
	SH
	SA
	WHITE_D
	WHITE_TP
	WHITE_TE
	WHITE_SH
	WHITE_SA
	WHITE_MT
)

func (p Building) ToString() string {
	names := []string{"D", "TP", "TE", "SH", "SA", "WHITE_D", "WHITE_TP", "WHITE_TE", "WHITE_SH", "WHITE_SA", "WHITE_MT"}

	pos := int(p)

	if pos == -1 {
		return ""
	}

	return names[pos]
}

func (p Building) Power() int {
	powers := []int{1, 2, 2, 3, 3, 1, 2, 2, 3, 3, 4}
	pos := int(p)

	return powers[pos]
}
