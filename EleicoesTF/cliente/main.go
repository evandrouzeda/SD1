package main

import (
	comands "SD1/EleicoesTF/pkg"
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func trataErros(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ln, err := net.Dial("tcp", "localhost:8081")

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	for {
		//Reads what the user type on prompt
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		//Divide the words on a array
		s := strings.Split(strings.Replace(text, "\n", "", 1), " ")

		switch s[0] {
		case "login":
			msg := comands.LOGIN(s[0], s[1], s[2])
			comands.SendMSG(ln, msg)

			cmd := comands.LOGINR()
			comands.WaitR(ln, cmd)
			break
		case "logout":
			msg := comands.LOGOUT(s[0])
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.LOGOUTR()
			comands.WaitR(ln, cmd)
			break
		case "list":
			msg := comands.LIST("list")
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.LISTR()
			comands.WaitR(ln, cmd)
			break
		case "cadas":
			//vou ter que colocar um for aqui dentro para ficar enviando a
			if s[1] == "decla" {
				qtd, _ := strconv.Atoi(s[2])
				cmd := comands.CriaCadasDecla(qtd)
				comands.SendMSG(ln, cmd)
			} else if s[1] == "cand" {
				cmd := comands.CriaCadasCand(s[2], s[3])
				comands.SendMSG(ln, cmd)
			}
			break
		case "inicia":
			msg := comands.INICIA()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.INICIAR()
			comands.WaitR(ln, cmd)
			break
		case "final":
			msg := comands.FINAL()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.FINALR()
			comands.WaitR(ln, cmd)
			break
		case "apura":
			msg := comands.FINAL()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.FINALR()
			comands.WaitR(ln, cmd)
			break
		default:
			fmt.Println("comando invalido")
			break
		}

	}
}
