package api

import (
	"net/http"
	"strings"

	"github.com/ark1790/alpha/model"
)

type toggleFollowPld struct {
	Profile string `json:"profile"`
}

func (t *toggleFollowPld) validate() *validationError {
	t.Profile = strings.TrimSpace(t.Profile)

	errV := validationError{}
	if t.Profile == "" {
		errV.add("profile", "is required")
	}

	if len(errV) > 0 {
		return &errV
	}

	return nil
}

func (rt *Router) ToggleFollow(w http.ResponseWriter, r *http.Request) {
	usr := getAuthUser(r)

	body := toggleFollowPld{}
	if err := parseBody(r, &body); err != nil {
		handleAPIError(w, newAPIError("Unable to parse body", errBadRequest, err))
		return
	}

	if err := body.validate(); err != nil {
		handleAPIError(w, newAPIError("Invalid data", errInvalidData, err))
		return
	}

	if usr == body.Profile {
		handleAPIError(w, newAPIError("Invalid data", errInvalidData, nil))
		return
	}

	flw := &model.Follow{
		Username: usr,
		Profile:  body.Profile,
	}

	if err := rt.followRepo.Toggle(flw); err != nil {
		panic(newAPIError("DB Failed", errInternalServer, err))
	}

	resp := response{
		code: http.StatusAccepted,
	}

	resp.serveJSON(w)
}
