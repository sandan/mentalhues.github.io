package main

import(
  "crypto/sha1"
  "math/rand"
  "net/http"
  "strings"
  "errors"
  "time"
  "log"
  "fmt"
)

/* function mappings to use in templates */
func getRandomColor(i int) string{

  var letters = []string{"0", "1", "2", "3", "4", "5", "6", "7",
                         "8", "9", "A", "B", "C", "D", "E", "F"}
  color := "#"

  rand.Seed(int64(i) + time.Now().UnixNano())

  for x := 0; x < 6; x++ {
    color += letters[rand.Intn(16)];
  }
  return color
}

func lower(s string) string{
  return strings.ToLower(s)
}

/* session utils */
// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
    u := new([16]byte)
    _, err := rand.Read(u[:])
    if err != nil {
        log.Fatalln("Cannot generate UUID", err)
    }

    // 0x40 is reserved variant from RFC 4122
    u[8] = (u[8] | 0x40) & 0x7F
    // Set the four most significant bits (bits 12 through 15) of the
    // time_hi_and_version field to the 4-bit version number.
    u[6] = (u[6] & 0xF) | (0x4 << 4)
    uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
    return
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess Session, err error) {
    cookie, err := request.Cookie("my_cookie")
    if err == nil {
        sess = Session{Uuid: cookie.Value}
        if ok, _ := sess.Check(); !ok {
            err = errors.New("Invalid session")
        }
    }
    return
}

// hash plaintext with SHA-1
// not good, see : https://blog.mozilla.org/webdev/2012/06/08/lets-talk-about-password-storage/
func Encrypt(plaintext string) (cryptext string) {
    cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
    return
}

/* Type utils */
// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
    rows, err := Db.Query("SELECT id, uuid, content, user_id, created_at FROM threads ORDER BY created_at DESC")
    if err != nil {
        return
    }
    for rows.Next() {
        conv := Thread{}
        if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
            return
        }
        threads = append(threads, conv)
    }
    rows.Close()
    return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
    conv = Thread{}
    err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
        Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
    statement := "delete from sessions"
    _, err = Db.Exec(statement)
    return
}

// Delete all users from database
func UserDeleteAll() (err error) {
    statement := "delete from users"
    _, err = Db.Exec(statement)
    return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
    rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
    if err != nil {
        return
    }
    for rows.Next() {
        user := User{}
        if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
            return
        }
        users = append(users, user)
    }
    rows.Close()
    return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}
