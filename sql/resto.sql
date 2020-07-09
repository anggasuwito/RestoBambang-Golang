-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: resto
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ekstra_menu`
--

DROP TABLE IF EXISTS `ekstra_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ekstra_menu` (
  `id_ekstra` varchar(100) NOT NULL,
  `nama_ekstra_menu` varchar(100) DEFAULT NULL,
  `harga_ekstra_menu` int DEFAULT NULL,
  PRIMARY KEY (`id_ekstra`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ekstra_menu`
--

LOCK TABLES `ekstra_menu` WRITE;
/*!40000 ALTER TABLE `ekstra_menu` DISABLE KEYS */;
/*!40000 ALTER TABLE `ekstra_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `harga_menu`
--

DROP TABLE IF EXISTS `harga_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `harga_menu` (
  `id_harga_menu` varchar(100) NOT NULL,
  `id_menu` varchar(100) DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `harga` int DEFAULT NULL,
  PRIMARY KEY (`id_harga_menu`),
  KEY `harga_menu_fk_idx` (`id_menu`),
  CONSTRAINT `harga_menu_fk` FOREIGN KEY (`id_menu`) REFERENCES `menu` (`id_menu`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `harga_menu`
--

LOCK TABLES `harga_menu` WRITE;
/*!40000 ALTER TABLE `harga_menu` DISABLE KEYS */;
INSERT INTO `harga_menu` VALUES ('54037c83-72d3-48a0-a180-7b848d19722f','b4e494e8-f418-447b-9035-b34f8c7f8cd8',NULL,15000),('b04f6ae6-472e-4b6a-97c5-4937a227064c','b4e494e8-f418-447b-9035-b34f8c7f8cd8',NULL,15000),('d0a0dcb0-f0e1-4efb-81a5-4254d0ee7687','ddbd7ded-be4e-4f73-9369-5a9bc9e9f959',NULL,15000);
/*!40000 ALTER TABLE `harga_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jenis`
--

DROP TABLE IF EXISTS `jenis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jenis` (
  `id_jenis` varchar(100) NOT NULL,
  `nama_jenis` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id_jenis`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jenis`
--

LOCK TABLES `jenis` WRITE;
/*!40000 ALTER TABLE `jenis` DISABLE KEYS */;
INSERT INTO `jenis` VALUES ('1ee4225f-1e16-4aaf-b2d9-1154f1c26a50','minuman');
/*!40000 ALTER TABLE `jenis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `list_menu`
--

DROP TABLE IF EXISTS `list_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `list_menu` (
  `id_list_menu` varchar(100) NOT NULL,
  `tanggal` date DEFAULT NULL,
  `id_jenis` varchar(100) DEFAULT NULL,
  `id_menu` varchar(100) DEFAULT NULL,
  `id_harga_menu` varchar(100) DEFAULT NULL,
  `id_stok_menu` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id_list_menu`),
  KEY `list_menu_fk_idx` (`id_menu`),
  KEY `list_jenis_fk_idx` (`id_jenis`),
  KEY `list_harga_fk_idx` (`id_harga_menu`),
  KEY `list_stok_fk_idx` (`id_stok_menu`),
  CONSTRAINT `list_harga_fk` FOREIGN KEY (`id_harga_menu`) REFERENCES `harga_menu` (`id_harga_menu`),
  CONSTRAINT `list_jenis_fk` FOREIGN KEY (`id_jenis`) REFERENCES `jenis` (`id_jenis`),
  CONSTRAINT `list_menu_fk` FOREIGN KEY (`id_menu`) REFERENCES `menu` (`id_menu`),
  CONSTRAINT `list_stok_fk` FOREIGN KEY (`id_stok_menu`) REFERENCES `stok_menu` (`id_stok_menu`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `list_menu`
--

LOCK TABLES `list_menu` WRITE;
/*!40000 ALTER TABLE `list_menu` DISABLE KEYS */;
/*!40000 ALTER TABLE `list_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `menu`
--

DROP TABLE IF EXISTS `menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `menu` (
  `id_menu` varchar(100) NOT NULL,
  `nama_menu` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id_menu`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `menu`
--

LOCK TABLES `menu` WRITE;
/*!40000 ALTER TABLE `menu` DISABLE KEYS */;
INSERT INTO `menu` VALUES ('b4e494e8-f418-447b-9035-b34f8c7f8cd8','nasi goreng'),('ddbd7ded-be4e-4f73-9369-5a9bc9e9f959','nasi baka goreng');
/*!40000 ALTER TABLE `menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stok_menu`
--

DROP TABLE IF EXISTS `stok_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `stok_menu` (
  `id_stok_menu` varchar(100) NOT NULL,
  `id_menu` varchar(100) DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  `stok` int DEFAULT NULL,
  PRIMARY KEY (`id_stok_menu`),
  KEY `stok_menu_idx` (`id_menu`),
  CONSTRAINT `stok_menu` FOREIGN KEY (`id_menu`) REFERENCES `menu` (`id_menu`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stok_menu`
--

LOCK TABLES `stok_menu` WRITE;
/*!40000 ALTER TABLE `stok_menu` DISABLE KEYS */;
/*!40000 ALTER TABLE `stok_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaksi`
--

DROP TABLE IF EXISTS `transaksi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaksi` (
  `id_transaksi` varchar(100) NOT NULL,
  `id_list_menu` varchar(100) DEFAULT NULL,
  `id_ekstra_menu` varchar(100) DEFAULT NULL,
  `tanggal` date DEFAULT NULL,
  PRIMARY KEY (`id_transaksi`),
  KEY `transaksi_ekstra_fk_idx` (`id_ekstra_menu`),
  KEY `transaksi_list_menu_idx` (`id_list_menu`),
  CONSTRAINT `transaksi_ekstra_menu` FOREIGN KEY (`id_ekstra_menu`) REFERENCES `ekstra_menu` (`id_ekstra`),
  CONSTRAINT `transaksi_list_menu` FOREIGN KEY (`id_list_menu`) REFERENCES `list_menu` (`id_list_menu`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaksi`
--

LOCK TABLES `transaksi` WRITE;
/*!40000 ALTER TABLE `transaksi` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaksi` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-10  5:41:13
