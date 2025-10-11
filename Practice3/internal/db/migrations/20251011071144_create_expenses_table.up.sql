create table expenses (
    id serial primary key ,
    user_id integer not null ,
    category_id integer not null ,
    amount numeric(12, 2) not null check ( amount > 0 ),
    currency char(3) not null ,
    spent_at timestamp not null,
    created_at timestamp default current_timestamp,
    note text,
    constraint fk_category_user foreign key (user_id) references users(id) on delete cascade ,
    constraint fk_categories foreign key (category_id) references categories(id) on delete cascade
);

create index idx_expenses_user_id on expenses(user_id);
create index idx_expenses_user_spent on expenses(user_id, spent_at);
