package comands

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//Candidato is as struct to represent a file
type Candidato struct {
	Nome string `json:"nome"`
	Num  string `json:"num"`
}

//Voto is a struct to represent a vote
type Voto struct {
	Numero string `json:"numero"`
}

//CriaCandidato create a file
func CriaCandidato(nome string, numero string) Candidato {
	return Candidato{nome, numero}
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
	Codigo string `json:"codigo"`
}

//LOGINR retuns a Loginr struct
func LOGINR() Loginr {
	return Loginr{""}
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
	Codigo string `json:"codigo"`
}

//LOGOUTR returns a Logout struct
func LOGOUTR() Logoutr {
	return Logoutr{""}
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
	Codigo string      `json:"codigo"`
	Lista  []Candidato `json:"lista"`
}

//LISTR returns
func LISTR() Listr {
	return Listr{}
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

//CriaCadasCand cria uma struct Cadas com os parametros necessarios para tipo declaracao
func CriaCadasCand(nome string, numero string) Cadas {
	return Cadas{Comando: "cadas", Type: "decla", Nome: nome, Num: numero}
}

//Cadasr 'e a struct para mensagem de cadastro
type Cadasr struct {
	Cod string `json:"cod"`
}

//CADASR returns
func CADASR() Cadasr {
	return Cadasr{}
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
	Cod string `json:"cod"`
}

//INICIAR returns Iniciar
func INICIAR() Iniciar {
	return Iniciar{}
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
	Cod string `json:"cod"`
}

//FINALR returns Finalr
func FINALR() Finalr {
	return Finalr{}
}

//SendMSG is function to send the message to server
func SendMSG(ln net.Conn, msg interface{}) {
	js, err := json.Marshal(msg)
	if err != nil {
		fmt.Print(err)
		return
	}
	s1 := string(js)
	fmt.Println(s1)
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
func WaitR(ln net.Conn, cmd interface{}) {
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

	fmt.Println(cmd)
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
