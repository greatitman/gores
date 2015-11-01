package main

import "fmt"
import "os"

var user = os.Getenv("user")
func main(){

	s := `hello`
	c :=[]byte(s)
	c[0] = 'j'
	var s2 = string(c)
	slice := []int{1,2,3}
	fmt.Printf("slice=\n",slice)
	fmt.Printf("%s\n",s2)
	if user==""{
		panic("no value for $USER")
	}
}
