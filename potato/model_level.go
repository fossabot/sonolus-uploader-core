/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package potato

// Level - A level provides a list of entities and powered by scripted behavior in engine, with skin, effect, background and particle, to create gameplay experience for players / It defines level for actual user play. It includes all data to play one level. https://github.com/NonSpicyBurrito/sonolus-wiki/wiki/Level
type Level struct {

	// english and number only name for searching
	Name string `json:"name,omitempty" validate:"omitempty,alphanum,min=1,max=50"`

	// Reserved for future update. current default is 1.
	Version int32 `json:"version,omitempty" validate:"omitempty,gte=1,lte=1"`

	// Difficulty of the level
	Rating int32 `json:"rating,omitempty" validate:"omitempty,gte=1,lte=1000"`

	Engine Engine `json:"engine,omitempty" validate:"omitempty"`

	UseSkin LevelUseSkin `json:"useSkin,omitempty" validate:"omitempty"`

	UseBackground LevelUseBackground `json:"useBackground,omitempty" validate:"omitempty"`

	UseEffect LevelUseEffect `json:"useEffect,omitempty" validate:"omitempty"`

	UseParticle LevelUseParticle `json:"useParticle,omitempty" validate:"omitempty"`

	// base title of this content
	Title string `json:"title,omitempty" validate:"omitempty,min=1,max=100"`

	// artist names of original music
	Artists string `json:"artists,omitempty" validate:"omitempty,min=1,max=100"`

	// author of this content
	Author string `json:"author,omitempty" validate:"omitempty,alphanumunicode,min=1,max=50"`

	Cover SonolusResourceLocator `json:"cover,omitempty" validate:"omitempty"`

	Bgm SonolusResourceLocator `json:"bgm,omitempty" validate:"omitempty"`

	Data SonolusResourceLocator `json:"data,omitempty" validate:"omitempty"`

	// 独自要素: 楽曲のジャンル
	Genre string `json:"genre,omitempty"`

	// 独自要素: 楽曲が全体公開かどうか
	Public bool `json:"public,omitempty"`

	// 独自要素: 譜面作成者のユーザーID
	UserId string `json:"userId,omitempty" validate:"omitempty,alphanum,min=1,max=50"`

	// 独自要素: 譜面内のノーツ数
	Notes int32 `json:"notes,omitempty" validate:"gte=1,lte=10000000"`

	// 独自要素: データを作成したエポックミリ秒(ソート用)
	CreatedTime int32 `json:"createdTime,omitempty" validate:"gte=1"`

	// 独自要素: データを更新したエポックミリ秒(ソート用)
	UpdatedTime int32 `json:"updatedTime,omitempty" validate:"gte=1"`

	// 独自要素: サイト内および譜面情報欄に表示される説明文
	Description string `json:"description,omitempty" validate:"min=1,max=3000"`
}
