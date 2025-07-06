# Go Proxy Implementation Plan

## Project Structure
```
proxy-server/
├── cmd/
│   ├── forward/
│   │   └── main.go
│   ├── reverse/
│   │   └── main.go
│   └── combined/
│       └── main.go
├── pkg/
│   ├── forward/
│   │   └── proxy.go
│   ├── reverse/
│   │   └── proxy.go
│   └── common/
│       └── utils.go
├── config/
│   └── config.yaml
├── test-servers/
│   ├── server1.go
│   ├── server2.go
│   └── server3.go
└── README.md
```

## Phase 1: Basic Forward Proxy

### Core Concepts
- Forward proxy receives HTTP requests from clients
- Extracts target URL from request
- Makes request to target server
- Returns response back to client
- Client must be configured to use the proxy

### Implementation Steps

#### Step 1: Basic HTTP Forward Proxy
```go
// Key components you'll implement:
// 1. HTTP server that listens for requests
// 2. Request parser to extract destination
// 3. HTTP client to make upstream requests
// 4. Response relay back to client
```

**Main functionality:**
- Listen on port (e.g., 8080)
- Handle `GET`, `POST`, `PUT`, `DELETE` requests
- Parse `Host` header and request URL
- Forward request to destination server
- Copy response headers and body back to client

#### Step 2: Add Request/Response Logging
```go
// Log format example:
// [2024-01-15 10:30:45] FORWARD GET http://example.com/api -> 200 OK (234ms)
```

#### Step 3: Add Basic Filtering
- Block certain domains (e.g., social media)
- Block based on request headers
- Rate limiting per client IP

### Key Go Packages You'll Use
- `net/http` - Core HTTP server/client
- `net/url` - URL parsing
- `io` - Copying request/response bodies
- `log` - Logging
- `time` - Timestamps and timeouts

## Phase 2: Basic Reverse Proxy

### Core Concepts
- Reverse proxy receives requests meant for backend services
- Routes requests based on rules (path, host, headers)
- Maintains pool of backend servers
- Load balances between multiple backends

### Implementation Steps

#### Step 1: Single Backend Reverse Proxy
```go
// Key components:
// 1. HTTP server listening on public port
// 2. Backend server configuration
// 3. Request routing logic
// 4. httputil.ReverseProxy for actual proxying
```

**Main functionality:**
- Listen on port (e.g., 8081)
- Define backend server (e.g., localhost:3000)
- Use `httputil.NewSingleHostReverseProxy()`
- Forward all requests to backend

#### Step 2: Path-Based Routing
```go
// Routing rules:
// /api/* -> localhost:3001
// /static/* -> localhost:3002
// /* -> localhost:3000
```

#### Step 3: Multiple Backend Support
- Round-robin load balancing
- Health checking
- Failover to backup servers

### Key Go Packages You'll Use
- `net/http/httputil` - ReverseProxy
- `net/http` - HTTP server
- `net/url` - URL parsing and manipulation
- `strings` - Path matching
- `sync` - Concurrent access to backend pools

## Phase 3: Advanced Features

### Forward Proxy Enhancements
1. **HTTPS Support (CONNECT method)**
   - Handle CONNECT requests for HTTPS tunneling
   - Proxy TCP connections for SSL/TLS

2. **Caching**
   - Cache GET responses
   - Implement cache headers (Cache-Control, ETag)
   - LRU cache eviction

3. **Authentication**
   - Basic auth for proxy access
   - Client certificate validation

### Reverse Proxy Enhancements
1. **Load Balancing Algorithms**
   - Round-robin
   - Least connections
   - Weighted round-robin

2. **Health Checks**
   - Periodic health checks to backends
   - Automatic failover
   - Backend recovery detection

3. **SSL Termination**
   - Accept HTTPS requests
   - Forward HTTP to backends
   - Certificate management

## Phase 4: Configuration and Testing

### Configuration System
```yaml
# config.yaml example
forward_proxy:
  port: 8080
  blocked_domains:
    - facebook.com
    - twitter.com
  rate_limit: 100 # requests per minute

reverse_proxy:
  port: 8081
  routes:
    - path: "/api/*"
      backends:
        - "localhost:3001"
        - "localhost:3002"
    - path: "/static/*"
      backends:
        - "localhost:3003"
    - path: "/*"
      backends:
        - "localhost:3000"
```

### Test Servers
Create simple HTTP servers for testing:
```go
// server1.go - API server on :3001
// server2.go - Static file server on :3002
// server3.go - Main web app on :3000
```

### Testing Strategy
1. **Unit Tests**
   - Test routing logic
   - Test request/response handling
   - Test configuration parsing

2. **Integration Tests**
   - Start test servers
   - Test proxy functionality end-to-end
   - Test failure scenarios

3. **Manual Testing**
   - Configure browser to use forward proxy
   - Test reverse proxy with curl/Postman
   - Monitor logs and metrics

## Phase 5: Monitoring and Metrics

### Metrics to Track
- Request count by method/status
- Response times
- Error rates
- Active connections
- Backend health status

### Implementation
```go
// Use channels for metrics collection
// Implement simple HTTP endpoint for metrics
// Add Prometheus metrics support (optional)
```

## Learning Outcomes

By completing this project, you'll understand:
- HTTP protocol details (headers, methods, status codes)
- Network programming in Go
- Request routing and load balancing
- Proxy authentication and security
- Configuration management
- Testing network applications
- Performance monitoring

## Getting Started Checklist

1. Set up Go workspace and modules
2. Create basic project structure
3. Implement simple forward proxy (HTTP GET only)
4. Test with curl: `curl -x localhost:8080 http://httpbin.org/get`
5. Implement simple reverse proxy
6. Create test backend servers
7. Add logging and basic error handling
8. Gradually add advanced features

## Useful Resources
- Go `net/http` documentation
- HTTP/1.1 specification (RFC 7230-7237)
- Proxy-related RFCs (7230 section 5.7)
- Go concurrency patterns for handling multiple connections
