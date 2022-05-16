
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
    sender varchar(45),
    addressee varchar(45),
    date timestamp not null,
    theme varchar(100) DEFAULT null,
    text text DEFAULT null,
    files varchar(100) DEFAULT null,
    read bool DEFAULT FALSE,
    foreign key (sender) references overflow.users(username) on delete set null,
    foreign key (addressee) references overflow.users(username) on delete set null
);

/* для папок */
CREATE TABLE overflow.folders (
  id serial not null primary key,
  name varchar(50) not null,
  user_id int not null,
  created_at timestamp not null DEFAULT NOW(),
  foreign key(user_id) references overflow.users(id) on delete cascade
);

/* для вложений */
CREATE TABLE overflow.attaches (
   mail_id serial not null references overflow.mails(id) on delete cascade,
   filename varchar(100) not null
);

CREATE OR REPLACE FUNCTION trg_ab_upbef_nulldel()
  RETURNS trigger
  LANGUAGE plpgsql AS
$func$
BEGIN
   DELETE FROM overflow.mails WHERE id = NEW.id;
   RETURN NULL;  -- to cancel UPDATE
END
$func$;

CREATE TRIGGER upbef_null2del
BEFORE UPDATE OF sender, addressee ON overflow.mails
FOR EACH ROW
WHEN (NEW.sender IS NULL AND NEW.addressee IS NULL)
EXECUTE PROCEDURE trg_ab_upbef_nulldel();

/* связь многие ко многим вида папка-письмо */
CREATE TABLE overflow.folder_to_mail (
  id SERIAL not null PRIMARY KEY,
  folder_id INTEGER NOT NULL,
  mail_id INTEGER NOT NULL,
  only_folder bool NOT NULL,
  /* удалить запись, если какой либо из ключей был удален */
  FOREIGN KEY ("folder_id") REFERENCES overflow.folders(id) on delete cascade,
  FOREIGN KEY ("mail_id") REFERENCES overflow.mails(id) on delete cascade 
);

CREATE UNIQUE INDEX "UI_folder_to_mail_mail_id_folder_id"
  ON overflow.folder_to_mail
  USING btree ("mail_id", "folder_id");