package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/pcn/clutch-custom-gateway/backend/cmd/assets"
	amiibomod "github.com/pcn/clutch-custom-gateway/backend/module/amiibo"
	"github.com/pcn/clutch-custom-gateway/backend/module/echo"
	amiiboservice "github.com/pcn/clutch-custom-gateway/backend/service/amiibo"
)

func main() {
	flags := gateway.ParseFlags()

	components := gateway.CoreComponentFactory

	// Add custom components.
	components.Modules[echo.Name] = echo.New
	components.Modules[amiibomod.Name] = amiibomod.New
	components.Servcies[amiiboservice.Name] = amiiboservice.New

	gateway.Run(flags, components, assets.VirtualFS)
}
