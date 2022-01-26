package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Golang")) //polymorphic

	//incrementer
	myint := IntCounter(0)
	var inc Incrementer = &myint
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	// Composition call here
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello everyone this is test bytes"))
	wc.Close()

	r, ok := wc.(*BufferedWriterCloser) //type conversion with pointer
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	//empty interface to handle value not pointer in line 23
	var obj interface{} = NewBufferedWriterCloser()
	if wc, ok := obj.(WriterCloser); ok {
		wc.Write([]byte("Hello everyone this empty interface testing"))
		wc.Close()
	}
	n, ok := obj.(io.Reader)
	if ok {
		fmt.Println(n)
	} else {
		fmt.Println("Conversion failed")
	}

}

type Writer interface { // by conventio name should be postfix ..._er
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

//composition with interface
type Closer interface {
	Close() error
}
type WriterCloser interface { //composition
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}

	}
	return n, nil
}
func (bwc *BufferedWriterCloser) Close() error {

	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

//constructor to init
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

/** Guidelines for interface
1. Method set of value is all methods with value receivers
2. Method set of pointer is all methods, regardless of reciever type

*/
