/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50712
Source Host           : 127.0.0.1:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50712
File Encoding         : 65001

Date: 2018-07-20 14:32:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for area
-- ----------------------------
DROP TABLE IF EXISTS `area`;
CREATE TABLE `area` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Aid` bigint(20) NOT NULL DEFAULT '0' COMMENT '区域ID',
  `Pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '区域父ID',
  `Level` int(255) NOT NULL DEFAULT '0' COMMENT '层级 0:国家 1：省  2：市 3：区/县  4:镇  5：村委会',
  `Name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `SimName` varchar(255) NOT NULL DEFAULT '' COMMENT '名字缩写',
  `NamePath` varchar(255) NOT NULL DEFAULT '' COMMENT '全名路径',
  `SimNamePath` varchar(255) NOT NULL DEFAULT '' COMMENT '缩写全名路径',
  `TelCode` varchar(255) NOT NULL DEFAULT '' COMMENT '电话区号',
  `ZipCode` varchar(255) NOT NULL DEFAULT '' COMMENT '邮政编码',
  `Pinyin` varchar(255) NOT NULL DEFAULT '' COMMENT '拼音',
  `SimPinyin` varchar(255) NOT NULL DEFAULT '' COMMENT '拼音缩写',
  `FirstPinyin` varchar(255) NOT NULL DEFAULT '' COMMENT '拼音首字母',
  `Lng` decimal(8,8) NOT NULL DEFAULT '0.00000000' COMMENT '经度',
  `Lat` decimal(8,8) NOT NULL DEFAULT '0.00000000' COMMENT '纬度',
  `Status` int(255) NOT NULL DEFAULT '0' COMMENT '有效无效',
  `Url` varchar(255) NOT NULL DEFAULT '' COMMENT '国家地区URL',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
