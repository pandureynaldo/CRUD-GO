/*
Navicat MySQL Data Transfer

Source Server         : Local
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : sa_db

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2020-01-25 10:34:10
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `article`
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `content` text,
  `created_at` datetime DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` varchar(255) DEFAULT NULL,
  `status` smallint(6) DEFAULT '0',
  `tag` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES ('1', 'Hello World', 'Ini adalah artikel pertama', '2020-01-20 23:53:24', 'admin', '2020-01-21 08:21:23', 'admin', '1', 'Test');
INSERT INTO `article` VALUES ('2', 'Ribut Ibu-ibu 2 Ekor Ayam Dihargai Rp 800 Ribu', 'Medan - Ibu-ibu terlibat keributan karena dua ekor ayam yang dihargai Rp 800 ribu. Video keributannya viral di media sosial (medsos).\r\n\r\nVideo ini banyak tersebar dan menuai beragam tanggapan. Diketahui peristiwa ribut-ribut itu terjadi di Sidikalang, ibu kota Kabupaten Diari, Sumatera Utara (Sumut).\r\n\r\nPihak restoran angkat bicara menjelaskan video yang viral. Pemilik Rumah Makan (RM) Malau Napinadar Sidikalang, Lambok Roy Marteen Malau mengatakan peristiwa itu terjadi sehari setelah tahun baru, Kamis (2/1).', '2020-01-21 08:15:11', 'user', '2020-01-21 08:21:30', 'user', '1', 'Berita');
INSERT INTO `article` VALUES ('3', 'Siapkah Kau Tuk Jatuh Cinta Lagi - HiVi', 'ketika ku mendengar bahwa\r\nkini kau tak lagi dengannya\r\ndalam benakku timbul tanya\r\n\r\nmasihkah ada dia di hatimu bertahta\r\natau ini saat bagiku\r\nuntuk singgah di hatimu\r\n\r\nnamun siapkah kau tuk jatuh cinta lagi\r\n\r\nmeski bibir ini tak berkata\r\nbukan berarti ku tak merasa\r\nada yang berbeda di antara kita\r\n\r\ndan tak mungkin ku melewatkanmu\r\nhanya karena diriku tak mampu untuk bicara\r\nbahwa aku inginkan kau ada di hidupku\r\n\r\nkini ku tak lagi dengannya\r\nsudah tak ada lagi rasa antara aku dengan dia (dengan dia)\r\nsiapkah kau bertahta di hatiku, adinda\r\nkarena ini saat yang tepat untuk singgah di hatiku\r\nnamun siapkah kau tuk jatuh cinta lagi oooh\r\n\r\nmeski bibir ini tak berkata\r\nbukan berarti ku tak merasa ada yang berbeda di antara kita\r\ndan tak mungkin ku melewatkanmu hanya karena\r\ndiriku tak mampu untuk bicara bahwa aku inginkan kau ada di hidupku', '2020-01-21 21:53:22', 'user', '2020-01-25 00:28:43', 'user', '1', 'Liri');
INSERT INTO `article` VALUES ('5', 'Dimana kita berada', 'Dimana anak kambing saya, anak kambing saya ada dirumah tangga', '2020-01-24 21:31:54', 'user', '2020-01-24 21:31:54', 'user', '0', 'Lagu');
INSERT INTO `article` VALUES ('6', 'Hidup  - Khalil Gibran', 'Hidup adalah kegelapan jika tanpa hasrat dan keinginan. Dan semua hasrat serta keinginan adalah buta, jika tidak disertai pengetahuan. Dan pengetahuan adalah hampa jika tidak diikuti pelajaran. Dan setiap pelajaran akan sia-sia jika tidak disertai cinta.', '2020-01-24 22:25:51', 'user', '2020-01-24 22:25:51', 'user', '0', 'Sajak');
INSERT INTO `article` VALUES ('7', 'Jakarta Diguyur Hujan, Mulai Muncul Genangan di Sejumlah Titik', 'Jakarta -\r\nSejumlah wilayah DKI Jakarta diguyur hujan deras pagi ini. Genangan air muncul di beberapa titik di Ibu Kota.\r\n\r\nAkun Twitter @TMCPoldaMetro mengabarkan ada genangan air di Jalan DI Panjaitan, Jakarta Timur. Per pukul 09.55 WIB, TMC Polda Metro Jaya mengabarkan kendaraan masih bisa melewati jalan tersebut.\r\n\r\n\"09.55 Situasi lalu lintas di Jl DI Panjaitan Jaktim terpantau ada genangan air sekitar 10cm dan masih bisa di lintasi pengendara,\" tulis TMC Polda Metro.', '2020-01-24 22:37:37', 'user', '0000-00-00 00:00:00', '2020-01-24 22:37:37', '0', 'Berita');
INSERT INTO `article` VALUES ('8', 'Judul baru', 'test', '2020-01-24 23:01:43', 'user', '2020-01-25 00:45:22', '2020-01-24 23:01:43', '0', 'yes');

-- ----------------------------
-- Table structure for `contact`
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `message` text,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of contact
-- ----------------------------
INSERT INTO `contact` VALUES ('1', 'user', 'pandureynaldo02@Gmail.com', 'Halo ini pesan pertama saya', '2020-01-25 01:00:35');
INSERT INTO `contact` VALUES ('2', 'user', 'client@gmail.com', 'kami coba coba saja', '2020-01-25 01:01:57');

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT NULL,
  `password` text,
  `created_at` datetime DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `status` smallint(2) DEFAULT NULL COMMENT '1=aktif, 0 = nonaktif',
  `role` varchar(255) DEFAULT NULL COMMENT 'admin, user',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'admin', 'admin', '2020-01-16 21:42:40', '2020-01-20 20:49:45', '1', 'admin');
INSERT INTO `users` VALUES ('2', 'user', 'user', '2020-01-16 21:43:05', '2020-01-20 20:49:51', '1', 'user');
