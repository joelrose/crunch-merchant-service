/* whitelist */
insert into public.whitelist (identifier, created_at)
values  ('+491743645092', '2022-09-22 08:44:45.837476 +00:00');

/* users  */
insert into public.users (id, firebase_id, language_code, firstname, lastname, created_at)
values  (1, '97ud9dLHEtUDTe5j6GVjR3rhXro2', 'DE', 'Joel', 'Rose', '2022-09-21 13:05:49.996927 +00:00');

/* stores */
insert into public.stores (id, name, description, address, average_pickup_time, average_review, review_count, google_maps_link, phone_number, stripe_account_id, stripe_account_status, fee, is_open, image_url, merchant_user_id)
values  ('9142ac52-e5a4-4ad8-8811-240c1f389ece', 'Name', 'Description', 'Address', 5, 5, 5, 'https://google.com/maps', '0123456789', null, null, 0.5, true, 'https://resizer.staging.deliverect.com/0KGNmtuUuldFR1S3V4W1KHs_KLOQCytSTr77pZgBGbU/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvRnJpZWQtQ2hpY2tlbi0yLTItNjIyYTJjZWJkYjU5ODYwMDFlYmY1OGY3LmpwZWc=', 'auth0|6325f524bda241ddc6fca32e');

/* deliverect */
insert into public.deliverect_channels (status, store_id, deliverect_link_id, location_id)
values  (0, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 'linkId', 'location');

