// +build js,wasm

package wasm

/*
https://www.w3.org/TR/uievents-key/
Types for KeyboardEvent.key
*/

// Special Keys
// https://www.w3.org/TR/uievents-key/#keys-special
const (
	KeyUnidentified = "Unidentified"
)

// Modifier Keys
// https://www.w3.org/TR/uievents-key/#keys-modifier
const (
	KeyAlt        = "Alt"
	KeyAltGraph   = "AltGraph"
	KeyCapsLock   = "CapsLock"
	KeyControl    = "Control"
	KeyFn         = "Fn"
	KeyFnLock     = "FnLock"
	KeyMeta       = "Meta"
	KeyNumLock    = "NumLock"
	KeyScrollLock = "ScrollLock"
	KeyShift      = "Shift"
	KeySymbol     = "Symbol"
	KeySymbolLock = "SymbolLock"
)

//Whitespace Keys
// https://www.w3.org/TR/uievents-key/#keys-whitespace
const (
	KeyEnter = "Enter"
	KeyTab   = "Tab"
	KeySpace = " "
)

// Navigation Keys
// https://www.w3.org/TR/uievents-key/#keys-navigation
const (
	KeyArrowDown  = "ArrowDown"
	KeyArrowLeft  = "ArrowLeft"
	KeyArrowRight = "ArrowRight"
	KeyArrowUp    = "ArrowUp"
	KeyEnd        = "End"
	KeyHome       = "Home"
	KeyPageDown   = "PageDown"
	KeyPageUp     = "PageUp"
)

// Editing keys
// https://www.w3.org/TR/uievents-key/#keys-editing
const (
	KeyBackspace = "Backspace"
	KeyClear     = "Clear"
	KeyCopy      = "Copy"
	KeyCrSel     = "CrSel"
	KeyCut       = "Cut"
	KeyDelete    = "Delete"
	KeyEraseEof  = "EraseEof"
	KeyExSel     = "ExSel"
	KeyInsert    = "Insert"
	KeyPaste     = "Paste"
	KeyRedo      = "Redo"
	KeyUndo      = "Undo"
)

// UI Keys
// https://www.w3.org/TR/uievents-key/#keys-ui
const (
	KeyAccept      = "Accept"
	KeyAgain       = "Again"
	KeyAttn        = "Attn"
	KeyCancel      = "Cancel"
	KeyContextMenu = "ContextMenu"
	KeyEscape      = "Escape"
	KeyExecute     = "Execute"
	KeyFind        = "Find"
	KeyHelp        = "Help"
	KeyPause       = "Pause"
	KeyPlay        = "Play"
	KeyProps       = "Props"
	KeySelect      = "Select"
	KeyZoomIn      = "ZoomIn"
	KeyZoomOut     = "ZoomOut"
)

// Device Keys
// https://www.w3.org/TR/uievents-key/#keys-device
const (
	KeyBrightnessDown = "BrightnessDown"
	KeyBrightnessUp   = "BrightnessUp"
	KeyEject          = "Eject"
	KeyLogOff         = "LogOff"
	KeyPower          = "Power"
	KeyPowerOff       = "PowerOff"
	KeyPrintScreen    = "PrintScreen"
	KeyHibernate      = "Hibernate"
	KeyStandby        = "Standby"
	KeyWakeUp         = "WakeUp"
)

// IME and Composition Keys
// https://www.w3.org/TR/uievents-key/#keys-composition
const (
	KeyAllCandidates     = "AllCandidates"
	KeyAlphanumeric      = "Alphanumeric"
	KeyCodeInput         = "CodeInput"
	KeyCompose           = "Compose"
	KeyConvert           = "Convert"
	KeyDead              = "Dead"
	KeyFinalMode         = "FinalMode"
	KeyGroupFirst        = "GroupFirst"
	KeyGroupLast         = "GroupLast"
	KeyGroupNext         = "GroupNext"
	KeyGroupPrevious     = "GroupPrevious"
	KeyModeChange        = "ModeChange"
	KeyNextCandidate     = "NextCandidate"
	KeyNonConvert        = "NonConvert"
	KeyPreviousCandidate = "PreviousCandidate"
	KeyProcess           = "Process"
	KeySingleCandidate   = "SingleCandidate"
	KeyHangulMode        = "HangulMode"
	KeyHanjaMode         = "HanjaMode"
	KeyJunjaMode         = "JunjaMode"
	KeyEisu              = "Eisu"
	KeyHankaku           = "Hankaku"
	KeyHiragana          = "Hiragana"
	KeyHiraganaKatakana  = "HiraganaKatakana"
	KeyKanaMode          = "KanaMode"
	KeyKanjiMode         = "KanjiMode"
	KeyKatakana          = "Katakana"
	KeyRomaji            = "Romaji"
	KeyZenkaku           = "Zenkaku"
	KeyZenkakuHankaku    = "ZenkakuHankaku"
)

