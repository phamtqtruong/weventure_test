create table users (
  id varchar(255) not null primary key,
  password varchar(255) not null,
  max_todo integer default 5 not null
);

create table tasks (
  id integer not null primary key auto_increment,
  content text not null,
  user_id varchar(255) not null,
  created_date varchar(255) not null,
  progress integer not null default 1, -- 1: open, 2: doing, 3: resolved, 4: overdue
  assigner varchar(255) not null,
  assignee varchar(255),
  due_date varchar(255),
  constraint tasks_user_id_fk foreign key (user_id) references users(id),
  constraint tasks_assigner_fk foreign key (assigner) references users(id),
  constraint tasks_assignee_fk foreign key (assignee) references users(id)
);

insert into users (id, password) values ('firstUser', 'example');

insert into tasks (content, user_id, created_date, assigner, assignee, due_date) values ('content', 'firstUser', '2021-03-04', 'firstUser', 'firstUser', '2021-03-05');
