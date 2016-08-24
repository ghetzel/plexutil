package client

type Setting struct {
	Advanced   int    `xml:"advanced,attr,omitempty"`   // "1"
	Default    string `xml:"default,attr,omitempty"`    // ""
	EnumValues string `xml:"enumValues,attr,omitempty"` // "never:never|scheduled:as a scheduled task|asap:as a scheduled task and when media is added"
	Group      string `xml:"group,attr,omitempty"`      // "channels"
	Hidden     int    `xml:"hidden,attr,omitempty"`     // "0"
	Id         string `xml:"id,attr,omitempty"`         // "iPhotoLibraryXmlPath"
	Label      string `xml:"label,attr,omitempty"`      // "iPhoto library XML path"
	Summary    string `xml:"summary,attr,omitempty"`    // ""
	Type       string `xml:"type,attr,omitempty"`       // "text"
	Value      string `xml:"value,attr,omitempty"`      // ""
}

type Stream struct {
	BitDepth           int     `xml:"bitDepth,attr,omitempty"`           // "8"
	Bitrate            int     `xml:"bitrate,attr,omitempty"`            // "1375"
	Cabac              int     `xml:"cabac,attr,omitempty"`              // "1"
	ChromaSubsampling  string  `xml:"chromaSubsampling,attr,omitempty"`  // "4:2:0"
	Codec              string  `xml:"codec,attr,omitempty"`              // "h264"
	CodecID            string  `xml:"codecID,attr,omitempty"`            // "V_MPEG4/ISO/AVC"
	Default            int     `xml:"default,attr,omitempty"`            // "1"
	Duration           int     `xml:"duration,attr,omitempty"`           // "1451326"
	FrameRate          float32 `xml:"frameRate,attr,omitempty"`          // "23.976"
	FrameRateMode      string  `xml:"frameRateMode,attr,omitempty"`      // "cfr"
	HasScalingMatrix   int     `xml:"hasScalingMatrix,attr,omitempty"`   // "0"
	Height             int     `xml:"height,attr,omitempty"`             // "358"
	Id                 int     `xml:"id,attr,omitempty"`                 // "291338"
	Index              int     `xml:"index,attr,omitempty"`              // "0"
	Level              int     `xml:"level,attr,omitempty"`              // "31"
	PixelFormat        string  `xml:"pixelFormat,attr,omitempty"`        // "yuv420p"
	Profile            string  `xml:"profile,attr,omitempty"`            // "high"
	RefFrames          int     `xml:"refFrames,attr,omitempty"`          // "5"
	RequiredBandwidths string  `xml:"requiredBandwidths,attr,omitempty"` // "1830,1830,1830,1830,1830,1830,1830,1830"
	ScanType           string  `xml:"scanType,attr,omitempty"`           // "progressive"
	StreamType         int     `xml:"streamType,attr,omitempty"`         // "1"
	Width              int     `xml:"width,attr,omitempty"`              // "640"
}

type Part struct {
	AudioProfile        string   `xml:"audioProfile,attr,omitempty"`        // "lc"
	Container           string   `xml:"container,attr,omitempty"`           // "mkv"
	DeepAnalysisVersion int      `xml:"deepAnalysisVersion,attr,omitempty"` // "1"
	Duration            int      `xml:"duration,attr,omitempty"`            // "1451337"
	File                string   `xml:"file,attr,omitempty"`                // "/cortex/central/videos/lib/TV/The Mind of a Chef/S01E12 - Fresh.mkv"
	Id                  int      `xml:"id,attr,omitempty"`                  // "154993"
	Indexes             string   `xml:"indexes,attr,omitempty"`             // "sd"
	Key                 string   `xml:"key,attr,omitempty"`                 // "/library/parts/154993/1393015691/file.mkv"
	RequiredBandwidths  string   `xml:"requiredBandwidths,attr,omitempty"`  // "1947,1947,1947,1947,1947,1947,1947,1947"
	Size                int      `xml:"size,attr,omitempty"`                // "249516275"
	Streams             []Stream `xml:"Stream"`
	VideoProfile        string   `xml:"videoProfile,attr,omitempty"` // "high"
}

