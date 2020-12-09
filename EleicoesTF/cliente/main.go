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
			comands.WaitR(ln, &cmd)
			if cmd.Codigo == "Ok" {
				fmt.Println("logado")
			}
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
			comands.WaitR(ln, &cmd)
			if cmd.Cod == "Ok" {
				for i := 0; i < len(cmd.Lista); i++ {
					fmt.Printf("Candidato: %v - %v \n", cmd.Lista[i].Nome, cmd.Lista[i].Num)
				}
			} else {
				fmt.Println(cmd.Cod)
			}
			break
		case "cadas":
			//vou ter que colocar um for aqui dentro para ficar enviando a
			if s[1] == "decla" {
				qtd, _ := strconv.Atoi(s[2])
				cmd := comands.CriaCadasDecla(qtd)
				comands.SendMSG(ln, cmd)
				for i := 0; i < qtd; i++ {
					reader := bufio.NewReader(os.Stdin)
					fmt.Printf("Candidato %v: ", i)
					text, _ := reader.ReadString('\n')
					cand := strings.Split(strings.Replace(text, "\n", "", 1), " ")
					cmd := comands.CriaCadasCand(cand[0], cand[1])
					comands.SendMSG(ln, cmd)
				}
				//Wait for the reply
				reply := comands.CADASR()
				comands.WaitR(ln, &reply)
				if reply.Cod == "Ok" {
					fmt.Println("Candidatos inseridos com sucesso!!!")
				} else {
					fmt.Println("Erro ao inserir candidatos...")
				}
			} else {
				fmt.Println("ainda nao foi declarado a quantidade de candidatos")
			}
			break
		case "inicia":
			msg := comands.INICIA()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			reply := comands.INICIAR()
			comands.WaitR(ln, &reply)
			if reply.Cod == "Ok" {
				fmt.Println("Eleicao iniciada com sucesso!!!")
			} else {
				fmt.Println("Erro ao inciar eleicao...")
			}
			break
		case "final":
			msg := comands.FINAL()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			reply := comands.FINALR()
			comands.WaitR(ln, &reply)
			if reply.Cod == "Ok" {
				fmt.Println("Eleicao finalizada com sucesso!!!")
			} else {
				fmt.Println("Erro ao finalizar eleicao...")
			}
			break
		case "apura":
			msg := comands.APURA()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			reply := comands.APURAR()
			comands.WaitR(ln, &reply)
			if reply.Cod == "Ok" {
				for i := 0; i < len(reply.Apuracao); i++ {
					fmt.Printf("Candidato: %v - %v \n", reply.Apuracao[i].Nome, reply.Apuracao[i].Votos)
				}
			} else {
				fmt.Println("Erro ao Apurar eleicao...")
			}
			break
		case "votar":
			msg := comands.VOTAR(s[1])
			comands.SendMSG(ln, msg)

			//Wait for the reply
			reply := comands.VOTARR()
			comands.WaitR(ln, &reply)
			if reply.Cod == "Ok" {
				fmt.Println("Voto contabilizado com sucesso!!!")
			} else {
				fmt.Println("Erro ao contabilizar voto...")
			}
			break
		case "resul":
			msg := comands.RESUL()
			comands.SendMSG(ln, msg)

			//Wait for the reply
			reply := comands.RESULR()
			comands.WaitR(ln, &reply)
			if reply.Cod == "Ok" {
				for i := 0; i < len(reply.Resultado); i++ {
					fmt.Printf("Candidato: %v - %v \n", reply.Resultado[i].Nome, reply.Resultado[i].Votos)
				}
			} else {
				go func() {
					reply := comands.RESULR()
					comands.WaitR(ln, &reply)
					fmt.Print("recebeu alguma Coisa")
					if reply.Cod == "Ok" {
						for i := 0; i < len(reply.Resultado); i++ {
							fmt.Printf("Candidato: %v - %v \n", reply.Resultado[i].Nome, reply.Resultado[i].Votos)
						}
					}
				}()
				fmt.Println("Esperando a Eleicao terminar")
			}
			break
		default:
			fmt.Println("comando invalido")
			break
		}

	}
}
