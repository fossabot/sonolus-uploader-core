package potato_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/PurplePalette/sonolus-uploader-core/potato"
	"github.com/PurplePalette/sonolus-uploader-core/utils/server"
	"github.com/PurplePalette/sonolus-uploader-core/utils/tests"
	"github.com/stretchr/testify/assert"
)

// CreateLevelsServer creates background server for testing
func CreateLevelsServer() *httptest.Server {
	firebase := server.NewFirebaseClient()
	firestore := server.NewFirebaseFirestoreClient(firebase)
	auth := server.NewFirebaseAuthorizationClient(firebase)
	cache := potato.NewCacheService(firestore)
	if err := cache.InitCache(); err != nil {
		panic(err)
	}
	LevelsApiService := potato.NewLevelsApiService(firestore, cache)
	LevelsApiController := potato.NewLevelsApiController(LevelsApiService)
	router := server.NewRouterWithTestInject(auth, LevelsApiController)
	return httptest.NewServer(router)
}

func TestAddLevel(t *testing.T) {
	s := CreateLevelsServer()
	defer s.Close()
	lv := potato.Level{
		Name:    "testLevel",
		Version: 1,
		Rating:  10,
		Engine: potato.Engine{
			Name:     "defaultEngine",
			Version:  1,
			Title:    "デフォルトエンジン",
			Subtitle: "FromPurplePalette",
			Author:   "Dev",
			Thumbnail: potato.SonolusResourceLocator{
				Type: "EngineThumbnail",
				Url:  "https://example.com",
			},
			Data: potato.SonolusResourceLocator{
				Type: "EngineData",
				Url:  "https://example.com",
			},
			Configuration: potato.SonolusResourceLocator{
				Type: "EngineConfiguration",
				Url:  "https://example.com",
			},
			Skin: potato.Skin{
				Name:     "defaultSkin",
				Version:  1,
				Title:    "デフォルトスキン",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "SkinThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "SkinData",
					Url:  "https://example.com",
				},
				Texture: potato.SonolusResourceLocator{
					Type: "SkinTexture",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			Background: potato.Background{
				Name:     "defaultBackground",
				Version:  1,
				Title:    "デフォルト背景",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "BackgroundThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "BackgroundData",
					Url:  "https://example.com",
				},
				Image: potato.SonolusResourceLocator{
					Type: "BackgroundImage",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			Effect: potato.Effect{
				Name:     "defaultEffect",
				Version:  1,
				Title:    "デフォルトエフェクト",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "EffectThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "EffectData",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			Particle: potato.Particle{
				Name:     "defaultParticle",
				Version:  1,
				Title:    "デフォルトパーティクル",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "ParticleThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "ParticleData",
					Url:  "https://example.com",
				},
				Texture: potato.SonolusResourceLocator{
					Type: "ParticleTexture",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			CreatedTime: int32(time.Now().Unix()),
			UpdatedTime: int32(time.Now().Unix()),
			UserId:      "omado",
			Description: "説明文",
		},
		UseSkin: potato.LevelUseSkin{
			UseDefault: true,
		},
		UseBackground: potato.LevelUseBackground{
			UseDefault: true,
		},
		UseEffect: potato.LevelUseEffect{
			UseDefault: true,
		},
		UseParticle: potato.LevelUseParticle{
			UseDefault: true,
		},
		Title:   "テスト用譜面",
		Artists: "Dev",
		Author:  "お窓",
		Cover: potato.SonolusResourceLocator{
			Type: "LevelCover",
			Url:  "https://example.com",
		},
		Bgm: potato.SonolusResourceLocator{
			Type: "LevelBgm",
			Url:  "https://example.com",
		},
		Data: potato.SonolusResourceLocator{
			Type: "LevelData",
			Url:  "https://example.com",
		},
		Genre:       "general",
		Public:      false,
		UserId:      "YnaKWRpbanfyn1ge6FKQChqocyyn",
		Notes:       1,
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		Description: "テスト用です一覧に表示されたら成功です",
	}
	bgJson, err := json.Marshal(lv)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(
		http.MethodPost,
		"/levels/myLevel",
		bytes.NewBuffer(bgJson),
	)
	req = tests.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestEditLevel(t *testing.T) {
	s := CreateLevelsServer()
	defer s.Close()
	lv := potato.Level{
		Name:    "testLevel2",
		Version: 1,
		Rating:  20,
		Engine: potato.Engine{
			Name:     "defaultEngine",
			Version:  1,
			Title:    "デフォルトエンジン",
			Subtitle: "FromPurplePalette",
			Author:   "Dev",
			Thumbnail: potato.SonolusResourceLocator{
				Type: "EngineThumbnail",
				Url:  "https://example.com",
			},
			Data: potato.SonolusResourceLocator{
				Type: "EngineData",
				Url:  "https://example.com",
			},
			Configuration: potato.SonolusResourceLocator{
				Type: "EngineConfiguration",
				Url:  "https://example.com",
			},
			Skin: potato.Skin{
				Name:     "defaultSkin",
				Version:  1,
				Title:    "デフォルトスキン",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "SkinThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "SkinData",
					Url:  "https://example.com",
				},
				Texture: potato.SonolusResourceLocator{
					Type: "SkinTexture",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			Background: potato.Background{
				Name:     "defaultBackground",
				Version:  1,
				Title:    "デフォルト背景",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "BackgroundThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "BackgroundData",
					Url:  "https://example.com",
				},
				Image: potato.SonolusResourceLocator{
					Type: "BackgroundImage",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			Effect: potato.Effect{
				Name:     "defaultEffect",
				Version:  1,
				Title:    "デフォルトエフェクト",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "EffectThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "EffectData",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			Particle: potato.Particle{
				Name:     "defaultParticle",
				Version:  1,
				Title:    "デフォルトパーティクル",
				Subtitle: "FromPurplePalette",
				Author:   "Dev",
				Thumbnail: potato.SonolusResourceLocator{
					Type: "ParticleThumbnail",
					Url:  "https://example.com",
				},
				Data: potato.SonolusResourceLocator{
					Type: "ParticleData",
					Url:  "https://example.com",
				},
				Texture: potato.SonolusResourceLocator{
					Type: "ParticleTexture",
					Url:  "https://example.com",
				},
				CreatedTime: int32(time.Now().Unix()),
				UpdatedTime: int32(time.Now().Unix()),
				UserId:      "omado",
				Description: "説明文",
			},
			CreatedTime: int32(time.Now().Unix()),
			UpdatedTime: int32(time.Now().Unix()),
			UserId:      "omado",
			Description: "説明文",
		},
		UseSkin: potato.LevelUseSkin{
			UseDefault: true,
		},
		UseBackground: potato.LevelUseBackground{
			UseDefault: true,
		},
		UseEffect: potato.LevelUseEffect{
			UseDefault: true,
		},
		UseParticle: potato.LevelUseParticle{
			UseDefault: true,
		},
		Title:   "テスト用譜面",
		Artists: "Dev",
		Author:  "お窓",
		Cover: potato.SonolusResourceLocator{
			Type: "LevelCover",
			Url:  "https://example.com",
		},
		Bgm: potato.SonolusResourceLocator{
			Type: "LevelBgm",
			Url:  "https://example.com",
		},
		Data: potato.SonolusResourceLocator{
			Type: "LevelData",
			Url:  "https://example.com",
		},
		Genre:       "general",
		Public:      false,
		UserId:      "YnaKWRpbanfyn1ge6FKQChqocyyn",
		Notes:       1,
		CreatedTime: int32(time.Now().Unix()),
		UpdatedTime: int32(time.Now().Unix()),
		Description: "テスト用2です一覧に表示されたら成功です",
	}
	bgJson, _ := json.Marshal(lv)
	req := httptest.NewRequest(
		http.MethodPatch,
		"/levels/myLevel",
		bytes.NewBuffer(bgJson),
	)
	req = tests.SetUserAuthorizationToHeader(req)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetLevel(t *testing.T) {
	s := CreateLevelsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/levels/myLevel", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetLevelList(t *testing.T) {
	s := CreateLevelsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/levels/list", nil)
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetLevelListWithSpecifyKeyword(t *testing.T) {
	s := CreateLevelsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/levels/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "スキン2")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetLevelListWithSort(t *testing.T) {
	s := CreateLevelsServer()
	defer s.Close()
	req := httptest.NewRequest(http.MethodGet, "/levels/list", nil)
	params := req.URL.Query()
	params.Add("keywords", "sort:d order:d")
	req.URL.RawQuery = params.Encode()
	rec := httptest.NewRecorder()
	s.Config.Handler.ServeHTTP(rec, req)
	t.Log(rec.Body)
	assert.Equal(t, http.StatusOK, rec.Code)
}
