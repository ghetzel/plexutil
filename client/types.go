package client

type Setting struct {
	Advanced   int    `xml:"advanced,attr,omitempty" json:"advanced,omitempty"`     // "1"
	Default    string `xml:"default,attr,omitempty" json:"default,omitempty"`       // ""
	EnumValues string `xml:"enumValues,attr,omitempty" json:"enumValues,omitempty"` // "never:never|scheduled:as a scheduled task|asap:as a scheduled task and when media is added"
	Group      string `xml:"group,attr,omitempty" json:"group,omitempty"`           // "channels"
	Hidden     int    `xml:"hidden,attr,omitempty" json:"hidden,omitempty"`         // "0"
	Id         string `xml:"id,attr,omitempty" json:"id,omitempty"`                 // "iPhotoLibraryXmlPath"
	Label      string `xml:"label,attr,omitempty" json:"label,omitempty"`           // "iPhoto library XML path"
	Summary    string `xml:"summary,attr,omitempty" json:"summary,omitempty"`       // ""
	Type       string `xml:"type,attr,omitempty" json:"type,omitempty"`             // "text"
	Value      string `xml:"value,attr,omitempty" json:"value,omitempty"`           // ""
}

type Location struct {
	LibrarySectionID int    `xml:"librarySectionID,attr,omitempty" json:"librarySectionID,omitempty"` // "4"
	Path             string `xml:"path,attr,omitempty" json:"path,omitempty"`                         // "/cortex/central/videos/lib/TV/The A-Team"
}

