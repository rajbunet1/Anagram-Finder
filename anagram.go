
package main
import (
	"fmt"
	"strings"
)

func main(){
	for{
	var word string
	fmt.Println("Enter a word")
	fmt.Scan(&word)
	var word2 string
	fmt.Println("Enter a second word")
	fmt.Scan(&word2)
		if len(word) == len(word2){
			for i:=0;i<len(word);i++{	
				if !strings.Contains(strings.ToLower(word2), string(strings.ToLower(word)[i])){
					fmt.Println("not an anagrammer")
					break	
				}else if i ==len(word)-1 && strings.Contains(strings.ToLower(word2), string(strings.ToLower(word)[i])){
					println(i)
					fmt.Println("This is an anagram!") 
				}
			}
		}
	}
}
