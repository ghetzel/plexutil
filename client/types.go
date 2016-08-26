package client

type Setting struct {
	Advanced   int    `xml:"advanced,attr,omitempty" json:"Advanced,omitempty"`     // "1"
	Default    string `xml:"default,attr,omitempty" json:"Default,omitempty"`       // ""
	EnumValues string `xml:"enumValues,attr,omitempty" json:"EnumValues,omitempty"` // "never:never|scheduled:as a scheduled task|asap:as a scheduled task and when media is added"
	Group      string `xml:"group,attr,omitempty" json:"Group,omitempty"`           // "channels"
	Hidden     int    `xml:"hidden,attr,omitempty" json:"Hidden,omitempty"`         // "0"
	Id         string `xml:"id,attr,omitempty" json:"Id,omitempty"`                 // "iPhotoLibraryXmlPath"
	Label      string `xml:"label,attr,omitempty" json:"Label,omitempty"`           // "iPhoto library XML path"
	Summary    string `xml:"summary,attr,omitempty" json:"Summary,omitempty"`       // ""
	Type       string `xml:"type,attr,omitempty" json:"Type,omitempty"`             // "text"
	Value      string `xml:"value,attr,omitempty" json:"Value,omitempty"`           // ""
}

type Location struct {
	LibrarySectionID int    `xml:"librarySectionID,attr,omitempty" json:"LibrarySectionID,omitempty"` // "4"
	Path             string `xml:"path,attr,omitempty" json:"Path,omitempty"`                         // "/cortex/central/videos/lib/TV/The A-Team"
}

