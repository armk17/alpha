package api

import (
	"net/http"
	"strings"
)

func (rt *Router) GetFeeds(w http.ResponseWriter, r *http.Request) {
	usr := getAuthUser(r)

	tp := strings.TrimSpace(r.URL.Query().Get("type"))
	uName := strings.TrimSpace(r.URL.Query().Get("username"))

	if uName != "" {
		usr = uName
	}

	pgr := newPager(r.URL, 20)

	twts, err := rt.feedRepo.List(usr, tp, pgr.offset(), pgr.limit())
	if err != nil {
		panic(newAPIError("Internal Server Error", errInternalServer, err))
	}
	resp := response{
		code: http.StatusOK,
		Data: twts,
		Meta: pgr,
	}

	resp.serveJSON(w)

}
