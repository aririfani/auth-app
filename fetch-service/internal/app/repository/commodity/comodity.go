package commodity

type repo struct {
	DB         Repository
	httpClient *httpClient
}

func NewRepo(db Repository) (returnData Repository) {
	return &repo{
		DB: db,
	}
}

func (r *repo) FetchDataCommodity() (returnData []Commodity, err error) {
	returnData, err = r.DB.FetchDataCommodity()
	if err != nil {
		return
	}
	return
}
