/*
 Navicat Premium Dump SQL

 Source Server         : db@local
 Source Server Type    : MySQL
 Source Server Version : 50726 (5.7.26)
 Source Host           : localhost:3306
 Source Schema         : dbetrain

 Target Server Type    : MySQL
 Target Server Version : 50726 (5.7.26)
 File Encoding         : 65001

 Date: 29/10/2025 15:36:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gen_column
-- ----------------------------
DROP TABLE IF EXISTS `gen_column`;
CREATE TABLE `gen_column`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `table_id` int(11) NULL DEFAULT NULL,
  `data_column` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `column_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `data_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `data_type_lang` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `data_size` int(10) UNSIGNED NULL DEFAULT NULL,
  `title` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `dom_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_primary_key` tinyint(4) NULL DEFAULT NULL,
  `auto_increment` tinyint(4) NULL DEFAULT NULL,
  `is_index` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_nullable` tinyint(4) NULL DEFAULT NULL,
  `placeholder` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `option` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `search_able` tinyint(4) NULL DEFAULT 1,
  `create_able` tinyint(4) NULL DEFAULT 1,
  `update_able` tinyint(4) NULL DEFAULT 1,
  `sort_able` tinyint(4) NULL DEFAULT 1,
  `hidden` tinyint(4) NULL DEFAULT 0,
  `serializer` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `reference` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '外键',
  `rawdata` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_gen_column_table`(`table_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_column
-- ----------------------------

-- ----------------------------
-- Table structure for gen_mapper
-- ----------------------------
DROP TABLE IF EXISTS `gen_mapper`;
CREATE TABLE `gen_mapper`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `lang` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '语言',
  `dbtype` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据类型',
  `fromtype` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '原始类型',
  `totype` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '新类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_mapper
-- ----------------------------

-- ----------------------------
-- Table structure for gen_project
-- ----------------------------
DROP TABLE IF EXISTS `gen_project`;
CREATE TABLE `gen_project`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目名称用作服务发现',
  `title` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目标题',
  `dbtype` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据库类型',
  `dbname` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据库名称',
  `dsn` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据连接',
  `prefix` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '数据库前缀',
  `author` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '作者',
  `tpl_id` mediumint(9) NULL DEFAULT NULL COMMENT '模板ID',
  `package` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名称',
  `sorting` bigint(20) NULL DEFAULT 0 COMMENT '排序位',
  `dirsave` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '项目存放目录',
  `lang` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'golang' COMMENT '默认语言',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_project
-- ----------------------------

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS `gen_table`;
CREATE TABLE `gen_table`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `project_id` int(11) NULL DEFAULT NULL COMMENT '项目ID',
  `name` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '标题',
  `module` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应模块',
  `method` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '支持的方法',
  `columns` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '全部字段',
  `types` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '支持的数据类型',
  `tree_option` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '树结构配置',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_gen_table_project`(`project_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_table
-- ----------------------------

-- ----------------------------
-- Table structure for gen_tpldata
-- ----------------------------
DROP TABLE IF EXISTS `gen_tpldata`;
CREATE TABLE `gen_tpldata`  (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `tplset_id` int(11) NULL DEFAULT NULL COMMENT '所属群组',
  `title` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板名称',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '模板内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_tpldata
-- ----------------------------

-- ----------------------------
-- Table structure for gen_tplset
-- ----------------------------
DROP TABLE IF EXISTS `gen_tplset`;
CREATE TABLE `gen_tplset`  (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板名称',
  `memo` varchar(220) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板描述',
  `status` smallint(6) NULL DEFAULT NULL COMMENT '模板状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_tplset
-- ----------------------------

-- ----------------------------
-- Table structure for sys_area
-- ----------------------------
DROP TABLE IF EXISTS `sys_area`;
CREATE TABLE `sys_area`  (
  `area_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'areaId',
  `pid` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '父级ID',
  `name` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `type` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
  `rank` bigint(20) NULL DEFAULT NULL COMMENT '级别',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `deleted` int(11) NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`area_id`) USING BTREE,
  INDEX `idx_sys_area_pid`(`pid`) USING BTREE,
  INDEX `idx_sys_area_type`(`type`) USING BTREE,
  INDEX `idx_sys_area_rank`(`rank`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_area
-- ----------------------------

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示',
  `value` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '值',
  `org_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '组织ID',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `deleted` int(11) NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_config_org_id`(`org_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, 'site.title', '站点名称', '湖南云立方', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (2, 'site.memo', '站点描述', '明德、精业、求实、创新', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (3, 'site.orgname', '公司名称', '湖南云立方智能科技有限公司', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (4, 'site.orgno', '营业执照', '91430111MA4R1P7M7B', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (5, 'site.orgaddr', '公司地址', '湖南省长沙市岳麓区岳麓街道后湖艺术园D区62栋210号', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (6, 'site.linkname', '联系人', '胡文林', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (7, 'site.linkphone', '联系电话', '15367151352', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (8, 'site.linkemail', '邮箱地址', 'winlion@turingdance.com', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (9, 'site.linkqq', 'qq地址', '271151388', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (10, 'site.qrcodewx', '二维码', 'https://www.turingdance.com/assets/images/winlion.jpg', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (11, 'site.logo0', '128*128logo', 'https://www.turingdance.com/assets/images/logo.png', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (12, 'site.logo1', '128*320logo', 'https://www.turingdance.com/assets/images/logo.png', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);
INSERT INTO `sys_config` VALUES (13, 'site.beianno', '备案号', '湘ICP备16002978号-4', 1, '2025-05-29 14:12:22', '2025-05-29 14:12:22', 0);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门ID',
  `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称名称',
  `pid` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父级ID',
  `master_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '部门负责人ID',
  `status` int(11) NULL DEFAULT 1 COMMENT '部门状态',
  `level` int(11) NULL DEFAULT 1 COMMENT '部门等级',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dept_pid`(`pid`) USING BTREE,
  INDEX `idx_sys_dept_master_id`(`master_id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名字',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '值得',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `deleted` int(11) NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
INSERT INTO `sys_dict` VALUES (1, 'status_problem', '问题状态', '[{\"value\":\"idle\",\"label\":\"草稿\",\"score\":0},{\"value\":\"confirmed\",\"label\":\"待处理\",\"score\":0},{\"value\":\"progress\",\"label\":\"整改中\",\"score\":0},{\"value\":\"complete\",\"label\":\"已完成\",\"score\":0}]', '2023-09-05 15:18:09', '2025-05-29 15:17:08', 1);
INSERT INTO `sys_dict` VALUES (2, 'trade', '行业', '[{\"value\":\"5\",\"label\":\"批发和零售业\",\"score\":0},{\"value\":\"10\",\"label\":\"租赁和商务服务业\",\"score\":0}]', '2024-05-26 21:14:14', '2025-05-29 15:32:34', 1);
INSERT INTO `sys_dict` VALUES (3, 'category', '分类', '[{\"value\":\"1\",\"label\":\"水\",\"score\":0},{\"value\":\"2\",\"label\":\"土\",\"score\":0},{\"value\":\"3\",\"label\":\"气\",\"score\":0}]', '2024-05-27 16:31:26', '2025-05-29 15:30:48', 1);
INSERT INTO `sys_dict` VALUES (4, 'flow_status', '流程状态', '[{\"value\":\"idle\",\"label\":\"待处理\",\"score\":0},{\"value\":\"stop\",\"label\":\"已停止\",\"score\":0},{\"value\":\"approvaling\",\"label\":\"审批中\",\"score\":0},{\"value\":\"ok\",\"label\":\"已完成\",\"score\":0}]', '2024-06-02 17:46:12', '2025-05-29 15:16:55', 1);
INSERT INTO `sys_dict` VALUES (5, 'sys_rights_type', '权限类型', '[{\"value\":\"group\",\"label\":\"业务分组\",\"score\":0},{\"value\":\"view\",\"label\":\"页面菜单\",\"score\":0},{\"value\":\"widget\",\"label\":\"按钮操作\",\"score\":0},{\"value\":\"api\",\"label\":\"服务接口\",\"score\":0}]', '2024-11-03 08:40:10', '2024-11-03 08:40:10', 0);
INSERT INTO `sys_dict` VALUES (6, 'state_status', '审核状态', '[{\"value\":\"idle\",\"label\":\"待处理\",\"score\":0},{\"value\":\"rejected\",\"label\":\"已拒绝\",\"score\":0},{\"value\":\"resolved\",\"label\":\"已同意\",\"score\":0}]', '2024-11-05 12:22:10', '2025-05-29 15:33:42', 1);
INSERT INTO `sys_dict` VALUES (7, 'sys_enable', '使能状态', '[{\"value\":\"1\",\"label\":\"启用\",\"score\":0},{\"value\":\"0\",\"label\":\"停用\",\"score\":0}]', '2025-05-29 15:37:41', '2025-05-29 15:37:41', 0);
INSERT INTO `sys_dict` VALUES (8, 'sys_media_type', '资源类型', '[{\"value\":\"video\",\"label\":\"视频资源\",\"score\":0},{\"value\":\"image\",\"label\":\"图片资源\",\"score\":0},{\"value\":\"pdf\",\"label\":\"PDF资源\",\"score\":0},{\"value\":\"docx\",\"label\":\"word资源\",\"score\":0},{\"value\":\"pptx\",\"label\":\"演示文档\",\"score\":0},{\"value\":\"zip\",\"label\":\"压缩文件\",\"score\":0},{\"value\":\"audio\",\"label\":\"音频文件\",\"score\":0},{\"value\":\"exe\",\"label\":\"可执行文件\",\"score\":0},{\"value\":\"url\",\"label\":\"外链资源\",\"score\":0},{\"value\":\"unkown\",\"label\":\"未知文件\",\"score\":0}]', '2025-06-01 10:06:52', '2025-06-01 10:06:52', 0);
INSERT INTO `sys_dict` VALUES (9, 'sys_lesson_cate', '课程分类', '[{\"value\":\"metal\",\"label\":\"虚拟仿真实训系统\",\"score\":0},{\"value\":\"acrylic\",\"label\":\"3D打印和装配过程\",\"score\":0},{\"value\":\"soft\",\"label\":\"教学公开课\",\"score\":0}]', '2025-06-10 21:47:43', '2025-06-10 21:47:43', 0);

-- ----------------------------
-- Table structure for sys_org
-- ----------------------------
DROP TABLE IF EXISTS `sys_org`;
CREATE TABLE `sys_org`  (
  `org_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '名称',
  `pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'logo',
  `memo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'sologo',
  `status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '使能状态',
  `linker` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系人',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '详细描述',
  `tel` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系方式',
  `province` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '省份',
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '市',
  `district` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '区',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '详细地址',
  `lat` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '纬度',
  `lng` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '经度',
  `cate` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型',
  `org_no` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '统一社会信用代码',
  `user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '机构创建者',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `deleted` int(11) NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`org_id`) USING BTREE,
  INDEX `idx_sys_org_user_id`(`user_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_org
-- ----------------------------
INSERT INTO `sys_org` VALUES (1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 0);

-- ----------------------------
-- Table structure for sys_rights
-- ----------------------------
DROP TABLE IF EXISTS `sys_rights`;
CREATE TABLE `sys_rights`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `biz` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '业务标识符',
  `title` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限名称',
  `pid` bigint(20) NULL DEFAULT NULL COMMENT '父级',
  `uri` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源路径',
  `type` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资源类型',
  `sort_index` decimal(10, 0) NULL DEFAULT 0 COMMENT '排序位',
  `component` varchar(240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '組件路径',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
  `always_show` smallint(6) NULL DEFAULT NULL COMMENT '是否始终展示',
  `hidden` smallint(6) NULL DEFAULT NULL COMMENT '是否隐藏',
  `redirect` varchar(140) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '如果跳转',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_sys_rights_biz`(`biz`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_rights
-- ----------------------------
INSERT INTO `sys_rights` VALUES (1, 'board', '控制面板', 0, '/board', 'group', 0, 'Layout', 'dashboard', 0, 0, '/dashboard');
INSERT INTO `sys_rights` VALUES (2, 'dashboard', '控制面板', 1, '/dashboard', 'view', 1, 'dashboard/index', NULL, 0, 0, NULL);
INSERT INTO `sys_rights` VALUES (3, 'system', '系统管理', 0, '/system', 'group', 1000, 'Layout', 'setting', 0, 0, NULL);
INSERT INTO `sys_rights` VALUES (4, 'system:team', '组织架构', 3, '/system/team', 'view', 1, 'sys/team/index', NULL, 0, 0, NULL);
INSERT INTO `sys_rights` VALUES (5, 'system:role', '角色管理', 3, '/system/role', 'view', 0, 'sys/role/index', NULL, 0, 0, NULL);
INSERT INTO `sys_rights` VALUES (6, 'system:rights', '权限管理', 3, '/system/rights', 'view', 0, 'sys/rights/index', NULL, 0, 0, NULL);
INSERT INTO `sys_rights` VALUES (7, 'system:dict', '系统字典', 3, '/system/dict', 'view', 0, 'sys/dict/index', NULL, 0, 0, NULL);
INSERT INTO `sys_rights` VALUES (8, 'system:accpro', '个人资料', 3, '/system/accpro', 'view', 0, 'sys/acc/profile', NULL, 0, 1, NULL);
INSERT INTO `sys_rights` VALUES (9, 'system:teampro', '团队资料', 3, '/system/teampro', 'view', 0, 'sys/team/profile', NULL, 0, 1, NULL);
INSERT INTO `sys_rights` VALUES (21, 'system:role:grant', '角色授权', 5, 'post:/sys/role/grant', 'widget', 0, NULL, NULL, NULL, 0, NULL);
INSERT INTO `sys_rights` VALUES (22, 'ecls', '实验教学', 0, '/ecls', 'group', 4, 'Layout', 'monitor', NULL, 0, '/ecls/lesson');
INSERT INTO `sys_rights` VALUES (23, 'ecls:lesson', '课件资源', 22, '/ecls/lesson', 'view', 0, 'ecls/lesson/index', '', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (24, 'ecls:lesson:search', '课件资源搜索', 23, '/ecls/lesson/search', 'widget', 1, NULL, 'search', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (25, 'ecls:lesson:getOne', '课件资源详情', 23, 'get:/ecls/lesson/{pkId}', 'widget', 2, NULL, 'tickets', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (26, 'ecls:lesson:create', '课件资源新增', 23, 'post:/ecls/lesson', 'widget', 3, NULL, 'plus', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (27, 'ecls:lesson:update', '课件资源修改', 23, 'put:/ecls/lesson', 'widget', 4, NULL, 'edit', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (28, 'ecls:lesson:delete', '课件资源删除单条', 23, 'delete:/ecls/lesson/{pkId}', 'widget', 5, NULL, 'delete', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (29, 'ecls:lesson:export', '课件资源导出', 23, 'get,post:/ecls/lesson/export', 'widget', 7, NULL, 'export', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (30, 'ecls:lesson:meta', '课件资源元数据', 23, 'get:/ecls/lesson/meta', 'widget', 8, NULL, 'meta', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (31, 'ecls:record', '上课记录', 22, '/ecls/record', 'view', 0, 'ecls/record/index', '', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (32, 'ecls:record:search', '上课记录搜索', 31, '/ecls/record/search', 'widget', 1, NULL, 'search', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (33, 'ecls:record:getOne', '上课记录详情', 31, 'get:/ecls/record/{pkId}', 'widget', 2, NULL, 'tickets', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (34, 'ecls:record:create', '上课记录新增', 31, 'post:/ecls/record', 'widget', 3, NULL, 'plus', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (35, 'ecls:record:update', '上课记录修改', 31, 'put:/ecls/record', 'widget', 4, NULL, 'edit', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (36, 'ecls:record:delete', '上课记录删除单条', 31, 'delete:/ecls/record/{pkId}', 'widget', 5, NULL, 'delete', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (37, 'ecls:record:export', '上课记录导出', 31, 'get,post:/ecls/record/export', 'widget', 7, NULL, 'export', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (38, 'ecls:record:meta', '上课记录元数据', 31, 'get:/ecls/record/meta', 'widget', 8, NULL, 'meta', NULL, 0, '');
INSERT INTO `sys_rights` VALUES (39, 'system:config', '系統參數', 3, '/system/config', 'view', 0, 'sys/config/index', NULL, NULL, 0, NULL);
INSERT INTO `sys_rights` VALUES (45, 'system:acc', '用户管理', 3, '/sys/acc', 'view', 2, 'sys/acc/index', NULL, 0, 0, '');
INSERT INTO `sys_rights` VALUES (46, 'ecls:record:mine', '我的记录', 22, '/ecls/record/mine', 'view', 1, 'ecls/record/mine', NULL, 0, 0, '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `code` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '索引admin/patient/worker',
  `rights_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '权限列表',
  `right_ids` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '权限列表',
  `state` int(11) NULL DEFAULT 1 COMMENT '状态1可用,0不可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'admin', NULL, '[1,2,3,4,5,21,6,7,8,9,39,45,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,46]', 1);
INSERT INTO `sys_role` VALUES (2, '学生', 'member', NULL, '[1,2,3,8,22,46]', 1);
INSERT INTO `sys_role` VALUES (3, '注册租户', 'tenant', NULL, '[3,8,22,46]', 1);
INSERT INTO `sys_role` VALUES (4, '普通访客', 'guest', NULL, '[3]', 1);

-- ----------------------------
-- Table structure for sys_smstask
-- ----------------------------
DROP TABLE IF EXISTS `sys_smstask`;
CREATE TABLE `sys_smstask`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '验证码',
  `phone_numbers` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `templete_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '模板ID',
  `result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '发送结果',
  `create_at` datetime NULL DEFAULT NULL COMMENT '发布时间',
  `success` int(11) NULL DEFAULT 0 COMMENT '是否成功',
  `send_at` datetime NULL DEFAULT NULL COMMENT '发送时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_smstask
-- ----------------------------

-- ----------------------------
-- Table structure for sys_userinfo
-- ----------------------------
DROP TABLE IF EXISTS `sys_userinfo`;
CREATE TABLE `sys_userinfo`  (
  `user_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户ID',
  `username` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `pic` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '头像',
  `role_id` bigint(20) NULL DEFAULT NULL COMMENT '角色ID',
  `dept_id` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '所属部门',
  `status` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '1' COMMENT '是否可用',
  `create_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `delete_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `deleted` int(11) NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `idx_sys_userinfo_username`(`username`) USING BTREE,
  INDEX `idx_sys_userinfo_mobile`(`mobile`) USING BTREE,
  INDEX `idx_sys_userinfo_dept_id`(`dept_id`) USING BTREE,
  INDEX `fk_sys_userinfo_role`(`role_id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_userinfo
-- ----------------------------
INSERT INTO `sys_userinfo` VALUES ('8y9jasdi8uy8933uiuhisdhak', 'admin', '15367151352', '$2a$04$zq79Oq2Szs7RHgcByH3eDeJ9waA89IYh9mfqysJH6bg/XODacuaMq', '创世主', 'mnt/2025/05/28/202505282042311748434560.png', 1, NULL, '1', '2025-05-27 15:31:03', '2025-05-27 15:31:05', 0);
INSERT INTO `sys_userinfo` VALUES ('212', 'winlion', NULL, '$2a$04$zz50pAbZIwZy0LdDdrPdcuwQuwz5OlphNsk9yxfl9FXxEEpY3PDqG', '胡大力', NULL, 2, NULL, '1', NULL, NULL, 0);
INSERT INTO `sys_userinfo` VALUES ('720462452218793984', 'z0019', '', '$2a$04$t/oyVfavXJnH.HLwKc264uAF9OPzUGz.pC79VOJwBLR6uK81ll8xW', '张三19', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452210405376', 'z0018', '', '$2a$04$IDuNJO.TerjFcZKHAPBHJOygmcHbYkJMQNn.CpQJCWNPuXaR8.SNK', '张三18', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452197822464', 'z0017', '', '$2a$04$nYUiDZ8rgMJm4fwbof9DbOqMbFV0kaiZ.kZPz28NQbR28qyxl6XPy', '张三17', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452189433856', 'z0016', '', '$2a$04$wzXKCJ3dvL5KSLmtEOMLz.zFlZA66Rp1ggdiT5z.WQeAGc3RmygZi', '张三16', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452185239552', 'z0015', '', '$2a$04$cq/lQwq.OTQH83ocU8PTq.8aMaOJQisqKyPoA8LZD.c1NIMbgwXPG', '张三15', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452176850944', 'z0014', '', '$2a$04$YUvBOXkR71hzCS.jK.8ORO66J1d0uincTkoJi2eY/bG09FtQz6owS', '张三14', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452168462336', 'z0013', '', '$2a$04$vih.TzTdeH7vdkgc1rVZWexeNWmYDbMgdvwg8Y59GiBf78dqsn7Q.', '张三13', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452160073728', 'z0012', '', '$2a$04$7X1A0CknzbQU5pm0kamYpeuM/EjU0YEa.SO42ExTOxpu5uIzJ6drO', '张三12', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452151685120', 'z0011', '', '$2a$04$b61saiY7SjdwTVk2WZHWf.4URepElMZ2ktC4CjOpdIiG4zoLVOsG2', '张三11', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452147490816', 'z0010', '', '$2a$04$EixR9De3yXfx28LuAgCDKODn9e2YlRm6ioSOs1gEByAw87n/qzK0m', '张三10', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452134907904', 'z0009', '', '$2a$04$scVJgIEQvnWpDtKhAstAPOsRKCQ3A/8q7EfFRCA6EkpC5olsYF3Im', '张三09', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452126519296', 'z0008', '', '$2a$04$bVT8jpk7BQOIH9rxzBYMeeUI6MBnKAfQsZot5E1KMuceWVQst9BvS', '张三08', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452118130688', 'z0007', '', '$2a$04$gZ3fZq/sjhDWcEN.ChNmGeaLVBVYRMmCLAss7AL46qVfKMX/ANhTm', '张三07', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452113936384', 'z0006', '', '$2a$04$l44NODSDAaHl/wiyk1EZUeo5bzmfeVQVeL4W8wt2FsO2HSzTlDz/6', '张三06', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452105547776', 'z0005', '', '$2a$04$17UVNSxJuA.ehBbD3nF5Kujm4PpunEO4vv68BiSyPUj09uQ.l9xIW', '张三05', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452097159168', 'z0004', '', '$2a$04$Nnxu0ZYI28Ry9rpo/q6hmO1.2zH4NcHDFNBS7zm17Boz.K0e.KCpO', '张三04', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452092964864', 'z0003', '', '$2a$04$F0KwbvzaZDq.7s6bFHNkceemyUCRKnMpPpZM77ujBBjBvGsqw1/Te', '张三03', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452080381952', 'z0001', '', '$2a$04$7KguXA.NH1CqWDcSCccb7OHaoLsVotGlCvN6R4AY4M0.wP4sabQi.', '张三01', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452084576256', 'z0002', '', '$2a$04$zBxj3ZH46.AIVWdF86vrHu69yzXoBpxylEa3uA8McR80cNGmNyF4a', '张三02', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452227182592', 'z0020', '', '$2a$04$qQFQvQG3JuH4SUmOhTYF1eomH/29nPteR.Awtrdm2RazgjKpShCai', '张三20', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('720462452231376896', 'z0021', '', '$2a$04$QpnRvYkzgyeQ9i.EgwKSHuhXWrWdDY2A0eePOmyOsyIfxzud.w.Xy', '张三21', '', 2, '', '1', '2025-06-11 10:20:34', '2025-06-11 10:20:34', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980272365568', 'h0003', '', '$2a$04$AhuFl22F/laPktcgibsOsuWTZnYYvYqpHs.EFRe5iV6l9wVcnVtum', '张三01', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980276559872', 'h0004', '', '$2a$04$s/MmvsHsFg6W7ri0vipxouQrQ8BGbAz6Jf1mK8nrPBBj4SMWxIHT.', '张三02', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980280754176', 'h0005', '', '$2a$04$PpFzyedoDDK4SM10vHXQ/.AVy7JctVnhvfcz1yoD.D6DnjGra0Jd.', '张三03', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980289142784', 'h0006', '', '$2a$04$WZ7JU9To9AGtC2kM6L4PFu18JcofgGwhRMPzf5OFTVnLempMFc512', '张三04', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980297531392', 'h0007', '', '$2a$04$qnOPUfDO7PxbRIfTIlBo9uIGfPtpaKJjqOCb7zVob8JlXCfiIIKCC', '张三05', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980305920000', 'h0008', '', '$2a$04$VuCnBGPzpfvz0uYIDYAbjOoeyk9X7RHpqBD2C30jpGIGBOjuAfumq', '张三06', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980305920001', 'h0009', '', '$2a$04$yIumJxhGspaBU1UKEWpdreWIvKd8fQ6a2yyLQH5xBHkgUt36oatNC', '张三07', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980318502912', 'h0010', '', '$2a$04$V0AL0hOv055ksPyqfH9LlOmDGq4TAVtZ7rj2oLNjeOe8mf/xp8io.', '张三08', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980322697216', 'h0011', '', '$2a$04$c.Bawypq4.rQjb8QTyUBtOetE5lQ2l0Jf.7YdY1Do47Ww1BVdswq2', '张三09', '', 2, '', '1', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980326891520', 'h0012', '', '$2a$04$ZK.vFkvuyqpfrW/ni1K6dObFocl7NC1OopaLRAI.sZDObHux8bhL6', '张三10', '', 2, '', '255', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980335280128', 'h0013', '', '$2a$04$oauPQbrlkpiTIUlyX.7cfe03cXIgz7bYcOfTsx94xxCDfeZWQTmoa', '张三11', '', 2, '', '255', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980343668736', 'h0014', '', '$2a$04$slkKz7VVU1pvsDZht7zD9eJrTS8bh7kd/2nA/fAVTNgPpUj0ngdYq', '张三12', '', 2, '', '255', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);
INSERT INTO `sys_userinfo` VALUES ('721533980352057344', 'h0015', '', '$2a$04$CEenQ/6ka8WnDj2ajw2lr.S87IZ8XTrWBEZfk8hYI38ANIEpt6APW', '张三13', '', 2, '', '255', '2025-06-14 09:18:26', '2025-06-14 09:18:26', 0);

-- ----------------------------
-- View structure for record_view
-- ----------------------------
DROP VIEW IF EXISTS `record_view`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `record_view` AS select `a`.`id` AS `id`,`a`.`create_at` AS `create_at`,`a`.`user_id` AS `user_id`,`a`.`lesson_id` AS `lesson_id`,`a`.`create_date` AS `create_date`,`b`.`title` AS `title`,`b`.`memo` AS `memo`,`b`.`cate` AS `cate`,`b`.`media` AS `media`,`b`.`cover` AS `cover`,`b`.`uri` AS `uri` from (`ecls_record` `a` left join `ecls_lesson` `b` on((`a`.`lesson_id` = `b`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
