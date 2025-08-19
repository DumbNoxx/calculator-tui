package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	ErrorNotDivisible = errors.New("No puedes dividir entre 0")
	ErrorIsZero       = errors.New("No puedes hacer operaciones si los dos numeros son 0")
)

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrorNotDivisible
	}
	return a / b, nil
}

func Sumar(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrorIsZero
	}
	return a + b, nil
}

func Restar(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, ErrorIsZero
	}
	return a - b, nil
}

func Multiplicar(a, b float64) (float64, error) {
	if a == 0 && b == 0 {
		return 0, ErrorIsZero
	}
	return a * b, nil
}

func ErrorWrapper(err error, typeError string) error {
	return fmt.Errorf("Ha ocurrido un error con la %s, contexto: %w", typeError, err)
}

func UnwrapperError(err error) error {
	return errors.Unwrap(err)
}

func continues(option string) {
	for {
		fmt.Print("Deseas continuar? [y/yes/n/no]: ")
		fmt.Scanln(&option)
		switch option {
		case "y", "yes":
			return
		case "n", "no":
			os.Exit(0)
		default:
			fmt.Println("Escoge alguna de las opciones")
			continue
		}
	}
}

func Menu() {
	var option int
	fmt.Println("Bienvenido a la calculadora")
	fmt.Println("1. Sumar")
	fmt.Println("2. Restar")
	fmt.Println("3. Multiplicar")
	fmt.Println("4. Dividir")
	fmt.Println("5. Salir")
	fmt.Print("Opcion: ")
	fmt.Scanln(&option)
	switch option {
	case 1:
		var number1, number2 int
		var optionContinue1 string
		fmt.Print("Escribe el primer numero: ")
		fmt.Scanln(&number1)
		fmt.Print("Escribe el segundo numero: ")
		fmt.Scanln(&number2)
		resultado, err := Sumar(number1, number2)
		if err != nil {
			wrapperError := ErrorWrapper(err, "suma")
			unWrapperError := UnwrapperError(wrapperError)
			if errors.Is(err, unWrapperError) {
				defer func() {
					recover()
					fmt.Println(wrapperError)
				}()
				panic(err)
			}
		}
		fmt.Println("El resultado es:", resultado)
		continues(optionContinue1)
	case 2:
		var number1, number2 int
		var optionContinue1 string
		fmt.Print("Escribe el primer numero: ")
		fmt.Scanln(&number1)
		fmt.Print("Escribe el segundo numero: ")
		fmt.Scanln(&number2)
		resultado, err := Restar(number1, number2)
		if err != nil {
			wrapperError := ErrorWrapper(err, "resta")
			unWrapperError := UnwrapperError(wrapperError)
			if errors.Is(err, unWrapperError) {
				defer func() {
					recover()
					fmt.Println(wrapperError)
				}()
				panic(err)
			}
		}
		fmt.Println("El resultado es:", resultado)
		continues(optionContinue1)
	case 3:
		var number1, number2 float64
		var optionContinue1 string
		fmt.Print("Escribe el primer numero: ")
		fmt.Scanln(&number1)
		fmt.Print("Escribe el segundo numero: ")
		fmt.Scanln(&number2)
		resultado, err := Multiplicar(number1, number2)
		if err != nil {
			wrapperError := ErrorWrapper(err, "multiplicacion")
			unWrapperError := UnwrapperError(wrapperError)
			if errors.Is(err, unWrapperError) {
				defer func() {
					recover()
					fmt.Println(wrapperError)
				}()
				panic(err)
			}
		}
		fmt.Println("El resultado es:", resultado)
		continues(optionContinue1)
	case 4:
		var number1, number2 float64
		var optionContinue1 string
		fmt.Print("Escribe el primer numero: ")
		fmt.Scanln(&number1)
		fmt.Print("Escribe el segundo numero: ")
		fmt.Scanln(&number2)
		resultado, err := Dividir(number1, number2)
		if err != nil {
			wrapperError := ErrorWrapper(err, "division")
			unWrapperError := UnwrapperError(wrapperError)
			if errors.Is(err, unWrapperError) {
				defer func() {
					cadena := recover()
					fmt.Println(wrapperError, cadena)
				}()
				panic(err)
			}
		}
		fmt.Println("El resultado es:", resultado)
		continues(optionContinue1)
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Debes de escoger alguna estas opciones")
	}
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("\033[H\033[2J")
	}
}

func main() {

	for {
		ClearTerminal()
		Menu()
	}
}
