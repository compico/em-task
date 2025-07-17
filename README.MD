


### Создание тестовых подписок

```sql
CREATE OR REPLACE FUNCTION generate_fake_subscriptions()
    RETURNS void AS $$
DECLARE
    services TEXT[] := ARRAY[
        'Yandex Plus',
        'Kinopoisk',
        'Netflix',
        'Yandex Music',
        'Spotify',
        'IVI'
        ];
    num_users INT := 800;
    total_rows INT := 1000;
    user_ids UUID[];
    i INT;
    uid UUID;
    service TEXT;
    price INT;
    month INT;
    day DATE;
BEGIN
    -- Генерируем 800 уникальных UUID и сохраняем их в массив
    FOR i IN 1..num_users LOOP
            user_ids := array_append(user_ids, gen_random_uuid());
        END LOOP;

    -- Создаём 1000 случайных подписок
    FOR i IN 1..total_rows LOOP
            uid := user_ids[1 + floor(random() * array_length(user_ids, 1))::int];
            service := services[1 + floor(random() * array_length(services, 1))::int];

            -- Цена от 199 до 899
            price := 100 * (2 + floor(random() * 8))::int - 1;

            -- Случайный месяц в 2025
            month := 1 + floor(random() * 12)::int;
            day := make_date(2025, month, 1);

            INSERT INTO subscriptions (service_name, price, user_id, start_date)
            VALUES (service, price, uid, day);
        END LOOP;
END;
$$ LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS pgcrypto;
SELECT generate_fake_subscriptions();
```