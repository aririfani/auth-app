package helper

import (
	"github.com/aririfani/auth-app/fetch-service/internal/app/repository/commodity"
	"sort"
)

func SortCommodityByPrice(commodity []commodity.Commodity) {
	sort.Slice(commodity, func(i, j int) bool {
		return commodity[i].PriceUSD > commodity[j].PriceUSD
	})
}
