package main

import "fmt"
import "strconv"
func main(){
	fmt.Println("Enter your card number")
	var cardNumber string	
	fmt.Scanln(&cardNumber)
	cardObject:=card{cardNumber:cardNumber}
	fmt.Println("this is your card number :",cardNumber)
	fmt.Println(cardObject.isValid())
}

type card struct {
    cardNumber string
}
func(card *card) isValid() bool{
	sum:=0
	index:=1
	val:=0
	for i := len(card.cardNumber)-1; i >=0 ; i-- {
		y,err := strconv.Atoi(string(card.cardNumber[i]))	
				if(err!=nil){
					return false
					}
		 if(index%2==0){
			val=y*2
		}else{
			val=y
		}		
		
		s:=strconv.Itoa(val)
for z := len(s)-1; z >=0 ; z-- {
	b,err:=strconv.Atoi(string(s[z]))	

	if(err!=nil){
					return false
					}
					sum=sum+b
}
	fmt.Println(sum)
	index=index+1
 }
if(sum%10==0){
	return true
}
 return false
}

