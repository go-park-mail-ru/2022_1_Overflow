CREATE OR REPLACE FUNCTION IsValidEmail(text) returns BOOLEAN AS
    'select $1 ~ ''^[^@\s]+@[^@\s]+(\.[^@\s]+)+$'' as result
         ' LANGUAGE sql;

CREATE OR REPLACE  FUNCTION  IsValidName(name text) returns  BOOLEAN AS $$
    begin
        if (name != regexp_replace(name, '([^A-Za-z])', '', 'g')) then
            return false;
        else
        return true;
        end if;
    end;
$$ language plpgsql;

CREATE OR REPLACE FUNCTION check_user()
RETURNS TRIGGER
AS $$
BEGIN

    if isvalidname(new.first_name) != true then
        raise exception 'incorrect first_name';
    end if;
    if isvalidname(new.last_name) != true then
        raise exception 'incorrect last_name';
    end if;

    if IsValidEmail(new.email) != true then
         raise exception 'incorrect email';
    end if;

    return new;
END;
$$ LANGUAGE PLPGSQL;

drop trigger if exists check_user on overflow.users;

CREATE TRIGGER check_user BEFORE INSERT ON overflow.users
FOR ROW EXECUTE PROCEDURE check_user();


CREATE OR REPLACE FUNCTION check_mail()
RETURNS TRIGGER
AS $$
BEGIN
    if IsValidEmail(new.sender) != true then
         raise exception 'incorrect sender';
    end if;

    if  IsValidEmail(new.addressee) != true  then
         raise exception 'incorrect addressee';
    end if;

    return new;
END;
$$ LANGUAGE PLPGSQL;

drop trigger if exists check_mail on overflow.mails;

CREATE TRIGGER check_mail BEFORE  INSERT ON overflow.mails
FOR ROW EXECUTE PROCEDURE check_mail();

insert into overflow.users(first_name, last_name, email, password)
values ('Mikhail', 'Rabinovich','animelover69@overflow.ru',  '12312312');
insert into overflow.mails(client_id, sender, addressee,theme,  text, files, date) values
(1,'animelover69@overflow.ru', 'animelover69@overflow.ru','adasd', 'pr23323', 'dropbox.ru/id1233', '01-10-2002');


Select * from overflow.users where id = 5

