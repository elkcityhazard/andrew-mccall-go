-- First update created_at with random dates
UPDATE posts 
SET created_at = DATE_ADD('2024-01-01', INTERVAL FLOOR(RAND() * 365) DAY);

-- Then update updated_at to be after created_at
UPDATE posts 
SET updated_at = DATE_ADD(created_at, INTERVAL FLOOR(RAND() * 30) DAY);

