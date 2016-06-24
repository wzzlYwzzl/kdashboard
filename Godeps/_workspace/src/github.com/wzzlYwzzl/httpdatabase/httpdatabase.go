package main

import (
	//"encoding/json"
	//"flag"
	"fmt"
	"log"
	"net/http"
	//"database/sql"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/pflag"
	"github.com/wzzlYwzzl/httpdatabase/handler"
	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

var (
	argPort         = pflag.Int("port", 9080, "The port to listen to for incoming HTTP requests, default 9080")
	argDatabaseHost = pflag.String("database-host", "localhost:3306", "The address is the backend database address, eg. mysql. "+
		"address:port. If not specified, the assumption is that the database is running locally. ")
	argUsername = pflag.String("username", "", "The username of the user to login to the mysql.")
	argPassword = pflag.String("password", "", "The password of the mysql user.")
)

func main() {
	pflag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	dbconf := new(sqlop.MysqlCon)
	dbconf.Host = *argDatabaseHost
	dbconf.Name = *argUsername
	dbconf.Password = *argPassword

	log.Printf("Starting HTTP server on port %d", *argPort)
	log.Printf("mysql username: %s, password is %s", *argUsername, *argPassword)

	apiHandler := new(handler.ApiHandler)
	apiHandler.DBconf = dbconf

	http.Handle("/api/", apiHandler.CreateApiHandler())
	log.Print(http.ListenAndServe(fmt.Sprintf(":%d", *argPort), nil))

}