// General-Purpose Function Keys
// https://www.w3.org/TR/uievents-key/#keys-function
const (
	KeyF1    = "F1"
	KeyF2    = "F2"
	KeyF3    = "F3"
	KeyF4    = "F4"
	KeyF5    = "F5"
	KeyF6    = "F6"
	KeyF7    = "F7"
	KeyF8    = "F8"
	KeyF9    = "F9"
	KeyF10   = "F10"
	KeyF11   = "F11"
	KeyF12   = "F12"
	KeySoft1 = "Soft1"
	KeySoft2 = "Soft2"
	KeySoft3 = "Soft3"
	KeySoft4 = "Soft4"
)

// Multimedia Keys
// https://www.w3.org/TR/uievents-key/#keys-multimedia
const (
	KeyChannelDown        = "ChannelDown"
	KeyChannelUp          = "ChannelUp"
	KeyClose              = "Close"
	KeyMailForward        = "MailForward"
	KeyMailReply          = "MailReply"
	KeyMailSend           = "MailSend"
	KeyMediaClose         = "MediaClose"
	KeyMediaFastForward   = "MediaFastForward"
	KeyMediaPause         = "MediaPause"
	KeyMediaPlay          = "MediaPlay"
	KeyMediaPlayPause     = "MediaPlayPause"
	KeyMediaRecord        = "MediaRecord"
	KeyMediaRewind        = "MediaRewind"
	KeyMediaStop          = "MediaStop"
	KeyMediaTrackNext     = "MediaTrackNext"
	KeyMediaTrackPrevious = "MediaTrackPrevious"
	KeyNew                = "New"
	KeyOpen               = "Open"
	KeyPrint              = "Print"
	KeySave               = "Save"
	KeySpellCheck         = "SpellCheck"
)

// Multimedia Numpad Keys
// https://www.w3.org/TR/uievents-key/#keys-multimedia-numpad
const (
	KeyKey11 = "Key11"
	KeyKey12 = "Key12"
)

// Audio Keys
// https://www.w3.org/TR/uievents-key/#keys-audio
const (
	KeyAudioBalanceLeft      = "AudioBalanceLeft"
	KeyAudioBalanceRight     = "AudioBalanceRight"
	KeyAudioBassBoostDown    = "AudioBassBoostDown"
	KeyAudioBassBoostToggle  = "AudioBassBoostToggle"
	KeyAudioBassBoostUp      = "AudioBassBoostUp"
	KeyAudioFaderFront       = "AudioFaderFront"
	KeyAudioFaderRear        = "AudioFaderRear"
	KeyAudioSurroundModeNext = "AudioSurroundModeNext"
	KeyAudioTrebleDown       = "AudioTrebleDown"
	KeyAudioTrebleUp         = "AudioTrebleUp"
	KeyAudioVolumeDown       = "AudioVolumeDown"
	KeyAudioVolumeUp         = "AudioVolumeUp"
	KeyAudioVolumeMute       = "AudioVolumeMute"
	KeyMicrophoneToggle      = "MicrophoneToggle"
	KeyMicrophoneVolumeDown  = "MicrophoneVolumeDown"
	KeyMicrophoneVolumeUp    = "MicrophoneVolumeUp"
	KeyMicrophoneVolumeMute  = "MicrophoneVolumeMute"
)

// Speech Keys
// https://www.w3.org/TR/uievents-key/#keys-speech
const (
	KeySpeechCorrectionList = "SpeechCorrectionList"
	KeySpeechInputToggle    = "SpeechInputToggle"
)

