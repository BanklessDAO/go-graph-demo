package interfaces

import (
	"go-graph-demo/app/infrastructure/persistence"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repositories persistence.Repositories
}
