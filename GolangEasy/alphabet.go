package GolangEasy

import (
	"math/rand"
	"time"
)
//all alphabet
var UpLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var DownLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var AllLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// n - length of the word
// if up = "none" -> the tpe doesn't matter and var AllLetters
// if tpe = "up" -> var UpLetters
// if tpe = "down" -> var DownLetters
func RandomWord(n int, tpe string) string {
	b := make([]rune, n)
	if tpe == "none"{
    	for i := range b {
        b[i] = AllLetters[rand.Intn(len(AllLetters))]
		}
	}else if tpe == "up"{
    	for i := range b {
        b[i] = UpLetters[rand.Intn(len(UpLetters))]
		}
	}else if tpe == "down"{
    	for i := range b {
        b[i] = DownLetters[rand.Intn(len(DownLetters))]
		}
	}
    return string(b)
}
func main(){
	rand.Seed(time.Now().UnixNano())
}
