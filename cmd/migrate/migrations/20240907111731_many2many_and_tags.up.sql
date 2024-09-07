ALTER TABLE events
DROP COLUMN participants;

CREATE TABLE participants (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    event_id INT REFERENCES events(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, event_id)
);

