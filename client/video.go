package client

type Stream struct {
	BitDepth           int     `xml:"bitDepth,attr,omitempty" json:"BitDepth,omitempty"`                     // "8"
	Bitrate            int     `xml:"bitrate,attr,omitempty" json:"Bitrate,omitempty"`                       // "1375"
	Cabac              int     `xml:"cabac,attr,omitempty" json:"Cabac,omitempty"`                           // "1"
	ChromaSubsampling  string  `xml:"chromaSubsampling,attr,omitempty" json:"ChromaSubsampling,omitempty"`   // "4:2:0"
	Codec              string  `xml:"codec,attr,omitempty" json:"Codec,omitempty"`                           // "h264"
	CodecID            string  `xml:"codecID,attr,omitempty" json:"CodecID,omitempty"`                       // "V_MPEG4/ISO/AVC"
	Default            int     `xml:"default,attr,omitempty" json:"Default,omitempty"`                       // "1"
	Duration           int     `xml:"duration,attr,omitempty" json:"Duration,omitempty"`                     // "1451326"
	FrameRate          float32 `xml:"frameRate,attr,omitempty" json:"FrameRate,omitempty"`                   // "23.976"
	FrameRateMode      string  `xml:"frameRateMode,attr,omitempty" json:"FrameRateMode,omitempty"`           // "cfr"
	HasScalingMatrix   int     `xml:"hasScalingMatrix,attr,omitempty" json:"HasScalingMatrix,omitempty"`     // "0"
	Height             int     `xml:"height,attr,omitempty" json:"Height,omitempty"`                         // "358"
	Id                 int     `xml:"id,attr,omitempty" json:"Id,omitempty"`                                 // "291338"
	Index              int     `xml:"index,attr,omitempty" json:"Index,omitempty"`                           // "0"
	Level              int     `xml:"level,attr,omitempty" json:"Level,omitempty"`                           // "31"
	PixelFormat        string  `xml:"pixelFormat,attr,omitempty" json:"PixelFormat,omitempty"`               // "yuv420p"
	Profile            string  `xml:"profile,attr,omitempty" json:"Profile,omitempty"`                       // "high"
	RefFrames          int     `xml:"refFrames,attr,omitempty" json:"RefFrames,omitempty"`                   // "5"
	RequiredBandwidths string  `xml:"requiredBandwidths,attr,omitempty" json:"RequiredBandwidths,omitempty"` // "1830,1830,1830,1830,1830,1830,1830,1830"
	ScanType           string  `xml:"scanType,attr,omitempty" json:"ScanType,omitempty"`                     // "progressive"
	StreamType         int     `xml:"streamType,attr,omitempty" json:"StreamType,omitempty"`                 // "1"
	Width              int     `xml:"width,attr,omitempty" json:"Width,omitempty"`                           // "640"
}

type Part struct {
	AudioProfile        string   `xml:"audioProfile,attr,omitempty" json:"AudioProfile,omitempty"`               // "lc"
	Container           string   `xml:"container,attr,omitempty" json:"Container,omitempty"`                     // "mkv"
	DeepAnalysisVersion int      `xml:"deepAnalysisVersion,attr,omitempty" json:"DeepAnalysisVersion,omitempty"` // "1"
	Duration            int      `xml:"duration,attr,omitempty" json:"Duration,omitempty"`                       // "1451337"
	File                string   `xml:"file,attr,omitempty" json:"File,omitempty"`                               // "/cortex/central/videos/lib/TV/The Mind of a Chef/S01E12 - Fresh.mkv"
	Id                  int      `xml:"id,attr,omitempty" json:"Id,omitempty"`                                   // "154993"
	Indexes             string   `xml:"indexes,attr,omitempty" json:"Indexes,omitempty"`                         // "sd"
	Key                 string   `xml:"key,attr,omitempty" json:"Key,omitempty"`                                 // "/library/parts/154993/1393015691/file.mkv"
	RequiredBandwidths  string   `xml:"requiredBandwidths,attr,omitempty" json:"RequiredBandwidths,omitempty"`   // "1947,1947,1947,1947,1947,1947,1947,1947"
	Size                int      `xml:"size,attr,omitempty" json:"Size,omitempty"`                               // "249516275"
	Streams             []Stream `xml:"Stream" json:"Stream,omitempty"`
	VideoProfile        string   `xml:"videoProfile,omitempty" json:"VideoProfile,omitempty"` // "high"
}

