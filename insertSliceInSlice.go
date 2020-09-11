package main
import "fmt"

func insertSlice(s, insertion []string, index int) []string {
	// iniate new arr
	res := make([]string, len(s)+len(insertion))
	// get position for inserting in res
	temp := copy(res, s[:index])
	// copy value of insertion to res in position temp
	temp += copy(res[temp:], insertion)
	// copy value from str to res in position
	copy(res[temp:], s[index:])
	return res
}

func main(){
	str := []string{"M", "N", "O", "P", "Q", "R"}
	strinsert := []string{"d", "e", "f"}
	res := insertSlice(str, strinsert, 0)
	fmt.Println(res)
	res = insertSlice(str, strinsert, 1)
	fmt.Println(res)
}