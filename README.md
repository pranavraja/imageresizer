# Summary

A simple image resizer service in Go. Supports JPEG and PNG, and uses [https://github.com/nfnt/resize](http://github.com/nfnt/resize). No caching (you could use a CDN in front of this for production).

`go get github.com/pranavraja/imageresizer`

# Starting the server

Assuming `$GOPATH/bin` is in your `$PATH`, just run `imageresizer`. The process will start in the foreground, to run in production, use something like upstart to daemonize the process.

# Usage

Go to [http://localhost:8080/?source=http://i.imgur.com/B6TRfAf.jpg&width=250](http://localhost:8080/?source=http://i.imgur.com/B6TRfAf.jpg&width=250) in a browser, and you should see a cool image.

You may also pass in `algorithm` as an additional query string parameter, the following are supported. See [https://github.com/nfnt/resize](http://github.com/nfnt/resize) for details.

- `nearestNeighbour` (default)
- `bilinear`
- `bicubic`
- `mitchellNetravali`
- `lanczos2`
- `lanczos3`

The server uses `net/http/pprof`, so you can use [http://localhost:8080/debug/pprof](http://localhost:8080/debug/pprof) to inspect the current running goroutines and stack.

# Test

    ¯\_(ツ)_/¯

