package business_error

import (
	"net/http"
)

const (
	Unknown = iota
	ParamsError
	SignatureError
	Unauthorized
	InvalidToken
	PermissionError

	RecordInvalid
	NotSaved
	NotFound
	NotDestroyed
	NotUnique
	InvalidForeignKey
	NotNull
	ValueTooLarge
	RangeError
)

type c = Common

var CommonErrors = map[int]Renderable{
	Unknown:         &c{-1, "unknown error", http.StatusInternalServerError},
	ParamsError:     &c{101001, "params validation failed", http.StatusBadRequest},
	SignatureError:  &c{101002, "signature error", http.StatusBadRequest},
	Unauthorized:    &c{101003, "unauthorized", http.StatusUnauthorized},
	InvalidToken:    &c{101010, "unauthorized", http.StatusUnauthorized},
	PermissionError: &c{101011, "insufficient permission", http.StatusForbidden},

	RecordInvalid:     &c{101020, "data validation failed", http.StatusOK},
	NotSaved:          &c{101021, "failed to save the record", http.StatusOK},
	NotFound:          &c{101022, "data not found", http.StatusOK},
	NotDestroyed:      &c{101023, "destroy failed", http.StatusOK},
	NotUnique:         &c{101024, "duplicate data", http.StatusOK},
	InvalidForeignKey: &c{101025, "association not found", http.StatusOK},
	NotNull:           &c{101026, "something should be not null", http.StatusOK},
	ValueTooLarge:     &c{101027, "data so large", http.StatusOK},
	RangeError:        &c{101028, "data out of range", http.StatusOK},
}
