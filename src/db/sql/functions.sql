DROP FUNCTION getincomemails(integer);
drop function if exists getIncomeMails();
create or replace function getIncomeMails(find_id integer)
returns table
        (
            sender varchar(234),
            theme varchar(30),
            text text,
            files varchar(20),
            date date
        ) as $$
declare
    client_email varchar(234);
begin
    select  email into client_email from overflow.users where id = find_id;
    return query select overflow.mails.sender, overflow.mails.theme, overflow.mails.text, overflow.mails.files, overflow.mails.date
    from overflow.mails
    where find_id = overflow.mails.client_id and overflow.mails.addressee = client_email;
end;
$$ language PLPGSQL;

SELECT *
FROM getIncomeMails(1);


drop function if exists getOutcomeMails();
create or replace function getOutcomeMails(find_id integer)
returns table
        (
            sender varchar(234),
            theme varchar(30),
            text text,
            files varchar(20),
            date date
        ) as $$
declare
    client_email varchar(234);
begin
    select  email into client_email from overflow.users where id = find_id;
    return query select overflow.mails.sender, overflow.mails.theme, overflow.mails.text, overflow.mails.files, overflow.mails.date
    from overflow.mails
    where find_id = overflow.mails.client_id and overflow.mails.sender = client_email;
end;
$$ language PLPGSQL;

SELECT *
FROM getOutcomeMails(1);

-- insert into overflow.users(first_name, last_name, password, email) values ('alex', 'alex', '123', 'alex@over.com');
--
insert into overflow.mails(client_id, sender, addressee, theme,  text, files, date) values(9, 'alex123@over.com', 'alex@over.com', 'Тема сообщения ', 'содержимое письма тут', '123', '2002-10-01')