package main

/*
  Definitions for objects mapped to tables
 */

import (
    "time"
)

// User resource
// Maps to the Users table
type User struct {
  Id       uint64
  Uuid     string
  Name     string
  Email    string
  Password string
  CreatedAt time.Time
}

// Maps to the Sessions table
type Session struct {
    Id        int
    Uuid      string
    Email     string
    UserId    int
    CreatedAt time.Time
}

// Thread resource
// Maps to the Thread table
type Thread struct {
    Id        int
    Uuid      string
    Topic     string
    UserId    int
    CreatedAt time.Time
}

// Post resource
// Maps to the Post table
type Post struct {
    Id        int
    Uuid      string
    Body      string
    UserId    int
    ThreadId  int
    CreatedAt time.Time
}

// Maps to the Hue table
type Hue struct {
  Id         uint64
  Uuid       string
  Body       string // html or delta
  Title      string // optional title
  Featured   bool
  Images     []string // list of references to images
  UserId     uint64 // Id of the User
  CreatedAt  time.Time
}

type Hues struct {
  Featured []Hue
}
