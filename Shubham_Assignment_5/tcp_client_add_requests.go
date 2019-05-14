package main

import "net"
import "fmt"
import "bufio"
// import "os"
import "strings"
//import "time"
import "sync"
import "io/ioutil"


func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  var wg sync.WaitGroup
  b, err := ioutil.ReadFile("input.txt") // b has type []byte
  check(err)
  words := strings.Split(string(b), "\n")

  //words := []string{"dog", "cool", "dogmatic", "fussy", "good"}
  wg.Add(100)
  for i := range words {
    conn, _ := net.Dial("tcp", "127.0.0.1:8082")
    go send_request(words[i], conn, &wg)
  }
  wg.Wait()
}

func send_request(word string, conn net.Conn, wg *sync.WaitGroup) {

  defer wg.Done()
  fmt.Fprintf(conn, word+"\n")
  m,_ := bufio.NewReader(conn).ReadString('\n')
  fmt.Println(m)


}
