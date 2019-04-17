# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 10.9.22.163 (MySQL 5.5.24-ucloudrel1-log)
# Database: devops
# Generation Time: 2019-04-17 16:16:34 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table app
# ------------------------------------------------------------

CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_important` int(11) NOT NULL DEFAULT '0',
  `name` varchar(32) NOT NULL DEFAULT '',
  `url` varchar(256) NOT NULL DEFAULT '',
  `desc` varchar(512) NOT NULL DEFAULT '',
  `repository_url` varchar(512) NOT NULL DEFAULT '',
  `deploy_dir` varchar(512) NOT NULL DEFAULT '',
  `monitor_url` varchar(512) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `repository_id` int(11) NOT NULL DEFAULT '0',
  `port` int(11) NOT NULL DEFAULT '0',
  `tag` varchar(32) NOT NULL DEFAULT '',
  `internal_url` varchar(256) NOT NULL DEFAULT '',
  `update_code_cmd` varchar(512) NOT NULL DEFAULT '',
  `reload_service_cmd` varchar(512) NOT NULL DEFAULT '',
  `check_service_cmd` varchar(512) NOT NULL DEFAULT '',
  `cmd_name` varchar(512) NOT NULL DEFAULT '',
  `cmd_dir` varchar(512) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_name_idx` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table cloud_available_region
# ------------------------------------------------------------

CREATE TABLE `cloud_available_region` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `region` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_idx` (`region`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table computer
# ------------------------------------------------------------

CREATE TABLE `computer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cpu` int(11) NOT NULL DEFAULT '0',
  `ram` int(11) NOT NULL DEFAULT '0',
  `private_ip` varchar(32) NOT NULL DEFAULT '',
  `public_ip` varchar(32) NOT NULL DEFAULT '',
  `host_id` varchar(32) DEFAULT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `host_tag` varchar(32) NOT NULL DEFAULT '',
  `host_name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `private_ip_idx` (`private_ip`),
  UNIQUE KEY `host_id_idx` (`host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table computer_role
# ------------------------------------------------------------

CREATE TABLE `computer_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `register_status` int(11) unsigned NOT NULL DEFAULT '0',
  `host_id` varchar(32) NOT NULL DEFAULT '',
  `app_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_host_id_app_id` (`host_id`,`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table deploy
# ------------------------------------------------------------

CREATE TABLE `deploy` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `commit_id` varchar(64) NOT NULL DEFAULT '',
  `commit_tag` varchar(64) NOT NULL DEFAULT '',
  `status` varchar(12) NOT NULL DEFAULT '',
  `app_id` int(11) NOT NULL DEFAULT '0',
  `desc` varchar(256) NOT NULL DEFAULT '',
  `rollback_tag` varchar(64) NOT NULL DEFAULT '',
  `rollback_id` varchar(64) NOT NULL DEFAULT '',
  `max` int(11) NOT NULL DEFAULT '0',
  `started_at` datetime DEFAULT NULL,
  `finished_at` datetime DEFAULT NULL,
  `interval` int(11) NOT NULL DEFAULT '0',
  `hosts` varchar(256) NOT NULL DEFAULT '',
  `binary_url` varchar(256) NOT NULL DEFAULT '' COMMENT '二进制包下载地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table disk
# ------------------------------------------------------------

CREATE TABLE `disk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `size` int(11) NOT NULL DEFAULT '0',
  `left` int(11) NOT NULL DEFAULT '0',
  `computer_id` int(11) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table domain
# ------------------------------------------------------------

CREATE TABLE `domain` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `name` varchar(256) NOT NULL DEFAULT '',
  `private` tinyint(4) NOT NULL DEFAULT '0',
  `host` varchar(256) NOT NULL DEFAULT '',
  `ip` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table service
# ------------------------------------------------------------

CREATE TABLE `service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `is_important` int(11) NOT NULL DEFAULT '0',
  `name` varchar(32) NOT NULL DEFAULT '',
  `url` varchar(256) NOT NULL DEFAULT '',
  `desc` varchar(512) NOT NULL DEFAULT '',
  `repository_url` varchar(512) NOT NULL DEFAULT '',
  `deploy_dir` varchar(512) NOT NULL DEFAULT '',
  `monitor_url` varchar(512) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table tag
# ------------------------------------------------------------

CREATE TABLE `tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '''''',
  `kind` varchar(32) NOT NULL DEFAULT '''''',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_idex` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
