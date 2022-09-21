alter table stores drop column merchant_user_id text;

CREATE TABLE merchants (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

alter table stores add column merchant_id bigint;

alter table stores
	drop constraint stores_merchant_id_fkey
		foreign key (merchant_id) references merchants
			on delete cascade;
