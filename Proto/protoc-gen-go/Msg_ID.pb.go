// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Msg_ID.proto

/*
Package Msg_ID is a generated protocol buffer package.

It is generated from these files:
	Msg_ID.proto

It has these top-level messages:
*/
package Msg_ID

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MessageID int32

const (
	MessageID_MSG_BEGIN MessageID = 0
	// ============================================================================================
	// 				一般消息定义
	// ============================================================================================
	MessageID_MSG_NORMAL_MSGID_BEGIN     MessageID = 1000000
	MessageID_MSG_CHECK_VERSION_REQ      MessageID = 1000001
	MessageID_MSG_CHECK_VERSION_ACK      MessageID = 1000002
	MessageID_MSG_ACCOUNT_REG_REQ        MessageID = 1000003
	MessageID_MSG_ACCOUNT_REG_ACK        MessageID = 1000004
	MessageID_MSG_ACCOUNT_LOGIN_REQ      MessageID = 1000005
	MessageID_MSG_ACCOUNT_LOGIN_ACK      MessageID = 1000006
	MessageID_MSG_DB_EXE_SQL_REQ         MessageID = 1000007
	MessageID_MSG_DB_EXE_SQL_ACK         MessageID = 1000008
	MessageID_MSG_SERVER_LIST_REQ        MessageID = 1000009
	MessageID_MSG_SERVER_LIST_ACK        MessageID = 1000010
	MessageID_MSG_SELECT_SERVER_REQ      MessageID = 1000011
	MessageID_MSG_SELECT_SERVER_ACK      MessageID = 1000012
	MessageID_MSG_LOGIC_REGTO_LOGIN_REQ  MessageID = 1000013
	MessageID_MSG_LOGIC_REGTO_LOGIN_ACK  MessageID = 1000014
	MessageID_MSG_SEAL_ACCOUNT_REQ       MessageID = 1000015
	MessageID_MSG_SEAL_ACCOUNT_ACK       MessageID = 1000016
	MessageID_MSG_GAME_REGTO_LOGIC_REQ   MessageID = 1000017
	MessageID_MSG_GAME_REGTO_LOGIC_ACK   MessageID = 1000018
	MessageID_MSG_LOGIC_REGTO_CENTER_REQ MessageID = 1000019
	MessageID_MSG_LOGIC_REGTO_CENTER_ACK MessageID = 1000020
	MessageID_MSG_LOG_DATA_NTF           MessageID = 1000021
	MessageID_MSG_PHP_GM_COMMAND_REQ     MessageID = 1000022
	MessageID_MSG_PHP_GM_COMMAND_ACK     MessageID = 1000023
	MessageID_MSG_WATCH_HEART_BEAT_REQ   MessageID = 1000024
	MessageID_MSG_WATCH_HEART_BEAT_ACK   MessageID = 1000025
	MessageID_MSG_WATCH_UPDATE_SVR_REQ   MessageID = 1000026
	MessageID_MSG_WATCH_UPDATE_SVR_ACK   MessageID = 1000027
	MessageID_MSG_WATCH_START_SVR_REQ    MessageID = 1000028
	MessageID_MSG_WATCH_START_SVR_ACK    MessageID = 1000029
	MessageID_MSG_WATCH_STOP_SVR_REQ     MessageID = 1000030
	MessageID_MSG_WATCH_STOP_SVR_ACK     MessageID = 1000031
	MessageID_MSG_DISCONNECT_NTY         MessageID = 1000032
	MessageID_MSG_ROLE_DATA_CHANGE_NTY   MessageID = 1000033
	MessageID_MSG_NORMAL_MSGID_END       MessageID = 1999999
	// ============================================================================================
	// 				Cient <==> LogicSVR
	// ============================================================================================
	MessageID_MSG_LOGICSVR_MSGID_BEGIN  MessageID = 2000000
	MessageID_MSG_ROLE_LIST_REQ         MessageID = 2000001
	MessageID_MSG_ROLE_LIST_ACK         MessageID = 2000002
	MessageID_MSG_ROLE_CREATE_REQ       MessageID = 2000003
	MessageID_MSG_ROLE_CREATE_ACK       MessageID = 2000004
	MessageID_MSG_ROLE_DELETE_REQ       MessageID = 2000005
	MessageID_MSG_ROLE_DELETE_ACK       MessageID = 2000006
	MessageID_MSG_ROLE_LOGIN_REQ        MessageID = 2000007
	MessageID_MSG_ROLE_LOGIN_ACK        MessageID = 2000008
	MessageID_MSG_ROLE_LOGOUT_REQ       MessageID = 2000009
	MessageID_MSG_ROLE_LOGOUT_ACK       MessageID = 2000010
	MessageID_MSG_DATA_UPDATE           MessageID = 2000011
	MessageID_MSG_NOTIFY_INTO_SCENE     MessageID = 2000014
	MessageID_MSG_COPYINFO_REPORT_REQ   MessageID = 2000016
	MessageID_MSG_MAIN_COPY_REQ         MessageID = 2000017
	MessageID_MSG_MAIN_COPY_ACK         MessageID = 2000018
	MessageID_MSG_COPY_ABORT_REQ        MessageID = 2000019
	MessageID_MSG_COPY_ABORT_ACK        MessageID = 2000020
	MessageID_MSG_BACK_TO_CITY_REQ      MessageID = 2000021
	MessageID_MSG_BACK_TO_CITY_ACK      MessageID = 2000022
	MessageID_MSG_BATTLE_RESULT_NTY     MessageID = 2000023
	MessageID_MSG_MAINCOPY_RESULT_NTY   MessageID = 2000024
	MessageID_MSG_ROLE_CHANG_NTY        MessageID = 2000025
	MessageID_MSG_BAG_CHANGE_NTY        MessageID = 2000026
	MessageID_MSG_PET_CHANGE_NTY        MessageID = 2000027
	MessageID_MSG_EQUIP_CHANGE_NTY      MessageID = 2000028
	MessageID_MSG_PARTNER_CHANGE_NTY    MessageID = 2000029
	MessageID_MSG_MOUNT_CHANGE_NTY      MessageID = 2000030
	MessageID_MSG_ROLE_OTHER_LOGIN_NTY  MessageID = 2000031
	MessageID_MSG_CHAT_MESSAGE_REQ      MessageID = 2000032
	MessageID_MSG_CHAT_MESSAGE_ACK      MessageID = 2000033
	MessageID_MSG_CHAT_MESSAGE_NTY      MessageID = 2000034
	MessageID_MSG_SETUP_EQUIP_REQ       MessageID = 2000035
	MessageID_MSG_SETUP_EQUIP_ACK       MessageID = 2000036
	MessageID_MSG_UNSET_EQUIP_REQ       MessageID = 2000037
	MessageID_MSG_UNSET_EQUIP_ACK       MessageID = 2000038
	MessageID_MSG_SETUP_PET_REQ         MessageID = 2000039
	MessageID_MSG_SETUP_PET_ACK         MessageID = 2000040
	MessageID_MSG_SETUP_PARTNER_REQ     MessageID = 2000041
	MessageID_MSG_SETUP_PARTNER_ACK     MessageID = 2000042
	MessageID_MSG_UNSET_PARTNER_REQ     MessageID = 2000043
	MessageID_MSG_UNSET_PARTNER_ACK     MessageID = 2000044
	MessageID_MSG_SETUP_MOUNT_REQ       MessageID = 2000045
	MessageID_MSG_SETUP_MOUNT_ACK       MessageID = 2000046
	MessageID_MSG_UNSET_MOUNT_REQ       MessageID = 2000047
	MessageID_MSG_UNSET_MOUNT_ACK       MessageID = 2000048
	MessageID_MSG_USE_ITEM_REQ          MessageID = 2000049
	MessageID_MSG_USE_ITEM_ACK          MessageID = 2000050
	MessageID_MSG_SETUP_GEM_REQ         MessageID = 2000051
	MessageID_MSG_SETUP_GEM_ACK         MessageID = 2000052
	MessageID_MSG_UNSET_GEM_REQ         MessageID = 2000053
	MessageID_MSG_UNSET_GEM_ACK         MessageID = 2000054
	MessageID_MSG_ROLE_RECONNECT_REQ    MessageID = 2000055
	MessageID_MSG_ROLE_RECONNECT_ACK    MessageID = 2000056
	MessageID_MSG_MAIL_CHANGE_NTY       MessageID = 2000057
	MessageID_MSG_TASK_CHANGE_NTY       MessageID = 2000058
	MessageID_MSG_BROAD_MESSAGE_NOTIFY  MessageID = 2000060
	MessageID_MSG_GASVR_REGTO_PROXY_REQ MessageID = 2000061
	MessageID_MSG_GASVR_REGTO_PROXY_ACK MessageID = 2000062
	MessageID_MSG_STORE_BUY_REQ         MessageID = 2000063
	MessageID_MSG_STORE_BUY_ACK         MessageID = 2000064
	MessageID_MSG_GEM_CHANGE_NTY        MessageID = 2000065
	MessageID_MSG_TEST_ADD_ITEM         MessageID = 2999998
	MessageID_MSG_LOGICSVR_MSGID_END    MessageID = 2999999
	// ============================================================================================
	// 				Cient <==> SceneSVR
	// ============================================================================================
	MessageID_MSG_SCENESVR_MSGID_BEGIN  MessageID = 3000000
	MessageID_MSG_CREATE_SCENE_REQ      MessageID = 3000001
	MessageID_MSG_CREATE_SCENE_ACK      MessageID = 3000002
	MessageID_MSG_DELETE_SCENE_REQ      MessageID = 3000003
	MessageID_MSG_DELETE_SCENE_ACK      MessageID = 3000004
	MessageID_MSG_SCENE_DESTROY_REQ     MessageID = 3000005
	MessageID_MSG_SCENE_DESTROY_ACK     MessageID = 3000006
	MessageID_MSG_TRANSFER_DATA_REQ     MessageID = 3000007
	MessageID_MSG_TRANSFER_DATA_ACK     MessageID = 3000008
	MessageID_MSG_NOTIFY_ROLE_ENTER_REQ MessageID = 3000009
	MessageID_MSG_NOTIFY_ROLE_ENTER_ACK MessageID = 3000010
	MessageID_MSG_ENTER_SCENE_REQ       MessageID = 3000011
	MessageID_MSG_ENTER_SCENE_ACK       MessageID = 3000012
	MessageID_MSG_LEAVE_SCENE_REQ       MessageID = 3000013
	MessageID_MSG_LEAVE_SCENE_ACK       MessageID = 3000014
	MessageID_MSG_SKILL_CAST_REQ        MessageID = 3000015
	MessageID_MSG_SKILL_CAST_ACK        MessageID = 3000016
	MessageID_MSG_SKILL_CAST_NTF        MessageID = 3000017
	MessageID_MSG_SKILL_RESULT_NTF      MessageID = 3000018
	MessageID_MSG_OBJECT_NEW_NTF        MessageID = 3000019
	MessageID_MSG_OBJECT_CHANGE_NTF     MessageID = 3000020
	MessageID_MSG_OBJECT_REMOVE_NTF     MessageID = 3000021
	MessageID_MSG_OBJECT_ACTION_REQ     MessageID = 3000022
	MessageID_MSG_OBJECT_ACTION_ACK     MessageID = 3000023
	MessageID_MSG_HEART_BEAT_REQ        MessageID = 3000024
	MessageID_MSG_HEART_BEAT_ACK        MessageID = 3000025
	MessageID_MSG_USE_HP_BOOTTLE_REQ    MessageID = 3000026
	MessageID_MSG_USE_HP_BOOTTLE_ACK    MessageID = 3000027
	MessageID_MSG_USE_MP_BOOTTLE_REQ    MessageID = 3000028
	MessageID_MSG_USE_MP_BOOTTLE_ACK    MessageID = 3000029
	MessageID_MSG_OBJECT_DIE_NOTIFY     MessageID = 3000030
	MessageID_MSG_BATTLE_CHAT_REQ       MessageID = 3000031
	MessageID_MSG_BATTLE_CHAT_ACK       MessageID = 3000032
	MessageID_MSG_BULLET_NEW_NTF        MessageID = 3000033
	MessageID_MSG_SWITCH_MOUNT_REQ      MessageID = 3000034
	MessageID_MSG_SWITCH_MOUNT_ACK      MessageID = 3000035
	MessageID_MSG_SCENEOBJ_CHAGE_NTF    MessageID = 3000036
	MessageID_MSG_SCENESVR_MSGID_END    MessageID = 3999999
	MessageID_MSG_REQ_USEITEM           MessageID = 1103
	MessageID_MSG_ACK_USEITEM           MessageID = 1104
	MessageID_MSG_REQ_DRESS_GEM         MessageID = 1109
	MessageID_MSG_ACK_DRESS_GEM         MessageID = 1110
	MessageID_MSG_REQ_UNLOAD_GEM        MessageID = 1111
	MessageID_MSG_ACK_UNLOAD_GEM        MessageID = 1112
	MessageID_MSG_REQ_COMPOSE_CHIP      MessageID = 1113
	MessageID_MSG_ACK_COMPOSE_CHIP      MessageID = 1114
	MessageID_MSG_REQ_STRENGTHEN_EQUIP  MessageID = 1115
	MessageID_MSG_ACK_STRENGTHEN_EQUIP  MessageID = 1116
	MessageID_MSG_REQ_ADVANCE_EQUIP     MessageID = 1117
	MessageID_MSG_ACK_ADVANCE_EQUIP     MessageID = 1118
	MessageID_MSG_REQ_UPSTAR_EQUIP      MessageID = 1119
	MessageID_MSG_ACK_UPSTAR_EQUIP      MessageID = 1120
	MessageID_MSG_REQ_STRENGTHEN_GEM    MessageID = 1121
	MessageID_MSG_ACK_STRENGTHEN_GEM    MessageID = 1122
	MessageID_MSG_REQ_ONEKEYTODRESSGEM  MessageID = 1151
	MessageID_MSG_ACK_ONEKEYTODRESSGEM  MessageID = 1152
	MessageID_MSG_REQ_ONEKEYTOUNLOADGEM MessageID = 1153
	MessageID_MSG_ACK_ONEKEYTOUNLOADGEM MessageID = 1154
	MessageID_MSG_REQ_GET_CHAPTERAWARD  MessageID = 1201
	MessageID_MSG_ACK_GET_CHAPTERAWARD  MessageID = 1202
	MessageID_MSG_REQ_BATTLE_CHECK      MessageID = 1203
	MessageID_MSG_ACK_BATTLE_CHECK      MessageID = 1204
	MessageID_MSG_REQ_PASSCOPY          MessageID = 1205
	MessageID_MSG_ACK_PASSCOPY          MessageID = 1206
	MessageID_MSG_REQ_SETMOUNT          MessageID = 1301
	MessageID_MSG_ACK_SETMOUNT          MessageID = 1302
	MessageID_MSG_REQ_CHARGE_RELICS     MessageID = 1401
	MessageID_MSG_ACK_CHARGE_RELICS     MessageID = 1402
	MessageID_MSG_REQ_UPGRADE_RELICS    MessageID = 1403
	MessageID_MSG_ACK_UPGRADE_RELICS    MessageID = 1404
	MessageID_MSG_REQ_BATTLE_RELICS     MessageID = 1405
	MessageID_MSG_ACK_BATTLE_RELICS     MessageID = 1406
	MessageID_MSG_REQ_UNLOAD_RELICS     MessageID = 1407
	MessageID_MSG_ACK_UNLOAD_RELICS     MessageID = 1408
	MessageID_MSG_REQ_BUY_STORE         MessageID = 1501
	MessageID_MSG_ACK_BUY_STORE         MessageID = 1502
	MessageID_MSG_REQ_UPGRADE_PET       MessageID = 1601
	MessageID_MSG_ACK_UPGRADE_PET       MessageID = 1602
	MessageID_MSG_REQ_BATTLE_PET        MessageID = 1603
	MessageID_MSG_ACK_BATTLE_PET        MessageID = 1604
	MessageID_MSG_REQ_UNLOAD_PET        MessageID = 1605
	MessageID_MSG_ACK_UNLOAD_PET        MessageID = 1606
	MessageID_MSG_REQ_CHANGE_PARTNER    MessageID = 1607
	MessageID_MSG_ACK_CHANGE_PARTNER    MessageID = 1608
	MessageID_MSG_REQ_ADVANVE_PARTNER   MessageID = 1609
	MessageID_MSG_ACK_ADVANVE_PARTNER   MessageID = 1610
	MessageID_MSG_REQ_UPGRADE_PARTNER   MessageID = 1611
	MessageID_MSG_ACK_UPGRADE_PARTNER   MessageID = 1612
	MessageID_MSG_REQ_SUBMIT_TASK       MessageID = 1613
	MessageID_MSG_ACK_SUBMIT_TASK       MessageID = 1614
	MessageID_MSG_REQ_ADDHERO_EXP       MessageID = 2001
	MessageID_MSG_ACK_ADDHERO_EXP       MessageID = 2002
	MessageID_MSG_ENTER_VIEW_NTF        MessageID = 3001
	MessageID_MSG_LEAVE_VIEW_NTF        MessageID = 3002
	MessageID_MSG_ACTOR_VALUECHANGE_NTF MessageID = 3003
	MessageID_MSG_ACTOR_SKILLCAST_NTF   MessageID = 3004
	MessageID_MSG_ACTOR_DIE_NTF         MessageID = 3005
)

