package handlers

import "github.com/cawauchi6204/hexagonal-architecture-todo/pkg/service_locater"

type BaseHandler struct {
	ServiceLocater *service_locater.ServiceLocater
}
