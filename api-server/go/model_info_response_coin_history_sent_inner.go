// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * API Avito shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi

// InfoResponseCoinHistorySentInner описывает исходящую транзакцию.
type InfoResponseCoinHistorySentInner struct {
	// Имя пользователя, которому отправлены монеты.
	ToUser string `json:"toUser,omitempty"`
	// Количество отправленных монет.
	Amount int32 `json:"amount,omitempty"`
}

// AssertInfoResponseCoinHistorySentInnerRequired проверяет, что обязательные поля заполнены.
func AssertInfoResponseCoinHistorySentInnerRequired(obj InfoResponseCoinHistorySentInner) error {
	// Аналогично можно проверить, что Amount > 0
	return nil
}

// AssertInfoResponseCoinHistorySentInnerConstraints проверяет ограничения значений полей.
func AssertInfoResponseCoinHistorySentInnerConstraints(obj InfoResponseCoinHistorySentInner) error {
	// Можно добавить дополнительную логику, если нужно
	return nil
}
