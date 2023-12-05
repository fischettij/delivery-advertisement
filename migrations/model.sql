CREATE TABLE IF NOT EXISTS establishments
(id           SERIAL PRIMARY KEY,
 establishment_id VARCHAR(255) NOT NULL,
 latitude DOUBLE PRECISION NOT NULL,
 longitude DOUBLE PRECISION NOT NULL,
 availability_radius DOUBLE PRECISION NOT null
);

CREATE INDEX IF NOT EXISTS idx_latitude ON establishments (latitude);

CREATE INDEX IF NOT EXISTS idx_longitude ON establishments (longitude);
