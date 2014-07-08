// this is dummy test demo which firstly writes the file and then read the same file.
package main
import (
	"fmt"
	"io/ioutil"
)
type Page struct {
	FileName string
	Data []byte
}
func (p *Page) saveFile() error{
	filename := p.FileName + ".txt"
	return ioutil.WriteFile(filename,p.Data,0600)
} 

func loadPage(name string)  (*Page, error) {
	fileName := name + ".txt"
	data,err :=ioutil.ReadFile(fileName)
	if err !=nil {
		return nil,err
	}
	return &Page{FileName: fileName, Data: data},nil
}

func main() {
	p1 := &Page{FileName: "TestPage",Data: []byte("This is an Sample Demo")}
	p1.saveFile()
	p2,err :=loadPage("TestPage")
	if err!=nil{
		fmt.Println(string("Error occured"))
	}else{
		fmt.Println(string(p2.Data))	
	}
}
