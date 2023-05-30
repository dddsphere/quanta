CREATE TABLE events (
  stream_id text NOT NULL,
  stream_version text NOT NULL,
  event_id text NOT NULL,
  event_name text NOT NULL,
  event_data bytea NOT NULL,
  occurred_at timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (stream_id, stream_version, event_id),
);
