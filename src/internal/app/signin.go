package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Login(w http.ResponseWriter, r *http.Request) {
	var loginDTO dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	///**
	//Parsing Session Data from Context
	//*/
	//userSession, err := utils.ParseCookie(r)
	//if err == nil || userSession != nil {
	//	response.ErrorResponse(w, http.StatusInternalServerError, messages.AlreadyLoggedIn)
	//	return
	//}

	userData, err := m.authService.SignIn(loginDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusUnauthorized, messages.UnsuccessfulLogin)
		return
	}
	//
	////sessionId, err := m.sessionService.CreateUserSession(*userData)
	//if err != nil {
	//	response.ErrorResponse(w, http.StatusForbidden, messages.AlreadyLoggedIn)
	//	return
	//}
	//
	//var (
	//	hashKey  = []byte(os.Getenv("HASH_KEY"))
	//	blockKey = []byte(os.Getenv("BLOCK_KEY"))
	//	s        = securecookie.New(hashKey, blockKey)
	//)
	//
	//encoded, err := s.Encode("sessionId", *sessionId)
	//if err != nil {
	//	fmt.Println(err)
	//	response.ErrorResponse(w, http.StatusInternalServerError, messages.FailToEncodeCookie)
	//	return
	//}
	//
	//responseCookie := http.Cookie{
	//	Name:     "sessionId",
	//	Value:    encoded,
	//	Expires:  time.Now().Add(1 * time.Hour),
	//	Secure:   true,
	//	HttpOnly: true,
	//	Path:     "/",
	//}
	//http.SetCookie(w, &responseCookie)

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulLogin, userData)
	return
}
