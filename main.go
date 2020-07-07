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
	"time"
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
	Order	string
}

type questHandlers2 struct{
	sync.Mutex
	store map[string]Quest2
}

func (q *questHandlers2) getQuestion(w http.ResponseWriter, r *http.Request){ 
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
	var stock = rand.Intn(20)
	var duration int = 1
    	bodyString := string(bodyBytes)
	dec := json.NewDecoder(strings.NewReader(bodyString))

	    for {
		var quest2 Quest2

		err := dec.Decode(&quest2)
		fmt.Println(err)
		if err == io.EOF {
		    // all done
		    break
		}
		if err != nil {
		    log.Fatal(err)
		}
		i, err := strconv.Atoi(quest2.Order)
		    if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("bullet need numeric, error '%s'",err)))
			return
		    }
		w.Write([]byte(	fmt.Sprintf("%v left in stock\n", stock)))		
		stock-=i
		if(stock <= 1){
		time.Sleep(time.Duration(duration) * time.Second)
		stock++
		w.Write([]byte(	fmt.Sprintf("factory need to slept for %v seconds\n", duration)))		
		}
	    }

	w.Header().Add("content-type","application/json")
	w.WriteHeader(http.StatusOK)
}

func newQuestion2() *questHandlers2{
	return &questHandlers2{
		store: map[string]Quest2{
		},
	}
}

func newQuestion3(w http.ResponseWriter, r *http.Request) {
	var angka = [6][8]int{
		{1,1,1,1,1,1,1,1},
		{1,0,0,0,0,0,0,1},
		{1,0,1,1,1,0,0,1},
		{1,0,0,0,1,0,1,1},
		{1,2,1,0,0,0,0,1},
		{1,1,1,1,1,1,1,1},}
		result := 0

		var x = 4
		var y = 1
		for i := 1; i < 5; i++ {
		if angka[x-i][y] == 0 {
		for j := 1; j < 5; j++ {
		if angka[x-i][y+j]== 0 {
		for k := 1; k < 5; k++ {
		if angka[x-i+k][y+j]== 0 {	
		w.Write([]byte(	fmt.Sprintf("titik yang menjadi kemungkinan lokasi kunci rumah Joni adalah %v , %v \n", x-i+k,y+j)))	
		result++;
		}else{
		break
		}
		}
		}else{
		break
		}
		}
		}else{
		break
		}
		}	
	w.Write([]byte(	fmt.Sprintf(" banyak titik yang menjadi kemungkinan lokasi kunci rumah Joni adalah %v \n", result)))

	w.WriteHeader(http.StatusOK)
	
}

func main() {
	question1 := newQuestion1()
	question2 := newQuestion2()
	http.HandleFunc("/question1",question1.getQuestion)
	http.HandleFunc("/question2",question2.getQuestion)
	http.HandleFunc("/question3",newQuestion3)
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}
