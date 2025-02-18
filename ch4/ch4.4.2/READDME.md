### 4.4.2. Сравнение структур

Если все поля структуры сравниваемы, то сравниваема и сама структура, так что два выражения этого типа можно сравнивать
с помощью оператора == или != .
Оператор == поочерёдно сравнивает соответствующие поля двух структур, 
так что приведённые ниже выражения сравнивания эквивалентны:
```go
type Point struct{X, Y int}
p := Point{1,2}
q := Point{2,1}
fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
fmt.Println(p == q)                   // "false"
```

---

Сравниваемые структурные типы, как и любые сравниваемые типы, могут использоваться в качестве типа ключа в отображении:
```go
type address struct {
	hostname string
	port     int
}
hits := make(map[address]int)
hits[address{"golang.org", 443}]++
```