type Media struct {
	AspectRatio     float32 `xml:"aspectRatio,attr,omitempty"`   // "1.78"
	AudioChannels   int     `xml:"audioChannels,attr,omitempty"` // "6"
	AudioCodec      string  `xml:"audioCodec,attr,omitempty"`    // "ac3"
	Bitrate         int     `xml:"bitrate,attr,omitempty"`       // "4322"
	Container       string  `xml:"container,attr,omitempty"`     // "mkv"
	Duration        int     `xml:"duration,attr,omitempty"`      // "1260427"
	Height          int     `xml:"height,attr,omitempty"`        // "720"
	Id              int     `xml:"id,attr,omitempty"`            // "77758"
	Parts           []Part  `xml:"Part"`
	VideoCodec      string  `xml:"videoCodec,attr,omitempty"`      // "h264"
	VideoFrameRate  string  `xml:"videoFrameRate,attr,omitempty"`  // "24p"
	VideoProfile    string  `xml:"videoProfile,attr,omitempty"`    // "high"
	VideoResolution int     `xml:"videoResolution,attr,omitempty"` // "720"
	Width           int     `xml:"width,attr,omitempty"`           // "1280"
}

type Chapter struct {
	Id              int    `xml:"id,attr,omitempty"`
	Tag             string `xml:"tag,attr,omitempty"`
	Index           int    `xml:"index,attr,omitempty"`
	StartTimeOffset int    `xml:"startTimeOffset,attr,omitempty"`
	EndTimeOffset   int    `xml:"endTimeOffset,attr,omitempty"`
}

type Role struct {
	Id    int    `xml:"id,attr,omitempty"`
	Tag   string `xml:"tag,attr,omitempty"`
	Role  string `xml:"role,attr,omitempty"`
	Count int    `xml:"count,attr,omitempty"`
}

type Identifier struct {
	Id    int    `xml:"id,attr,omitempty"`
	Tag   string `xml:"tag,attr,omitempty"`
	Count int    `xml:"count,attr,omitempty"`
}

type Video struct {
	AddedAt               int              `xml:"addedAt,attr,omitempty"`             // "1456373737"
	Art                   string           `xml:"art,attr,omitempty"`                 // "/library/metadata/21666/art/1461210754"
	AudienceRating        float32          `xml:"audienceRating,attr,omitempty"`      // "7.5"
	AudienceRatingImage   string           `xml:"audienceRatingImage,attr,omitempty"` // "rottentomatoes://image.rating.upright"
	Chapters              []Chapter        `xml:"Chapter"`
	ChapterSource         string           `xml:"chapterSource,attr,omitempty"` // ""
	ContentRating         string           `xml:"contentRating,attr,omitempty"` // "TV-MA"
	Country               Identifier       `xml:"Country"`
	Directors             []Identifier     `xml:"Director"`
	Duration              int              `xml:"duration,attr,omitempty"` // "1260427"
	Extras                []MediaContainer `xml:"Extras"`
	Genres                []Identifier     `xml:"Genre"`
	GrandparentArt        string           `xml:"grandparentArt,attr,omitempty"`       // "/library/metadata/21666/art/1461210754"
	GrandparentKey        string           `xml:"grandparentKey,attr,omitempty"`       // "/library/metadata/21666"
	GrandparentRatingKey  int              `xml:"grandparentRatingKey,attr,omitempty"` // "21666"
	GrandparentTheme      string           `xml:"grandparentTheme,attr,omitempty"`     // "/library/metadata/21666/theme/1461210754"
	GrandparentThumb      string           `xml:"grandparentThumb,attr,omitempty"`     // "/library/metadata/21666/thumb/1461210754"
	GrandparentTitle      string           `xml:"grandparentTitle,attr,omitempty"`     // "Broad City"
	Index                 int              `xml:"index,attr,omitempty"`                // "2"
	Key                   string           `xml:"key,attr,omitempty"`                  // "/library/metadata/60627"
	LastViewedAt          int              `xml:"lastViewedAt,attr,omitempty"`         // "1471313734"
	LibrarySectionID      int              `xml:"librarySectionID,attr,omitempty"`     // "4"
	Media                 Media            `xml:"Media"`
	OriginallyAvailableAt string           `xml:"originallyAvailableAt,attr,omitempty"` // "2016-02-24"
	ParentIndex           int              `xml:"parentIndex,attr,omitempty"`           // "3"
	ParentKey             string           `xml:"parentKey,attr,omitempty"`             // "/library/metadata/60207"
	ParentRatingKey       int              `xml:"parentRatingKey,attr,omitempty"`       // "60207"
	PrimaryExtraKey       string           `xml:"primaryExtraKey,attr,omitempty"`       // "/library/metadata/99398"
	Producers             []Identifier     `xml:"Producer"`
	Rating                float32          `xml:"rating,attr,omitempty"`      // "8.0"
	RatingImage           string           `xml:"ratingImage,attr,omitempty"` // "rottentomatoes://image.rating.certified"
	RatingKey             int              `xml:"ratingKey,attr,omitempty"`   // "60627"
	Related               []MediaContainer `xml:"Related"`
	Roles                 []Role           `xml:"Role"`
	Studio                string           `xml:"studio,attr,omitempty"`     // "United Artists"
	Summary               string           `xml:"summary,attr,omitempty"`    // "Abbi pretends to be Ilana to cover her shift at the food co-op, while Ilana goes to an important doctor's appointment on Long Island."
	Tagline               string           `xml:"tagline,attr,omitempty"`    // "Is it a game, or is it real?"
	Thumb                 string           `xml:"thumb,attr,omitempty"`      // "/library/metadata/60627/thumb/1470739012"
	Title                 string           `xml:"title,attr,omitempty"`      // "Co-Op"
	Type                  string           `xml:"type,attr,omitempty"`       // "episode"
	UpdatedAt             int              `xml:"updatedAt,attr,omitempty"`  // "1470739012"
	ViewCount             int              `xml:"viewCount,attr,omitempty"`  // "1"
	ViewOffset            int              `xml:"viewOffset,attr,omitempty"` // "102000"
	Writers               []Identifier     `xml:"Writer"`
	Year                  int              `xml:"year,attr,omitempty"` // "2016"
}

