CREATE TABLE IF NOT EXISTS reactor(
    id UUID PRIMARY KEY,
    reagent_id UUID REFERENCES reagent (id)
);
