package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"strings"
	"net/url"
)

func main() {
	// Initialize a new html engine
	engine := html.New("./views", ".html")
	engine.AddFunc("trimSuffix", func(s, suffix string) string {
    	return strings.TrimSuffix(s, suffix)
    })

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Assuming video files are in a directory named "videos"
		files, err := ioutil.ReadDir("./videos")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var videoNames []string
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".mp4" {
				videoNames = append(videoNames, f.Name())
			}
		}

		// Render index template
		return c.Render("index", fiber.Map{
			"VideoNames": videoNames,
		})
	})

	app.Get("/video/:name", func(c *fiber.Ctx) error {
		name, _ := url.QueryUnescape(c.Params("name"))

		// Render video template
		return c.Render("video", fiber.Map{
			"Name": name,
		})
	})

	// Assuming video files are in a directory named "videos"
	app.Static("/videos", "./videos")
	app.Static("/assets", "./assets")

	app.Listen(":3000")
}