type Directory struct {
	AddedAt               int        `xml:"addedAt,attr,omitempty" json:"AddedAt,omitempty"` // "1471546619"
	Agent                 string     `xml:"agent,attr,omitempty" json:"Agent,omitempty"`
	AllowSync             int        `xml:"allowSync,attr,omitempty" json:"AllowSync,omitempty"`   // "1"
	Art                   string     `xml:"art,attr,omitempty" json:"Art,omitempty"`               // "/library/metadata/13814/art/1471546640"
	Banner                string     `xml:"banner,attr,omitempty" json:"Banner,omitempty"`         // "/library/metadata/13814/banner/1471546640"
	ChildCount            int        `xml:"childCount,attr,omitempty" json:"ChildCount,omitempty"` // "5"
	Composite             string     `xml:"composite,attr,omitempty" json:"Composite,omitempty"`
	ContentRating         string     `xml:"contentRating,attr,omitempty" json:"ContentRating,omitempty"` // "TV-PG"
	CreatedAt             int        `xml:"createdAt,attr,omitempty" json:"CreatedAt,omitempty"`         // "1471546640"
	Duration              int        `xml:"duration,attr,omitempty" json:"Duration,omitempty"`           // "2700000"
	Filters               int        `xml:"filters,attr,omitempty" json:"Filters,omitempty"`
	Guid                  string     `xml:"guid,attr,omitempty" json:"Guid,omitempty"`   // "com.plexapp.agents.thetvdb://77904"
	Index                 int        `xml:"index,attr,omitempty" json:"Index,omitempty"` // "3"
	Key                   string     `xml:"key,attr,omitempty" json:"Key,omitempty"`     // "/library/metadata/20448/children"
	Language              string     `xml:"language,attr,omitempty" json:"Language,omitempty"`
	LeafCount             int        `xml:"leafCount,attr,omitempty" json:"LeafCount,omitempty"`               // "98"
	LibrarySectionID      int        `xml:"librarySectionID,attr,omitempty" json:"LibrarySectionID,omitempty"` // "4"
	Locations             []Location `xml:"Location" json:"Location"`
	OriginallyAvailableAt string     `xml:"originallyAvailableAt,omitempty" json:"OriginallyAvailableAt,omitempty"` // "1983-01-23"
	ParentIndex           int        `xml:"parentIndex,attr,omitempty" json:"ParentIndex,omitempty"`                // "1"
	ParentKey             string     `xml:"parentKey,attr,omitempty" json:"ParentKey,omitempty"`                    // "/library/metadata/14207"
	ParentRatingKey       int        `xml:"parentRatingKey,attr,omitempty" json:"ParentRatingKey,omitempty"`        // "14207"
	ParentTheme           string     `xml:"parentTheme,attr,omitempty" json:"ParentTheme,omitempty"`                // "/library/metadata/14207/theme/1471544783"
	ParentThumb           string     `xml:"parentThumb,attr,omitempty" json:"ParentThumb,omitempty"`                // "/library/metadata/14207/thumb/1471544783"
	ParentTitle           string     `xml:"parentTitle,attr,omitempty" json:"ParentTitle,omitempty"`                // "The O.C."
	Rating                float32    `xml:"rating,attr,omitempty" json:"Rating,omitempty"`                          // "7.8"
	RatingKey             int        `xml:"ratingKey,attr,omitempty" json:"RatingKey,omitempty"`                    // "20448"
	Refreshing            int        `xml:"refreshing,attr,omitempty" json:"Refreshing,omitempty"`
	Scanner               string     `xml:"scanner,attr,omitempty" json:"Scanner,omitempty"`
	Studio                string     `xml:"studio,attr,omitempty" json:"Studio,omitempty"`       // "NBC"
	Summary               string     `xml:"summary,attr,omitempty" json:"Summary,omitempty"`     // "The A-Team is about a group of ex-United States Army Special Forces personnel who work as soldiers of fortune, while on the run from the Army after being branded as war criminals for a crime they didn't commit."
	Theme                 string     `xml:"theme,attr,omitempty" json:"Theme,omitempty"`         // "/library/metadata/13814/theme/1471546640"
	Thumb                 string     `xml:"thumb,attr,omitempty" json:"Thumb,omitempty"`         // "/library/metadata/13814/thumb/1471546640"
	Title                 string     `xml:"title,attr,omitempty" json:"Title,omitempty"`         // "The A-Team"
	TitleSort             string     `xml:"titleSort,attr,omitempty" json:"TitleSort,omitempty"` // "A-Team"
	Type                  string     `xml:"type,attr,omitempty" json:"Type,omitempty"`           // "show"
	UpdatedAt             int        `xml:"updatedAt,attr,omitempty" json:"UpdatedAt,omitempty"` // "1471546640"
	UUID                  string     `xml:"uuid,attr,omitempty" json:"Uuid,omitempty"`
	ViewedLeafCount       int        `xml:"viewedLeafCount,attr,omitempty" json:"ViewedLeafCount,omitempty"` // "52"
	Year                  int        `xml:"year,attr,omitempty" json:"Year,omitempty"`                       // "1983"
}

type Hub struct {
	Videos        []Video     `xml:"Video" json:"Video"`
	Directories   []Directory `xml:"Directory" json:"Directory"`
	Key           string      `xml:"key,omitempty" json:"Key,omitempty"`                          // "/hubs/home/recentlyAdded?type=1&amp;personal=1"
	Type          string      `xml:"type,attr,omitempty" json:"Type,omitempty"`                   // "clip"
	HubIdentifier string      `xml:"hubIdentifier,attr,omitempty" json:"HubIdentifier,omitempty"` // "home.videos.recent"
	Size          int         `xml:"size,attr,omitempty" json:"Size,omitempty"`                   // "0"
	Title         string      `xml:"title,attr,omitempty" json:"Title,omitempty"`                 // "Recently Added Home Videos"
	More          int         `xml:"more,attr,omitempty" json:"More,omitempty"`                   // "0"
}

