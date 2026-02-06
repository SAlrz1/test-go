package main

import "fmt"

type Formatter interface {
	Format() string
}
type Person struct {
	Name string
	Age  int
}
type Book struct {
	Title string
	Price float64
	Person
}

func (p *Person) Format() string {
	return fmt.Sprintf("%s is %d years old ", p.Name, p.Age)
}
func (b *Book) Format() string {
	return fmt.Sprintf("%s is %.2f %v", b.Title, b.Price, b.Person.Format())
}
func BatchFormat[T Formatter](items []T) []string {
	result := make([]string, 0, len(items))
	for _, v := range items {
		result = append(result, v.Format()) //批处理
	}
	return result
}
func main() {
	person := []*Person{
		{"liruzhen", 22},
		{"wanglonghui", 22},
	}
	book := []*Book{
		{"hdlwl", 20.3, Person{"liruzhen", 22}},
		{"gtszylcd", 50, Person{"wanglonghui", 22}},
	}
	fmt.Println(BatchFormat(book))
	fmt.Println(BatchFormat(person))
}
