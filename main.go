package main

import 
(
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
    	"strconv"
	"sync"
)

type Quest1 struct{
	Bullet	string
}

func (q *Quest1) getQuestion(w http.ResponseWriter, r *http.Request){ 

	if r.Method != "POST"{
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
	return	
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil{
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("need Content-Type 'application/json', but got '%s'", ct)))
		return
	}
	var magazine = rand.Intn(10)
    	bodyString := string(bodyBytes)
	dec := json.NewDecoder(strings.NewReader(bodyString))

	    for {
		var quest1 Quest1

		err := dec.Decode(&quest1)
		fmt.Println(err)
		if err == io.EOF {
		    // all done
		    break
		}
		if err != nil {
		    log.Fatal(err)
		}
		i, err := strconv.Atoi(quest1.Bullet)
		    if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("bullet need numeric, error '%s'",err)))
			return
		    }
		if magazine == i{
		w.Write([]byte("magazine is full"))
		break
		} else {
		w.Write([]byte("magazine is not full\n"))
		}
	    }

	
	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}

func newQuestion1() *Quest1{
	return &Quest1{
	}
}

type Quest2 struct{
	order	string
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


type Quest3 struct{
	tes	string	`json:"name"`
}

type questHandlers3 struct{
	sync.Mutex
	store map[string]Quest3
}

func (q *questHandlers3) getQuestion(w http.ResponseWriter, r *http.Request){ 
	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}

func newQuestion3() *questHandlers3{
	return &questHandlers3{
		store: map[string]Quest3{
		},
	}
}

func main() {
	question1 := newQuestion1()
	question2 := newQuestion2()
	question3 := newQuestion3()
	http.HandleFunc("/question1",question1.getQuestion)
	http.HandleFunc("/question2",question2.getQuestion)
	http.HandleFunc("/question3",question3.getQuestion)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}
