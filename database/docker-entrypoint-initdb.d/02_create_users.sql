DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(191),
  `email` VARCHAR(191),
  PRIMARY KEY (`id`)
);

INSERT INTO users (id, name, email) VALUES (1, "dungnv", "dungnv@example.com");
INSERT INTO users (id, name, email) VALUES (2, "dungnv1", "dungnv1@example.com");