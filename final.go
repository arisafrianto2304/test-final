package main
import "fmt"

func main(){
	var x = []int{12,444,54643,3155,667543,8637,0,369,7516,335}
	// var y int = 4
	for i := 0; i < len(x); i++{
			fmt.Printf("\n+======+\n")
			fmt.Printf("|")
			fmt.Print(bintang(x[i]),x[i])
			fmt.Printf("|")
			fmt.Printf("\n+======+")
	}
}

// find mount digit
func jumlahDigit(a int)int{
	var count int = 0
	if a == 0{
		a = 1
	}
	for a > 0{
		a = a/10
		count++
	}
	return count
}

// digit < 6, left '*'
func bintang(a int) string{
	var b int = jumlahDigit(a)
	var tempa int
	var hasil string
	for i:=0; i < b; i++ {
		tempa = 6 - b	
	}
	for j:=0; j < tempa; j++{
		hasil += "*"
	}
	return hasil
}