// Application Keys
// https://www.w3.org/TR/uievents-key/#keys-apps
const (
	KeyLaunchApplication1  = "LaunchApplication1"
	KeyLaunchApplication2  = "LaunchApplication2"
	KeyLaunchCalendar      = "LaunchCalendar"
	KeyLaunchContacts      = "LaunchContacts"
	KeyLaunchMail          = "LaunchMail"
	KeyLaunchMediaPlayer   = "LaunchMediaPlayer"
	KeyLaunchMusicPlayer   = "LaunchMusicPlayer"
	KeyLaunchPhone         = "LaunchPhone"
	KeyLaunchScreenSaver   = "LaunchScreenSaver"
	KeyLaunchSpreadsheet   = "LaunchSpreadsheet"
	KeyLaunchWebBrowser    = "LaunchWebBrowser"
	KeyLaunchWebCam        = "LaunchWebCam"
	KeyLaunchWordProcessor = "LaunchWordProcessor"
)

// Browser Keys
// https://www.w3.org/TR/uievents-key/#keys-browser
const (
	KeyBrowserBack      = "BrowserBack"
	KeyBrowserFavorites = "BrowserFavorites"
	KeyBrowserForward   = "BrowserForward"
	KeyBrowserHome      = "BrowserHome"
	KeyBrowserRefresh   = "BrowserRefresh"
	KeyBrowserSearch    = "BrowserSearch"
	KeyBrowserStop      = "BrowserStop"
)

// Mobile Phone Keys
// https://www.w3.org/TR/uievents-key/#keys-mobile
const (
	KeyAppSwitch        = "AppSwitch"
	KeyCall             = "Call"
	KeyCamera           = "Camera"
	KeyCameraFocus      = "CameraFocus"
	KeyEndCall          = "EndCall"
	KeyGoBack           = "GoBack"
	KeyGoHome           = "GoHome"
	KeyHeadsetHook      = "HeadsetHook"
	KeyLastNumberRedial = "LastNumberRedial"
	KeyNotification     = "Notification"
	KeyMannerMode       = "MannerMode"
	KeyVoiceDial        = "VoiceDial"
)

// TV Keys
// https://www.w3.org/TR/uievents-key/#keys-tv
const (
	KeyTV                        = "TV"
	KeyTV3DMode                  = "TV3DMode"
	KeyTVAntennaCable            = "TVAntennaCable"
	KeyTVAudioDescription        = "TVAudioDescription"
	KeyTVAudioDescriptionMixDown = "TVAudioDescriptionMixDown"
	KeyTVAudioDescriptionMixUp   = "TVAudioDescriptionMixUp"
	KeyTVContentsMenu            = "TVContentsMenu"
	KeyTVDataService             = "TVDataService"
	KeyTVInput                   = "TVInput"
	KeyTVInputComponent1         = "TVInputComponent1"
	KeyTVInputComponent2         = "TVInputComponent2"
	KeyTVInputComposite1         = "TVInputComposite1"
	KeyTVInputComposite2         = "TVInputComposite2"
	KeyTVInputHDMI1              = "TVInputHDMI1"
	KeyTVInputHDMI2              = "TVInputHDMI2"
	KeyTVInputHDMI3              = "TVInputHDMI3"
	KeyTVInputHDMI4              = "TVInputHDMI4"
	KeyTVInputVGA1               = "TVInputVGA1"
	KeyTVMediaContext            = "TVMediaContext"
	KeyTVNetwork                 = "TVNetwork"
	KeyTVNumberEntry             = "TVNumberEntry"
	KeyTVPower                   = "TVPower"
	KeyTVRadioService            = "TVRadioService"
	KeyTVSatellite               = "TVSatellite"
	KeyTVSatelliteBS             = "TVSatelliteBS"
	KeyTVSatelliteCS             = "TVSatelliteCS"
	KeyTVSatelliteToggle         = "TVSatelliteToggle"
	KeyTVTerrestrialAnalog       = "TVTerrestrialAnalog"
	KeyTVTerrestrialDigital      = "TVTerrestrialDigital"
	KeyTVTimer                   = "TVTimer"
)

