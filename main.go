package main
import(
	"fmt"
	"strconv"
)

// a function to return two
func Two(x int ) int{
	return 2
}

func main(){
	two := strconv.Itoa( Two(50000) )
	fmt.Printf("hello world! " + two )
}