type Location struct {
	LibrarySectionID int    `xml:"librarySectionID,attr,omitempty"` // "4"
	Path             string `xml:"path,attr,omitempty"`             // "/cortex/central/videos/lib/TV/The A-Team"
}

type Directory struct {
	AddedAt               int        `xml:"addedAt,attr,omitempty"` // "1471546619"
	Agent                 string     `xml:"agent,attr,omitempty"`
	AllowSync             int        `xml:"allowSync,attr,omitempty"`  // "1"
	Art                   string     `xml:"art,attr,omitempty"`        // "/library/metadata/13814/art/1471546640"
	Banner                string     `xml:"banner,attr,omitempty"`     // "/library/metadata/13814/banner/1471546640"
	ChildCount            int        `xml:"childCount,attr,omitempty"` // "5"
	Composite             string     `xml:"composite,attr,omitempty"`
	ContentRating         string     `xml:"contentRating,attr,omitempty"` // "TV-PG"
	CreatedAt             int        `xml:"createdAt,attr,omitempty"`     // "1471546640"
	Duration              int        `xml:"duration,attr,omitempty"`      // "2700000"
	Filters               int        `xml:"filters,attr,omitempty"`
	Guid                  string     `xml:"guid,attr,omitempty"`  // "com.plexapp.agents.thetvdb://77904"
	Index                 int        `xml:"index,attr,omitempty"` // "3"
	Key                   string     `xml:"key,attr,omitempty"`   // "/library/metadata/20448/children"
	Language              string     `xml:"language,attr,omitempty"`
	LeafCount             int        `xml:"leafCount,attr,omitempty"`        // "98"
	LibrarySectionID      int        `xml:"librarySectionID,attr,omitempty"` // "4"
	Locations             []Location `xml:"Location"`
	OriginallyAvailableAt string     `xml:"originallyAvailableAt,attr,omitempty"` // "1983-01-23"
	ParentIndex           int        `xml:"parentIndex,attr,omitempty"`           // "1"
	ParentKey             string     `xml:"parentKey,attr,omitempty"`             // "/library/metadata/14207"
	ParentRatingKey       int        `xml:"parentRatingKey,attr,omitempty"`       // "14207"
	ParentTheme           string     `xml:"parentTheme,attr,omitempty"`           // "/library/metadata/14207/theme/1471544783"
	ParentThumb           string     `xml:"parentThumb,attr,omitempty"`           // "/library/metadata/14207/thumb/1471544783"
	ParentTitle           string     `xml:"parentTitle,attr,omitempty"`           // "The O.C."
	Rating                float32    `xml:"rating,attr,omitempty"`                // "7.8"
	RatingKey             int        `xml:"ratingKey,attr,omitempty"`             // "20448"
	Refreshing            int        `xml:"refreshing,attr,omitempty"`
	Scanner               string     `xml:"scanner,attr,omitempty"`
	Studio                string     `xml:"studio,attr,omitempty"`    // "NBC"
	Summary               string     `xml:"summary,attr,omitempty"`   // "The A-Team is about a group of ex-United States Army Special Forces personnel who work as soldiers of fortune, while on the run from the Army after being branded as war criminals for a crime they didn't commit."
	Theme                 string     `xml:"theme,attr,omitempty"`     // "/library/metadata/13814/theme/1471546640"
	Thumb                 string     `xml:"thumb,attr,omitempty"`     // "/library/metadata/13814/thumb/1471546640"
	Title                 string     `xml:"title,attr,omitempty"`     // "The A-Team"
	TitleSort             string     `xml:"titleSort,attr,omitempty"` // "A-Team"
	Type                  string     `xml:"type,attr,omitempty"`      // "show"
	UpdatedAt             int        `xml:"updatedAt,attr,omitempty"` // "1471546640"
	UUID                  string     `xml:"uuid,attr,omitempty"`
	ViewedLeafCount       int        `xml:"viewedLeafCount,attr,omitempty"` // "52"
	Year                  int        `xml:"year,attr,omitempty"`            // "1983"
}

