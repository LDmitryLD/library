package modules

import "projects/LDmitryLD/library/app/internal/modules/library/controller"

type Controllers struct {
	Library controller.Libraryer
}

func NewControllers(services *Services) *Controllers {
	libraryController := controller.NewLibraryController(services.Library)

	return &Controllers{
		Library: libraryController,
	}
}