/* opening_hours */ 
insert into public.store_opening_hours (id, day_of_week, start_timestamp, end_timestamp, store_id)
values  ('d2ec8c56-79e0-4eba-9efc-45cef24e09e4', 0, 0, 1439, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('71199068-2595-4a7b-92aa-10192a6b8899', 3, 0, 1439, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('031fd9c9-4603-404c-9efb-944ad3de0d6e', 4, 540, 720, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('fe0947e2-4f10-4415-8419-158447b74f4a', 5, 0, 1439, '9142ac52-e5a4-4ad8-8811-240c1f389ece');

/* categories */
insert into public.menu_categories (id, name, description, image_url, sort_order, store_id)
values  ('a75a1502-f970-4400-8e1d-ce0dedaf7ea3', 'Chicken', 'Delicious chicken, cooked to perfection', 'https://resizer.staging.deliverect.com/0KGNmtuUuldFR1S3V4W1KHs_KLOQCytSTr77pZgBGbU/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvRnJpZWQtQ2hpY2tlbi0yLTItNjIyYTJjZWJkYjU5ODYwMDFlYmY1OGY3LmpwZWc=', 0, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('5e20b60a-547c-4b39-997c-ca97791d1968', 'Pizza', 'Proper Pizza', 'https://resizer.staging.deliverect.com/NGRplb9bgmesum0bP3XbTaU9ShkUytbjMQGYa2XIzzk/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY2xhc3NpY19kaWF2b2xhX3BlcHBlcm9uaS02MjIxZTE0ODA3ZjAzODAwMTk4NTkyYmQuanBn', 1, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('003933bb-3e7a-4b4c-8162-90f592486941', 'Sides', 'Choose an accompaniment', '', 2, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('9c4efe6b-01b9-4346-8dff-740561b424eb', 'Drinks', 'Classic beverages, straight out the fridge,', '', 3, '9142ac52-e5a4-4ad8-8811-240c1f389ece'),
        ('926ee1b7-92be-49d6-bec2-feae63436b73', 'Steak & Burgers', 'Something from the Grill', 'https://resizer.staging.deliverect.com/9bQOP0kZwDLLIrqM-jiyoZJJG85dMIycm3rvTCkg7qw/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvZ3JpbGwtNjIyYTJjNDhjNzE1YjQwM2IwZWI3MmJlLmpwZWc=', 4, '9142ac52-e5a4-4ad8-8811-240c1f389ece');

/* orders */
insert into public.orders (id, status, estimated_pickup_time, price, stripe_order_id, is_paid, created_at, store_id, user_id, fee)
values  (1, 1, '2022-09-21 13:10:26.456517 +00:00', 47440, 'pi_3LkSh4K5PYFPbUlO1Nt9Z7qF', false, '2022-09-21 13:10:26.462087 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (2, 1, '2022-09-22 07:14:17.968163 +00:00', 47440, 'pi_3LkjbxK5PYFPbUlO1JAcSgOy', false, '2022-09-22 07:14:17.973700 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (3, 1, '2022-09-22 07:14:22.359918 +00:00', 47440, 'pi_3Lkjc2K5PYFPbUlO1f5oGfvA', false, '2022-09-22 07:14:22.360912 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (4, 1, '2022-09-22 07:14:24.173741 +00:00', 47440, 'pi_3Lkjc4K5PYFPbUlO0ArkwAla', false, '2022-09-22 07:14:24.174450 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (5, 1, '2022-09-22 07:58:53.392894 +00:00', 6350, 'pi_3LkkJ7K5PYFPbUlO0DyQB709', false, '2022-09-22 07:58:53.400320 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (6, 1, '2022-09-22 08:02:32.404134 +00:00', 800, 'pi_3LkkMeK5PYFPbUlO0lrU36vU', false, '2022-09-22 08:02:32.402635 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (7, 1, '2022-09-22 08:13:04.299018 +00:00', 800, 'pi_3LkkWqK5PYFPbUlO11lUJSId', false, '2022-09-22 08:13:04.303708 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5),
        (8, 1, '2022-09-22 08:13:44.240305 +00:00', 800, 'pi_3LkkXUK5PYFPbUlO09QgDtjR', true, '2022-09-22 08:13:44.245166 +00:00', '9142ac52-e5a4-4ad8-8811-240c1f389ece', 1, 0.5);

/* order_items */
insert into public.order_items (id, plu, name, price, quantity, order_id, parent_id)
values  ('e3322b41-35fe-45dd-9bfc-789b4b05a3d0', '1', 'Eins', 1000, 1, 1, null),
        ('2fb09171-4a79-40dc-b266-d6ea04343da5', '3', 'Drei', 1000, 1, 1, 'e3322b41-35fe-45dd-9bfc-789b4b05a3d0'),
        ('8f0b0922-ffdb-4086-a4bd-67ccd23aa1c0', '6', 'Sechs', 8888, 5, 1, '2fb09171-4a79-40dc-b266-d6ea04343da5'),
        ('c3b1e456-30d7-4283-971b-18ddc3c65fa4', '2', 'Zwei', 1000, 1, 1, null),
        ('17033b5d-8f25-4016-80e7-de2fb68848ba', '1', 'Eins', 1000, 1, 2, null),
        ('c5ba1107-63b9-4f47-8ee3-084d498ea49d', '3', 'Drei', 1000, 1, 2, '17033b5d-8f25-4016-80e7-de2fb68848ba'),
        ('dcca7f32-81b7-4dcb-a07d-26755091915a', '6', 'Sechs', 8888, 5, 2, 'c5ba1107-63b9-4f47-8ee3-084d498ea49d'),
        ('939b8307-9df6-4f4e-88bc-e73edc638868', '2', 'Zwei', 1000, 1, 2, null),
        ('b6cf4973-dcb1-4a16-96fd-8a2327662b7b', '1', 'Eins', 1000, 1, 3, null),
        ('f72fd8cc-1455-4b6d-a893-71fc1d602d00', '3', 'Drei', 1000, 1, 3, 'b6cf4973-dcb1-4a16-96fd-8a2327662b7b'),
        ('cbe761c8-6d41-425e-a731-373661dcd1fa', '6', 'Sechs', 8888, 5, 3, 'f72fd8cc-1455-4b6d-a893-71fc1d602d00'),
        ('c6337774-d7bf-4d7e-b25e-b743ff5dc961', '2', 'Zwei', 1000, 1, 3, null),
        ('64633f72-3cd5-45d7-aee5-17b9d934abaf', '1', 'Eins', 1000, 1, 4, null),
        ('d576cebd-c1a1-48a9-b7af-7082489a78d4', '3', 'Drei', 1000, 1, 4, '64633f72-3cd5-45d7-aee5-17b9d934abaf'),
        ('0c346db6-3506-4fe6-8b93-0d33664d9480', '6', 'Sechs', 8888, 5, 4, 'd576cebd-c1a1-48a9-b7af-7082489a78d4'),
        ('d4b389ec-87fd-4e1c-a43d-dcd40bbb4d13', '2', 'Zwei', 1000, 1, 4, null),
        ('992804ad-d476-480b-9d84-087a6c4a7369', 'VAR-PROD-1-#D1#-', 'Chicken Tenders', 800, 1, 5, null),
        ('8dffb77a-4e9f-4571-b7be-5dfca7e0b388', 'VAR-1-#D1#V0#-', '3 Pieces', 0, 1, 5, '992804ad-d476-480b-9d84-087a6c4a7369'),
        ('209e1ba9-7317-4341-8adc-0065d5a8d94d', 'P-SATE-#D1#-', 'Chicken Sate', 450, 1, 5, null),
        ('577a7de7-84ec-4e2d-a694-ad4400318b1d', 'P-SATE-#D1#-', 'Chicken Sate', 450, 1, 5, null),
        ('9faa2996-c942-4fa8-88c3-8d416ba4d839', 'RICE-01-#D2#-', 'White Rice', 450, 1, 5, '577a7de7-84ec-4e2d-a694-ad4400318b1d'),
        ('f0c730db-a4ec-433a-967d-3a7b1fa45b96', 'RICE-02-#D1#-', 'Yellow Rice', 450, 1, 5, '577a7de7-84ec-4e2d-a694-ad4400318b1d'),
        ('86c53518-43da-495a-9a02-5d1ce33875c5', 'P-BURG-CHK-#D1#-', 'Chicken Burger', 800, 1, 5, null),
        ('4e47040b-c311-4322-8930-c75795738b11', 'P-BURG-VEG-#D1#-', 'Veggie Burger', 750, 1, 5, null),
        ('86109b2c-2fa7-44d7-8cea-87f6f6e918c6', 'STK-01-#D1#-', 'Delicious Steak Frites ***NEW', 2000, 1, 5, null),
        ('e34971a5-8e93-4644-98d0-19acf6a52c68', 'COOK-02-#D1#-', 'Medium Rare', 0, 1, 5, '86109b2c-2fa7-44d7-8cea-87f6f6e918c6'),
        ('b02eb42f-ec3a-4d84-a28e-c5b44a28ae21', 'COOK-03-#D1#-', 'Well Done', 0, 1, 5, '86109b2c-2fa7-44d7-8cea-87f6f6e918c6'),
        ('0fae0f36-40b6-489e-acbb-468b5c98e879', 'SI-01-#D1#-', 'Fries', 0, 1, 5, '86109b2c-2fa7-44d7-8cea-87f6f6e918c6'),
        ('4934b085-0be8-4391-85bf-8318dfb2e5f9', 'SI-02-#D1#-', 'Salad', 200, 1, 5, '86109b2c-2fa7-44d7-8cea-87f6f6e918c6'),
        ('28a2e33a-ea97-4d43-bd6d-b5b9b0eeb7d6', 'VAR-PROD-1-#D1#-', 'Chicken Tenders', 800, 1, 6, null),
        ('e8997972-6695-422d-805f-b3cefe9ab9db', 'VAR-1-#D1#V0#-', '3 Pieces', 0, 1, 6, '28a2e33a-ea97-4d43-bd6d-b5b9b0eeb7d6'),
        ('43bcb646-eae4-423a-9688-386456c885a4', 'VAR-PROD-1-#D1#-', 'Chicken Tenders', 800, 1, 7, null),
        ('327f1867-cd5a-41fb-8833-04e8aef2df29', 'VAR-1-#D1#V0#-', '3 Pieces', 0, 1, 7, '43bcb646-eae4-423a-9688-386456c885a4'),
        ('d5e09098-eb12-46f6-8a9d-e13eaa00d649', 'VAR-PROD-1-#D1#-', 'Chicken Tenders', 800, 1, 8, null),
        ('5f6f22a6-faef-415a-a0a4-836d7b165a44', 'VAR-1-#D1#V0#-', '3 Pieces', 0, 1, 8, 'd5e09098-eb12-46f6-8a9d-e13eaa00d649');

/* menu product */
insert into public.menu_product (id, name, plu, price, description, snoozed, tax, image_url, max, min, multiply, product_type, sort_order, visible, store_id, multi_max)
values  ('08cdafa3-595e-4d59-ba6e-aa1d5e024aef', 'Yellow Rice', 'RICE-02-#D1#-', 450, 'White rice with Saffron', false, 9000, 'https://resizer.staging.deliverect.com/X4k35ro8oeQ3ZLi2IvtokseerUmAF1OwS6KxivGgxiw/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvTmFzaS1rdW5pbmctMy0xNjEwMjg5NzI2OTQ2LmpwZw==', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('5e7c5f6f-c44e-41cd-90f8-4549b9ad3de8', 'Chicken Burger', 'P-BURG-CHK-#D1#-###', 0, 'Crispy coated chicken thigh, iceberg lettuce, pickles, slice of cheese & mayo, all in a toasted brioche bun.', false, 9000, 'https://resizer.staging.deliverect.com/VvD90nU-bTCKE2uoFBeCmZ-PhVAjyXZ4fu9JvXmiZTg/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY2hrYnVyZ2VyLTYyMjhjMWRjZGI1OTg2MDAxZWJmNThkZi5qcGVn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('8bb70955-dfd5-421e-a0c1-8491fa5091e7', '9 Pieces', 'VAR-3-#D1#V550#-', 550, '', false, 9000, '', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('b9e099f2-cd46-4bbb-8b07-a35cd352bc68', 'Add a Drink? (not included)', 'UPSLL-01-#D2#-', 0, '', false, 9000, '', 0, 0, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 99),
        ('b857373a-0656-49a1-a7ea-566c1a549d04', 'Medium Rare', 'COOK-02-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('3e4d8830-926b-4c44-a6bd-688866db2757', 'Diet Coke', 'DRNK-02-#D1#-', 400, 'Cola flavoured aspartame and caffeine', false, 9000, 'https://resizer.staging.deliverect.com/z-Jf1m_zG3tBblCT94q7f8UNlQWoKOrqALvYC4CVQDY/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvZGlldGNva2UtNjIyODU0Y2U4YzUwNmYwMTViZTYwMThjLmpwZWc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('1ad0a226-df3f-48fd-b078-c02f6d5f013c', 'Choose a sauce', 'MG-SAUCES-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 99),
        ('89189372-970f-4bf8-a5c2-54554268e07d', 'Chicken Sate', 'P-SATE-#D1#-', 450, 'Product with Nested Modifiers - Multimax variables - Allergens (tags)', false, 9000, 'https://resizer.staging.deliverect.com/YqtYD9yqbRnLe7xj7pFOBApLPmWJSse_G6DWTvAu9Hc/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvc2F0YXktNjIyODRlM2M4YzUwNmYwMTViZTYwMTg0LmpwZWc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('c0f6470a-5763-4798-bac7-90524f00f358', 'White Rice', 'RICE-01-#D2#-', 450, 'White coloured rice', false, 9000, 'https://resizer.staging.deliverect.com/u3xhTntUZcuuNfaoOVaFDaW9jgbfAUCdzk09ckckvWk/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvbmFzaXB1dGktMTYxMDI5MDE0MDQ5NC5qcGc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('5d3c3ed6-36aa-40eb-830a-6f6dfdd4b49b', 'Cheeseburger', 'P-BURG-CHE-#D1#-', 850, '100% beef patty, cheddar, caramelized onions, mayonnaise, pickles in a Pretzel bun', false, 9000, 'https://resizer.staging.deliverect.com/pd2c_4pm6KclLUrp4cUiAXCVxzNYPGgNybqCdAhrPyg/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY2hlZXNlYnVyZ2VyLTYyMjg2ZTI2ZGI1OTg2MDAxZWJmNThkNy5qcGc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('76a4f294-fa55-46c9-bb45-22e5be268750', 'Build your own Pizza', 'PIZZ-00-#D1#S1#-', 800, 'Build your own pizza, first topping is free!', false, 9000, 'https://resizer.staging.deliverect.com/wTku9di12lFUdycy-D8h8LEbHgew9oIf8nKhcjvEcxY/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvcGl6emEtNjIyODUyNWViMzAzZmMwM2ExNDhkZTQ2LmpwZWc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('a7bd445e-c30c-421f-b0c2-9bf0653eedb0', 'Pepperoni', 'PEPP-#D1#O0#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('a2b82fa9-d0b5-43d8-bf33-68694ce2329f', 'Chicken Burger', 'P-BURG-CHK-#D1#-', 800, 'Crispy coated chicken thigh, iceberg lettuce, pickles, slice of cheese & mayo, all in a toasted brioche bun.', false, 9000, 'https://resizer.staging.deliverect.com/VvD90nU-bTCKE2uoFBeCmZ-PhVAjyXZ4fu9JvXmiZTg/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY2hrYnVyZ2VyLTYyMjhjMWRjZGI1OTg2MDAxZWJmNThkZi5qcGVn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('2792f195-813b-4f63-937e-81605382c475', 'Cooking instructions *** HOT***', 'MOD-01-#D1#-', 0, '', false, 9000, '', 3, 0, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('4498023f-c798-4373-8228-ed3c8b0d0a7b', 'Mashed Potato', 'SI-03-#D1#-', 100, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('1ba621ab-b658-4572-baa7-5289d6254cc7', 'Egg Noodles', 'NOOD-01-#D1#-', 450, 'Egg noodles and veggies fried and tossed with a delicious sauce', false, 9000, 'https://resizer.staging.deliverect.com/e-krmZr2fbU9SSpfxgC2KhnQ4WgAJKfKMrHgddqtQDo/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvYmFtaWdvcmVuZy0xNjEwMjg5OTIyOTY5LmpwZw==', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('1c731615-6e85-4db7-ae42-d1aac7b9b85c', 'Veggie Burger', 'P-BURG-VEG-#D1#-###', 0, 'Black bean burgers with sweet potato, mushrooms, quinoa, and pecans.', false, 9000, 'https://resizer.staging.deliverect.com/OUBF0w3h7eq9b1YE1mNKMjil4dNH-4cpDjQcLr9okEE/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvdmVnZ2llYnVyZ2VyLTYyMjg2Y2JhYzcxNWI0MDNiMGViNzI5NC5qcGVn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('aad23c2a-c30f-4ccd-a866-733fbf27276d', 'Red Onion', 'RONION-#D1#O1#-', 100, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('dd056a61-0817-4a9b-85aa-7f3de6f3eba9', '3 Pieces', 'VAR-1-#D1#V0#-', 0, '', false, 9000, '', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('4e67345c-a756-413c-aabf-322f55297ebd', 'French Fries', 'P-FRS-S-#D1#-', 200, 'Plain fries from France', false, 9000, '', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('3ccf7bc6-4d56-4b95-940d-b8c31ee66631', '6 Pieces', 'VAR-2-#D1#V300#-', 300, '', false, 9000, '', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('d7227149-91ad-4dc6-985d-660514ee477f', 'Delicious Steak Frites ***NEW', 'STK-01-#D1#-', 2000, 'Basic Example Product with - Modifier groups - min/max variables - default selection - translations **** VERY ***', false, 9000, 'https://resizer.staging.deliverect.com/W0QA9MN33DOxZWfnO8vZrHvZ-3Qt-kL7TwM-leyQfoM/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvc3RlYWstNjIyODYyNTg4YzUwNmYwMTViZTYwMThlLmpwZWc=', 0, 0, 1, 1, 0, true, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('bc8d999a-4bac-496d-9628-bb81767b7319', 'Burger Combo (Drink not Included)', 'P-BRGR-#D1#-', 950, 'Combo with Bundles - Modifier Groups as Upsell', false, 9000, 'https://resizer.staging.deliverect.com/mZfsUa6kbr8sozzJv_DpMK7WGRdUz6vZSnhN-39xFsY/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy82MWIwNTE3ZGEzNjVhMjY2ODEyM2ZiZTAvbmV3LW1lbnUtaXRlbS02MjMwYWY3OGIzMDNmYzAzYTE0OGRlN2EuanBn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('4b7e6a29-bac2-43c4-bba0-061c012ed07c', 'Cheeseburger', 'P-BURG-CHE-#D1#-###', 0, '100% beef patty, cheddar, caramelized onions, mayonnaise, pickles in a Pretzel bun', false, 9000, 'https://resizer.staging.deliverect.com/pd2c_4pm6KclLUrp4cUiAXCVxzNYPGgNybqCdAhrPyg/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY2hlZXNlYnVyZ2VyLTYyMjg2ZTI2ZGI1OTg2MDAxZWJmNThkNy5qcGc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('f11b24b3-2fb9-4ff8-a538-ba03007a9ec9', 'Hot Sauce', 'SAUCE-02-#D1#-', 50, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('f85a96b3-2d4e-414a-9e27-4f65cc1db7c3', 'Seasoned Fries', 'P-FRS-L-#D1#-', 250, 'Plain fries, but a bit fancier', false, 9000, '', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('c9d67daa-858c-44bc-93a6-0c2f15dc25a6', 'Fries', 'SI-01-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('43d0185d-f05a-4fc5-a295-dca745cc40ee', 'Red Pepper', 'REDPEPP-#D1#O1#-', 100, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('6fe3bc5b-3cd6-4bd5-970a-22f537f50a29', 'Red Onion', 'RONION-#D1#O0#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('ed66370f-ce45-4061-b279-78cd86f1d3dd', 'Add a side', 'MOD-02-#D1#-', 0, '', false, 9000, '', 0, 1, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('2f54b42f-239d-4c5e-b725-e696b5843b7e', 'Sate Sauce', 'SAUCE-01-#D1#-', 50, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('3188fc15-f3c0-4210-b7a3-e04078cb037a', 'Bacon', 'BAC-#D1#O1#-', 100, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('3fee3575-a587-4be0-bc28-1793d75ca25f', 'Bacon', 'BAC-#D1#O0#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('ba7e4af2-d5f0-4d7c-bfa5-5515f1ec92f9', 'Mushroom', 'MUSH-#D1#O1#-', 100, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('62a45a69-f337-4479-ac17-8222e2abab40', 'How many pieces?', 'MG-VAR-1-#D1#-', 0, '', false, 9000, '', 1, 1, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('78ad748e-225a-4680-a6c0-37ed3da78d23', 'Veggie Burger', 'P-BURG-VEG-#D1#-', 750, 'Black bean burgers with sweet potato, mushrooms, quinoa, and pecans.', false, 9000, 'https://resizer.staging.deliverect.com/OUBF0w3h7eq9b1YE1mNKMjil4dNH-4cpDjQcLr9okEE/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvdmVnZ2llYnVyZ2VyLTYyMjg2Y2JhYzcxNWI0MDNiMGViNzI5NC5qcGVn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('8bc215da-22eb-49ed-815e-464b1da1b77f', 'Fries Selection', 'MG-FRS-#D1#-', 0, '', false, 9000, '', 1, 1, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('11d77ab4-7188-4090-bf81-feb6d260bd58', 'Mushroom', 'MUSH-#D1#O0#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('24e3b293-30f6-479d-9bef-dddd5e4a885d', 'Well Done', 'COOK-03-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('df1c29b6-8a0a-4b59-9c93-dda28cf9da8a', 'Ramen Noodles', 'NOOD-02-#D1#-', 450, 'Chinese-style wheat noodles', false, 9000, 'https://resizer.staging.deliverect.com/KZyNSHX41i_DcF9byp9bCr8wXXIQC3fl4HF3FUcAILA/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvTWlob2VuLTItMS0xNjEwMjg5ODcwMTU3LmpwZw==', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('2a1e34d3-2dc1-48eb-9010-1acf6b2677fb', 'Pepperoni', 'PEPP-#D1#O1#-', 100, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('50c2cba9-0db3-4987-9cb6-fd5b596314c1', 'Burger Selection', 'BNDL-BRG-#D1#-', 0, '', false, 9000, '', 1, 1, 1, 4, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('3eccf3e0-26ba-4e05-ade5-a48a4030321b', 'Chicken Tenders', 'VAR-PROD-1-#D1#-', 800, 'Variant prices for different sizes will show cheapaest on top level product', false, 9000, 'https://resizer.staging.deliverect.com/wK0RxpTvMLSdQq93iHIcK5K852q_1kBcZbf-0KqYB-o/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY2hpY2tlbi02MjI4NWY5MGRiNTk4NjAwMWViZjU4ZDUuanBn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('57340c98-2eff-407e-9c25-1e8638a8ef4a', 'Curly Fries', 'P-FRS-M-#D1#-', 200, 'Spiralised potatoes, fried', false, 9000, '', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('86fc70fe-6d01-48b3-b3f1-e52274547901', 'Rice Selection', 'MG-RICE-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 99),
        ('c510eb79-b4dc-495b-ac38-6f8a6a8fb25f', 'Salad', 'SI-02-#D1#-', 200, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('28dd1cc7-8a68-4f04-bdc6-0285c55b59d5', 'Red Pepper', 'REDPEPP-#D1#O0#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('63badf23-d63d-4ce5-b946-4940544fc963', 'Ginger Beer', 'DRNK-03-#D1#-', 500, 'Australia''s favourite ginger beer!', false, 9000, 'https://resizer.staging.deliverect.com/vmsR-mNeVB_uqqxNLllvd_s7tenFkD1yT7u3bc51e3M/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvZ2luZ2VyYmVlci02MjI4NTU0OGRiNTk4NjAwMWViZjU4ZDEuanBn', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('baabeddb-9be1-4195-a815-26d53cb45033', 'Coca Cola', 'DRNK-01-#D1#-', 100, 'Cola flavoured sugar and caffeine', false, 9000, 'https://resizer.staging.deliverect.com/iNQFZHRi8XpOpb-Lp4Gou6JDXeGo1rjVMpS18bJ99Qg/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvY29jYWNvbGEtNjIyODU0YTc4YzUwNmYwMTViZTYwMThhLmpwZWc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('b2a970ee-0c51-4af1-b17b-9cf8005773ef', 'Add extra toppings', 'PIZZ-TOP-#D1#S1#-', 0, '', false, 9000, '', 0, 0, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 10),
        ('a03b1b5d-f9e5-4d22-bcc4-14db8d40e591', 'Noodles Selection', 'MG-NOODLES-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 99),
        ('e559daa5-fd66-4c07-b85f-3f9c64d9d099', 'Choose your First Topping', 'FREE-TOP-#D1#S1#-', 0, '', false, 9000, '', 1, 1, 1, 3, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('e58452c9-2859-4ec5-8db8-aca088e88ce0', 'Rare **** RED', 'COOK-01-#D1#-', 0, '', false, 9000, '', 0, 0, 1, 2, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0),
        ('42cdad60-e247-4277-8227-47868253e8c7', 'The Hawaiian', 'PIZZ-01-#D1#-', 800, 'Italy''s favourite Pizza!', false, 9000, 'https://resizer.staging.deliverect.com/X4s-6fsgdWm7RXLko7ef1YW84XCbR1ZGvognct0bIWI/rt:fill/g:ce/el:0/cb:f324f1661db6463e9552f15f61752f1a/aHR0cHM6Ly9zdG9yYWdlLmdvb2dsZWFwaXMuY29tL2lrb25hLWJ1Y2tldC1zdGFnaW5nL2ltYWdlcy81ZmY2ZWUwODkzMjhjOGFlZmVlYWJlMzMvaGF3YWlpYW4tNjIyODU1YzdiMzAzZmMwM2ExNDhkZTQ4LmpwZWc=', 0, 0, 1, 1, 0, false, '9142ac52-e5a4-4ad8-8811-240c1f389ece', 0);

    /* product product relation */
insert into public.product_product_relation (parent_product_id, child_product_id)
values  ('d7227149-91ad-4dc6-985d-660514ee477f', '2792f195-813b-4f63-937e-81605382c475'),
        ('d7227149-91ad-4dc6-985d-660514ee477f', 'ed66370f-ce45-4061-b279-78cd86f1d3dd'),
        ('bc8d999a-4bac-496d-9628-bb81767b7319', '50c2cba9-0db3-4987-9cb6-fd5b596314c1'),
        ('bc8d999a-4bac-496d-9628-bb81767b7319', '8bc215da-22eb-49ed-815e-464b1da1b77f'),
        ('bc8d999a-4bac-496d-9628-bb81767b7319', 'b9e099f2-cd46-4bbb-8b07-a35cd352bc68'),
        ('ed66370f-ce45-4061-b279-78cd86f1d3dd', 'c9d67daa-858c-44bc-93a6-0c2f15dc25a6'),
        ('ed66370f-ce45-4061-b279-78cd86f1d3dd', 'c510eb79-b4dc-495b-ac38-6f8a6a8fb25f'),
        ('ed66370f-ce45-4061-b279-78cd86f1d3dd', '4498023f-c798-4373-8228-ed3c8b0d0a7b'),
        ('62a45a69-f337-4479-ac17-8222e2abab40', 'dd056a61-0817-4a9b-85aa-7f3de6f3eba9'),
        ('62a45a69-f337-4479-ac17-8222e2abab40', '3ccf7bc6-4d56-4b95-940d-b8c31ee66631'),
        ('62a45a69-f337-4479-ac17-8222e2abab40', '8bb70955-dfd5-421e-a0c1-8491fa5091e7'),
        ('8bc215da-22eb-49ed-815e-464b1da1b77f', '4e67345c-a756-413c-aabf-322f55297ebd'),
        ('8bc215da-22eb-49ed-815e-464b1da1b77f', '57340c98-2eff-407e-9c25-1e8638a8ef4a'),
        ('8bc215da-22eb-49ed-815e-464b1da1b77f', 'f85a96b3-2d4e-414a-9e27-4f65cc1db7c3'),
        ('df1c29b6-8a0a-4b59-9c93-dda28cf9da8a', '1ad0a226-df3f-48fd-b078-c02f6d5f013c'),
        ('50c2cba9-0db3-4987-9cb6-fd5b596314c1', '5e7c5f6f-c44e-41cd-90f8-4549b9ad3de8'),
        ('50c2cba9-0db3-4987-9cb6-fd5b596314c1', '4b7e6a29-bac2-43c4-bba0-061c012ed07c'),
        ('50c2cba9-0db3-4987-9cb6-fd5b596314c1', '1c731615-6e85-4db7-ae42-d1aac7b9b85c'),
        ('3eccf3e0-26ba-4e05-ade5-a48a4030321b', '62a45a69-f337-4479-ac17-8222e2abab40'),
        ('86fc70fe-6d01-48b3-b3f1-e52274547901', 'c0f6470a-5763-4798-bac7-90524f00f358'),
        ('86fc70fe-6d01-48b3-b3f1-e52274547901', '08cdafa3-595e-4d59-ba6e-aa1d5e024aef'),
        ('a03b1b5d-f9e5-4d22-bcc4-14db8d40e591', '1ba621ab-b658-4572-baa7-5289d6254cc7'),
        ('a03b1b5d-f9e5-4d22-bcc4-14db8d40e591', 'df1c29b6-8a0a-4b59-9c93-dda28cf9da8a'),
        ('e559daa5-fd66-4c07-b85f-3f9c64d9d099', 'a7bd445e-c30c-421f-b0c2-9bf0653eedb0'),
        ('e559daa5-fd66-4c07-b85f-3f9c64d9d099', '3fee3575-a587-4be0-bc28-1793d75ca25f'),
        ('e559daa5-fd66-4c07-b85f-3f9c64d9d099', '6fe3bc5b-3cd6-4bd5-970a-22f537f50a29'),
        ('e559daa5-fd66-4c07-b85f-3f9c64d9d099', '11d77ab4-7188-4090-bf81-feb6d260bd58'),
        ('e559daa5-fd66-4c07-b85f-3f9c64d9d099', '28dd1cc7-8a68-4f04-bdc6-0285c55b59d5'),
        ('b2a970ee-0c51-4af1-b17b-9cf8005773ef', '2a1e34d3-2dc1-48eb-9010-1acf6b2677fb'),
        ('b2a970ee-0c51-4af1-b17b-9cf8005773ef', '3188fc15-f3c0-4210-b7a3-e04078cb037a'),
        ('b2a970ee-0c51-4af1-b17b-9cf8005773ef', 'aad23c2a-c30f-4ccd-a866-733fbf27276d'),
        ('b2a970ee-0c51-4af1-b17b-9cf8005773ef', 'ba7e4af2-d5f0-4d7c-bfa5-5515f1ec92f9'),
        ('b2a970ee-0c51-4af1-b17b-9cf8005773ef', '43d0185d-f05a-4fc5-a295-dca745cc40ee'),
        ('b9e099f2-cd46-4bbb-8b07-a35cd352bc68', 'baabeddb-9be1-4195-a815-26d53cb45033'),
        ('b9e099f2-cd46-4bbb-8b07-a35cd352bc68', '3e4d8830-926b-4c44-a6bd-688866db2757'),
        ('b9e099f2-cd46-4bbb-8b07-a35cd352bc68', '63badf23-d63d-4ce5-b946-4940544fc963'),
        ('08cdafa3-595e-4d59-ba6e-aa1d5e024aef', '1ad0a226-df3f-48fd-b078-c02f6d5f013c'),
        ('1ad0a226-df3f-48fd-b078-c02f6d5f013c', '2f54b42f-239d-4c5e-b725-e696b5843b7e'),
        ('1ad0a226-df3f-48fd-b078-c02f6d5f013c', 'f11b24b3-2fb9-4ff8-a538-ba03007a9ec9'),
        ('76a4f294-fa55-46c9-bb45-22e5be268750', 'e559daa5-fd66-4c07-b85f-3f9c64d9d099'),
        ('76a4f294-fa55-46c9-bb45-22e5be268750', 'b2a970ee-0c51-4af1-b17b-9cf8005773ef'),
        ('89189372-970f-4bf8-a5c2-54554268e07d', '86fc70fe-6d01-48b3-b3f1-e52274547901'),
        ('89189372-970f-4bf8-a5c2-54554268e07d', 'a03b1b5d-f9e5-4d22-bcc4-14db8d40e591'),
        ('c0f6470a-5763-4798-bac7-90524f00f358', '1ad0a226-df3f-48fd-b078-c02f6d5f013c'),
        ('2792f195-813b-4f63-937e-81605382c475', 'e58452c9-2859-4ec5-8db8-aca088e88ce0'),
        ('2792f195-813b-4f63-937e-81605382c475', 'b857373a-0656-49a1-a7ea-566c1a549d04'),
        ('2792f195-813b-4f63-937e-81605382c475', '24e3b293-30f6-479d-9bef-dddd5e4a885d'),
        ('1ba621ab-b658-4572-baa7-5289d6254cc7', '1ad0a226-df3f-48fd-b078-c02f6d5f013c');

/* category product relation */
insert into public.category_product_relation (menu_category_id, menu_product_id)
values  ('a75a1502-f970-4400-8e1d-ce0dedaf7ea3', '3eccf3e0-26ba-4e05-ade5-a48a4030321b'),
        ('a75a1502-f970-4400-8e1d-ce0dedaf7ea3', '89189372-970f-4bf8-a5c2-54554268e07d'),
        ('5e20b60a-547c-4b39-997c-ca97791d1968', '76a4f294-fa55-46c9-bb45-22e5be268750'),
        ('5e20b60a-547c-4b39-997c-ca97791d1968', '42cdad60-e247-4277-8227-47868253e8c7'),
        ('003933bb-3e7a-4b4c-8162-90f592486941', '1ba621ab-b658-4572-baa7-5289d6254cc7'),
        ('003933bb-3e7a-4b4c-8162-90f592486941', 'df1c29b6-8a0a-4b59-9c93-dda28cf9da8a'),
        ('003933bb-3e7a-4b4c-8162-90f592486941', 'c0f6470a-5763-4798-bac7-90524f00f358'),
        ('003933bb-3e7a-4b4c-8162-90f592486941', '08cdafa3-595e-4d59-ba6e-aa1d5e024aef'),
        ('9c4efe6b-01b9-4346-8dff-740561b424eb', 'baabeddb-9be1-4195-a815-26d53cb45033'),
        ('9c4efe6b-01b9-4346-8dff-740561b424eb', '63badf23-d63d-4ce5-b946-4940544fc963'),
        ('926ee1b7-92be-49d6-bec2-feae63436b73', 'bc8d999a-4bac-496d-9628-bb81767b7319'),
        ('926ee1b7-92be-49d6-bec2-feae63436b73', '5d3c3ed6-36aa-40eb-830a-6f6dfdd4b49b'),
        ('926ee1b7-92be-49d6-bec2-feae63436b73', 'a2b82fa9-d0b5-43d8-bf33-68694ce2329f'),
        ('926ee1b7-92be-49d6-bec2-feae63436b73', '78ad748e-225a-4680-a6c0-37ed3da78d23'),
        ('926ee1b7-92be-49d6-bec2-feae63436b73', 'd7227149-91ad-4dc6-985d-660514ee477f');