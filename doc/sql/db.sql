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

SET FOREIGN_KEY_CHECKS = 1;
