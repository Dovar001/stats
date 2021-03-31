package stats

import (
	"github.com/Dovar001/bank/pkg/types"
	
)

//Avg расчитивает средную сумму платежа

func Avg(payments [] types.Payment)  types.Money{

	var mid types.Money
	var a types.Money
	for _, operation := range payments {

		mid+=operation.Amount
		a++

	}
	mid=mid/a
	return mid
}

func TotalInCategory(payments []types.Payment , category types.Category) types.Money{

       var totalincat types.Money	
     for _, operation := range payments {
		 
		if (operation.Category == category){
			totalincat +=operation.Amount

		}
		 
	 }
	 return totalincat 
}