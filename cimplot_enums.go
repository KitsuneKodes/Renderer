// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// original name: ImAxis_
type PlotAxisEnum int

const (
	AxisX1    = 0
	AxisX2    = 1
	AxisX3    = 2
	AxisY1    = 3
	AxisY2    = 4
	AxisY3    = 5
	AxisCOUNT = 6
)

// original name: ImPlotAxisFlags_
type PlotAxisFlags int

const (
	PlotAxisFlagsNone          = 0
	PlotAxisFlagsNoLabel       = 1
	PlotAxisFlagsNoGridLines   = 2
	PlotAxisFlagsNoTickMarks   = 4
	PlotAxisFlagsNoTickLabels  = 8
	PlotAxisFlagsNoInitialFit  = 16
	PlotAxisFlagsNoMenus       = 32
	PlotAxisFlagsNoSideSwitch  = 64
	PlotAxisFlagsNoHighlight   = 128
	PlotAxisFlagsOpposite      = 256
	PlotAxisFlagsForeground    = 512
	PlotAxisFlagsInvert        = 1024
	PlotAxisFlagsAutoFit       = 2048
	PlotAxisFlagsRangeFit      = 4096
	PlotAxisFlagsPanStretch    = 8192
	PlotAxisFlagsLockMin       = 16384
	PlotAxisFlagsLockMax       = 32768
	PlotAxisFlagsLock          = 49152
	PlotAxisFlagsNoDecorations = 15
	PlotAxisFlagsAuxDefault    = 258
)

// original name: ImPlotBarGroupsFlags_
type PlotBarGroupsFlags int

const (
	PlotBarGroupsFlagsNone       = 0
	PlotBarGroupsFlagsHorizontal = 1024
	PlotBarGroupsFlagsStacked    = 2048
)

// original name: ImPlotBarsFlags_
type PlotBarsFlags int

const (
	PlotBarsFlagsNone       = 0
	PlotBarsFlagsHorizontal = 1024
)

// original name: ImPlotBin_
type PlotBin int

const (
	PlotBinSqrt    = -1
	PlotBinSturges = -2
	PlotBinRice    = -3
	PlotBinScott   = -4
)

// original name: ImPlotCol_
type PlotCol int

const (
	PlotColLine          = 0
	PlotColFill          = 1
	PlotColMarkerOutline = 2
	PlotColMarkerFill    = 3
	PlotColErrorBar      = 4
	PlotColFrameBg       = 5
	PlotColPlotBg        = 6
	PlotColPlotBorder    = 7
	PlotColLegendBg      = 8
	PlotColLegendBorder  = 9
	PlotColLegendText    = 10
	PlotColTitleText     = 11
	PlotColInlayText     = 12
	PlotColAxisText      = 13
	PlotColAxisGrid      = 14
	PlotColAxisTick      = 15
	PlotColAxisBg        = 16
	PlotColAxisBgHovered = 17
	PlotColAxisBgActive  = 18
	PlotColSelection     = 19
	PlotColCrosshairs    = 20
	PlotColCOUNT         = 21
)

// original name: ImPlotColormapScaleFlags_
type PlotColormapScaleFlags int

const (
	PlotColormapScaleFlagsNone     = 0
	PlotColormapScaleFlagsNoLabel  = 1
	PlotColormapScaleFlagsOpposite = 2
	PlotColormapScaleFlagsInvert   = 4
)

// original name: ImPlotColormap_
type PlotColormap int

const (
	PlotColormapDeep     = 0
	PlotColormapDark     = 1
	PlotColormapPastel   = 2
	PlotColormapPaired   = 3
	PlotColormapViridis  = 4
	PlotColormapPlasma   = 5
	PlotColormapHot      = 6
	PlotColormapCool     = 7
	PlotColormapPink     = 8
	PlotColormapJet      = 9
	PlotColormapTwilight = 10
	PlotColormapRdBu     = 11
	PlotColormapBrBG     = 12
	PlotColormapPiYG     = 13
	PlotColormapSpectral = 14
	PlotColormapGreys    = 15
)

// original name: ImPlotCond_
type PlotCond int

const (
	PlotCondNone   = 0
	PlotCondAlways = 1
	PlotCondOnce   = 2
)

// original name: ImPlotDateFmt_
type PlotDateFmt int

const (
	PlotDateFmtNone    = 0
	PlotDateFmtDayMo   = 1
	PlotDateFmtDayMoYr = 2
	PlotDateFmtMoYr    = 3
	PlotDateFmtMo      = 4
	PlotDateFmtYr      = 5
)

// original name: ImPlotDigitalFlags_
type PlotDigitalFlags int

