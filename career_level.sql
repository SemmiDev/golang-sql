CREATE TABLE IF NOT EXISTS `career_level` (
  `id` varchar(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `job_title` varchar(255) NOT NULL,
  `age` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `career_level` (`id`, `name`, `job_title`, `age`) VALUES
('A1', 'Sammi Aldhi Yanto',   'Technical Architect',           19),
('A2', 'Aditya Andika Putra', 'Chief Operating Officer',       19),
('A3', 'Abdul Raud', 'Chief Compliance Officer',               19),
('A4', 'Aditya Fauzan Nul Haq', 'Chief Executive Officer',     19),
('A5', 'Ayatullah Ramadhan Jacoeb', 'Chief Financial Officer', 19),
('A6', 'Dandi Arnanda', 'Chief Technology Officer',            19),
('A7', 'ALharidt Mahmudi', 'Chief Decision Officer',           19),
('A8', 'Gusnur', 'Chief Marketing Officer',                    19),
('A9', 'Fadhil Ramadhan', 'Principal Engineer',                19);

ALTER TABLE `career_level` ADD PRIMARY KEY (`id`);