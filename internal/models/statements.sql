CREATE TABLE IF NOT EXISTS users (
    userID INTEGER PRIMARY KEY,
    firstname TEXT,
    lastname TEXT,
    email TEXT UNIQUE,
    hashedPassword TEXT,
    token TEXT,
    expiry DATETIME
);

CREATE TABLE IF NOT EXISTS posts (
    postID INTEGER PRIMARY KEY,
    userID INTEGER REFERENCES users(userID),
    author TEXT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    likes_count NUMBER,
    dislikes_count NUMBER,
    comments_count NUMBER,
    tags TEXT NOT NULL

);

CREATE TRIGGER IF NOT EXISTS set_default_created_datetime
AFTER INSERT ON posts
BEGIN
    UPDATE posts SET created = DATETIME('now', '+6 hours') WHERE rowid = NEW.rowid;
END;



CREATE TABLE IF NOT EXISTS likes (postid INTEGER, likedby INTEGER);
CREATE TABLE IF NOT EXISTS dislikes (postid INTEGER, dislikedby INTEGER);



CREATE TABLE IF NOT EXISTS comments (
    postID INTEGER REFERENCES posts(postID),
    commentID INTEGER PRIMARY KEY,
    authorID INTEGER REFERENCES users(userID),
    comment TEXT,
    likes_count NUMBER,
    dislikes_count NUMBER
);

CREATE TABLE IF NOT EXISTS comment_likes (
    postid INTEGER,
    commentid INTEGER,
    likedby INTEGER
);

CREATE TABLE IF NOT EXISTS comment_dislikes (
    postid INTEGER,
    commentid INTEGER,
    dislikedby INTEGER
);

CREATE TABLE IF NOT EXISTS  categories (
	postid INTEGER,
	category TEXT
);

-- fake posts
-- INSERT INTO
-- 	posts (userID, author, title, content, created, likes_count,dislikes_count, comments_count, tags)
-- VALUES
-- 	(
-- 		'123123',
-- 		'Matsuo Bashō',
-- 		'An old silent pond',
-- 		'An old silent pond...
--     A frog jumps into the pond,
--     splash! Silence again.',
-- 		CURRENT_TIMESTAMP,
--         '85',
--         '3',
--         '0', 
--         "#Other"
-- 	);
-- INSERT INTO
-- 	posts (userID, author, title, content, created, likes_count,dislikes_count, comments_count, tags)
-- VALUES
-- 	(
-- 		'123124',
-- 		'Natsume Soseki',
-- 		'Over the wintry forest',
-- 		'Over the wintry
--     forest, winds howl in rage
--     with no leaves to blow.',
-- 		CURRENT_TIMESTAMP,
--         '9562',
--         '785456',
--         '0',
--         "#Other"
-- 	);
-- INSERT INTO
-- 	posts (userID, author, title, content, created, likes_count, dislikes_count, comments_count, tags)
-- VALUES
-- 	(
-- 		'123125',
-- 		'Murakami Kijo',
-- 		'First autumn morning',
-- 		'First autumn morning
--     the mirror I stare into
--     shows my father''s face.',
-- 		CURRENT_TIMESTAMP,
--         '123',
--         '5',
--         '0', 
--         "#Other"
-- 	);

