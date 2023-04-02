create table users (
  id varchar(255) not null primary key,
  name varchar(255) not null,
  email varchar(255) not null,
  password varchar(255) not null,
  created_at int not null,
  updated_at int
);