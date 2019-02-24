package routes

import(
  "fmt"
  "net/http"
  "go-webapp/utils"
  "go-webapp/models"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
  categories, err := models.GetCategories()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Internal Server Error"))
  }
  products, err := models.GetProducts()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Internal Server Error"))
  }
  utils.ExecuteTemplate(w, "home.html", struct{
    Categories []models.Category
    Products []models.Product
  }{
    Categories: categories,
    Products: products,
  })
}

func homePostHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type", "text/html; charset=UTF-8")
  r.ParseForm()
  search := r.PostForm.Get("search")
  products, err := models.SearchProducts(search)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("Internal Server Error"))
  }
  count := len(products)
  var html string = ""
  if count > 0 {
    html += "<table class='table table-bordered'>"
    html += fmt.Sprintf("<th> Id </th> <th> Categoria </th> <th> Nome </th> <th> Preço </th> <th> Quantidade </th> <th> Valor total</th>")
    for _, p := range products {
      html += "<tr>"
      html += fmt.Sprintf("<td> %d </td> <td> %s </td> <td> %s </td> <td> %.2f R$ </td> <td> %d </td> <td> %.2f </td>", p.Id, 
      p.Category.Description, p.Name, p.Price, p.Quantity, p.Amount)
      html += "</tr>"
    } 
    html += "</table>"
  } else {
    html += fmt.Sprintf(`<p class='alert alert-info'> Nada encotrado com <code>"<strong> %s </strong> </code>"</p>`, search)
  }

  w.Write([]byte(html))
}
