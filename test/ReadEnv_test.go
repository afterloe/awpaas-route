package test

import (
	"os"
	"fmt"
	"testing"
)

func Test_readEnv(t *testing.T) {
	arg_num := len(os.Args)
	fmt.Printf("the num of input is %d\n",arg_num)

	fmt.Printf("they are :\n")
	for i := 0 ; i < arg_num ;i++{
		fmt.Println(os.Args[i])
	}

	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	fmt.Println("------------------------------------------------------------\n")
	logname := os.Getenv("info")
	fmt.Printf("logname is %s\n",logname)
}