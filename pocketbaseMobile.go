package pocketbaseMobile

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// Java Callbacks, make sure to register them before starting pocketbase
// to expose any method to java, add that with FirstLetterCapital
var nativeBridge NativeBridge
var version string = "0.0.2"
var app *pocketbase.PocketBase

func RegisterNativeBridgeCallback(c NativeBridge) { nativeBridge = c }

func StartPocketbase(path string, hostname string, port string, getApiLogs bool, staticFilesPath *string) {
	os.Args = append(os.Args, "serve", "--http", hostname+":"+port)
	appConfig := pocketbase.Config{
		DefaultDataDir: path,
	}

	if app != nil {
		sendCommand("log", "Pocketbase is already running")
		StopPocketbase()
	}

	app := pocketbase.NewWithConfig(appConfig)
	setupPocketbaseCallbacks(app, getApiLogs, staticFilesPath)

	serverUrl := "http://" + hostname + ":" + port

	sendCommand("onServerStarting", fmt.Sprintln("Server starting at:", serverUrl+"\n",
		"➜ REST API: ", serverUrl+"/api/\n",
		"➜ Dashboard: ", serverUrl+"/_/"))

	if err := app.Start(); err != nil {
		sendCommand("error", fmt.Sprintln("Error: ", "Failed to start pocketbase server: ", err))
	}
}

func StopPocketbase() {
	sendCommand("log", "Stopping pocketbase...")
	if app == nil {
		sendCommand("log", "Pocketbase is not running")
		return
	}
	app.OnTerminate().Trigger(&core.TerminateEvent{App: app})
	app = nil
	sendCommand("log", "Pocketbase stopped")
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
func setupPocketbaseCallbacks(app *pocketbase.PocketBase, getApiLogs bool, staticFilesPath *string) {

	// Setup callbacks
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Set logs middleware if required
		if getApiLogs {
			se.Router.BindFunc(func(e *core.RequestEvent) error {
				fullPath := e.Request.URL.Host + e.Request.URL.Path + "?" + e.Request.URL.RawQuery
				sendCommand("apiLogs", fullPath)
				return e.Next()
			})
		}

		if staticFilesPath != nil {
			se.Router.GET("/{path...}", apis.Static(os.DirFS(*staticFilesPath), false))
		}

		// registers new "GET /hello" route
		se.Router.GET("/api/nativeGet", func(re *core.RequestEvent) error {
			var data = sendCommand("nativeGetRequest", re.Request.URL.Query().Encode())
			return re.JSON(http.StatusOK, map[string]string{
				"success": data,
			})
		})

		se.Router.POST("/api/nativePost", func(re *core.RequestEvent) error {
			var err = re.Request.ParseForm()
			if err == nil {
				var data = sendCommand("nativePostRequest", re.Request.Form.Encode())
				return re.JSON(http.StatusOK, map[string]string{
					"success": data,
				})
			} else {
				return re.JSON(http.StatusOK, map[string]string{
					"error": err.Error(),
				})
			}
		})

		return se.Next()
	})

	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		if err := e.Next(); err != nil {
			return err
		}
		sendCommand("OnBootstrap", "")
		return nil
	})

	app.OnTerminate().BindFunc(func(e *core.TerminateEvent) error {
		sendCommand("OnTerminate", "")
		return e.Next()
	})

}
