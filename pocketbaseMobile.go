package pocketbaseMobile

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// Java Callbacks, make sure to register them before starting pocketbase
// to expose any method to java, add that with FirstLetterCapital
var nativeBridge NativeBridge
var version string = "0.0.1"

func RegisterNativeBridgeCallback(c NativeBridge) { nativeBridge = c }

func StartPocketbase(path string, hostname string, port string, getApiLogs bool) {
	os.Args = append(os.Args, "serve", "--http", hostname+":"+port)
	appConfig := pocketbase.Config{
		DefaultDataDir: path,
	}
	app := pocketbase.NewWithConfig(&appConfig)
	setupPocketbaseCallbacks(app, getApiLogs)

	serverUrl := "http://" + hostname + ":" + port
	sendCommand("onServerStarting", fmt.Sprintln("Server starting at:", serverUrl+"\n",
		"➜ REST API: ", serverUrl+"/api/\n",
		"➜ Admin UI: ", serverUrl+"/_/"))
	if err := app.Start(); err != nil {
		sendCommand("error", fmt.Sprintln("Error: ", "Failed to start pocketbase server: ", err))
	}
}

func StopPocketbase() {
	sendCommand("log", "Stopping pocketbase...")
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}

func GetVersion() string {
	return version
}

// Helper methods
type NativeBridge interface {
	HandleCallback(string, string) string
}

// send command to native and return the response
func sendCommand(command string, data string) string {
	return nativeBridge.HandleCallback(command, data)
}

// Hooks :https://pocketbase.io/docs/event-hooks/
func setupPocketbaseCallbacks(app *pocketbase.PocketBase, getApiLogs bool) {
	// Setup callbacks
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		sendCommand("OnBeforeServe", "")
		if getApiLogs {
			e.Router.Use(ApiLogsMiddleWare(app))
		}
		// setup a native Get request handler
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/nativeGet",
			Handler: func(context echo.Context) error {
				var data = sendCommand("nativeGetRequest", context.QueryParams().Encode())
				return context.JSON(http.StatusOK, map[string]string{
					"success": data,
				})
			},
		})
		// setup a native Post request handler
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/nativePost",
			Handler: func(context echo.Context) error {
				form, error := context.FormValues()
				if error != nil {
					return context.JSON(http.StatusBadRequest, map[string]string{
						"error": error.Error(),
					})
				}
				var data = sendCommand("nativePostRequest", form.Encode())
				return context.JSON(http.StatusOK, map[string]string{
					"success": data,
				})
			},
		})
		return nil
	})
	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		sendCommand("OnBeforeBootstrap", "")
		return nil
	})
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		sendCommand("OnAfterBootstrap", "")
		return nil
	})
	app.OnTerminate().Add(func(e *core.TerminateEvent) error {
		sendCommand("OnTerminate", "")
		return nil
	})
}

// Middleware, this will log all api calls
func ApiLogsMiddleWare(app core.App) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			fullPath := request.URL.Host + request.URL.Path + "?" + request.URL.RawQuery
			sendCommand("apiLogs", fullPath)
			return next(c)
		}
	}
}
