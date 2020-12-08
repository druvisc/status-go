ALTER TABLE raw_messages ADD COLUMN expired BOOL default false;
UPDATE raw_messages SET expired = 0;