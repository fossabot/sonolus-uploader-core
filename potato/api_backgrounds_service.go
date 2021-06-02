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
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/PurplePalette/sonolus-uploader-core/utils/request"
)

// BackgroundsApiService is a service that implents the logic for the BackgroundsApiServicer
// This service should implement the business logic for every endpoint for the BackgroundsApi API.
// Include any external packages or services that will be required by this service.
type BackgroundsApiService struct {
	firestore *firestore.Client
	cache     *CacheService
}

// NewBackgroundsApiService creates a default api service
func NewBackgroundsApiService(firestore *firestore.Client, cache *CacheService) BackgroundsApiServicer {
	return &BackgroundsApiService{firestore: firestore, cache: cache}
}

// AddBackground - Add background
func (s *BackgroundsApiService) AddBackground(ctx context.Context, backgroundName string, background Background) (ImplResponse, error) {
	if !request.IsLoggedIn(ctx) {
		return Response(http.StatusUnauthorized, nil), nil
	}
	if !request.IsValidName(backgroundName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	if s.cache.IsBackgroundExist(backgroundName) {
		return Response(http.StatusConflict, nil), nil
	}
	// Force set parameter to valid
	userId, _ := request.GetUserId(ctx)
	background.UserId = userId
	background.Name = backgroundName
	col := s.firestore.Collection("backgrounds")
	// Add background to firestore
	if _, err := col.Doc(backgroundName).Set(ctx, background); err != nil {
		log.Fatalln("Error posting background:", err)
		return Response(500, nil), nil
	}
	// Add background to cache
	s.cache.backgroundList.Add(backgroundName, background)
	return Response(200, nil), nil
}

// EditBackground - Edit background
func (s *BackgroundsApiService) EditBackground(ctx context.Context, backgroundName string, background Background) (ImplResponse, error) {
	if !request.IsLoggedIn(ctx) {
		return Response(http.StatusUnauthorized, nil), nil
	}
	if !request.IsValidName(backgroundName) {
		return Response(http.StatusBadRequest, nil), nil
	}
	userId, _ := request.GetUserId(ctx)
	match, err := s.cache.backgroundList.IsOwnerMatch(backgroundName, userId)
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	if !match {
		return Response(http.StatusForbidden, nil), nil
	}
	// Update background data in firestore
	col := s.firestore.Collection("backgrounds")
	if _, err := col.Doc(backgroundName).Set(ctx, background); err != nil {
		log.Fatalln("Error posting background:", err)
		return Response(500, nil), nil
	}
	// Update background data in cache
	s.cache.backgroundList.Set(backgroundName, background)
	return Response(200, nil), nil
}

// GetBackground - Get background
func (s *BackgroundsApiService) GetBackground(ctx context.Context, backgroundName string) (ImplResponse, error) {
	bg, err := s.cache.backgroundList.Get(backgroundName)
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	resp := GetBackgroundResponse{
		Item:        bg.(Background),
		Description: "",
		Recommended: []Background{},
	}
	return Response(200, resp), nil
}

// GetBackgroundList - Get background list
func (s *BackgroundsApiService) GetBackgroundList(ctx context.Context, localization string, page int32, keywords string) (ImplResponse, error) {
	query := request.ParseSearchQuery(keywords)
	pages := s.cache.backgroundList.Pages()
	items, err := s.cache.backgroundList.GetPage(page, query)
	if err != nil {
		log.Fatal(err)
		return Response(500, nil), nil
	}
	var backgrounds []Background
	err = json.Unmarshal(items, &backgrounds)
	if err != nil {
		return Response(500, nil), nil
	}
	resp := GetBackgroundListResponse{
		PageCount: pages,
		Items:     backgrounds,
	}
	return Response(200, resp), nil
}
