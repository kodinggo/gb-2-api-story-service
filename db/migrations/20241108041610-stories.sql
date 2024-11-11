
-- +migrate Up
CREATE TABLE `stories` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `content` text NOT NULL,
    `thumbnail_url` varchar(255) NOT NULL,
    `category_id` int(11) NOT NULL,
    `user_id` int(11) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT NOW(),
    `updated_at` timestamp NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`category_id`) REFERENCES categories (`id`)
);
-- +migrate Down
DROP TABLE IF EXISTS `stories`;
