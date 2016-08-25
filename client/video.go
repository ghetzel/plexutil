package client

type Stream struct {
	BitDepth           int     `xml:"bitDepth,attr,omitempty" json:"bitDepth,omitempty"`                     // "8"
	Bitrate            int     `xml:"bitrate,attr,omitempty" json:"bitrate,omitempty"`                       // "1375"
	Cabac              int     `xml:"cabac,attr,omitempty" json:"cabac,omitempty"`                           // "1"
	ChromaSubsampling  string  `xml:"chromaSubsampling,attr,omitempty" json:"chromaSubsampling,omitempty"`   // "4:2:0"
	Codec              string  `xml:"codec,attr,omitempty" json:"codec,omitempty"`                           // "h264"
	CodecID            string  `xml:"codecID,attr,omitempty" json:"codecID,omitempty"`                       // "V_MPEG4/ISO/AVC"
	Default            int     `xml:"default,attr,omitempty" json:"default,omitempty"`                       // "1"
	Duration           int     `xml:"duration,attr,omitempty" json:"duration,omitempty"`                     // "1451326"
	FrameRate          float32 `xml:"frameRate,attr,omitempty" json:"frameRate,omitempty"`                   // "23.976"
	FrameRateMode      string  `xml:"frameRateMode,attr,omitempty" json:"frameRateMode,omitempty"`           // "cfr"
	HasScalingMatrix   int     `xml:"hasScalingMatrix,attr,omitempty" json:"hasScalingMatrix,omitempty"`     // "0"
	Height             int     `xml:"height,attr,omitempty" json:"height,omitempty"`                         // "358"
	Id                 int     `xml:"id,attr,omitempty" json:"id,omitempty"`                                 // "291338"
	Index              int     `xml:"index,attr,omitempty" json:"index,omitempty"`                           // "0"
	Level              int     `xml:"level,attr,omitempty" json:"level,omitempty"`                           // "31"
	PixelFormat        string  `xml:"pixelFormat,attr,omitempty" json:"pixelFormat,omitempty"`               // "yuv420p"
	Profile            string  `xml:"profile,attr,omitempty" json:"profile,omitempty"`                       // "high"
	RefFrames          int     `xml:"refFrames,attr,omitempty" json:"refFrames,omitempty"`                   // "5"
	RequiredBandwidths string  `xml:"requiredBandwidths,attr,omitempty" json:"requiredBandwidths,omitempty"` // "1830,1830,1830,1830,1830,1830,1830,1830"
	ScanType           string  `xml:"scanType,attr,omitempty" json:"scanType,omitempty"`                     // "progressive"
	StreamType         int     `xml:"streamType,attr,omitempty" json:"streamType,omitempty"`                 // "1"
	Width              int     `xml:"width,attr,omitempty" json:"width,omitempty"`                           // "640"
}

type Part struct {
	AudioProfile        string   `xml:"audioProfile,attr,omitempty" json:"audioProfile,omitempty"`               // "lc"
	Container           string   `xml:"container,attr,omitempty" json:"container,omitempty"`                     // "mkv"
	DeepAnalysisVersion int      `xml:"deepAnalysisVersion,attr,omitempty" json:"deepAnalysisVersion,omitempty"` // "1"
	Duration            int      `xml:"duration,attr,omitempty" json:"duration,omitempty"`                       // "1451337"
	File                string   `xml:"file,attr,omitempty" json:"file,omitempty"`                               // "/cortex/central/videos/lib/TV/The Mind of a Chef/S01E12 - Fresh.mkv"
	Id                  int      `xml:"id,attr,omitempty" json:"id,omitempty"`                                   // "154993"
	Indexes             string   `xml:"indexes,attr,omitempty" json:"indexes,omitempty"`                         // "sd"
	Key                 string   `xml:"key,attr,omitempty" json:"key,omitempty"`                                 // "/library/parts/154993/1393015691/file.mkv"
	RequiredBandwidths  string   `xml:"requiredBandwidths,attr,omitempty" json:"requiredBandwidths,omitempty"`   // "1947,1947,1947,1947,1947,1947,1947,1947"
	Size                int      `xml:"size,attr,omitempty" json:"size,omitempty"`                               // "249516275"
	Streams             []Stream `xml:"Stream" json:"Stream"`
	VideoProfile        string   `xml:"videoProfile,omitempty" json:"videoProfile,omitempty"` // "high"
}

