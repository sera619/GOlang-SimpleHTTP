package main


import(
	"fmt"
	"log"
	"net/http"
	"html/template"
)


var temp *template.Template

func init(){
	temp = template.Must(template.ParseGlob("public/html/*.html"))

}
func main(){
	fileserver := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileserver)
	http.HandleFunc("/home", home)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("\nWebServer is running on http://localhost:8080 ...\n")
	if err := http.ListenAndServe((":8080"), nil); err != nil{
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/home"{
		http.Error(w, "Not found", 404)
		return
	}
	fmt.Fprintf(w, "Hello World")


}
func formHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form"{
		http.Error(w, "Not found", 404)
		return
	}
	temp.ExecuteTemplate(w, "home.html",nil)
}

