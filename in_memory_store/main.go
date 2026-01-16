package main

func main(){
	store := INewInMemoryStore()
	store.Put("name", "akangkha")
	if value, exists := store.Get("name"); exists {
		println("Got value:", value)
	} else {
		println("Key not found")
	}
	store.Delete("name")
	if _, exists := store.Get("name"); !exists {
		println("Key successfully deleted")
	}

}