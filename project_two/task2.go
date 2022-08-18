package main

import (
    "fmt"
)

func swap(x, y string)(string, string){
    return y, x
}


func main() {
    var i, j int
    var myarrayint = [5]int {1, 2, 3, 4}
    var myIdentity = [5]string {"fname-> sheriff",
     "lname-> yusuf", "stack-> backend dev", "lang-> go lang"}

    fmt.Println("My details are:-")
    for i = 0; i < 5; i++ {

        fmt.Println(myIdentity[i])
    }

    for j = 0; j < 5; j++ {

        fmt.Print(myarrayint[j], " ")
    }
    

   
}