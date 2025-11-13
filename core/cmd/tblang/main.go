package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tblang/core/internal/engine"
)

var (
	// Color definitions
	successColor = color.New(color.FgGreen, color.Bold)
	errorColor   = color.New(color.FgRed, color.Bold)
	warningColor = color.New(color.FgYellow, color.Bold)
	infoColor    = color.New(color.FgCyan, color.Bold)
	headerColor  = color.New(color.FgMagenta, color.Bold)
)

var rootCmd = &cobra.Command{
	Use:   "tblang",
	Short: "TBLang - Infrastructure as Code Language",
	Long: `TBLang is a domain-specific language for Infrastructure as Code.
It provides a simple, readable syntax for managing cloud infrastructure
with a plugin-based architecture supporting multiple cloud providers.`,
	Version: "1.1.1",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			printCredits()
		}
	},
}

var planCmd = &cobra.Command{
	Use:   "plan [file.tbl]",
	Short: "Show what infrastructure changes will be made",
	Long:  `Analyze the TBLang configuration file and show what resources will be created, updated, or destroyed.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runWithEngine(func(ctx context.Context, engine *engine.Engine) error {
			infoColor.Println("Planning infrastructure changes...")
			return engine.Plan(ctx, args[0])
		})
	},
}

var applyCmd = &cobra.Command{
	Use:   "apply [file.tbl]",
	Short: "Apply infrastructure changes",
	Long:  `Create, update, or delete infrastructure resources as defined in the TBLang configuration file.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runWithEngine(func(ctx context.Context, engine *engine.Engine) error {
			infoColor.Println("Applying infrastructure changes...")
			return engine.Apply(ctx, args[0])
		})
	},
}

var destroyCmd = &cobra.Command{
	Use:   "destroy [file.tbl]",
	Short: "Destroy infrastructure",
	Long:  `Destroy all infrastructure resources defined in the TBLang configuration file.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runWithEngine(func(ctx context.Context, engine *engine.Engine) error {
			warningColor.Println("Destroying infrastructure...")
			return engine.Destroy(ctx, args[0])
		})
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current infrastructure state",
	Long:  `Display the current state of managed infrastructure resources.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runWithEngine(func(ctx context.Context, engine *engine.Engine) error {
			infoColor.Println("Current infrastructure state:")
			return engine.Show()
		})
	},
}

var graphCmd = &cobra.Command{
	Use:   "graph [file.tbl]",
	Short: "Show dependency graph",
	Long:  `Display the dependency graph and deployment order for resources.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runWithEngine(func(ctx context.Context, engine *engine.Engine) error {
			infoColor.Println("Analyzing dependency graph...")
			return engine.Graph(ctx, args[0])
		})
	},
}

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Plugin management commands",
	Long:  `Manage TBLang provider plugins.`,
}

var pluginsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available plugins",
	Long:  `List all available provider plugins.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runWithEngine(func(ctx context.Context, engine *engine.Engine) error {
			plugins := engine.ListPlugins()
			if len(plugins) == 0 {
				warningColor.Println("No plugins found. Install plugins in .tblang/plugins/ directory.")
			} else {
				headerColor.Println("Available plugins:")
				for _, plugin := range plugins {
					successColor.Printf("  %s\n", plugin)
				}
			}
			return nil
		})
	},
}

func printCredits() {
	cyan := color.New(color.FgCyan, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)
	yellow := color.New(color.FgYellow)
	
	cyan.Println("\n╔════════════════════════════════════════════════════════╗")
	magenta.Println("║              Developed with ❤️  by                     ║")
	yellow.Println("║                                                        ║")
	successColor.Println("║           🚀 Swan Htet Aung Phyo                       ║")
	successColor.Println("║           🚀 Aung Zayar Moe                            ║")
	yellow.Println("║                                                        ║")
	cyan.Println("╚════════════════════════════════════════════════════════╝")
}

func init() {
	// Custom version template
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
{{- with .Short}}
{{.}}{{end}}

` + getCreditsString())
	
	// Add subcommands
	rootCmd.AddCommand(planCmd)
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(destroyCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(graphCmd)
	rootCmd.AddCommand(pluginsCmd)
	
	// Add plugin subcommands
	pluginsCmd.AddCommand(pluginsListCmd)
	
	// Global flags
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable colored output")
}

func getCreditsString() string {
	cyan := color.New(color.FgCyan, color.Bold)
	magenta := color.New(color.FgMagenta, color.Bold)
	green := color.New(color.FgGreen, color.Bold)
	
	return cyan.Sprint("╔════════════════════════════════════════════════════════╗\n") +
		magenta.Sprint("║              Developed with ❤️  by                     ║\n") +
		cyan.Sprint("║                                                        ║\n") +
		green.Sprint("║           🚀 Swan Htet Aung Phyo                       ║\n") +
		green.Sprint("║           🚀 Aung Zayar Moe                            ║\n") +
		cyan.Sprint("║                                                        ║\n") +
		cyan.Sprint("╚════════════════════════════════════════════════════════╝")
}

func runWithEngine(fn func(context.Context, *engine.Engine) error) error {
	// Create context that cancels on interrupt
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		warningColor.Println("\nReceived interrupt signal, shutting down...")
		cancel()
	}()

	// Check for no-color flag
	if noColor, _ := rootCmd.PersistentFlags().GetBool("no-color"); noColor {
		color.NoColor = true
	}

	// Initialize TBLang engine
	tblangEngine := engine.New()
	defer tblangEngine.Shutdown()

	if err := tblangEngine.Initialize(ctx); err != nil {
		errorColor.Printf("Failed to initialize engine: %v\n", err)
		return err
	}

	return fn(ctx, tblangEngine)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		errorColor.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}