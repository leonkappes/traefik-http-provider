package routes

import (
	"encoding/json"
	"log"

	jsonUtils "github.com/leonkappes/go-traefik-daemon/pkg/json"
	"github.com/leonkappes/go-traefik-daemon/pkg/types"
	"github.com/valyala/fasthttp"
)

func ProviderHandler(ctx *fasthttp.RequestCtx) {
	//Routers
	routersJson := jsonUtils.ReadJsonFile("configs/routers.json")
	var routers map[string]types.Router
	if err := json.Unmarshal(routersJson, &routers); err != nil {
		log.Printf("Error parsing json: %s", err)
	}

	//Middlewares
	middlewaresJson := jsonUtils.ReadJsonFile("configs/middlewares.json")
	var middlewares map[string]types.Middleware
	if err := json.Unmarshal(middlewaresJson, &middlewares); err != nil {
		log.Printf("Error parsing json: %s", err)
	}

	//Services
	serviceJson := jsonUtils.ReadJsonFile("configs/services.json")
	var services map[string]types.Service
	if err := json.Unmarshal(serviceJson, &services); err != nil {
		log.Printf("Error parsing json: %s", err)
	}

	jsonUtils.ValueChanged = false

	httpResponse := &types.Http{
		Services:    services,
		Middlewares: middlewares,
		Routers:     routers,
	}

	data, _ := json.Marshal(httpResponse)
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Write(data)
}