type Hub struct {
	Videos        []Video     `xml:"Video"`
	Directories   []Directory `xml:"Directory"`
	Key           string      `xml:"key,attr,omitempty"`           // "/hubs/home/recentlyAdded?type=1&amp;personal=1"
	Type          string      `xml:"type,attr,omitempty"`          // "clip"
	HubIdentifier string      `xml:"hubIdentifier,attr,omitempty"` // "home.videos.recent"
	Size          int         `xml:"size,attr,omitempty"`          // "0"
	Title         string      `xml:"title,attr,omitempty"`         // "Recently Added Home Videos"
	More          int         `xml:"more,attr,omitempty"`          // "0"
}

type MediaContainer struct {
	AllowSync                     int         `xml:"allowSync,attr,omitempty"` // "1"
	Art                           string      `xml:"art,attr,omitempty"`
	Directories                   []Directory `xml:"Directory"`
	Hubs                          []Hub       `xml:"Hub"`
	Identifier                    string      `xml:"identifier,attr,omitempty"`       // "com.plexapp.plugins.library"
	LibrarySectionID              int         `xml:"librarySectionID,attr,omitempty"` // "4"
	LibrarySectionTitle           string      `xml:"librarySectionTitle,attr,omitempty"`
	LibrarySectionUUID            string      `xml:"librarySectionUUID,attr,omitempty"`
	MediaTagPrefix                string      `xml:"mediaTagPrefix,attr,omitempty"`  // "/system/bundle/media/flags/"
	MediaTagVersion               int         `xml:"mediaTagVersion,attr,omitempty"` // "1466734815"
	NoCache                       int         `xml:"nocache,attr,omitempty"`
	Offset                        int         `xml:"offset,attr,omitempty"`
	Settings                      []Setting   `xml:"Setting"`
	Size                          int         `xml:"size,attr,omitempty"` // "8"
	SortAscending                 int         `xml:"sortAsc,attr,omitempty"`
	Thumb                         string      `xml:"thumb,attr,omitempty"`
	Title1                        string      `xml:"title1,attr,omitempty"` // "Plex Library"
	Title2                        string      `xml:"title2,attr,omitempty"` // "Plex Library"
	TotalSize                     int         `xml:"totalSize,attr,omitempty"`
	Videos                        []Video     `xml:"Video"`
	ViewGroup                     string      `xml:"viewGroup,attr,omitempty"`
	ViewMode                      int         `xml:"viewMode,attr,omitempty"`
	AllowCameraUpload             int         `xml:"allowCameraUpload,attr,omitempty"`             // "1"
	AllowChannelAccess            int         `xml:"allowChannelAccess,attr,omitempty"`            // "1"
	AllowSharing                  int         `xml:"allowSharing,attr,omitempty"`                  // "1"
	BackgroundProcessing          int         `xml:"backgroundProcessing,attr,omitempty"`          // "1"
	Certificate                   int         `xml:"certificate,attr,omitempty"`                   // "1"
	CompanionProxy                int         `xml:"companionProxy,attr,omitempty"`                // "1"
	EventStream                   int         `xml:"eventStream,attr,omitempty"`                   // "1"
	FriendlyName                  string      `xml:"friendlyName,attr,omitempty"`                  // "Cortex"
	HubSearch                     int         `xml:"hubSearch,attr,omitempty"`                     // "1"
	MachineIdentifier             string      `xml:"machineIdentifier,attr,omitempty"`             // "91dbc3a3ec55607b993cb1ec6b10983f9f4bb407"
	Multiuser                     int         `xml:"multiuser,attr,omitempty"`                     // "1"
	MyPlex                        int         `xml:"myPlex,attr,omitempty"`                        // "1"
	MyPlexMappingState            string      `xml:"myPlexMappingState,attr,omitempty"`            // "mapped"
	MyPlexSigninState             string      `xml:"myPlexSigninState,attr,omitempty"`             // "ok"
	MyPlexSubscription            int         `xml:"myPlexSubscription,attr,omitempty"`            // "1"
	MyPlexUsername                string      `xml:"myPlexUsername,attr,omitempty"`                // "garyhetzel@gmail.com"
	OwnerFeatures                 string      `xml:"ownerFeatures,attr,omitempty"`                 // "camera_upload,cloudsync,content_filter,home,lyrics,music_videos,pass,premium_music_metadata,sync,trailers"
	Platform                      string      `xml:"platform,attr,omitempty"`                      // "FreeBSD"
	PlatformVersion               string      `xml:"platformVersion,attr,omitempty"`               // "10.3-RELEASE (FreeBSD 10.3-RELEASE #0 r297264: Fri Mar 25 02:10:02 UTC 2016 root@releng1.nyi.freebsd.org:/usr/obj/usr/src/sys/GENERIC)"
	PluginHost                    int         `xml:"pluginHost,attr,omitempty"`                    // "1"
	ReadOnlyLibraries             int         `xml:"readOnlyLibraries,attr,omitempty"`             // "0"
	RequestParametersInCookie     int         `xml:"requestParametersInCookie,attr,omitempty"`     // "1"
	Sync                          int         `xml:"sync,attr,omitempty"`                          // "1"
	TranscoderActiveVideoSessions int         `xml:"transcoderActiveVideoSessions,attr,omitempty"` // "0"
	TranscoderAudio               int         `xml:"transcoderAudio,attr,omitempty"`               // "1"
	TranscoderLyrics              int         `xml:"transcoderLyrics,attr,omitempty"`              // "1"
	TranscoderPhoto               int         `xml:"transcoderPhoto,attr,omitempty"`               // "1"
	TranscoderSubtitles           int         `xml:"transcoderSubtitles,attr,omitempty"`           // "1"
	TranscoderVideo               int         `xml:"transcoderVideo,attr,omitempty"`               // "1"
	TranscoderVideoBitrates       string      `xml:"transcoderVideoBitrates,attr,omitempty"`       // "64,96,208,320,720,1500,2000,3000,4000,8000,10000,12000,20000"
	TranscoderVideoQualities      string      `xml:"transcoderVideoQualities,attr,omitempty"`      // "0,1,2,3,4,5,6,7,8,9,10,11,12"
	TranscoderVideoResolutions    string      `xml:"transcoderVideoResolutions,attr,omitempty"`    // "128,128,160,240,320,480,768,720,720,1080,1080,1080,1080"
	UpdatedAt                     int         `xml:"updatedAt,attr,omitempty"`                     // "1471466714"
	Updater                       int         `xml:"updater,attr,omitempty"`                       // "1"
	Version                       string      `xml:"version,attr,omitempty"`                       // "1.0.0.2261-a17e99e"
}
