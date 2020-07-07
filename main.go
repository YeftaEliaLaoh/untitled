package main

import 
(
	"net/http"
	"sync"
)

type Quest1 struct{
	tes 	string
}

type questHandlers struct{
	sync.Mutex
	store map[string]Quest1
}

func (q *questHandlers) getQuestion(w http.ResponseWriter, r *http.Request){ 
	
	w.Write([]byte("<html><h1>tes1</h1></html>"))
}

func newQuestion1() *questHandlers{
	return &questHandlers{
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
