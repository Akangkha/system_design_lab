package main

import "log"
//interface : abstraction what has to be done
type UserRepository interface {
	Save(user People) error
}


//Data model
type People struct{
	Name string
}

//implementation of interface

type PostgresUserRepository struct{}
type MongoUserRepository struct{}


func (p PostgresUserRepository) Save(user People) error {
	log.Println("Saving user to Postgres database:", user.Name)
	return nil
}

func(m MongoUserRepository) Save(user People) error {
	log.Println("Saving user to MongoDB database:", user.Name)
	return nil
}



//business logic
func RegisterUser(repo UserRepository, user People) error {
	return repo.Save(user)
}

func main() {
	user := People{Name: "akangkha"}
	postgresRepo := PostgresUserRepository{}
	mongoRepo := MongoUserRepository{}
	err := RegisterUser(postgresRepo, user)
	if err != nil {
		log.Println("Error registering user in Postgres:", err)
	}
	err = RegisterUser(mongoRepo, user)
	if err != nil {
		log.Println("Error registering user in MongoDB:", err)
	}
}




