package cmd

import (
	"fmt"
	"os"

	"caflax/internal/server"

	"github.com/spf13/cobra"
)

var (
	port     int
	host     string
	dir      string
	spa      bool
	cors     bool
	noCache  bool
	compress bool
	quiet    bool
)

var rootCmd = &cobra.Command{
	Use:   "caflax [dir]",
	Short: "Simple static file server",
	Long: `Caflax is a lightweight static file server inspired by Vercel/Serve,
designed for quickly serving static websites or files locally.
Supports Single Page Applications (SPA), CORS, compression, and more.`,
	Example: `  caflax                     # Use current directory
  caflax dist                # Use specified directory
  caflax -p 8080 ./public   # Specify port and directory
  caflax --spa              # Enable SPA mode
  caflax --cors --compress  # Enable CORS and compression`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// If positional argument is provided, use it as directory
		if len(args) > 0 {
			dir = args[0]
		}

		serverOpts := server.Options{
			Port:     port,
			Host:     host,
			Dir:      dir,
			SPA:      spa,
			CORS:     cors,
			NoCache:  noCache,
			Compress: compress,
			Quiet:    quiet,
		}

		if err := server.Start(serverOpts); err != nil {
			fmt.Fprintln(os.Stderr, "Error starting server:", err)
			os.Exit(1)
		}
	},
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Set common options
	rootCmd.Flags().IntVarP(&port, "port", "p", 3000, "Server port")
	rootCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Server hostname")
	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Directory to serve")
	rootCmd.Flags().BoolVarP(&spa, "spa", "s", false, "Enable Single Page Application mode (returns index.html for unmatched requests)")
	rootCmd.Flags().BoolVarP(&cors, "cors", "c", false, "Enable Cross-Origin Resource Sharing (CORS)")
	rootCmd.Flags().BoolVarP(&noCache, "no-cache", "n", false, "Disable browser caching")
	rootCmd.Flags().BoolVarP(&compress, "compress", "z", true, "Enable gzip/brotli compression (enabled by default)")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Quiet mode, reduce log output")

	// Simplify command line usage template
	rootCmd.SetUsageTemplate(`Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Options:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Options:
  {{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)
}
