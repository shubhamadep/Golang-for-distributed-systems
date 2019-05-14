package main 

import "net/rpc"
import "fmt"
import "log"
import "sync"
import "io/ioutil"
import "strings"
// import "reflect"
// import "/Users/shubhamadep/Documents/Distributed Systems/Shubham_Assignment_4/server"

func main() {

  var wg sync.WaitGroup
  b, _  := ioutil.ReadFile("input.txt") // b has type []byte
  words := strings.Split(string(b), "\n")
  
  wg.Add(100)

  for i := range words {
    client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
    if err != nil {
      log.Fatal("dialing:", err)
    } 
    go send_request(words[i], client, &wg)
  }
  wg.Wait()
}

func send_request(word string, client *rpc.Client, wg *sync.WaitGroup) {

  defer wg.Done()
  var reply string
  err := client.Call("Arith.GetMeaning", &word, &reply)
  if err != nil {
    log.Fatal("arith error:", err)
  }
  fmt.Println(reply)
  

}