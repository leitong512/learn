
create table `conf_watertrees_act` (
    `id` int(11) not null auto_increment,
    `activity_name` varchar(30) collate utf8mb4_unicode_ci default null comment '活动名',
    `begin_date` datetime DEFAULT NULL COMMENT '开始时间',
    `end_date` datetime DEFAULT NULL COMMENT '结束时间',
    `activity_gift` varchar(30) collate utf8mb4_unicode_ci default null comment '活动礼物',
    `gift_id` int(11) default null comment '礼物ID',
    `time_interval` smallint(4) DEFAULT NULL COMMENT '滚屏结束时间(s)',
    `rolling_content` text COLLATE utf8mb4_unicode_ci COMMENT '滚屏文案',
    `end_rolling_content` text COLLATE utf8mb4_unicode_ci COMMENT '结束滚屏文案',
    `end_sys_msg` text COLLATE utf8mb4_unicode_ci COMMENT '结束系统消息文案',
    `status` smallint(5) DEFAULT NULL COMMENT '状态 1-进行中 2-已结束',
    PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



create table `watertrees_act_user` (
    `id` int(11) not null auto_increment,
    `activity_id`   int(11) not null,
    `activity_name` varchar(30) collate utf8mb4_unicode_ci default null comment '活动名',
    `user_id`   int(11) comment '用户ID',
    `user_name` varchar(30) default '' comment '用户名',
    `gift_count` int(11) default '0' comment '礼物数量',
    `begin_date` datetime default null comment '活动开始时间',
    `end_date` datetime DEFAULT NULL COMMENT '结束时间',
    `update_date` datetime default null comment '更新数量时间',
    `status` smallint(5) DEFAULT NULL COMMENT '状态 1-进行中 2-已结束',
    primary key ('id')
) engine=innod default  charset=utf8mb4 collate=utf8mb4_unicode_ci;

create table `log_manager_broadcast_room` (
    `id` int(11) not null auto_increment,
    `room_id` int(11) not null comment '房间ID',
    `pretty_num` int(11) default '0' comment '靓号',
    `room_name` varchar (50) default null comment '房间名称',
    `nickname` varchar (50) default null comment '房主昵称',
    `create_at` datetime default null comment  '创建时间',
    primary key (`id`),
    unique key `room_id` (`room_id`)
) engine=innod default  charset=utf8mb4 collate=utf8mb4_unicode_ci;