package stats

import (
	"github.com/Dovar001/bank/v2/pkg/types"
	"fmt"

	
)

func ExampleAvg() {

	cards := []types.Payment{

		{
		ID: 4444,
		Amount: 10_000,
		Category: "Машина",
		Status: "FAIL",
		},
		{
		ID: 3333,
		Amount: 100_000,
		Category: "Ресторан",
		Status: "OK",
		},
		
		{
		ID: 1111,
		Amount: 30_000,
		Category: "Магазин",
		Status: "OK",
		},
		
	}

	 mid := Avg(cards)
	 fmt.Println(mid)

	//Output:65000

	
}

func ExampleTotalInCategory() {

	cards := []types.Payment{

		{
		ID: 4444,
		Amount: 10_000,
		Category: "Машина",
		Status: "FAIL",
		},
		{
		ID: 3333,
		Amount: 30_000,
		Category: "Ресторан",
		},
		{
			ID: 2222,
			Amount: 50_000,
			Category: "Машина",
			},
		
		{
		ID: 1111,
		Amount: 30_000,
		Category: "Магазин",
		},
		{
		ID: 0000,
		Amount: 100_000,
		Category: "Ресторан",
		},
		
	};
	var cat types.Category
	cat = "Машина"

	 mid := TotalInCategory(cards,cat)
	 fmt.Println(mid)

	//Output:50000
	

	
}
