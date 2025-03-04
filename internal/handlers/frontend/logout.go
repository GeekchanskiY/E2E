package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug(
		"frontend.logout.handler",
		zap.String("event", "got request"),
	)

	authCookie := http.Cookie{
		Name:     "user",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	saltCookie := http.Cookie{
		Name:     "salt",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &authCookie)
	http.SetCookie(w, &saltCookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	return

}
