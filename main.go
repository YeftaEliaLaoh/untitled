package main

import 
(
	"net/http"
	"sync"
)

type Quest1 struct{
	tes	string	`json:"name"`
}

type questHandlers1 struct{
	sync.Mutex
	store map[string]Quest1
}

func (q *questHandlers1) getQuestion(w http.ResponseWriter, r *http.Request){ 
	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}

func newQuestion1() *questHandlers1{
	return &questHandlers1{
		store: map[string]Quest1{
		},
	}
}

func main() {
	question1 := newQuestion1()
	http.HandleFunc("/question1",question1.getQuestion)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}
