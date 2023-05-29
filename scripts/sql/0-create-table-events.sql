CREATE TABLE streams (
  stream_id text NOT NULL,
  stream_name text NOT NULL,
  stream_version int NOT NULL,
  PRIMARY KEY (stream_id, stream_name, stream_version)
);

CREATE TABLE events (
  stream_id text NOT NULL,
  event_id text NOT NULL,
  event_name text NOT NULL,
  event_data bytea NOT NULL,
  occurred_at timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (stream_id, event_id),
  FOREIGN KEY (stream_id) REFERENCES streams (stream_id)
);
