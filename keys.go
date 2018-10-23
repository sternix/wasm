// +build js,wasm

package wasm

/*
https://www.w3.org/TR/uievents-key/
Types for KeyboardEvent.key
*/

type Key string

// Special Keys
// https://www.w3.org/TR/uievents-key/#keys-special
const (
	KeyUnidentified Key = "Unidentified"
)

// Modifier Keys
// https://www.w3.org/TR/uievents-key/#keys-modifier
const (
	KeyAlt        Key = "Alt"
	KeyAltGraph   Key = "AltGraph"
	KeyCapsLock   Key = "CapsLock"
	KeyControl    Key = "Control"
	KeyFn         Key = "Fn"
	KeyFnLock     Key = "FnLock"
	KeyMeta       Key = "Meta"
	KeyNumLock    Key = "NumLock"
	KeyScrollLock Key = "ScrollLock"
	KeyShift      Key = "Shift"
	KeySymbol     Key = "Symbol"
	KeySymbolLock Key = "SymbolLock"
)

//Whitespace Keys
// https://www.w3.org/TR/uievents-key/#keys-whitespace
const (
	KeyEnter Key = "Enter"
	KeyTab   Key = "Tab"
	KeySpace Key = " "
)

// Navigation Keys
// https://www.w3.org/TR/uievents-key/#keys-navigation
const (
	KeyArrowDown  Key = "ArrowDown"
	KeyArrowLeft  Key = "ArrowLeft"
	KeyArrowRight Key = "ArrowRight"
	KeyArrowUp    Key = "ArrowUp"
	KeyEnd        Key = "End"
	KeyHome       Key = "Home"
	KeyPageDown   Key = "PageDown"
	KeyPageUp     Key = "PageUp"
)

// Editing keys
// https://www.w3.org/TR/uievents-key/#keys-editing
const (
	KeyBackspace Key = "Backspace"
	KeyClear     Key = "Clear"
	KeyCopy      Key = "Copy"
	KeyCrSel     Key = "CrSel"
	KeyCut       Key = "Cut"
	KeyDelete    Key = "Delete"
	KeyEraseEof  Key = "EraseEof"
	KeyExSel     Key = "ExSel"
	KeyInsert    Key = "Insert"
	KeyPaste     Key = "Paste"
	KeyRedo      Key = "Redo"
	KeyUndo      Key = "Undo"
)

// UI Keys
// https://www.w3.org/TR/uievents-key/#keys-ui
const (
	KeyAccept      Key = "Accept"
	KeyAgain       Key = "Again"
	KeyAttn        Key = "Attn"
	KeyCancel      Key = "Cancel"
	KeyContextMenu Key = "ContextMenu"
	KeyEscape      Key = "Escape"
	KeyExecute     Key = "Execute"
	KeyFind        Key = "Find"
	KeyHelp        Key = "Help"
	KeyPause       Key = "Pause"
	KeyPlay        Key = "Play"
	KeyProps       Key = "Props"
	KeySelect      Key = "Select"
	KeyZoomIn      Key = "ZoomIn"
	KeyZoomOut     Key = "ZoomOut"
)

// Device Keys
// https://www.w3.org/TR/uievents-key/#keys-device
const (
	KeyBrightnessDown Key = "BrightnessDown"
	KeyBrightnessUp   Key = "BrightnessUp"
	KeyEject          Key = "Eject"
	KeyLogOff         Key = "LogOff"
	KeyPower          Key = "Power"
	KeyPowerOff       Key = "PowerOff"
	KeyPrintScreen    Key = "PrintScreen"
	KeyHibernate      Key = "Hibernate"
	KeyStandby        Key = "Standby"
	KeyWakeUp         Key = "WakeUp"
)

