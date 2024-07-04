package main
import "fmt"

func main(){
	var lan string 
	var rate uint
	var Email string
	fmt.Print("\nFrom which language you used : ")
	fmt.Scan(&lan)
	fmt.Print("\nRate GO in scale of 10 : ")
	fmt.Scan(&rate)
	fmt.Print("\nYour Email : ")
	fmt.Scan(&Email)
	var lst [] string
	//lst[0]="....."
	lst=append(lst,lan + " " + " "+ Email)
	fmt.Printf("\n %v \n\tThanks for Your Time",lst)}
