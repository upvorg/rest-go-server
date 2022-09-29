CREATE DATABASE IF NOT EXISTS `YSZM` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `YSZM`;
-- MySQL dump 10.13  Distrib 5.7.38, for Linux (x86_64)
--
-- Host: localhost    Database: yszm
-- ------------------------------------------------------
-- Server version	5.7.38

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `collections`
--

DROP TABLE IF EXISTS `collections`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `collections` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `pid` varchar(36) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collections`
--

LOCK TABLES `collections` WRITE;
/*!40000 ALTER TABLE `collections` DISABLE KEYS */;
INSERT INTO `collections` VALUES (2,1,'4','2022-05-13 13:04:55');
/*!40000 ALTER TABLE `collections` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) DEFAULT NULL,
  `target_id` int(11) DEFAULT NULL,
  `pid` varchar(36) NOT NULL,
  `uid` int(11) NOT NULL,
  `vid` int(11) DEFAULT '1',
  `content` varchar(200) NOT NULL,
  `color` varchar(10) DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
INSERT INTO `comments` VALUES (1,0,0,'3',2,1,'留个沙发','','2022-05-24 08:44:05',NULL);
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `feedbacks`
--

DROP TABLE IF EXISTS `feedbacks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `feedbacks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) DEFAULT '',
  `name` varchar(15) DEFAULT '佚名',
  `display_name` varchar(15) DEFAULT '佚名',
  `email` varchar(50) DEFAULT '',
  `message` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `feedbacks`
--

LOCK TABLES `feedbacks` WRITE;
/*!40000 ALTER TABLE `feedbacks` DISABLE KEYS */;
INSERT INTO `feedbacks` VALUES (1,'127.0.0.1','root','不可以涩涩','','不可以涩涩 哒咩哒咩','2022-05-12 09:23:06'),(2,'127.0.0.1','佚名','月色真美群捏！','2990592294@qq.com','呜呜呜呜呜的星河呜呜呜','2022-06-25 11:59:10');
/*!40000 ALTER TABLE `feedbacks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `likes`
--

DROP TABLE IF EXISTS `likes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `likes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `pid` varchar(36) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `likes`
--

