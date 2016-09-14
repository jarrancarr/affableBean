package main

import (
	"net/http"
	"strconv"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
)

var aBean *website.Site
var acs *website.AccountService
var ecs *ecommerse.ECommerseService

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
	aBean = website.CreateSite("AffableBean", "localhost:8070")
	addScripts();
	addMenus();
	addServices();
	addPages();
	addTemplatePages();
	
	//viewCart := aBean.AddPage("viewCart", "viewCart", "/AffableBean/viewCart")
	//viewCart.AddBypassSiteProcessor("secure")
}

func addScripts() {
		
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
}

func addMenus() {	
	aBean.AddMenu("lang").
		AddItem("English", "/lang?english").
		AddItem("Spanish", "/lang?spanish").
		Add("", "display: inline;", "")
}

func addTemplatePages() {
	aBean.AddPage("head", "head", "")
	aBean.AddPage("category", "category", "")
}

func addProducts(ecs *ecommerse.ECommerseService) {
	ecs.AddCategory("Meats","Cured and Fresh Meats","meats.jpg")
	ecs.AddCategory("Bakery","Bakery Items","bakery.jpg")
	ecs.AddCategory("Dairy","Locally Farmed Dairy","dairy.jpg")
	ecs.AddCategory("Fruit & Veg","Non GMO Fruits and Vegetables","fruit & veg.jpg")

	ecs.AddProduct("Meats", "organic meat patties", "rolled in fresh herbs 2 patties (250g)", "organic meat patties.png", 229, 34)
	ecs.AddProduct("Meats", "parma ham", "organic (70g)", "parma ham.png", 349, 11)
	ecs.AddProduct("Meats", "chicken leg", "free range (250g)", "chicken leg.png", 259, 61)
	ecs.AddProduct("Meats", "sausages", "reduced fat, pork 3 sausages (350g)", "sausages.png", 359, 48)

	ecs.AddProduct("Dairy", "milk", "semi skimmed (1L)", "milk.png", 170, 14)
	ecs.AddProduct("Dairy", "cheese", "mild cheddar (330g)", "cheese.png", 239, 51)
	ecs.AddProduct("Dairy", "butter", "unsalted (250g)", "butter.png", 109, 21)
	ecs.AddProduct("Dairy", "free range eggs", "medium-sized (6 eggs)", "eggs.png", 176, 30)

	ecs.AddProduct("Bakery", "sunflower seed loaf", "600g", "sunflower seed loaf.png", 189, 8)
	ecs.AddProduct("Bakery", "sesame seed bagel", "4 bagels", "sesame seed bagel.png", 119, 11)
	ecs.AddProduct("Bakery", "pumpkin seed bun", "4 buns", "pumpkin seed bun.png", 115, 14)
	ecs.AddProduct("Bakery", "chocolate cookies", "contain peanuts (3 cookies)", "chocolate cookies.png", 239, 27)

	ecs.AddProduct("Fruits & Veg", "corn on the cob", "2 pieces", "corn on the cob.png", 159, 14)
	ecs.AddProduct("Fruits & Veg", "red currants", "150g", "red currants.png", 249, 14)
	ecs.AddProduct("Fruits & Veg", "broccoli", "500g", "broccoli.png", 129, 14)
	ecs.AddProduct("Fruits & Veg", "seedless watermelon", "250g", "seedless watermelon.png", 149, 14)
}

func addServices() {
	acs = website.CreateAccountService()
	aBean.AddService("account", acs)
	aBean.AddSiteProcessor("secure", acs.CheckSecure)
	ecs = ecommerse.CreateService(acs)
	aBean.AddService("ecommerse", ecs)
}

func addPages() {
	main := aBean.AddPage("AffableBean", "main", "/")
	main.AddBypassSiteProcessor("secure")

	products := aBean.AddPage("products", "products", "/AffableBean/product")
	products.AddBypassSiteProcessor("secure")
	products.AddPostHandler("addToCart", ecs.AddToCart)
	for index, cat := range []string{"Meats", "Dairy", "Bakery", "Fruit & Veg"} {
		products.AddData("category", `<a class="categoryButton"  href="/AffableBean/product?category=`+strconv.Itoa(index)+`" ><span class="categoryText">`+cat+`</span></a>`)
	}
	ecs.AddPage("products", products)
	addProducts(ecs)
}