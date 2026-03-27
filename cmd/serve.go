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

type ProductPage struct {
	User string
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
			tmpl, err := template.New("productpage").Parse(string(data))
			if err != nil {
				c.Data(http.StatusInternalServerError, defaultContentType, []byte(err.Error()))
			}
			productPage := ProductPage{User: "admin"}
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
