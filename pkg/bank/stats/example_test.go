package stats

import (
	"fmt"

	"github.com/mrGreatProgrammer/tenlecture/pkg/bank/types"
)

func ExampleAvg() {

	var paymetns = []types.Payment{
		{
			ID:       1,
			Amount:   200,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   100,
			Category: "auto",
		},
		{
			ID:       3,
			Amount:   300,
			Category: "auto",
		},
		{
			ID:       4,
			Amount:   500,
			Category: "auto",
		},
		{
			ID:       5,
			Amount:   400,
			Category: "auto",
		},
	}

	fmt.Println(Avg(paymetns))

	//Output:
	// 300

}