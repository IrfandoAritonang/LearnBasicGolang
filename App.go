package main

import "fmt"

func main() {
    // Mencetak teks sederhana ke konsol
    fmt.Println("Hello, World!")

    // Mencetak nilai variabel dengan formatting
    name := "John"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    // Format dan simpan dalam variabel string
    formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(formattedString)
}
