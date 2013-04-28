# Summary

A simple JPEG image resizer service in Go. Uses [https://github.com/nfnt/resize](http://github.com/nfnt/resize). No caching (you could use a CDN in front of this for production).

# Build

`go get && go install`

# Run

`./resizer`

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

# Known issues

- Progressive JPEG support is only available in Go 1.1