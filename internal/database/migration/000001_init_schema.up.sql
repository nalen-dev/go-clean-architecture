CREATE TABLE users(
    id              BIGSERIAL PRIMARY KEY,
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50),
    email           VARCHAR(100) NOT NULL,
    password        VARCHAR(100) NOT NULL,
    is_verified     BOOLEAN DEFAULT false,
    profile_picture VARCHAR(100),
    created_at      TIMESTAMP DEFAULT now(),
    updated_at      TIMESTAMP DEFAULT now()
);

CREATE TABLE jobs(
    id              BIGSERIAL PRIMARY KEY,
    user_id         BIGSERIAL,
    title           VARCHAR(100) NOT NULL,
    location        VARCHAR(100) NOT NULL,
    description     TEXT,
    created_at      TIMESTAMP DEFAULT now(),
    updated_at      TIMESTAMP DEFAULT now()
);

ALTER TABLE "jobs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");