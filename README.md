# smsir-golang
this module is developed for
https://sms.ir
# How to Use
this package for now contains only sending messages
not getting reports from the api or change setting

# Bulk send
for bulk send, the only thing you need to do, is to make a slice of your phone numbers, and make an object of Bulk struct
```go
package main

import (
	"fmt"
	"github.com/Aliizi83/smsir-golang/sendsms/bulk"
	"time"
)

func main() {
	phones := []string{
		"phone number 1",
		"phone number 2",
		"phone number 3",
	}

	var t *time.Time

	b := bulk.NewBulk("your-api-key", "line number", "message text", phones, t)
	resp, err := b.Send()
	if err != nil {
		// do something
	}
	fmt.Println(resp)
}


```

for schedule sending of bulk messages you have to make a pointer of time.Time object and give it to the last parameter of the Send method

# Delete Schedule Sending

```go
package main

import (
	"fmt"
	"github.com/Aliizi83/smsir-golang/sendsms/delete"
)

func main() {
	response, err := delete.Delete("your-api-key", "pack_id")
	if err != nil {
		// do something
	}
	fmt.Println(response)
}
```

# Sending Like To Like messages

```go
package main

import (
	"fmt"
	"github.com/Aliizi83/smsir-golang/sendsms/likeTolike"
)

func main() {
	
	messages := []string{
		"first message",
		"second message",
		//...
	}
	
	phones := []string{
		"first phone",
		"second phone",
		//...
	}
	
	l := likeTolike.New("your-api-key", "line_number", messages, phones)
	resp, err := l.Send()
	if	err != nil{
		//do something
		return
	}
	fmt.Println(resp)
}
```
in this part, length of the phones and messages slice must be equal

# Verify

for sending verification, we have to structs to use, 
the first one, is parameter struct
that it has two fields, Name and Value,
which Name is the name of your parameter in your template in the smsir panel
and Value is the value of that parameter

```go
package main

import (
	"fmt"
	"github.com/Aliizi83/smsir-golang/sendsms/verify"
)

func main() {
	var parameters []verify.Parameter
	
	param1 := verify.NewParam("name1", "value1")
	param2 := verify.NewParam("name2", "value 2")
	//...
	
	parameters = append(parameters, *param1)
	parameters = append(parameters, *param2)
	
	v := verify.NewVerify("your-api-key", "the mobile you whant to send", "your template id", parameters)
	
	resp, err := v.Send()
	if err != nil{
		// do something
	}
	fmt.Println(resp)
}
```