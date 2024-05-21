INSERT INTO users (name, nick, email, password)
VALUES
("User 1", "user_1", "user1@gmail.com", "$2a$10$DLBM3HA7wCmhxZmj94f.ce2j9UNOXZCC6cBy2DGT/cqT2xuB3HT7O"), -- user 1
("User 2", "user_2", "user2@gmail.com", "$2a$10$DLBM3HA7wCmhxZmj94f.ce2j9UNOXZCC6cBy2DGT/cqT2xuB3HT7O"), -- user 2
("User 3", "user_3", "user3@gmail.com", "$2a$10$DLBM3HA7wCmhxZmj94f.ce2j9UNOXZCC6cBy2DGT/cqT2xuB3HT7O"); -- user 3

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(3, 1);