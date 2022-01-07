# sitemap.xml Gin Middleware
Provides sitemap.xml support for Gin.

Shamelessly modified clone of [gorobots](https://github.com/vasiliyaltunin/gorobots) :grimacing: (originally based on [Favicon Middleware for Gin](https://github.com/thinkerou/favicon))


# Usage
Assuming you've already prepared your `sitemap.xml` file, you can proceed with the steps below.

## Step 1
Download and install the package

```
go get https://github.com/dragstor/ginsitemap
```

## Step 2
Add the following line to your `import` section:

```go
import "github.com/dragstor/ginsitemap"
```

## Step 3
Add the middleware to your `router`:
```go
r.Use(ginsitemap.New("./path/to/sitemap/sitemap.xml"))
```



# Example Code
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/dragstor/ginsitemap"
)

func main() {
    r := gin.Default()
    r.Use(ginsitemap.New("./sitemap.xml"))

    r.Run(":8080")
}
```

# License
MIT