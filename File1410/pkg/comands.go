package comands

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//Arquivo is as struct to represent a file
type Arquivo struct {
	Nome string `json:"nome"`
	Dono string `json:"dono"`
}

//CriaArquivo create a file
func CriaArquivo(nome string, dono string) Arquivo {
	return Arquivo{nome, dono}
}

//Login is a struct to send a login comand
type Login struct {
	Comando string `json:"comando"`
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
	Comando string `json:"comando"`
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

//Upload struct to be transport
type Upload struct {
	Comando string  `json:"comando"`
	Arquivo Arquivo `json:"arquivo"`
}

//UPLOAD returns a upload struct
func UPLOAD(comando string, arquivo Arquivo) Upload {
	return Upload{comando, arquivo}
}

//Uploadr returns a upload struct
type Uploadr struct {
	Codigo string `json:"codigo"`
	Res    string `json:"res"`
}

//UPLOADR returns
func UPLOADR() Uploadr {
	return Uploadr{}
}

//Search returns a upload struct
type Search struct {
	Comando string `json:"comando"`
	Nome    string `json:"nome"`
}

//SEARCH returns a upload struct
func SEARCH(comando string, nome string) Search {
	return Search{comando, nome}
}

//Searchr returns a upload struct
type Searchr struct {
	Codigo string `json:"codigo"`
	Res    string `json:"res"`
}

//SEARCHR returns
func SEARCHR() Searchr {
	return Searchr{}
}

//Download returns a upload struct
type Download struct {
	Comando string `json:"comando"`
	Nome    string `json:"nome"`
}

//DOWNLOAD returns a upload struct
func DOWNLOAD(comando string, nome string) Download {
	return Download{comando, nome}
}

//Downloadr returns a upload struct
type Downloadr struct {
	Codigo  string  `json:"codigo"`
	Arquivo Arquivo `json:"arquivo"`
}

//DOWNLOADR returns
func DOWNLOADR() Downloadr {
	return Downloadr{}
}

//List returns a upload struct
type List struct {
	Comando string `json:"comando"`
}

//LIST returns a List struct
func LIST(comando string) List {
	return List{comando}
}

//Listr returns a upload struct
type Listr struct {
	Codigo string    `json:"codigo"`
	Lista  []Arquivo `json:"lista"`
}

//LISTR returns
func LISTR() Listr {
	return Listr{}
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
		os.Exit(3)
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
		if k == "comando" {
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