type Media struct {
	AspectRatio     float32 `xml:"aspectRatio,attr,omitempty" json:"AspectRatio,omitempty"`     // "1.78"
	AudioChannels   int     `xml:"audioChannels,attr,omitempty" json:"AudioChannels,omitempty"` // "6"
	AudioCodec      string  `xml:"audioCodec,attr,omitempty" json:"AudioCodec,omitempty"`       // "ac3"
	Bitrate         int     `xml:"bitrate,attr,omitempty" json:"Bitrate,omitempty"`             // "4322"
	Container       string  `xml:"container,attr,omitempty" json:"Container,omitempty"`         // "mkv"
	Duration        int     `xml:"duration,attr,omitempty" json:"Duration,omitempty"`           // "1260427"
	Height          int     `xml:"height,attr,omitempty" json:"Height,omitempty"`               // "720"
	Id              int     `xml:"id,attr,omitempty" json:"Id,omitempty"`                       // "77758"
	Parts           []Part  `xml:"Part" json:"Part,omitempty"`
	VideoCodec      string  `xml:"videoCodec,omitempty" json:"VideoCodec,omitempty"`                // "h264"
	VideoFrameRate  string  `xml:"videoFrameRate,attr,omitempty" json:"VideoFrameRate,omitempty"`   // "24p"
	VideoProfile    string  `xml:"videoProfile,attr,omitempty" json:"VideoProfile,omitempty"`       // "high"
	VideoResolution int     `xml:"videoResolution,attr,omitempty" json:"VideoResolution,omitempty"` // "720"
	Width           int     `xml:"width,attr,omitempty" json:"Width,omitempty"`                     // "1280"
}

type Chapter struct {
	Id              int    `xml:"id,attr,omitempty" json:"Id,omitempty"`
	Tag             string `xml:"tag,attr,omitempty" json:"Tag,omitempty"`
	Index           int    `xml:"index,attr,omitempty" json:"Index,omitempty"`
	StartTimeOffset int    `xml:"startTimeOffset,attr,omitempty" json:"StartTimeOffset,omitempty"`
	EndTimeOffset   int    `xml:"endTimeOffset,attr,omitempty" json:"EndTimeOffset,omitempty"`
}

type Role struct {
	Id    int    `xml:"id,attr,omitempty" json:"Id,omitempty"`
	Tag   string `xml:"tag,attr,omitempty" json:"Tag,omitempty"`
	Role  string `xml:"role,attr,omitempty" json:"Role,omitempty"`
	Count int    `xml:"count,attr,omitempty" json:"Count,omitempty"`
}

type Identifier struct {
	Id    int    `xml:"id,attr,omitempty" json:"Id,omitempty"`
	Tag   string `xml:"tag,attr,omitempty" json:"Tag,omitempty"`
	Count int    `xml:"count,attr,omitempty" json:"Count,omitempty"`
}

