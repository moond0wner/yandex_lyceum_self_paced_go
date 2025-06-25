package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getStudentName(name string, occupied bool) string {
	if occupied {
		return name
	}
	return "-"
}

func main() {
	var student1, student2, student3, student4, student5 string
	var occupied1, occupied2, occupied3, occupied4, occupied5 bool

	numberOfStudents := 5
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		parts := strings.Fields(input)

		if len(parts) == 1 {
			command := parts[0]
			switch command {
			case "конец":
				fmt.Printf("1. %s\n", getStudentName(student1, occupied1))
				fmt.Printf("2. %s\n", getStudentName(student2, occupied2))
				fmt.Printf("3. %s\n", getStudentName(student3, occupied3))
				fmt.Printf("4. %s\n", getStudentName(student4, occupied4))
				fmt.Printf("5. %s\n", getStudentName(student5, occupied5))
				return
			case "очередь":
				fmt.Printf("1. %s\n", getStudentName(student1, occupied1))
				fmt.Printf("2. %s\n", getStudentName(student2, occupied2))
				fmt.Printf("3. %s\n", getStudentName(student3, occupied3))
				fmt.Printf("4. %s\n", getStudentName(student4, occupied4))
				fmt.Printf("5. %s\n", getStudentName(student5, occupied5))
			case "количество":
				occupiedCount := 0
				if occupied1 {
					occupiedCount++
				}
				if occupied2 {
					occupiedCount++
				}
				if occupied3 {
					occupiedCount++
				}
				if occupied4 {
					occupiedCount++
				}
				if occupied5 {
					occupiedCount++
				}

				freeCount := numberOfStudents - occupiedCount
				fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", freeCount, occupiedCount)

			default:
				fmt.Println("некорректный ввод")
			}
		} else if len(parts) == 2 {
			name := parts[0]
			numberStr := parts[1]

			number, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", numberStr)
				continue
			}

			if number < 1 || number > numberOfStudents {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", number)
				continue
			}

			occupiedCount := 0
			if occupied1 {
				occupiedCount++
			}
			if occupied2 {
				occupiedCount++
			}
			if occupied3 {
				occupiedCount++
			}
			if occupied4 {
				occupiedCount++
			}
			if occupied5 {
				occupiedCount++
			}

			switch number {
			case 1:
				if occupiedCount == numberOfStudents {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
					continue
				}
				if occupied1 {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				student1 = name
				occupied1 = true
			case 2:
				if occupiedCount == numberOfStudents {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
					continue
				}
				if occupied2 {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				student2 = name
				occupied2 = true
			case 3:
				if occupiedCount == numberOfStudents {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
					continue
				}
				if occupied3 {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				student3 = name
				occupied3 = true
			case 4:
				if occupiedCount == numberOfStudents {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
					continue
				}
				if occupied4 {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				student4 = name
				occupied4 = true
			case 5:
				if occupiedCount == numberOfStudents {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", number)
					continue
				}
				if occupied5 {
					fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", number)
					continue
				}
				student5 = name
				occupied5 = true
			}

		}
	}
}