const (
	PlotDigitalFlagsNone = 0
)

// original name: ImPlotDragToolFlags_
type PlotDragToolFlags int

const (
	PlotDragToolFlagsNone      = 0
	PlotDragToolFlagsNoCursors = 1
	PlotDragToolFlagsNoFit     = 2
	PlotDragToolFlagsNoInputs  = 4
	PlotDragToolFlagsDelayed   = 8
)

// original name: ImPlotDummyFlags_
type PlotDummyFlags int

const (
	PlotDummyFlagsNone = 0
)

// original name: ImPlotErrorBarsFlags_
type PlotErrorBarsFlags int

const (
	PlotErrorBarsFlagsNone       = 0
	PlotErrorBarsFlagsHorizontal = 1024
)

// original name: ImPlotFlags_
type PlotFlags int

const (
	PlotFlagsNone        = 0
	PlotFlagsNoTitle     = 1
	PlotFlagsNoLegend    = 2
	PlotFlagsNoMouseText = 4
	PlotFlagsNoInputs    = 8
	PlotFlagsNoMenus     = 16
	PlotFlagsNoBoxSelect = 32
	PlotFlagsNoChild     = 64
	PlotFlagsNoFrame     = 128
	PlotFlagsEqual       = 256
	PlotFlagsCrosshairs  = 512
	PlotFlagsCanvasOnly  = 55
)

// original name: ImPlotHeatmapFlags_
type PlotHeatmapFlags int

const (
	PlotHeatmapFlagsNone     = 0
	PlotHeatmapFlagsColMajor = 1024
)

// original name: ImPlotHistogramFlags_
type PlotHistogramFlags int

const (
	PlotHistogramFlagsNone       = 0
	PlotHistogramFlagsHorizontal = 1024
	PlotHistogramFlagsCumulative = 2048
	PlotHistogramFlagsDensity    = 4096
	PlotHistogramFlagsNoOutliers = 8192
	PlotHistogramFlagsColMajor   = 16384
)

// original name: ImPlotImageFlags_
type PlotImageFlags int

const (
	PlotImageFlagsNone = 0
)

// original name: ImPlotInfLinesFlags_
type PlotInfLinesFlags int

const (
	PlotInfLinesFlagsNone       = 0
	PlotInfLinesFlagsHorizontal = 1024
)

// original name: ImPlotItemFlags_
type PlotItemFlags int

const (
	PlotItemFlagsNone     = 0
	PlotItemFlagsNoLegend = 1
	PlotItemFlagsNoFit    = 2
)

// original name: ImPlotLegendFlags_
type PlotLegendFlags int

const (
	PlotLegendFlagsNone            = 0
	PlotLegendFlagsNoButtons       = 1
	PlotLegendFlagsNoHighlightItem = 2
	PlotLegendFlagsNoHighlightAxis = 4
	PlotLegendFlagsNoMenus         = 8
	PlotLegendFlagsOutside         = 16
	PlotLegendFlagsHorizontal      = 32
	PlotLegendFlagsSort            = 64
)

// original name: ImPlotLineFlags_
type PlotLineFlags int

const (
	PlotLineFlagsNone     = 0
	PlotLineFlagsSegments = 1024
	PlotLineFlagsLoop     = 2048
	PlotLineFlagsSkipNaN  = 4096
	PlotLineFlagsNoClip   = 8192
	PlotLineFlagsShaded   = 16384
)

// original name: ImPlotLocation_
type PlotLocation int

const (
	PlotLocationCenter    = 0
	PlotLocationNorth     = 1
	PlotLocationSouth     = 2
	PlotLocationWest      = 4
	PlotLocationEast      = 8
	PlotLocationNorthWest = 5
	PlotLocationNorthEast = 9
	PlotLocationSouthWest = 6
	PlotLocationSouthEast = 10
)

// original name: ImPlotMarker_
type PlotMarker int

const (
	PlotMarkerNone     = -1
	PlotMarkerCircle   = 0
	PlotMarkerSquare   = 1
	PlotMarkerDiamond  = 2
	PlotMarkerUp       = 3
	PlotMarkerDown     = 4
	PlotMarkerLeft     = 5
	PlotMarkerRight    = 6
	PlotMarkerCross    = 7
	PlotMarkerPlus     = 8
	PlotMarkerAsterisk = 9
	PlotMarkerCOUNT    = 10
)

// original name: ImPlotMouseTextFlags_
type PlotMouseTextFlags int

const (
	PlotMouseTextFlagsNone       = 0
	PlotMouseTextFlagsNoAuxAxes  = 1
	PlotMouseTextFlagsNoFormat   = 2
	PlotMouseTextFlagsShowAlways = 4
)

