package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

//for check functionality:
//go run . --type json
//go run . --type xml
func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "type",
			Usage: "json or xml",
		},
	}

	app.Action = func(c *cli.Context) error {
		Handle(c.String("type"))
		// switch c.String("type") {
		// case "xml": // if its the 'xml' command
		//вынести в одну функцию

		// case "json": // if its the 'json' command
		// Handle()
		// default:
		// return fmt.Errorf("unexpected value! expected 'xml' or 'json' commands")
		// }
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//Вопросы/проблемы
//Флаги: если использовать библиотеку urfave/cli, то в ней флаги указываются вторым параметром, после Name.
//То есть нужно создать основной параметр (как у меня: --type), а уже после него задавать значение флага: json/xml
//То есть сделать полностью в соответствии с заданием: go run main.go -json НЕ получилось
//плюс даже не получилось установить флаг: -json / -xml (получилось просто json / xml)
//так как блиотека urfave/cli имеет другой функционал для параметров заданных через -

//мы думали, что json.Unmarshal проверяет наш json и можно больше не добавлять проверок
//но когда я попробовал добавить просто числовое значение, например: 123456789 то Unmarshal не ругается (проверка прошла без ошибок)
//но ведь это не json :)

//я решил создать отдельный файл file.txt, в который записываю хэш сумму (если это новая/уникальная хэш сумма)
//для простоты использования и чтения данных с этого файла - я записываю каждые данные в новую строку (это просто удобно)
//но вот красиво/правильно реализовать чтение с этого файла, проверку уникальности этих данных и запись новых данных мне не удалось
//написал через 'попу': пришлось два раза открывать файл + для добавления данных с новой строки ТУПО канкатонировал строку с /n
//должен быть способ это сделать красиво, но знаний не хватает

//workshop мы не прошли полностью. Дошли до 7 части включительно. Сам сделал 8 часть и делаю 9. По 9 есть вопросы
//также есть вопросы по тестам:
//1) пройтись по тестовому файлу и разобраться
//2) какие еще тесты нужно написать для первого воркшопа?
//Вопрос не связанный с тестами: Ваня сказал, что я не правильно создал каталоги проектов: go/src/мои_проекты
//Нужно создать каталог: github.com. То есть должно быть: go/src/github.com/мои_проекты ? Или как?
//Почему нужно именно использовать github.com? Для чего?

func Handle(flag string) {
	v := NewValidator(flag)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter json: ")
	scanner.Scan()
	jsonText := scanner.Text()

	err := v.Validate(jsonText)
	// err := ValidateJson(jsonText)
	if err != nil {
		log.Fatal(err)
	}

	err = Save(flag, jsonText)
	if err != nil {
		log.Fatal(err)
	}
}
