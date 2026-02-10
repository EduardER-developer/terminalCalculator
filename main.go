package main

import "fmt"

func main()  {
	var mode string

	var num1 int
	var num2 int

	fmt.Printf("Hello. This a calculator console programm. This programm have *multiplication*, *division*, *plus*, *minus* \n")

	fmt.Printf("Hava a start! Select mode: (choose multiplication or division or plus or minus)\n")
	fmt.Scanln(&mode)

	fmt.Printf("%s \n", mode);

	switch mode {
	case "division":
		fmt.Printf("Write first number: \n")
		fmt.Scanln(&num1)
		fmt.Printf("Write second number: \n")
		fmt.Scanln(&num2)
		sum := num1 * num2
		fmt.Printf("Your division sum number is %d\n", sum)
	}
}