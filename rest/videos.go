package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"personal-website-server/env"
)

type VideosHandler struct {
	RestHandler
}

type VideosLoginRequest struct {
	Password string `json:"password"`
}

func (vh VideosHandler) login(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondWithError(createHTTPBodyError(err), http.StatusBadRequest, &w)
		return
	}

	request := VideosLoginRequest{}
	err = json.Unmarshal(data, &request)
	if err != nil {
		respondWithError(createJSONParseError(err), http.StatusBadRequest, &w)
		return
	}

	if request.Password == vh.env.VideosPassword {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func newVideosHandler(env *env.Env) *VideosHandler {
	return &VideosHandler{
		RestHandler: RestHandler{
			env,
		},
	}
}
