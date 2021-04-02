package raindrops

import (
	"fmt"
	"os/exec"
)

func Script() {
	lim := 25
	for i:=1; i <= lim; i++{
		err:=exec.Command("say", Convert(i)).Run()
		fmt.Printf("error %v\n",err )
		exec.Command("sleep", "1").Run()
		fmt.Printf("error %v\n",err )

	}
}
