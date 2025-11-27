package main


import (
    "fmt"
    myerr "github.com/jfb0301/performance_optimization/core/errors"     // 1 
)

func shouldFail() bool { return true }

func noErrCanHappen() int {
    return 204
}

func doOrErr() error {
    if shouldFail() {
        return myerr.New("ups, XYZ2 failed")
    }
    return nil
}

func nestedDoOrErr() error {
    if err := doOrErr(); err != nil {        
        return myerr.Wrap(err, "do")      //  3
    }
    return nil
}

func main() {
    if err := nestedDoOrErr(); err != nil {  // 2 
        fmt.Printf("%+v\n", err)
    }
}
