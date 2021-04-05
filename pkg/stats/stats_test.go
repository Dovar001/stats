package stats

import (
	
	"reflect"
	"testing"

	"github.com/Dovar001/bank/v2/pkg/types"

)

func TestCategoriesAvg(t *testing.T) {

	payments := []types.Payment{

		{
			Category: "Машина",
			Amount: 10_000,
		},
		{
			Category: "Машина",
			Amount: 30_000,
		},
		{
			Category: "Магазин",
			Amount: 50_000,
		},
		{
			Category: "Ресторан",
			Amount: 80_000,
		},
		{
			Category: "Магазин",
			Amount: 90_000,
		},
	}
	expect :=  map[types.Category]types.Money{
		"Машина": 20_000,
		"Ресторан": 80_000,
		"Магазин" : 70_000,

	}
	

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expect,result){
		t.Errorf("invalid result, expected: %v, actual:%v", expect , result)
	}
	

}

func TestPeriodsDynamic(t *testing.T) {

	first := map[types.Category]types.Money{

			"Машина": 10_000,
		    "Ресторан": 20_000,
	
	}
	second := map[types.Category]types.Money{

		"Машина": 10_000,
		"Ресторан":10_000,
		
}
	expect :=  map[types.Category]types.Money{
		"Машина": 0_000,
		"Ресторан": -10_000,
		
	}
	

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expect,result){
		t.Errorf("invalid result, expected: %v, actual:%v", expect , result)
	}
	

}