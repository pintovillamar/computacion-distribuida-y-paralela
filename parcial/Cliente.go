package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	//"strconv"
	//"math"
)

// send 3 messages to the server

// func sendData(f func(float64), float64, a, b float64, n int) float64 {
// 	data := 0
// 	h := (b - a) / float64(n)
// 	sum := 0.5 * (f(a) + f(b))
// 	for i := 1; i < n; i++ {
// 		sum += f(a + float64(i)*h)
// 	}
// 	data := sum * h
// 	data := strconv.Itoa(data)
// 	return data
// }

func main() {

	// f := func(x float64) float64 {
	// 	return ((math.Pow(x, 2) + 1) / 2)
	// }

	arguments := os.Args
	//Verifica que se ingrese una dirección y un puerto. Ejemplo: 127.0.0.1:1234
	if len(arguments) == 1 {
		fmt.Println("Por favor ingrese un host y puerto de la siguiente manera:  host:port ")
		return
	}
	CONNECT := arguments[1]

	//Devuelve una direccion UDP
	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("El servidor UDP es %s\n", c.RemoteAddr().String())
	//Con defer se asegura que lo último que se realice sea  cerrar la conección
	defer c.Close()

	//Se crea un ciclo for(true)
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Mensaje a enviar:  ")
		text, _ := reader.ReadString('\n')
		data := []byte(text)
		_, err = c.Write(data)
		//Cuando se ingrese el mensaje STOP, se detiene el cliente
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Saliendo del cliente UDP")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
