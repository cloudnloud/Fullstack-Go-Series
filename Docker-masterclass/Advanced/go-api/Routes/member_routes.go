package Routes

import (
	"github.com/gorilla/mux"
	"github.com/newlinedeveloper/go-api/Controllers"
)

func MemberRoutes(router *mux.Router) {
	router.HandleFunc("/member", Controllers.CreateMember()).Methods("POST")
	router.HandleFunc("/members", Controllers.GetAllMembers()).Methods("GET")
	router.HandleFunc("/member/{id}", Controllers.GetMember()).Methods("GET")
	router.HandleFunc("/update-member/{id}", Controllers.UpdateMember()).Methods("PUT")
	router.HandleFunc("/delete-member/{id}", Controllers.DeleteMember()).Methods("DELETE")
}