CREATE TABLE IF NOT EXISTS reagent_instance (
    id UUID NOT NULL,
    namespace VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (namespace, name)
);
