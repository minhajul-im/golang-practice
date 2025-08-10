package main

import "fmt"

type User struct {
	Name string 
	Age int
}


func main() {
    u := User{Name: "Minhaj", Age: 24}


    userMap := make(map[string]*User)
	userMap["minhaj@gmail.com"] = &u
	userMap["john@gmail.com"] = &User{Name: "John", Age: 24}
	
	user, exist := userMap["minhaj@gmail.coms"];

	if  exist {

		fmt.Println(user.Name, user.Age)
	} else {
		fmt.Println("No user Found")
	}

	users := []*User{
        &User{Name: "Minhaj", Age: 24},
        &User{Name: "John", Age: 30},
    }

    i := 0

	// for  i < len(users) && users[0].Age <= 25 {
	// 	fmt.Println(i)
	// 	i++
	// }

    for {
        if i >= len(users) {
            fmt.Println("No user named Alice found")
            break
        }
        if users[i].Name == "John" {
            fmt.Printf("Found Alice at index %d\n", i)
            break
        }
        i++
    }
}