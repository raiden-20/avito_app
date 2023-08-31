package main

import (
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"os"
)

var cert, _ = tls.LoadX509KeyPair("/cert.pem", "/key.pem")
var client = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: os.Getenv("REDIS_PASSWORD"),
	DB:       0,
	TLSConfig: &tls.Config{
		MinVersion:         tls.VersionTLS12,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	},
})

func handleSetKeyRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			return
		}

		key := r.FormValue("key")
		val := r.FormValue("val")

		err = client.Set(key, val, 0).Err()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		fmt.Fprintf(w, "Пользователь \"%s\" добавлен с номером телефона %s", key, val)
	} else {
		fmt.Fprintf(w, "Метод запроса должен быть POST")
	}
}

func handleGetKeyRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		key := r.URL.Query().Get("key")

		val, err := client.Get(key).Result()
		if err != nil {
			fmt.Fprintf(w, "Такого пользователя нет в системе")
		} else {
			fmt.Fprintf(w, "У пользователя \"%s\" номер телефона: %s", key, val)
		}
	} else {
		fmt.Fprintf(w, "Метод запроса должен быть GET")
	}
}

func handleDeleteKeyRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		key := r.URL.Query().Get("key")
		err := client.Del(key).Err()
		if err != nil {
			fmt.Fprintf(w, "Такого пользователя нет в системе")
		} else {
			fmt.Fprintf(w, "Учетная запись %s удалена из списка", key)
		}
	} else {
		fmt.Fprintf(w, "Метод запроса должен быть DELETE")
	}
}

func handleOtherRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden) //403
}

func main() {
	http.HandleFunc("/set_key", handleSetKeyRequest)
	http.HandleFunc("/get_key", handleGetKeyRequest)
	http.HandleFunc("/del_key", handleDeleteKeyRequest)
	http.HandleFunc("/", handleOtherRequest)
	http.ListenAndServe(":8080", nil)
}
