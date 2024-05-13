package bvg_tickets

// Thanks to https://www.reddit.com/r/golang/comments/uvpygm/comment/iab61oh/
type Zone int

const (
	ZoneABC Zone = iota
	ZoneAB
	ZoneBC
)

func (z Zone) String() string {
	return []string{"ABC", "AB", "BC"}[z]
}
