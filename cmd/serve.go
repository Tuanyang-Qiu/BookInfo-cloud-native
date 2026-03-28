/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type Base struct {
	Status int
	Error  string
}

type ProductPage struct {
	User    string
	Product Product
	Details Details
	Reviews Reviews
}

type Product struct {
	Title           string
	DescriptionHtml string
}

type Details struct {
	Base
	ISBN10    string
	Publisher string
	Pages     int
	Type      string
	Language  string
}
type Rating struct {
	Base
	Stars int
	Color string
}
type Review struct {
	Rating   Rating
	Text     string
	Reviewer string
}
type Reviews struct {
	Base
	Reviews     []Review
	PodName     string
	ClusterName string
}

const defaultContentType string = "text/html,charset=utf-8"

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "server",
	Long:  `Start BookInfo server.`,
	Run: func(cmd *cobra.Command, args []string) {

		r := gin.Default()

		r.Static("/static", "./static")

		r.GET("/productpage", func(c *gin.Context) {
			data, _ := os.ReadFile("templates/productPage.html")

			funcs := template.FuncMap{"html_format": func(s string) template.HTML {
				return template.HTML(s)
			}, "inRange": func(stars int) []int {
				return make([]int, stars)
			}, "sub": func(a, b int) int {
				return a - b
			}}

			tmpl, err := template.New("productpage").Funcs(funcs).Parse(string(data))
			if err != nil {
				c.Data(http.StatusInternalServerError, defaultContentType, []byte(err.Error()))
			}
			productPage := ProductPage{
				User: "admin",
				Product: Product{
					Title:           "The Comedy of Errors",
					DescriptionHtml: `<a href="https://en.wikipedia.org/wiki/The_Comedy_of_Errors">Wikipedia Summary</a>: The Comedy of Errors is one of <b>William Shakespeare\'s</b> early plays. It is his shortest and one of his most farcical comedies, with a major part of the humour coming from slapstick and mistaken identity, in addition to puns and word play.`,
				},
				Details: Details{
					Base:      Base{Status: 200},
					ISBN10:    "1234567890",
					Publisher: "PublisherA",
					Pages:     200,
					Type:      "paperback",
					Language:  "English",
				},
				Reviews: Reviews{
					Base: Base{Status: 200},
					Reviews: []Review{
						{
							Rating: Rating{
								Base:  Base{Status: 200},
								Stars: 4,
								Color: "black",
							},
							Text:     "good",
							Reviewer: "test user 1",
						},
						{
							Rating: Rating{
								Base:  Base{Status: 200},
								Stars: 5,
								Color: "red",
							},
							Text:     "very good",
							Reviewer: "test user 2",
						},
					},
					PodName:     "test pod 1",
					ClusterName: "test cluster 1",
				},
			}
			buf := bytes.Buffer{}
			err = tmpl.Execute(&buf, productPage)
			if err != nil {
				c.Data(http.StatusInternalServerError, defaultContentType, []byte(err.Error()))
				return
			}
			c.Data(http.StatusOK, defaultContentType, buf.Bytes())
		})

		r.GET("/", func(c *gin.Context) {
			data, _ := os.ReadFile("templates/index.html")
			c.Data(http.StatusOK, defaultContentType, data)
		})

		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		r.Run()
		fmt.Println("serve called!!!")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
