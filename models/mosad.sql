

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID(学号)',
  `username` varchar(50) NOT NULL COMMENT '用户姓名',
  `password` varchar(100) NOT NULL COMMENT '用户密码',
  
  PRIMARY KEY (`user_id`) 
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Table structure for course
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`  (
  `course_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '课程ID',
  `course_name` varchar(50) NOT NULL COMMENT '课程名称',
  `location` varchar(100) NOT NULL COMMENT '课程地址',
  `week_time` varchar(100) NOT NULL COMMENT '课程时间',
  `term_time` varchar(100) NOT NULL COMMENT '学期时间',
  `symbol` int NOT NULL COMMENT '标志',
  PRIMARY KEY (`course_id`) 
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Table structure for user_course
-- ----------------------------
DROP TABLE IF EXISTS `user_course`;
CREATE TABLE `user_course`  (
  `id` int NOT NULL AUTO_INCREMENT auto_increment,
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID(学号)',
  `course_id` int UNSIGNED NOT NULL COMMENT '课程ID',
  
  PRIMARY KEY (`id`)
  -- alter table 'course' add constraint 'user_course' foreign key ('course_id') references 'course' ('course_id')

) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- ----------------------------
-- Table structure for bill
-- ----------------------------
DROP TABLE IF EXISTS `bill`;
CREATE TABLE `bill`  (
  `bill_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '账单ID',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID(学号)',
  `money` int  NOT NULL COMMENT '账单价格',
  `bill_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '账单时间',
  `tag` varchar(100) NOT NULL COMMENT '账单类型',

  PRIMARY KEY (`bill_id`) 
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- ----------------------------

-- Table structure for homework
-- ----------------------------
DROP TABLE IF EXISTS `homework`;
CREATE TABLE `homework`  (
  `hw_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '作业ID',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户ID(学号)',
  `course_id` int UNSIGNED NOT NULL COMMENT '课程ID',
  `content` varchar(100) NOT NULL COMMENT '作业内容',
  `deadline` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '截止时间',

  PRIMARY KEY (`hw_id`) 
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- ----------------------------