// IME and Composition Keys
// https://www.w3.org/TR/uievents-key/#keys-composition
const (
	KeyAllCandidates     Key = "AllCandidates"
	KeyAlphanumeric      Key = "Alphanumeric"
	KeyCodeInput         Key = "CodeInput"
	KeyCompose           Key = "Compose"
	KeyConvert           Key = "Convert"
	KeyDead              Key = "Dead"
	KeyFinalMode         Key = "FinalMode"
	KeyGroupFirst        Key = "GroupFirst"
	KeyGroupLast         Key = "GroupLast"
	KeyGroupNext         Key = "GroupNext"
	KeyGroupPrevious     Key = "GroupPrevious"
	KeyModeChange        Key = "ModeChange"
	KeyNextCandidate     Key = "NextCandidate"
	KeyNonConvert        Key = "NonConvert"
	KeyPreviousCandidate Key = "PreviousCandidate"
	KeyProcess           Key = "Process"
	KeySingleCandidate   Key = "SingleCandidate"
	KeyHangulMode        Key = "HangulMode"
	KeyHanjaMode         Key = "HanjaMode"
	KeyJunjaMode         Key = "JunjaMode"
	KeyEisu              Key = "Eisu"
	KeyHankaku           Key = "Hankaku"
	KeyHiragana          Key = "Hiragana"
	KeyHiraganaKatakana  Key = "HiraganaKatakana"
	KeyKanaMode          Key = "KanaMode"
	KeyKanjiMode         Key = "KanjiMode"
	KeyKatakana          Key = "Katakana"
	KeyRomaji            Key = "Romaji"
	KeyZenkaku           Key = "Zenkaku"
	KeyZenkakuHankaku    Key = "ZenkakuHankaku"
)

// General-Purpose Function Keys
// https://www.w3.org/TR/uievents-key/#keys-function
const (
	KeyF1    Key = "F1"
	KeyF2    Key = "F2"
	KeyF3    Key = "F3"
	KeyF4    Key = "F4"
	KeyF5    Key = "F5"
	KeyF6    Key = "F6"
	KeyF7    Key = "F7"
	KeyF8    Key = "F8"
	KeyF9    Key = "F9"
	KeyF10   Key = "F10"
	KeyF11   Key = "F11"
	KeyF12   Key = "F12"
	KeySoft1 Key = "Soft1"
	KeySoft2 Key = "Soft2"
	KeySoft3 Key = "Soft3"
	KeySoft4 Key = "Soft4"
)

// Multimedia Keys
// https://www.w3.org/TR/uievents-key/#keys-multimedia
const (
	KeyChannelDown        Key = "ChannelDown"
	KeyChannelUp          Key = "ChannelUp"
	KeyClose              Key = "Close"
	KeyMailForward        Key = "MailForward"
	KeyMailReply          Key = "MailReply"
	KeyMailSend           Key = "MailSend"
	KeyMediaClose         Key = "MediaClose"
	KeyMediaFastForward   Key = "MediaFastForward"
	KeyMediaPause         Key = "MediaPause"
	KeyMediaPlay          Key = "MediaPlay"
	KeyMediaPlayPause     Key = "MediaPlayPause"
	KeyMediaRecord        Key = "MediaRecord"
	KeyMediaRewind        Key = "MediaRewind"
	KeyMediaStop          Key = "MediaStop"
	KeyMediaTrackNext     Key = "MediaTrackNext"
	KeyMediaTrackPrevious Key = "MediaTrackPrevious"
	KeyNew                Key = "New"
	KeyOpen               Key = "Open"
	KeyPrint              Key = "Print"
	KeySave               Key = "Save"
	KeySpellCheck         Key = "SpellCheck"
)

// Multimedia Numpad Keys
// https://www.w3.org/TR/uievents-key/#keys-multimedia-numpad
const (
	KeyKey11 Key = "Key11"
	KeyKey12 Key = "Key12"
)

// Audio Keys
// https://www.w3.org/TR/uievents-key/#keys-audio
const (
	KeyAudioBalanceLeft      Key = "AudioBalanceLeft"
	KeyAudioBalanceRight     Key = "AudioBalanceRight"
	KeyAudioBassBoostDown    Key = "AudioBassBoostDown"
	KeyAudioBassBoostToggle  Key = "AudioBassBoostToggle"
	KeyAudioBassBoostUp      Key = "AudioBassBoostUp"
	KeyAudioFaderFront       Key = "AudioFaderFront"
	KeyAudioFaderRear        Key = "AudioFaderRear"
	KeyAudioSurroundModeNext Key = "AudioSurroundModeNext"
	KeyAudioTrebleDown       Key = "AudioTrebleDown"
	KeyAudioTrebleUp         Key = "AudioTrebleUp"
	KeyAudioVolumeDown       Key = "AudioVolumeDown"
	KeyAudioVolumeUp         Key = "AudioVolumeUp"
	KeyAudioVolumeMute       Key = "AudioVolumeMute"
	KeyMicrophoneToggle      Key = "MicrophoneToggle"
	KeyMicrophoneVolumeDown  Key = "MicrophoneVolumeDown"
	KeyMicrophoneVolumeUp    Key = "MicrophoneVolumeUp"
	KeyMicrophoneVolumeMute  Key = "MicrophoneVolumeMute"
)

// Speech Keys
// https://www.w3.org/TR/uievents-key/#keys-speech
const (
	KeySpeechCorrectionList Key = "SpeechCorrectionList"
	KeySpeechInputToggle    Key = "SpeechInputToggle"
)

