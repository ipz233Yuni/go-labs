package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ---------------------- STRUCTS ----------------------

type Worker struct {
	Name      string
	Year      int
	Month     int
	WorkPlace Company
}

type Company struct {
	Name     string
	Position string
	Salary   float64
}

// ---------------------- CONSTRUCTORS ----------------------

func NewCompany(name, position string, salary float64) Company {
	return Company{name, position, salary}
}

func NewWorker(name string, year, month int, company Company) Worker {
	return Worker{name, year, month, company}
}

// ---------------------- GET/SET METHODS ----------------------

func (w *Worker) GetName() string       { return w.Name }
func (w *Worker) GetYear() int          { return w.Year }
func (w *Worker) GetMonth() int         { return w.Month }
func (w *Worker) GetWorkPlace() Company { return w.WorkPlace }

func (w *Worker) SetName(name string)          { w.Name = name }
func (w *Worker) SetYear(year int)             { w.Year = year }
func (w *Worker) SetMonth(month int)           { w.Month = month }
func (w *Worker) SetWorkPlace(company Company) { w.WorkPlace = company }

func (c *Company) GetCompanyName() string { return c.Name }
func (c *Company) GetPosition() string    { return c.Position }
func (c *Company) GetSalary() float64     { return c.Salary }

func (c *Company) SetCompanyName(name string) { c.Name = name }
func (c *Company) SetPosition(pos string)     { c.Position = pos }
func (c *Company) SetSalary(s float64)        { c.Salary = s }

// ---------------------- COMPANY METHODS ----------------------

func (c Company) GetInfo() string {
	return fmt.Sprintf("Компанія: %s, Посада: %s, Оклад: %.2f", c.Name, c.Position, c.Salary)
}

// ---------------------- WORKER METHODS ----------------------

func (w Worker) GetWorkerPosition() string {
	return w.WorkPlace.GetPosition()
}

func (w Worker) GetWorkExperience() int {
	start := time.Date(w.Year, time.Month(w.Month), 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	months := (now.Year()-start.Year())*12 + int(now.Month()) - int(start.Month())
	if months < 0 {
		months = 0
	}
	return months
}

func (w Worker) GetTotalMoney() float64 {
	return float64(w.GetWorkExperience()) * w.WorkPlace.GetSalary()
}

// ---------------------- INPUT HELPERS ----------------------

var reader = bufio.NewReader(os.Stdin)

func ReadString(prompt string) string {
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text != "" {
			return text
		}
		fmt.Println("Поле не може бути пустим. Спробуйте ще раз.")
	}
}

func ReadInt(prompt string, min, max int) int {
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.Atoi(text)
		if err == nil && num >= min && num <= max {
			return num
		}
		fmt.Println("Некоректне значення. Спробуйте ще раз.")
	}
}

func ReadFloat(prompt string, min float64) float64 {
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		num, err := strconv.ParseFloat(text, 64)
		if err == nil && num >= min {
			return num
		}
		fmt.Println("Некоректне значення. Спробуйте ще раз.")
	}
}

// ---------------------- MAIN FUNCTIONS ----------------------

func ReadWorkersArray() []Worker {
	n := ReadInt("Введіть кількість працівників: ", 1, 1000)
	workers := make([]Worker, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\n--- Працівник #%d ---\n", i+1)

		name := ReadString("ПІБ працівника: ")
		year := ReadInt("Рік початку роботи: ", 1950, time.Now().Year())
		month := ReadInt("Місяць початку роботи (1-12): ", 1, 12)

		companyName := ReadString("Назва компанії: ")
		position := ReadString("Посада: ")
		salary := ReadFloat("Зарплата за місяць: ", 0)

		company := NewCompany(companyName, position, salary)
		workers[i] = NewWorker(name, year, month, company)
	}

	return workers
}

func PrintWorker(w Worker) {
	fmt.Println("-------------")
	fmt.Println("ПІБ:", w.Name)
	fmt.Println(w.WorkPlace.GetInfo())
	fmt.Println("Досвід роботи (міс.):", w.GetWorkExperience())
	fmt.Println("Загальна зарплата:", w.GetTotalMoney())
	fmt.Println("-------------")
}

func PrintWorkers(workers []Worker) {
	for _, w := range workers {
		PrintWorker(w)
	}
}

func GetWorkersInfo(workers []Worker) (float64, float64) {
	if len(workers) == 0 {
		return 0, 0
	}

	max := workers[0].WorkPlace.GetSalary()
	min := workers[0].WorkPlace.GetSalary()

	for _, w := range workers {
		s := w.WorkPlace.GetSalary()
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}
	return max, min
}

// ---------------------- MAIN ----------------------

func main() {
	workers := ReadWorkersArray()
	PrintWorkers(workers)

	max, min := GetWorkersInfo(workers)
	fmt.Printf("\nНайбільша зарплата: %.2f\n", max)
	fmt.Printf("Найменша зарплата: %.2f\n", min)
}
