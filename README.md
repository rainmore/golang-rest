# GoLang REST example

## Database Config

```sql
CREATE database go_lang_example;

\c go_lang_example

CREATE SCHEMA example;

CREATE USER go_lang_example WITH ENCRYPTED PASSWORD '{change-it}';
GRANT ALL PRIVILEGES ON DATABASE go_lang_example TO go_lang_example;
GRANT ALL ON SCHEMA example TO go_lang_example;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA example TO go_lang_example;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA example TO go_lang_example;

SET search_path TO example;

CREATE TABLE IF NOT EXISTS users (
  id              SERIAL PRIMARY KEY,
  first_name      VARCHAR(100) NOT NULL,
  last_name       VARCHAR(100) NOT NULL,
  email           VARCHAR(255) NOT NULL,
  date_of_birth   DATE NULL,
  created_at      TIMESTAMP WITH TIME ZONE NOT NULL
);


CREATE INDEX users_name ON users (first_name, last_name);
CREATE UNIQUE INDEX users_email ON users (email);
CREATE INDEX users_created_at ON users (created_at);



insert into users (first_name, last_name, email, date_of_birth, created_at) values('felix', 'rong', 'rainmore24@gmail.com', '1978-09-24', NOW());

select * from users;
```

Drop Database and User

```sql
REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA example FROM go_lang_example;
REVOKE ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA example FROM go_lang_example;
-- REVOKE ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA example FROM go_lang_example;
REVOKE ALL PRIVILEGES ON SCHEMA example FROM go_lang_example ;

DROP DATABASE IF EXISTS go_lang_example;

DROP USER go_lang_example;

```