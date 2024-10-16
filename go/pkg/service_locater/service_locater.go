package service_locater

import (
	"database/sql"
	"log"

	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/domain/repository"
	"github.com/cawauchi6204/hexagonal-architecture-todo/pkg/infra"
)

type ServiceLocater struct {
	Db *sql.DB

	// repositories
	userRepository repository.UserRepository
}

func BuildServiceLocater(env string) *ServiceLocater {
	var s *ServiceLocater
	switch env {
	case "local":
		s = buildLocalServiceLocater()
	case "test":
		s = buildTestServiceLocater()
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

func (s *ServiceLocater) registerRepositories() {
	db := s.Db

	s.userRepository = repository_impl.NewUserRepository(db)
}
