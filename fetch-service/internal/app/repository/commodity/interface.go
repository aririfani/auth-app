package commodity

type Repository interface {
	FetchDataCommodity() (res []Commodity, err error)
}
