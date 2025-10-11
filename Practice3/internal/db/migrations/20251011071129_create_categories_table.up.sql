create table categories (
    id SERIAL PRIMARY KEY ,
    name text not null ,
    user_id integer ,
    constraint fk_category_user foreign key (user_id) references users(id) on delete cascade ,
    constraint unique_user_category unique (user_id, name)
);

create index idx_categories_user_id on categories(user_id);