LOCK TABLES `likes` WRITE;
/*!40000 ALTER TABLE `likes` DISABLE KEYS */;
INSERT INTO `likes` VALUES (2,1,'1','2022-05-13 07:50:03'),(3,1,'2','2022-05-13 07:50:14'),(4,1,'3','2022-05-13 07:50:21'),(8,1,'5','2022-05-14 09:20:25'),(9,1,'163f51f8-76a4-4e41-8eeb-f27295520829','2022-05-24 11:47:11'),(10,1,'4','2022-05-25 05:49:05'),(11,1,'131478f2-31cd-446a-9cb2-15158b8c3efe','2022-05-30 16:28:28');
/*!40000 ALTER TABLE `likes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post_rankings`
--

DROP TABLE IF EXISTS `post_rankings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `post_rankings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` varchar(36) NOT NULL,
  `hits` int(11) DEFAULT '0',
  `hits_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '日｜月｜总',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=228 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post_rankings`
--

LOCK TABLES `post_rankings` WRITE;
/*!40000 ALTER TABLE `post_rankings` DISABLE KEYS */;
INSERT INTO `post_rankings` VALUES (1,'1',7,'2022-05-12 07:53:36'),(2,'2',3,'2022-05-12 08:02:43'),(3,'3',5,'2022-05-12 08:11:14'),(4,'3',6,'2022-05-13 07:48:40'),(5,'1',2,'2022-05-13 07:49:51'),(6,'2',5,'2022-05-13 07:50:12'),(7,'4',19,'2022-05-13 10:01:45'),(8,'5',4,'2022-05-14 09:19:44'),(9,'3',2,'2022-05-14 09:42:36'),(10,'4',2,'2022-05-14 09:44:19'),(11,'1',1,'2022-05-14 13:27:10'),(12,'6',1,'2022-05-14 18:12:03'),(13,'4',3,'2022-05-15 02:41:02'),(14,'1',1,'2022-05-15 23:23:53'),(15,'6',1,'2022-05-15 23:24:14'),(16,'3',1,'2022-05-16 06:00:19'),(17,'5',1,'2022-05-16 10:55:17'),(18,'4',2,'2022-05-16 10:55:38'),(19,'6',2,'2022-05-16 11:14:18'),(20,'1',1,'2022-05-16 11:53:30'),(21,'4',6,'2022-05-17 01:47:59'),(22,'4',3,'2022-05-18 02:55:12'),(23,'5',1,'2022-05-18 03:02:26'),(24,'3',1,'2022-05-18 03:16:07'),(25,'8',3,'2022-05-18 07:28:04'),(26,'3',2,'2022-05-19 04:49:39'),(27,'8',8,'2022-05-19 05:03:43'),(28,'4',1,'2022-05-19 07:54:26'),(29,'5',1,'2022-05-19 12:44:13'),(30,'1',2,'2022-05-19 12:44:36'),(31,'9',2,'2022-05-20 07:19:32'),(32,'4',1,'2022-05-20 07:50:14'),(33,'8',1,'2022-05-20 13:11:07'),(34,'9',1,'2022-05-21 05:41:15'),(35,'10',31,'2022-05-21 05:54:24'),(36,'2',1,'2022-05-21 17:26:45'),(37,'10',7,'2022-05-22 07:39:02'),(38,'6',1,'2022-05-22 15:39:10'),(39,'5',1,'2022-05-22 22:50:21'),(40,'5',3,'2022-05-24 00:16:49'),(41,'3',5,'2022-05-24 04:01:22'),(42,'4',1,'2022-05-24 05:51:40'),(43,'10',6,'2022-05-24 06:04:10'),(44,'8',12,'2022-05-24 08:30:31'),(45,'9',2,'2022-05-24 09:04:57'),(46,'163f51f8-76a4-4e41-8eeb-f27295520829',25,'2022-05-24 10:54:16'),(47,'1',1,'2022-05-24 20:04:20'),(48,'5',3,'2022-05-25 03:05:44'),(49,'1',1,'2022-05-25 05:13:48'),(50,'4',2,'2022-05-25 05:48:05'),(51,'163f51f8-76a4-4e41-8eeb-f27295520829',10,'2022-05-25 06:12:05'),(52,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-05-25 07:51:10'),(53,'131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-05-25 07:51:20'),(54,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-05-26 04:48:56'),(55,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-05-26 04:49:02'),(56,'1',1,'2022-05-26 06:40:57'),(57,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-05-27 10:55:19'),(58,'10',3,'2022-05-27 23:37:29'),(59,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-05-28 02:22:09'),(60,'8',4,'2022-05-28 02:22:46'),(61,'1',1,'2022-05-28 02:23:04'),(62,'10',1,'2022-05-28 09:08:52'),(63,'4',1,'2022-05-28 10:31:05'),(64,'5',1,'2022-05-28 23:54:52'),(65,'4',2,'2022-05-29 15:20:36'),(66,'2',4,'2022-05-29 15:50:14'),(67,'10',3,'2022-05-29 16:01:43'),(68,'8',1,'2022-05-30 13:33:54'),(69,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-05-30 16:13:09'),(70,'131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-05-30 16:13:26'),(71,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',1,'2022-05-30 16:40:24'),(72,'9',1,'2022-05-31 12:28:37'),(73,'1',2,'2022-06-02 00:50:06'),(74,'131478f2-31cd-446a-9cb2-15158b8c3efe',3,'2022-06-03 02:46:49'),(75,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',2,'2022-06-03 04:56:03'),(76,'5',1,'2022-06-03 05:08:39'),(77,'1',1,'2022-06-03 05:09:06'),(78,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',6,'2022-06-03 05:14:15'),(79,'8',1,'2022-06-03 11:59:13'),(80,'5',3,'2022-06-04 01:08:00'),(81,'6',1,'2022-06-04 10:55:40'),(82,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',3,'2022-06-05 05:25:58'),(83,'8',1,'2022-06-05 07:40:28'),(84,'10',2,'2022-06-05 07:40:37'),(85,'4',1,'2022-06-05 08:56:47'),(86,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-05 14:11:22'),(87,'1',1,'2022-06-06 07:41:04'),(88,'131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-06-06 22:46:22'),(89,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-08 11:39:53'),(90,'4',1,'2022-06-09 09:40:22'),(91,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-06-09 14:28:43'),(92,'8',1,'2022-06-09 14:36:52'),(93,'10',1,'2022-06-10 12:51:16'),(94,'1',1,'2022-06-10 16:58:04'),(95,'2',1,'2022-06-10 23:14:24'),(96,'8',3,'2022-06-11 03:42:22'),(97,'131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-06-11 04:45:19'),(98,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',2,'2022-06-11 06:37:21'),(99,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',1,'2022-06-11 07:02:09'),(100,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-11 16:03:39'),(101,'9',1,'2022-06-11 17:08:46'),(102,'10',1,'2022-06-11 17:10:03'),(103,'7',1,'2022-06-11 17:10:09'),(104,'163f51f8-76a4-4e41-8eeb-f27295520829',2,'2022-06-11 17:45:02'),(105,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-06-11 23:41:55'),(106,'6',3,'2022-06-12 02:14:26'),(107,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-06-12 09:32:09'),(108,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-12 20:20:25'),(109,'2',1,'2022-06-13 15:16:25'),(110,'2',1,'2022-06-14 10:59:29'),(111,'2',1,'2022-06-16 07:08:12'),(112,'9',1,'2022-06-16 07:25:18'),(113,'163f51f8-76a4-4e41-8eeb-f27295520829',2,'2022-06-16 12:05:53'),(114,'1',1,'2022-06-16 12:06:17'),(115,'5',1,'2022-06-17 01:51:49'),(116,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',3,'2022-06-17 01:52:03'),(117,'8',1,'2022-06-17 01:52:13'),(118,'4',1,'2022-06-17 03:57:24'),(119,'2',1,'2022-06-17 05:34:17'),(120,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-06-18 14:16:29'),(121,'1',2,'2022-06-19 04:42:36'),(122,'8',1,'2022-06-19 05:24:46'),(123,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',2,'2022-06-19 06:57:42'),(124,'10',1,'2022-06-19 11:00:59'),(125,'4',1,'2022-06-19 11:01:29'),(126,'2',1,'2022-06-19 22:43:41'),(127,'4',2,'2022-06-20 03:55:47'),(128,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-20 07:47:38'),(129,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-06-20 08:32:40'),(130,'9',1,'2022-06-20 15:40:59'),(131,'9',3,'2022-06-21 05:25:03'),(132,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',4,'2022-06-21 05:43:32'),(133,'10',7,'2022-06-21 05:44:22'),(134,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-21 18:03:50'),(135,'4',4,'2022-06-22 02:15:43'),(136,'10',2,'2022-06-22 04:51:17'),(137,'1',1,'2022-06-22 05:44:55'),(138,'4',2,'2022-06-23 00:32:25'),(139,'9',1,'2022-06-23 03:08:32'),(140,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-24 16:02:05'),(141,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-24 16:02:27'),(142,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-25 05:01:12'),(143,'5',1,'2022-06-25 05:01:18'),(144,'4',2,'2022-06-25 11:27:33'),(145,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',30,'2022-06-25 15:48:03'),(146,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',9,'2022-06-26 00:04:02'),(147,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-26 17:17:20'),(148,'4',1,'2022-06-27 15:21:15'),(149,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',7,'2022-06-27 15:21:22'),(150,'9',1,'2022-06-27 15:23:04'),(151,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-28 16:00:04'),(152,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',2,'2022-06-28 16:00:46'),(153,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-06-28 16:00:52'),(154,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-06-28 16:01:06'),(155,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',2,'2022-06-29 07:40:20'),(156,'2',1,'2022-07-04 01:41:22'),(157,'9',2,'2022-07-04 07:00:00'),(158,'4',1,'2022-07-05 03:37:09'),(159,'6',1,'2022-07-06 13:56:53'),(160,'10',1,'2022-07-06 22:16:37'),(161,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-07-07 20:41:54'),(162,'2',1,'2022-07-07 21:27:26'),(163,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',1,'2022-07-07 21:29:32'),(164,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-07-07 21:46:34'),(165,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-07-07 22:44:31'),(166,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-07-07 22:59:01'),(167,'4',4,'2022-07-08 02:26:05'),(168,'1',1,'2022-07-09 03:52:17'),(169,'4',1,'2022-07-10 01:52:43'),(170,'4',1,'2022-07-11 16:10:48'),(171,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',2,'2022-07-13 02:51:02'),(172,'4',1,'2022-07-13 06:07:15'),(173,'131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-07-14 09:48:06'),(174,'131478f2-31cd-446a-9cb2-15158b8c3efe',13,'2022-07-16 13:17:26'),(175,'131478f2-31cd-446a-9cb2-15158b8c3efe',2,'2022-07-17 11:44:08'),(176,'131478f2-31cd-446a-9cb2-15158b8c3efe',2,'2022-07-18 04:42:10'),(177,'2',1,'2022-07-18 10:13:16'),(178,'9',1,'2022-07-19 03:51:09'),(179,'131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-07-20 11:38:47'),(180,'163f51f8-76a4-4e41-8eeb-f27295520829',2,'2022-07-21 01:30:09'),(181,'5',1,'2022-07-21 01:30:43'),(182,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',1,'2022-07-21 01:31:01'),(183,'4',2,'2022-07-21 15:38:10'),(184,'10',2,'2022-07-21 15:42:06'),(185,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',1,'2022-07-24 07:08:22'),(186,'4',3,'2022-07-25 05:11:54'),(187,'1',1,'2022-07-26 05:40:31'),(188,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-07-26 14:38:33'),(189,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-07-29 13:07:11'),(190,'9',1,'2022-07-30 07:12:07'),(191,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',2,'2022-07-31 08:26:24'),(192,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-07-31 11:08:31'),(193,'2',1,'2022-08-01 00:24:40'),(194,'5',3,'2022-08-01 07:03:12'),(195,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-08-01 07:08:41'),(196,'9',1,'2022-08-01 11:16:13'),(197,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',4,'2022-08-02 02:21:44'),(198,'5',1,'2022-08-02 02:24:48'),(199,'1',1,'2022-08-02 02:40:27'),(200,'10',1,'2022-08-02 23:55:02'),(201,'10',1,'2022-08-03 17:51:32'),(202,'10',1,'2022-08-04 07:59:41'),(203,'4',1,'2022-08-08 02:30:52'),(204,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',1,'2022-08-11 18:16:51'),(205,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',1,'2022-08-11 22:55:25'),(206,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',1,'2022-08-12 01:53:45'),(207,'1',1,'2022-08-12 09:25:40'),(208,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-08-12 09:40:40'),(209,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-08-12 18:53:51'),(210,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',1,'2022-08-12 21:02:55'),(211,'2',1,'2022-08-12 22:10:17'),(212,'2',1,'2022-08-13 16:47:06'),(213,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',2,'2022-08-18 16:45:15'),(214,'2',2,'2022-08-18 17:30:17'),(215,'163f51f8-76a4-4e41-8eeb-f27295520829',1,'2022-08-20 12:59:22'),(216,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',2,'2022-08-24 07:16:44'),(217,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',2,'2022-08-25 17:50:29'),(218,'1',1,'2022-08-25 19:54:15'),(219,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',1,'2022-08-26 08:51:35'),(220,'77745a95-b5db-4c0b-ad08-0a55a6e70f49',1,'2022-08-27 04:01:27'),(221,'9',2,'2022-08-27 21:19:14'),(222,'3f118cbc-c6a4-4368-b472-cf6677ecd02b',2,'2022-08-27 21:50:00'),(223,'1',1,'2022-08-27 22:48:20'),(224,'2',1,'2022-08-29 09:30:08'),(225,'9',1,'2022-08-29 19:57:58'),(226,'5589f28e-42c1-4c4a-bc58-f41e7f1d22f3',1,'2022-08-30 12:01:56'),(227,'a7a1345e-1bc8-4793-9650-a0ee89b6d74b',1,'2022-08-30 22:34:02');
/*!40000 ALTER TABLE `post_rankings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `posts` (
  `id` varchar(36) NOT NULL,
  `cover` varchar(200) DEFAULT '',
  `title` varchar(60) NOT NULL,
  `content` text,
  `uid` int(11) DEFAULT NULL,
  `tags` varchar(40) DEFAULT '' COMMENT '标签:最多四个 原创搞笑运动励志热血战斗竞技校园青春爱情恋爱冒险后宫百合治愈萝莉魔法悬疑推理奇幻科幻游戏神魔恐怖血腥机战战争犯罪历史社会职场剧情伪娘耽美童年教育亲子真人歌舞肉番美少女轻小说吸血鬼女性向泡面番欢乐向',
  `status` tinyint(1) DEFAULT '3' COMMENT '1=>删除 | 2=>下架 | 3=>待审核| 4=>正常',
  `type` varchar(8) DEFAULT 'post' COMMENT 'post | video',
  `is_pined` tinyint(1) DEFAULT '1' COMMENT '1=> | 2=>置顶',
  `is_recommend` tinyint(1) DEFAULT '1' COMMENT '1=> | 2=>推荐',
  `is_original` tinyint(1) DEFAULT '1' COMMENT '1=>否 | 2=>是',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;
/*!40000 ALTER TABLE `posts` DISABLE KEYS */;
INSERT INTO `posts` VALUES ('1','https://ae01.alicdn.com/kf/Ue6647d254f55449b9872794e31d5ae13B.jpg','异种族风俗娘评鉴指南','这里不仅是人类，还有艾尔夫、兽人、恶魔、天使，所有异种族混在一起生活的世界。那里当然也有各种异族的写生店…。\n在提供足不出户的服务的店里工作的人类冒险者史坦克，有一天因种族间(性意义上的)感性差异，与恶友艾洛埃尔夫·塞尔发生冲突。\n终结的方法是……女儿的评论!?以交叉评论的方式对所有异种族姑娘的服务进行评分，作为给其他同伴的“角色勃”信息提供的斯坦克们的活跃，简直就像性战士一样!\n今天，女士们也为了寻找新的快乐而开始旅行……。',1,'其他',4,'video',1,2,1,'2022-05-12 07:44:00','2022-05-12 08:40:13',NULL),('10','http://pic.ku-img.com/upload/vod/20220507-1/3f028116b79860ca1d1bbd7560129225.jpg','爱，死亡和机器人第三季','艾美奖获奖动画选集《爱、死亡 & 机器人》第三部回归，由蒂姆·米勒（《死侍》《终结者：黑暗命运》）和大卫·芬奇（《心灵猎人》《曼克》）担任监制。恐怖、想象力和美在新剧集中完美融合，从揭示古老的邪恶力量到喜剧般的末日，剧集以标志性的巧思和创造性的视觉效果，为观众带来令人震惊的奇幻、恐怖和科幻短篇故事。',1,'R15',4,'video',1,2,1,'2022-05-21 05:44:16','2022-05-21 05:54:18',NULL),('131478f2-31cd-446a-9cb2-15158b8c3efe','https://tu.11114.cc/zxdy/uploads/allimg/210916/2bdb9d09305fc071.jpg','你的名字','',1,'其他',4,'video',1,2,1,'2022-05-24 13:00:39','2022-05-24 13:02:47',NULL),('163f51f8-76a4-4e41-8eeb-f27295520829','https://img1.imgtp.com/2022/05/24/vbAqTytU.png','式守同学不只可爱而已','超级“帅气女友“登场！ 和泉是一名拥有不幸体质的高中男生，他有一个和他同班的女朋友，叫做式守。 式守的笑容十分甜美、温柔，跟和泉在一起的时候脸上总是洋溢着幸福。她平时可爱动人，心中满是爱情，但只要看到和泉遇到危险，她就会……摇身一变，变成“帅气女友”！满是可爱×帅气的式守跟和泉将与他们的小伙伴一起带来无限愉快的日常！1000%美好的爱情喜剧，即将开幕！',36,'其他 R15',4,'video',1,2,1,'2022-05-24 10:50:22','2022-05-24 13:01:42',NULL),('18d5c720-6f19-4095-9dae-ca8a072fecc2','https://img1.com','多彩时间的明天','',1,'其他',3,'video',1,1,1,'2022-05-26 06:29:53','2022-05-26 06:29:53',NULL),('2','https://s2.loli.net/2022/05/01/nzFqRv3HPryWupg.png','恋爱要在征服世界后','说明一下吧！以世界和平为目标的英雄战队“杰拉特5”的队长·相川不动，与企图征服世界的秘密结社“格格车”的战斗员队长·“死神公主”祸原死亡美。宿敌之间的他们，有着超越组织之墙的深深因缘！！其实这两个人…在交往！两人没有向社会和朋友公开，开始秘密交往。但是，对于恋爱初学者的纯洁的二人来说，一切都是新的。一旦被发现，比赛就结束的禁断的爱情喜剧开始了',1,'其他',4,'video',1,2,1,'2022-05-12 08:00:40','2022-05-12 08:40:19',NULL),('3','https://s2.loli.net/2021/12/17/z8XctaHY2qlmLVG.png','君之名','![suo](https://s2.loli.net/2021/12/17/z8XctaHY2qlmLVG.png)\\n\\n别名：你的名字/君之名/Your Name/Kimi no na wa.\\n\\n清晰：剧场版\\n\\n类型：[奇幻](1) [治愈](1) [穿越](1) [新海诚系列](1)\\n\\n主演：[神木隆之介](1) [上白石萌音](1) [长泽雅美](1) [市原悦子](1)\\n\\n导演：[新海诚](1)\\n\\n国家/地区：日本\\n\\n语言/字幕：国语对白 中文字幕\\n\\n更新时间：2021-11-07 08:29:06\\n',1,'其他',4,'video',1,1,1,'2022-05-12 08:11:08','2022-05-24 13:03:54','2022-05-24 13:03:55'),('3f118cbc-c6a4-4368-b472-cf6677ecd02b','https://img1.imgtp.com/2022/05/25/EBc456d9.jpg','借东西的小人阿莉埃蒂','',36,'治愈 恋爱',4,'video',1,1,1,'2022-05-25 06:29:45','2022-05-26 06:30:00',NULL),('4','https://tvax3.sinaimg.cn/large/007Pu4zFly1h1lyoe5m8pj30h30m8q4s.jpg','杜鹃的婚约',' 讲述因抱错婴儿，让就读名门私立学校的男子高中生「海野凪」与酒店总裁千金「天野绘里香」被父母订下婚约，但凪暗恋着「濑川弥」，而凪的妹妹也对他有着不一样的心情，因此展开一场四角关系的恋爱喜剧。',1,'其他',4,'video',1,2,1,'2022-05-13 10:00:05','2022-05-13 12:50:24',NULL),('5','','TikTok_v24.3.5_MOD 版','TikTok_v24.3.5_MOD版\n由 DMITRY Nechoporenko 修改\n此更新有什么新功能？\n- 一种绕过安全检查的新方法\n描述:\n- 删除了视频谷歌和其他广告\n- 将无水印的视频和 GIF 下载到 Movies/TikTok 文件夹而不是 DCIM/Camera\n- 删除了所有下载限制，您可以下载任何视频\n- 删除了许多其他限制\n- 应用程序已尽可能清理\n- 最大压缩 + ZipAlign\n- 禁用不必要的活动\n- 删除对二重奏、拼接和动态壁纸的限制\n- 现在任何视频都可以倒带\n- 电池消耗优化\n- 删除区域限制\n- 修复 Facebook 授权\n- 修复 VK 授权\n-固定谷歌授权\n- 禁用自动启动\n- 启用高品质音频\n- 启用高品质视频\n- 启用超分辨率\n- 启用抗锯齿\n- 隐藏根权限\n- 禁用 InAppBillingService\n- 禁用所有类型的分析\n- 禁用测量 ',1,'其他',4,'post',1,1,1,'2022-05-14 09:10:51','2022-05-22 03:51:03',NULL),('5589f28e-42c1-4c4a-bc58-f41e7f1d22f3','https://s1.ax1x.com/2022/06/03/XNxrHx.png','投稿教程','// TOODO\n---\n![15](https://s1.ax1x.com/2022/06/03/XNxrHx.png)',1,'',4,'post',1,1,1,'2022-06-03 05:11:46','2022-06-03 05:13:58',NULL),('6','https://cdn.04pic.com/image/6261987222ef2.jpg','夏日重现','',1,'其他',4,'video',1,2,1,'2022-05-14 18:11:56','2022-05-18 07:32:36',NULL),('7','https://cdn.tupianla.cc/images/5ikmj/uploads/allimg/210821/85f7787b37c25437.jpg','我想吃掉你的胰脏','没有名字的我，没有未来的她”对他人毫无兴趣，总是独自一人读书的高中生“我”。这样的“我”有一天，偶然捡到一册写着《共病文库》的文库本。那是，天真烂漫的班上人气王·山内樱良私下记录的日记本。里面记载着她身患胰脏的疾病，已经时日无多……。隐藏自己的疾病度过日常的樱良，与知晓其秘密的“我”。——两人的距离，还没有名字。',1,'其他',4,'video',1,2,1,'2022-05-18 07:22:49','2022-05-18 07:33:11',NULL),('77745a95-b5db-4c0b-ad08-0a55a6e70f49','https://91m.pilipata.com/upload/vod/20220524-1/552c0b86fab876803f789b73ea22fe11.jpg','言叶之庭','',1,'其他',4,'video',1,2,1,'2022-05-24 12:51:22','2022-05-24 13:01:49',NULL),('8','https://pic.url.cn/qqgameedu/0/9cff033db0e36fbbd22f3eeda8371306/0','间谍过家家','数々の漫画賞を受賞している人気漫画『SPY×FAMILY』（作者：遠藤達哉）が、2022年にテレビアニメ化されることが決定した。キャスト・スタッフ情報も公開され、主人公のロイド・フォージャー役を江口拓也、監督は古橋一浩氏（代表作：『機動戦士ガンダムUC』）が担当し、制作はWIT STUDIO×CloverWorksによる共同制作となる。',1,'其他',4,'video',1,2,1,'2022-05-18 07:25:31','2022-05-18 07:28:22',NULL),('9','https://cdn.tupianla.cc/images/5ikmj/uploads/allimg/210823/25469c1f97c06ad6.jpg','秒速5厘米','',1,'其他',4,'video',1,2,1,'2022-05-20 04:57:32','2022-05-20 07:19:26',NULL),('a7a1345e-1bc8-4793-9650-a0ee89b6d74b','','新 QQ 群','点击链接加入群聊【月色真美】：[https://jq.qq.com/?_wv=1027&k=EFpgexC8](https://jq.qq.com/?_wv=1027&k=EFpgexC8)',1,'公告',4,'post',1,1,1,'2022-06-25 15:47:53','2022-06-25 15:50:19',NULL),('b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7','https://pic.monidai.com/img/55208d35c4e9b.jpg','可塑性记忆','',1,'其他',4,'video',1,2,1,'2022-06-03 04:51:29','2022-06-03 04:55:56',NULL);
/*!40000 ALTER TABLE `posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  `synopsis` varchar(200) DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,1,'其他','','2022-05-12 07:41:35'),(2,1,'R15','','2022-05-12 11:48:48'),(3,36,'恋爱','','2022-05-24 14:27:32'),(4,36,'漫画改','','2022-05-24 14:27:53'),(5,36,'日常','','2022-05-24 14:27:58'),(6,36,'校园','','2022-05-24 14:28:45'),(7,1,'新海诚','','2022-05-25 05:30:50'),(8,1,'搞笑','','2022-05-25 05:41:37'),(9,1,'励志','','2022-05-25 05:41:40'),(10,1,'治愈','','2022-05-25 05:41:44'),(11,1,'公告','','2022-06-03 05:11:29');
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(16) NOT NULL,
  `nickname` varchar(16) DEFAULT '',
  `avatar` varchar(100) DEFAULT 'https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640',
  `pwd` varchar(100) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `bio` varchar(100) DEFAULT '这个人很酷，什么都没有留下',
  `level` tinyint(1) NOT NULL DEFAULT '4' COMMENT '1=>超级管理员 | 2=>管理员 | 3=>创作者 | 4=>普通用户',
  `status` tinyint(1) DEFAULT '1' COMMENT '1=>正常 | 2=>封禁',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `qq_openid` varchar(36) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'root','YUESE','https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640','$2a$04$SKhBaklLI3Ky1c5K8O0hbOqnKTobiq1Q1LLoQSXPjOl5RUVUfujM.',NULL,'这个人很酷，什么都没有留下',1,1,'2022-05-12 05:51:42','2022-05-18 09:36:34',NULL,NULL),(2,'mufunka','mufunka','https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640','$2a$04$N5GsQlcoAurKtXClfwbmkONObDNuSxQ6AextN9a5i29rdmWInapWO','','这个人很懒，什么都没有留下',4,2,'2022-05-19 08:19:36','2022-05-19 08:19:36',NULL,NULL),(36,'Keain','Keain','https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640','$2a$04$JRneMnxwOPpzKVK08izh4.ClG9sR1goiwnH6Ansfvat7FM/rkMBYi',NULL,'这个人很酷，什么都没有留下',3,2,'2022-05-24 10:15:28','2022-05-24 10:19:44',NULL,NULL),(37,'Tranaki','Tranaki','https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640','$2a$04$bEQsHOrYe1Qqr5.g4MJOquUZG/1Mp1XlyvbE49WViHSb2sE33TjEG',NULL,'这个人很酷，什么都没有留下',4,2,'2022-06-05 07:40:26','2022-06-05 07:40:26',NULL,NULL),(38,'136621','136621','https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640','$2a$04$iBwdtu5/q75DfhzcGVJX4eoljG8jlQI4eEVfO.J8od577Cyq40Fue',NULL,'这个人很酷，什么都没有留下',4,2,'2022-06-25 16:03:24','2022-06-25 16:03:24',NULL,NULL),(39,'muciku','muciku','https://q1.qlogo.cn/g?b=qq&nk=7619376472&s=640','$2a$04$1uhyQsEQ9kBVez1a2WYKcuMUFglyy/oc3ZQuT0SsQ5sCgxV45etcS',NULL,'这个人很酷，什么都没有留下',4,2,'2022-07-23 09:25:49','2022-07-23 09:25:49',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video_metas`
--

DROP TABLE IF EXISTS `video_metas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `video_metas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` varchar(36) NOT NULL,
  `title_japanese` varchar(60) DEFAULT '',
  `title_romanji` varchar(60) DEFAULT '',
  `genre` varchar(10) NOT NULL COMMENT '番剧|动画电影|电影|电视剧',
  `region` varchar(10) DEFAULT '美国｜日本｜中国',
  `is_end` tinyint(1) DEFAULT '1' COMMENT '1=>未完结 | 2=>完结',
  `episodes` int(11) DEFAULT '0' COMMENT '共几集',
  `publish_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `updated_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '每周几更新',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video_metas`
--

LOCK TABLES `video_metas` WRITE;
/*!40000 ALTER TABLE `video_metas` DISABLE KEYS */;
INSERT INTO `video_metas` VALUES (1,'1','','','番剧','日本',2,12,NULL,NULL,NULL),(2,'2','','','番剧','日本',1,5,'2022-04-20 00:00:00','2022-05-13 00:00:00',NULL),(3,'3','','','动画电影','日本',2,1,NULL,NULL,'2022-05-24 13:03:55'),(4,'4','','','番剧','日本',1,1,NULL,'2022-05-14 00:00:00',NULL),(5,'6','','','番剧','日本',1,1,NULL,NULL,NULL),(6,'7','','','动画电影','中国',2,1,NULL,NULL,NULL),(7,'8','','','番剧','日本',1,0,NULL,NULL,NULL),(8,'9','','','番剧','日本',1,0,NULL,NULL,NULL),(9,'10','','','电视剧','美国',1,0,NULL,NULL,NULL),(10,'163f51f8-76a4-4e41-8eeb-f27295520829','','','番剧','日本',1,0,'2022-04-19 00:00:00','2022-05-24 00:00:00',NULL),(11,'77745a95-b5db-4c0b-ad08-0a55a6e70f49','','','番剧','日本',1,1,NULL,NULL,NULL),(12,'131478f2-31cd-446a-9cb2-15158b8c3efe','','','动画电影','其他',1,1,NULL,NULL,NULL),(13,'3f118cbc-c6a4-4368-b472-cf6677ecd02b','','','动画电影','日本',2,1,NULL,NULL,NULL),(14,'18d5c720-6f19-4095-9dae-ca8a072fecc2','','','番剧','日本',1,1,NULL,NULL,NULL),(15,'b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7','','','番剧','日本',2,13,'2015-01-01 00:00:00',NULL,NULL);
/*!40000 ALTER TABLE `video_metas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `videos`
--

DROP TABLE IF EXISTS `videos`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `videos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cover` varchar(200) DEFAULT '',
  `episode` int(11) DEFAULT '1',
  `title` varchar(60) DEFAULT '',
  `title_japanese` varchar(60) DEFAULT '',
  `title_romanji` varchar(60) DEFAULT '',
  `video_url` varchar(1024) NOT NULL,
  `synopsis` varchar(200) DEFAULT '',
  `pid` varchar(36) NOT NULL,
  `uid` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `videos`
--

LOCK TABLES `videos` WRITE;
/*!40000 ALTER TABLE `videos` DISABLE KEYS */;
INSERT INTO `videos` VALUES (1,'',2,'','','','https://mysource-anfunsapi-bangumi.anfuns.cn/api/v3/file/source/3717/[UHA-WINGS][Ishuzoku%20Rebyuazu][01][x264][CHT][AT-X%201280x720].mp4?sign=oX8Y0a-QUkxFy4cj3WRJAN33c8SdhxjtNFebTsSEv_0=:0','','1',1,'2022-05-12 07:48:13','2022-05-12 07:51:09'),(2,'',1,'','','','https://mysource-anfunsapi-bangumi.anfuns.cn/api/v3/file/source/3711/[UHA-WINGS][Ishuzoku%20Rebyuazu][02][x264][CHT][AT-X%201280x720].mp4?sign=OU4umKZdXve0uctJ3OhRwbIrXCGozvIIgvqUNzVkz5s=:0','','1',1,'2022-05-12 07:48:23','2022-05-12 07:48:23'),(3,'',3,'','','','https://mysource-anfunsapi-bangumi.anfuns.cn/api/v3/file/source/3710/[UHA-WINGS][Ishuzoku%20Rebyuazu][03][x264][CHT][AT-X%201280x720].mp4?sign=Zl0YY4jSTC5KqP8G6BiO2g0H4mF5TpekbfgKIJ1Pf8o=:0','','1',1,'2022-05-12 07:51:35','2022-05-12 07:51:35'),(4,'',2,'','','','https://mysource-anfunsapi-bangumi.anfuns.cn/api/v3/file/source/3305/[NC-Raws] 愛在征服世界後 - 02 (Baha 1920x1080 AVC AAC MP4).mp4?sign=shtVVkwuVt0MjaRrgPS-XDnzkWOg9CpU9gKejwEFcL8=:0','','2',1,'2022-05-12 08:00:55','2022-05-12 08:04:28'),(5,'',1,'','','','https://mysource-anfunsapi-bangumi.anfuns.cn/api/v3/file/source/3187/[NC-Raws] 愛在征服世界後 - 01 (Baha 1920x1080 AVC AAC MP4).mp4?sign=hIF55CgPftirBN06BF5cFOZRN52MnIICgBsgJtAe5c0=:0','','2',1,'2022-05-12 08:02:04','2022-05-12 08:03:28'),(6,'',3,'','','','https://mysource-anfunsapi-bangumi.anfuns.cn/api/v3/file/source/3490/[NC-Raws] Koi wa Sekai Seifuku no Ato de - 03 (Baha 1920x1080 AVC AAC MP4).mp4?sign=sw1jafQ6osw1hQseskk2-qYl6IhXp0iK0gLYHxVRXTE=:0','','2',1,'2022-05-12 08:04:45','2022-05-12 08:04:45'),(7,'',4,'','','','https://s2.monidai.com/ppvod/D5213B2CC36C72170278F4C55D6816D0.m3u81','','2',1,'2022-05-12 08:04:57','2022-05-12 08:07:04'),(8,'',1,'','','','https://svp.cdn.qq.com/0b53o4akaaaaemakbx3o6nrja56dub3qbiaa.f0.mp4?dis_k=c48661fc7823c752057af41be8d752c7&dis_t=1651376213','','2',1,'2022-05-12 08:06:00','2022-05-12 08:06:00'),(9,'',1,'','','','https://m3u8.yyd.me/statictt/MmUeIixzRaTiFeFdMdUYxQjFGMUIxRDFC/ddd78eba64c3995567cb14f27d9c72de/643631626336386234.m3u8','','3',1,'2022-05-12 08:11:28','2022-05-12 08:11:28'),(10,'',1,'','','','https://s2.monidai.com/ppvod/E86CABD04478859BA65F6761D9AF58C9.m3u8','','4',1,'2022-05-13 10:00:39','2022-05-13 10:28:03'),(12,'',2,'','','','https://s2.monidai.com/ppvod/FCD7885D5BE89FDA071A063380B93C25.m3u8','','4',1,'2022-05-13 10:11:58','2022-05-17 09:59:43'),(13,'',3,'','','','https://s2.monidai.com/ppvod/700A4258AEF616893FF08919E86A13C5.m3u8','','4',1,'2022-05-13 10:28:28','2022-05-17 09:59:51'),(14,'',4,'','','','https://s2.monidai.com/20220515/jqNVG7pD/index.m3u8','','4',1,'2022-05-18 06:54:01','2022-05-18 06:54:01'),(15,'',1,'','','','https://v.kdcdn.net/20210928/eoB5uhLE/index.m3u8','','7',1,'2022-05-18 07:23:02','2022-05-18 07:23:02'),(16,'',1,'','','','https://c2.monidai.com/20220410/rWAsAYym/index.m3u8','','8',1,'2022-05-18 07:25:45','2022-05-18 07:25:45'),(17,'',2,'','','','https://c2.monidai.com/20220416/deoaaVzJ/index.m3u8','','8',1,'2022-05-18 07:25:59','2022-05-18 07:26:33'),(18,'',3,'','','','https://s2.monidai.com/20220423/hcx3Be1E/index.m3u8','','8',1,'2022-05-18 07:26:28','2022-05-18 07:26:28'),(19,'',4,'','','','https://s2.monidai.com/20220430/8RVpOfBn/index.m3u8','','8',1,'2022-05-18 07:26:45','2022-05-18 07:26:45'),(20,'',5,'','','','https://s2.monidai.com/20220507/pqyaT7mD/index.m3u8','','8',1,'2022-05-18 07:26:57','2022-05-18 07:27:26'),(21,'',6,'','','','https://s2.monidai.com/20220514/fLVGkkWB/index.m3u8','','8',1,'2022-05-18 07:27:21','2022-05-18 07:27:21'),(24,'',1,'','','','https://ukzyvod3.ukubf5.com/20220410/yAU8vUFg/2000kb/hls/index.m3u8','','9',1,'2022-05-20 07:22:42','2022-05-24 12:39:35'),(26,'',1,'','','','https://v.v1kd.com/20220507/AQzU93SJ/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:44:52','2022-05-21 05:49:14'),(27,'',2,'','','','https://v.v1kd.com/20220507/EZbmRzt5/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:45:27','2022-05-21 05:51:58'),(28,'',3,'','','','https://v.v1kd.com/20220507/Q7RVHMuw/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:52:14','2022-05-21 05:52:14'),(29,'',4,'','','','https://v.v1kd.com/20220507/sYAgfN3v/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:52:31','2022-05-21 05:52:31'),(30,'',5,'','','','https://v.v1kd.com/20220507/mEDIMkUB/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:52:49','2022-05-21 05:52:49'),(31,'',6,'','','','https://v.v1kd.com/20220507/ayG1jXcK/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:53:06','2022-05-21 05:53:06'),(32,'',7,'','','','https://v.v1kd.com/20220507/Y30hQlmb/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:53:20','2022-05-21 05:53:20'),(33,'',8,'','','','https://v.v1kd.com/20220507/0x7Hlc7r/2000kb/hls/index.m3u8','','10',1,'2022-05-21 05:53:44','2022-05-21 05:53:44'),(35,'',1,'','','','https://cdn.zoubuting.com/20210706/2N8rtI9R/1200kb/hls/index.m3u8','','77745a95-b5db-4c0b-ad08-0a55a6e70f49',1,'2022-05-24 12:51:34','2022-05-24 12:51:34'),(36,'',1,'','','','https://baidu.sd-play.com/20211015/hJx6vRoz/hls/index.m3u8','','131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-05-24 13:00:44','2022-05-24 13:00:44'),(37,'',2,'','','','https://dm.sszyplay.com/20220324/8uYGGlsD/2000kb/hls/index.m3u8','','131478f2-31cd-446a-9cb2-15158b8c3efe',1,'2022-05-24 13:06:49','2022-05-24 13:06:49'),(39,'https://img1.imgtp.com/2022/05/24/UlhL9KtI.jpg',1,'第1集 我的女朋友非常可爱','','','https://v9-default.ixigua.com/163ac3943d850ce496f7366abd826104/628dea30/video/tos/cn/tos-cn-v-c9e10a/d72784b247664af5befffcf24f66809b/','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 13:48:08','2022-05-25 07:41:06'),(40,'https://img1.imgtp.com/2022/05/24/afmZ9scb.jpg',2,'第2集 风起云涌、球技大赛！','','','https://v9-default.ixigua.com/b50a0f99387477a637963536aebb8b4d/628cf3fc/video/tos/cn/tos-cn-v-c9e10a/88939415c54a4c4e9e92307869469ef7/','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 13:50:23','2022-05-24 13:52:22'),(41,'https://img1.imgtp.com/2022/05/24/zdM0oztm.jpg',3,'第3集 不幸转晴','','','https://v6-default.ixigua.com/45fba06d2d87c7b595925e90c33747a1/628cfde0/video/tos/cn/tos-cn-v-c9e10a/1084e3fadf7c411d91b20c5d4b5fa805/','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 13:53:21','2022-05-24 14:36:57'),(42,'https://img1.imgtp.com/2022/05/24/gJmWdeP4.png',4,'第4集 立夏、各自的心意','','','https://v3-default.ixigua.com/a1812a8e14eee2b74cde9d00f963e846/628d01d3/video/tos/cn/tos-cn-v-6f4170/72cd544617df42c2ac39d5aec5c0b8dd/?libvio.cc&filename=1.mp4','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 14:40:58','2022-05-24 14:40:58'),(43,'',5,'第5集','','','https://v6-default.ixigua.com/5144192003a5223ba0f94270a5f5f94b/628d020f/video/tos/cn/tos-cn-v-6f4170/220a246152a7437889dba477b3db72bc/?libvio.cc&filename=1.mp4','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 14:41:54','2022-05-24 14:41:54'),(44,'',6,'第6集','','','https://v3-default.ixigua.com/c80991a19d798115bc887cb3ef71c22a/628cf989/video/tos/cn/tos-cn-v-6f4170/19292dd9f7b6482595f98d2dfa89a95d/','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 14:43:14','2022-05-24 14:43:14'),(45,'',7,'第6.5集','','','https://v3-default.ixigua.com/50676779d33bf54019bd8494e5f12aa2/628d0298/video/tos/cn/tos-cn-v-6f4170/901edc995bc64bce8e9d44b58e27342f/?libvio.cc&filename=1.mp4','','163f51f8-76a4-4e41-8eeb-f27295520829',36,'2022-05-24 14:46:03','2022-05-24 14:46:03'),(46,'',5,'','','','https://s2.monidai.com/20220515/jqNVG7pD/index.m3u8','','4',1,'2022-05-25 05:37:03','2022-05-25 05:37:03'),(47,'',9,'','','','https://s2.monidai.com/20220525/LWVWnDf9/index.m3u8','','4',1,'2022-05-25 05:38:21','2022-05-25 05:38:21'),(48,'',1,'','','','https://b1.szjal.cn/20210103/bDNXuOrD/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:51:49','2022-06-03 04:51:49'),(49,'',2,'','','','https://b1.szjal.cn/20210103/HpA44OaS/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:52:34','2022-06-03 04:52:34'),(50,'',3,'','','','https://v5.szjal.cn/20210103/gRMs8WxQ/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:52:46','2022-06-03 04:52:46'),(51,'',4,'','','','https://b1.szjal.cn/20210103/vOLP3w7M/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:53:01','2022-06-03 04:53:14'),(52,'',5,'','','','https://b1.szjal.cn/20210103/K7D6Xhz4/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:53:35','2022-06-03 04:53:35'),(53,'',6,'','','','https://b1.szjal.cn/20210103/vxij97Ln/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:53:51','2022-06-03 04:53:51'),(54,'',7,'','','','https://b1.szjal.cn/20210103/OpJ5Sztb/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:54:10','2022-06-03 04:54:10'),(55,'',8,'','','','https://b1.szjal.cn/20210103/zCqZkHsx/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:54:33','2022-06-03 04:54:33'),(56,'',9,'','','','https://b1.szjal.cn/20210103/MIT8ZOai/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:54:47','2022-06-03 04:54:47'),(57,'',10,'','','','https://b1.szjal.cn/20210103/7RySu6zc/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:55:00','2022-06-03 04:55:00'),(58,'',11,'','','','https://b1.szjal.cn/20210103/r62pQU6U/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:55:14','2022-06-03 04:55:14'),(59,'',12,'','','','https://b1.szjal.cn/20210103/DiHJvgeD/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:55:26','2022-06-03 04:55:26'),(60,'',13,'','','','https://b1.szjal.cn/20210103/lyExKYMd/index.m3u8','','b2cf7c39-3f62-45c0-bbe1-a9522ff24cb7',1,'2022-06-03 04:55:41','2022-06-03 04:55:41');
/*!40000 ALTER TABLE `videos` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-09-07  8:11:04
