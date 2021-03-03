package registry

import "github.com/yoskeoka/go-example/mock/domain"

// ServiceRegistryInterface initializes a service with ensureing dependent on others.
type ServiceRegistryInterface interface {
	User() domain.User
	UserGroup() domain.UserGroup
}

// ServiceRegistry implements service registry interface.
type ServiceRegistry struct {
	// dependencies here...
}

// User returns User service.
func (sr *ServiceRegistry) User() domain.User {
	// initializes User service using it's dependencies.
	return nil
}
