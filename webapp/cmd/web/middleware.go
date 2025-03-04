package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

type contextKey string

const contextUserKey contextKey = "user_ip"

func (app *application) ipFromContext(ctx context.Context) string {
	return ctx.Value(contextUserKey).(string)
}

func (app *application) addIPToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()

		// get the ip (as accurately as possible)
		ip, err := GetIp(r)
		if err != nil {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}
			ctx = context.WithValue(r.Context(), contextUserKey, ip)
		} else {
			ctx = context.WithValue(r.Context(), contextUserKey, ip)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetIp(r *http.Request) (string, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown", err
	}

	userIp := net.ParseIP(ip)
	if userIp == nil {
		return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}

	forward := r.Header.Get("X-Forwarded-For")
	if len(forward) > 0 {
		ip = forward
	}

	if len(ip) == 0 {
		ip = "forward"
	}

	return ip, nil
}
