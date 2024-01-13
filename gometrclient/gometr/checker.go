package gometr

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Checker struct {
	items []Checkable
}

type Data struct {
	DataPing   bool
	IDs        string
	DataHealth bool
}

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

func (c *Checker) Add(elements ...Checkable) {
	c.items = append(c.items, elements...)
}

func (c *Checker) GetAllIDs() {
	for i := range c.items {
		fmt.Println(i)
	}
}

func (c *Checker) String() string {
	var ids []string
	for _, item := range c.items {
		ids = append(ids, item.GetID())
	}
	return fmt.Sprintf("Checker with IDs: %s", strings.Join(ids, ", "))
}

func (d *Data) GetMetrics() string {
	// URL, по которому вы хотите сделать GET-запрос
	url := "http://127.0.0.1:8000/metrics"

	// Отправка GET-запроса
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprint("Ошибка при отправке запроса:", err)
	}
	defer response.Body.Close()

	// Чтение тела ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprint("Ошибка при чтении ответа:", err)

	}

	// Преобразование тела ответа в строку
	content := string(body)

	// Использование регулярного выражения для поиска метрики go_memstats_frees_total
	re := regexp.MustCompile(`go_memstats_frees_total (\d+)`)
	matches := re.FindStringSubmatch(content)

	// Вывод найденной метрики
	if len(matches) >= 2 {
		return fmt.Sprintf("Метрика go_memstats_frees_total: %s\n", matches[1])
	} else {
		return fmt.Sprintf("Метрика go_memstats_frees_total не найдена.")
	}
}

func (d *Data) Ping() error {
	return nil
}

func (d *Data) GetID() string {
	var userInput string

	fmt.Print("Введите ID: ")
	fmt.Scanln(&userInput)

	return strings.TrimSpace(userInput)
}

func (d *Data) Health() bool {
	return true
}
