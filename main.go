package main

import "github.com/Chaebin-Park/Golang/url"

// func checkError(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func main() {
	urls := []string {
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	url.GoRoutineHitURL(urls)
	// url.HitURL(urls)

	// account := accounts.NewAccount("nico")
	// account.Deposit(1000)

	// account.ChangeOwner("chaebin")
	// fmt.Println(account.Owner(), account.Balance())
	// fmt.Println(account)

	// dictionary := mydict.Dictionary{}
	// dictionary["hello"] = "hello"

	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	// err := dictionary.Add("hello", "Greeting")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// definition, _ := dictionary.Search("hello")
	// fmt.Println(definition)

	// err2 := dictionary.Add("hello", "Greet")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	
}