type Directory struct {
	AddedAt               int        `xml:"addedAt,attr,omitempty" json:"addedAt,omitempty"` // "1471546619"
	Agent                 string     `xml:"agent,attr,omitempty" json:"agent,omitempty"`
	AllowSync             int        `xml:"allowSync,attr,omitempty" json:"allowSync,omitempty"`   // "1"
	Art                   string     `xml:"art,attr,omitempty" json:"art,omitempty"`               // "/library/metadata/13814/art/1471546640"
	Banner                string     `xml:"banner,attr,omitempty" json:"banner,omitempty"`         // "/library/metadata/13814/banner/1471546640"
	ChildCount            int        `xml:"childCount,attr,omitempty" json:"childCount,omitempty"` // "5"
	Composite             string     `xml:"composite,attr,omitempty" json:"composite,omitempty"`
	ContentRating         string     `xml:"contentRating,attr,omitempty" json:"contentRating,omitempty"` // "TV-PG"
	CreatedAt             int        `xml:"createdAt,attr,omitempty" json:"createdAt,omitempty"`         // "1471546640"
	Duration              int        `xml:"duration,attr,omitempty" json:"duration,omitempty"`           // "2700000"
	Filters               int        `xml:"filters,attr,omitempty" json:"filters,omitempty"`
	Guid                  string     `xml:"guid,attr,omitempty" json:"guid,omitempty"`   // "com.plexapp.agents.thetvdb://77904"
	Index                 int        `xml:"index,attr,omitempty" json:"index,omitempty"` // "3"
	Key                   string     `xml:"key,attr,omitempty" json:"key,omitempty"`     // "/library/metadata/20448/children"
	Language              string     `xml:"language,attr,omitempty" json:"language,omitempty"`
	LeafCount             int        `xml:"leafCount,attr,omitempty" json:"leafCount,omitempty"`               // "98"
	LibrarySectionID      int        `xml:"librarySectionID,attr,omitempty" json:"librarySectionID,omitempty"` // "4"
	Locations             []Location `xml:"Location" json:"Location"`
	OriginallyAvailableAt string     `xml:"originallyAvailableAt,omitempty" json:"originallyAvailableAt,omitempty"` // "1983-01-23"
	ParentIndex           int        `xml:"parentIndex,attr,omitempty" json:"parentIndex,omitempty"`                // "1"
	ParentKey             string     `xml:"parentKey,attr,omitempty" json:"parentKey,omitempty"`                    // "/library/metadata/14207"
	ParentRatingKey       int        `xml:"parentRatingKey,attr,omitempty" json:"parentRatingKey,omitempty"`        // "14207"
	ParentTheme           string     `xml:"parentTheme,attr,omitempty" json:"parentTheme,omitempty"`                // "/library/metadata/14207/theme/1471544783"
	ParentThumb           string     `xml:"parentThumb,attr,omitempty" json:"parentThumb,omitempty"`                // "/library/metadata/14207/thumb/1471544783"
	ParentTitle           string     `xml:"parentTitle,attr,omitempty" json:"parentTitle,omitempty"`                // "The O.C."
	Rating                float32    `xml:"rating,attr,omitempty" json:"rating,omitempty"`                          // "7.8"
	RatingKey             int        `xml:"ratingKey,attr,omitempty" json:"ratingKey,omitempty"`                    // "20448"
	Refreshing            int        `xml:"refreshing,attr,omitempty" json:"refreshing,omitempty"`
	Scanner               string     `xml:"scanner,attr,omitempty" json:"scanner,omitempty"`
	Studio                string     `xml:"studio,attr,omitempty" json:"studio,omitempty"`       // "NBC"
	Summary               string     `xml:"summary,attr,omitempty" json:"summary,omitempty"`     // "The A-Team is about a group of ex-United States Army Special Forces personnel who work as soldiers of fortune, while on the run from the Army after being branded as war criminals for a crime they didn't commit."
	Theme                 string     `xml:"theme,attr,omitempty" json:"theme,omitempty"`         // "/library/metadata/13814/theme/1471546640"
	Thumb                 string     `xml:"thumb,attr,omitempty" json:"thumb,omitempty"`         // "/library/metadata/13814/thumb/1471546640"
	Title                 string     `xml:"title,attr,omitempty" json:"title,omitempty"`         // "The A-Team"
	TitleSort             string     `xml:"titleSort,attr,omitempty" json:"titleSort,omitempty"` // "A-Team"
	Type                  string     `xml:"type,attr,omitempty" json:"type,omitempty"`           // "show"
	UpdatedAt             int        `xml:"updatedAt,attr,omitempty" json:"updatedAt,omitempty"` // "1471546640"
	UUID                  string     `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
	ViewedLeafCount       int        `xml:"viewedLeafCount,attr,omitempty" json:"viewedLeafCount,omitempty"` // "52"
	Year                  int        `xml:"year,attr,omitempty" json:"year,omitempty"`                       // "1983"
}

type Hub struct {
	Videos        []Video     `xml:"Video" json:"Video"`
	Directories   []Directory `xml:"Directory" json:"Directory"`
	Key           string      `xml:"key,omitempty" json:"key,omitempty"`                          // "/hubs/home/recentlyAdded?type=1&amp;personal=1"
	Type          string      `xml:"type,attr,omitempty" json:"type,omitempty"`                   // "clip"
	HubIdentifier string      `xml:"hubIdentifier,attr,omitempty" json:"hubIdentifier,omitempty"` // "home.videos.recent"
	Size          int         `xml:"size,attr,omitempty" json:"size,omitempty"`                   // "0"
	Title         string      `xml:"title,attr,omitempty" json:"title,omitempty"`                 // "Recently Added Home Videos"
	More          int         `xml:"more,attr,omitempty" json:"more,omitempty"`                   // "0"
}

type MediaContainer struct {
	AllowSync                     int         `xml:"allowSync,attr,omitempty" json:"allowSync,omitempty"` // "1"
	Art                           string      `xml:"art,attr,omitempty" json:"art,omitempty"`
	Directories                   []Directory `xml:"Directory" json:"Directory"`
	Hubs                          []Hub       `xml:"Hub" json:"Hub"`
	Identifier                    string      `xml:"identifier,omitempty" json:"identifier,omitempty"`                  // "com.plexapp.plugins.library"
	LibrarySectionID              int         `xml:"librarySectionID,attr,omitempty" json:"librarySectionID,omitempty"` // "4"
	LibrarySectionTitle           string      `xml:"librarySectionTitle,attr,omitempty" json:"librarySectionTitle,omitempty"`
	LibrarySectionUUID            string      `xml:"librarySectionUUID,attr,omitempty" json:"librarySectionUUID,omitempty"`
	MediaTagPrefix                string      `xml:"mediaTagPrefix,attr,omitempty" json:"mediaTagPrefix,omitempty"`   // "/system/bundle/media/flags/"
	MediaTagVersion               int         `xml:"mediaTagVersion,attr,omitempty" json:"mediaTagVersion,omitempty"` // "1466734815"
	NoCache                       int         `xml:"nocache,attr,omitempty" json:"nocache,omitempty"`
	Offset                        int         `xml:"offset,attr,omitempty" json:"offset,omitempty"`
	Settings                      []Setting   `xml:"Setting" json:"Setting"`
	Size                          int         `xml:"size,omitempty" json:"size,omitempty"` // "8"
	SortAscending                 int         `xml:"sortAsc,attr,omitempty" json:"sortAsc,omitempty"`
	Thumb                         string      `xml:"thumb,attr,omitempty" json:"thumb,omitempty"`
	Title1                        string      `xml:"title1,attr,omitempty" json:"title1,omitempty"` // "Plex Library"
	Title2                        string      `xml:"title2,attr,omitempty" json:"title2,omitempty"` // "Plex Library"
	TotalSize                     int         `xml:"totalSize,attr,omitempty" json:"totalSize,omitempty"`
	Videos                        []Video     `xml:"Video" json:"Video"`
	ViewGroup                     string      `xml:"viewGroup,omitempty" json:"viewGroup,omitempty"`
	ViewMode                      int         `xml:"viewMode,attr,omitempty" json:"viewMode,omitempty"`
	AllowCameraUpload             int         `xml:"allowCameraUpload,attr,omitempty" json:"allowCameraUpload,omitempty"`                         // "1"
	AllowChannelAccess            int         `xml:"allowChannelAccess,attr,omitempty" json:"allowChannelAccess,omitempty"`                       // "1"
	AllowSharing                  int         `xml:"allowSharing,attr,omitempty" json:"allowSharing,omitempty"`                                   // "1"
	BackgroundProcessing          int         `xml:"backgroundProcessing,attr,omitempty" json:"backgroundProcessing,omitempty"`                   // "1"
	Certificate                   int         `xml:"certificate,attr,omitempty" json:"certificate,omitempty"`                                     // "1"
	CompanionProxy                int         `xml:"companionProxy,attr,omitempty" json:"companionProxy,omitempty"`                               // "1"
	EventStream                   int         `xml:"eventStream,attr,omitempty" json:"eventStream,omitempty"`                                     // "1"
	FriendlyName                  string      `xml:"friendlyName,attr,omitempty" json:"friendlyName,omitempty"`                                   // "Cortex"
	HubSearch                     int         `xml:"hubSearch,attr,omitempty" json:"hubSearch,omitempty"`                                         // "1"
	MachineIdentifier             string      `xml:"machineIdentifier,attr,omitempty" json:"machineIdentifier,omitempty"`                         // "91dbc3a3ec55607b993cb1ec6b10983f9f4bb407"
	Multiuser                     int         `xml:"multiuser,attr,omitempty" json:"multiuser,omitempty"`                                         // "1"
	MyPlex                        int         `xml:"myPlex,attr,omitempty" json:"myPlex,omitempty"`                                               // "1"
	MyPlexMappingState            string      `xml:"myPlexMappingState,attr,omitempty" json:"myPlexMappingState,omitempty"`                       // "mapped"
	MyPlexSigninState             string      `xml:"myPlexSigninState,attr,omitempty" json:"myPlexSigninState,omitempty"`                         // "ok"
	MyPlexSubscription            int         `xml:"myPlexSubscription,attr,omitempty" json:"myPlexSubscription,omitempty"`                       // "1"
	MyPlexUsername                string      `xml:"myPlexUsername,attr,omitempty" json:"myPlexUsername,omitempty"`                               // "garyhetzel@gmail.com"
	OwnerFeatures                 string      `xml:"ownerFeatures,attr,omitempty" json:"ownerFeatures,omitempty"`                                 // "camera_upload,cloudsync,content_filter,home,lyrics,music_videos,pass,premium_music_metadata,sync,trailers"
	Platform                      string      `xml:"platform,attr,omitempty" json:"platform,omitempty"`                                           // "FreeBSD"
	PlatformVersion               string      `xml:"platformVersion,attr,omitempty" json:"platformVersion,omitempty"`                             // "10.3-RELEASE (FreeBSD 10.3-RELEASE #0 r297264: Fri Mar 25 02:10:02 UTC 2016 root@releng1.nyi.freebsd.org:/usr/obj/usr/src/sys/GENERIC)"
	PluginHost                    int         `xml:"pluginHost,attr,omitempty" json:"pluginHost,omitempty"`                                       // "1"
	ReadOnlyLibraries             int         `xml:"readOnlyLibraries,attr,omitempty" json:"readOnlyLibraries,omitempty"`                         // "0"
	RequestParametersInCookie     int         `xml:"requestParametersInCookie,attr,omitempty" json:"requestParametersInCookie,omitempty"`         // "1"
	Sync                          int         `xml:"sync,attr,omitempty" json:"sync,omitempty"`                                                   // "1"
	TranscoderActiveVideoSessions int         `xml:"transcoderActiveVideoSessions,attr,omitempty" json:"transcoderActiveVideoSessions,omitempty"` // "0"
	TranscoderAudio               int         `xml:"transcoderAudio,attr,omitempty" json:"transcoderAudio,omitempty"`                             // "1"
	TranscoderLyrics              int         `xml:"transcoderLyrics,attr,omitempty" json:"transcoderLyrics,omitempty"`                           // "1"
	TranscoderPhoto               int         `xml:"transcoderPhoto,attr,omitempty" json:"transcoderPhoto,omitempty"`                             // "1"
	TranscoderSubtitles           int         `xml:"transcoderSubtitles,attr,omitempty" json:"transcoderSubtitles,omitempty"`                     // "1"
	TranscoderVideo               int         `xml:"transcoderVideo,attr,omitempty" json:"transcoderVideo,omitempty"`                             // "1"
	TranscoderVideoBitrates       string      `xml:"transcoderVideoBitrates,attr,omitempty" json:"transcoderVideoBitrates,omitempty"`             // "64,96,208,320,720,1500,2000,3000,4000,8000,10000,12000,20000"
	TranscoderVideoQualities      string      `xml:"transcoderVideoQualities,attr,omitempty" json:"transcoderVideoQualities,omitempty"`           // "0,1,2,3,4,5,6,7,8,9,10,11,12"
	TranscoderVideoResolutions    string      `xml:"transcoderVideoResolutions,attr,omitempty" json:"transcoderVideoResolutions,omitempty"`       // "128,128,160,240,320,480,768,720,720,1080,1080,1080,1080"
	UpdatedAt                     int         `xml:"updatedAt,attr,omitempty" json:"updatedAt,omitempty"`                                         // "1471466714"
	Updater                       int         `xml:"updater,attr,omitempty" json:"updater,omitempty"`                                             // "1"
	Version                       string      `xml:"version,attr,omitempty" json:"version,omitempty"`                                             // "1.0.0.2261-a17e99e"
}

type User struct {
	Id    int    `xml:"id,attr,omitempty" json:"id,omitempty"`
	Title string `xml:"title,attr,omitempty" json:"title,omitempty"`
}

type Player struct {
	Address           string `xml:"address,attr,omitempty" json:"address,omitempty"`                     // "73.78.220.221"
	Device            string `xml:"device,attr,omitempty" json:"device,omitempty"`                       // "Windows"
	MachineIdentifier string `xml:"machineIdentifier,attr,omitempty" json:"machineIdentifier,omitempty"` // "4648a824-5ef2-4e40-a044-86fdd6b76bac"
	Model             string `xml:"model,attr,omitempty" json:"model,omitempty"`                         // ""
	Platform          string `xml:"platform,attr,omitempty" json:"platform,omitempty"`                   // "Chrome"
	PlatformVersion   string `xml:"platformVersion,attr,omitempty" json:"platformVersion,omitempty"`     // "52.0"
	Product           string `xml:"product,attr,omitempty" json:"product,omitempty"`                     // "Plex Web"
	Profile           string `xml:"profile,attr,omitempty" json:"profile,omitempty"`                     // "Web"
	State             string `xml:"state,attr,omitempty" json:"state,omitempty"`                         // "playing"
	Title             string `xml:"title,attr,omitempty" json:"title,omitempty"`                         // "Plex Web (Chrome)"
	Vendor            string `xml:"vendor,attr,omitempty" json:"vendor,omitempty"`                       // ""
	Version           string `xml:"version,attr,omitempty" json:"version,omitempty"`                     // "2.7.8"
}

type TranscodeSession struct {
	AudioChannels int     `xml:"audioChannels,attr,omitempty" json:"audioChannels"` // "2"
	AudioCodec    string  `xml:"audioCodec,attr,omitempty" json:"audioCodec"`       // "mp3"
	AudioDecision string  `xml:"audioDecision,attr,omitempty" json:"audioDecision"` // "copy"
	Complete      int     `xml:"complete,attr,omitempty" json:"complete"`           // "0"
	Container     string  `xml:"container,attr,omitempty" json:"container"`         // "mkv"
	Context       string  `xml:"context,attr,omitempty" json:"context"`             // "streaming"
	Duration      int     `xml:"duration,attr,omitempty" json:"duration"`           // "3717000"
	Height        int     `xml:"height,attr,omitempty" json:"height"`               // "352"
	Key           string  `xml:"key,attr,omitempty" json:"key"`                     // "fs3uwoq9w1yv7o5bj0d15g66r"
	Progress      float64 `xml:"progress,attr,omitempty" json:"progress"`           // "26.899999618530273"
	Protocol      string  `xml:"protocol,attr,omitempty" json:"protocol"`           // "http"
	Remaining     int     `xml:"remaining,attr,omitempty" json:"remaining"`         // "19543"
	Speed         int     `xml:"speed,attr,omitempty" json:"speed"`                 // "0"
	Throttled     int     `xml:"throttled,attr,omitempty" json:"throttled"`         // "1"
	VideoCodec    string  `xml:"videoCodec,attr,omitempty" json:"videoCodec"`       // "h264"
	VideoDecision string  `xml:"videoDecision,attr,omitempty" json:"videoDecision"` // "transcode"
	Width         int     `xml:"width,attr,omitempty" json:"width"`                 // "640"
}
