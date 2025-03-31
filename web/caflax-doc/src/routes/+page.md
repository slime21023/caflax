---
title: Caflax
heroImage: /caflax-logo.png
tagline: A lightweight static file server implemented in Golang with the Iris web framework
actions:
  - label: Get Started
    type: primary
    to: /guide/
  - label: GitHub
    type: secondary
    to: https://github.com/yourname/caflax
    external: true
features:
  - title: Fast & Efficient
    description: Built with Go and Iris framework for high-performance static content serving
  - title: SPA Mode
    description: Perfect support for Single Page Applications, automatically redirecting routes to index.html
  - title: Security First
    description: Automatically adds security-related headers to protect your application
  - title: Developer-Friendly
    description: Intuitive CLI, CORS support, compression, and cache controls for seamless development
---

# Caflax

A lightweight static file server implemented in Golang with the Iris web framework, inspired by [Vercel/Serve](https://github.com/vercel/serve). Caflax is designed for quickly serving static websites or files locally.

## Basic Usage

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

## Advanced Options

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

## Features

- **Fast** - Built with Go and Iris framework for high-performance static content serving
- **SPA Mode** - Perfect support for Single Page Applications, automatically redirecting routes to index.html
- **Secure** - Automatically adds security-related headers to protect your application
- **CORS Support** - Optional Cross-Origin Resource Sharing
- **Compression** - Supports gzip and brotli compression to reduce transfer size
- **Cache Control** - Optional browser cache disabling for development
- **Directory Listing** - Support for displaying directory content listings
- **Easy to Use** - Intuitive command-line interface suitable for various development scenarios
