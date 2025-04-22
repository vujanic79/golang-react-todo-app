package health

import (
	"github.com/vujanic79/golang-react-todo-app/pkg/controller/util"
	"net/http"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJson(w, http.StatusOK, map[string]string{"status": "ok"})
}
