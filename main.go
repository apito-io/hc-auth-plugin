package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-auth-plugin] Starting Authentication plugin...")

	plugin := sdk.Init("hc-auth-plugin", "1.0.0", "apito-plugin-key")

	// Stub query - replace with actual auth implementation
	statusType := sdk.NewObjectType("AuthStatus", "Authentication plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()

	plugin.RegisterQuery("getAuthStatus",
		sdk.ComplexObjectField("Get authentication plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{
				"status":  "ready",
				"version": "1.0.0",
			}, nil
		})

	log.Printf("🚀 [hc-auth-plugin] Plugin ready")
	plugin.Serve()
}
