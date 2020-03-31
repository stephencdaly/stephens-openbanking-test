CREATE TABLE charges (
  id serial primary key,
  external_id char(26),
  amount bigint not null,
  reference varchar(255) not null,
  description varchar(255) not null,
  created_date timestamp without timezone not null,
  return_url text not null,
  status VARCHAR(50)
);