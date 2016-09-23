package main

import (
	"net/http"
	"strconv"
	//"fmt"

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
	aBean = website.CreateSite("AffableBean", "localhost:8070","en")
	aBean.AddParamList("languages", []string{"en", "cs", "fr", "sp"})
	addScripts();
	addMenus();
	addServices();
	addPages();
	addTemplatePages();
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
	aBean.AddParamTriggerHandler("language",chooseLanguage)
}

func addTemplatePages() {
	aBean.AddPage("", "head", "")
	aBean.AddPage("", "header", "")
	aBean.AddPage("", "footer", "")
	aBean.AddPage("category", "category", "")
	aBean.AddPage("", "lang", "")
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
	ecs.AddProduct("Dairy", "free range eggs", "medium-sized (6 eggs)", "free range eggs.png", 176, 30)

	ecs.AddProduct("Bakery", "sunflower seed loaf", "600g", "sunflower seed loaf.png", 189, 8)
	ecs.AddProduct("Bakery", "sesame seed bagel", "4 bagels", "sesame seed bagel.png", 119, 11)
	ecs.AddProduct("Bakery", "pumpkin seed bun", "4 buns", "pumpkin seed bun.png", 115, 14)
	ecs.AddProduct("Bakery", "chocolate cookies", "contain peanuts (3 cookies)", "chocolate cookies.png", 239, 27)

	ecs.AddProduct("Fruits & Veg", "corn on the cob", "2 pieces", "corn on the cob.png", 159, 14)
	ecs.AddProduct("Fruits & Veg", "red currants", "150g", "red currants.png", 249, 14)
	ecs.AddProduct("Fruits & Veg", "broccoli", "500g", "broccoli.png", 129, 14)
	ecs.AddProduct("Fruits & Veg", "seedless watermelon", "250g", "seedless watermelon.png", 149, 14)
	addProductsCs(ecs)
	addProductsFr(ecs)
	addProductsSp(ecs)
}

func addProductsCs(ecs *ecommerse.ECommerseService) {
	ecs.AddCategory("maso","Cured and Fresh Meats","meats.jpg")
	ecs.AddCategory("pečivo","Bakery Items","bakery.jpg")
	ecs.AddCategory("mléčné výrobky","mléčné výrobky","dairy.jpg")
	ecs.AddCategory("ovoce a zeleniny","Non GMO Fruits and Vegetables","fruit & veg.jpg")

	ecs.AddProduct("maso", "kořeněné pirohy", "v čerstvých bylinkách 2 placky (250g)", "organic meat patties.png", 229, 34)
	ecs.AddProduct("maso", "pršut", "vyzrálé, bio vepřové maso (70g)", "parma ham.png", 349, 11)
	ecs.AddProduct("maso", "kuřecí stehno", "z volného chovu (250g)", "chicken leg.png", 259, 61)
	ecs.AddProduct("maso", "klobásy", "s nízkým obsahem tuku 3 vepřové klobásky (350g)", "sausages.png", 359, 48)

	ecs.AddProduct("mléčné výrobky", "mléko", "polotučné (1L)", "milk.png", 170, 14)
	ecs.AddProduct("mléčné výrobky", "sýr", "jemný čedar (330g)", "cheese.png", 239, 51)
	ecs.AddProduct("mléčné výrobky", "máslo", "máslo ", "butter.png", 109, 21)
	ecs.AddProduct("mléčné výrobky", "vejce z volných pastvin", "vejce z volných pastvin", "free range eggs.png", 176, 30)

	ecs.AddProduct("pečivo", "slunečnicový chléb", "600g", "sunflower seed loaf.png", 189, 8)
	ecs.AddProduct("pečivo", "sezamová bageta", "4 bagely", "sesame seed bagel.png", 119, 11)
	ecs.AddProduct("pečivo", "semínková dýňová", "4 buchty", "pumpkin seed bun.png", 115, 14)
	ecs.AddProduct("pečivo", "čokoládové sušenky", "obsahují arašídy (3 sušenky)", "chocolate cookies.png", 239, 27)

	ecs.AddProduct("ovoce a zeleniny", "kukuřice", "2 klasy", "corn on the cob.png", 159, 14)
	ecs.AddProduct("ovoce a zeleniny", "červený rybíz", "150g", "red currants.png", 249, 14)
	ecs.AddProduct("ovoce a zeleniny", "brokolice", "500g", "broccoli.png", 129, 14)
	ecs.AddProduct("ovoce a zeleniny", "meloun bez semínek", "250g", "seedless watermelon.png", 149, 14)
}

