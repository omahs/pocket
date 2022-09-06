// Package rpc provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package rpc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// SendRawTxParams defines model for SendRawTxParams.
type SendRawTxParams struct {
	Address     string `json:"address"`
	RawHexBytes string `json:"raw_hex_bytes"`
}

// StringResult defines model for StringResult.
type StringResult = string

// PostV1ClientBroadcastTxSyncJSONBody defines parameters for PostV1ClientBroadcastTxSync.
type PostV1ClientBroadcastTxSyncJSONBody = SendRawTxParams

// PostV1ClientBroadcastTxSyncJSONRequestBody defines body for PostV1ClientBroadcastTxSync for application/json ContentType.
type PostV1ClientBroadcastTxSyncJSONRequestBody = PostV1ClientBroadcastTxSyncJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /v1/client/broadcast_tx_sync)
	PostV1ClientBroadcastTxSync(ctx echo.Context) error

	// (GET /v1/health)
	GetV1Health(ctx echo.Context) error

	// (GET /v1/version)
	GetV1Version(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostV1ClientBroadcastTxSync converts echo context to params.
func (w *ServerInterfaceWrapper) PostV1ClientBroadcastTxSync(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostV1ClientBroadcastTxSync(ctx)
	return err
}

// GetV1Health converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1Health(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1Health(ctx)
	return err
}

// GetV1Version converts echo context to params.
func (w *ServerInterfaceWrapper) GetV1Version(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetV1Version(ctx)
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

	router.POST(baseURL+"/v1/client/broadcast_tx_sync", wrapper.PostV1ClientBroadcastTxSync)
	router.GET(baseURL+"/v1/health", wrapper.GetV1Health)
	router.GET(baseURL+"/v1/version", wrapper.GetV1Version)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xUUW/bNhD+Kwd2DyugSE7aAgOftrbZVmDYgjoNMARBeqbOFleJ5I6n2Fqg/z6QkpM4",
	"Mfa0J1Hkdx/vu7uP98r4LnhHTqLS9yqahjrMyyW5+jNuL3cXyNjlrcA+EIul/Id1zRTzUoZASqsobN1G",
	"jYVi3N42tLtdDULHEAlCf/eWqVb6+oHqeeDNWKhlDvlMsW8lMdEOu9AmstNyUS5U8YJ7LBTthNhh+9Gb",
	"dP1YKOvWPoUb7wTNxNShbZVWGKwQdj/GLW42xKX1SUJN0bANYr1TWunabzRog6JBM65WVjRcNjaCjYAQ",
	"c04Qie+I4YIkiuf9fwnwp+/BoIO1dTX4XqBLx7hKy+V0LaDAdSMSdFU9ZnLz/Yut1+AZvINry6ZcM5Hz",
	"NZWOpIBXM+pIVGXZVK9LgJ89g6TEp5wLGObc+kggDQEGC99ogK8xkLHYnnyj4SuIB6EoE6KXxrP9B1Nx",
	"YG1bIY6lKlRrDblIqbgOu9SUnwKahuAsN6rnVO45te12W2I+LT1vqjk0Vr99+nD++/L85KxclI10beqF",
	"EHfxj/WS+M4aeuR4Ii9DqjQNVvJ07Mu6b4Yq1B1xnNo5jc5YKB/IYbBKqzfzNAWUJo9sdXdamdaSk2rF",
	"HmuDUW5ldxsHZ7IbfMxjlDyRK/GpVlpd+ChXpx9y3Pt92OVumYKmoaco73097IeRXGbBEFprMk/1V0xJ",
	"7t2YVt8xrZVWr6pHu1azV6vnRh0P3SXcU96IwacCK+36tk2YJLAhbKVJV2zoiJpfSK5Of50wBxz36myx",
	"SJ9Dm0zQIVX27eLty/MvbrpwgBP44pjQNLhqKeHfHeN7il9O5jpn9pxEzgoemvpfEq5m0DMNNa1xflb+",
	"n1Y8fatyhodyPk73wT6JLGMs1PRORKWv7w9N0nqDbeOj6HeLH96cVWq8Gf8NAAD//7SlU5e0BQAA",
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
