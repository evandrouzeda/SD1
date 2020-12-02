package comands

import (
	"fmt"
	"io"
)

// Leitor implements buffering for an io.Reader object.
type Leitor struct {
	Buf  []byte
	rd   io.Reader // reader provided by the client
	r, w int       // buf read and write positions
}

//CriaLeitor create a reader
func CriaLeitor(rd io.Reader, size int) *Leitor {
	r := new(Leitor)
	r.reset(make([]byte, size), rd)
	return r
}

func (b *Leitor) reset(buf []byte, r io.Reader) {
	*b = Leitor{
		Buf: buf,
		rd:  r,
	}
}

//Buf returns the buf from a message, but it`s not working as expected
func Buf(buf *Leitor) []byte {
	n, err := buf.rd.Read(buf.Buf[buf.w:])
	if err != nil {
		fmt.Println(err)
	}
	return buf.Buf[:n]
}
