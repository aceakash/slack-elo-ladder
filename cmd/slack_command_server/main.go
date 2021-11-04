package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := ":8081"
	err := http.ListenAndServe(addr, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			log.Printf("Error parsing form from request: %v", err)
			return
		}
		for s, strings := range request.PostForm {
			fmt.Println(s, strings)
		}

		/*
			channel_name [directmessage]
			command [/chessbot]
			trigger_id [2696922552737.815361447776.535988f2fca7ceb6aad85381c476f8cd]
			team_id [TPZAMD5NU]
			channel_id [D01QQCCAPKN]
			user_id [U01QJDF6YSX]
			user_name [akash.kurdekar]
			text [help]
			api_app_id [A02KXCVMZC6]
			is_enterprise_install [false]
			response_url [https://hooks.slack.com/commands/TPZAMD5NU/2684329489299/pfXSb32y9muvXDO7UfqIVjot]
			token [xxxxxxxxxx]
			team_domain [saltpay]

		*/

		_, err = fmt.Fprintf(writer, "Hello from the server! The time is %v", time.Now())
		if err != nil {
			log.Printf("Error writing HTTP response: %v\n", err)
		}
	}))
	if err != nil {
		log.Fatalf("Error starting server on %s: %v\n", addr, err)
	}
}
