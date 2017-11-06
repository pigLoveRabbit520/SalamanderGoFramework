package middlewares

import (
	"net/http"

	"github.com/relaxgo/tangram/param"

	"bitbucket.org/reewoow_web/jiyin/components/errs"
	"bitbucket.org/reewoow_web/jiyin/controls"
	"bitbucket.org/reewoow_web/jiyin/models"
	"bitbucket.org/reewoow_web/jiyin/utils"
)

func UserAuthCheck(check bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if !check {
				next.ServeHTTP(w, r)
				return
			}
			p := controls.NewRequestValue(r)
			userid := param.Int(p, "userid")
			if userid == 0 {
				cookie, err := r.Cookie("userid")
				if err == nil {
					userid = utils.StrToInt(cookie.Value)
				}
			}
			token := param.String(p, "token")
			if token == "" {
				cookie, err := r.Cookie("token")
				if err == nil {
					token = cookie.Value
				}
			}
			if userid == 0 || !models.UserTokenValid(userid, token) {
				controls.Respond(w, r, nil, errs.CMLoginNeed)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}

}
