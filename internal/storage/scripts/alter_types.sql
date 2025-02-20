CREATE TYPE new_access_level AS ENUM('owner', 'full', 'read');

--
ALTER TABLE user_permission
    ALTER COLUMN level
        TYPE new_access_level
        USING user_permission.level::text::new_access_level;


DROP TYPE access_level;

ALTER TYPE new_access_level RENAME TO access_level;