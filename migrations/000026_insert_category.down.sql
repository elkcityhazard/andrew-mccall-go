DELETE FROM categories
WHERE EXISTS (SELECT * FROM categories);
