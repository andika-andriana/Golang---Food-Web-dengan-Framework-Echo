/*
SQLyog Ultimate v12.5.1 (64 bit)
MySQL - 10.4.10-MariaDB : Database - db_pizza
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_pizza` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `db_pizza`;

/*Table structure for table `tbl_menu` */

DROP TABLE IF EXISTS `tbl_menu`;

CREATE TABLE `tbl_menu` (
  `id_menu` int(11) NOT NULL AUTO_INCREMENT,
  `nama_menu` varchar(100) DEFAULT NULL,
  `deskripsi` text DEFAULT NULL,
  `url_gambar` text DEFAULT NULL,
  `jenis` enum('Pizza','Pasta','Appetizer','Dessert') DEFAULT NULL,
  `harga` double DEFAULT NULL,
  PRIMARY KEY (`id_menu`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*Data for the table `tbl_menu` */

insert  into `tbl_menu`(`id_menu`,`nama_menu`,`deskripsi`,`url_gambar`,`jenis`,`harga`) values 
(1,'Beef Lasagna','Beef Lasagna Adalah Pasta Terbaik Di Kota Ini','beeflasagna.png','Pasta',89000),
(2,'Pizza Cheese Burger','Pizza Cheese Burger Adalah Pizza Dengan Campuran Burger Dan Keju','cheeseburger.png','Pizza',210000),
(3,'Chicken Canneloni','Chicken Canneloni Adalah Pasta Dengan Campuran Ayam','chickencanneloni.png','Pasta',75000),
(4,'Pizza Chicken Loverz','Pizza Chicken Loverz adalah pizza untuk pecinta ayam','chickenloverz.png','Pizza',240000),
(5,'Neapolitan Dream','Neapolitan Dream adalah menu yang cocok untuk hidangan penutup','neapolitandream.png','Dessert',75000),
(6,'Potato Wedges','Potato Wedges  adalah menu yang cocok untuk makanan ringan','potatoweges.png','Appetizer',17000);

/*Table structure for table `tbl_order` */

DROP TABLE IF EXISTS `tbl_order`;

CREATE TABLE `tbl_order` (
  `id_order` int(11) NOT NULL AUTO_INCREMENT,
  `id_menu` int(11) DEFAULT NULL,
  `nama_pemesan` varchar(100) DEFAULT NULL,
  `nomor_telephone` varchar(30) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `jumlah_pemesan` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_order`),
  KEY `id_menu` (`id_menu`),
  CONSTRAINT `tbl_order_ibfk_1` FOREIGN KEY (`id_menu`) REFERENCES `tbl_menu` (`id_menu`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

/*Data for the table `tbl_order` */

insert  into `tbl_order`(`id_order`,`id_menu`,`nama_pemesan`,`nomor_telephone`,`alamat`,`jumlah_pemesan`) values 
(4,1,'Andika','081234567890','Jl. Manggarai Jakarta',20),
(5,3,'Kiki','081234556789','Jl. Studio Alam Depok',12),
(6,5,'Budi','08312345678','Jl. Buana 2 Jakarta Timur',7),
(7,5,'Yunan','08124567892','Jl. Bina Negara, Surabaya',19),
(8,1,'Umar','08145678922','Jl. Nusa Indah, Jakarta',7),
(9,2,'Via Shakinah','08112233445','Jl. Dermaga V, Jakarta',15),
(11,4,'Andika Andriana','08112233445','Jl. xxxyyyxxx',5),
(12,6,'Bayu','081234556789','Jl. YYYZZZYYY',10);

/*Table structure for table `vw_totalorder` */

DROP TABLE IF EXISTS `vw_totalorder`;

/*!50001 DROP VIEW IF EXISTS `vw_totalorder` */;
/*!50001 DROP TABLE IF EXISTS `vw_totalorder` */;

/*!50001 CREATE TABLE  `vw_totalorder`(
 `id_menu` int(11) ,
 `nama_menu` varchar(100) ,
 `deskripsi` text ,
 `url_gambar` text ,
 `jenis` enum('Pizza','Pasta','Appetizer','Dessert') ,
 `harga` double ,
 `total_older` decimal(32,0) 
)*/;

/*View structure for view vw_totalorder */

/*!50001 DROP TABLE IF EXISTS `vw_totalorder` */;
/*!50001 DROP VIEW IF EXISTS `vw_totalorder` */;

/*!50001 CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `vw_totalorder` AS select `tbl_menu`.`id_menu` AS `id_menu`,`tbl_menu`.`nama_menu` AS `nama_menu`,`tbl_menu`.`deskripsi` AS `deskripsi`,`tbl_menu`.`url_gambar` AS `url_gambar`,`tbl_menu`.`jenis` AS `jenis`,`tbl_menu`.`harga` AS `harga`,coalesce(sum(`tbl_order`.`jumlah_pemesan`),0) AS `total_older` from (`tbl_menu` left join `tbl_order` on(`tbl_order`.`id_menu` = `tbl_menu`.`id_menu`)) group by `tbl_menu`.`id_menu` */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
