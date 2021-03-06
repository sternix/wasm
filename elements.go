// +build js,wasm

package wasm

import (
	"time"
)

type (
	// https://www.w3.org/TR/html52/sections.html#htmlbodyelement
	HTMLBodyElement interface {
		HTMLElement
		WindowEventHandlers
	}

	// https://www.w3.org/TR/html52/sections.html#htmlheadingelement
	HTMLHeadingElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmlparagraphelement
	HTMLParagraphElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmlhrelement
	HTMLHRElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmlpreelement
	HTMLPreElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmlquoteelement
	HTMLQuoteElement interface {
		HTMLElement

		Cite() string
		SetCite(string)
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmlolistelement
	HTMLOListElement interface {
		HTMLElement

		Reversed() bool
		SetReversed(bool)
		Start() int
		SetStart(int)
		Type() string
		SetType(string)
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmlulistelement
	HTMLUListElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmllielement
	HTMLLIElement interface {
		HTMLElement

		Value() int
		SetValue(int)
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmldlistelement
	HTMLDListElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/grouping-content.html#htmldivelement
	HTMLDivElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/textlevel-semantics.html#htmlanchorelement
	HTMLAnchorElement interface {
		HTMLElement
		HTMLHyperlinkElementUtils

		Target() string
		SetTarget(string)
		Download() string
		SetDownload(string)
		Rel() string
		SetRel(string)
		Rev() string
		SetRev(string)
		RelList() DOMTokenList
		HrefLang() string
		SetHrefLang(string)
		Type() string
		SetType(string)
		Text() string
		SetText(string)
		ReferrerPolicy() string
		SetReferrerPolicy(string)
	}

	// https://www.w3.org/TR/html52/links.html#htmlhyperlinkelementutils
	HTMLHyperlinkElementUtils interface {
		Href() string
		SetHref(string)
		Origin() string
		Protocol() string
		SetProtocol(string)
		Username() string
		SetUsername(string)
		Password() string
		SetPassword(string)
		Host() string
		SetHost(string)
		Hostname() string
		SetHostname(string)
		Port() string
		SetPort(string)
		Pathname() string
		SetPathname(string)
		Search() string
		SetSearch(string)
		Hash() string
		SetHash(string)
	}

	// https://www.w3.org/TR/html52/textlevel-semantics.html#htmldataelement
	HTMLDataElement interface {
		HTMLElement

		Value() string
		SetValue(string)
	}

	// https://www.w3.org/TR/html52/textlevel-semantics.html#htmltimeelement
	HTMLTimeElement interface {
		HTMLElement

		DateTime() string
		SetDateTime(string)
	}

	// https://www.w3.org/TR/html52/textlevel-semantics.html#htmlspanelement
	HTMLSpanElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/textlevel-semantics.html#htmlbrelement
	HTMLBRElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/edits.html#htmlmodelement
	HTMLModElement interface {
		HTMLElement

		Cite() string
		SetCite(string)
		DateTime() string
		SetDateTime(string)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlpictureelement
	HTMLPictureElement interface {
		HTMLElement
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlsourceelement
	HTMLSourceElement interface {
		HTMLElement

		Src() string
		SetSrc(string)
		Type() string
		SetType(string)
		SrcSet() string
		SetSrcSet(string)
		Sizes() string
		SetSizes(string)
		Media() string
		SetMedia(string)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlimageelement
	HTMLImageElement interface {
		HTMLElement

		Alt() string
		SetAlt(string)
		Src() string
		SetSrc(string)
		SrcSet() string
		SetSrcSet(string)
		Sizes() string
		SetSizes(string)
		CrossOrigin() string
		SetCrossOrigin(string)
		UseMap() string
		SetUseMap(string)
		LongDesc() string
		SetLongDesc(string)
		IsMap() bool
		SetIsMap(bool)
		Width() uint
		SetWidth(uint)
		Height() uint
		SetHeight(uint)
		NaturalWidth() uint
		NaturalHeight() uint
		Complete() bool
		CurrentSrc() string
		ReferrerPolicy() string
		SetReferrerPolicy(string)

		X() int
		Y() int
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmliframeelement
	HTMLIFrameElement interface {
		HTMLElement

		Src() string
		SetSrc(string)
		SrcDoc() string
		SetSrcDoc(string)
		Name() string
		SetName(string)
		Sandbox() DOMTokenList
		AllowFullScreen() bool
		SetAllowFullScreen(bool)
		AllowPaymentRequest() bool
		SetAllowPaymentRequest(bool)
		Width() string
		SetWidth(string)
		Height() string
		SetHeight(string)
		ReferrerPolicy() string
		SetReferrerPolicy(string)
		ContentDocument() Document
		ContentWindow() WindowProxy
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlembedelement
	HTMLEmbedElement interface {
		HTMLElement

		Src() string
		SetSrc(string)
		Type() string
		SetType(string)
		Width() string
		SetWidth(string)
		Height() string
		SetHeight(string)
		// legacycaller any (any... arguments) ????
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlobjectelement
	HTMLObjectElement interface {
		HTMLElement

		Data() string
		SetData(string)
		Type() string
		SetType(string)
		TypeMustMatch() bool
		SetTypeMustMatch(bool)
		Name() string
		SetName(string)
		Form() HTMLFormElement
		Width() string
		SetWidth(string)
		Height() string
		SetHeight(string)
		ContentDocument() Document
		ContentWindow() WindowProxy
		WillValidate() bool
		Validity() ValidityState
		ValidationMessage() string
		CheckValidity() bool
		ReportValidity() bool
		SetCustomValidity(string)
	}

	// https://www.w3.org/TR/html52/sec-forms.html#validitystate
	ValidityState interface {
		ValueMissing() bool
		TypeMismatch() bool
		PatternMismatch() bool
		TooLong() bool
		TooShort() bool
		RangeUnderflow() bool
		RangeOverflow() bool
		StepMismatch() bool
		BadInput() bool
		CustomError() bool
		Valid() bool
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlparamelement
	HTMLParamElement interface {
		HTMLElement

		Name() string
		SetName(string)
		Value() string
		SetValue(string)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlvideoelement
	HTMLVideoElement interface {
		HTMLMediaElement

		Width() uint
		SetWidth(uint)
		Height() uint
		SetHeight(uint)
		VideoWidth() uint
		VideoHeight() uint
		Poster() string
		SetPoster(string)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlaudioelement
	HTMLAudioElement interface {
		HTMLMediaElement
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmltrackelement
	HTMLTrackElement interface {
		HTMLElement

		Kind() string
		SetKind(string)
		Src() string
		SetSrc(string)
		SrcLang() string
		SetSrcLang(string)
		Label() string
		SetLabel(string)
		Default() bool
		SetDefault(bool)
		ReadyState() HTMLTrackElementReadyState
		Track() TextTrack
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#texttrack
	TextTrack interface {
		EventTarget

		Kind() TextTrackKind
		Label() string
		Language() string
		Id() string
		InBandMetadataTrackDispatchType() string
		Mode() TextTrackMode
		SetMode(TextTrackMode)
		Cues() TextTrackCueList
		ActiveCues() TextTrackCueList
		AddCue(TextTrackCue)
		RemoveCue(TextTrackCue)
		OnCueChange(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#texttrackcuelist
	TextTrackCueList interface {
		Length() uint
		Item(uint) TextTrackCue
		CueById(string) TextTrackCue
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#texttrackcue
	TextTrackCue interface {
		EventTarget

		Track() TextTrack
		Id() string
		SetId(string)
		StartTime() float64
		SetStartTime(float64)
		EndTime() float64
		SetEndTime(float64)
		PauseOnExit() bool
		SetPauseOnExit(bool)
		OnEnter(func(Event)) EventHandler
		OnExit(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlmapelement
	HTMLMapElement interface {
		HTMLElement

		Name() string
		SetName(string)
		Areas() []HTMLAreaElement
		Images() []HTMLElement
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlareaelement
	HTMLAreaElement interface {
		HTMLElement
		HTMLHyperlinkElementUtils

		Alt() string
		SetAlt(string)
		Coords() string
		SetCoords(string)
		Shape() string
		SetShape(string)
		Target() string
		SetTarget(string)
		Download() string
		SetDownload(string)
		Rel() string
		SetRel(string)
		RelList() DOMTokenList
		HrefLang() string
		SetHrefLang(string)
		Type() string
		SetType(string)
		ReferrerPolicy() string
		SetReferrerPolicy(string)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlmediaelement
	HTMLMediaElement interface {
		HTMLElement

		Error() MediaError
		Src() string
		SetSrc(string)
		SrcObject() MediaProvider
		SetSrcObject(MediaProvider)
		CurrentSrc() string
		CrossOrigin() string
		SetCrossOrigin(string)
		NetworkState() MediaNetworkState
		Preload() string
		SetPreload(string)
		Buffered() TimeRanges
		Load()
		CanPlayType(string) CanPlayTypeResult
		ReadyState() MediaReadyState
		Seeking() bool
		CurrentTime() float64
		SetCurrentTime(float64)
		FastSeek(float64)
		Duration() float64
		StartDate() time.Time // js Date
		Paused() bool
		DefaultPlaybackRate() float64
		SetDefaultPlaybackRate(float64)
		PlaybackRate() float64
		SetPlaybackRate(float64)
		Played() TimeRanges
		Seekable() TimeRanges
		Ended() bool
		AutoPlay() bool
		SetAutoPlay(bool)
		Loop() bool
		SetLoop(bool)
		Play()
		Pause()
		Controls() bool
		SetControls(bool)
		Volume() bool
		SetVolume(bool)
		Muted() bool
		SetMuted(bool)
		DefaultMuted() bool
		SetDefaultMuted(bool)
		AudioTracks() AudioTrackList
		VideoTracks() VideoTrackList
		TextTracks() TextTrackList
		AddTextTrack(TextTrackKind, ...string) TextTrack
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#audiotracklist
	AudioTrackList interface {
		EventTarget

		Length() uint
		Item(uint) AudioTrack
		TrackById(string) AudioTrack
		OnChange(func(Event)) EventHandler
		OnAddTrack(func(Event)) EventHandler
		OnRemoveTrack(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#audiotrack
	AudioTrack interface {
		Id() string
		Kind() string
		Label() string
		Language() string
		Enabled() bool
		SetEnabled(bool)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#videotracklist
	VideoTrackList interface {
		EventTarget

		Length() uint
		Item(uint) VideoTrack
		TrackById(string) VideoTrack
		SelectedIndex() int
		OnChange(func(Event)) EventHandler
		OnAddTrack(func(Event)) EventHandler
		OnRemoveTrack(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#videotrack
	VideoTrack interface {
		Id() string
		Kind() string
		Label() string
		Language() string
		Selected() bool
		SetSelected(bool)
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#texttracklist
	TextTrackList interface {
		EventTarget

		Length() uint
		Item(uint) TextTrack
		TrackById(string) TextTrack
		OnChange(func(Event)) EventHandler
		OnAddTrack(func(Event)) EventHandler
		OnRemoveTrack(func(Event)) EventHandler
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#timeranges
	TimeRanges interface {
		Length() uint
		Start(uint) float64
		End(uint) float64
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#mediaerror
	MediaError interface {
		Code() MediaErrorCode
	}

	// https://www.w3.org/TR/html52/semantics-embedded-content.html#typedefdef-mediaprovider
	// typedef (MediaStream or MediaSource or Blob) MediaProvider;
	MediaProvider interface {
		JSValue() jsValue
	}

	// https://www.w3.org/TR/media-source/#idl-def-mediasource
	MediaSource interface {
		EventTarget

		SourceBuffers() SourceBufferList
		ActiveSourceBuffers() SourceBufferList
		ReadyState() ReadyState
		Duration() float64
		SetDuration(float64)
		OnSourceOpen(func(Event)) EventHandler
		OnSourceEnded(func(Event)) EventHandler
		OnSourceClose(func(Event)) EventHandler
		AddSourceBuffer(string) SourceBuffer
		RemoveSourceBuffer(SourceBuffer)
		EndOfStream(...EndOfStreamError)
		SetLiveSeekableRange(float64, float64)
		ClearLiveSeekableRange()
		IsTypeSupported(string) bool // static
	}

	// https://www.w3.org/TR/media-source/#idl-def-sourcebufferlist
	SourceBufferList interface {
		Length() uint
		OnAddSourceBuffer(func(Event)) EventHandler
		OnRemoveSourceBuffer(func(Event)) EventHandler
		Item(uint) SourceBuffer // getter
	}

	// https://www.w3.org/TR/media-source/#idl-def-sourcebuffer
	SourceBuffer interface {
		EventTarget

		Mode() AppendMode
		SetMode(AppendMode)
		Updating() bool
		Buffered() TimeRanges
		TimestampOffset() float64
		SetTimestampOffset(float64)
		AudioTracks() AudioTrackList
		VideoTracks() VideoTrackList
		TextTracks() TextTrackList
		AppendWindowStart() float64
		SetAppendWindowStart(float64)
		AppendWindowEnd() float64
		SetAppendWindowEnd(float64)
		OnUpdatestart(fn func(Event)) EventHandler
		OnUpdate(fn func(Event)) EventHandler
		OnUpdateend(fn func(Event)) EventHandler
		OnError(fn func(Event)) EventHandler
		OnAbort(fn func(Event)) EventHandler
		AppendBuffer(BufferSource)
		Abort()
		Remove(float64, float64)
	}
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlmediaelement
type MediaReadyState uint16

const (
	MediaReadyStateHaveNothing     MediaReadyState = 0
	MediaReadyStateHaveMetadata    MediaReadyState = 1
	MediaReadyStateHaveCurrentData MediaReadyState = 2
	MediaReadyStateHaveFutureData  MediaReadyState = 3
	MediaReadyStateHaveEnoughData  MediaReadyState = 4
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmlmediaelement
type MediaNetworkState uint16

const (
	MediaNetworkStateEmpty    MediaNetworkState = 0
	MediaNetworkStateIdLe     MediaNetworkState = 1
	MediaNetworkStateLoading  MediaNetworkState = 2
	MediaNetworkStateNoSource MediaNetworkState = 3
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#mediaerror
type MediaErrorCode uint16

const (
	MediaErrorCodeAborted         MediaErrorCode = 1
	MediaErrorCodeNetwork         MediaErrorCode = 2
	MediaErrorCodeDecode          MediaErrorCode = 3
	MediaErrorCodeSrcNotSupported MediaErrorCode = 4
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#enumdef-canplaytyperesult
type CanPlayTypeResult string

const (
	CanPlayTypeResultEmpty    CanPlayTypeResult = ""
	CanPlayTypeResultMaybe    CanPlayTypeResult = "maybe"
	CanPlayTypeResultProbably CanPlayTypeResult = "probably"
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#htmltrackelement
type HTMLTrackElementReadyState uint16

const (
	HTMLTrackElementReadyStateNone    HTMLTrackElementReadyState = 0
	HTMLTrackElementReadyStateLoading HTMLTrackElementReadyState = 1
	HTMLTrackElementReadyStateLoaded  HTMLTrackElementReadyState = 2
	HTMLTrackElementReadyStateError   HTMLTrackElementReadyState = 3
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#enumdef-texttrackkind
type TextTrackKind string

const (
	TextTrackKindSubtitles    TextTrackKind = "subtitles"
	TextTrackKindCaptions     TextTrackKind = "captions"
	TextTrackKindDescriptions TextTrackKind = "descriptions"
	TextTrackKindChapters     TextTrackKind = "chapters"
	TextTrackKindMetada       TextTrackKind = "metadata"
)

// https://www.w3.org/TR/html52/semantics-embedded-content.html#enumdef-texttrackmode
type TextTrackMode string

const (
	TextTrackModeDisabled TextTrackMode = "disabled"
	TextTrackModeHidden   TextTrackMode = "hidden"
	TextTrackModeShowing  TextTrackMode = "showing"
)

// https://www.w3.org/TR/media-source/#idl-def-readystate
type ReadyState string

const (
	ReadyStateClosed ReadyState = "closed"
	ReadyStateOpen   ReadyState = "open"
	ReadyStateEnded  ReadyState = "ended"
)

// https://www.w3.org/TR/media-source/#idl-def-appendmode
type AppendMode string

const (
	AppendModeSegments AppendMode = "segments"
	AppendModeSequence AppendMode = "sequence"
)

// -------------8<---------------------------------------

// https://www.w3.org/TR/media-source/#idl-def-endofstreamerror
type EndOfStreamError string

const (
	EndOfStreamErrorNetwork EndOfStreamError = "network"
	EndOfStreamErrorDecode  EndOfStreamError = "decode"
)

// -------------8<---------------------------------------
