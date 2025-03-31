---
title: Getting Started with Caflax
---

# Getting Started with Caflax

Caflax is a lightweight static file server that makes it easy to serve your web projects locally. This guide will help you get started with basic and advanced usage.

## Basic Usage

The most basic way to use Caflax is to run it in the directory you want to serve:

```bash
caflax
```

This will start a server on the default port (3000) serving files from the current directory. You can access your files by navigating to `http://localhost:3000` in your browser.

### Serving Specific Directories

To serve files from a specific directory:

```bash
caflax ./dist
```

Or with the `-d` flag:

```bash
caflax -d ./public
```

### Customizing Port and Host

You can specify a different port using the `-p` or `--port` flag:

```bash
caflax -p 8080
```

To make the server accessible from other devices on your network, use the `-H` flag to set the host:

```bash
caflax -H 0.0.0.0
```

Combining port and host:

```bash
caflax -H 0.0.0.0 -p 8000 ./public
```

## Advanced Features

### SPA Mode

If you're working with a Single Page Application (React, Vue, Angular, etc.), you can enable SPA mode:

```bash
caflax --spa
```

This ensures that all requests for non-existent files will be redirected to `index.html`, allowing your client-side routing to handle navigation properly.

### CORS Support

To enable Cross-Origin Resource Sharing:

```bash
caflax --cors
```

This allows your served content to be accessed from different domains, which is useful when developing APIs or front-end applications that need to communicate with external services.

### Caching Control

During development, you might want to disable browser caching to ensure you always see the latest changes:

```bash
caflax --no-cache
```

### Compression

By default, Caflax compresses responses using gzip and brotli to reduce bandwidth usage. Compression is enabled by default, but you can control it:

```bash
# Explicitly enable compression (already on by default)
caflax --compress

# Disable compression
caflax --compress=false
```

### Quiet Mode

For less verbose output:

```bash
caflax --quiet
```

### Checking Version

To display version information:

```bash
caflax --version
```

## Complete Command Reference

Here's a complete list of available options:

| Option | Short | Default | Description |
|--------|-------|---------|-------------|
| `--port` | `-p` | 3000 | Server port |
| `--host` | `-H` | localhost | Server hostname |
| `--dir` | `-d` | . | Directory to serve |
| `--spa` | `-s` | false | Enable Single Page Application mode |
| `--cors` | `-c` | false | Enable Cross-Origin Resource Sharing |
| `--no-cache` | `-n` | false | Disable browser caching |
| `--compress` | `-z` | true | Enable gzip/brotli compression |
| `--quiet` | `-q` | false | Quiet mode, reduce log output |
| `--version` | `-v` | false | Show version information |

For more help, run:

```bash
caflax --help
```
