package controllers

type user struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUser() {

	var user user

	user.Id = "1"
	user.Name = "Gustavo"
	user.Age = 18
}
