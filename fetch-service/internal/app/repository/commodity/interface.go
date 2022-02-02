package commodity

type Repository interface {
	FetchDataCommodity() (res []Commodity, err error)
	FetchDataCommodityByProvince(provinceName string, sorted bool, sortedType string) (res []Commodity, err error)
}
