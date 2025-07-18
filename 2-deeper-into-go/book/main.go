package main

import "fmt"

// func main() {
//     fmt.Println(getBookTitle())
// }

// func getBookTitle() string {
//     return "Harry Potter"
// }

////////////////////////////////////////

// type bookName string

// func (book bookName) printTitle() {
//     fmt.Println(bookName)
// }

// func main() {
//     var book bookName = "Harry Potter"
//     book.printTitle()
// }

////////////////////////////////////////

// type bookSize float64

// func (this bookSize) getSizeOfBook() bookSize {
//     return this
// }

// func main() {
//     var bookSize bookSize = 6.5
//     fmt.Println("Book size:", bookSize.getSizeOfBook())
// }

////////////////////////////////////////

// func main() {
//     title, pages, weight := getBookInfo()
//     fmt.Println(title, pages, weight)
// }

// func getBookInfo() (string, int, float64){
//     return "War and Peace", 1000, 2.5
// }

////////////////////////////////////////

// func main() {
//    b := book("Harry Potter")

//    fmt.Println(b.describe("is an awesome book"))
// }

// type book string

// func (b book) describe(description string) string {
//    return string(b) + " " + description
// }

////////////////////////////////////////

func main (){
    b := bookInfo{
        title: "Harry Potter",
        pages: 1000,
        weight: 2.5,
    }
    fmt.Println(b.getBookInfo())
}

type bookInfo struct {
    title string
    pages int
    weight float64
}

func (b bookInfo) getBookInfo() (string, int, float64) {
    return b.title, b.pages, b.weight
}