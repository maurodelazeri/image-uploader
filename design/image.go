package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
) // User defines the media type used to render users. The mediaType gives all fields and the views break it down to specific responses

var _ = Resource("image", func() {
	BasePath("/images")

	Action("upload", func() {
		Routing(POST("/"))
		Description("Upload multiple images in multipart request")
		Response(OK, CollectionOf(ImageMedia))
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Description("Show an image metadata")
		Params(func() {
			Param("id", Integer, "Image ID")
		})
		Response(OK, ImageMedia)
		Response(NotFound)
	})

	Files("/download/*filename", "images/") // Serve files from the "images" directory
})

var ImageMedia = MediaType("application/json", func() {
	Description("Image metadata")
	TypeName("ImageMedia")
	Attributes(func() {
		Attribute("id", Integer, "Image ID")
		Attribute("filename", String, "Image filename")
		Attribute("uploaded_at", DateTime, "Upload timestamp")
		Required("id", "filename", "uploaded_at")
	})
	View("default", func() {
		Attribute("id")
		Attribute("filename")
		Attribute("uploaded_at")
	})
})
