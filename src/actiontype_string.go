// Code generated by "stringer -type=actionType"; DO NOT EDIT.

package fzf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[actIgnore-0]
	_ = x[actStart-1]
	_ = x[actClick-2]
	_ = x[actInvalid-3]
	_ = x[actChar-4]
	_ = x[actMouse-5]
	_ = x[actBeginningOfLine-6]
	_ = x[actAbort-7]
	_ = x[actAccept-8]
	_ = x[actAcceptNonEmpty-9]
	_ = x[actAcceptOrPrintQuery-10]
	_ = x[actBackwardChar-11]
	_ = x[actBackwardDeleteChar-12]
	_ = x[actBackwardDeleteCharEof-13]
	_ = x[actBackwardWord-14]
	_ = x[actCancel-15]
	_ = x[actChangeBorderLabel-16]
	_ = x[actChangeListLabel-17]
	_ = x[actChangeInputLabel-18]
	_ = x[actChangeHeader-19]
	_ = x[actChangeHeaderLabel-20]
	_ = x[actChangeMulti-21]
	_ = x[actChangePreviewLabel-22]
	_ = x[actChangePrompt-23]
	_ = x[actChangeQuery-24]
	_ = x[actChangeNth-25]
	_ = x[actClearScreen-26]
	_ = x[actClearQuery-27]
	_ = x[actClearSelection-28]
	_ = x[actClose-29]
	_ = x[actDeleteChar-30]
	_ = x[actDeleteCharEof-31]
	_ = x[actEndOfLine-32]
	_ = x[actFatal-33]
	_ = x[actForwardChar-34]
	_ = x[actForwardWord-35]
	_ = x[actKillLine-36]
	_ = x[actKillWord-37]
	_ = x[actUnixLineDiscard-38]
	_ = x[actUnixWordRubout-39]
	_ = x[actYank-40]
	_ = x[actBackwardKillWord-41]
	_ = x[actSelectAll-42]
	_ = x[actDeselectAll-43]
	_ = x[actToggle-44]
	_ = x[actToggleSearch-45]
	_ = x[actToggleAll-46]
	_ = x[actToggleDown-47]
	_ = x[actToggleUp-48]
	_ = x[actToggleIn-49]
	_ = x[actToggleOut-50]
	_ = x[actToggleTrack-51]
	_ = x[actToggleTrackCurrent-52]
	_ = x[actToggleHeader-53]
	_ = x[actToggleWrap-54]
	_ = x[actToggleMultiLine-55]
	_ = x[actToggleHscroll-56]
	_ = x[actTrackCurrent-57]
	_ = x[actUntrackCurrent-58]
	_ = x[actDown-59]
	_ = x[actUp-60]
	_ = x[actPageUp-61]
	_ = x[actPageDown-62]
	_ = x[actPosition-63]
	_ = x[actHalfPageUp-64]
	_ = x[actHalfPageDown-65]
	_ = x[actOffsetUp-66]
	_ = x[actOffsetDown-67]
	_ = x[actOffsetMiddle-68]
	_ = x[actJump-69]
	_ = x[actJumpAccept-70]
	_ = x[actPrintQuery-71]
	_ = x[actRefreshPreview-72]
	_ = x[actReplaceQuery-73]
	_ = x[actToggleSort-74]
	_ = x[actShowPreview-75]
	_ = x[actHidePreview-76]
	_ = x[actTogglePreview-77]
	_ = x[actTogglePreviewWrap-78]
	_ = x[actTransform-79]
	_ = x[actTransformBorderLabel-80]
	_ = x[actTransformListLabel-81]
	_ = x[actTransformInputLabel-82]
	_ = x[actTransformHeader-83]
	_ = x[actTransformHeaderLabel-84]
	_ = x[actTransformPreviewLabel-85]
	_ = x[actTransformPrompt-86]
	_ = x[actTransformQuery-87]
	_ = x[actPreview-88]
	_ = x[actChangePreview-89]
	_ = x[actChangePreviewWindow-90]
	_ = x[actPreviewTop-91]
	_ = x[actPreviewBottom-92]
	_ = x[actPreviewUp-93]
	_ = x[actPreviewDown-94]
	_ = x[actPreviewPageUp-95]
	_ = x[actPreviewPageDown-96]
	_ = x[actPreviewHalfPageUp-97]
	_ = x[actPreviewHalfPageDown-98]
	_ = x[actPrevHistory-99]
	_ = x[actPrevSelected-100]
	_ = x[actPrint-101]
	_ = x[actPut-102]
	_ = x[actNextHistory-103]
	_ = x[actNextSelected-104]
	_ = x[actExecute-105]
	_ = x[actExecuteSilent-106]
	_ = x[actExecuteMulti-107]
	_ = x[actSigStop-108]
	_ = x[actFirst-109]
	_ = x[actLast-110]
	_ = x[actReload-111]
	_ = x[actReloadSync-112]
	_ = x[actDisableSearch-113]
	_ = x[actEnableSearch-114]
	_ = x[actSelect-115]
	_ = x[actDeselect-116]
	_ = x[actUnbind-117]
	_ = x[actRebind-118]
	_ = x[actBecome-119]
	_ = x[actShowHeader-120]
	_ = x[actHideHeader-121]
}

