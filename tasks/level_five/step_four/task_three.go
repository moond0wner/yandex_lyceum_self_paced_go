// "Кадровый учёт"

package main

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrInvalidPosition = errors.New("invalid worker position")
	ErrEmptyList       = errors.New("empty list workers")
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, income, experience uint) error
	SortWorkers() ([]string, error)
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Worker struct {
	name       string
	position   string
	salary     uint
	experience uint
	income     uint
}

type Company struct {
	Workers []Worker
}

func NewCompany() *Company {
	return &Company{
		Workers: []Worker{},
	}
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	switch position {
	case "директор", "зам. директора", "начальник цеха", "мастер", "рабочий":
		worker := Worker{
			name:       name,
			position:   position,
			salary:     salary,
			experience: experience,
			income:     salary * experience,
		}
		c.Workers = append(c.Workers, worker)
	default:
		return ErrInvalidPosition
	}
	return nil
}

func positionPriority(position string) int {
	switch position {
	case "директор":
		return 5
	case "зам. директора":
		return 4
	case "начальник цеха":
		return 3
	case "мастер":
		return 2
	case "рабочий":
		return 1
	default:
		return 0
	}
}

func (c *Company) SortWorkers() ([]string, error) {
	if len(c.Workers) == 0 {
		return nil, ErrEmptyList
	}
	sort.Slice(c.Workers, func(i, j int) bool {
		if c.Workers[i].income > c.Workers[j].income {
			return true
		}
		if c.Workers[i].income < c.Workers[j].income {
			return false
		}
		posPriorityI := positionPriority(c.Workers[i].position)
		posPriorityJ := positionPriority(c.Workers[j].position)
		if posPriorityI > posPriorityJ {
			return true
		} else if posPriorityI < posPriorityJ {
			return false
		}
		return c.Workers[i].name < c.Workers[j].name
	})
	result := []string{}
	for _, worker := range c.Workers {
		result = append(result, fmt.Sprintf("%s — %d — %s", worker.name, worker.income, worker.position))
	}
	return result, nil
}

func main() {
	company := NewCompany()

	// 1. Добавление сотрудников:
	err := company.AddWorkerInfo("Михаил", "директор", 1000, 12)
	if err != nil {
		fmt.Println("Ошибка при добавлении сотрудника:", err)
		return
	}

	err = company.AddWorkerInfo("Андрей", "мастер", 900, 13)
	if err != nil {
		fmt.Println("Ошибка при добавлении сотрудника:", err)
		return
	}

	err = company.AddWorkerInfo("Игорь", "зам. директора", 800, 14)
	if err != nil {
		fmt.Println("Ошибка при добавлении сотрудника:", err)
		return
	}

	// 2. Обработка ошибки при добавлении сотрудника с недопустимой должностью:
	err = company.AddWorkerInfo("Сергей", "менеджер", 600, 10)
	if err != nil {
		fmt.Println("Ошибка при добавлении сотрудника:", err) // Вывод: Ошибка при добавлении сотрудника: invalid worker position
	}

	// 4. Сортировка сотрудников:
	sortedWorkers, err := company.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка при сортировке сотрудников:", err)
		return
	}

	// 5. Вывод информации о сотрудниках после сортировки:
	fmt.Println("\nСотрудники после сортировки:")
	for _, workerInfo := range sortedWorkers {
		fmt.Println(workerInfo)
	}

	// 6. Пример с добавлением нескольких сотрудников:
	company.AddWorkerInfo("Елена", "рабочий", 400, 15)
	company.AddWorkerInfo("Дмитрий", "начальник цеха", 700, 11)

	// 7. Повторная сортировка и вывод:
	sortedWorkers, err = company.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка при повторной сортировке сотрудников:", err)
		return
	}

	fmt.Println("\nСотрудники после повторной сортировки:")
	for _, workerInfo := range sortedWorkers {
		fmt.Println(workerInfo)
	}

	// 8. Проверка на пустой список сотрудников:
	emptyCompany := NewCompany()
	sortedWorkers, err = emptyCompany.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка при сортировке в пустой компании:", err)
		return
	}

	fmt.Println("\nСотрудники в пустой компании:")
	for _, workerInfo := range sortedWorkers {
		fmt.Println(workerInfo) // Ничего не выводится
	}

	// 9. Добавление одинаковых должностей
	err = company.AddWorkerInfo("Ольга", "директор", 1200, 12)
	if err != nil {
		fmt.Println("Ошибка при добавлении сотрудника:", err)
		return
	}

	sortedWorkers, err = company.SortWorkers()
	if err != nil {
		fmt.Println("Ошибка при повторной сортировке сотрудников:", err)
		return
	}

	fmt.Println("\nСотрудники после добавления директора:")
	for _, workerInfo := range sortedWorkers {
		fmt.Println(workerInfo)
	}
}
