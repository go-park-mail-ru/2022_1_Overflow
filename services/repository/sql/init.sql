
DROP SCHEMA IF EXISTS overflow CASCADE ;
CREATE SCHEMA overflow

CREATE TABLE overflow.users (
    id serial  not null primary key,
    first_name varchar(45) not null,
    last_name varchar(45) not null,
    password varchar(45) not null,
    username varchar(45) not null unique
);

CREATE TABLE overflow.mails (
    id serial not null primary key,
    client_id serial not null,
    sender varchar(45) not null ,
    addressee varchar(45) not null ,
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


/* Функции */

drop function if exists getIncomeMails();
drop function if exists getIncomeMails(integer);
create or replace function getIncomeMails(find_id integer, limit_num integer, offset_num integer)
returns table
        (
            sender varchar(234),
            theme varchar(30),
            text text,
            files varchar(20),
            date timestamp,
            read bool,
            id int
        ) as $$
declare
    client_user varchar(234);
begin
    select  username into client_user from overflow.users where overflow.users.id = find_id;
    return query select overflow.mails.sender, overflow.mails.theme, overflow.mails.text, overflow.mails.files, overflow.mails.date, overflow.mails.read, overflow.mails.id
    from overflow.mails
    where overflow.mails.addressee = client_user and overflow.mails.only_folder = FALSE
    order by date desc offset offset_num limit limit_num;
end;
$$ language PLPGSQL;


drop function if exists getOutcomeMails();
drop function if exists getOutcomemails(integer);
create or replace function getOutcomeMails(find_id integer, limit_num integer, offset_num integer)
returns table
        (
            addressee varchar(234),
            theme varchar(30),
            text text,
            files varchar(20),
            date timestamp,
            id int
        ) as $$
declare
    client_user varchar(234);
begin
    select  username into client_user from overflow.users where overflow.users.id = find_id;
    return query select overflow.mails.addressee, overflow.mails.theme, overflow.mails.text, overflow.mails.files, overflow.mails.date, overflow.mails.id
    from overflow.mails
    where find_id = overflow.mails.client_id and overflow.mails.sender = client_user and overflow.mails.only_folder = FALSE
    order by date desc offset offset_num limit limit_num;
end;
$$ language PLPGSQL;
