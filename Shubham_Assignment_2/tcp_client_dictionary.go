package main

import "net"
import "fmt"
import "bufio"
//import "os"
//import "strings"


func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  //scanner := bufio.NewScanner(os.Stdin)
  var goahead = "y"
  var i = 100000
  var text = ""
  for goahead == "y" {
    if i < 1{
      break
    }
    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter a word: ")
    // text, err := reader.ReadString('\n')
    //check(err)
    text = "zoo"+"\n"
    fmt.Fprintf(conn, text)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Meaning of the word ", text," is : "+message)
    i -= 1

    // if strings.Compare(message, "word incorrect" + "\n") != 0 {
    //   // fmt.Print("Meaning of the word ", text," is : "+message)
    //   // fmt.Println("Do you want to continue ? y/n")
    //   // scanner.Scan()
    //   // goahead = scanner.Text()
    //
    // } else{
    //   fmt.Println("Please check your spelling.")
    //   goahead = "y"
    // }
  }
  fmt.Fprintf(conn, "n")
  conn.Close()
}
