CREATE OR REPLACE FUNCTION create_code(o varchar, p varchar) RETURNS varchar(8) language plpgsql AS $$
declare try_value varchar(8);
declare inserted varchar(8);
declare inserted_count int;
begin
LOOP
select random_between(10000000, 99999999)::text into try_value;
INSERT INTO codes(owner, payload, value, valid_between)
VALUES (o, p, try_value, tstzrange(now(), now() + (120 ||' seconds')::interval, '[)'))
    ON CONFLICT ON CONSTRAINT unique_code_between_time_window DO NOTHING returning value into inserted;
GET DIAGNOSTICS inserted_count = ROW_COUNT;
EXIT WHEN inserted_count > 0;
END LOOP;
return inserted;
end; $$
