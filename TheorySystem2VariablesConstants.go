package main

import (
	"fmt"
	"reflect"
	"time"
)

func TS2() {
	var s = "щ"
	fmt.Println(reflect.TypeOf(s))
	var now = time.Now()
	fmt.Println(now)
	pi, e := 3.1415, 2.7183 //множественное обьявление
	fmt.Println(pi, e)

	//Константы
	const (
		id  = 100
		id2 = 200
	) //Группировка (как и переменных)
	var i int64 = id
	var f float64 = id
	fmt.Println("i=", i, "f=", f) //Ошибки нет потому что типизация не такая сильная для констант

	//Инкремент iota
	const (
		Black = iota
		Gray
		White
	)
	// счётчик обнуляется
	const (
		Yellow = iota
		Red
		Green = iota // это присваивание не обнулит iota
		Blue
	)
	const (
		Purple   = 0
		darkBlue = iota //может быть не в начале
	)
	//Лайфхак
	const (
		_ = iota * 10 //Можно пропускать константы
		ten
		twenty
		thirty
	)
	//Иногда правильнее определить пользовательский тип для констант чтобы не запутаться:
	type Weekday int

	const (
		Monday Weekday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
}
