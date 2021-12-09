package factory

import "github.com/LuannMateus/fc-fs-5-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