type MediaContainer struct {
	AllowSync                     int         `xml:"allowSync,attr,omitempty" json:"AllowSync,omitempty"` // "1"
	Art                           string      `xml:"art,attr,omitempty" json:"Art,omitempty"`
	Directories                   []Directory `xml:"Directory" json:"Directory"`
	Hubs                          []Hub       `xml:"Hub" json:"Hub"`
	Identifier                    string      `xml:"identifier,omitempty" json:"Identifier,omitempty"`                  // "com.plexapp.plugins.library"
	LibrarySectionID              int         `xml:"librarySectionID,attr,omitempty" json:"LibrarySectionID,omitempty"` // "4"
	LibrarySectionTitle           string      `xml:"librarySectionTitle,attr,omitempty" json:"LibrarySectionTitle,omitempty"`
	LibrarySectionUUID            string      `xml:"librarySectionUUID,attr,omitempty" json:"LibrarySectionUUID,omitempty"`
	MediaTagPrefix                string      `xml:"mediaTagPrefix,attr,omitempty" json:"MediaTagPrefix,omitempty"`   // "/system/bundle/media/flags/"
	MediaTagVersion               int         `xml:"mediaTagVersion,attr,omitempty" json:"MediaTagVersion,omitempty"` // "1466734815"
	NoCache                       int         `xml:"nocache,attr,omitempty" json:"Nocache,omitempty"`
	Offset                        int         `xml:"offset,attr,omitempty" json:"Offset,omitempty"`
	Settings                      []Setting   `xml:"Setting" json:"Setting"`
	Size                          int         `xml:"size,omitempty" json:"Size,omitempty"` // "8"
	SortAscending                 int         `xml:"sortAsc,attr,omitempty" json:"SortAsc,omitempty"`
	Thumb                         string      `xml:"thumb,attr,omitempty" json:"Thumb,omitempty"`
	Title1                        string      `xml:"title1,attr,omitempty" json:"Title1,omitempty"` // "Plex Library"
	Title2                        string      `xml:"title2,attr,omitempty" json:"Title2,omitempty"` // "Plex Library"
	TotalSize                     int         `xml:"totalSize,attr,omitempty" json:"TotalSize,omitempty"`
	Videos                        []Video     `xml:"Video" json:"Video"`
	ViewGroup                     string      `xml:"viewGroup,omitempty" json:"ViewGroup,omitempty"`
	ViewMode                      int         `xml:"viewMode,attr,omitempty" json:"ViewMode,omitempty"`
	AllowCameraUpload             int         `xml:"allowCameraUpload,attr,omitempty" json:"AllowCameraUpload,omitempty"`                         // "1"
	AllowChannelAccess            int         `xml:"allowChannelAccess,attr,omitempty" json:"AllowChannelAccess,omitempty"`                       // "1"
	AllowSharing                  int         `xml:"allowSharing,attr,omitempty" json:"AllowSharing,omitempty"`                                   // "1"
	BackgroundProcessing          int         `xml:"backgroundProcessing,attr,omitempty" json:"BackgroundProcessing,omitempty"`                   // "1"
	Certificate                   int         `xml:"certificate,attr,omitempty" json:"Certificate,omitempty"`                                     // "1"
	CompanionProxy                int         `xml:"companionProxy,attr,omitempty" json:"CompanionProxy,omitempty"`                               // "1"
	EventStream                   int         `xml:"eventStream,attr,omitempty" json:"EventStream,omitempty"`                                     // "1"
	FriendlyName                  string      `xml:"friendlyName,attr,omitempty" json:"FriendlyName,omitempty"`                                   // "Cortex"
	HubSearch                     int         `xml:"hubSearch,attr,omitempty" json:"HubSearch,omitempty"`                                         // "1"
	MachineIdentifier             string      `xml:"machineIdentifier,attr,omitempty" json:"MachineIdentifier,omitempty"`                         // "91dbc3a3ec55607b993cb1ec6b10983f9f4bb407"
	Multiuser                     int         `xml:"multiuser,attr,omitempty" json:"Multiuser,omitempty"`                                         // "1"
	MyPlex                        int         `xml:"myPlex,attr,omitempty" json:"MyPlex,omitempty"`                                               // "1"
	MyPlexMappingState            string      `xml:"myPlexMappingState,attr,omitempty" json:"MyPlexMappingState,omitempty"`                       // "mapped"
	MyPlexSigninState             string      `xml:"myPlexSigninState,attr,omitempty" json:"MyPlexSigninState,omitempty"`                         // "ok"
	MyPlexSubscription            int         `xml:"myPlexSubscription,attr,omitempty" json:"MyPlexSubscription,omitempty"`                       // "1"
	MyPlexUsername                string      `xml:"myPlexUsername,attr,omitempty" json:"MyPlexUsername,omitempty"`                               // "garyhetzel@gmail.com"
	OwnerFeatures                 string      `xml:"ownerFeatures,attr,omitempty" json:"OwnerFeatures,omitempty"`                                 // "camera_upload,cloudsync,content_filter,home,lyrics,music_videos,pass,premium_music_metadata,sync,trailers"
	Platform                      string      `xml:"platform,attr,omitempty" json:"Platform,omitempty"`                                           // "FreeBSD"
	PlatformVersion               string      `xml:"platformVersion,attr,omitempty" json:"PlatformVersion,omitempty"`                             // "10.3-RELEASE (FreeBSD 10.3-RELEASE #0 r297264: Fri Mar 25 02:10:02 UTC 2016 root@releng1.nyi.freebsd.org:/usr/obj/usr/src/sys/GENERIC)"
	PluginHost                    int         `xml:"pluginHost,attr,omitempty" json:"PluginHost,omitempty"`                                       // "1"
	ReadOnlyLibraries             int         `xml:"readOnlyLibraries,attr,omitempty" json:"ReadOnlyLibraries,omitempty"`                         // "0"
	RequestParametersInCookie     int         `xml:"requestParametersInCookie,attr,omitempty" json:"RequestParametersInCookie,omitempty"`         // "1"
	Sync                          int         `xml:"sync,attr,omitempty" json:"Sync,omitempty"`                                                   // "1"
	TranscoderActiveVideoSessions int         `xml:"transcoderActiveVideoSessions,attr,omitempty" json:"TranscoderActiveVideoSessions,omitempty"` // "0"
	TranscoderAudio               int         `xml:"transcoderAudio,attr,omitempty" json:"TranscoderAudio,omitempty"`                             // "1"
	TranscoderLyrics              int         `xml:"transcoderLyrics,attr,omitempty" json:"TranscoderLyrics,omitempty"`                           // "1"
	TranscoderPhoto               int         `xml:"transcoderPhoto,attr,omitempty" json:"TranscoderPhoto,omitempty"`                             // "1"
	TranscoderSubtitles           int         `xml:"transcoderSubtitles,attr,omitempty" json:"TranscoderSubtitles,omitempty"`                     // "1"
	TranscoderVideo               int         `xml:"transcoderVideo,attr,omitempty" json:"TranscoderVideo,omitempty"`                             // "1"
	TranscoderVideoBitrates       string      `xml:"transcoderVideoBitrates,attr,omitempty" json:"TranscoderVideoBitrates,omitempty"`             // "64,96,208,320,720,1500,2000,3000,4000,8000,10000,12000,20000"
	TranscoderVideoQualities      string      `xml:"transcoderVideoQualities,attr,omitempty" json:"TranscoderVideoQualities,omitempty"`           // "0,1,2,3,4,5,6,7,8,9,10,11,12"
	TranscoderVideoResolutions    string      `xml:"transcoderVideoResolutions,attr,omitempty" json:"TranscoderVideoResolutions,omitempty"`       // "128,128,160,240,320,480,768,720,720,1080,1080,1080,1080"
	UpdatedAt                     int         `xml:"updatedAt,attr,omitempty" json:"UpdatedAt,omitempty"`                                         // "1471466714"
	Updater                       int         `xml:"updater,attr,omitempty" json:"Updater,omitempty"`                                             // "1"
	Version                       string      `xml:"version,attr,omitempty" json:"Version,omitempty"`                                             // "1.0.0.2261-a17e99e"
}

