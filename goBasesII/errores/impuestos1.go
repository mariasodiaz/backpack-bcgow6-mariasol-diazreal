package main

import ("fmt"
 "os")

type SalaryError struct{
	status int
	message string
}

func (e *SalaryError) Error() string{
	return fmt.Sprintf("%v - %v\n",e.status,e.message)
}

func AplicaImpuesto (salary int) (int,error){
	if salary > 150000{
		return 200, nil
	} 
	return 500,&SalaryError{
		status:500,
		message: "error, el salario ingresado no alcanza el minimo imponible",
	}
}

func main(){

	salary := 200000
	_,error := AplicaImpuesto(salary)
	if error != nil{
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
	
}