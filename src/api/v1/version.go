package v1

import (
	"encoding/json"
	"../../config"
	"net/http"
	log "github.com/sirupsen/logrus"
)

type ServerVersion struct {
	Version string `json:"version"`
}

var versionJSON []byte

func init() {
	var err error
	versionJSON, err =json.Marshal(ServerVersion{
		Version: config.Version,
	})
	if err != nil {
		panic(err)
	}
}

func VersionHandler(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
	if _,err := w.Write(versionJSON); err != nil {
		log.WithError(err).Debug("Error writing version")
	}
}