type User struct {
	Id    int    `xml:"id,attr,omitempty" json:"Id,omitempty"`
	Title string `xml:"title,attr,omitempty" json:"Title,omitempty"`
}

type Player struct {
	Address           string `xml:"address,attr,omitempty" json:"Address,omitempty"`                     // "73.78.220.221"
	Device            string `xml:"device,attr,omitempty" json:"Device,omitempty"`                       // "Windows"
	MachineIdentifier string `xml:"machineIdentifier,attr,omitempty" json:"MachineIdentifier,omitempty"` // "4648a824-5ef2-4e40-a044-86fdd6b76bac"
	Model             string `xml:"model,attr,omitempty" json:"Model,omitempty"`                         // ""
	Platform          string `xml:"platform,attr,omitempty" json:"Platform,omitempty"`                   // "Chrome"
	PlatformVersion   string `xml:"platformVersion,attr,omitempty" json:"PlatformVersion,omitempty"`     // "52.0"
	Product           string `xml:"product,attr,omitempty" json:"Product,omitempty"`                     // "Plex Web"
	Profile           string `xml:"profile,attr,omitempty" json:"Profile,omitempty"`                     // "Web"
	State             string `xml:"state,attr,omitempty" json:"State,omitempty"`                         // "playing"
	Title             string `xml:"title,attr,omitempty" json:"Title,omitempty"`                         // "Plex Web (Chrome)"
	Vendor            string `xml:"vendor,attr,omitempty" json:"Vendor,omitempty"`                       // ""
	Version           string `xml:"version,attr,omitempty" json:"Version,omitempty"`                     // "2.7.8"
}

