package main
import(
  "time"
)

/*
  API for each type that is mapped to tables in types.go
 */

/* User events */
// Create a new user, save user info into the database
func (user *User) Create() (err error) {
    // Postgres does not automatically return the last insert id, because it would be wrong to assume
    // you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
    // information from postgres.
    statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    // use QueryRow to return a row and scan the returned id into the User struct
    err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
    return
}

// Delete user from database
func (user *User) Delete() (err error) {
    statement := "delete from users where id = $1"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Id)
    return
}

// Update user information in the database
func (user *User) Update() (err error) {
    statement := "update users set name = $2, email = $3 where id = $1"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Id, user.Name, user.Email)
    return
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
    statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
    session = Session{}
    err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Id).
        Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    return
}

// Create a new thread from User
func (user *User) CreateThread(topic string) (conv Thread, err error) {
    statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
    statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
    return
}

// Create a new Hue
func (user *User) CreateHue(body, title  string, featured bool) (h Hue, err error) {
    s := "insert into hues (uuid, body, title, featured, user_id, created_at) values ($1, $2, $3, $4, $5, $6) returning id, uuid, body, title, featured, user_id, created_at"
    stmt, err := Db.Prepare(s)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmt.QueryRow(createUUID(), body, title, featured, user.Id, time.Now()).Scan(&h.Id, &h.Uuid, &h.Body, &h.Title, &h.Featured, &h.UserId, &h.CreatedAt)
    return
}

/* Session events */
// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
    err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
        Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    if err != nil {
        valid = false
        return
    }
    if session.Id != 0 {
        valid = true
    }
    return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
    statement := "delete from sessions where uuid = $1"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(session.Uuid)
    return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", session.UserId).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}

/* Threads */
// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
    return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
    rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
    if err != nil {
        return
    }
    for rows.Next() {
        if err = rows.Scan(&count); err != nil {
            return
        }
    }
    rows.Close()
    return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
    rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = $1", thread.Id)
    if err != nil {
        return
    }
    for rows.Next() {
        post := Post{}
        if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
            return
        }
        posts = append(posts, post)
    }
    rows.Close()
    return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
    user = User{}
    Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", thread.UserId).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}

/* Post */
func (post *Post) CreatedAtDate() string {
    return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Get the user who wrote the post
func (post *Post) User() (user User) {
    user = User{}
    Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", post.UserId).
        Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}
