package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/lloistborn/stockinvent/inventory"
)

// Routes object.
type Routes struct {
	Inventory *inventory.InventoryModule
}

// Route return all api endpoint from object Routes.
func Route() *Routes {
	return &Routes{
		Inventory: inventory.NewInventoryModule(),
	}
}

// RegisterAPI will register all available API Endpoint.
func (routes *Routes) RegisterAPI(r *httprouter.Router) {
	// Blouse REST endpoint.
	r.POST("/items/blouses", routes.Inventory.InsertBlouse)
	r.GET("/items/blouses", routes.Inventory.GetBlouses)
	r.GET("/items/blouses/:id", routes.Inventory.GetBlouse)
	r.PUT("/items/blouses/:id", routes.Inventory.UpdateBlouse)
	r.DELETE("/items/blouses/:id", routes.Inventory.DeleteBlouse)

	r.POST("/items/blouses/inbound", routes.Inventory.InsertInboundBlouse)
}
