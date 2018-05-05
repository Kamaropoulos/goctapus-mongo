package goctapus

import (
	"gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	Log "github.com/sirupsen/logrus"
)

var AppPort string
var DBHost string
var DBPort string
var DBUser string
var DBPass string

var Database *mgo.Session

var e *echo.Echo

func getArgs(args []string) (string, string, string, string, string) {

	// Set default configuration values
	AppPort := "8000"
	DBHost := "localhost"
	DBPort := "27017"
	DBUser := ""
	DBPass := ""

	switch argsCount := len(args); argsCount {
	case 1:
		// There was 1 argument passed, use it as
		// the port for the app to listen to
		if res, err := isUsablePort(args[0]); res {
			AppPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

	case 2:
		// There were 2 arguments passed
		// Use them as DB username and password
		DBUser = args[0]
		DBPass = args[1]

	case 3:
		// There were 3 arguments passed
		// Use them as App Port, DB Username and Password
		if res, err := isUsablePort(args[0]); res {
			AppPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

		DBUser = args[1]
		DBPass = args[2]

	case 4:
		// There were 4 arguments passed
		// Use them as App Port, DB Username, Password and Hostname/IP
		if res, err := isUsablePort(args[0]); res {
			AppPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

		DBHost = args[3]
		DBUser = args[1]
		DBPass = args[2]

	case 5:
		// There were 5 arguments passed covering all configuration parameters
		// Use them as App Port, DB Username, Password, Hostname/IP and Port
		if res, err := isUsablePort(args[0]); res {
			AppPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

		DBHost = args[3]
		DBPort = args[4]
		DBUser = args[1]

		if res, err := isValidPort(args[2]); res {
			AppPort = args[2]
		} else {
			panic("\"" + args[2] + "\" can't be used as a port number: " + err)
		}

		DBPass = args[2]
	}

	return AppPort, DBHost, DBPort, DBUser, DBPass
}

func Init(args []string, logLevel string, database string) {

	InitLogger(logLevel)

	Log.Info("Initializing Goctapus...")

	AppPort, DBHost, DBPort, DBUser, DBPass = getArgs(args[1:])

	Log.WithFields(Log.Fields{
		"AppPort": AppPort,
		"DBHost":  DBHost,
		"DBPort":  DBPort,
		"DBUser":  DBUser,
		"DBPass":  DBPass,
	}).Debug("Current server configuration:")

	// Databases = make(map[string]*sql.DB)
	// ConnectDB(DBHost, DBPort, database, DBUser, DBPass)

	e = echo.New()

	Log.Debug("Goctapus Initialization done.")
}

func Start() {
	Log.Info("Starting up web server...")
	e.Start(":" + AppPort)
	Log.Debug("Web Server started succesfully.")
}

func GET(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	e.GET(path, handler)
}

func POST(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	e.POST(path, handler)
}

func PUT(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	e.PUT(path, handler)
}

func DELETE(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	e.DELETE(path, handler)
}

func File(path, file string) {
	e.File(path, file)
}
