package main


import (
  "flag"
  "fmt"
  "os"
  "strings"
)

func main() {
    charCount := flag.Bool("c", false, "parse byte counts")
    lineCount := flag.Bool("l", false, "get line counts")
    wordCount := flag.Bool("w", false , "get word counts")
    flag.Parse()
    tails := flag.Args()
    fileName := tails[0]
     
   
    content , err := os.ReadFile(fileName)
    if err!=nil{
      fmt.Println(err)
      os.Exit(-1)
    }
    if *charCount {
        fmt.Fprintf(os.Stdout, "%d %s\n", len(content), fileName)
    }
    if *lineCount {
        newLine := byte('\n') 
        count := 0 
    for _ , val := range content {
      if val == newLine {
        count++
      } 
    }
    fmt.Fprintf(os.Stdout, "%d %s\n",count, fileName)
    }

    if *wordCount {
      count := 0 
      strContent := string(content)
      for _, line := range strings.Split(strContent , "\n"){
        count += len(strings.Fields(line))
      }
      fmt.Fprintf(os.Stdout, "%d %s\n", count, fileName)
    }


}
