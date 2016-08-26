package main

import (
	"net/http"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
)

var Shelf []ecommerse.Category
var aBean *website.Site

func main() {
	website.ResourceDir = ".."
	setup()

	http.HandleFunc("/js/", website.ServeResource)
	http.HandleFunc("/css/", website.ServeResource)
	http.HandleFunc("/img/", website.ServeResource)
	http.ListenAndServe(":8070", nil)
}

func setup() {
	//website
	aBean = website.CreateSite("AffableBean", "localhost:8090")
	aBean.AddMenu("lang").
		AddItem("English", "/lang?english").
		AddItem("Spanish", "/lang?spanish").
		Add("", "display: inline;", "")
		
	aBean.AddScript("script",`$('a.categoryButton').hover(
		function () {$(this).animate({backgroundColor: '#b2d2d2'})},
		function () {$(this).animate({backgroundColor: '#d3ede8'})}  );`)
	aBean.AddScript("script",`$('div.categoryBox').hover(over, out); `)
	aBean.AddScript("script",`function over() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.3});
		$(span[1]).animate({color: 'white'}); } `)
	aBean.AddScript("script",`function out() {
		var span = this.getElementsByTagName('span');
		$(span[0]).animate({opacity: 0.7});
		$(span[1]).animate({color: '#444'}); } `)
	// services
	acs := website.CreateAccountService()
	aBean.AddService("account", acs)
	aBean.AddSiteProcessor("secure", acs.CheckSecure)
	ecs := ecommerse.CreateService(acs)
	aBean.AddService("ecommerse", ecs)

	// template subpages
	aBean.AddPage("head", "head", "")
	aBean.AddPage("category", "category", "")

	// pages
	main := aBean.AddPage("AffableBean", "main", "/")
	
	for _, cat := range []string{"Meats", "Dairy", "Bakery", "Fruits&Veg"} {
		main.AddData("category", `<div class="categoryButton"><span class="categoryText">`+cat+`</span></div>`)
	}
	
	main.AddBypassSiteProcessor("secure")
		
		
	products := aBean.AddPage("products", "products", "/AffableBean/category")
	
	products.AddBypassSiteProcessor("secure")
}