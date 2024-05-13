package zone

// Thanks to https://www.reddit.com/r/golang/comments/uvpygm/comment/iab61oh/
type Zone int

const (
	ABC Zone = iota
	AB
	BC
)

func (z Zone) String() string {
	return []string{"ABC", "AB", "BC"}[z]
}
