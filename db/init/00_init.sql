CREATE TABLE IF NOT EXISTS items (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO items(name) VALUES
  ('mise is nice'),
  ('go + postgres'),
  ('php + node')
ON CONFLICT DO NOTHING;
