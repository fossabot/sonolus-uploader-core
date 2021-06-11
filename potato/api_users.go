/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A UsersApiController binds http requests to an api service and writes the service results to the http response
type UsersApiController struct {
	service UsersApiServicer
}

// NewUsersApiController creates a default api controller
func NewUsersApiController(s UsersApiServicer) Router {
	return &UsersApiController{service: s}
}

// Routes returns all of the api route for the UsersApiController
func (c *UsersApiController) Routes() Routes {
	return Routes{
		{
			"GetUserList",
			strings.ToUpper("Get"),
			"/users/list",
			c.GetUserList,
		},
		{
			"EditUser",
			strings.ToUpper("Patch"),
			"/users/{userId}",
			c.EditUser,
		},
		{
			"GetUser",
			strings.ToUpper("Get"),
			"/users/{userId}",
			c.GetUser,
		},
		{
			"GetUserServerInfo",
			strings.ToUpper("Get"),
			"/users/{userId}/info",
			c.GetUserServerInfo,
		},
		{
			"GetUsersBackgrounds",
			strings.ToUpper("Get"),
			"/users/{userId}/backgrounds/list",
			c.GetUsersBackgrounds,
		},
		{
			"GetUsersEffects",
			strings.ToUpper("Get"),
			"/users/{userId}/effects/list",
			c.GetUsersEffects,
		},
		{
			"GetUsersEngines",
			strings.ToUpper("Get"),
			"/users/{userId}/engines/list",
			c.GetUsersEngines,
		},
		{
			"GetUsersLevels",
			strings.ToUpper("Get"),
			"/users/{userId}/levels/list",
			c.GetUsersLevels,
		},
		{
			"GetUsersParticles",
			strings.ToUpper("Get"),
			"/users/{userId}/particles/list",
			c.GetUsersParticles,
		},
		{
			"GetUsersSkins",
			strings.ToUpper("Get"),
			"/users/{userId}/skins/list",
			c.GetUsersSkins,
		},
		{
			"GetUsersBackground",
			strings.ToUpper("Get"),
			"/users/{userId}/backgrounds/{backgroundName}",
			c.GetUsersBackground,
		},
		{
			"GetUsersEffect",
			strings.ToUpper("Get"),
			"/users/{userId}/effects/{effectName}",
			c.GetUsersEffect,
		},
		{
			"GetUsersEngine",
			strings.ToUpper("Get"),
			"/users/{userId}/engines/{engineName}",
			c.GetUsersEngine,
		},
		{
			"GetUsersLevel",
			strings.ToUpper("Get"),
			"/users/{userId}/levels/{levelName}",
			c.GetUsersLevel,
		},
		{
			"GetUsersParticle",
			strings.ToUpper("Get"),
			"/users/{userId}/particles/{particleName}",
			c.GetUsersParticle,
		},
		{
			"GetUsersSkin",
			strings.ToUpper("Get"),
			"/users/{userId}/skins/{skinName}",
			c.GetUsersSkin,
		},
	}
}

// EditUser - Edit user
func (c *UsersApiController) EditUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.EditUser(r.Context(), userId, *user)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUser - Get user
func (c *UsersApiController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	result, err := c.service.GetUser(r.Context(), userId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUserList - Get user list
func (c *UsersApiController) GetUserList(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetUserList(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUserServerInfo - Get user server info
func (c *UsersApiController) GetUserServerInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	result, err := c.service.GetUserServerInfo(r.Context(), userId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersBackgrounds - Get backgrounds for test
func (c *UsersApiController) GetUsersBackgrounds(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userId := params["userId"]

	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetUsersBackgrounds(r.Context(), userId, localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersEffects - Get effects for test
func (c *UsersApiController) GetUsersEffects(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userId := params["userId"]

	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetUsersEffects(r.Context(), userId, localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersEngines - Get engines for test
func (c *UsersApiController) GetUsersEngines(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userId := params["userId"]

	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetUsersEngines(r.Context(), userId, localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersLevels - Get levels for test
func (c *UsersApiController) GetUsersLevels(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userId := params["userId"]

	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetUsersLevels(r.Context(), userId, localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersParticles - Get particles for test
func (c *UsersApiController) GetUsersParticles(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userId := params["userId"]

	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetUsersParticles(r.Context(), userId, localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersSkins - Get skins for test
func (c *UsersApiController) GetUsersSkins(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	userId := params["userId"]

	localization := query.Get("localization")
	page, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keywords := query.Get("keywords")
	result, err := c.service.GetUsersSkins(r.Context(), userId, localization, page, keywords)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersBackground - Get users background
func (c *UsersApiController) GetUsersBackground(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	backgroundName := params["backgroundName"]

	result, err := c.service.GetUsersBackground(r.Context(), userId, backgroundName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersEffect - Get users effect
func (c *UsersApiController) GetUsersEffect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	effectName := params["effectName"]

	result, err := c.service.GetUsersEffect(r.Context(), userId, effectName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersEngine - Get users engine
func (c *UsersApiController) GetUsersEngine(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	engineName := params["engineName"]

	result, err := c.service.GetUsersEngine(r.Context(), userId, engineName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersLevel - Get users level
func (c *UsersApiController) GetUsersLevel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	levelName := params["levelName"]

	result, err := c.service.GetUsersLevel(r.Context(), userId, levelName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersParticle - Get users particle
func (c *UsersApiController) GetUsersParticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	particleName := params["particleName"]

	result, err := c.service.GetUsersParticle(r.Context(), userId, particleName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetUsersSkin - Get users skin
func (c *UsersApiController) GetUsersSkin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	skinName := params["skinName"]

	result, err := c.service.GetUsersSkin(r.Context(), userId, skinName)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
