package engine

import (
	"context"
	"fmt"
	"os"

	"github.com/tblang/core/internal/compiler"
)

func (e *Engine) loadRequiredPlugins(ctx context.Context, program *compiler.Program) error {

	for providerName, config := range program.CloudVendors {
		infoColor.Printf("Found provider: %s\n", providerName)
		fmt.Printf("  Region: %v\n", config.Properties["region"])

		if profile, exists := config.Properties["profile"]; exists {
			if profileStr, ok := profile.(string); ok {
				os.Setenv("AWS_PROFILE", profileStr)
				infoColor.Printf("  Profile: %s\n", profileStr)
			}
		}

		if accountID, exists := config.Properties["account_id"]; exists {
			fmt.Printf("  Account ID: %v\n", accountID)
		}

		successColor.Printf("Provider %s configured (mock mode)\n", providerName)
	}

	return nil
}

func (e *Engine) loadAndConfigurePlugins(ctx context.Context, program *compiler.Program) error {
	for providerName, config := range program.CloudVendors {
		infoColor.Printf("Found provider: %s\n", providerName)
		fmt.Printf("  Region: %v\n", config.Properties["region"])

		if profile, exists := config.Properties["profile"]; exists {
			if profileStr, ok := profile.(string); ok {
				os.Setenv("AWS_PROFILE", profileStr)
				infoColor.Printf("  Profile: %s\n", profileStr)
			}
		}

		if accountID, exists := config.Properties["account_id"]; exists {
			fmt.Printf("  Account ID: %v\n", accountID)
		}

		_, err := e.pluginManager.LoadPlugin(ctx, providerName)
		if err != nil {
			return fmt.Errorf("failed to load plugin %s: %w", providerName, err)
		}

		if err := e.pluginManager.ConfigurePlugin(ctx, providerName, config.Properties); err != nil {
			return fmt.Errorf("failed to configure plugin %s: %w", providerName, err)
		}

		successColor.Printf("Provider %s loaded and configured\n", providerName)
	}

	return nil
}
