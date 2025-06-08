package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses -> Should go in a separate file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website`
}

//fakedb

var courses []Course

//middleware or helper for this usecase

func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

//controllers -> Should go in a separate file
// Serve home route
func main() {
	fmt.Println("First API without AI ")
	r := mux.NewRouter()

	//Seeding of data

	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Om Kulkarni", Website: "LCO.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 199, Author: &Author{Fullname: "Om Kulkarni", Website: "go.dev"}})

	//Routing 
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET") 
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") 
	r.HandleFunc("/course", createOneCourse).Methods("POST") 
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT") 
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE") 
	
	//Listen to a port
	log.Fatal(http.ListenAndServe(":5000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api by learncodeonline</h1>"))
}

// Notice that this name of the function starts with a lowercase letter
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //seeding the db with the fake values
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course ")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//Loop through courses -> find the matching id -> Return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with give id ")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//What if entire body is empty ?
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	//Abt the data == {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("please send some data")
		return
	}


	//Duplicate check -> 
	//loop -> title matches with existing -> JSON Response
	input := course.CourseName
	for _, c := range courses{
		if c.CourseName == input {
			json.NewEncoder(w).Encode("The Course with the same name already exists")
			return
		}
	}

	//Generate a unique id -> Convert to a string -> append course into courses

	rn := rand.New(rand.NewSource(time.Now().UnixNano()))

	course.CourseId = strconv.Itoa(rn.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	//Find id from request
	params := mux.Vars(r)

	// Loop through the values -> Once we get the id -> remove that value -> Add new value with id
	var flag bool = false

	for index, course := range courses {
		if course.CourseId == params["id"] {
			slices.Delete(courses, index, index+1)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			flag = true
			return

		}
	}

	//Send response when id is not found

	if !flag {
		json.NewEncoder(w).Encode("requested id not found")
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop -> id -> remove

	for index, course := range courses {
		if course.CourseId == params["id"] {
			slices.Delete(courses, index, index+1)
			break
		}
	}

	_= json.NewEncoder(w).Encode("Deleted Successfully!")
}

