package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("ISprime Origins API", func() {
	Title("ISprime")
	Description("This API exposes an image resource that allows uploading and downloading images")
	Contact(func() {
		Name("Mauro Delazeri")
		Email("mauro.delazeri@image-uploader.com")
	})
	Host("localhost:8080")
	Origin("*", func() {
		Methods("GET", "POST", "DELETE", "PUT")
		Headers("x-auth, Content-Type")
	})
	Scheme("http")
	BasePath("/api")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("swagger", func() {
	Description("The swagger definition")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})
