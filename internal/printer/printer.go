package printer

import (
	"fmt"

	"github.com/fatih/color"
)



var (
	success = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	warning = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
	error   = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	info    = color.New(color.FgWhite).SprintFunc()
)

func Success(message string){
	fmt.Println(success(message))
}

func Warning(message string){
	fmt.Println(warning(message))
}

func Error(message string) {
	fmt.Println(error(message))
}

func Info(message string) {
	fmt.Println(info(message))
}


func Gradient(message string){
	//Color-gradient from blue-to-white (Kustom color)
	startRGB := [3]int{51,165,195}
	endRGB := [3]int{255, 255, 255} 

	for i, ch := range message {
		r := startRGB[0] + (endRGB[0]-startRGB[0])*i/len(message)
		g := startRGB[1] + (endRGB[1]-startRGB[1])*i/len(message)
		b := startRGB[2] + (endRGB[2]-startRGB[2])*i/len(message)
		fmt.Printf("\x1b[1;38;2;%d;%d;%dm%c", r, g, b, ch)
	}
	fmt.Print("\x1b[0m\n") 
}
