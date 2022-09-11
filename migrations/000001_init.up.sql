CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    firebase_id   VARCHAR(50) UNIQUE NOT NULL,
    language_code VARCHAR(2)         NOT NULL,
    firstname     TEXT               NOT NULL,
    lastname      TEXT               NOT NULL,
    created_at    TIMESTAMPTZ        NOT NULL DEFAULT NOW()
);

CREATE TABLE merchants
(
    id         SERIAL PRIMARY KEY,
    name       TEXT        NOT NULL,
    active     BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE orders
(
    id                    SERIAL PRIMARY KEY,
    status                int         NOT NULL,
    estimated_pickup_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    price                 int         NOT NULL,
    stripeOrderId         VARCHAR(30) NOT NULL,
    isPaid                BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    merchant_id           int         NOT NULL,
    user_id               int         NOT NULL,

    FOREIGN KEY (merchant_id) REFERENCES merchants (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE order_items
(
    id        SERIAL PRIMARY KEY,
    plu       TEXT NOT NULL,
    name      TEXT NOT NULL,
    price     int  NOT NULL,
    quantity  int  NOT NULL,
    order_id  int  NOT NULL,
    parent_id int,

    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES order_items (id) ON DELETE CASCADE
);

CREATE TABLE stores
(
    id                  SERIAL PRIMARY KEY,
    name                TEXT    NOT NULL,
    address             TEXT    NOT NULL,
    averagePickupTime   int,
    averageReview       float,
    reviewCount         int,
    googleMapsLink      TEXT    NOT NULL,
    phoneNumber         TEXT    NOT NULL,
    stripeAccountId     TEXT,
    stripeAccountStatus int,
    fee                 float   NOT NULL,
    isOpen              BOOLEAN NOT NULL DEFAULT FALSE,
    menuId              int
);

CREATE TABLE store_opening_hours
(
    id              SERIAL PRIMARY KEY,
    day_of_week     int NOT NULL,
    start_timestamp int NOT NULL,
    end_timestamp   int NOT NULL,
    store_id        int NOT NULL,

    FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
);

CREATE TABLE channels
(
    id                         SERIAL PRIMARY KEY,
    type                       int NOT NULL,
    merchant_id                int NOT NULL,
    store_id                   int NOT NULL,
    deliverect_channel_link_id TEXT,
    deliverect_location_id     TEXT,

    FOREIGN KEY (merchant_id) REFERENCES merchants (id) ON DELETE CASCADE,
    FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
);

CREATE TABLE menus
(
    id          SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    name        TEXT NOT NULL,
    image_url   TEXT NOT NULL
);

CREATE TABLE menu_categories
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    description TEXT NOT NULL,
    image_url   TEXT NOT NULL,
    sort_order  int  NOT NULL,
    menu_id     int  NOT NULL,
    FOREIGN KEY (menu_id) REFERENCES menus (id) ON DELETE CASCADE
);

CREATE TABLE menu_product
(
    id                SERIAL PRIMARY KEY,
    name              TEXT    NOT NULL,
    plu               TEXT    NOT NULL,
    price             int     NOT NULL,
    description       TEXT    NOT NULL,
    snoozed           BOOLEAN NOT NULL DEFAULT FALSE,
    tax               int,
    image_url         TEXT,
    max               int,
    min               int,
    multiply          int,
    product_type      int,
    sort_order        int              DEFAULT 0,
    visible           BOOLEAN NOT NULL DEFAULT TRUE,
    menu_id           int,

    FOREIGN KEY (menu_id) REFERENCES menus (id) ON DELETE CASCADE
);

CREATE TABLE product_product_relation
(
    parent_product_id int NOT NULL,
    child_product_id  int NOT NULL,

    PRIMARY KEY (child_product_id, parent_product_id),
    FOREIGN KEY (parent_product_id) REFERENCES menu_product (id) ON DELETE CASCADE,
    FOREIGN KEY (child_product_id) REFERENCES menu_product (id) ON DELETE CASCADE
);

CREATE TABLE category_product_relation
(
    menu_category_id int NOT NULL,
    menu_product_id  int NOT NULL,

    PRIMARY KEY (menu_product_id, menu_category_id),
    FOREIGN KEY (menu_category_id) REFERENCES menu_categories (id) ON DELETE CASCADE,
    FOREIGN KEY (menu_product_id) REFERENCES menu_product (id) ON DELETE CASCADE
);