// Application Keys
// https://www.w3.org/TR/uievents-key/#keys-apps
const (
	KeyLaunchApplication1  Key = "LaunchApplication1"
	KeyLaunchApplication2  Key = "LaunchApplication2"
	KeyLaunchCalendar      Key = "LaunchCalendar"
	KeyLaunchContacts      Key = "LaunchContacts"
	KeyLaunchMail          Key = "LaunchMail"
	KeyLaunchMediaPlayer   Key = "LaunchMediaPlayer"
	KeyLaunchMusicPlayer   Key = "LaunchMusicPlayer"
	KeyLaunchPhone         Key = "LaunchPhone"
	KeyLaunchScreenSaver   Key = "LaunchScreenSaver"
	KeyLaunchSpreadsheet   Key = "LaunchSpreadsheet"
	KeyLaunchWebBrowser    Key = "LaunchWebBrowser"
	KeyLaunchWebCam        Key = "LaunchWebCam"
	KeyLaunchWordProcessor Key = "LaunchWordProcessor"
)

// Browser Keys
// https://www.w3.org/TR/uievents-key/#keys-browser
const (
	KeyBrowserBack      Key = "BrowserBack"
	KeyBrowserFavorites Key = "BrowserFavorites"
	KeyBrowserForward   Key = "BrowserForward"
	KeyBrowserHome      Key = "BrowserHome"
	KeyBrowserRefresh   Key = "BrowserRefresh"
	KeyBrowserSearch    Key = "BrowserSearch"
	KeyBrowserStop      Key = "BrowserStop"
)

// Mobile Phone Keys
// https://www.w3.org/TR/uievents-key/#keys-mobile
const (
	KeyAppSwitch        Key = "AppSwitch"
	KeyCall             Key = "Call"
	KeyCamera           Key = "Camera"
	KeyCameraFocus      Key = "CameraFocus"
	KeyEndCall          Key = "EndCall"
	KeyGoBack           Key = "GoBack"
	KeyGoHome           Key = "GoHome"
	KeyHeadsetHook      Key = "HeadsetHook"
	KeyLastNumberRedial Key = "LastNumberRedial"
	KeyNotification     Key = "Notification"
	KeyMannerMode       Key = "MannerMode"
	KeyVoiceDial        Key = "VoiceDial"
)

// TV Keys
// https://www.w3.org/TR/uievents-key/#keys-tv
const (
	KeyTV                        Key = "TV"
	KeyTV3DMode                  Key = "TV3DMode"
	KeyTVAntennaCable            Key = "TVAntennaCable"
	KeyTVAudioDescription        Key = "TVAudioDescription"
	KeyTVAudioDescriptionMixDown Key = "TVAudioDescriptionMixDown"
	KeyTVAudioDescriptionMixUp   Key = "TVAudioDescriptionMixUp"
	KeyTVContentsMenu            Key = "TVContentsMenu"
	KeyTVDataService             Key = "TVDataService"
	KeyTVInput                   Key = "TVInput"
	KeyTVInputComponent1         Key = "TVInputComponent1"
	KeyTVInputComponent2         Key = "TVInputComponent2"
	KeyTVInputComposite1         Key = "TVInputComposite1"
	KeyTVInputComposite2         Key = "TVInputComposite2"
	KeyTVInputHDMI1              Key = "TVInputHDMI1"
	KeyTVInputHDMI2              Key = "TVInputHDMI2"
	KeyTVInputHDMI3              Key = "TVInputHDMI3"
	KeyTVInputHDMI4              Key = "TVInputHDMI4"
	KeyTVInputVGA1               Key = "TVInputVGA1"
	KeyTVMediaContext            Key = "TVMediaContext"
	KeyTVNetwork                 Key = "TVNetwork"
	KeyTVNumberEntry             Key = "TVNumberEntry"
	KeyTVPower                   Key = "TVPower"
	KeyTVRadioService            Key = "TVRadioService"
	KeyTVSatellite               Key = "TVSatellite"
	KeyTVSatelliteBS             Key = "TVSatelliteBS"
	KeyTVSatelliteCS             Key = "TVSatelliteCS"
	KeyTVSatelliteToggle         Key = "TVSatelliteToggle"
	KeyTVTerrestrialAnalog       Key = "TVTerrestrialAnalog"
	KeyTVTerrestrialDigital      Key = "TVTerrestrialDigital"
	KeyTVTimer                   Key = "TVTimer"
)

