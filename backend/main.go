package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/pcn/clutch-custom-gateway/backend/cmd/assets"
	"github.com/pcn/clutch-custom-gateway/backend/module/echo"
	amiibomod "github.com/pcn/clutch/backend/module/amiibo"
	amiiboservice "github.com/pcn/clutch/backend/service/amiibo"
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
