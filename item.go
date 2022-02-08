package avl

type Item interface {
	Less(b Item) int
}

type Int int

func (a Int) Less(b Item) int {
	if a == b {
		return 0
	} else if a < b.(Int) {
		return -1
	} else {
		return 1
	}
}
