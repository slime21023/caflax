# Caflax

A lightweight static file server implemented in Golang with the Iris web framework, inspired by [Vercel/Serve](https://github.com/vercel/serve). Caflax is designed for quickly serving static websites or files locally.

## Features

- **Fast** - Built with Go and Iris framework for high-performance static content serving
- **SPA Mode** - Perfect support for Single Page Applications, automatically redirecting routes to index.html
- **Secure** - Automatically adds security-related headers to protect your application
- **CORS Support** - Optional Cross-Origin Resource Sharing
- **Compression** - Supports gzip and brotli compression to reduce transfer size
- **Cache Control** - Optional browser cache disabling for development
- **Directory Listing** - Support for displaying directory content listings
- **Easy to Use** - Intuitive command-line interface suitable for various development scenarios

## Usage

### Basic Usage

```bash
# Start the server in the current directory
caflax

# Specify directory to serve
caflax ./dist

# Specify port
caflax -p 8080

# Specify hostname and port
caflax -H 0.0.0.0 -p 8000 ./public
```

### Advanced Options

```bash
# Enable SPA mode (all non-existent paths will be redirected to index.html)
caflax --spa

# Enable CORS
caflax --cors

# Disable browser caching
caflax --no-cache

# Enable compression (enabled by default)
caflax --compress

# Quiet mode, reduce log output
caflax --quiet

# Show version information
caflax --version
```

### Complete Help Information

Run `caflax --help` to view the complete command-line options.

## Use Cases

### Development Environment
- Local static website development
- Frontend project showcasing
- Testing single page applications
- Temporary sharing of static files or documents
- Quickly starting a local server

### Production Environment Considerations

Caflax can be used in the following production environments, but with some limitations:

**Suitable Scenarios**:
- Static file serving for small internal websites or tools
- Static content websites with small to medium traffic
- Static resource distribution nodes combined with CDN
- Hosting company internal documents or knowledge bases

**Usage Recommendations**:
- In production environments, it's recommended to configure Nginx/Caddy or other reverse proxies in front of Caflax to provide HTTPS support
- For high-traffic websites, consider using professional static file hosting services
- When dealing with sensitive data, additional authentication and authorization layers should be added
- Use monitoring tools to track server performance and status