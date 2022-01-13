package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

var verbose *bool

func proxyStaticFile(url string, c *fiber.Ctx) error {
	if *verbose {
		println("proxy static url : " + url)
	}

	c.Request().Header.SetContentType("application/javascript")

	if err := proxy.Do(c, url); err != nil {
		return err
	}
	// Remove Server header from response
	c.Response().Header.Del(fiber.HeaderServer)
	c.Response().Header.SetContentType("application/javascript")
	return nil
}

func proxyPassPath(url string, c *fiber.Ctx) error {
	if *verbose {
		println("proxy pass path url : " + url)
	}

	if err := proxy.Do(c, url); err != nil {
		return err
	}
	// Remove Server header from response
	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}

func main() {
	// proxy
	port := flag.Int("port", 2000, "Proxy Port. Normaly 2000. Auth should redirect there ")
	verbose = flag.Bool("verbose", false, "Display more information, access files")
	help := flag.Bool("help", false, "Display default commands")

	// clin-frontend
	frontendHost := flag.String("frontend-host", "http://0.0.0.0:2002", "clin-frontend host name or ip plus the port if not 80")
	frontendStaticPath := flag.String("frontend-staticpath", "/static", "clin-frontend development static ressources url")
	frontendStaticFiles := []string{"/config.js", "/manifest.json"}

	// clin-ui
	uiHost := flag.String("ui-host", "http://0.0.0.0:2005", "clin-ui host name or ip plus the port if not 80")
	uiStaticPath := flag.String("clinui-staticpath", "/clinui-static", "clin-frontend development static ressources url")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}

	app := fiber.New()

	app.Get(*uiStaticPath+"/*", func(c *fiber.Ctx) error {
		url := *uiHost + *uiStaticPath + "/" + c.Params("*")
		return proxyStaticFile(url, c)
	})

	app.Get("/search/*", func(c *fiber.Ctx) error {
		url := *uiHost + c.Params("*")
		return proxyPassPath(url, c)
	})

	app.Get(*frontendStaticPath+"/*", func(c *fiber.Ctx) error {
		url := *frontendHost + *frontendStaticPath + "/" + c.Params("*")
		return proxyStaticFile(url, c)
	})

	app.Get("/patient/search/*", func(c *fiber.Ctx) error {
		url := *frontendHost + "/patient/search/" + c.Params("*")
		return proxyPassPath(url, c)
	})

	for _, file := range frontendStaticFiles {
		app.Get(file, func(c *fiber.Ctx) error {
			url := *frontendHost + file
			return proxyStaticFile(url, c)
		})
	}
	app.Get("/*", func(c *fiber.Ctx) error {
		url := *frontendHost + "/" + c.Params("*")
		return proxyPassPath(url, c)
	})

	listenTo := fmt.Sprintf(":%d", *port)
	println("_______________________________\n")
	println(" Dev Proxy started")

	app.Listen(listenTo)
}
