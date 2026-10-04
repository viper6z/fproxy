# Reverse proxy checkpoints

## 1. Forward to one hardcoded backend

- [x] GET through the proxy gives the same answer as direct
- [x] Request headers copied to the outbound request
- [x] Path and query survive: curl "localhost:8080/foo?x=1"
- [x] Method and body survive: curl -d "hello" localhost:8080/foo
- [ ] Large download works

## 2. Forward correctly

- [x] No automatic gzip (DisableCompression)
- [ ] Redirects are passed to the client, not followed
- [ ] Hop-by-hop headers stripped, both directions
- [ ] X-Forwarded-For carries the client IP
- [ ] Host: decide what the backend should see
- [ ] Streaming: a backend sending one line per second arrives one line per second
- [ ] Client disconnect cancels the backend request

## 3. Fail properly

- [ ] Backend down gives 502, not an empty 200
- [ ] Slow backend gives 504 after a timeout
- [ ] Nothing can hang forever
- [ ] Errors are logged

## 4. Several backends

- [ ] Round robin across three backends

## 5. Health

- [ ] Dead backend is taken out of rotation
- [ ] It comes back when healthy again

## 6. Routing

- [ ] Host and path prefix choose a backend group
- [ ] Read from a config file

## 7. Filter chain

- [ ] Access log
- [ ] One deny rule

## 8. Lifecycle

- [ ] Graceful shutdown lets in-flight requests finish
