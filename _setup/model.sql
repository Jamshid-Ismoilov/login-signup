--MODEL

create table login_bodies (
	login_id serial not null,
	email varchar(128),
	password varchar(64)
);

comment on table login is 'login ma''lumotlari';
