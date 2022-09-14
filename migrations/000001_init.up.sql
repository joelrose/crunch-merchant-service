CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    firebase_id TEXT UNIQUE NOT NULL,
    language_code TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE merchants (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE stores (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    address TEXT NOT NULL,
    average_pickup_time int,
    average_review float,
    review_count int,
    google_maps_link TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    stripe_account_id TEXT,
    stripe_account_status int,
    fee float NOT NULL,
    is_open BOOLEAN NOT NULL DEFAULT FALSE,
    image_url TEXT NOT NULL,
    merchant_id BIGINT NOT NULL,
    FOREIGN KEY (merchant_id) REFERENCES merchants (id) ON DELETE CASCADE
);

CREATE TABLE store_opening_hours (
    id BIGSERIAL PRIMARY KEY,
    day_of_week int NOT NULL,
    start_timestamp int NOT NULL,
    end_timestamp int NOT NULL,
    store_id uuid NOT NULL,
    FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
);

CREATE TABLE deliverect_channels (
    status int NOT NULL,
    store_id uuid NOT NULL,
    deliverect_link_id TEXT NOT NULL,
    location_id TEXT NOT NULL,
    PRIMARY KEY (store_id, deliverect_link_id),
    FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
);

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    status int NOT NULL,
    estimated_pickup_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    price int NOT NULL,
    stripe_order_id TEXT NOT NULL,
    is_paid BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    store_id uuid NOT NULL,
    user_id bigint NOT NULL,
    FOREIGN KEY (store_id) REFERENCES stores (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE order_items (
    id BIGSERIAL PRIMARY KEY,
    plu TEXT NOT NULL,
    name TEXT NOT NULL,
    price int NOT NULL,
    quantity int NOT NULL,
    order_id bigint NOT NULL,
    parent_id bigint,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES order_items (id) ON DELETE CASCADE
);

CREATE TABLE menu_categories (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT NOT NULL,
    sort_order int NOT NULL,
    store_id uuid NOT NULL,
    FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
);

CREATE TABLE menu_product (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    plu TEXT NOT NULL,
    price int NOT NULL,
    description TEXT NOT NULL,
    snoozed BOOLEAN NOT NULL DEFAULT FALSE,
    tax int,
    image_url TEXT,
    max int,
    min int,
    multiply int,
    product_type int,
    sort_order int DEFAULT 0,
    visible BOOLEAN NOT NULL DEFAULT TRUE,
    store_id uuid NOT NULL,
    FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
);

CREATE TABLE product_product_relation (
    parent_product_id bigint NOT NULL,
    child_product_id bigint NOT NULL,
    PRIMARY KEY (child_product_id, parent_product_id),
    FOREIGN KEY (parent_product_id) REFERENCES menu_product (id) ON DELETE CASCADE,
    FOREIGN KEY (child_product_id) REFERENCES menu_product (id) ON DELETE CASCADE
);

CREATE TABLE category_product_relation (
    menu_category_id bigint NOT NULL,
    menu_product_id bigint NOT NULL,
    PRIMARY KEY (menu_product_id, menu_category_id),
    FOREIGN KEY (menu_category_id) REFERENCES menu_categories (id) ON DELETE CASCADE,
    FOREIGN KEY (menu_product_id) REFERENCES menu_product (id) ON DELETE CASCADE
);