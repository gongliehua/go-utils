/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80012 (8.0.12)
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80012 (8.0.12)
 File Encoding         : 65001

 Date: 19/08/2025 11:34:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_roles
-- ----------------------------
DROP TABLE IF EXISTS `admin_roles`;
CREATE TABLE `admin_roles`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '管理员id',
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色id',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员关联角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_roles
-- ----------------------------
INSERT INTO `admin_roles` VALUES (6, 1, 1, '2025-08-13 17:17:41', '2025-08-13 17:17:41');
INSERT INTO `admin_roles` VALUES (9, 36, 1, '2025-08-18 17:43:57', '2025-08-18 17:43:57');
INSERT INTO `admin_roles` VALUES (10, 36, 3, '2025-08-18 17:43:57', '2025-08-18 17:43:57');

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `sex` tinyint(3) NOT NULL DEFAULT 0 COMMENT '性别:0=未知,1=男,2=女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `status` tinyint(3) NOT NULL DEFAULT 1 COMMENT '状态:0=禁用,1=启用',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admins
-- ----------------------------
INSERT INTO `admins` VALUES (1, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 'Admin', 2, '/uploads/2025/07/18/1752828347512521.jpg', 1, '2025-07-15 17:51:48', '2025-08-15 15:45:47', NULL);
INSERT INTO `admins` VALUES (36, 'gongliehua', 'e10adc3949ba59abbe56e057f20f883e', '龚烈华', 2, '', 1, '2025-08-13 17:16:12', '2025-08-18 17:43:57', NULL);
INSERT INTO `admins` VALUES (37, '11sa', 'e10adc3949ba59abbe56e057f20f883e', 's', 1, '', 1, '2025-08-19 10:52:59', '2025-08-19 10:52:59', NULL);
INSERT INTO `admins` VALUES (38, 'aaad', 'e10adc3949ba59abbe56e057f20f883e', 'fda', 1, '', 1, '2025-08-19 10:53:50', '2025-08-19 10:53:50', '2025-08-19 10:54:04');

-- ----------------------------
-- Table structure for configs
-- ----------------------------
DROP TABLE IF EXISTS `configs`;
CREATE TABLE `configs`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `var` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '变量名',
  `type` tinyint(3) NOT NULL DEFAULT 0 COMMENT '类型:0=单行文本,1=多行文本,2=单选按钮,3=复选框,4=下拉框',
  `option` json NULL COMMENT '配置项(针对于:单选按钮,复选框,下拉框)',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '配置值',
  `weight` bigint(20) NOT NULL DEFAULT 100 COMMENT '排序(权重)',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `var`(`var` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of configs
-- ----------------------------
INSERT INTO `configs` VALUES (6, '标题', 'title', 0, '[]', '3', 100, '2025-08-15 14:34:08', '2025-08-15 15:46:32', NULL);
INSERT INTO `configs` VALUES (7, '关键字', 'keywords', 1, '[]', '3', 100, '2025-08-15 14:34:25', '2025-08-15 15:46:32', NULL);
INSERT INTO `configs` VALUES (8, '描述', 'description', 1, '[]', '3', 100, '2025-08-15 14:34:50', '2025-08-15 15:46:32', NULL);
INSERT INTO `configs` VALUES (9, '维护模式', 'switch', 2, '[{\"key\": \"up\", \"value\": \"上线\"}, {\"key\": \"down\", \"value\": \"下线\"}]', 'up', 100, '2025-08-15 14:37:46', '2025-08-15 15:46:32', NULL);
INSERT INTO `configs` VALUES (10, '可选主题', 'themes', 3, '[{\"key\": \"a\", \"value\": \"白色\"}, {\"key\": \"b\", \"value\": \"黄色\"}, {\"key\": \"c\", \"value\": \"绿色\"}, {\"key\": \"d\", \"value\": \"黑色\"}]', '[\"a\",\"b\",\"c\",\"d\"]', 100, '2025-08-15 14:39:57', '2025-08-15 15:46:32', NULL);
INSERT INTO `configs` VALUES (11, '性别', 'sex', 4, '[{\"key\": 0, \"value\": \"未知\"}, {\"key\": 1, \"value\": \"男\"}, {\"key\": \"2\", \"value\": \"女\"}]', '2', 100, '2025-08-15 14:41:02', '2025-08-15 15:46:32', NULL);

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级权限',
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `icon` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'URL',
  `is_menu` tinyint(3) NOT NULL DEFAULT 0 COMMENT '菜单:0=否,1=是',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `weight` bigint(20) NOT NULL DEFAULT 100 COMMENT '排序(权重)',
  `status` tinyint(3) NOT NULL DEFAULT 1 COMMENT '状态:0=禁用,1=启用',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of permissions
-- ----------------------------
INSERT INTO `permissions` VALUES (1, 0, '控制面板', 'fa-tachometer', '/admin', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (2, 0, '管理员管理', 'fa-users', '', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (3, 2, '管理员列表', '', '/admin/admin', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (4, 2, '角色管理', '', '/admin/role', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (5, 2, '权限管理', '', '/admin/permission', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (6, 0, '系统设置', 'fa-cogs', '', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (7, 6, '配置列表', '', '/admin/config', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);
INSERT INTO `permissions` VALUES (8, 6, '配置管理', '', '/admin/config/setting', 1, '', 100, 1, '2025-08-19 00:00:00', '2025-08-19 00:00:00', NULL);

-- ----------------------------
-- Table structure for role_permissions
-- ----------------------------
DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色id',
  `permission_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '权限id',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色关联权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_permissions
-- ----------------------------

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint(3) NOT NULL DEFAULT 0 COMMENT '状态:0=禁用,1=启用',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (1, '默认角色', '', 1, '2025-08-13 17:04:19', '2025-08-18 17:44:04');
INSERT INTO `roles` VALUES (3, '普通用户', '', 1, '2025-08-13 17:15:49', '2025-08-18 18:03:54');

SET FOREIGN_KEY_CHECKS = 1;
