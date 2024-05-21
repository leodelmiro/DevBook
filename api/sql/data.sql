INSERT INTO users (name, nick, email, password)
VALUES
("User 1", "user_1", "user1@gmail.com", "$2a$10$Z7QxRdAjDdkHAIW04.4rOe6hHEoYCEVh4TDqWZwNCMkzau2ttQRNa"), -- user 1
("User 2", "user_2", "user2@gmail.com", "$2a$10$Z7QxRdAjDdkHAIW04.4rOe6hHEoYCEVh4TDqWZwNCMkzau2ttQRNa"), -- user 2
("User 3", "user_3", "user3@gmail.com", "$2a$10$Z7QxRdAjDdkHAIW04.4rOe6hHEoYCEVh4TDqWZwNCMkzau2ttQRNa"); -- user 3

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(1, 3),
(3, 1);