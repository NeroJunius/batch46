package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Title          string
	Author         string
	StartDate      string
	EndDate        string
	ContentProject string
	NodeJs         string
	NodeReact      string
	TypeScript     string
	ReactJs        string
}

var dataProject = []Project{
	{
		Title:          "My First Time Open World Game",
		Author:         "Nafiisan N. Achmad",
		StartDate:      "5/2/2020",
		EndDate:        "6/2/2020",
		ContentProject: "COGNOSPHERE -- Genshin Impact, adalah game yang sedang hype selama 2 tahun terakhir. Saya sebagai pemain yang main dari awal rilis secara global, hingga saat ini masih memainkannya. Genre open world dengan basic ARPG tema anime, yang mana saya sendiri suka anime, menjadikannya daya tarik tersendiri.",
	},
	{
		Title:          "My Second Time Open World Game",
		Author:         "Nero",
		StartDate:      "6/2/2020",
		EndDate:        "7/2/2020",
		ContentProject: "COGNOSPHERE -- Genshin Impact, adalah game yang sedang hype selama 2 tahun terakhir. Saya sebagai pemain yang main dari awal rilis secara global, hingga saat ini masih memainkannya. Genre open world dengan basic ARPG tema anime, yang mana saya sendiri suka anime, menjadikannya daya tarik tersendiri.",
	},
	{
		Title:          "My 3rd Time Open World Game",
		Author:         "Junius",
		StartDate:      "6/2/2020",
		EndDate:        "7/2/2020",
		ContentProject: "COGNOSPHERE -- Genshin Impact, adalah game yang sedang hype selama 2 tahun terakhir. Saya sebagai pemain yang main dari awal rilis secara global, hingga saat ini masih memainkannya. Genre open world dengan basic ARPG tema anime, yang mana saya sendiri suka anime, menjadikannya daya tarik tersendiri.",
	},
}

func main() {
	// create new echo instance
	e := echo.New()

	// serve static files from public
	e.Static("/assets", "assets")

	// routing
	e.GET("/", home)
	e.GET("/contactMe", contactMe)
	e.GET("/post", post)
	e.POST("/add-post", addPost)
	e.GET("/blogDetail/:id", blogDetail)
	e.POST("/editProject/:id", editProject)
	e.GET("/deleteProject/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tab/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}
	posts := map[string]interface{}{
		"Posts": dataProject,
	}
	return tmpl.Execute(c.Response(), posts)
}

func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tab/contact-me.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func post(c echo.Context) error {
	var tmpl, err = template.ParseFiles("tab/blog.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var tmpl, err = template.ParseFiles("tab/blog.detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	var DataProject = Project{}
	for index, data := range dataProject {
		if id == index {
			DataProject = Project{
				Title:          data.Title,
				Author:         data.Author,
				StartDate:      data.StartDate,
				EndDate:        data.EndDate,
				ContentProject: data.ContentProject,
			}
		}
	}

	data := map[string]interface{}{
		"Project": DataProject,
	}

	return tmpl.Execute(c.Response(), data)
}

func addPost(c echo.Context) error {
	title := c.FormValue("title")
	contentProject := c.FormValue("contentProject")
	nodeJs := c.FormValue("nodeJs")
	nodeReact := c.FormValue("nodeReact")
	TypeScript := c.FormValue("TypeScript")
	reactJs := c.FormValue("reactJs")

	var addPost = Project{
		Title:          title,
		Author:         "Nafiisan N. Achmad",
		StartDate:      time.Now().String(),
		EndDate:        time.Now().String(),
		ContentProject: contentProject,
		NodeJs:         nodeJs,
		NodeReact:      nodeReact,
		TypeScript:     TypeScript,
		ReactJs:        reactJs,
	}

	fmt.Println(addPost)
	dataProject = append(dataProject, addPost)

	// fmt.Println(title)
	// fmt.Println(startDate)
	// fmt.Println(endDate)
	// fmt.Println(contentProject)
	// fmt.Println(nodeJs)
	// fmt.Println(nodeReact)
	// fmt.Println(TypeScript)
	// fmt.Println(reactJs)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
