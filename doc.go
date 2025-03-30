// This package provides an interface to interact with the Telegram Bot API.
// All requests sent with [SendRequest] or [SendRequestWith] either store the response in the provided destination
// or return [ErrTelegramAPIFailedRequest], providing details on the failure,
// or [ErrFailedValidation], if the request body failed validation.
//
// Additionally, utility functions are available for working with JSON encoding.
//
// Licensed under the MIT License. See LICENSE file for details.
package gotely
