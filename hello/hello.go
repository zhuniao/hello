package hello

import "fmt"

func SayHello(language string) {
	switch language {
	case "cn":
		fmt.Println("求知若渴，求贤若愚。")
	case "en":
		fmt.Println("Stay hungry,stay foolish.")
	default:
		fmt.Println("only support cn and en")
	}
}