type Media struct {
	AspectRatio     float32 `xml:"aspectRatio,attr,omitempty" json:"aspectRatio,omitempty"`     // "1.78"
	AudioChannels   int     `xml:"audioChannels,attr,omitempty" json:"audioChannels,omitempty"` // "6"
	AudioCodec      string  `xml:"audioCodec,attr,omitempty" json:"audioCodec,omitempty"`       // "ac3"
	Bitrate         int     `xml:"bitrate,attr,omitempty" json:"bitrate,omitempty"`             // "4322"
	Container       string  `xml:"container,attr,omitempty" json:"container,omitempty"`         // "mkv"
	Duration        int     `xml:"duration,attr,omitempty" json:"duration,omitempty"`           // "1260427"
	Height          int     `xml:"height,attr,omitempty" json:"height,omitempty"`               // "720"
	Id              int     `xml:"id,attr,omitempty" json:"id,omitempty"`                       // "77758"
	Parts           []Part  `xml:"Part" json:"Part"`
	VideoCodec      string  `xml:"videoCodec,omitempty" json:"videoCodec,omitempty"`                // "h264"
	VideoFrameRate  string  `xml:"videoFrameRate,attr,omitempty" json:"videoFrameRate,omitempty"`   // "24p"
	VideoProfile    string  `xml:"videoProfile,attr,omitempty" json:"videoProfile,omitempty"`       // "high"
	VideoResolution int     `xml:"videoResolution,attr,omitempty" json:"videoResolution,omitempty"` // "720"
	Width           int     `xml:"width,attr,omitempty" json:"width,omitempty"`                     // "1280"
}

type Chapter struct {
	Id              int    `xml:"id,attr,omitempty" json:"id,omitempty"`
	Tag             string `xml:"tag,attr,omitempty" json:"tag,omitempty"`
	Index           int    `xml:"index,attr,omitempty" json:"index,omitempty"`
	StartTimeOffset int    `xml:"startTimeOffset,attr,omitempty" json:"startTimeOffset,omitempty"`
	EndTimeOffset   int    `xml:"endTimeOffset,attr,omitempty" json:"endTimeOffset,omitempty"`
}

type Role struct {
	Id    int    `xml:"id,attr,omitempty" json:"id,omitempty"`
	Tag   string `xml:"tag,attr,omitempty" json:"tag,omitempty"`
	Role  string `xml:"role,attr,omitempty" json:"role,omitempty"`
	Count int    `xml:"count,attr,omitempty" json:"count,omitempty"`
}

type Identifier struct {
	Id    int    `xml:"id,attr,omitempty" json:"id,omitempty"`
	Tag   string `xml:"tag,attr,omitempty" json:"tag,omitempty"`
	Count int    `xml:"count,attr,omitempty" json:"count,omitempty"`
}

