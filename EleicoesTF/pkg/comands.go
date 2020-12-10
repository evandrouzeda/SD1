package comands

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//Candidato is as struct to represent a file
type Candidato struct {
	Nome  string `json:"nome"`
	Num   string `json:"num"`
	Votos int    `json:"votos"`
}

//Voto is a struct to represent a vote
type Voto struct {
	Num string `json:"num"`
}

//CriaCandidato create a file
func CriaCandidato(nome string, num string) Candidato {
	return Candidato{Nome: nome, Num: num, Votos: 0}
}

//Login is a struct to send a login comand
type Login struct {
	Comando string `json:"cmd"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

//LOGIN returns a login struct
func LOGIN(id string, usuario string, senha string) Login {
	return Login{id, usuario, senha}
}

//Loginr is a struct to reply a login request
type Loginr struct {
	Comando string `json:"cmd"`
	Codigo  string `json:"cod"`
}

//LOGINR retuns a Loginr struct
func LOGINR() Loginr {
	return Loginr{Comando: "loginr"}
}

//Logout is a struct to send a logout comand
type Logout struct {
	Comando string `json:"cmd"`
}

//LOGOUT returns a Logout struct
func LOGOUT(id string) Logout {
	return Logout{id}
}

//Logoutr is a struct to send a logout comand
type Logoutr struct {
	Comando string `json:"cmd"`
	Codigo  string `json:"cod"`
}

//LOGOUTR returns a Logout struct
func LOGOUTR() Logoutr {
	return Logoutr{Comando: "logoutr"}
}

//List returns a upload struct
type List struct {
	Comando string `json:"cmd"`
}

//LIST returns a List struct
func LIST(comando string) List {
	return List{comando}
}

//Listr returns a upload struct
type Listr struct {
	Comando string      `json:"cmd"`
	Cod     string      `json:"cod"`
	Lista   []Candidato `json:"lista"`
}

//LISTR returns
func LISTR() Listr {
	return Listr{Comando: "listr"}
}

//Cadas 'e a struct para mensagem de cadastro
type Cadas struct {
	Comando string `json:"cmd"`
	Type    string `json:"type"`
	Qtd     int    `json:"qtd"`
	Num     string `json:"num"`
	Nome    string `json:"nome"`
}

//CriaCadasDecla cria uma struct Cadas com os parametros necessarios para tipo declaracao
func CriaCadasDecla(qtd int) Cadas {
	return Cadas{Comando: "cadas", Type: "decla", Qtd: qtd}
}

//CriaCadasCand cria uma struct Cadas com os parametros necessarios para tipo candidato
func CriaCadasCand(nome string, numero string) Cadas {
	return Cadas{Comando: "cadas", Type: "candi", Nome: nome, Num: numero}
}

//Cadasr 'e a struct para mensagem de cadastro
type Cadasr struct {
	Comando string `json:"cmd"`
	Cod     string `json:"cod"`
}

//CADASR returns
func CADASR() Cadasr {
	return Cadasr{Comando: "cadasr"}
}

//Inicia 'e a struct para mensagem de cadastro
type Inicia struct {
	Comando string `json:"cmd"`
}

//INICIA returns Inicia
func INICIA() Inicia {
	return Inicia{Comando: "inicia"}
}

//Iniciar 'e a struct para fazer o replay do inicia
type Iniciar struct {
	Comando string `json:"cmd"`
	Cod     string `json:"cod"`
}

//INICIAR returns Iniciar
func INICIAR() Iniciar {
	return Iniciar{Comando: "iniciar"}
}

//Final 'e a struct para mensagem de cadastro
type Final struct {
	Comando string `json:"cmd"`
}

//FINAL returns Final
func FINAL() Final {
	return Final{Comando: "final"}
}

//Finalr 'e a struct para fazer o replay do Final
type Finalr struct {
	Comando string `json:"cmd"`
	Cod     string `json:"cod"`
}

//FINALR returns Finalr
func FINALR() Finalr {
	return Finalr{Comando: "finalr"}
}

//Apura struct
type Apura struct {
	Comando string `json:"cmd"`
}

//APURA returns Final
func APURA() Apura {
	return Apura{Comando: "apura"}
}

//Apurar 'e a struct para fazer o replay do Final
type Apurar struct {
	Comando  string      `json:"cmd"`
	Cod      string      `json:"cod"`
	Apuracao []Candidato `json:"apuracao"`
}

//APURAR returns Finalr
func APURAR() Apurar {
	return Apurar{Comando: "apurar"}
}

//Votar 'e a struct para fazer o replay do Final
type Votar struct {
	Comando string `json:"cmd"`
	Num     string `json:"num"`
}

//VOTAR returns Finalr
func VOTAR(num string) Votar {
	return Votar{Comando: "votar", Num: num}
}

//Votarr struct
type Votarr struct {
	Comando string `json:"cmd"`
	Cod     string `json:"cod"`
}

//VOTARR returns Final
func VOTARR() Votarr {
	return Votarr{Comando: "votarr"}
}

//Resul struct
type Resul struct {
	Comando string `json:"cmd"`
}

//RESUL returns a List struct
func RESUL() Resul {
	return Resul{Comando: "resul"}
}

//Resulr returns a upload struct
type Resulr struct {
	Comando   string      `json:"cmd"`
	Cod       string      `json:"cod"`
	Resultado []Candidato `json:"resul"`
}

//RESULR returns
func RESULR() Resulr {
	return Resulr{Comando: "resulr"}
}

//SendMSG is function to send the message to server
func SendMSG(ln net.Conn, msg interface{}) {
	js, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	//s1 := string(js)
	//fmt.Println(s1)
	ln.Write(js)
}

//Wait is a function that wait for a message from the client
func Wait(ln net.Conn) (*Leitor, int) {
	msg := CriaLeitor(ln, 4096)
	n, err := msg.rd.Read(msg.Buf[msg.w:])
	if err != nil {
		fmt.Println("Cliente desconectou", err)
	}
	return msg, n
}

//WaitR is a function that wait for Replay
func WaitR(ln net.Conn, cmd interface{}) interface{} {
	msg := CriaLeitor(ln, 4096)
	n, err := msg.rd.Read(msg.Buf[msg.w:])
	if err != nil {
		fmt.Println("Cliente desconectou", err)
		os.Exit(3)
	}

	erro := json.Unmarshal(msg.Buf[:n], &cmd)
	if erro != nil {
		fmt.Println(erro)
	}

	//fmt.Println(cmd)
	TornaStruct(msg.Buf[:n], &cmd)
	return cmd
}

//WaitResul is a function that wait for Replay
func WaitResul(ln net.Conn, data interface{}) {
	msg := CriaLeitor(ln, 4096)
	n, err := msg.rd.Read(msg.Buf[msg.w:])
	if err != nil {
		fmt.Println("Cliente desconectou", err)
		os.Exit(3)
	}
	erro := json.Unmarshal(msg.Buf[:n], &data)
	if erro != nil {
		fmt.Println(erro)
	}
	var cmd = FindComando(data)

	if cmd == "resulr" {
		TornaStruct(msg.Buf[:n], &data)
	}
	//fmt.Println(cmd)
}

//FindComando procura o atributo comando dentro da interface
func FindComando(msg interface{}) string {
	mapa := msg.(map[string]interface{})
	for k, v := range mapa {
		if k == "cmd" {
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}

//TornaStruct torna um json em struct
func TornaStruct(msg []byte, destino interface{}) {
	erro := json.Unmarshal(msg, &destino)
	if erro != nil {
		fmt.Println(erro)
	}
}
