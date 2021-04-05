package stats

import (
	"github.com/Dovar001/bank/v2/pkg/types"
)

//Avg расчитивает средную сумму платежа

func Avg(payments []types.Payment) types.Money {

	var mid types.Money
	var a types.Money
	for _, operation := range payments {

		if !(operation.Status == types.StatusFail) {
			mid += operation.Amount
			a++
		}
	}
	mid = mid / a
	return mid
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {

	var totalincat types.Money
	for _, operation := range payments {

		if (operation.Category == category && !(operation.Status == types.StatusFail)) {
			totalincat += operation.Amount

		}

		 
	 }
	 return totalincat 

}


func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {


	result := map[types.Category]types.Money{}
	count := map[types.Category]types.Money {}
  
	  for _, payment := range payments {
  
		  result[payment.Category] +=payment.Amount
		  count[payment.Category] +=1
		  
	  }
  
	  for key := range result {
		  
		  result[key] /=count[key] 
		  
	  }
	  return result
  
}  



func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money,
   ) map[types.Category]types.Money {

	result := map[types.Category]types.Money{}

	for key := range second {
	
		result[key]+=second[key]
	}
		

	for key := range first {
	
		result[key]-=first[key]
	}
	return result
   }
   

