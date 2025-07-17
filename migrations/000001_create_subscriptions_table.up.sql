CREATE TABLE subscriptions (
    id              SERIAL PRIMARY KEY,
    service_name    TEXT NOT NULL,
    price           INTEGER NOT NULL CHECK (price >= 0),
    user_id         UUID NOT NULL,
    start_date      DATE NOT NULL
);

CREATE INDEX idx_subs_full ON subscriptions(start_date, user_id, service_name);
CREATE INDEX idx_subs_service ON subscriptions(start_date, service_name);
CREATE INDEX idx_service_name on subscriptions(service_name);
CREATE INDEX idx_user_id ON subscriptions(user_id);
