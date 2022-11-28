package main

import (
	"fmt"
	"math"
	"net"
	"os"
	"strings"
	"time"
)

func suma(a, b int) int {
	return a + b
}

func TrapezoidRule(f func(float64) float64, a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := 0.5 * (f(a) + f(b))
	for i := 1; i < n; i++ {
		sum += f(a + float64(i)*h)
	}

	return sum * h
}

func worker(jobs chan int, results chan float64) {
	f := func(x float64) float64 {
		return ((math.Pow(x, 2) + 1) / 2)
	}

	for n := range jobs {
		results <- TrapezoidRule(f, 5, 20, n)
	}
}

func main() {
	arguments := os.Args
	//Verifica que se ingrese un puerto. Ejemplo: 1234
	if len(arguments) == 1 {
		fmt.Println("Por favor ingrese un puerto")
		return
	}
	PORT := ":" + arguments[1]
	//Devuelve una direccion UDP
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ListenPack para redes UDP
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on localhost, port " + PORT)
	fmt.Println("Waiting for a message...")

	//Con defer se asegura que lo último que se realice sea  cerrar la conección
	defer connection.Close()

	// buffer := make([]byte, 1024)
	// asign suma to buffer
	buffer := make([]byte, 1024)

	//Se crea un ciclo for(true)
	for {

		//Lee el mensaje y adicionalmente nos devuelve la dirección del Cliente
		n, addr, err := connection.ReadFromUDP(buffer)

		//Cuando llegue el mensaje STOP, se detiene el servidor
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Saliendo del servidor UDP", addr)
			return
		}
		if strings.TrimSpace(string(buffer[0:n])) == "trapecio" {
			fmt.Println("trapecio")
			n := 1000

			jobs := make(chan int, n)
			results := make(chan float64, n)

			go worker(jobs, results)
			go worker(jobs, results)

			final_times := 0
			new_result := 0.0

			start := time.Now()
			for i := 1; i <= n; i++ {
				start_per_job := time.Now()
				jobs <- i
				elapsed_per_job := time.Since(start_per_job).Nanoseconds()
				// fmt.Println("Area in trapezoid[", i, "]", <-results, ":", elapsed_per_job, "ns")
				final_times += int(elapsed_per_job)
				new_result = <-results
			}
			close(jobs)
			elapsed := time.Since(start).Nanoseconds()
			a := fmt.Sprintf("%.2f", new_result)
			fmt.Println("Total Area with", n, "trapezoids:", a)
			fmt.Println("Total Time[native]: ", elapsed)
			fmt.Println("Total Time[optimz]: ", final_times)

			// for a := 1; a <= 9; a++ {
			// 	fmt.Println(<-results)
			// }
		}
		// fmt.Printf("Mensaje recibido de: %s\n", addr)
		// fmt.Print("Mensaje recibido: ", string(buffer[0:n]))

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
