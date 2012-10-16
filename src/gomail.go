package main

import (
	"fmt"
	"net/smtp" 
	"strings"
	"flag"
	"os"
)

func main()  {

	var smtp_server  = flag.String( "server", "", "the server to use smtp (example smtp.gmail.com)");
	var smtp_port  = flag.String( "port", "", "the server port to use smtp (example 587)");
    var user = flag.String( "u", "", "the user for the server");
	var pass = flag.String( "p", "", "the password for the user");
	var tos = flag.String( "t", "", "comma seperated list of to emails (example: jo@jo.com,who@who.org");
	var subject = flag.String( "s", "", "the subject for the email");
	var body = flag.String( "b", "", "the body for the email");

	flag.Parse()

	if *smtp_server == ""  {

		fmt.Println("you didn't specify a server.  do so with -server");
		os.Exit(1);

	}

	if *smtp_port == ""  {

		fmt.Println("you didn't specify a server port.  do so with -port");
		os.Exit(1);
	}



	if *user == ""  {

		fmt.Println("you didn't specify a user.  do so with -u");
		os.Exit(1);
	}

	if *pass == ""  {

		fmt.Println("you didn't specify a password for the user.  do so with -p");
		os.Exit(1);
	}

	if *tos == ""  {

		fmt.Println("you didn't specify at least one user.  do so with -t");
		os.Exit(1);
	}

	toArr := strings.Split(*tos, ",");



	auth := smtp.PlainAuth("", *user, *pass, *smtp_server);
	err := smtp.SendMail(*smtp_server + ":" + *smtp_port,
						 auth,
						 *user,
						 toArr,
						 []byte("subject:"+ *subject +"\n" + *body),
	)

	if err != nil {

		fmt.Println(err); 
	}

	fmt.Println("email sent sucessfully");
}
