/*
 * SweetPotato Server API
 *
 * Sonolusの基本APIを拡張する感じ。 ユーザー認証はFirebaseAuthorizationを通してやる。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gen

type EngineDataNodes struct {

	// Function name. See https://github.com/NonSpicyBurrito/sonolus-wiki/wiki/Node-Functions
	Func string `json:"func"`

	// Indexes of argument nodes in nodes.
	Args []map[string]interface{} `json:"args"`
}
