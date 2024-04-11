CREATE TABLE blogs (
   id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
   userId UUID NOT NULL,
   title VARCHAR(255) NOT NULL,
   content TEXT NOT NULL,
   date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
   FOREIGN KEY (userId) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_user_id ON blogs (userId);
