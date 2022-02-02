package commodity

import "github.com/aririfani/auth-app/fetch-service/internal/app/repository/commodity"

type Service interface {
	FetchDataCommodity() (returnData []commodity.Commodity, err error)
	FetchDataCommodityByProvince(provinceName string, sorted bool, sortedType string) (returnData []commodity.Commodity, err error)
}
