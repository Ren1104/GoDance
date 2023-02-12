建库建表：

视频、用户、点赞、关注表

视频表：

```mysql
CREATE TABLE `videos` (
`id` bigint NOT NULL AUTO_INCREMENT,
`author_id` bigint DEFAULT NULL COMMENT '作者id',
`play_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '视频url',
`cover_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '封面url',
`create_at` bigint DEFAULT NULL COMMENT '创建时间',
`favorite_count` bigint DEFAULT NULL COMMENT '点赞数',
`comment_count` bigint DEFAULT NULL COMMENT '评论数',
`title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '文案',
 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
```

用户表:


```mysql
CREATE TABLE `users` (
 `id` bigint NOT NULL AUTO_INCREMENT,
 `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
 `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '密码',
 `follow_count` bigint DEFAULT NULL COMMENT '关注人数',
 `follower_count` bigint DEFAULT NULL COMMENT '粉丝数',
 `total_favorited` bigint DEFAULT NULL COMMENT '总喜欢数',
 `favorite_count` bigint DEFAULT NULL COMMENT '点赞数',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
```

点赞表:
```mysql
CREATE TABLE `favorite` (
 `id` bigint NOT NULL AUTO_INCREMENT,
 `user_id` bigint DEFAULT NULL COMMENT '用户id',
 `video_id` bigint DEFAULT NULL COMMENT '喜欢的视频id',
 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
```

关注表:
```mysql

CREATE TABLE `following` (
 `id` int NOT NULL AUTO_INCREMENT,
 `author_id` bigint DEFAULT NULL COMMENT '视频发布者id',
 `follower_id` bigint DEFAULT NULL COMMENT '关注者id',
 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
```