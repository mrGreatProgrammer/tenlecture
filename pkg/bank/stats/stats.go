package stats

import (
	"github.com/mrGreatProgrammer/tenlecture/pkg/bank/types"
)

var theMoney types.Money

func Avg(paymetns []types.Payment) types.Money {
	for _, v := range paymetns {
		theMoney += v.Amount
	}

	calcIt := int(theMoney) / len(paymetns)

	return types.Money(calcIt)
}