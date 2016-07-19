package lightning

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

const (
	Version = "1.0.0"
	Logo    = `__________        ______ _____       _____
___  /__(_)______ ___  /___  /__________(_)_____________ _
__  /__  /__  __  /_  __ \  __/_  __ \_  /__  __ \_  __  /
_  / _  / _  /_/ /_  / / / /_ _  / / /  / _  / / /  /_/ /
/_/  /_/  _\__, / /_/ /_/\__/ /_/ /_//_/  /_/ /_/_\__, /
          /____/                                 /____/`
)

func Run() {
	fmt.Printf("\x1b[36;1m%s %s\x1b[0m\n\n\x1b[32;1mStarted at %s\x1b[0m\n", Logo, Version, time.Now())
}

func ListenAndServe(addr string, handler fasthttp.RequestHandler) error {
	Run()
	return fasthttp.ListenAndServe(addr, handler)
}

func ListenAndServeUNIX(addr string, mode os.FileMode, handler fasthttp.RequestHandler) error {
	Run()
	return fasthttp.ListenAndServeUNIX(addr, mode, handler)
}

func ListenAndServeTLS(addr, certFile, keyFile string, handler fasthttp.RequestHandler) error {
	Run()
	return fasthttp.ListenAndServeTLS(addr, certFile, keyFile, handler)
}

func ListenAndServeTLSEmbed(addr string, certData, keyData []byte, handler fasthttp.RequestHandler) error {
	Run()
	return fasthttp.ListenAndServeTLSEmbed(addr, certData, keyData, handler)
}
