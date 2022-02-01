package commodity

import (
	"github.com/aririfani/auth-app/fetch-service/internal/app/repository"
	"github.com/aririfani/auth-app/fetch-service/internal/app/repository/commodity"
)

type srv struct {
	Repo repository.Repositories
}

func NewSrv(repo repository.Repositories) (returnData Service) {
	return &srv{
		Repo: repo,
	}
}

func (s *srv) FetchDataCommodity() (res []commodity.Commodity, err error) {
	res, err = s.Repo.Commodity().FetchDataCommodity()
	if err != nil {
		return
	}

	return
}
