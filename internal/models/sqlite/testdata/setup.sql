INSERT INTO users (name, email, hashed_password, created) VALUES (
    'Example Name',
    'alice@example.com',
    '$2a$12$rA8DKAcsvPecP6CFwL.lJe1338p8435dUdfPc7c7LKYcAeznUutxC',
    '2023-01-01 10:00:00'
);

INSERT INTO bookmarks (title, url, content, read, user_id, created) VALUES (
    'An Article Title',
    'https://example.com',
    '<p>some text here</p>',
    '1',
    '1',
    '2023-01-01 10:00:00'
);

INSERT INTO bookmarks (title, url, content, read, user_id, created) VALUES (
    'An Unread Article Title',
    'https://example.com/unreadarticle',
    '<p>This article is unread</p>',
    '0',
    '1',
    '2023-01-01 10:00:00'
);
