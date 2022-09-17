package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	s := 0
	forsize := l
	for forsize != nil {
		forsize = forsize.Next
		s++
	}
	if pos > s {
		return nil
	} else if l == nil {
		return nil
	} else if l.Next == nil {
		return l
	}
	for pos != 0 {
		l = l.Next
		pos--
	}
	return l
}
