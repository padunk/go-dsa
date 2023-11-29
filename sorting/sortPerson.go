package sorting

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int               { return len(a) }
func (a ByAge) Swap(i int, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i int, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name > a[j].Name }
