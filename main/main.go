package main
import (
	"fmt"
	"github.com/satori/go.uuid"
)
func main(){
	u := uuid.NewV4();
	fmt.Printf("hello, world %s \n", u)
}
