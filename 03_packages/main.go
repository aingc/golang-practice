package main

// If multiple packages, then separate by white space
// If saved without using the package, then package import will disappear
import (
	"fmt"
	"math"

	"github.com/aingc/golang-practice/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}