type Video struct {
	AddedAt               int              `xml:"addedAt,attr,omitempty" json:"AddedAt,omitempty"`                         // "1456373737"
	Art                   string           `xml:"art,attr,omitempty" json:"Art,omitempty"`                                 // "/library/metadata/21666/art/1461210754"
	AudienceRating        float32          `xml:"audienceRating,attr,omitempty" json:"AudienceRating,omitempty"`           // "7.5"
	AudienceRatingImage   string           `xml:"audienceRatingImage,attr,omitempty" json:"AudienceRatingImage,omitempty"` // "rottentomatoes://image.rating.upright"
	Chapters              []Chapter        `xml:"Chapter" json:"Chapters,omitempty"`
	ChapterSource         string           `xml:"chapterSource,omitempty" json:"ChapterSource,omitempty"`      // ""
	ContentRating         string           `xml:"contentRating,attr,omitempty" json:"ContentRating,omitempty"` // "TV-MA"
	Country               Identifier       `xml:"Country" json:"Country,omitempty"`
	Directors             []Identifier     `xml:"Director" json:"Directors,omitempty"`
	Duration              int              `xml:"duration,omitempty" json:"Duration,omitempty"` // "1260427"
	Extras                []MediaContainer `xml:"Extras" json:"Extras,omitempty"`
	Genres                []Identifier     `xml:"Genre" json:"Genres,omitempty"`
	GrandparentArt        string           `xml:"grandparentArt,omitempty" json:"GrandparentArt,omitempty"`                  // "/library/metadata/21666/art/1461210754"
	GrandparentKey        string           `xml:"grandparentKey,attr,omitempty" json:"GrandparentKey,omitempty"`             // "/library/metadata/21666"
	GrandparentRatingKey  int              `xml:"grandparentRatingKey,attr,omitempty" json:"GrandparentRatingKey,omitempty"` // "21666"
	GrandparentTheme      string           `xml:"grandparentTheme,attr,omitempty" json:"GrandparentTheme,omitempty"`         // "/library/metadata/21666/theme/1461210754"
	GrandparentThumb      string           `xml:"grandparentThumb,attr,omitempty" json:"GrandparentThumb,omitempty"`         // "/library/metadata/21666/thumb/1461210754"
	GrandparentTitle      string           `xml:"grandparentTitle,attr,omitempty" json:"GrandparentTitle,omitempty"`         // "Broad City"
	Index                 int              `xml:"index,attr,omitempty" json:"Index,omitempty"`                               // "2"
	Key                   string           `xml:"key,attr,omitempty" json:"Key,omitempty"`                                   // "/library/metadata/60627"
	LastViewedAt          int              `xml:"lastViewedAt,attr,omitempty" json:"LastViewedAt,omitempty"`                 // "1471313734"
	LibrarySectionID      int              `xml:"librarySectionID,attr,omitempty" json:"LibrarySectionID,omitempty"`         // "4"
	Media                 Media            `xml:"Media" json:"Media,omitempty"`
	OriginallyAvailableAt string           `xml:"originallyAvailableAt,omitempty" json:"OriginallyAvailableAt,omitempty"` // "2016-02-24"
	ParentIndex           int              `xml:"parentIndex,attr,omitempty" json:"ParentIndex,omitempty"`                // "3"
	ParentKey             string           `xml:"parentKey,attr,omitempty" json:"ParentKey,omitempty"`                    // "/library/metadata/60207"
	ParentRatingKey       int              `xml:"parentRatingKey,attr,omitempty" json:"ParentRatingKey,omitempty"`        // "60207"
	Player                Player           `xml:"Player" json:"Player,omitempty"`
	PrimaryExtraKey       string           `xml:"primaryExtraKey,attr,omitempty" json:"PrimaryExtraKey,omitempty"` // "/library/metadata/99398"
	Producers             []Identifier     `xml:"Producer" json:"Producers,omitempty"`
	Rating                float32          `xml:"rating,omitempty" json:"Rating,omitempty"`                // "8.0"
	RatingImage           string           `xml:"ratingImage,attr,omitempty" json:"RatingImage,omitempty"` // "rottentomatoes://image.rating.certified"
	RatingKey             int              `xml:"ratingKey,attr,omitempty" json:"RatingKey,omitempty"`     // "60627"
	Related               []MediaContainer `xml:"Related" json:"Related,omitempty"`
	Roles                 []Role           `xml:"Role" json:"Role,omitempty"`
	Studio                string           `xml:"studio,omitempty" json:"Studio,omitempty"`        // "United Artists"
	Summary               string           `xml:"summary,attr,omitempty" json:"Summary,omitempty"` // "Abbi pretends to be Ilana to cover her shift at the food co-op, while Ilana goes to an important doctor's appointment on Long Island."
	Tagline               string           `xml:"tagline,attr,omitempty" json:"Tagline,omitempty"` // "Is it a game, or is it real?"
	Thumb                 string           `xml:"thumb,attr,omitempty" json:"Thumb,omitempty"`     // "/library/metadata/60627/thumb/1470739012"
	Title                 string           `xml:"title,attr,omitempty" json:"Title,omitempty"`     // "Co-Op"
	TranscodeSession      TranscodeSession `xml:"TranscodeSession" json:"TranscodeSession,omitempty"`
	Type                  string           `xml:"type,attr,omitempty" json:"Type,omitempty"`           // "episode"
	UpdatedAt             int              `xml:"updatedAt,attr,omitempty" json:"UpdatedAt,omitempty"` // "1470739012"
	User                  User             `xml:"User,omitempty" json:"User,omitempty"`
	ViewCount             int              `xml:"viewCount,attr,omitempty" json:"ViewCount,omitempty"`   // "1"
	ViewedAt              int64            `xml:"viewedAt,attr,omitempty" json:"ViewedAt,omitempty"`     // "1472089998"
	ViewOffset            int              `xml:"viewOffset,attr,omitempty" json:"ViewOffset,omitempty"` // "102000"
	Writers               []Identifier     `xml:"Writer" json:"Writers,omitempty"`
	Year                  int              `xml:"year,omitempty" json:"Year,omitempty"` // "2016"
}