// Media Controller Keys
// https://www.w3.org/TR/uievents-key/#keys-media-controller
const (
	KeyAVRInput            Key = "AVRInput"
	KeyAVRPower            Key = "AVRPower"
	KeyColorF0Red          Key = "ColorF0Red"
	KeyColorF1Green        Key = "ColorF1Green"
	KeyColorF2Yellow       Key = "ColorF2Yellow"
	KeyColorF3Blue         Key = "ColorF3Blue"
	KeyColorF4Grey         Key = "ColorF4Grey"
	KeyColorF5Brown        Key = "ColorF5Brown"
	KeyClosedCaptionToggle Key = "ClosedCaptionToggle"
	KeyDimmer              Key = "Dimmer"
	KeyDisplaySwap         Key = "DisplaySwap"
	KeyDVR                 Key = "DVR"
	KeyExit                Key = "Exit"
	KeyFavoriteClear0      Key = "FavoriteClear0"
	KeyFavoriteClear1      Key = "FavoriteClear1"
	KeyFavoriteClear2      Key = "FavoriteClear2"
	KeyFavoriteClear3      Key = "FavoriteClear3"
	KeyFavoriteRecall0     Key = "FavoriteRecall0"
	KeyFavoriteRecall1     Key = "FavoriteRecall1"
	KeyFavoriteRecall2     Key = "FavoriteRecall2"
	KeyFavoriteRecall3     Key = "FavoriteRecall3"
	KeyFavoriteStore0      Key = "FavoriteStore0"
	KeyFavoriteStore1      Key = "FavoriteStore1"
	KeyFavoriteStore2      Key = "FavoriteStore2"
	KeyFavoriteStore3      Key = "FavoriteStore3"
	KeyGuide               Key = "Guide"
	KeyGuideNextDay        Key = "GuideNextDay"
	KeyGuidePreviousDay    Key = "GuidePreviousDay"
	KeyInfo                Key = "Info"
	KeyInstantReplay       Key = "InstantReplay"
	KeyLink                Key = "Link"
	KeyListProgram         Key = "ListProgram"
	KeyLiveContent         Key = "LiveContent"
	KeyLock                Key = "Lock"
	KeyMediaApps           Key = "MediaApps"
	KeyMediaAudioTrack     Key = "MediaAudioTrack"
	KeyMediaLast           Key = "MediaLast"
	KeyMediaSkipBackward   Key = "MediaSkipBackward"
	KeyMediaSkipForward    Key = "MediaSkipForward"
	KeyMediaStepBackward   Key = "MediaStepBackward"
	KeyMediaStepForward    Key = "MediaStepForward"
	KeyMediaTopMenu        Key = "MediaTopMenu"
	KeyNavigateIn          Key = "NavigateIn"
	KeyNavigateNext        Key = "NavigateNext"
	KeyNavigateOut         Key = "NavigateOut"
	KeyNavigatePrevious    Key = "NavigatePrevious"
	KeyNextFavoriteChannel Key = "NextFavoriteChannel"
	KeyNextUserProfile     Key = "NextUserProfile"
	KeyOnDemand            Key = "OnDemand"
	KeyPairing             Key = "Pairing"
	KeyPinPDown            Key = "PinPDown"
	KeyPinPMove            Key = "PinPMove"
	KeyPinPToggle          Key = "PinPToggle"
	KeyPinPUp              Key = "PinPUp"
	KeyPlaySpeedDown       Key = "PlaySpeedDown"
	KeyPlaySpeedReset      Key = "PlaySpeedReset"
	KeyPlaySpeedUp         Key = "PlaySpeedUp"
	KeyRandomToggle        Key = "RandomToggle"
	KeyRcLowBattery        Key = "RcLowBattery"
	KeyRecordSpeedNext     Key = "RecordSpeedNext"
	KeyRfBypass            Key = "RfBypass"
	KeyScanChannelsToggle  Key = "ScanChannelsToggle"
	KeyScreenModeNext      Key = "ScreenModeNext"
	KeySettings            Key = "Settings"
	KeySplitScreenToggle   Key = "SplitScreenToggle"
	KeySTBInput            Key = "STBInput"
	KeySTBPower            Key = "STBPower"
	KeySubtitle            Key = "Subtitle"
	KeyTeletext            Key = "Teletext"
	KeyVideoModeNext       Key = "VideoModeNext"
	KeyWink                Key = "Wink"
	KeyZoomToggle          Key = "ZoomToggle"
	KeyMediaNextTrack      Key = "MediaNextTrack"
	KeyMediaPreviousTrack  Key = "MediaPreviousTrack"
)
