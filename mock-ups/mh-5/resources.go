package main

import(
)

// An interface for REST resources in our application
type Resource interface{
  fetch(id uint64) (error)
  create() (error)
  update() (error)
  del() (error)
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

func (h *Post) fetch(d uint64) error {
  return nil

}

func (u *Post) create() error {
  return nil

}

func (u *Post) update() error {
  return nil

}

func (u *Post) del() error {
  return nil
}