// original name: ImPlotPieChartFlags_
type PlotPieChartFlags int

const (
	PlotPieChartFlagsNone      = 0
	PlotPieChartFlagsNormalize = 1024
)

// original name: ImPlotScale_
type PlotScale int

const (
	PlotScaleLinear = 0
	PlotScaleTime   = 1
	PlotScaleLog10  = 2
	PlotScaleSymLog = 3
)

// original name: ImPlotScatterFlags_
type PlotScatterFlags int

const (
	PlotScatterFlagsNone   = 0
	PlotScatterFlagsNoClip = 1024
)

// original name: ImPlotShadedFlags_
type PlotShadedFlags int

const (
	PlotShadedFlagsNone = 0
)

// original name: ImPlotStairsFlags_
type PlotStairsFlags int

const (
	PlotStairsFlagsNone    = 0
	PlotStairsFlagsPreStep = 1024
	PlotStairsFlagsShaded  = 2048
)

// original name: ImPlotStemsFlags_
type PlotStemsFlags int

const (
	PlotStemsFlagsNone       = 0
	PlotStemsFlagsHorizontal = 1024
)

// original name: ImPlotStyleVar_
type PlotStyleVar int

const (
	PlotStyleVarLineWeight         = 0
	PlotStyleVarMarker             = 1
	PlotStyleVarMarkerSize         = 2
	PlotStyleVarMarkerWeight       = 3
	PlotStyleVarFillAlpha          = 4
	PlotStyleVarErrorBarSize       = 5
	PlotStyleVarErrorBarWeight     = 6
	PlotStyleVarDigitalBitHeight   = 7
	PlotStyleVarDigitalBitGap      = 8
	PlotStyleVarPlotBorderSize     = 9
	PlotStyleVarMinorAlpha         = 10
	PlotStyleVarMajorTickLen       = 11
	PlotStyleVarMinorTickLen       = 12
	PlotStyleVarMajorTickSize      = 13
	PlotStyleVarMinorTickSize      = 14
	PlotStyleVarMajorGridSize      = 15
	PlotStyleVarMinorGridSize      = 16
	PlotStyleVarPlotPadding        = 17
	PlotStyleVarLabelPadding       = 18
	PlotStyleVarLegendPadding      = 19
	PlotStyleVarLegendInnerPadding = 20
	PlotStyleVarLegendSpacing      = 21
	PlotStyleVarMousePosPadding    = 22
	PlotStyleVarAnnotationPadding  = 23
	PlotStyleVarFitPadding         = 24
	PlotStyleVarPlotDefaultSize    = 25
	PlotStyleVarPlotMinSize        = 26
	PlotStyleVarCOUNT              = 27
)

// original name: ImPlotSubplotFlags_
type PlotSubplotFlags int

const (
	PlotSubplotFlagsNone       = 0
	PlotSubplotFlagsNoTitle    = 1
	PlotSubplotFlagsNoLegend   = 2
	PlotSubplotFlagsNoMenus    = 4
	PlotSubplotFlagsNoResize   = 8
	PlotSubplotFlagsNoAlign    = 16
	PlotSubplotFlagsShareItems = 32
	PlotSubplotFlagsLinkRows   = 64
	PlotSubplotFlagsLinkCols   = 128
	PlotSubplotFlagsLinkAllX   = 256
	PlotSubplotFlagsLinkAllY   = 512
	PlotSubplotFlagsColMajor   = 1024
)

// original name: ImPlotTextFlags_
type PlotTextFlags int

const (
	PlotTextFlagsNone     = 0
	PlotTextFlagsVertical = 1024
)

// original name: ImPlotTimeFmt_
type PlotTimeFmt int

const (
	PlotTimeFmtNone     = 0
	PlotTimeFmtUs       = 1
	PlotTimeFmtSUs      = 2
	PlotTimeFmtSMs      = 3
	PlotTimeFmtS        = 4
	PlotTimeFmtMinSMs   = 5
	PlotTimeFmtHrMinSMs = 6
	PlotTimeFmtHrMinS   = 7
	PlotTimeFmtHrMin    = 8
	PlotTimeFmtHr       = 9
)

// original name: ImPlotTimeUnit_
type PlotTimeUnit int

const (
	PlotTimeUnitUs    = 0
	PlotTimeUnitMs    = 1
	PlotTimeUnitS     = 2
	PlotTimeUnitMin   = 3
	PlotTimeUnitHr    = 4
	PlotTimeUnitDay   = 5
	PlotTimeUnitMo    = 6
	PlotTimeUnitYr    = 7
	PlotTimeUnitCOUNT = 8
)
