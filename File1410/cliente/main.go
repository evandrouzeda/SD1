package main

import (
	comands "SD1/File1410/pkg"
	"bufio"
	"fmt"
	"net"
	"os"
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
		case "upload":
			arquivo := comands.CriaArquivo(s[1], s[2])
			msg := comands.UPLOAD(s[0], arquivo)
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.UPLOADR()
			comands.WaitR(ln, cmd)
			break
		case "search":
			msg := comands.SEARCH(s[0], s[1])
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.SEARCHR()
			comands.WaitR(ln, cmd)
			break
		case "download":
			msg := comands.DOWNLOAD(s[0], s[1])
			comands.SendMSG(ln, msg)

			//Wait for the reply
			cmd := comands.DOWNLOADR()
			comands.WaitR(ln, cmd)
			break
		default:
			fmt.Println("comando invalido")
			break
		}

	}
}
