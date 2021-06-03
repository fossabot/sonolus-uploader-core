package potato

import (
	"errors"

	"cloud.google.com/go/firestore"
)

type CacheService struct {
	// cacheInitService
	init *CacheInitService
	// backgrounds stores background
	backgrounds Cache
	// effects stores effect
	effects Cache
	// engines store engine
	engines Cache
	// levels stores level
	levels Cache
	// particles stores particle
	particles Cache
	// skins stores skin
	skins Cache
	// users stores user
	users Cache
	// news stores fake levels
	news Cache
}

func NewCacheService(firestore *firestore.Client) *CacheService {
	s := NewCacheInitService(firestore)
	return &CacheService{
		init: s,
	}
}

func (s *CacheService) InitNews() {
	s.news.Data = make(map[string]interface{})
	s.news.Add(
		"sweetPotatoWelcome",
		NewNews(
			"SweetPotatoサーバーへようこそ!",
			"もっと をタップして一覧を表示してください",
			"sweetPotatoWelcome",
			"PurplePalette DevTeam",
			410,
			"Cocoa",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/cocoa.jpg",
			"※この譜面は遊べません",
		),
	)
	s.news.Add(
		"sweetPotatoUserWelcome",
		NewNews(
			"SweetPotato ユーザー個別サーバー",
			"Placeholder",
			"sweetPotatoUserWelcome",
			"PurplePalette DevTeam",
			1204,
			"Chino",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/chino.jpg",
			"※この譜面は遊べません",
		),
	)
	s.news.Add(
		"sweetPotatoTestWelcome",
		NewNews(
			"SweetPotato テストサーバー",
			"Placeholder",
			"sweetPotatoTestWelcome",
			"PurplePalette DevTeam",
			214,
			"Rize",
			"https://gochiusa.com/core_sys/images/main/cont/special/37/t_icon4/rize.jpg",
			"※この譜面は遊べません",
		),
	)
}

func (s *CacheService) InitCache() error {
	backgrounds, err := s.init.LoadDatabaseFromFirebase("backgrounds")
	if err != nil {
		return errors.New("could not get backgrounds from firestore")
	}
	s.backgrounds.Data = backgrounds
	effects, err := s.init.LoadDatabaseFromFirebase("effects")
	if err != nil {
		return errors.New("could not get effects from firestore")
	}
	s.effects.Data = effects
	engines, err := s.init.LoadDatabaseFromFirebase("engines")
	if err != nil {
		return errors.New("could not get engines from firestore")
	}
	s.engines.Data = engines
	levels, err := s.init.LoadDatabaseFromFirebase("levels")
	if err != nil {
		return errors.New("could not get levels from firestore")
	}
	s.levels.Data = levels
	particles, err := s.init.LoadDatabaseFromFirebase("particles")
	if err != nil {
		return errors.New("could not get particles from firestore")
	}
	s.particles.Data = particles
	skins, err := s.init.LoadDatabaseFromFirebase("skins")
	if err != nil {
		return errors.New("could not get skins from firestore")
	}
	s.skins.Data = skins
	users, err := s.init.LoadDatabaseFromFirebase("users")
	if err != nil {
		return errors.New("could not get user list from firestore")
	}
	s.users.Data = users
	s.InitNews()
	return nil
}
