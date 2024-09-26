package port

type FavouriteRepositoryItf interface {
	NewFavouriteRepositoryClient(tx bool) FavouriteRepositoryClientItf
}

type FavouriteRepositoryClientItf interface {
	Commit() error
	Rollback() error
}

type FavouriteServiceItf interface {
}