var MessageID_name = map[int32]string{
	0:       "MSG_BEGIN",
	1000000: "MSG_NORMAL_MSGID_BEGIN",
	1000001: "MSG_CHECK_VERSION_REQ",
	1000002: "MSG_CHECK_VERSION_ACK",
	1000003: "MSG_ACCOUNT_REG_REQ",
	1000004: "MSG_ACCOUNT_REG_ACK",
	1000005: "MSG_ACCOUNT_LOGIN_REQ",
	1000006: "MSG_ACCOUNT_LOGIN_ACK",
	1000007: "MSG_DB_EXE_SQL_REQ",
	1000008: "MSG_DB_EXE_SQL_ACK",
	1000009: "MSG_SERVER_LIST_REQ",
	1000010: "MSG_SERVER_LIST_ACK",
	1000011: "MSG_SELECT_SERVER_REQ",
	1000012: "MSG_SELECT_SERVER_ACK",
	1000013: "MSG_LOGIC_REGTO_LOGIN_REQ",
	1000014: "MSG_LOGIC_REGTO_LOGIN_ACK",
	1000015: "MSG_SEAL_ACCOUNT_REQ",
	1000016: "MSG_SEAL_ACCOUNT_ACK",
	1000017: "MSG_GAME_REGTO_LOGIC_REQ",
	1000018: "MSG_GAME_REGTO_LOGIC_ACK",
	1000019: "MSG_LOGIC_REGTO_CENTER_REQ",
	1000020: "MSG_LOGIC_REGTO_CENTER_ACK",
	1000021: "MSG_LOG_DATA_NTF",
	1000022: "MSG_PHP_GM_COMMAND_REQ",
	1000023: "MSG_PHP_GM_COMMAND_ACK",
	1000024: "MSG_WATCH_HEART_BEAT_REQ",
	1000025: "MSG_WATCH_HEART_BEAT_ACK",
	1000026: "MSG_WATCH_UPDATE_SVR_REQ",
	1000027: "MSG_WATCH_UPDATE_SVR_ACK",
	1000028: "MSG_WATCH_START_SVR_REQ",
	1000029: "MSG_WATCH_START_SVR_ACK",
	1000030: "MSG_WATCH_STOP_SVR_REQ",
	1000031: "MSG_WATCH_STOP_SVR_ACK",
	1000032: "MSG_DISCONNECT_NTY",
	1000033: "MSG_ROLE_DATA_CHANGE_NTY",
	1999999: "MSG_NORMAL_MSGID_END",
	2000000: "MSG_LOGICSVR_MSGID_BEGIN",
	2000001: "MSG_ROLE_LIST_REQ",
	2000002: "MSG_ROLE_LIST_ACK",
	2000003: "MSG_ROLE_CREATE_REQ",
	2000004: "MSG_ROLE_CREATE_ACK",
	2000005: "MSG_ROLE_DELETE_REQ",
	2000006: "MSG_ROLE_DELETE_ACK",
	2000007: "MSG_ROLE_LOGIN_REQ",
	2000008: "MSG_ROLE_LOGIN_ACK",
	2000009: "MSG_ROLE_LOGOUT_REQ",
	2000010: "MSG_ROLE_LOGOUT_ACK",
	2000011: "MSG_DATA_UPDATE",
	2000014: "MSG_NOTIFY_INTO_SCENE",
	2000016: "MSG_COPYINFO_REPORT_REQ",
	2000017: "MSG_MAIN_COPY_REQ",
	2000018: "MSG_MAIN_COPY_ACK",
	2000019: "MSG_COPY_ABORT_REQ",
	2000020: "MSG_COPY_ABORT_ACK",
	2000021: "MSG_BACK_TO_CITY_REQ",
	2000022: "MSG_BACK_TO_CITY_ACK",
	2000023: "MSG_BATTLE_RESULT_NTY",
	2000024: "MSG_MAINCOPY_RESULT_NTY",
	2000025: "MSG_ROLE_CHANG_NTY",
	2000026: "MSG_BAG_CHANGE_NTY",
	2000027: "MSG_PET_CHANGE_NTY",
	2000028: "MSG_EQUIP_CHANGE_NTY",
	2000029: "MSG_PARTNER_CHANGE_NTY",
	2000030: "MSG_MOUNT_CHANGE_NTY",
	2000031: "MSG_ROLE_OTHER_LOGIN_NTY",
	2000032: "MSG_CHAT_MESSAGE_REQ",
	2000033: "MSG_CHAT_MESSAGE_ACK",
	2000034: "MSG_CHAT_MESSAGE_NTY",
	2000035: "MSG_SETUP_EQUIP_REQ",
	2000036: "MSG_SETUP_EQUIP_ACK",
	2000037: "MSG_UNSET_EQUIP_REQ",
	2000038: "MSG_UNSET_EQUIP_ACK",
	2000039: "MSG_SETUP_PET_REQ",
	2000040: "MSG_SETUP_PET_ACK",
	2000041: "MSG_SETUP_PARTNER_REQ",
	2000042: "MSG_SETUP_PARTNER_ACK",
	2000043: "MSG_UNSET_PARTNER_REQ",
	2000044: "MSG_UNSET_PARTNER_ACK",
	2000045: "MSG_SETUP_MOUNT_REQ",
	2000046: "MSG_SETUP_MOUNT_ACK",
	2000047: "MSG_UNSET_MOUNT_REQ",
	2000048: "MSG_UNSET_MOUNT_ACK",
	2000049: "MSG_USE_ITEM_REQ",
	2000050: "MSG_USE_ITEM_ACK",
	2000051: "MSG_SETUP_GEM_REQ",
	2000052: "MSG_SETUP_GEM_ACK",
	2000053: "MSG_UNSET_GEM_REQ",
	2000054: "MSG_UNSET_GEM_ACK",
	2000055: "MSG_ROLE_RECONNECT_REQ",
	2000056: "MSG_ROLE_RECONNECT_ACK",
	2000057: "MSG_MAIL_CHANGE_NTY",
	2000058: "MSG_TASK_CHANGE_NTY",
	2000060: "MSG_BROAD_MESSAGE_NOTIFY",
	2000061: "MSG_GASVR_REGTO_PROXY_REQ",
	2000062: "MSG_GASVR_REGTO_PROXY_ACK",
	2000063: "MSG_STORE_BUY_REQ",
	2000064: "MSG_STORE_BUY_ACK",
	2000065: "MSG_GEM_CHANGE_NTY",
	2999998: "MSG_TEST_ADD_ITEM",
	2999999: "MSG_LOGICSVR_MSGID_END",
	3000000: "MSG_SCENESVR_MSGID_BEGIN",
	3000001: "MSG_CREATE_SCENE_REQ",
	3000002: "MSG_CREATE_SCENE_ACK",
	3000003: "MSG_DELETE_SCENE_REQ",
	3000004: "MSG_DELETE_SCENE_ACK",
	3000005: "MSG_SCENE_DESTROY_REQ",
	3000006: "MSG_SCENE_DESTROY_ACK",
	3000007: "MSG_TRANSFER_DATA_REQ",
	3000008: "MSG_TRANSFER_DATA_ACK",
	3000009: "MSG_NOTIFY_ROLE_ENTER_REQ",
	3000010: "MSG_NOTIFY_ROLE_ENTER_ACK",
	3000011: "MSG_ENTER_SCENE_REQ",
	3000012: "MSG_ENTER_SCENE_ACK",
	3000013: "MSG_LEAVE_SCENE_REQ",
	3000014: "MSG_LEAVE_SCENE_ACK",
	3000015: "MSG_SKILL_CAST_REQ",
	3000016: "MSG_SKILL_CAST_ACK",
	3000017: "MSG_SKILL_CAST_NTF",
	3000018: "MSG_SKILL_RESULT_NTF",
	3000019: "MSG_OBJECT_NEW_NTF",
	3000020: "MSG_OBJECT_CHANGE_NTF",
	3000021: "MSG_OBJECT_REMOVE_NTF",
	3000022: "MSG_OBJECT_ACTION_REQ",
	3000023: "MSG_OBJECT_ACTION_ACK",
	3000024: "MSG_HEART_BEAT_REQ",
	3000025: "MSG_HEART_BEAT_ACK",
	3000026: "MSG_USE_HP_BOOTTLE_REQ",
	3000027: "MSG_USE_HP_BOOTTLE_ACK",
	3000028: "MSG_USE_MP_BOOTTLE_REQ",
	3000029: "MSG_USE_MP_BOOTTLE_ACK",
	3000030: "MSG_OBJECT_DIE_NOTIFY",
	3000031: "MSG_BATTLE_CHAT_REQ",
	3000032: "MSG_BATTLE_CHAT_ACK",
	3000033: "MSG_BULLET_NEW_NTF",
	3000034: "MSG_SWITCH_MOUNT_REQ",
	3000035: "MSG_SWITCH_MOUNT_ACK",
	3000036: "MSG_SCENEOBJ_CHAGE_NTF",
	3999999: "MSG_SCENESVR_MSGID_END",
	1103:    "MSG_REQ_USEITEM",
	1104:    "MSG_ACK_USEITEM",
	1109:    "MSG_REQ_DRESS_GEM",
	1110:    "MSG_ACK_DRESS_GEM",
	1111:    "MSG_REQ_UNLOAD_GEM",
	1112:    "MSG_ACK_UNLOAD_GEM",
	1113:    "MSG_REQ_COMPOSE_CHIP",
	1114:    "MSG_ACK_COMPOSE_CHIP",
	1115:    "MSG_REQ_STRENGTHEN_EQUIP",
	1116:    "MSG_ACK_STRENGTHEN_EQUIP",
	1117:    "MSG_REQ_ADVANCE_EQUIP",
	1118:    "MSG_ACK_ADVANCE_EQUIP",
	1119:    "MSG_REQ_UPSTAR_EQUIP",
	1120:    "MSG_ACK_UPSTAR_EQUIP",
	1121:    "MSG_REQ_STRENGTHEN_GEM",
	1122:    "MSG_ACK_STRENGTHEN_GEM",
	1151:    "MSG_REQ_ONEKEYTODRESSGEM",
	1152:    "MSG_ACK_ONEKEYTODRESSGEM",
	1153:    "MSG_REQ_ONEKEYTOUNLOADGEM",
	1154:    "MSG_ACK_ONEKEYTOUNLOADGEM",
	1201:    "MSG_REQ_GET_CHAPTERAWARD",
	1202:    "MSG_ACK_GET_CHAPTERAWARD",
	1203:    "MSG_REQ_BATTLE_CHECK",
	1204:    "MSG_ACK_BATTLE_CHECK",
	1205:    "MSG_REQ_PASSCOPY",
	1206:    "MSG_ACK_PASSCOPY",
	1301:    "MSG_REQ_SETMOUNT",
	1302:    "MSG_ACK_SETMOUNT",
	1401:    "MSG_REQ_CHARGE_RELICS",
	1402:    "MSG_ACK_CHARGE_RELICS",
	1403:    "MSG_REQ_UPGRADE_RELICS",
	1404:    "MSG_ACK_UPGRADE_RELICS",
	1405:    "MSG_REQ_BATTLE_RELICS",
	1406:    "MSG_ACK_BATTLE_RELICS",
	1407:    "MSG_REQ_UNLOAD_RELICS",
	1408:    "MSG_ACK_UNLOAD_RELICS",
	1501:    "MSG_REQ_BUY_STORE",
	1502:    "MSG_ACK_BUY_STORE",
	1601:    "MSG_REQ_UPGRADE_PET",
	1602:    "MSG_ACK_UPGRADE_PET",
	1603:    "MSG_REQ_BATTLE_PET",
	1604:    "MSG_ACK_BATTLE_PET",
	1605:    "MSG_REQ_UNLOAD_PET",
	1606:    "MSG_ACK_UNLOAD_PET",
	1607:    "MSG_REQ_CHANGE_PARTNER",
	1608:    "MSG_ACK_CHANGE_PARTNER",
	1609:    "MSG_REQ_ADVANVE_PARTNER",
	1610:    "MSG_ACK_ADVANVE_PARTNER",
	1611:    "MSG_REQ_UPGRADE_PARTNER",
	1612:    "MSG_ACK_UPGRADE_PARTNER",
	1613:    "MSG_REQ_SUBMIT_TASK",
	1614:    "MSG_ACK_SUBMIT_TASK",
	2001:    "MSG_REQ_ADDHERO_EXP",
	2002:    "MSG_ACK_ADDHERO_EXP",
	3001:    "MSG_ENTER_VIEW_NTF",
	3002:    "MSG_LEAVE_VIEW_NTF",
	3003:    "MSG_ACTOR_VALUECHANGE_NTF",
	3004:    "MSG_ACTOR_SKILLCAST_NTF",
	3005:    "MSG_ACTOR_DIE_NTF",
}
var MessageID_value = map[string]int32{
	"MSG_BEGIN":                  0,
	"MSG_NORMAL_MSGID_BEGIN":     1000000,
	"MSG_CHECK_VERSION_REQ":      1000001,
	"MSG_CHECK_VERSION_ACK":      1000002,
	"MSG_ACCOUNT_REG_REQ":        1000003,
	"MSG_ACCOUNT_REG_ACK":        1000004,
	"MSG_ACCOUNT_LOGIN_REQ":      1000005,
	"MSG_ACCOUNT_LOGIN_ACK":      1000006,
	"MSG_DB_EXE_SQL_REQ":         1000007,
	"MSG_DB_EXE_SQL_ACK":         1000008,
	"MSG_SERVER_LIST_REQ":        1000009,
	"MSG_SERVER_LIST_ACK":        1000010,
	"MSG_SELECT_SERVER_REQ":      1000011,
	"MSG_SELECT_SERVER_ACK":      1000012,
	"MSG_LOGIC_REGTO_LOGIN_REQ":  1000013,
	"MSG_LOGIC_REGTO_LOGIN_ACK":  1000014,
	"MSG_SEAL_ACCOUNT_REQ":       1000015,
	"MSG_SEAL_ACCOUNT_ACK":       1000016,
	"MSG_GAME_REGTO_LOGIC_REQ":   1000017,
	"MSG_GAME_REGTO_LOGIC_ACK":   1000018,
	"MSG_LOGIC_REGTO_CENTER_REQ": 1000019,
	"MSG_LOGIC_REGTO_CENTER_ACK": 1000020,
	"MSG_LOG_DATA_NTF":           1000021,
	"MSG_PHP_GM_COMMAND_REQ":     1000022,
	"MSG_PHP_GM_COMMAND_ACK":     1000023,
	"MSG_WATCH_HEART_BEAT_REQ":   1000024,
	"MSG_WATCH_HEART_BEAT_ACK":   1000025,
	"MSG_WATCH_UPDATE_SVR_REQ":   1000026,
	"MSG_WATCH_UPDATE_SVR_ACK":   1000027,
	"MSG_WATCH_START_SVR_REQ":    1000028,
	"MSG_WATCH_START_SVR_ACK":    1000029,
	"MSG_WATCH_STOP_SVR_REQ":     1000030,
	"MSG_WATCH_STOP_SVR_ACK":     1000031,
	"MSG_DISCONNECT_NTY":         1000032,
	"MSG_ROLE_DATA_CHANGE_NTY":   1000033,
	"MSG_NORMAL_MSGID_END":       1999999,
	"MSG_LOGICSVR_MSGID_BEGIN":   2000000,
	"MSG_ROLE_LIST_REQ":          2000001,
	"MSG_ROLE_LIST_ACK":          2000002,
	"MSG_ROLE_CREATE_REQ":        2000003,
	"MSG_ROLE_CREATE_ACK":        2000004,
	"MSG_ROLE_DELETE_REQ":        2000005,
	"MSG_ROLE_DELETE_ACK":        2000006,
	"MSG_ROLE_LOGIN_REQ":         2000007,
	"MSG_ROLE_LOGIN_ACK":         2000008,
	"MSG_ROLE_LOGOUT_REQ":        2000009,
	"MSG_ROLE_LOGOUT_ACK":        2000010,
	"MSG_DATA_UPDATE":            2000011,
	"MSG_NOTIFY_INTO_SCENE":      2000014,
	"MSG_COPYINFO_REPORT_REQ":    2000016,
	"MSG_MAIN_COPY_REQ":          2000017,
	"MSG_MAIN_COPY_ACK":          2000018,
	"MSG_COPY_ABORT_REQ":         2000019,
	"MSG_COPY_ABORT_ACK":         2000020,
	"MSG_BACK_TO_CITY_REQ":       2000021,
	"MSG_BACK_TO_CITY_ACK":       2000022,
	"MSG_BATTLE_RESULT_NTY":      2000023,
	"MSG_MAINCOPY_RESULT_NTY":    2000024,
	"MSG_ROLE_CHANG_NTY":         2000025,
	"MSG_BAG_CHANGE_NTY":         2000026,
	"MSG_PET_CHANGE_NTY":         2000027,
	"MSG_EQUIP_CHANGE_NTY":       2000028,
	"MSG_PARTNER_CHANGE_NTY":     2000029,
	"MSG_MOUNT_CHANGE_NTY":       2000030,
	"MSG_ROLE_OTHER_LOGIN_NTY":   2000031,
	"MSG_CHAT_MESSAGE_REQ":       2000032,
	"MSG_CHAT_MESSAGE_ACK":       2000033,
	"MSG_CHAT_MESSAGE_NTY":       2000034,
	"MSG_SETUP_EQUIP_REQ":        2000035,
	"MSG_SETUP_EQUIP_ACK":        2000036,
	"MSG_UNSET_EQUIP_REQ":        2000037,
	"MSG_UNSET_EQUIP_ACK":        2000038,
	"MSG_SETUP_PET_REQ":          2000039,
	"MSG_SETUP_PET_ACK":          2000040,
	"MSG_SETUP_PARTNER_REQ":      2000041,
	"MSG_SETUP_PARTNER_ACK":      2000042,
	"MSG_UNSET_PARTNER_REQ":      2000043,
	"MSG_UNSET_PARTNER_ACK":      2000044,
	"MSG_SETUP_MOUNT_REQ":        2000045,
	"MSG_SETUP_MOUNT_ACK":        2000046,
	"MSG_UNSET_MOUNT_REQ":        2000047,
	"MSG_UNSET_MOUNT_ACK":        2000048,
	"MSG_USE_ITEM_REQ":           2000049,
	"MSG_USE_ITEM_ACK":           2000050,
	"MSG_SETUP_GEM_REQ":          2000051,
	"MSG_SETUP_GEM_ACK":          2000052,
	"MSG_UNSET_GEM_REQ":          2000053,
	"MSG_UNSET_GEM_ACK":          2000054,
	"MSG_ROLE_RECONNECT_REQ":     2000055,
	"MSG_ROLE_RECONNECT_ACK":     2000056,
	"MSG_MAIL_CHANGE_NTY":        2000057,
	"MSG_TASK_CHANGE_NTY":        2000058,
	"MSG_BROAD_MESSAGE_NOTIFY":   2000060,
	"MSG_GASVR_REGTO_PROXY_REQ":  2000061,
	"MSG_GASVR_REGTO_PROXY_ACK":  2000062,
	"MSG_STORE_BUY_REQ":          2000063,
	"MSG_STORE_BUY_ACK":          2000064,
	"MSG_GEM_CHANGE_NTY":         2000065,
	"MSG_TEST_ADD_ITEM":          2999998,
	"MSG_LOGICSVR_MSGID_END":     2999999,
	"MSG_SCENESVR_MSGID_BEGIN":   3000000,
	"MSG_CREATE_SCENE_REQ":       3000001,
	"MSG_CREATE_SCENE_ACK":       3000002,
	"MSG_DELETE_SCENE_REQ":       3000003,
	"MSG_DELETE_SCENE_ACK":       3000004,
	"MSG_SCENE_DESTROY_REQ":      3000005,
	"MSG_SCENE_DESTROY_ACK":      3000006,
	"MSG_TRANSFER_DATA_REQ":      3000007,
	"MSG_TRANSFER_DATA_ACK":      3000008,
	"MSG_NOTIFY_ROLE_ENTER_REQ":  3000009,
	"MSG_NOTIFY_ROLE_ENTER_ACK":  3000010,
	"MSG_ENTER_SCENE_REQ":        3000011,
	"MSG_ENTER_SCENE_ACK":        3000012,
	"MSG_LEAVE_SCENE_REQ":        3000013,
	"MSG_LEAVE_SCENE_ACK":        3000014,
	"MSG_SKILL_CAST_REQ":         3000015,
	"MSG_SKILL_CAST_ACK":         3000016,
	"MSG_SKILL_CAST_NTF":         3000017,
	"MSG_SKILL_RESULT_NTF":       3000018,
	"MSG_OBJECT_NEW_NTF":         3000019,
	"MSG_OBJECT_CHANGE_NTF":      3000020,
	"MSG_OBJECT_REMOVE_NTF":      3000021,
	"MSG_OBJECT_ACTION_REQ":      3000022,
	"MSG_OBJECT_ACTION_ACK":      3000023,
	"MSG_HEART_BEAT_REQ":         3000024,
	"MSG_HEART_BEAT_ACK":         3000025,
	"MSG_USE_HP_BOOTTLE_REQ":     3000026,
	"MSG_USE_HP_BOOTTLE_ACK":     3000027,
	"MSG_USE_MP_BOOTTLE_REQ":     3000028,
	"MSG_USE_MP_BOOTTLE_ACK":     3000029,
	"MSG_OBJECT_DIE_NOTIFY":      3000030,
	"MSG_BATTLE_CHAT_REQ":        3000031,
	"MSG_BATTLE_CHAT_ACK":        3000032,
	"MSG_BULLET_NEW_NTF":         3000033,
	"MSG_SWITCH_MOUNT_REQ":       3000034,
	"MSG_SWITCH_MOUNT_ACK":       3000035,
	"MSG_SCENEOBJ_CHAGE_NTF":     3000036,
	"MSG_SCENESVR_MSGID_END":     3999999,
	"MSG_REQ_USEITEM":            1103,
	"MSG_ACK_USEITEM":            1104,
	"MSG_REQ_DRESS_GEM":          1109,
	"MSG_ACK_DRESS_GEM":          1110,
	"MSG_REQ_UNLOAD_GEM":         1111,
	"MSG_ACK_UNLOAD_GEM":         1112,
	"MSG_REQ_COMPOSE_CHIP":       1113,
	"MSG_ACK_COMPOSE_CHIP":       1114,
	"MSG_REQ_STRENGTHEN_EQUIP":   1115,
	"MSG_ACK_STRENGTHEN_EQUIP":   1116,
	"MSG_REQ_ADVANCE_EQUIP":      1117,
	"MSG_ACK_ADVANCE_EQUIP":      1118,
	"MSG_REQ_UPSTAR_EQUIP":       1119,
	"MSG_ACK_UPSTAR_EQUIP":       1120,
	"MSG_REQ_STRENGTHEN_GEM":     1121,
	"MSG_ACK_STRENGTHEN_GEM":     1122,
	"MSG_REQ_ONEKEYTODRESSGEM":   1151,
	"MSG_ACK_ONEKEYTODRESSGEM":   1152,
	"MSG_REQ_ONEKEYTOUNLOADGEM":  1153,
	"MSG_ACK_ONEKEYTOUNLOADGEM":  1154,
	"MSG_REQ_GET_CHAPTERAWARD":   1201,
	"MSG_ACK_GET_CHAPTERAWARD":   1202,
	"MSG_REQ_BATTLE_CHECK":       1203,
	"MSG_ACK_BATTLE_CHECK":       1204,
	"MSG_REQ_PASSCOPY":           1205,
	"MSG_ACK_PASSCOPY":           1206,
	"MSG_REQ_SETMOUNT":           1301,
	"MSG_ACK_SETMOUNT":           1302,
	"MSG_REQ_CHARGE_RELICS":      1401,
	"MSG_ACK_CHARGE_RELICS":      1402,
	"MSG_REQ_UPGRADE_RELICS":     1403,
	"MSG_ACK_UPGRADE_RELICS":     1404,
	"MSG_REQ_BATTLE_RELICS":      1405,
	"MSG_ACK_BATTLE_RELICS":      1406,
	"MSG_REQ_UNLOAD_RELICS":      1407,
	"MSG_ACK_UNLOAD_RELICS":      1408,
	"MSG_REQ_BUY_STORE":          1501,
	"MSG_ACK_BUY_STORE":          1502,
	"MSG_REQ_UPGRADE_PET":        1601,
	"MSG_ACK_UPGRADE_PET":        1602,
	"MSG_REQ_BATTLE_PET":         1603,
	"MSG_ACK_BATTLE_PET":         1604,
	"MSG_REQ_UNLOAD_PET":         1605,
	"MSG_ACK_UNLOAD_PET":         1606,
	"MSG_REQ_CHANGE_PARTNER":     1607,
	"MSG_ACK_CHANGE_PARTNER":     1608,
	"MSG_REQ_ADVANVE_PARTNER":    1609,
	"MSG_ACK_ADVANVE_PARTNER":    1610,
	"MSG_REQ_UPGRADE_PARTNER":    1611,
	"MSG_ACK_UPGRADE_PARTNER":    1612,
	"MSG_REQ_SUBMIT_TASK":        1613,
	"MSG_ACK_SUBMIT_TASK":        1614,
	"MSG_REQ_ADDHERO_EXP":        2001,
	"MSG_ACK_ADDHERO_EXP":        2002,
	"MSG_ENTER_VIEW_NTF":         3001,
	"MSG_LEAVE_VIEW_NTF":         3002,
	"MSG_ACTOR_VALUECHANGE_NTF":  3003,
	"MSG_ACTOR_SKILLCAST_NTF":    3004,
	"MSG_ACTOR_DIE_NTF":          3005,
}

