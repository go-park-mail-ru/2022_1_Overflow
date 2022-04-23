drop function if exists getIncomeMails();
drop function if exists getIncomeMails(integer);
create or replace function getIncomeMails(find_id integer)
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
    where overflow.mails.addressee = client_user;
end;
$$ language PLPGSQL;


drop function if exists getOutcomeMails();
drop function if exists getOutcomemails(integer);
create or replace function getOutcomeMails(find_id integer)
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
    where find_id = overflow.mails.client_id and overflow.mails.sender = client_user;
end;
$$ language PLPGSQL;
