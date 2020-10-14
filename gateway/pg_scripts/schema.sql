create table if not exists widgets (
	id integer primary key,
	color varchar(10) not null,
	size integer
);

create index if not exists idx_color on widgets (color);
create index if not exists idx_size on widgets (size);

create table if not exists prices (
	widget_id integer primary key,
	price integer
);

create index if not exists idx_price on prices (price);

create table if not exists inventory (
	widget_id integer primary key,
	inventory integer
);

create index if not exists idx_inventory on inventory (inventory);
