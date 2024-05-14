package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User!"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Users!"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting User!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User!"))
}