package main

import (
	"log"
	"os"

	"github.com/jgsqware/iptv-parser/handlers"
	"github.com/jgsqware/iptv-parser/models"
	"github.com/labstack/echo"
)

var tplt = `#EXTM3U
{{range .}}#EXTINF:-1 tvg-id="{{.TVGID}}" tvg-name="{{.TVGName}}" tvg-logo="{{.TVGLogo}}" group-title="{{.GroupTitle}}",{{.Name}}
{{.URL}}
{{end}}
`

func main() {

	channels, err := models.Parse(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/groups", handlers.GetGroups(channels))
	e.GET("/groups/:id/channels", handlers.GetChannels(channels))
	//e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	//e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })
	e.Logger.Fatal(e.Start(":1323"))
}

//func cmd() {
//	app := cli.NewApp()
//	app.Name = "iptv-parser"
//	app.Usage = "Manage your IPTV playlist"
//
//	app.Flags = []cli.Flag{
//		cli.StringFlag{
//			Name:  "file, f",
//			Usage: "M3U playlist `FILE`",
//		},
//		cli.StringFlag{
//			Name:  "output, o",
//			Usage: "Output M3U playlist `FILE`",
//		},
//	}
//
//	app.Action = func(c *cli.Context) error {
//		f := c.String("file")
//		o := c.String("output")
//		channels, err := Parse(f)
//
//		if err != nil {
//			return err
//		}
//
//		tmpl := template.New("iptv")
//
//		tmpl, err = tmpl.Parse(tplt)
//		if err != nil {
//			return err
//		}
//		fo := os.Stdout
//		if o != "" {
//			fo, err = os.Create(o)
//
//			if err != nil {
//				return err
//			}
//		}
//		err1 := tmpl.Execute(fo, channels)
//		if err1 != nil {
//			return err
//		}
//
//		g := channelByGroup(channels)
//
//		for k, d := range g {
//			fmt.Printf("%s: %d\n", k, len(d))
//		}
//
//		return nil
//	}
//
//	app.Run(os.Args)
//}
