package utils

import (
	"errors"
)

var (
	// InternalServerErr is returned when there was unknown error.
	InternalServerErr = errors.New("internal server error")

	// UnauthorizedErr is returned when user credentials are invalid.
	UnauthorizedErr = errors.New("unauthorized")

	// ConflictErr is returned when the service cannot proceed with the requests due to conflict of data.
	ConflictErr = errors.New("conflict")

	// NotFoundErr is returned when the service cannot find the resource.
	NotFoundErr = errors.New("not found")
)

const (
	// InternalServerErrorMessage should be sent when there is an [InternalServerErr]
	InternalServerErrorMessage = "Internal Server Error"

	// UnauthorizedErrorMessage should be sent when there is an [UnauthorizedErr]
	UnauthorizedErrorMessage = "Unauthorized"

	// ConflictErrorMessage should be sent where there was an [ConflictErr]
	ConflictErrorMessage = "Conflict"

	// BadRequestErrorMessage when http body couldn't be parsed.
	BadRequestErrorMessage = "Bad Request"

	// NotFoundErrorMessage should be sent when there is an [NotFoundErr]
	NotFoundErrorMessage = "Not Found"
)
