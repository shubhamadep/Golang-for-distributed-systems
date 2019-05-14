package main

import "net"
import "fmt"
import "bufio"
import "os"
// import "time"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  // connect to this socket
  // p := make([] byte, 2048)
  addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
  conn, err := net.DialUDP("udp", nil, addr)
  check(err)
  scanner := bufio.NewScanner(os.Stdin)
  reader := bufio.NewReader(os.Stdin)
  var goahead = "y"
  buffer := make([]byte, 512)

  for goahead == "y"{
    // time.Sleep(1000*time.Millisecond)
    fmt.Print("Enter a word: ")
    text, err := reader.ReadString('\n')
    check(err)

    n, err := conn.Write([]byte(text))
    check(err)
    
    n, err = conn.Read(buffer)
    // fmt.Println(string(buffer[:n]))
    var message = string(buffer[:n])
    if message != "incorrect" {
      fmt.Print("Message from server: "+message)
      fmt.Println("Do you want to continue ? y/n")
      scanner.Scan()
      goahead = scanner.Text()
    } else{
      fmt.Println("Please check your spelling.")
      goahead = "y"
    }
  }
  _, cerr := conn.Write([]byte(goahead))
  check(cerr)
  conn.Close()

}