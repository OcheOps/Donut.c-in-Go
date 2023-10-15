package main


import (
	"fmt"
	"math"
	//"os"
	"time"
)

func main(){
	for i := 0; i < 360; i++ {
		degree := float64(i)
        radian := degree * (math.Pi / 180.0)
        z := math.Sin(radian)
        z2 := math.Sin(radian / 2)
        for j := 0; j < 360; j += 2 {
            degree2 := float64(j)
            radian2 := degree2 * (math.Pi / 180.0)
            x := int(40 + (z + z2*math.Sin(radian2)) * 10)
            y := int(12 + z2*math.Cos(radian2) * 20)
            if x >= 0 && x < 80 && y >= 0 && y < 24 {
                fmt.Printf("\033[%d;%dH*", y+1, x+1)
            }
        }
        time.Sleep(100 * time.Millisecond) // Sleep for animation effect
        fmt.Print("\033[2J") // Clear screen
	}
}

