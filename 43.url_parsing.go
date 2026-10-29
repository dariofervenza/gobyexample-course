package main

// https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// this url includes:
	//	  - a scheme
	//	  - a auth info
	//	  - a host
	//	  - a port
	//	  - a path
	//	  - a query params
	//	  - a query fragment
	// hence its the most complex example of a url
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// access each part
	fmt.Println("scheme: ", u.Scheme)

	// access auth info (user + password) via User attr
	user_data := u.User
	fmt.Println("All auth: ", user_data)
	fmt.Println("Only username: ", user_data.Username())
	p, _ := user_data.Password()
	fmt.Println("Only password: ", p)

	// the host contains both host:port
	fmt.Println("Entire host: ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Splitted host: ", host)
	fmt.Println("Splitted port: ", port)

	fmt.Println("path: ", u.Path)
	fmt.Println("FRAGMENT: ", u.Fragment)

	// rawquery (?) is k=v
	fmt.Println("Raw query: ", u.RawQuery)

	// use parsekey to get a map of query key: value
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