type Video struct {
	AddedAt               int              `xml:"addedAt,attr,omitempty" json:"addedAt,omitempty"`                         // "1456373737"
	Art                   string           `xml:"art,attr,omitempty" json:"art,omitempty"`                                 // "/library/metadata/21666/art/1461210754"
	AudienceRating        float32          `xml:"audienceRating,attr,omitempty" json:"audienceRating,omitempty"`           // "7.5"
	AudienceRatingImage   string           `xml:"audienceRatingImage,attr,omitempty" json:"audienceRatingImage,omitempty"` // "rottentomatoes://image.rating.upright"
	Chapters              []Chapter        `xml:"Chapter" json:"Chapter"`
	ChapterSource         string           `xml:"chapterSource,omitempty" json:"chapterSource,omitempty"`      // ""
	ContentRating         string           `xml:"contentRating,attr,omitempty" json:"contentRating,omitempty"` // "TV-MA"
	Country               Identifier       `xml:"Country" json:"Country"`
	Directors             []Identifier     `xml:"Director" json:"Director"`
	Duration              int              `xml:"duration,omitempty" json:"duration,omitempty"` // "1260427"
	Extras                []MediaContainer `xml:"Extras" json:"Extras"`
	Genres                []Identifier     `xml:"Genre" json:"Genre"`
	GrandparentArt        string           `xml:"grandparentArt,omitempty" json:"grandparentArt,omitempty"`                  // "/library/metadata/21666/art/1461210754"
	GrandparentKey        string           `xml:"grandparentKey,attr,omitempty" json:"grandparentKey,omitempty"`             // "/library/metadata/21666"
	GrandparentRatingKey  int              `xml:"grandparentRatingKey,attr,omitempty" json:"grandparentRatingKey,omitempty"` // "21666"
	GrandparentTheme      string           `xml:"grandparentTheme,attr,omitempty" json:"grandparentTheme,omitempty"`         // "/library/metadata/21666/theme/1461210754"
	GrandparentThumb      string           `xml:"grandparentThumb,attr,omitempty" json:"grandparentThumb,omitempty"`         // "/library/metadata/21666/thumb/1461210754"
	GrandparentTitle      string           `xml:"grandparentTitle,attr,omitempty" json:"grandparentTitle,omitempty"`         // "Broad City"
	Index                 int              `xml:"index,attr,omitempty" json:"index,omitempty"`                               // "2"
	Key                   string           `xml:"key,attr,omitempty" json:"key,omitempty"`                                   // "/library/metadata/60627"
	LastViewedAt          int              `xml:"lastViewedAt,attr,omitempty" json:"lastViewedAt,omitempty"`                 // "1471313734"
	LibrarySectionID      int              `xml:"librarySectionID,attr,omitempty" json:"librarySectionID,omitempty"`         // "4"
	Media                 Media            `xml:"Media" json:"Media"`
	OriginallyAvailableAt string           `xml:"originallyAvailableAt,omitempty" json:"originallyAvailableAt,omitempty"` // "2016-02-24"
	ParentIndex           int              `xml:"parentIndex,attr,omitempty" json:"parentIndex,omitempty"`                // "3"
	ParentKey             string           `xml:"parentKey,attr,omitempty" json:"parentKey,omitempty"`                    // "/library/metadata/60207"
	ParentRatingKey       int              `xml:"parentRatingKey,attr,omitempty" json:"parentRatingKey,omitempty"`        // "60207"
	Player                Player           `xml:"Player" json:"Player"`
	PrimaryExtraKey       string           `xml:"primaryExtraKey,attr,omitempty" json:"primaryExtraKey,omitempty"` // "/library/metadata/99398"
	Producers             []Identifier     `xml:"Producer" json:"Producer"`
	Rating                float32          `xml:"rating,omitempty" json:"rating,omitempty"`                // "8.0"
	RatingImage           string           `xml:"ratingImage,attr,omitempty" json:"ratingImage,omitempty"` // "rottentomatoes://image.rating.certified"
	RatingKey             int              `xml:"ratingKey,attr,omitempty" json:"ratingKey,omitempty"`     // "60627"
	Related               []MediaContainer `xml:"Related" json:"Related"`
	Roles                 []Role           `xml:"Role" json:"Role"`
	Studio                string           `xml:"studio,omitempty" json:"studio,omitempty"`        // "United Artists"
	Summary               string           `xml:"summary,attr,omitempty" json:"summary,omitempty"` // "Abbi pretends to be Ilana to cover her shift at the food co-op, while Ilana goes to an important doctor's appointment on Long Island."
	Tagline               string           `xml:"tagline,attr,omitempty" json:"tagline,omitempty"` // "Is it a game, or is it real?"
	Thumb                 string           `xml:"thumb,attr,omitempty" json:"thumb,omitempty"`     // "/library/metadata/60627/thumb/1470739012"
	Title                 string           `xml:"title,attr,omitempty" json:"title,omitempty"`     // "Co-Op"
	TranscodeSession      TranscodeSession `xml:"TranscodeSession" json:"TranscodeSession,omitempty"`
	Type                  string           `xml:"type,attr,omitempty" json:"type,omitempty"`           // "episode"
	UpdatedAt             int              `xml:"updatedAt,attr,omitempty" json:"updatedAt,omitempty"` // "1470739012"
	User                  User             `xml:"User,omitempty" json:"user,omitempty"`
	ViewCount             int              `xml:"viewCount,attr,omitempty" json:"viewCount,omitempty"`   // "1"
	ViewOffset            int              `xml:"viewOffset,attr,omitempty" json:"viewOffset,omitempty"` // "102000"
	Writers               []Identifier     `xml:"Writer" json:"Writer"`
	Year                  int              `xml:"year,omitempty" json:"year,omitempty"` // "2016"
}
