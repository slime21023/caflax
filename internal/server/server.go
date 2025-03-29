package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kataras/iris/v12"
)

// Options contains server configuration options
type Options struct {
	Port     int    // Server port
	Host     string // Host name
	Dir      string // Static file directory
	SPA      bool   // Single Page Application mode
	CORS     bool   // Enable CORS
	NoCache  bool   // Disable cache
	Compress bool   // Enable compression
	Quiet    bool   // Quiet mode
}

// CORS middleware
func corsMiddleware(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	if ctx.Method() == "OPTIONS" {
		ctx.StatusCode(http.StatusOK)
		return
	}

	ctx.Next()
}

// noCacheMiddleware disables browser caching
func noCacheMiddleware(ctx iris.Context) {
	ctx.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Next()
}

// Start launches the static file server
func Start(opts Options) error {
	// Create Iris application
	app := iris.New()

	// Configure logging
	if opts.Quiet {
		app.Logger().SetLevel("error")
	} else {
		app.Logger().SetLevel("info")
	}

	// Get absolute path
	absPath, err := filepath.Abs(opts.Dir)
	if err != nil {
		return fmt.Errorf("unable to get absolute path: %v", err)
	}

	// Check if directory exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", absPath)
	}

	if !opts.Quiet {
		app.Logger().Infof("Serving directory: %s", absPath)
	}

	// Enable CORS
	if opts.CORS {
		app.UseRouter(corsMiddleware)
		if !opts.Quiet {
			app.Logger().Info("CORS enabled")
		}
	}

	// Enable compression
	if opts.Compress {
		app.Use(iris.Compression)
		if !opts.Quiet {
			app.Logger().Info("Compression enabled")
		}
	}

	// Disable cache
	if opts.NoCache {
		app.Use(noCacheMiddleware)
		if !opts.Quiet {
			app.Logger().Info("Caching disabled")
		}
	}

	// Configure DirOptions (using sensible defaults)
	dirOptions := iris.DirOptions{
		IndexName: "index.html",
		ShowList:  true, // Show directory listing by default
		SPA:       opts.SPA,
	}

	// Configure Cache options (enabled by default for better performance)
	dirOptions.Cache = iris.DirCacheOptions{
		Enable:          true,
		CompressIgnore:  iris.MatchImagesAssets, // Don't compress images
		Encodings:       []string{"gzip", "br"}, // Most common compression formats
		CompressMinSize: 50,                     // Don't compress files smaller than 50 bytes
		Verbose:         0,                      // Don't show detailed logs by default
	}

	// Configure PushTargetsRegexp (enable HTTP/2 Server Push for common file types)
	pushTargetsRegexp := make(map[string]*regexp.Regexp)
	pushTargetsRegexp["/"] = iris.MatchCommonAssets
	dirOptions.PushTargetsRegexp = pushTargetsRegexp

	// Add custom 404 handler
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		path := ctx.Path()

		// In SPA mode, return index.html
		if opts.SPA && !strings.HasPrefix(path, "/api") && !strings.Contains(path, ".") {
			ctx.ServeFile(filepath.Join(absPath, "index.html"))
			return
		}

		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"error": fmt.Sprintf("Path not found: %s", path),
		})
	})

	// Serve static files
	app.HandleDir("/", absPath, dirOptions)

	// Start server
	addr := fmt.Sprintf("%s:%d", opts.Host, opts.Port)

	// Print access information
	if !opts.Quiet {
		app.Logger().Infof("Server started at http://%s", addr)
		app.Logger().Infof("Available at:")
		if opts.Host == "localhost" || opts.Host == "127.0.0.1" {
			app.Logger().Infof("- Local: http://localhost:%d", opts.Port)
		} else {
			app.Logger().Infof("- Local: http://%s:%d", opts.Host, opts.Port)
		}
		app.Logger().Infof("- Network: http://0.0.0.0:%d", opts.Port)

		if opts.SPA {
			app.Logger().Info("SPA mode enabled")
		}
	}

	return app.Listen(addr)
}
