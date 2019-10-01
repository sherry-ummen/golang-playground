package main

func main() {
	user_info := &UserInfo{UserName: "Sherry Ummen", Age: 33, EmployeeId: 139}
	dbService := NewDatabaseServiceImp(HOST, PORT, DBNAME, COLLECTIONNAME)
	dbService.Save(user_info)
}
