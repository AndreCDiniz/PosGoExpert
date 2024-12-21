package main

//Using locally workspaces
//It's important to use go mod tidy -e if u use a local package and online packages
import (
	"github.com/AndreCDiniz/PosGoExpert/5-Packaging/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
