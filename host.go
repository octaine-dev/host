package host

import (
	"log"
	"fmt"

	"github.com/octaine-dev/lib"
)

type RNG struct {}
func (r RNG) Next(query []int) ([]int, error) {
	fmt.Println("RNG callback ", query)
	return []int{4,5,6}, nil
}

func main() {
	resp, err := lib.Play(&lib.PlayRequest{lib.Bet{100,"0.01"},[]string{"state"},[]string{"payload"}}, RNG{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response from lib ", resp)

}
