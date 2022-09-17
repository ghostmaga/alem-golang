package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	wtref := &List{}
	for l.Head != nil {
		if l.Head.Data == data_ref {
			l.Head = l.Head.Next
			continue
		}
		ListPushBack(wtref, l.Head.Data)
		l.Head = l.Head.Next
	}
	*l = *wtref
}
