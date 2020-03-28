
drop table posts;
drop table hues;
drop table threads;
drop table sessions;
drop table users;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  avatar     varchar(255),
  created_at timestamp not null   
);
-- Set the system wide anonymous user to allow anonymous posts/hues
insert into users values(0, '', 'Anonymous', '', '', null, now());

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

--Wall threads
create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  content    text,
  user_id    integer references users(id),
  created_at timestamp not null       
);
--Posts per thread
create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id    integer references threads(id),
  created_at timestamp not null  
);
--table for the hues page
create table hues (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  title      varchar(64),
  featured   boolean,
  user_id    integer references users(id),
  created_at timestamp not null       
);
