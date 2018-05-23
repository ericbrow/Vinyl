package main

import (
    "fmt"
    "log"
    "regexp"
)

func main() {

    example := "#GoLang\"Code!$!"

    // Make a Regex to say we only want
    reg, err := regexp.Compile("[\"]+")
    if err != nil {
        log.Fatal(err)
    }
    processedString := reg.ReplaceAllString(example, "")

    fmt.Printf("A string of %s becomes %s \n", example, processedString)

    // Will output: 'A string of #GoLangCode!$! becomes GoLangCode'
}