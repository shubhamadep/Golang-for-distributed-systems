package main

import "net/rpc"
import "net"
import "fmt"
import "net/http"
import "log"
import "bufio"
import "strings"
import "os"


var dictionary = make(map[string]string)

type Arith string
var threadNumber int

func (t *Arith) GetMeaning(w *string, reply *string) error {
	word := *w
	threadNumber++
	fmt.Println("Thread : ", threadNumber, " serving word : ", *w)
	if dictionary[word] != "" {
		*reply = word+" : "+dictionary[word]
	} else{
		*reply = word+" : "+"Please check the spelling again."
	}
	return nil
}

func main() {

	make_dictionary()
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	var sem = make(chan int, 10)
	for {
		err := http.Serve(l, nil)
		if err != nil {
			log.Fatal("Error serving: ", err)
		}
	}

}

func make_dictionary() {

    f, err := os.Open("Oxford_English_Dictionary.txt")

    reader := bufio.NewReader(f)
    var line string
    line, err = reader.ReadString('\n')

    for {
      line, err = reader.ReadString('\n')
      word := strings.Split(line, " ")[0]
      meaning := strings.Join(strings.Split(line, " ")[1:], " ")
      dictionary[strings.ToLower(word)] = meaning
      if err != nil {
        break
      }
    }
    fmt.Println("Dictionary is ready.")
}
