// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "ISprime Origins API": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/maurodelazeri/image-uploader/design
// --out=$(GOPATH)/src/github.com/maurodelazeri/image-uploader
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ImageController is the controller interface for the Image actions.
type ImageController interface {
	goa.Muxer
	goa.FileServer
	Show(*ShowImageContext) error
	Upload(*UploadImageContext) error
}

// MountImageController "mounts" a Image resource controller on the given service.
func MountImageController(service *goa.Service, ctrl ImageController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/api/images/:id", ctrl.MuxHandler("preflight", handleImageOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/api/images/", ctrl.MuxHandler("preflight", handleImageOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/download/*filename", ctrl.MuxHandler("preflight", handleImageOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowImageContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleImageOrigin(h)
	service.Mux.Handle("GET", "/api/images/:id", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Show", "route", "GET /api/images/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadImageContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Upload(rctx)
	}
	h = handleImageOrigin(h)
	service.Mux.Handle("POST", "/api/images/", ctrl.MuxHandler("upload", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "action", "Upload", "route", "POST /api/images/")

	h = ctrl.FileHandler("/download/*filename", "images/")
	h = handleImageOrigin(h)
	service.Mux.Handle("GET", "/download/*filename", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "files", "images/", "route", "GET /download/*filename")

	h = ctrl.FileHandler("/download/", "images/index.html")
	h = handleImageOrigin(h)
	service.Mux.Handle("GET", "/download/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Image", "files", "images/index.html", "route", "GET /download/")
}

// handleImageOrigin applies the CORS response headers corresponding to the origin.
func handleImageOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger-ui/*filepath", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/index.html")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/index.html", "route", "GET /swagger-ui/")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT")
				rw.Header().Set("Access-Control-Allow-Headers", "x-auth, Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
