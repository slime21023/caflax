---
title: Caflax Features
---

# Caflax Features in Detail

Caflax offers a comprehensive set of features designed to make serving static content simple and efficient. This page provides detailed information about each feature.

## Core Features

### High Performance Static File Serving

Caflax is built using Go and the Iris web framework, which provides exceptional performance for static content serving. The implementation focuses on efficiency, making it suitable for serving everything from small websites to larger applications.

Key performance aspects:
- Efficient file handling with proper content type detection
- Optimized memory usage
- Fast request processing

### Single Page Application (SPA) Support

Modern web applications built with frameworks like React, Vue, or Angular often use client-side routing. Caflax's SPA mode ensures that all requests for non-existent files are redirected to `index.html`, allowing your SPA's router to handle the navigation.

Enable SPA mode with:
```bash
caflax --spa
```

When a request comes in for a path that doesn't match a physical file (e.g., `/dashboard`), Caflax will serve the `index.html` file instead, allowing your client-side application to handle the routing.

### Security Considerations

Caflax is built on top of the Iris web framework, which provides a solid foundation for serving content securely. While Caflax focuses on being lightweight and efficient, you should be aware of security considerations when serving content, especially in production environments.

For production deployments, it's recommended to place Caflax behind a reverse proxy like Nginx or Caddy, which can add additional security headers and provide HTTPS support.

### Cross-Origin Resource Sharing (CORS)

When enabled, Caflax adds the necessary CORS headers to allow cross-origin requests, which is essential for frontend applications that need to fetch resources from different domains.

Enable CORS with:
```bash
caflax --cors
```

This adds the following headers:
- `Access-Control-Allow-Origin: *`
- `Access-Control-Allow-Methods: GET, POST, OPTIONS, PUT, DELETE`
- `Access-Control-Allow-Headers: Origin, Content-Type, Accept`

The OPTIONS method is handled specially to support preflight requests, returning a 200 OK status.

### Compression

Caflax automatically compresses responses to reduce bandwidth usage and improve loading times. It supports both gzip and brotli compression algorithms and intelligently avoids compressing already compressed files like images.

Compression is enabled by default but can be controlled:
```bash
# Explicitly enable compression (default)
caflax --compress

# Disable compression
caflax --compress=false
```

The compression settings include:
- Smart detection of file types that benefit from compression
- Skipping compression for files smaller than 50 bytes
- No double-compression of already compressed formats

### Cache Control

For development purposes, you might want to disable browser caching to ensure you always see the latest changes. Caflax makes this easy with the `--no-cache` option.

```bash
caflax --no-cache
```

When enabled, this adds the following headers:
- `Cache-Control: no-store, no-cache, must-revalidate, proxy-revalidate`
- `Pragma: no-cache`
- `Expires: 0`

### Directory Listings

When accessing a directory without an index file, Caflax can display a clean directory listing, making it easy to navigate and access available files.

## Additional Features

### Custom Port and Host

Easily specify which port and hostname Caflax should use:

```bash
# Custom port
caflax -p 8080

# Custom host
caflax -H 0.0.0.0

# Both together
caflax -H 0.0.0.0 -p 8080
```

### Quiet Mode

Reduce log output for less verbose operation:

```bash
caflax --quiet
```

### Flexible Directory Selection

Serve files from any directory:

```bash
# Current directory
caflax

# Specific directory
caflax ./public

# Using the -d flag
caflax -d /path/to/files
```

### Version Information

Quickly check which version of Caflax you're running:

```bash
caflax --version
```

## HTTP/2 Support

Caflax leverages Iris's HTTP/2 capabilities, including Server Push for common web assets, providing better performance for compatible clients.
