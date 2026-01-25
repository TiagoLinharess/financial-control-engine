package repositories

import pgs "financialcontrol/internal/store/pgstore"

type Repository struct {
	store *pgs.Queries
}

func NewRepository(store *pgs.Queries) Repository {
	return Repository{store: store}
}
