package main

import "fmt"

// интерфейсы
type jsonService interface {
	SendJson() string
}

type xmlService interface {
	SendXml() string
}

// json-документ
type jsonDocument struct{}

func (j jsonDocument) SendJson() string {
	return "`json:`"
}

// xml-документ
type xmlDocument struct{}

func (x xmlDocument) SendXml() string {
	return "<xml>testXml</xml>"
}

// adapter делает jsonService совместимым с xmlService
type adapter struct {
	json jsonService
}

func (a adapter) SendXml() string {
	// оборачиваем JSON в XML
	return "<xml>" + a.json.SendJson() + "</xml>"
}

func main() {
	exampleJson := jsonDocument{}

	// работаем через xmlService, хотя внутри JSON
	var xmlDoc xmlService = adapter{json: exampleJson}

	fmt.Println(xmlDoc.SendXml())
}

/*Пояснение (по паттерну «Адаптер»)

Применимость:

Когда у нас есть библиотека/код, который ожидает один интерфейс, но мы не можем/не хотим переписывать существующую структуру.

Тогда мы «оборачиваем» объект адаптером, который конвертирует вызовы.

Плюсы:

Позволяет использовать уже существующий код, не меняя его.

Упрощает интеграцию несовместимых систем.

Работает как «прослойка», поэтому удобно расширять.

Минусы:

Добавляется слой абстракции → больше кода, может снижать производительность.

Иногда адаптеров становится слишком много (если много несовместимых интерфейсов).*/
