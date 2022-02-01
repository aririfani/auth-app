package commodity

import "github.com/aririfani/auth-app/fetch-service/internal/app/repository/commodity"

type Service interface {
	FetchDataCommodity() (returnData []commodity.Commodity, err error)
}
