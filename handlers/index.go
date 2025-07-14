/*
Package handlers: Handlers for website interactions
*/
package handlers

import (
	"aluUpgradeCalc/views"
	"net/http"

	"github.com/a-h/templ"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	templ.Handler(views.HomePage("", []string{})).ServeHTTP(w, r)
}
