// Code generated by "stringer -type=CmdID,CmdDirection,CmdCategory,CmdStatus -output=lbs_message.string.go"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[lbsLineCheck-24577]
	_ = x[lbsLogout-24578]
	_ = x[lbsShutDown-24579]
	_ = x[lbsVSUserLost-24580]
	_ = x[lbsSendMail-26372]
	_ = x[lbsRecvMail-26373]
	_ = x[lbsManagerMessage-26374]
	_ = x[lbsLoginType-24848]
	_ = x[lbsConnectionID-24833]
	_ = x[lbsAskConnectionID-24834]
	_ = x[lbsWarningMessage-24835]
	_ = x[lbsEncodeStart-24890]
	_ = x[lbsUserInfo1-24881]
	_ = x[lbsUserInfo2-24882]
	_ = x[lbsUserInfo3-24883]
	_ = x[lbsUserInfo4-24884]
	_ = x[lbsUserInfo5-24885]
	_ = x[lbsUserInfo6-24886]
	_ = x[lbsUserInfo7-24887]
	_ = x[lbsUserInfo8-24888]
	_ = x[lbsUserInfo9-24889]
	_ = x[lbsRegulationHeader-26656]
	_ = x[lbsRegulationText-26657]
	_ = x[lbsRegulationFooter-26658]
	_ = x[lbsUserHandle-24849]
	_ = x[lbsUserRegist-24850]
	_ = x[lbsUserDecide-24851]
	_ = x[lbsAskPlatformCode-24852]
	_ = x[lbsAskCountryCode-24853]
	_ = x[lbsAskGameCode-24854]
	_ = x[lbsAskGameVersion-24855]
	_ = x[lbsLoginOk-24856]
	_ = x[lbsAskBattleResult-24864]
	_ = x[lbsAskKDDICharges-24898]
	_ = x[lbsPostGameParameter-24899]
	_ = x[lbsWinLose-24901]
	_ = x[lbsRankRanking-24900]
	_ = x[lbsDeviceData-24904]
	_ = x[lbsServerMoney-24905]
	_ = x[lbsAskNewsTag-26625]
	_ = x[lbsNewsText-26626]
	_ = x[lbsInvitationTag-26640]
	_ = x[lbsTopRankingTag-26705]
	_ = x[lbsTopRankingSuu-26706]
	_ = x[lbsTopRanking-26707]
	_ = x[lbsAskPatchData-26721]
	_ = x[lbsPatchHeader-26722]
	_ = x[lbsPatchData6863-26723]
	_ = x[lbsCalcDownloadChecksum-26724]
	_ = x[lbsPatchPing-26725]
	_ = x[lbsStartLobby-24897]
	_ = x[lbsPlazaMax-25091]
	_ = x[lbsPlazaTitle-25092]
	_ = x[lbsPlazaJoin-25093]
	_ = x[lbsPlazaStatus-25094]
	_ = x[lbsPlazaExplain-25098]
	_ = x[lbsPlazaEntry-25095]
	_ = x[lbsPlazaExit-25350]
	_ = x[lbsLobbyJoin-25347]
	_ = x[lbsLobbyEntry-25349]
	_ = x[lbsLobbyExit-25608]
	_ = x[lbsLobbyMatchingJoin-25615]
	_ = x[lbsRoomMax-25601]
	_ = x[lbsRoomTitle-25602]
	_ = x[lbsRoomStatus-25604]
	_ = x[lbsRoomCreate-25607]
	_ = x[lbsPutRoomName-26121]
	_ = x[lbsEndRoomCreate-26124]
	_ = x[lbsRoomEntry-25606]
	_ = x[lbsRoomExit-25857]
	_ = x[lbsRoomLeaver-25858]
	_ = x[lbsRoomCommer-25859]
	_ = x[lbsMatchingEntry-25860]
	_ = x[lbsRoomRemove-25861]
	_ = x[lbsWaitJoin-25862]
	_ = x[lbsRoomUserReject-25863]
	_ = x[lbsPostChatMessage-26369]
	_ = x[lbsChatMessage-26370]
	_ = x[lbsUserSite-26371]
	_ = x[lbsLobbyRemove-25792]
	_ = x[lbsLobbyMatchingEntry-25614]
	_ = x[lbsGoToTop-25096]
	_ = x[lbsReadyBattle-26896]
	_ = x[lbsAskMatchingJoin-26897]
	_ = x[lbsAskPlayerSide-26898]
	_ = x[lbsAskPlayerInfo-26899]
	_ = x[lbsAskRuleData-26900]
	_ = x[lbsAskBattleCode-26901]
	_ = x[lbsAskMcsAddress-26902]
	_ = x[lbsAskMcsVersion-26903]
	_ = x[lbsMatchingCancel-24581]
	_ = x[lbsExtSyncSharedData-39168]
	_ = x[lbsPlatformInfo-39248]
	_ = x[lbsGamePatch-39264]
	_ = x[lbsP2PMatching-39265]
	_ = x[lbsP2PMatchingReport-39266]
	_ = x[lbsBattleUserCount-39269]
}

