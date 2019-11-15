package main

// ../bin/go run ../try_handle.go

import (
    "fmt"
    "io/ioutil"
)

func main() {
    tryx {
        bytes := ioutil.ReadFile("try_handle.go")
        _ = fmt.Println(string(bytes))
    } handlex err {
        fmt.Println("ERR1", err)
        return
    }

    // try {
    //     bytes, := ioutil.ReadFile("try_handle_DNE.go")
    //     _ = fmt.Println(string(bytes))
    // } handle err {
    //     fmt.Println("ERR2", err)
    //     return
    // }

    //     bytes, err := ioutil.ReadFile("try_handle.go")
//     if err != nil {
//         fmt.Println("ERR1", err)
//         return
//     }

//     err = fmt.Println(string(bytes))
//     if err != nil {
//         fmt.Println("ERR1", err)
//         return
//     }

//     bytes, err = ioutil.ReadFile("try_handle_DNE.go")
//     if err != nil {
//         fmt.Println("ERR2", err)
//         return
//     }

//     err = fmt.Println(string(bytes))
//     if err != nil {
//         fmt.Println("ERR2", err)
//         return
//     }
}
