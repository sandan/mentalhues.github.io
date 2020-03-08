package main

import(
  "time"
)

// An interface for REST resources in our application
type Resource interface{
  fetch(id uint64) (error)
  create() (error)
  update() (error)
  del() (error)
}

type User struct{
  Id       uint64
  Uuid     string
  Name     string
  Email    string
  Password string
  CreatedAt time.Time
}


func (u *User) fetch(d uint64) error {
  return nil
}

func (u *User) create() error {
  return nil
}

func (u *User) update() error {

  return nil
}

func (u *User) del() error {

  return nil
}

type Hue struct{
  Id      uint64
  Uuid    string
  Body    string // html or delta
  Title   string // optional title
  Images  string
  UserId  uint64 // Id of the User
  CreatedAt time.Time
}

type Hues struct{
  Featured []Hue
}

func (h *Hue) fetch(d uint64) error {
  return nil

}

func (u *Hue) create() error {
  return nil

}

func (u *Hue) update() error {
  return nil

}

func (u *Hue) del() error {
  return nil

}
