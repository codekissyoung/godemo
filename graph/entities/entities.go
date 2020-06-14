package entities

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	Name       string
	Email      string
	ext        int
	privileged bool
}

type Admin struct {
	user
	Level string
}

func (u user) notify() {
	fmt.Println("Sending user email to ", u.Name, u.Email)
}
func (u *user) changeEmail(email string) {
	u.Email = email
}

func (a Admin) notify() {
	fmt.Println("Sending user email to admin : ", a.Name, a.Email)
}

func SendNotification(n notifier) {
	n.notify()
}
