package commands

import (
	"encoding/json"
	"fmt"
	"log"

	jsonUtils "github.com/leonkappes/go-traefik-daemon/pkg/json"
	"github.com/leonkappes/go-traefik-daemon/pkg/types"
	"github.com/urfave/cli/v2"
)

func ServiceInfo(c *cli.Context) error {
	serviceJson := jsonUtils.ReadJsonFile("configs/services.json")
	var services map[string]types.Service
	if err := json.Unmarshal(serviceJson, &services); err != nil {
		log.Printf("Error parsing json: %s", err)
		return err
	}
	serviceName := c.Args().First()
	fmt.Printf("Service: %s\nServers:\n", serviceName)
	for _, value := range services[serviceName].LoadBalancer.Servers {
		fmt.Printf("\t%s\n", value.Url)
	}
	return nil
}
