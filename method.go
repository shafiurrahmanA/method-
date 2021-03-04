package main
import "fmt"


type person struct {
  name   string
  age    int
  gender string
}


func (p *person) describe() {
  fmt.Printf("%v is %v years old.", p.name, p.age)
}
func (p *person) setAge(age int) {
  p.age = age
}

func (p person) setName(name string) {
  p.name = name
}

func main() {
  pp := &person{name: "rahman", age: 25, gender: "Male"}
  pp.describe()
  
  pp.setAge(25)
  fmt.Println(pp.age)
  
  pp.setName("Hari")
  fmt.Println(pp.name)
  
}