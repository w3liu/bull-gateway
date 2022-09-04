package code

import "net/http"

func init() {
	register(Success, http.StatusOK, "Success")
	register(ErrUnknown, http.StatusInternalServerError, "Internal server error")
	register(ErrBind, http.StatusBadRequest, "Error occurred while binding the request body to the struct")
	register(ErrUnauthorized, http.StatusUnauthorized, "Unauthorized or Token invalid")
	register(ErrForbidden, http.StatusForbidden, "Access to this resource or interface is forbidden")
	register(ErrNotFound, http.StatusNotFound, "Page not found")
}
