package main

type meter interface {
	val() int
	inc()
	dec()
}

type tracker struct {
	value int
}

func (t *tracker) inc() {
	t.value = t.value + 1
}

func (t *tracker) dec() {
	t.value = t.value - 1
}

func (t *tracker) val() int {
	return t.value
}

func NewMeter() meter {
	return &tracker{}
}
func main() {

}
