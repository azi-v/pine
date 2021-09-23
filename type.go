package pine

import "strconv"

type Field interface {
	Input() *Str
}

type Str struct {
	Key string
	Val string
}

func (o Str) Input() *Str {
	return &o
}

type Int struct {
	Key string
	Val int
}

func (o Int) Input() *Str {
	return &Str{
		Key: o.Key,
		Val: strconv.FormatInt(int64(o.Val), 10),
	}
}
