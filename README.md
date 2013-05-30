# Summary

A simple image resizer service in Go. Supports JPEG and PNG, and uses [https://github.com/nfnt/resize](http://github.com/nfnt/resize). No caching (you could use a CDN in front of this for production).

`go get github.com/pranavraja/imageresizer`

# Starting the server

Assuming `$GOPATH/bin` is in your `$PATH`, just run `imageresizer`. The process will start in the foreground, to run in production, use something like upstart to daemonize the process.

# Usage

Go to [http://localhost:8080/?source=http://1.bp.blogspot.com/-w6AJJ5Xoulg/T3AymMKbFlI/AAAAAAAAACI/EslCtw42HHg/s1600/jpeg.jpg&width=250](http://localhost:8080/?source=http://1.bp.blogspot.com/-w6AJJ5Xoulg/T3AymMKbFlI/AAAAAAAAACI/EslCtw42HHg/s1600/jpeg.jpg&width=250) in a browser, and you should see a small image of the lion from Madagascar.

You may also pass in `algorithm` as an additional query string parameter, the following are supported. See [https://github.com/nfnt/resize](http://github.com/nfnt/resize) for details.

- `nearestNeighbour` (default)
- `bilinear`
- `bicubic`
- `mitchellNetravali`
- `lanczos2`
- `lanczos3`

# Test

    ¯\_(ツ)_/¯

