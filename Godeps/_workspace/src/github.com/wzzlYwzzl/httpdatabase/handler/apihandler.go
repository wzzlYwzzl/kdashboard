package handler

import (
	"fmt"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	//"github.com/wzzlYwzzl/httpdatabase/client"
	"github.com/wzzlYwzzl/httpdatabase/resource/user"
	"github.com/wzzlYwzzl/httpdatabase/sqlop"
)

type ApiHandler struct {
	userns *user.User
	DBconf *sqlop.MysqlCon
}

const (
	// RequestLogString is a template for request log message.
	RequestLogString = "Incoming %s %s %s request from %s"

	// ResponseLogString is a template for response log message.
	ResponseLogString = "Outcoming response to %s with %d status code"
)

// Web-service filter function used for request and response logging.
func wsLogger(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf(FormatRequestLog(req))
	chain.ProcessFilter(req, resp)
	log.Printf(FormatResponseLog(resp, req))
}

// FormatRequestLog formats request log string.
// TODO(maciaszczykm): Display request body.
func FormatRequestLog(req *restful.Request) string {
	reqURI := ""
	if req.Request.URL != nil {
		reqURI = req.Request.URL.RequestURI()
	}

	return fmt.Sprintf(RequestLogString, req.Request.Proto, req.Request.Method,
		reqURI, req.Request.RemoteAddr)
}

// FormatResponseLog formats response log string.
// TODO(maciaszczykm): Display response content.
func FormatResponseLog(resp *restful.Response, req *restful.Request) string {
	return fmt.Sprintf(ResponseLogString, req.Request.RemoteAddr, resp.StatusCode())
}

func (apiHandler *ApiHandler) CreateApiHandler() http.Handler {
	wsContainer := restful.NewContainer()
	log.Printf("CreateAPiHander")
	userWs := new(restful.WebService)
	userWs.Filter(wsLogger)
	userWs.Path("/api/v1/user").
		//Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	userWs.Route(userWs.GET("/{name}").
		To(apiHandler.judgeUser))
	userWs.Route(userWs.GET("/{name}/{namespaces}").
		To(apiHandler.createNS))
	userWs.Route(userWs.GET("/ns/{name}").
		To(apiHandler.getNS))
	userWs.Route(userWs.GET("/ns/all/{name}").
		To(apiHandler.getNSAll))
	userWs.Route(userWs.GET("/allinfo/{name}").
		To(apiHandler.getAllInfo))
	userWs.Route(userWs.DELETE("/{name}").
		To(apiHandler.deleteUser))
	userWs.Route(userWs.POST("").
		To(apiHandler.createUser).
		Reads(user.User{}))
	userWs.Route(userWs.POST("/resource").
		To(apiHandler.updateResource).
		Reads(user.User{}))
	userWs.Route(userWs.GET("/all").
		To(apiHandler.getAllUserInfo).
		Writes(user.UserList{}))

	wsContainer.Add(userWs)

	return wsContainer
}

func (apiHandler *ApiHandler) judgeUser(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	apiHandler.userns = new(user.User)
	apiHandler.userns.Name = name
	log.Printf("judgeUser")

	b, err := apiHandler.userns.JudgeExist(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	if b == false {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	response.WriteHeader(http.StatusOK)
}

func (apiHandler *ApiHandler) createNS(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	namespaces := request.PathParameter("namespaces")

	apiHandler.userns = new(user.User)
	apiHandler.userns.Name = name
	apiHandler.userns.Namespaces = append(apiHandler.userns.Namespaces, namespaces)

	log.Printf("createNS username is %s, namespace is %s", apiHandler.userns.Name, apiHandler.userns.Namespaces)

	err := apiHandler.userns.CreateNamespace(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (apiHandler *ApiHandler) getNS(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")

	apiHandler.userns = new(user.User)
	apiHandler.userns.Name = name

	err := apiHandler.userns.GetNamespaces(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	log.Println("GetNS result is", apiHandler.userns.Namespaces)
	response.WriteHeaderAndEntity(http.StatusCreated, apiHandler.userns.Namespaces)
}

func (apiHandler *ApiHandler) getNSAll(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")

	apiHandler.userns = new(user.User)
	apiHandler.userns.Name = name

	err := apiHandler.userns.GetNamespacesAll(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}
	log.Println("GetNS result is", apiHandler.userns.Namespaces)
	response.WriteHeaderAndEntity(http.StatusCreated, apiHandler.userns.Namespaces)
}

func (apiHandler *ApiHandler) getAllInfo(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	apiHandler.userns = new(user.User)
	apiHandler.userns.Name = name

	err := apiHandler.userns.GetAllInfo(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, apiHandler.userns)
}

func (apiHandler *ApiHandler) deleteUser(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	apiHandler.userns = new(user.User)
	apiHandler.userns.Name = name

	log.Println("deleteUser is called")
	err := apiHandler.userns.DeleteUser(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.WriteHeader(http.StatusOK)
}

func (apiHandler *ApiHandler) createUser(request *restful.Request, response *restful.Response) {
	user := new(user.User)
	err := request.ReadEntity(user)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	err = user.CreateUser(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (apiHandler *ApiHandler) updateResource(request *restful.Request, response *restful.Response) {
	user := new(user.User)
	err := request.ReadEntity(user)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	err = user.UpdateResource(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (apiHandler *ApiHandler) getAllUserInfo(request *restful.Request, response *restful.Response) {
	userlist := new(user.UserList)

	err := userlist.GetAllUserInfo(apiHandler.DBconf)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, userlist)
}

// Handler that writes the given error to the response and sets appropriate HTTP status headers.
func handleInternalError(response *restful.Response, err error) {
	log.Print(err)
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(http.StatusInternalServerError, err.Error()+"\n")
}