func addProductsFr(ecs *ecommerse.ECommerseService) {
	ecs.AddCategory("Viandes","Salaisons et Fresh Meats","meats.jpg")
	ecs.AddCategory("Boulangerie","Boulangerie Articles","bakery.jpg")
	ecs.AddCategory("Laitier","les produits laitiers d'élevage localement","dairy.jpg")
	ecs.AddCategory("Fruits et Légumes","Fruits et légumes OGM non","fruit & veg.jpg")

	ecs.AddProduct("Viandes", "galettes de viande bio", "roulé dans des herbes fraîches 2 galettes (250g)", "organic meat patties.png", 229, 34)
	ecs.AddProduct("Viandes", "jambon de parme", "bio (70g)", "parma ham.png", 349, 11)
	ecs.AddProduct("Viandes", "cuisses de poulet", "plage libre (250g)", "chicken leg.png", 259, 61)
	ecs.AddProduct("Viandes", "saucisses", "la réduction de graisse, porc 3 saucisses (350g)", "sausages.png", 359, 48)

	ecs.AddProduct("Laitier", "le lait", "semi skimmed (1L)", "milk.png", 170, 14)
	ecs.AddProduct("Laitier", "de fromage", "mild cheddar (330g)", "cheese.png", 239, 51)
	ecs.AddProduct("Laitier", "de beurre", "unsalted (250g)", "butter.png", 109, 21)
	ecs.AddProduct("Laitier", "œufs de libre parcours", "de taille moyenne (6 oeufs)", "free range eggs.png", 176, 30)

	ecs.AddProduct("Boulangerie", "Pain de graines de tournesol", "600g", "sunflower seed loaf.png", 189, 8)
	ecs.AddProduct("Boulangerie", "bagel aux graines de sésame", "4 bagel", "sesame seed bagel.png", 119, 11)
	ecs.AddProduct("Boulangerie", "bun de graines de citrouille", "4 pain", "pumpkin seed bun.png", 115, 14)
	ecs.AddProduct("Boulangerie", "des biscuits au chocolat", "contenir des arachides (3 cookies)", "chocolate cookies.png", 239, 27)

	ecs.AddProduct("Fruits et Légumes", "Maïs en épi", "2 pieces", "corn on the cob.png", 159, 14)
	ecs.AddProduct("Fruits et Légumes", "groseilles rouges", "150g", "red currants.png", 249, 14)
	ecs.AddProduct("Fruits et Légumes", "brocoli", "500g", "broccoli.png", 129, 14)
	ecs.AddProduct("Fruits et Légumes", "pastèque sans pépins", "250g", "seedless watermelon.png", 149, 14)
}

func addProductsSp(ecs *ecommerse.ECommerseService) {
	ecs.AddCategory("Carnes","Cured and Fresh Meats","meats.jpg")
	ecs.AddCategory("Panadería","Bakery Items","bakery.jpg")
	ecs.AddCategory("Lácteos","Locally Farmed Dairy","dairy.jpg")
	ecs.AddCategory("Frutas y verduras","Non GMO Fruits and Vegetables","fruit & veg.jpg")

	ecs.AddProduct("Carnes", "empanadas de carne orgánica", "rolled in fresh herbs 2 patties (250g)", "organic meat patties.png", 229, 34)
	ecs.AddProduct("Carnes", "jamón de Parma", "organic (70g)", "parma ham.png", 349, 11)
	ecs.AddProduct("Carnes", "de las piernas de pollo", "free range (250g)", "chicken leg.png", 259, 61)
	ecs.AddProduct("Carnes", "salchichas", "reduced fat, pork 3 sausages (350g)", "sausages.png", 359, 48)

	ecs.AddProduct("Lácteos", "leche", "semi skimmed (1L)", "milk.png", 170, 14)
	ecs.AddProduct("Lácteos", "queso", "mild cheddar (330g)", "cheese.png", 239, 51)
	ecs.AddProduct("Lácteos", "mantequilla", "unsalted (250g)", "butter.png", 109, 21)
	ecs.AddProduct("Lácteos", "huevos de corral", "medium-sized (6 eggs)", "free range eggs.png", 176, 30)

	ecs.AddProduct("Panadería", "Pan de semillas de girasol", "600g", "sunflower seed loaf.png", 189, 8)
	ecs.AddProduct("Panadería", "bagel de semilla de sésamo", "4 bagels", "sesame seed bagel.png", 119, 11)
	ecs.AddProduct("Panadería", "pan de semillas de calabaza", "4 buns", "pumpkin seed bun.png", 115, 14)
	ecs.AddProduct("Panadería", "galletas de chocolate", "contain peanuts (3 cookies)", "chocolate cookies.png", 239, 27)

	ecs.AddProduct("Frutas y verduras", "Mazorca de maíz", "2 pieces", "corn on the cob.png", 159, 14)
	ecs.AddProduct("Frutas y verduras", "grosellas rojas", "150g", "red currants.png", 249, 14)
	ecs.AddProduct("Frutas y verduras", "brócoli", "500g", "broccoli.png", 129, 14)
	ecs.AddProduct("Frutas y verduras", "sandía sin semillas", "250g", "seedless watermelon.png", 149, 14)
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
	for _, lang := range aBean.ParamList["languages"] {
		for index, cat := range products.Body[lang]["Category"] {
			products.AddData("category:"+lang, `<a class="categoryButton"  href="/AffableBean/product?category=`+
				strconv.Itoa(index)+`" ><span class="categoryText">`+cat+`</span></a>`)
		}
	}
	ecs.AddPage("products", products)
	addProducts(ecs)
	
	viewCart := aBean.AddPage("viewCart", "viewCart", "/AffableBean/viewCart")
	viewCart.AddBypassSiteProcessor("secure")
	viewCart.AddParamTriggerHandler("clear",ecs.ClearCart)
}

func chooseLanguage(w http.ResponseWriter, r *http.Request, s *website.Session, p *website.Page) (string, error) {
	s.Data["language"]=p.Param["language"]
	http.Redirect(w,r,s.Data["navigation"],302)
	return "", nil
}