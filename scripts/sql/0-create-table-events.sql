CREATE TABLE events (
  stream_id text NOT NULL,
  version text NOT NULL,
  event_id text NOT NULL,
  event_name text NOT NULL,
  event_data bytea NOT NULL,
  occurred_at timestamptz NOT NULL DEFAULT NOW(),
  PRIMARY KEY (stream_id, version, event_id),
);
