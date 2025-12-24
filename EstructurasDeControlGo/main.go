package main

import "fmt"

func main() {

	saludo := hello("Alex")
	fmt.Println(saludo)
	sum, mul := calc(4, 5)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La multiplicación es: ", mul)

	/*for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("Tarde!")
	} else {
		fmt.Println("Noche!")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")

	case "linux":
		fmt.Println("Go run -> Linux")

	case "darwin":
		fmt.Println("Go run -> MAC OS")

	default:
		fmt.Println("Go run -> Otro OS")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Mañana!")

	case t.Hour() < 17:
		fmt.Println("Tarde!")

	default:
		fmt.Println("Noche!")
	}*/

}

func hello(name string) string {
	return "Hola, " + name
}

func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}
