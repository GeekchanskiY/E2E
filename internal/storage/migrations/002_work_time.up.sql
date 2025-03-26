begin;

create table user_work (
    id serial primary key,
    name varchar(255),
    hourly_rate float null,
    worker int references users(id),

    unique (name, worker)
);

create table work_time (
  id serial primary key,
  work int references user_work(id) not null,
  start_time timestamp not null default current_timestamp,
  end_time timestamp null default null
);

end;
