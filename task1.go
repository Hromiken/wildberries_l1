package wildberries_l1

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func main() {
	example := Human{"Kate", 12}
	action := Action{
		Human: example,
	}
	action.Walk()
}

func (h Human) Walk() {
	fmt.Printf("%s is walking on\n", h.Name)
}
