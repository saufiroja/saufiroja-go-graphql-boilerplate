create table posts (
  id varchar(255) primary key,
  title varchar(255) not null,
  content text not null,
  created_at int not null,
  updated_at int
);