const _CmdID_name = "lbsLineChecklbsLogoutlbsShutDownlbsVSUserLostlbsMatchingCancellbsConnectionIDlbsAskConnectionIDlbsWarningMessagelbsLoginTypelbsUserHandlelbsUserRegistlbsUserDecidelbsAskPlatformCodelbsAskCountryCodelbsAskGameCodelbsAskGameVersionlbsLoginOklbsAskBattleResultlbsUserInfo1lbsUserInfo2lbsUserInfo3lbsUserInfo4lbsUserInfo5lbsUserInfo6lbsUserInfo7lbsUserInfo8lbsUserInfo9lbsEncodeStartlbsStartLobbylbsAskKDDIChargeslbsPostGameParameterlbsRankRankinglbsWinLoselbsDeviceDatalbsServerMoneylbsPlazaMaxlbsPlazaTitlelbsPlazaJoinlbsPlazaStatuslbsPlazaEntrylbsGoToToplbsPlazaExplainlbsLobbyJoinlbsLobbyEntrylbsPlazaExitlbsRoomMaxlbsRoomTitlelbsRoomStatuslbsRoomEntrylbsRoomCreatelbsLobbyExitlbsLobbyMatchingEntrylbsLobbyMatchingJoinlbsLobbyRemovelbsRoomExitlbsRoomLeaverlbsRoomCommerlbsMatchingEntrylbsRoomRemovelbsWaitJoinlbsRoomUserRejectlbsPutRoomNamelbsEndRoomCreatelbsPostChatMessagelbsChatMessagelbsUserSitelbsSendMaillbsRecvMaillbsManagerMessagelbsAskNewsTaglbsNewsTextlbsInvitationTaglbsRegulationHeaderlbsRegulationTextlbsRegulationFooterlbsTopRankingTaglbsTopRankingSuulbsTopRankinglbsAskPatchDatalbsPatchHeaderlbsPatchData6863lbsCalcDownloadChecksumlbsPatchPinglbsReadyBattlelbsAskMatchingJoinlbsAskPlayerSidelbsAskPlayerInfolbsAskRuleDatalbsAskBattleCodelbsAskMcsAddresslbsAskMcsVersionlbsExtSyncSharedDatalbsPlatformInfolbsGamePatchlbsP2PMatchinglbsP2PMatchingReportlbsBattleUserCount"

