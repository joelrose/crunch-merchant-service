alter table stores drop constraint stores_merchant_id_fkey;

alter table stores drop column merchant_id;

alter table stores add column merchant_user_id text;

drop table merchants;
