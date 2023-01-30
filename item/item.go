package item

import ( 
	"fmt"
)

type Item struct {
	Type string
	location string
	weight int  //gram
}

func (i *Item) Init(Type string, location string, weight int) {
	i.Type = Type
	i.location = location
	i.weight = weight
}

func (i *Item) SetDefault() {
	if (i.Type == "") {
		i.Type = "SomeItem"
	}
	if (i.location == "") {
		i.location = "Bag"
	}
	if (i.weight == 0) { 
		i.weight = 100
	}
}

func (i *Item) GetInfo()  {
	fmt.Printf("%+v\n", i)
}

func (i *Item) GetType() string {
	return i.Type;
}

func (i *Item) GetLocation() string {
	return i.location
}

func (i *Item) SetLocation(location string) {
	i.location = location
}
func (i *Item) GetWeight() int{
	return i.weight
}
func (i *Item) SetWeight(weight int) {
	i.weight = weight
}



// func main() {
// 	i1 := new(Item)
// 	i1.Init("phone", "bed", 200)
// 	i1.GetInfo()
// }