const _actionType_name = "actIgnoreactStartactClickactInvalidactCharactMouseactBeginningOfLineactAbortactAcceptactAcceptNonEmptyactAcceptOrPrintQueryactBackwardCharactBackwardDeleteCharactBackwardDeleteCharEofactBackwardWordactCancelactChangeBorderLabelactChangeListLabelactChangeInputLabelactChangeHeaderactChangeHeaderLabelactChangeMultiactChangePreviewLabelactChangePromptactChangeQueryactChangeNthactClearScreenactClearQueryactClearSelectionactCloseactDeleteCharactDeleteCharEofactEndOfLineactFatalactForwardCharactForwardWordactKillLineactKillWordactUnixLineDiscardactUnixWordRuboutactYankactBackwardKillWordactSelectAllactDeselectAllactToggleactToggleSearchactToggleAllactToggleDownactToggleUpactToggleInactToggleOutactToggleTrackactToggleTrackCurrentactToggleHeaderactToggleWrapactToggleMultiLineactToggleHscrollactTrackCurrentactUntrackCurrentactDownactUpactPageUpactPageDownactPositionactHalfPageUpactHalfPageDownactOffsetUpactOffsetDownactOffsetMiddleactJumpactJumpAcceptactPrintQueryactRefreshPreviewactReplaceQueryactToggleSortactShowPreviewactHidePreviewactTogglePreviewactTogglePreviewWrapactTransformactTransformBorderLabelactTransformListLabelactTransformInputLabelactTransformHeaderactTransformHeaderLabelactTransformPreviewLabelactTransformPromptactTransformQueryactPreviewactChangePreviewactChangePreviewWindowactPreviewTopactPreviewBottomactPreviewUpactPreviewDownactPreviewPageUpactPreviewPageDownactPreviewHalfPageUpactPreviewHalfPageDownactPrevHistoryactPrevSelectedactPrintactPutactNextHistoryactNextSelectedactExecuteactExecuteSilentactExecuteMultiactSigStopactFirstactLastactReloadactReloadSyncactDisableSearchactEnableSearchactSelectactDeselectactUnbindactRebindactBecomeactShowHeaderactHideHeader"

var _actionType_index = [...]uint16{0, 9, 17, 25, 35, 42, 50, 68, 76, 85, 102, 123, 138, 159, 183, 198, 207, 227, 245, 264, 279, 299, 313, 334, 349, 363, 375, 389, 402, 419, 427, 440, 456, 468, 476, 490, 504, 515, 526, 544, 561, 568, 587, 599, 613, 622, 637, 649, 662, 673, 684, 696, 710, 731, 746, 759, 777, 793, 808, 825, 832, 837, 846, 857, 868, 881, 896, 907, 920, 935, 942, 955, 968, 985, 1000, 1013, 1027, 1041, 1057, 1077, 1089, 1112, 1133, 1155, 1173, 1196, 1220, 1238, 1255, 1265, 1281, 1303, 1316, 1332, 1344, 1358, 1374, 1392, 1412, 1434, 1448, 1463, 1471, 1477, 1491, 1506, 1516, 1532, 1547, 1557, 1565, 1572, 1581, 1594, 1610, 1625, 1634, 1645, 1654, 1663, 1672, 1685, 1698}

func (i actionType) String() string {
	if i < 0 || i >= actionType(len(_actionType_index)-1) {
		return "actionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _actionType_name[_actionType_index[i]:_actionType_index[i+1]]
}