type TranscodeSession struct {
	AudioChannels int     `xml:"audioChannels,attr,omitempty" json:"AudioChannels,omitempty"` // "2"
	AudioCodec    string  `xml:"audioCodec,attr,omitempty" json:"AudioCodec,omitempty"`       // "mp3"
	AudioDecision string  `xml:"audioDecision,attr,omitempty" json:"AudioDecision,omitempty"` // "copy"
	Complete      int     `xml:"complete,attr,omitempty" json:"Complete,omitempty"`           // "0"
	Container     string  `xml:"container,attr,omitempty" json:"Container,omitempty"`         // "mkv"
	Context       string  `xml:"context,attr,omitempty" json:"Context,omitempty"`             // "streaming"
	Duration      int     `xml:"duration,attr,omitempty" json:"Duration,omitempty"`           // "3717000"
	Height        int     `xml:"height,attr,omitempty" json:"Height,omitempty"`               // "352"
	Key           string  `xml:"key,attr,omitempty" json:"Key,omitempty"`                     // "fs3uwoq9w1yv7o5bj0d15g66r"
	Progress      float64 `xml:"progress,attr,omitempty" json:"Progress,omitempty"`           // "26.899999618530273"
	Protocol      string  `xml:"protocol,attr,omitempty" json:"Protocol,omitempty"`           // "http"
	Remaining     int     `xml:"remaining,attr,omitempty" json:"Remaining,omitempty"`         // "19543"
	Speed         int     `xml:"speed,attr,omitempty" json:"Speed,omitempty"`                 // "0"
	Throttled     int     `xml:"throttled,attr,omitempty" json:"Throttled,omitempty"`         // "1"
	VideoCodec    string  `xml:"videoCodec,attr,omitempty" json:"VideoCodec,omitempty"`       // "h264"
	VideoDecision string  `xml:"videoDecision,attr,omitempty" json:"VideoDecision,omitempty"` // "transcode"
	Width         int     `xml:"width,attr,omitempty" json:"Width,omitempty"`                 // "640"
}
