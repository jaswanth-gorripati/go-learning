package goInterface

type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

type IntCollection struct {
	Elements []int
}

func (ic *IntCollection) Len() int {
	return len(ic.Elements)
}

func (ic *IntCollection) Less(i, j int) bool {
	return ic.Elements[i] < ic.Elements[j]
}

func (ic *IntCollection) Swap(i, j int) {
	ic.Elements[i], ic.Elements[j] = ic.Elements[j], ic.Elements[i]
}
