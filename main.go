package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

// Student model
type Student struct {
	Name   string
	Scores []int
}

var students []Student

func main() {
	options()
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func pressToContinue() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Press enter to continue")
	scanner.Scan()
}

func options() {
	var input string
	for input != "5" {
		clear()
		input = "0"
		fmt.Println("1) Capturar alumnos")
		fmt.Println("2) Guardar información en txt")
		fmt.Println("3) Abrir archivo txt")
		fmt.Println("4) Salir")
		fmt.Print("Select an option: ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			clear()
			captureStudents()
		case "2":
			clear()
			writeFile()
		case "3":
			clear()
			readFile()
			pressToContinue()
		case "4":
		}
	}
}

func captureStudents() {
	i := 0
	j := 0
	for len(students) < 5 {
		var student Student
		fmt.Println("Ingrese nombre del alumno")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			if str := scanner.Text(); len(str) != 0 {
				student.Name = str
				i++
				var scores []int
				for len(scores) < 3 {
					fmt.Println("Ingrese calificación " + strconv.Itoa(j+1) + " del alumno " + strconv.Itoa(i))
					scanner2 := bufio.NewScanner(os.Stdin)
					if scanner2.Scan() {
						if str := scanner2.Text(); len(str) != 0 {
							val, err := strconv.Atoi(str)
							if err != nil {
								fmt.Println("Ingrese un número entero")
							} else {
								scores = append(scores, val)
								j++
							}
						}
					}
				}
				j = 0
				student.Scores = scores
				students = append(students, student)
			}
		}
	}
	printStudents(students)
}

func writeFile() {
	if len(students) > 0 {
		file, err := os.Create("test.txt") // Truncates if file already exists, be careful!
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close() // Make sure to close the file when you're done
		spaces := "              "
		for i := 0; i < len(students); i++ {
			if i == 0 {
				_, err := file.WriteString("Nombre         Calificación 1  Calificación 2  Calificación 3")
				if err != nil {
					log.Fatalf("failed writing to file: %s", err)
				}
			}
			student := students[i]
			spacesName := ""
			for h := 0; h < 15-len(student.Name); h++ {
				spacesName += " "
			}
			_, err := file.WriteString("\n" + student.Name + spacesName + "       " + strconv.Itoa(student.Scores[0]) + spaces + strconv.Itoa(student.Scores[1]) + spaces + strconv.Itoa(student.Scores[2]))
			if err != nil {
				log.Fatalf("failed writing to file: %s", err)
			}
		}

		fmt.Println()
	} else {
		fmt.Println("No hay datos para escribir")
		pressToContinue()
	}
}

func readFile() {
	data, err := ioutil.ReadFile("test.txt")
	if err == nil {
		fmt.Printf("\n%s", data)
		fmt.Println()
	} else {
		fmt.Println(err)
		// log.Panicf("failed reading data from file: %s", err)
	}

}

func printStudents(students []Student) {
	if len(students) != 0 {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Println("STUDENTS")
		fmt.Fprintln(w, "\t  Nombre\t  Calificación 1\t  Calificación 2\t  Calificación 3\t")
		for _, student := range students {
			fmt.Fprintln(w, ("\t  " + student.Name + "\t  " + strconv.Itoa(student.Scores[0]) + "\t  " + strconv.Itoa(student.Scores[1]) + "\t  " + strconv.Itoa(student.Scores[2]) + "\t  "))
		}
		w.Flush()
	} else {
		fmt.Println("No records found")
	}
}