var _CmdID_map = map[CmdID]string{
	24577: _CmdID_name[0:12],
	24578: _CmdID_name[12:21],
	24579: _CmdID_name[21:32],
	24580: _CmdID_name[32:45],
	24581: _CmdID_name[45:62],
	24833: _CmdID_name[62:77],
	24834: _CmdID_name[77:95],
	24835: _CmdID_name[95:112],
	24848: _CmdID_name[112:124],
	24849: _CmdID_name[124:137],
	24850: _CmdID_name[137:150],
	24851: _CmdID_name[150:163],
	24852: _CmdID_name[163:181],
	24853: _CmdID_name[181:198],
	24854: _CmdID_name[198:212],
	24855: _CmdID_name[212:229],
	24856: _CmdID_name[229:239],
	24864: _CmdID_name[239:257],
	24881: _CmdID_name[257:269],
	24882: _CmdID_name[269:281],
	24883: _CmdID_name[281:293],
	24884: _CmdID_name[293:305],
	24885: _CmdID_name[305:317],
	24886: _CmdID_name[317:329],
	24887: _CmdID_name[329:341],
	24888: _CmdID_name[341:353],
	24889: _CmdID_name[353:365],
	24890: _CmdID_name[365:379],
	24897: _CmdID_name[379:392],
	24898: _CmdID_name[392:409],
	24899: _CmdID_name[409:429],
	24900: _CmdID_name[429:443],
	24901: _CmdID_name[443:453],
	24904: _CmdID_name[453:466],
	24905: _CmdID_name[466:480],
	25091: _CmdID_name[480:491],
	25092: _CmdID_name[491:504],
	25093: _CmdID_name[504:516],
	25094: _CmdID_name[516:530],
	25095: _CmdID_name[530:543],
	25096: _CmdID_name[543:553],
	25098: _CmdID_name[553:568],
	25347: _CmdID_name[568:580],
	25349: _CmdID_name[580:593],
	25350: _CmdID_name[593:605],
	25601: _CmdID_name[605:615],
	25602: _CmdID_name[615:627],
	25604: _CmdID_name[627:640],
	25606: _CmdID_name[640:652],
	25607: _CmdID_name[652:665],
	25608: _CmdID_name[665:677],
	25614: _CmdID_name[677:698],
	25615: _CmdID_name[698:718],
	25792: _CmdID_name[718:732],
	25857: _CmdID_name[732:743],
	25858: _CmdID_name[743:756],
	25859: _CmdID_name[756:769],
	25860: _CmdID_name[769:785],
	25861: _CmdID_name[785:798],
	25862: _CmdID_name[798:809],
	25863: _CmdID_name[809:826],
	26121: _CmdID_name[826:840],
	26124: _CmdID_name[840:856],
	26369: _CmdID_name[856:874],
	26370: _CmdID_name[874:888],
	26371: _CmdID_name[888:899],
	26372: _CmdID_name[899:910],
	26373: _CmdID_name[910:921],
	26374: _CmdID_name[921:938],
	26625: _CmdID_name[938:951],
	26626: _CmdID_name[951:962],
	26640: _CmdID_name[962:978],
	26656: _CmdID_name[978:997],
	26657: _CmdID_name[997:1014],
	26658: _CmdID_name[1014:1033],
	26705: _CmdID_name[1033:1049],
	26706: _CmdID_name[1049:1065],
	26707: _CmdID_name[1065:1078],
	26721: _CmdID_name[1078:1093],
	26722: _CmdID_name[1093:1107],
	26723: _CmdID_name[1107:1123],
	26724: _CmdID_name[1123:1146],
	26725: _CmdID_name[1146:1158],
	26896: _CmdID_name[1158:1172],
	26897: _CmdID_name[1172:1190],
	26898: _CmdID_name[1190:1206],
	26899: _CmdID_name[1206:1222],
	26900: _CmdID_name[1222:1236],
	26901: _CmdID_name[1236:1252],
	26902: _CmdID_name[1252:1268],
	26903: _CmdID_name[1268:1284],
	39168: _CmdID_name[1284:1304],
	39248: _CmdID_name[1304:1319],
	39264: _CmdID_name[1319:1331],
	39265: _CmdID_name[1331:1345],
	39266: _CmdID_name[1345:1365],
	39269: _CmdID_name[1365:1383],
}

func (i CmdID) String() string {
	if str, ok := _CmdID_map[i]; ok {
		return str
	}
	return "CmdID(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServerToClient-24]
	_ = x[ClientToServer-129]
}

const (
	_CmdDirection_name_0 = "ServerToClient"
	_CmdDirection_name_1 = "ClientToServer"
)

func (i CmdDirection) String() string {
	switch {
	case i == 24:
		return _CmdDirection_name_0
	case i == 129:
		return _CmdDirection_name_1
	default:
		return "CmdDirection(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CategoryQuestion-1]
	_ = x[CategoryAnswer-2]
	_ = x[CategoryNotice-16]
	_ = x[CategoryCustom-255]
}

const (
	_CmdCategory_name_0 = "CategoryQuestionCategoryAnswer"
	_CmdCategory_name_1 = "CategoryNotice"
	_CmdCategory_name_2 = "CategoryCustom"
)

var (
	_CmdCategory_index_0 = [...]uint8{0, 16, 30}
)

func (i CmdCategory) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _CmdCategory_name_0[_CmdCategory_index_0[i]:_CmdCategory_index_0[i+1]]
	case i == 16:
		return _CmdCategory_name_1
	case i == 255:
		return _CmdCategory_name_2
	default:
		return "CmdCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusError-4294967295]
	_ = x[StatusSuccess-16777215]
}

const (
	_CmdStatus_name_0 = "StatusSuccess"
	_CmdStatus_name_1 = "StatusError"
)

func (i CmdStatus) String() string {
	switch {
	case i == 16777215:
		return _CmdStatus_name_0
	case i == 4294967295:
		return _CmdStatus_name_1
	default:
		return "CmdStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
