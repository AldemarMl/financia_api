package api
import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter()(http.Handler,error)  {
	router := mux.NewRouter()
	return router, nil
}