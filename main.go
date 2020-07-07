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

type Quest2 struct{
	tes	string	`json:"name"`
}

type questHandlers2 struct{
	sync.Mutex
	store map[string]Quest2
}

func (q *questHandlers2) getQuestion(w http.ResponseWriter, r *http.Request){ 
	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}

func newQuestion2() *questHandlers2{
	return &questHandlers2{
		store: map[string]Quest2{
		},
	}
}

func main() {
	question1 := newQuestion1()
	question2 := newQuestion2()
	http.HandleFunc("/question1",question1.getQuestion)
	http.HandleFunc("/question2",question2.getQuestion)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}
