-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3307
-- Время создания: Янв 23 2022 г., 15:30
-- Версия сервера: 10.3.22-MariaDB
-- Версия PHP: 7.1.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `totalscore`
--

-- --------------------------------------------------------

--
-- Структура таблицы `totals`
--

CREATE TABLE `totals` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `total_first_place` int(11) DEFAULT NULL,
  `total_secondt_place` int(11) DEFAULT NULL,
  `total_third_place` int(11) DEFAULT NULL,
  `total_top10` int(11) DEFAULT NULL,
  `total_participate` int(11) DEFAULT NULL,
  `total_points` int(11) DEFAULT NULL,
  `total_score` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Дамп данных таблицы `totals`
--

INSERT INTO `totals` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `total_first_place`, `total_secondt_place`, `total_third_place`, `total_top10`, `total_participate`, `total_points`, `total_score`) VALUES
(326, '2022-01-16 21:18:18', '2022-01-16 21:18:18', NULL, 'gig51600', 1, 0, 0, 1, 1, 25, 30),
(327, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'daaleksandrov', 0, 1, 0, 1, 1, 19, 27),
(328, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'TravelerY', 0, 0, 1, 1, 1, 14, 24),
(329, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'TupikDenis', 0, 0, 0, 1, 1, 10, 24),
(330, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Chessmaste', 0, 0, 0, 1, 1, 7, 21),
(331, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Poltergeist777', 0, 0, 0, 1, 1, 5, 20),
(332, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Andreikaluga222', 0, 0, 0, 1, 1, 4, 18),
(333, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'RakhmonovNazar', 0, 0, 0, 1, 1, 3, 11),
(334, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Vladikbrest', 0, 0, 0, 1, 1, 2, 11),
(335, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Lebedevms', 0, 0, 0, 1, 1, 1, 8),
(336, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'stas527', 0, 0, 0, 0, 1, 0, 8),
(337, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'ggtikoblt', 0, 0, 0, 0, 1, 0, 7),
(338, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'MAXim_2016', 0, 0, 0, 0, 1, 0, 6),
(339, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'PashokKotov', 0, 0, 0, 0, 1, 0, 0),
(340, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'NastyaMorozova2015', 0, 0, 0, 0, 1, 0, 0),
(341, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Alex_Lavri', 0, 0, 0, 0, 1, 0, 0),
(342, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Bandirn', 0, 0, 0, 0, 1, 0, 0),
(343, '2022-01-16 21:18:19', '2022-01-16 21:18:19', NULL, 'Demon702', 0, 0, 0, 0, 1, 0, 0);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `totals`
--
ALTER TABLE `totals`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_totals_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `totals`
--
ALTER TABLE `totals`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=344;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
