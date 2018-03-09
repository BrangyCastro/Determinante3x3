package main

import "fmt"

func main() {

	var a11, a12, a13, a21, a22, a23, a31, a32, a33 int

	fmt.Println("Ingrese la primer fila de la matriz: ")

	fmt.Scanf("%d\n", &a11)
	fmt.Scanf("%d\n", &a12)
	fmt.Scanf("%d\n", &a13)

	fmt.Println("Ingrese la segunda fila de la matriz: ")

	fmt.Scanf("%d\n", &a21)
	fmt.Scanf("%d\n", &a22)
	fmt.Scanf("%d\n", &a23)

	fmt.Println("Ingrese la tercera fila de la matriz: ")

	fmt.Scanf("%d\n", &a31)
	fmt.Scanf("%d\n", &a32)
	fmt.Scanf("%d\n", &a33)

	var matriz [3][3] int

	matriz [0][0] = a11
	matriz [0][1] = a12
	matriz [0][2] = a13

	matriz [1][0] = a21
	matriz [1][1] = a22
	matriz [1][2] = a23

	matriz [2][0] = a31
	matriz [2][1] = a32
	matriz [2][2] = a33

	for i:= 0; i < len(matriz); i++ {

		fmt.Print("|")

			for y:= 0; y < len(matriz[i]); y++ {

				fmt.Print(matriz[i][y])

				if y != len(matriz[i]) -1{

					fmt.Print("\t")

				}

			}
		fmt.Print("|")

	}

	r11 := matriz[1][1] * matriz[2][2]
	r12 := matriz[2][1] * matriz[1][2]
	r1 := r11 - r12
	uno := matriz[0][0] + r1

	r21 := matriz[1][0] * matriz[2][2]
	r22 := matriz[2][0] * matriz[1][2]
	r2 := r21 - r22
	dos := -matriz[0][1] + r2

	r31 := matriz[1][0] * matriz[2][1]
	r32 := matriz[2][0] * matriz[1][1]
	r3 := r31 - r32
	tres := matriz[0][2] + r3

	derteminante := uno + dos + tres

	fmt.Println("")

	fmt.Printf("El Determinante es: %d\n", derteminante)

}