-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Feb 01, 2023 at 03:39 PM
-- Server version: 5.7.36
-- PHP Version: 7.4.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_moladin_project_learning_week`
--

-- --------------------------------------------------------

--
-- Table structure for table `albums`
--

DROP TABLE IF EXISTS `albums`;
CREATE TABLE IF NOT EXISTS `albums` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `year` bigint(20) UNSIGNED DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `albums`
--

INSERT INTO `albums` (`id`, `name`, `year`, `created_at`, `updated_at`) VALUES
(1, 'One Ok Rock', 2023, '2023-02-01 06:23:19.183', '2023-02-01 06:23:19.183'),
(2, 'Sheila On 7', 2023, '2023-02-01 06:25:19.183', '2023-02-01 06:25:19.183');

-- --------------------------------------------------------

--
-- Table structure for table `songs`
--

DROP TABLE IF EXISTS `songs`;
CREATE TABLE IF NOT EXISTS `songs` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `album_id` bigint(20) UNSIGNED DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(100) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_album_id` (`album_id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `songs`
--

INSERT INTO `songs` (`id`, `album_id`, `title`, `author`, `created_at`, `updated_at`) VALUES
(1, 1, 'The Beginning', 'One Ok Rock', '2023-02-01 06:24:50.343', '2023-02-01 06:24:50.343'),
(2, 1, 'Heartache', 'One Ok Rock', '2023-02-01 06:25:43.100', '2023-02-01 06:25:43.100'),
(3, 1, 'Chaosmyth', 'One Ok Rock', '2023-02-01 06:26:00.694', '2023-02-01 06:26:00.694'),
(4, 1, 'Cry Out', 'One Ok Rock', '2023-02-01 06:26:15.457', '2023-02-01 06:26:15.457'),
(5, 1, 'Clock\'s Strike', 'One Ok Rock', '2023-02-01 06:26:27.296', '2023-02-01 06:26:27.296'),
(6, 2, 'Dan', 'Duta', '2023-02-01 06:27:27.296', '2023-02-01 06:27:27.296'),
(7, 2, 'Tunggu Aku Di Jakarta', 'Duta', '2023-02-01 06:27:40.296', '2023-02-01 06:27:40.296');

--
-- Constraints for dumped tables
--

--
-- Constraints for table `songs`
--
ALTER TABLE `songs`
  ADD CONSTRAINT `fk_albums_songs` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`),
  ADD CONSTRAINT `songs_ibfk_1` FOREIGN KEY (`album_id`) REFERENCES `albums` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
