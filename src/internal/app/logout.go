package app

import (
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"net/http"
	"time"
)

func (m *MicroserviceServer) Logout(w http.ResponseWriter, r *http.Request) {

	expiredCookie := &http.Cookie{
		Name:     "sessionData",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, expiredCookie)

	_, err := r.Cookie("sessionData")
	if err != nil {
		response.ErrorResponse(w, http.StatusForbidden, messages.SessionExpired)
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogout, nil)
}
