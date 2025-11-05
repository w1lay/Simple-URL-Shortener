CREATE TABLE links
(
    id            bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    short_link    varchar(8) NOT NULL,
    long_link     varchar(100) NOT NULL
);

CREATE INDEX ix_links_short_link ON links (short_link);