create or replace function create_code(o varchar, p varchar) RETURNS varchar(10) language plpgsql AS $$
    declare try_value varchar(10);
    declare inserted varchar(10);
    declare inserted_count int;
begin
    loop
        select random_between(1000000000, 9999999999)::text into try_value;
        insert into codes(owner, payload, value, valid_between)
        values (o, p, try_value, tstzrange(now(), now() + (60 ||' seconds')::interval, '[)'))
            on conflict on constraint unique_code_between_time_window do nothing returning value into inserted;
        get diagnostics inserted_count = row_count;
        exit when inserted_count > 0;
    end loop ;
    return inserted;
end; $$
