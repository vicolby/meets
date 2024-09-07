DROP TABLE IF EXISTS participants;

ALTER TABLE events
ADD COLUMN participants INTEGER[];
