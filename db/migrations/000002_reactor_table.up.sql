CREATE TABLE IF NOT EXISTS reactor (
    id UUID PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    reagent_id UUID REFERENCES reagent (id),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL
);
