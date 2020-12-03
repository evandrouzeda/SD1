package main

import (
	comands "SD1/EleicoesTF/pkg"
	"encoding/json"
	"fmt"
	"net"
)

//Cliente guarda infos do cliente
type Cliente struct {
	id         int
	conexao    net.Conn
	autorizado bool
	admin      bool
}

//Estados diz qual 'e o estado da eleicao
type Estados struct {
	acontecendo bool
	inicio      string
	final       string
}

//Candidatos e uma lista de Candidatos
var Candidatos []comands.Candidato

//Votos e uma lista de votos
var Votos []comands.Voto

//Resultado 'e uma lista de candidatos com o numero votos
var Resultado []comands.Candidato

//Eleicao guarda os estados da eleicao
var Eleicao Estados

//funcao para procurar aquivo na lista
func procuraCandidato(numero string) bool {
	for i := 0; i < len(Candidatos); i++ {
		if Candidatos[i].Num == numero {
			return true
		}
	}
	return false
}

func trataErros(err error) {
	if err != nil {
		//fmt.Println(err)
	}
}

//aqui e onde as mensagem dos cliente sao tratadas
func fileMenagement(cliente Cliente) {
	for {
		var jsn interface{}

		msg, n := comands.Wait(cliente.conexao)
		erro := json.Unmarshal(msg.Buf[:n], &jsn)
		trataErros(erro)
		if jsn == nil {
			break
		}

		var cmd = comands.FindComando(jsn)
		if cliente.autorizado {
			if cliente.admin {
				switch cmd {
				case "cadas":
					var des comands.Cadas
					comands.TornaStruct(msg.Buf[:n], &des)
					fmt.Println(des)
					if des.Type == "decla" {
						for i := 0; i < des.Qtd; i++ {
							cand, n := comands.Wait(cliente.conexao)
							var candidato comands.Candidato
							comands.TornaStruct(cand.Buf[:n], &candidato)
							Candidatos = append(Candidatos, comands.CriaCandidato(candidato.Nome, candidato.Num))
							fmt.Println("candidato inserido")
						}
						msg := comands.CADASR()
						msg.Cod = "Ok"
						comands.SendMSG(cliente.conexao, msg)
					} else {
						msg := comands.CADASR()
						msg.Cod = "ERRO"
						comands.SendMSG(cliente.conexao, msg)
					}
					break
				case "inicia":
					msg := comands.INICIAR()
					if Eleicao.acontecendo {
						msg.Cod = "ERRO"
					} else {
						Eleicao.acontecendo = true
						msg.Cod = "Ok"
					}
					comands.SendMSG(cliente.conexao, msg)
					break
				case "final":
					msg := comands.FINALR()
					if Eleicao.acontecendo {
						Eleicao.acontecendo = false
						msg.Cod = "Ok"
					} else {
						msg.Cod = "ERRO"
					}
					comands.SendMSG(cliente.conexao, msg)
					break
				case "apura":
					msg := comands.APURAR()
					msg.Cod = "Ok"
					for i := 0; i < len(Votos); i++ {
						for j := 0; j < len(Candidatos); j++ {
							if Votos[i].Num == Candidatos[j].Num {
								Candidatos[j].Votos++
							}
						}
					}
					// tem que fazer um ordenador
					/* for j := 0; j < len(Candidatos); j++ {
						if Votos[i].Num == Candidatos[j].Num {
							Candidatos[j].Votos++
						}
					} */
					comands.SendMSG(cliente.conexao, msg)
					break
				}
			}
			switch cmd {
			case "logout":
				cliente.conexao.Close()
				break
			case "list":
				reply := comands.LISTR()
				if Eleicao.acontecendo {
					var des comands.List
					comands.TornaStruct(msg.Buf[:n], &des)
					reply.Cod = "OK"
					reply.Lista = Candidatos
				} else {
					reply.Cod = "ERRO"
					var vazioCand []comands.Candidato
					reply.Lista = vazioCand
				}
				comands.SendMSG(cliente.conexao, reply)
				break
			case "votar":
				reply := comands.VOTARR()
				if Eleicao.acontecendo {
					var des comands.Votar
					comands.TornaStruct(msg.Buf[:n], &des)
					if procuraCandidato(des.Num) {
						reply.Cod = "Ok"
						num := comands.Voto{Num: des.Num}
						Votos = append(Votos, num)
					} else {
						reply.Cod = "ERRO"
					}
				} else {
					reply.Cod = "ERRO"
				}
				comands.SendMSG(cliente.conexao, reply)
				break
			case "resul":
				reply := comands.RESULR()
				if !Eleicao.acontecendo {
					var des comands.Resul
					comands.TornaStruct(msg.Buf[:n], &des)
					reply.Cod = "OK"
					reply.Lista = Resultado
				} else {
					reply.Cod = "ERRO"
				}
				comands.SendMSG(cliente.conexao, reply)
				break
			default:
				fmt.Println("comando invalido")
				break
			}
		} else {
			switch cmd {
			case "login":
				var des comands.Login
				comands.TornaStruct(msg.Buf[:n], &des)
				if des.Usuario == "admin" {
					cliente.autorizado = true
					cliente.admin = true
				} else {
					cliente.autorizado = true
					cliente.admin = false
				}
				msg := comands.LOGINR()
				msg.Codigo = "OK"
				comands.SendMSG(cliente.conexao, msg)
				break
			default:
				fmt.Println("comando invalido")
				msg := comands.LOGINR()
				msg.Codigo = "INVALID"
				cliente.autorizado = true
				comands.SendMSG(cliente.conexao, msg)
				break
			}
		}
	}
}

func main() {
	fmt.Println("Servidor aguardando conexao...")
	ln, erro := net.Listen("tcp", ":8081")

	trataErros(erro)

	defer ln.Close()

	for {
		conexao, erro1 := ln.Accept()

		trataErros(erro1)

		fmt.Println("Conexao Aceita...")

		cliente := Cliente{
			id:         1,
			conexao:    conexao,
			autorizado: false,
		}

		go fileMenagement(cliente)
	}
}
