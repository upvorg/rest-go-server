CREATE DATABASE IF NOT EXISTS `gorm` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `gorm`;

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cover` varchar(200) DEFAULT '',
  `title` varchar(60) NOT NULL,
  `content` text DEFAULT NULL,
  `uid` int DEFAULT NULL,
  `tags` varchar(40) DEFAULT '' COMMENT '标签:最多四个 原创搞笑运动励志热血战斗竞技校园青春爱情恋爱冒险后宫百合治愈萝莉魔法悬疑推理奇幻科幻游戏神魔恐怖血腥机战战争犯罪历史社会职场剧情伪娘耽美童年教育亲子真人歌舞肉番美少女轻小说吸血鬼女性向泡面番欢乐向',
  `status` TINYINT(1) DEFAULT 3 COMMENT '1=>删除 | 2=>下架 | 3=>待审核| 4=>正常',
  `type` varchar(8) DEFAULT 'post' COMMENT 'post | video',
  `is_pined` TINYINT(1) DEFAULT 1 COMMENT '1=> | 2=>置顶',
  `is_recommend` TINYINT(1) DEFAULT 1 COMMENT '1=> | 2=>推荐',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `video_metas`;
CREATE TABLE `video_metas` (
  `id` int NOT NULL AUTO_INCREMENT,
  `pid` int NOT NULL,
  `title_japanese` varchar(60) DEFAULT '',
  `title_romanji` varchar(60) DEFAULT '',
  `genre` varchar(10) NOT NULL COMMENT '番剧|动画电影|电影|电视剧',
  `region` varchar(10) DEFAULT '美国｜日本｜中国',
  `is_end` TINYINT(1) DEFAULT 1 COMMENT '1=>未完结 | 2=>完结',
  `episodes` int DEFAULT 0 COMMENT "共几集",
  `publish_date` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `updated_date` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '每周几更新',
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `post_rankings`;
CREATE TABLE `post_rankings` (
`id` int NOT NULL AUTO_INCREMENT,
`pid` int NOT NULL,
`hits` int DEFAULT 0,
`hits_at` timestamp DEFAULT CURRENT_TIMESTAMP COMMENT '日｜月｜总',
 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `cover` varchar(200) DEFAULT '',
  `episode` int DEFAULT 1,
  `title` varchar(60) DEFAULT '',
  `title_japanese` varchar(60) DEFAULT '',
  `title_romanji` varchar(60) DEFAULT '',
  `video_url` varchar(200) NOT NULL,
  `synopsis` varchar(200) DEFAULT '',
  `pid` int NOT NULL,
  `uid` int NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `name` varchar(20) NOT NULL UNIQUE,
  `synopsis` varchar(200) DEFAULT '',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(16) NOT NULL UNIQUE,
  `nickname` varchar(16) DEFAULT '',
  `avatar` varchar(100) DEFAULT 'https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640',
  `pwd` varchar(100) NOT NULL,
  `email` varchar(50) DEFAULT NULL UNIQUE,
  `bio` varchar(100) DEFAULT '这个人很酷，什么都没有留下',
  `level` TINYINT(1) NOT NULL DEFAULT 4 COMMENT '1=>超级管理员 | 2=>管理员 | 3=>创作者 | 4=>普通用户',
  `status` TINYINT(1) DEFAULT 1 COMMENT '1=>正常 | 2=>封禁',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL,
  `target_id` int DEFAULT NULL,
  `pid` int NOT NULL,
  `uid` int NOT NULL,
  `vid` int DEFAULT 1,
  `content` varchar(200) NOT NULL,
  `color` varchar(10) DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `collections`;
CREATE TABLE `collections` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `pid` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL,
  `pid` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;


DROP TABLE IF EXISTS `feedbacks`;
CREATE TABLE `feedbacks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) DEFAULT '',
  `name` varchar(15) DEFAULT '佚名',
  `display_name` varchar(15) DEFAULT '佚名',
  `email` varchar(50) DEFAULT '',
  `message` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
