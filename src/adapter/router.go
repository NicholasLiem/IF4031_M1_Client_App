package adapter

import (
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/middleware"
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/routes"
	"github.com/NicholasLiem/IF4031_M1_Client_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/app"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(server app.MicroserviceServer) *mux.Router {

	router := mux.NewRouter()

	structs.AppRoutes = append(structs.AppRoutes,
		routes.UserRoutes(server),
		routes.AuthRoutes(server),
		routes.BookingRoutes(server),
	)

	for _, route := range structs.AppRoutes {

		//create sub route
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//for each sub route
		for _, subRoute := range route.SubRoutes {

			var handler http.Handler
			handler = subRoute.HandlerFunc

			if subRoute.Protected {
				handler = middleware.Middleware(subRoute.HandlerFunc) // use middleware
			}

			//register the route
			routePrefix.Path(subRoute.Pattern).Handler(handler).Methods(subRoute.Method).Name(subRoute.Name)
		}

	}

	return router
}
