package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kirill944/Go_Yandex_Lyceum/pkg/calculation"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func (a *Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("application was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Req struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	var resp Req
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal server error")
		fmt.Fprint(w, "\"error\": \"Internal server error\"")
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Internal server error")
		fmt.Fprint(w, "\"error\": \"Internal server error\"")
		return
	}
	ans, err := calculation.Calc(resp.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println("Expression is not valid")
		fmt.Fprint(w, "\"error\": \"Expression is not valid\"")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println("Good result")
		fmt.Fprint(w, "\"result\": \"", ans, "\"")
		return
	}
}

func (a *Application) RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalcHandler)
	log.Println("running server on port " + a.config.Addr)
	http.ListenAndServe(":8080", mux)
	return nil
}
