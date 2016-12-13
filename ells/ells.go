package main

import  (
  "fmt"
  "bufio"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  name, _ := reader.ReadString('\n')
  
  if name == "ells" {
    fmt.Printf("Goodbye Radek\n")
  } else if name == "giles.bentley" {
    fmt.Printf("Shut the door when you leave.\n")
  }
}
