package main

import (
	comands "SD1/File1410/pkg"
	"encoding/json"
	"fmt"
	"net"
)

//Cliente guarda infos do cliente
type Cliente struct {
	id         int
	conexao    net.Conn
	autorizado bool
}

//Lista de aqruivos
var Lista []comands.Arquivo

//funcao para procurar aquivo na lista
func procuraArquivo(nome string) (bool, int) {
	for i := 0; i < len(Lista); i++ {
		if Lista[i].Nome == nome {
			return true, i
		}
	}
	return false, -1
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

		var cmd = comands.FindComando(jsn)
		if cliente.autorizado {
			switch cmd {
			case "login":
				msg := comands.LOGINR()
				msg.Codigo = "OK"
				cliente.autorizado = true
				comands.SendMSG(cliente.conexao, msg)
				break
			case "logout":
				cliente.conexao.Close()
				break
			case "list":
				var des comands.List
				comands.TornaStruct(msg.Buf[:n], &des)

				msg := comands.LISTR()
				msg.Codigo = "OK"
				msg.Lista = Lista
				comands.SendMSG(cliente.conexao, msg)
				break
			case "upload":
				var des comands.Upload
				comands.TornaStruct(msg.Buf[:n], &des)
				fmt.Println(des.Arquivo)

				Lista = append(Lista, des.Arquivo)

				msg := comands.UPLOADR()
				msg.Codigo = "OK"
				msg.Res = "success"
				comands.SendMSG(cliente.conexao, msg)
				break
			case "search":
				var des comands.Search
				comands.TornaStruct(msg.Buf[:n], &des)
				res, _ := procuraArquivo(des.Nome)
				msg := comands.SEARCHR()
				if res {
					msg.Codigo = "OK"
					msg.Res = "encontrado"
				} else {
					msg.Codigo = "ERRO"
					msg.Res = "nao encontrado"
				}
				comands.SendMSG(cliente.conexao, msg)
				break
			case "download":
				var des comands.Download
				comands.TornaStruct(msg.Buf[:n], &des)
				res, i := procuraArquivo(des.Nome)
				msg := comands.DOWNLOADR()
				if res {
					msg.Codigo = "OK"
					msg.Arquivo = Lista[i]

				} else {
					msg.Codigo = "ERRO"
				}
				comands.SendMSG(cliente.conexao, msg)
				break
			default:
				fmt.Println("comando invalido")
				break
			}
		} else {
			switch cmd {
			case "login":
				msg := comands.LOGINR()
				msg.Codigo = "OK"
				cliente.autorizado = true
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

	conexao, erro1 := ln.Accept()

	trataErros(erro1)

	fmt.Println("Conexao Aceita...")

	cliente := Cliente{
		id:         1,
		conexao:    conexao,
		autorizado: false,
	}

	fileMenagement(cliente)
}
