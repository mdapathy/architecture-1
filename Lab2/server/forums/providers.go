package forums

import ("github.com/google/wire")

var Providers = wire.NewSet(HttpHandler)