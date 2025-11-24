package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World")

	var a int
	b := a         //короткая запись без явного указания типа
	fmt.Println(b) //как видно a было присвоено 0 по умолчанию

	var str string = "Hello World2"
	fmt.Println(str)

	var bl bool
	fmt.Println(bl) //false по умолчанию (а у String "")
	bl = true
	fmt.Println(bl)

	s := "абц"
	fmt.Println(len(s))                    //выдает 6))
	fmt.Println(s[0])                      //выдает байтовое значение первого символа
	fmt.Println(utf8.RuneCountInString(s)) //корректное кол-во символов с помощью специальной команды

	var stringFormattedVar string
	// следующие выражения равнозначны
	stringFormattedVar = "Hello,\nworld!\n\n\t\t\"quote!\""
	fmt.Println(stringFormattedVar)
	stringFormattedVar = `Hello,
world!

        "quote!"`
	fmt.Println(stringFormattedVar)

	//Пользовательские типы
	type Fruit string
	var fruit Fruit
	fmt.Println(fruit)

	var strr string
	strr = "Hello, world!"
	fmt.Println(strr[0]) //72, а нужна сама буква
	fmt.Println(string((strr[0])))

	//алиасы
	type MyString = string // MyString здесь — это псевдоним типа string
	var aaa string = "brr"
	var aab MyString
	aab = aaa // ошибки нет
	fmt.Println(aaa, aab)

}
