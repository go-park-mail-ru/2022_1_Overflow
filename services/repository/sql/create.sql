
DROP SCHEMA IF EXISTS overflow CASCADE ;
CREATE SCHEMA overflow

CREATE TABLE overflow.users (
    id serial  not null primary key,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    password varchar(30) not null,
    username varchar(234) not null unique
);

CREATE TABLE overflow.mails (
    id serial not null primary key,
    client_id serial not null,
    sender varchar(234) not null ,
    addressee varchar(234) not null ,
    date timestamp not null,
    theme varchar(100),
    text text not null,
    files varchar(100),
    read bool DEFAULT FALSE,
    only_folder bool DEFAULT FALSE,
    foreign key (client_id) references overflow.users(id) on delete cascade
);

/* для папок */
CREATE TABLE overflow.folders (
	id serial not null primary key,
	name varchar(50) not null,
	user_id int not null,
  created_at timestamp not null DEFAULT NOW(),
	foreign key(user_id) references overflow.users(id) on delete cascade
);

/* связь многие ко многим вида папка-письмо */
CREATE TABLE overflow.folder_to_mail (
  id SERIAL not null PRIMARY KEY,
  folder_id INTEGER NOT NULL,
  mail_id INTEGER NOT NULL,
  FOREIGN KEY ("folder_id") REFERENCES overflow.folders(id) on delete cascade,
  FOREIGN KEY ("mail_id") REFERENCES overflow.mails(id)
);

CREATE UNIQUE INDEX "UI_folder_to_mail_mail_id_folder_id"
  ON overflow.folder_to_mail
  USING btree ("mail_id", "folder_id");