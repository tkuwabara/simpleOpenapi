// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /amazons)
	FindAmazons(ctx echo.Context, params FindAmazonsParams) error

	// (POST /amazons)
	CreateAmazon(ctx echo.Context) error

	// (PATCH /amazons/active/{key})
	ActiveAmazon(ctx echo.Context, key Key) error

	// (PATCH /amazons/inactive/{key})
	InactiveAmazon(ctx echo.Context, key Key) error

	// (DELETE /amazons/{key})
	DeleteAmazon(ctx echo.Context, key Key) error

	// (GET /amazons/{key})
	FindAmazonByKey(ctx echo.Context, key Key) error

	// (PATCH /amazons/{key})
	PatchAmazon(ctx echo.Context, key Key) error

	// (PUT /amazons/{key})
	UpdateAmazon(ctx echo.Context, key Key) error

	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error

	// (POST /pets)
	AddPet(ctx echo.Context) error

	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id ID) error

	// (GET /pets/{id})
	FindPetById(ctx echo.Context, id ID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindAmazons converts echo context to params.
func (w *ServerInterfaceWrapper) FindAmazons(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindAmazonsParams
	// ------------- Optional query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, false, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", ctx.QueryParams(), &params.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindAmazons(ctx, params)
	return err
}

// CreateAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAmazon(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateAmazon(ctx)
	return err
}

// ActiveAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) ActiveAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ActiveAmazon(ctx, key)
	return err
}

// InactiveAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) InactiveAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.InactiveAmazon(ctx, key)
	return err
}

// DeleteAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteAmazon(ctx, key)
	return err
}

// FindAmazonByKey converts echo context to params.
func (w *ServerInterfaceWrapper) FindAmazonByKey(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindAmazonByKey(ctx, key)
	return err
}

// PatchAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) PatchAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PatchAmazon(ctx, key)
	return err
}

// UpdateAmazon converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateAmazon(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "key" -------------
	var key Key

	err = runtime.BindStyledParameterWithLocation("simple", false, "key", runtime.ParamLocationPath, ctx.Param("key"), &key)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateAmazon(ctx, key)
	return err
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, false, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id ID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetById converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id ID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetById(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/amazons", wrapper.FindAmazons)
	router.POST(baseURL+"/amazons", wrapper.CreateAmazon)
	router.PATCH(baseURL+"/amazons/active/:key", wrapper.ActiveAmazon)
	router.PATCH(baseURL+"/amazons/inactive/:key", wrapper.InactiveAmazon)
	router.DELETE(baseURL+"/amazons/:key", wrapper.DeleteAmazon)
	router.GET(baseURL+"/amazons/:key", wrapper.FindAmazonByKey)
	router.PATCH(baseURL+"/amazons/:key", wrapper.PatchAmazon)
	router.PUT(baseURL+"/amazons/:key", wrapper.UpdateAmazon)
	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
