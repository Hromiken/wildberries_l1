package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1024)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(len(justString))
}

/*


1. К каким негативным последствиям ?
Если просто взять срез v[:100], как в исходном варианте, justString будет ссылаться на всю исходную строку v в памяти.
При больших строках или повторяющихся операциях это может привести к неэффективному использованию памяти и потенциальным утечкам.

2.Что происходит с переменной justString?

justString — это срез строки v, который ссылается на те же данные в памяти, что и v.
Хотя мы видим только 100 байт, память под всю строку всё ещё удерживается.

3. Как это исправить ?
Использовать копирование данных

*/
