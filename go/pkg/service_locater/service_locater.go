package service_locater

import (
	"database/sql"
	"log"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/application/usecase"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra/repository_impl"
)

type ServiceLocater struct {
	Db *sql.DB

	// repositories
	UserRepository repository.UserRepository

	// usecases
	UserUsecase *usecase.UserUsecase
}

func BuildServiceLocater(env string) *ServiceLocater {
	var s *ServiceLocater
	switch env {
	case "local":
		s = buildLocalServiceLocater().registerRepositories().registerUsecases()
	case "test":
		s = buildTestServiceLocater().registerRepositories().registerUsecases()
	default:
		log.Fatalf("can not build service locater because env is unknown one: %s", env)
		return nil
	}
	return s
}

func buildLocalServiceLocater() *ServiceLocater {
	db := infra.InitDB()
	return &ServiceLocater{Db: db}
}

func buildTestServiceLocater() *ServiceLocater {
	db := infra.InitDB()
	return &ServiceLocater{Db: db}
}

func (s *ServiceLocater) registerRepositories() *ServiceLocater {
	db := s.Db

	s.UserRepository = repository_impl.NewUserRepository(db)

	return s
}

func (s *ServiceLocater) registerUsecases() *ServiceLocater {
	s.UserUsecase = &usecase.UserUsecase{
		UserRepository: s.UserRepository,
	}
	return s
}