func (x MessageID) String() string {
	return proto.EnumName(MessageID_name, int32(x))
}
func (MessageID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("MessageID", MessageID_name, MessageID_value)
}

func init() { proto.RegisterFile("Msg_ID.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x98, 0xc9, 0x93, 0xdc, 0x44,
	0x16, 0xc6, 0xc7, 0x97, 0x89, 0x2e, 0xb9, 0x26, 0xac, 0xd1, 0xd8, 0x2e, 0x77, 0x7b, 0x3d, 0x0f,
	0x11, 0x5c, 0x38, 0xd7, 0x21, 0x4b, 0xca, 0x56, 0xc9, 0x2d, 0x29, 0x55, 0x52, 0x56, 0xb5, 0xfb,
	0xa4, 0x80, 0x08, 0x87, 0x8f, 0x26, 0x30, 0xa7, 0x3a, 0xd9, 0xa6, 0x00, 0x17, 0x04, 0x04, 0x06,
	0x6c, 0xf6, 0xc5, 0xec, 0x3b, 0x78, 0xc7, 0xc6, 0xfb, 0xbe, 0xef, 0xfb, 0xd2, 0x5e, 0xe0, 0x4f,
	0xe0, 0xc2, 0xda, 0xc4, 0xcb, 0x4c, 0x29, 0x53, 0x6a, 0xa9, 0x8e, 0x9d, 0xbf, 0xfc, 0x52, 0xf9,
	0x5e, 0xbe, 0x97, 0xf9, 0x55, 0x6b, 0x55, 0x6f, 0xd5, 0x8a, 0xd8, 0xb1, 0x1e, 0x7d, 0xf2, 0xa9,
	0x95, 0x4f, 0xaf, 0xfc, 0xff, 0xa6, 0x47, 0xb4, 0x8a, 0xb7, 0x7c, 0xd5, 0xaa, 0xc7, 0x57, 0x2c,
	0x77, 0x2c, 0xe3, 0x3f, 0x5a, 0xc5, 0x8b, 0xec, 0xb8, 0x81, 0x6d, 0xc7, 0xd7, 0xff, 0x65, 0x2c,
	0xd0, 0xe6, 0xc2, 0x9f, 0x3e, 0x09, 0x3d, 0xe4, 0xc6, 0x5e, 0x64, 0x3b, 0x96, 0x60, 0x07, 0x7a,
	0x75, 0x63, 0xbe, 0x36, 0x07, 0xa8, 0xd9, 0xc4, 0xe6, 0x58, 0xdc, 0xc1, 0x61, 0xe4, 0x10, 0x3f,
	0x0e, 0x71, 0x4b, 0x3f, 0x58, 0x06, 0x91, 0x39, 0xa6, 0x1f, 0xea, 0xd5, 0x8d, 0x61, 0xed, 0x7f,
	0x00, 0x91, 0x69, 0x92, 0xb6, 0x4f, 0xe3, 0x10, 0xdb, 0x4c, 0x77, 0xb8, 0x18, 0x81, 0xea, 0x88,
	0x5c, 0x32, 0x41, 0x2e, 0xb1, 0x1d, 0xfe, 0xbd, 0xa3, 0x65, 0x10, 0x94, 0xc7, 0x7a, 0x75, 0x63,
	0x9e, 0x66, 0x00, 0xb4, 0x1a, 0x31, 0x5e, 0x86, 0xe3, 0xa8, 0xe5, 0x32, 0xd9, 0xf1, 0x42, 0x02,
	0x9a, 0x13, 0x72, 0x23, 0x11, 0x0e, 0x3b, 0x38, 0x8c, 0x5d, 0x27, 0xa2, 0x4c, 0x74, 0xb2, 0x18,
	0x81, 0xea, 0x94, 0xdc, 0x46, 0x84, 0x5d, 0x6c, 0xd2, 0x64, 0x06, 0xe8, 0x4e, 0x97, 0x41, 0x50,
	0x9e, 0xe9, 0xd5, 0x8d, 0xc5, 0xda, 0x30, 0x40, 0xd8, 0xb8, 0x09, 0x61, 0x53, 0xa2, 0x44, 0x78,
	0x76, 0xd0, 0x04, 0x58, 0xe1, 0x5c, 0xaf, 0x6e, 0x8c, 0x68, 0xb3, 0xf9, 0xf2, 0xc8, 0x55, 0xf2,
	0xd7, 0xd2, 0xcf, 0x97, 0x30, 0xd0, 0x5d, 0xe8, 0xd5, 0x8d, 0x45, 0xda, 0x3c, 0x60, 0x36, 0xf2,
	0xb0, 0xb2, 0xae, 0xc9, 0xb4, 0x17, 0x07, 0x70, 0xd0, 0x5f, 0xea, 0xd5, 0x8d, 0x25, 0xda, 0x48,
	0x7e, 0x63, 0x26, 0xf6, 0xa9, 0x08, 0xfc, 0xf2, 0xc0, 0x19, 0xb0, 0xc6, 0x95, 0x5e, 0xdd, 0x98,
	0xab, 0xe9, 0x62, 0x46, 0x6c, 0x21, 0x8a, 0x62, 0x9f, 0x8e, 0xea, 0x57, 0x7b, 0xf5, 0xa4, 0x02,
	0x83, 0x66, 0x10, 0xdb, 0x5e, 0x6c, 0x12, 0xcf, 0x43, 0xbe, 0xc5, 0xd6, 0xbd, 0x56, 0x4a, 0x61,
	0xcd, 0xeb, 0x72, 0xdf, 0xe3, 0x88, 0x9a, 0xcd, 0xb8, 0x89, 0x51, 0x48, 0xe3, 0x06, 0x46, 0x3c,
	0x27, 0x37, 0x06, 0x70, 0xd0, 0xdf, 0xcc, 0xf3, 0x76, 0x60, 0x21, 0x8a, 0xe3, 0xa8, 0xc3, 0xa3,
	0xba, 0x35, 0x80, 0x83, 0xfe, 0x76, 0xaf, 0x6e, 0x2c, 0xd4, 0x6a, 0x92, 0x47, 0x14, 0xd6, 0x4f,
	0xe4, 0x77, 0xca, 0x31, 0xa8, 0xef, 0xca, 0xd8, 0x12, 0x4c, 0x82, 0x54, 0x3c, 0x59, 0x4a, 0x41,
	0x7b, 0x4f, 0xa9, 0x6a, 0x27, 0x32, 0x89, 0xef, 0x43, 0xb1, 0xf9, 0x74, 0x42, 0xbf, 0x2f, 0xf7,
	0x1c, 0x12, 0x17, 0xf3, 0x44, 0x9b, 0x4d, 0xe4, 0xdb, 0x98, 0xf1, 0x07, 0xb2, 0x4e, 0x32, 0x1d,
	0x8f, 0x7d, 0x4b, 0x9f, 0x5a, 0xd7, 0x4d, 0xb4, 0xec, 0x14, 0xe1, 0x6b, 0xea, 0x7d, 0xb0, 0xba,
	0xdf, 0x35, 0x6a, 0xda, 0x7f, 0xd3, 0xb5, 0xd3, 0x7e, 0x59, 0x53, 0x04, 0x60, 0x9f, 0x6b, 0xfb,
	0xdd, 0xa4, 0x91, 0x18, 0x30, 0x43, 0x0c, 0x09, 0x04, 0xcd, 0x33, 0xc5, 0x08, 0x54, 0xbd, 0x1c,
	0xb2, 0xb0, 0x8b, 0x85, 0xea, 0xd9, 0x62, 0x04, 0xaa, 0xe7, 0xfa, 0xdd, 0x24, 0x27, 0x7c, 0x13,
	0x69, 0x63, 0x3d, 0x5f, 0x48, 0x40, 0xb3, 0x2e, 0xb7, 0x9c, 0x4b, 0x6c, 0xd2, 0xe6, 0x31, 0xf5,
	0x8b, 0x11, 0xa8, 0x5e, 0xe8, 0x77, 0x8d, 0x39, 0xda, 0x2c, 0x96, 0x7d, 0x48, 0x2f, 0x2f, 0x0b,
	0xfd, 0xc5, 0x7e, 0x37, 0xe9, 0x7e, 0x9f, 0x50, 0x67, 0x74, 0x22, 0x76, 0x7c, 0x4a, 0xe2, 0xc8,
	0xc4, 0x3e, 0xd6, 0x5f, 0xee, 0x77, 0x93, 0x62, 0x30, 0x49, 0x30, 0xe1, 0xf8, 0xa3, 0x24, 0x0e,
	0x71, 0x40, 0x42, 0xfe, 0xb5, 0x57, 0x64, 0x06, 0x3d, 0xe4, 0xf8, 0x6c, 0x0e, 0x03, 0xeb, 0x8b,
	0x00, 0x6c, 0xe2, 0x55, 0x19, 0x14, 0x1f, 0x6b, 0x24, 0x6b, 0xbd, 0x56, 0x48, 0x40, 0xf3, 0x7a,
	0xbf, 0x9b, 0x1c, 0x7e, 0x03, 0x99, 0x63, 0x31, 0xb4, 0xa8, 0x43, 0xf9, 0x87, 0x36, 0x94, 0x30,
	0xd0, 0x6d, 0x94, 0x91, 0x35, 0x10, 0xa5, 0x2e, 0x9c, 0x45, 0xd4, 0x76, 0x79, 0xc5, 0xbd, 0x21,
	0x23, 0x83, 0x1d, 0x8a, 0x9d, 0xa7, 0xf8, 0xcd, 0x5c, 0xf2, 0x59, 0x2d, 0x32, 0xf2, 0x96, 0x24,
	0x0d, 0x64, 0xab, 0x45, 0xfa, 0xb6, 0x24, 0x01, 0xa6, 0x2a, 0x79, 0x47, 0xee, 0x12, 0xb7, 0xda,
	0x4e, 0xa0, 0xb2, 0x77, 0xfb, 0xdd, 0xf4, 0xb2, 0x40, 0x21, 0xf5, 0x71, 0xa8, 0xd2, 0xf7, 0xa4,
	0xd2, 0x63, 0x37, 0xa3, 0xc2, 0xde, 0xef, 0x77, 0x33, 0x4d, 0x43, 0x68, 0x13, 0xee, 0x7c, 0x56,
	0x26, 0xc0, 0x3f, 0x90, 0x5a, 0xb3, 0x89, 0x68, 0xec, 0xe1, 0x28, 0x42, 0x36, 0xaf, 0xc8, 0x0f,
	0x4b, 0x18, 0xe4, 0x6d, 0x53, 0x09, 0x83, 0x35, 0x3f, 0x92, 0xf5, 0x15, 0x61, 0xda, 0x0e, 0x44,
	0x3c, 0xb0, 0xe4, 0xc7, 0xc5, 0x08, 0x56, 0xfc, 0x44, 0xa2, 0xb6, 0x1f, 0x61, 0xaa, 0xa8, 0x3e,
	0x2d, 0x46, 0xa0, 0xfa, 0x4c, 0x16, 0x11, 0x5f, 0x10, 0xb2, 0x0a, 0x9a, 0xcf, 0x8b, 0x00, 0x28,
	0xbe, 0x90, 0x27, 0x2e, 0x80, 0xc8, 0x28, 0xa8, 0xbe, 0x2c, 0x83, 0xa0, 0xfc, 0x4a, 0x42, 0xbe,
	0x0d, 0x55, 0xf9, 0x75, 0x19, 0x04, 0xe5, 0x37, 0xf9, 0xb0, 0xbd, 0xf4, 0x75, 0xfb, 0xb6, 0x18,
	0x81, 0xea, 0xbb, 0x7c, 0xd8, 0x52, 0xf5, 0x7d, 0x31, 0x02, 0xd5, 0x0f, 0xfd, 0x6e, 0xf2, 0x1c,
	0xb5, 0x23, 0x1c, 0x3b, 0x14, 0x7b, 0x4c, 0xb2, 0xb9, 0x60, 0x1c, 0xe6, 0x6f, 0xc9, 0x27, 0xca,
	0x16, 0x82, 0xad, 0x45, 0x00, 0x14, 0xdb, 0x24, 0xe0, 0x1f, 0x4f, 0x14, 0xdb, 0x8b, 0x00, 0x28,
	0x76, 0xc8, 0xfa, 0x65, 0x55, 0x18, 0xe2, 0xe4, 0x62, 0x07, 0xd9, 0xce, 0x52, 0x0a, 0xda, 0x1f,
	0x65, 0xa8, 0x1e, 0x72, 0x5c, 0xb5, 0xb8, 0x77, 0x49, 0x44, 0x51, 0x34, 0xa6, 0xa2, 0xdd, 0xb2,
	0xee, 0x1b, 0x21, 0x41, 0x96, 0x2c, 0x50, 0x76, 0x7f, 0xe9, 0x7b, 0xfa, 0xdd, 0xc4, 0x91, 0xd8,
	0x88, 0xbf, 0x4c, 0xf0, 0xac, 0x07, 0x21, 0x59, 0xc6, 0x2f, 0x8d, 0xbd, 0x83, 0x26, 0xc0, 0xbe,
	0xf6, 0x29, 0xe9, 0xa1, 0x24, 0xc4, 0x71, 0xa3, 0xcd, 0x95, 0xfb, 0x8b, 0x00, 0x28, 0x0e, 0xc8,
	0xde, 0x87, 0xc4, 0x28, 0xbb, 0x3d, 0xc8, 0x08, 0x93, 0x50, 0x0c, 0x0f, 0x8c, 0x65, 0xb1, 0x03,
	0xd2, 0xf7, 0xbd, 0xb4, 0x73, 0x86, 0xb1, 0x90, 0xe7, 0x26, 0xf7, 0x70, 0xc1, 0xb3, 0xb6, 0x1f,
	0xf0, 0x62, 0x1e, 0x26, 0xbb, 0x8c, 0xf3, 0xef, 0xda, 0x01, 0x98, 0x30, 0x5f, 0xf4, 0x29, 0x7f,
	0x86, 0xd8, 0x3c, 0xee, 0x73, 0xcb, 0x20, 0xf3, 0xb9, 0x0a, 0x14, 0x4f, 0x91, 0x54, 0x1e, 0x2e,
	0x83, 0xcc, 0xeb, 0x02, 0x5c, 0x20, 0x9a, 0x88, 0x8d, 0x5a, 0x38, 0xa2, 0x21, 0xe1, 0xd9, 0x39,
	0x5a, 0x4a, 0x99, 0xdb, 0x55, 0x28, 0x0d, 0x91, 0x1f, 0x8d, 0xe2, 0x90, 0xbf, 0x44, 0xcc, 0xf1,
	0x96, 0x52, 0xe6, 0x7a, 0x81, 0x2e, 0xe1, 0x47, 0x26, 0x5e, 0x29, 0x56, 0x4e, 0xd2, 0xcb, 0x9d,
	0x1c, 0x38, 0x83, 0x79, 0x60, 0x98, 0x31, 0xc2, 0x4b, 0x8a, 0x8f, 0xca, 0xa0, 0x4f, 0x97, 0x30,
	0xe6, 0x80, 0x15, 0xe6, 0x62, 0xd4, 0x51, 0x93, 0x75, 0xb6, 0x84, 0x31, 0xdf, 0x0b, 0x6c, 0x98,
	0xd7, 0x44, 0x34, 0xe6, 0xb8, 0x6e, 0x6c, 0x22, 0xe1, 0x3c, 0xce, 0x17, 0x23, 0xe6, 0x7a, 0x8b,
	0x11, 0x98, 0xce, 0x8b, 0xca, 0xc9, 0x70, 0x94, 0x3e, 0x58, 0xa3, 0xfa, 0x25, 0x45, 0x47, 0x1a,
	0x4b, 0x99, 0xb1, 0xc2, 0xe3, 0x0c, 0x5d, 0x56, 0x52, 0x2b, 0x50, 0x5a, 0x9f, 0xa3, 0xfa, 0x95,
	0xe9, 0x34, 0xc4, 0x1e, 0xe9, 0x70, 0x7a, 0x75, 0x3a, 0x45, 0x26, 0x4d, 0x7e, 0x4d, 0x5d, 0x2b,
	0xa5, 0xcc, 0xe8, 0x2a, 0x5b, 0xca, 0x7b, 0xdc, 0x62, 0xc4, 0xec, 0xad, 0xd2, 0x16, 0x70, 0x99,
	0x35, 0x83, 0xb8, 0x41, 0x88, 0x78, 0xbe, 0x5b, 0xfa, 0xad, 0x72, 0xcc, 0xcc, 0x6d, 0x0e, 0x7b,
	0x59, 0xf5, 0x9d, 0x72, 0xcc, 0xcc, 0xed, 0xf4, 0x78, 0x2c, 0x27, 0xbd, 0x57, 0x26, 0x95, 0xc3,
	0x16, 0x86, 0x82, 0xbd, 0x8f, 0xb0, 0xf0, 0xbd, 0x12, 0x06, 0xab, 0xde, 0x57, 0x82, 0x6d, 0xb4,
	0x5d, 0x17, 0xcb, 0xa3, 0x79, 0xa0, 0x1e, 0xe9, 0xb8, 0x03, 0x8e, 0x59, 0x3e, 0x04, 0x0f, 0xcb,
	0x20, 0x2c, 0xfa, 0xb3, 0x12, 0x09, 0xab, 0x39, 0xd2, 0x58, 0x0a, 0x9f, 0x14, 0xa7, 0xfa, 0x4b,
	0x1e, 0x67, 0x2f, 0x97, 0xa9, 0xf5, 0xbf, 0xce, 0x30, 0x66, 0x73, 0x33, 0x18, 0xe2, 0x16, 0xe4,
	0x82, 0xdd, 0x49, 0xe7, 0x87, 0x92, 0x51, 0x30, 0x53, 0xc9, 0xe8, 0x85, 0x21, 0x63, 0xae, 0xf0,
	0xc9, 0xb8, 0x15, 0x5b, 0x21, 0x8e, 0x22, 0xb8, 0xe5, 0xf4, 0xab, 0xe9, 0x38, 0xcc, 0x96, 0xe3,
	0xd7, 0x86, 0x8c, 0x9a, 0xf0, 0x4e, 0xb0, 0xb6, 0xef, 0xc2, 0x25, 0x0d, 0xe0, 0x7a, 0x0a, 0xd8,
	0xf2, 0x12, 0xdc, 0x18, 0x32, 0x86, 0x79, 0xa0, 0xa0, 0x30, 0x89, 0x17, 0x90, 0x08, 0x32, 0xe8,
	0x04, 0xfa, 0xcd, 0x14, 0x81, 0x26, 0x83, 0x6e, 0x0d, 0x19, 0x0b, 0x85, 0xff, 0xc1, 0xad, 0x38,
	0xa2, 0x21, 0xf6, 0x6d, 0xda, 0xc4, 0x3e, 0xf7, 0x10, 0xfa, 0xed, 0x14, 0x83, 0x72, 0x1a, 0xbe,
	0x33, 0x64, 0x8c, 0xf0, 0xa3, 0x06, 0x35, 0xb2, 0x3a, 0xc8, 0x37, 0xb1, 0x60, 0x77, 0x53, 0x06,
	0xd2, 0x2c, 0x9b, 0xcc, 0xec, 0xb5, 0x1d, 0xc0, 0xef, 0x23, 0x81, 0xee, 0x65, 0xf6, 0x9a, 0x41,
	0xf7, 0x87, 0x8c, 0xf9, 0xe2, 0x1d, 0xcc, 0xee, 0x15, 0xc2, 0x7f, 0x90, 0xc2, 0xdc, 0x4e, 0x01,
	0x3e, 0xcc, 0x44, 0x49, 0x7c, 0x3c, 0x86, 0x27, 0x28, 0x61, 0xd9, 0x06, 0x3c, 0x95, 0x89, 0x72,
	0x1a, 0x5e, 0x5d, 0x31, 0x16, 0xf1, 0x5b, 0x51, 0x55, 0xf3, 0xd4, 0x03, 0x5f, 0x93, 0x72, 0x55,
	0x2e, 0xf9, 0xda, 0x8a, 0xfa, 0x75, 0x9b, 0xfb, 0xda, 0x80, 0xe2, 0x10, 0x8d, 0xa3, 0xd0, 0xd2,
	0x37, 0x57, 0xd4, 0xaf, 0x4f, 0xc3, 0x5b, 0x2a, 0x6a, 0xae, 0xd2, 0xc6, 0xc0, 0xe6, 0x98, 0xbe,
	0xb5, 0xa2, 0xe6, 0x2a, 0x83, 0xb6, 0x55, 0x8c, 0x39, 0xdc, 0xcd, 0x80, 0x2a, 0x40, 0x51, 0x04,
	0xf6, 0x5c, 0xdf, 0x9e, 0x0e, 0x83, 0x22, 0x1d, 0xde, 0x91, 0x99, 0x1d, 0x61, 0xca, 0x9a, 0x44,
	0xdf, 0xa0, 0xa9, 0xb3, 0xd3, 0xe1, 0x8d, 0x9a, 0x7a, 0xea, 0x66, 0x13, 0x85, 0xcc, 0x11, 0xbb,
	0x8e, 0x19, 0xe9, 0xbf, 0x69, 0xea, 0xa9, 0x67, 0xd9, 0xef, 0x9a, 0x7a, 0x7e, 0xed, 0xc0, 0x0e,
	0x91, 0x95, 0xc2, 0x3f, 0x34, 0xf5, 0xfc, 0x72, 0xf0, 0xcf, 0xcc, 0x17, 0xd3, 0x5f, 0x22, 0x8c,
	0xfd, 0x95, 0xf9, 0x62, 0x96, 0xfd, 0x9d, 0xd1, 0x89, 0x66, 0x11, 0x6c, 0x2a, 0xa3, 0xcb, 0xb2,
	0xd5, 0x33, 0xd5, 0x6e, 0x05, 0x83, 0xc2, 0xac, 0x8a, 0x7e, 0x77, 0xa6, 0xda, 0xad, 0x72, 0x7c,
	0x72, 0xa6, 0x31, 0x4f, 0xfc, 0x62, 0x54, 0x22, 0x0b, 0x30, 0xd5, 0x0f, 0x56, 0x13, 0xa2, 0x86,
	0x05, 0xe4, 0x50, 0x55, 0xed, 0x70, 0xb1, 0x6f, 0x00, 0x87, 0xab, 0x6a, 0x87, 0x2b, 0xe0, 0x48,
	0xb5, 0xe0, 0x4e, 0x00, 0x70, 0xb4, 0x5a, 0x70, 0x27, 0x00, 0x38, 0x56, 0x55, 0x33, 0x2e, 0x5e,
	0x2c, 0xe1, 0xbc, 0xf5, 0xe3, 0x55, 0x35, 0xe3, 0x39, 0x78, 0xa2, 0x6a, 0x2c, 0xe0, 0x3f, 0xed,
	0xd2, 0xce, 0xee, 0x48, 0x7a, 0x32, 0xa5, 0x69, 0x6f, 0x2b, 0xf4, 0x54, 0x46, 0x9b, 0xc6, 0x2c,
	0xe8, 0xe9, 0x8c, 0x36, 0x4f, 0xcf, 0x54, 0xd5, 0x4c, 0x46, 0xed, 0x86, 0xe7, 0x50, 0xe6, 0x5e,
	0xf5, 0xb3, 0x99, 0x4c, 0xaa, 0xe4, 0x5c, 0x46, 0x83, 0x2c, 0xab, 0x89, 0x43, 0x12, 0xe3, 0x65,
	0x81, 0x7e, 0x71, 0x96, 0xaa, 0x51, 0xc9, 0xa5, 0x59, 0x49, 0xca, 0xb8, 0x57, 0xe9, 0x38, 0xe2,
	0x39, 0xd9, 0x55, 0x4b, 0x00, 0x37, 0x23, 0x29, 0xd8, 0x5d, 0x93, 0x5d, 0x4e, 0x49, 0x18, 0x77,
	0x90, 0xdb, 0xc6, 0x8a, 0x09, 0xf8, 0xa9, 0x26, 0xe3, 0x02, 0xce, 0xec, 0x45, 0x6a, 0x3c, 0xf6,
	0xd4, 0x64, 0xe5, 0x00, 0x65, 0x6f, 0x22, 0x1d, 0xd5, 0xf7, 0xd6, 0x9e, 0xf8, 0x37, 0xfb, 0x57,
	0xed, 0x63, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x22, 0x73, 0x8d, 0x90, 0xba, 0x15, 0x00, 0x00,
}
