package middlewares

import (
	"OverflowBackend/internal/session"
	"OverflowBackend/pkg"
	"OverflowBackend/proto/attach_proto"
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var allowedPaths = []string{
	"/signin",
	"/signup",
	"/swagger",
}

var minioPath = "/minio-storage"

var attach attach_proto.AttachClient

func CheckLogin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//проверка доступа к файлам вложений
		if strings.Contains(r.URL.Path, minioPath) && !CheckFilePermission(r) {
			log.Info("unauthorized")
			pkg.WriteJsonErrFull(w, &pkg.UNAUTHORIZED_ERR)
			return
		}

		allowed := false
		for _, path := range allowedPaths {
			allowed = allowed || strings.Contains(r.URL.Path, path)
		}

		if !allowed && !session.Manager.IsLoggedIn(r) {
			log.Info("unauthorized")
			pkg.WriteJsonErrFull(w, &pkg.UNAUTHORIZED_ERR)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func CheckFilePermission(r *http.Request) bool {
	if !session.Manager.IsLoggedIn(r) {
		return false
	}
	sess, err := session.Manager.GetData(r)
	if err != nil {
		log.Warning(err)
		return false
	}
	access, err := attach.CheckAttachPermission(context.Background(), &attach_proto.AttachPermissionRequest{
		Username: sess.Username,
		FileUrl:  r.URL.Path,
	})
	if err != nil {
		log.Warning(err)
		return false
	}
	return access.Access
}