// Media Controller Keys
// https://www.w3.org/TR/uievents-key/#keys-media-controller
const (
	KeyAVRInput            = "AVRInput"
	KeyAVRPower            = "AVRPower"
	KeyColorF0Red          = "ColorF0Red"
	KeyColorF1Green        = "ColorF1Green"
	KeyColorF2Yellow       = "ColorF2Yellow"
	KeyColorF3Blue         = "ColorF3Blue"
	KeyColorF4Grey         = "ColorF4Grey"
	KeyColorF5Brown        = "ColorF5Brown"
	KeyClosedCaptionToggle = "ClosedCaptionToggle"
	KeyDimmer              = "Dimmer"
	KeyDisplaySwap         = "DisplaySwap"
	KeyDVR                 = "DVR"
	KeyExit                = "Exit"
	KeyFavoriteClear0      = "FavoriteClear0"
	KeyFavoriteClear1      = "FavoriteClear1"
	KeyFavoriteClear2      = "FavoriteClear2"
	KeyFavoriteClear3      = "FavoriteClear3"
	KeyFavoriteRecall0     = "FavoriteRecall0"
	KeyFavoriteRecall1     = "FavoriteRecall1"
	KeyFavoriteRecall2     = "FavoriteRecall2"
	KeyFavoriteRecall3     = "FavoriteRecall3"
	KeyFavoriteStore0      = "FavoriteStore0"
	KeyFavoriteStore1      = "FavoriteStore1"
	KeyFavoriteStore2      = "FavoriteStore2"
	KeyFavoriteStore3      = "FavoriteStore3"
	KeyGuide               = "Guide"
	KeyGuideNextDay        = "GuideNextDay"
	KeyGuidePreviousDay    = "GuidePreviousDay"
	KeyInfo                = "Info"
	KeyInstantReplay       = "InstantReplay"
	KeyLink                = "Link"
	KeyListProgram         = "ListProgram"
	KeyLiveContent         = "LiveContent"
	KeyLock                = "Lock"
	KeyMediaApps           = "MediaApps"
	KeyMediaAudioTrack     = "MediaAudioTrack"
	KeyMediaLast           = "MediaLast"
	KeyMediaSkipBackward   = "MediaSkipBackward"
	KeyMediaSkipForward    = "MediaSkipForward"
	KeyMediaStepBackward   = "MediaStepBackward"
	KeyMediaStepForward    = "MediaStepForward"
	KeyMediaTopMenu        = "MediaTopMenu"
	KeyNavigateIn          = "NavigateIn"
	KeyNavigateNext        = "NavigateNext"
	KeyNavigateOut         = "NavigateOut"
	KeyNavigatePrevious    = "NavigatePrevious"
	KeyNextFavoriteChannel = "NextFavoriteChannel"
	KeyNextUserProfile     = "NextUserProfile"
	KeyOnDemand            = "OnDemand"
	KeyPairing             = "Pairing"
	KeyPinPDown            = "PinPDown"
	KeyPinPMove            = "PinPMove"
	KeyPinPToggle          = "PinPToggle"
	KeyPinPUp              = "PinPUp"
	KeyPlaySpeedDown       = "PlaySpeedDown"
	KeyPlaySpeedReset      = "PlaySpeedReset"
	KeyPlaySpeedUp         = "PlaySpeedUp"
	KeyRandomToggle        = "RandomToggle"
	KeyRcLowBattery        = "RcLowBattery"
	KeyRecordSpeedNext     = "RecordSpeedNext"
	KeyRfBypass            = "RfBypass"
	KeyScanChannelsToggle  = "ScanChannelsToggle"
	KeyScreenModeNext      = "ScreenModeNext"
	KeySettings            = "Settings"
	KeySplitScreenToggle   = "SplitScreenToggle"
	KeySTBInput            = "STBInput"
	KeySTBPower            = "STBPower"
	KeySubtitle            = "Subtitle"
	KeyTeletext            = "Teletext"
	KeyVideoModeNext       = "VideoModeNext"
	KeyWink                = "Wink"
	KeyZoomToggle          = "ZoomToggle"
	KeyMediaNextTrack      = "MediaNextTrack"
	KeyMediaPreviousTrack  = "MediaPreviousTrack"
)
