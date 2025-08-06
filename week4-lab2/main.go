package main
import "fmt"
//moodeng
func main() {
	//var name string = "Sippakorn"
	var age int = 20

	email := "butkot_s@su.ac.th"
	gpa := 3.5

	firstname, lastname := "Sippakorn", "Butkot"
	fmt.Printf("Name %s %s,age %d,email %s,gpa %.2f\n", firstname, lastname, age, email, gpa)
}