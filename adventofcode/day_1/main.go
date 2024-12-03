/*
NOTE:
Lo importante no es tanto aprender a leer el archivo, lo importante es entender:
Cómo organizar los arreglos de menor a mayor?
Qué es un algoritmo de ordenamiento?
Cuál algoritmo de Ordenamiento utilizar?
Por qué Quick_Sort y no Merge_Sort?
Como implementarlo?
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Abrir el archivo
	file, err := os.Open("day_1_input.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear una matriz para almacenar los valores
	var matriz [][]int
	var ejeX []int
	var ejeY []int

	// Leer línea por línea
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		// Dividir la línea en palabras separadas por espacios
		valores := strings.Fields(linea)

		// Convertir los valores a enteros
		fila := make([]int, len(valores))
		for i, valor := range valores {
			num, err := strconv.Atoi(valor)
			if err != nil {
				fmt.Println("Error al convertir el valor a entero:", err)
				return
			}
			fila[i] = num
		}
		// Agregar la fila a la matriz
		matriz = append(matriz, fila)
	}

	// Manejo de errores durante la lectura
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Crear los slices
	for _, fila := range matriz {
		ejeX = append(ejeX, fila[0])
		ejeY = append(ejeY, fila[1])
	}

	//Organizar los slices
	quickSort(&ejeX, 0, len(ejeX)-1)
	quickSort(&ejeY, 0, len(ejeY)-1)

	//Calcular distancias
	fmt.Printf("The distance is: %d\n", distance(ejeX, ejeY))

	//Calcular el similarity score
	fmt.Printf("The Similarity score is: %d\n", similarityScore(ejeX, ejeY))
}

func quickSort(array *[]int, left, right int) []int {
	// separar valores menores a left de pivot
	// y valores mayores a la right del pivot
	if left < right {
		arr := *array
		limit, origin := right, left
		pivot := arr[right]
		right--

		for left <= right {
			// buscar en left numero mayor que pivot
			for left < len(arr) && arr[left] < pivot {
				left++
			}
			// buscar en right numero menor que pivot
			for right >= 0 && arr[right] > pivot {
				right--
			}

			if left <= right {
				// intercambiar encontrados
				swap(array, left, right)
				// aumentar valores left
				left++
				// decrementar valores right
				right--
			}

		}
		// termina separacion left | right
		swap(array, left, limit)
		quickSort(array, origin, right)
		quickSort(array, left, limit)
	}

	return *array
}

func swap(array *[]int, left, right int) {
	// intercambiar valor left con right
	arr := *array
	temp := arr[left]
	arr[left] = arr[right]
	arr[right] = temp
}

func distance(arrayX []int, arrayY []int) int {
	distance := 0
	for i, x := range arrayX {
		d := x - arrayY[i]
		if d < 0 {
			d = d * -1
		}
		distance += d
	}

	return distance
}

func similarityScore(arrayX []int, arrayY []int) int {
	sS := 0

	for _, x := range arrayX {
		count := 0
		for _, y := range arrayY {
			if x == y {
				count++
			}
		}
		sS += x * count
	}

	return sS
}

/*
|1 - 2| = 1
|2 - 1| = 1
*/
