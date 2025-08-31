package main

import (
	"fmt"
	"os"
)

// Interface держит в себе три поля: data (сами данные), itab(тип + указатель на таблицу методов)
// Любая структура, которая имеет такие же методы, что и интерфейс, может стать им (утиная типизация)
// Как раз interface{} поддается любому типу данных, так как нет определенных методов, которых тип должен иметь
// С версии 1.18 появился алиас пустого интерфейса - any

func Foo() error {
	var err *os.PathError = nil
	// type - *os.PathError
	// data - nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        //Вывод: <nil>
	fmt.Println(err == nil) // Вывод: false
	// выводится из поля data, а там хранится nil, но сама переменная не является nil
}
