-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 27, 2023 at 04:15 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kreditplus`
--

-- --------------------------------------------------------

--
-- Table structure for table `konsumen`
--

CREATE TABLE `konsumen` (
  `id_konsumen` bigint(20) UNSIGNED NOT NULL,
  `nik` longtext DEFAULT NULL,
  `gender` longtext DEFAULT NULL,
  `full_name` longtext DEFAULT NULL,
  `legal_name` longtext DEFAULT NULL,
  `tempat_lahir` longtext DEFAULT NULL,
  `tanggal_lahir` longtext DEFAULT NULL,
  `gaji` bigint(20) UNSIGNED DEFAULT NULL,
  `foto_ktp` longtext DEFAULT NULL,
  `foto_selfie` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `konsumen`
--

INSERT INTO `konsumen` (`id_konsumen`, `nik`, `gender`, `full_name`, `legal_name`, `tempat_lahir`, `tanggal_lahir`, `gaji`, `foto_ktp`, `foto_selfie`) VALUES
(1, '1922511504970002', 'Laki-laki', 'Aria Shabry', 'Aria Shabry', 'Padang', '15-04-1997', 5000000, 'http://kreditplus.com/storage/data/image/1922511504970002_300.jpg', 'http://kreditplus.com/storage/data/image/1922511504970002_600.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `pinjaman`
--

CREATE TABLE `pinjaman` (
  `id_pinjaman` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `amount` bigint(20) NOT NULL,
  `tenor` bigint(20) NOT NULL,
  `status` longtext NOT NULL,
  `totaldebt` bigint(20) NOT NULL,
  `interest_rate` double NOT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `pinjaman`
--

INSERT INTO `pinjaman` (`id_pinjaman`, `user_id`, `amount`, `tenor`, `status`, `totaldebt`, `interest_rate`, `created_at`) VALUES
(1, 1, 500000, 5, 'Approved', 0, 0.02, '2023-06-26 23:47:55.707');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id_transaction` bigint(20) UNSIGNED NOT NULL,
  `loan_id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `amount` bigint(20) NOT NULL,
  `payment_date` datetime(3) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `konsumen`
--
ALTER TABLE `konsumen`
  ADD PRIMARY KEY (`id_konsumen`);

--
-- Indexes for table `pinjaman`
--
ALTER TABLE `pinjaman`
  ADD PRIMARY KEY (`id_pinjaman`),
  ADD KEY `fk_konsumen_pinjamans` (`user_id`);

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id_transaction`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `konsumen`
--
ALTER TABLE `konsumen`
  MODIFY `id_konsumen` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `pinjaman`
--
ALTER TABLE `pinjaman`
  MODIFY `id_pinjaman` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id_transaction` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `pinjaman`
--
ALTER TABLE `pinjaman`
  ADD CONSTRAINT `fk_konsumen_pinjamans` FOREIGN KEY (`user_id`) REFERENCES `konsumen` (`id_konsumen`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
