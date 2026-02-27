package dto

type NETEASEMusicDetailsStruct struct {
	Privileges []Privilege `json:"privileges"`
	Code       int         `json:"code"`
	Songs      []Song      `json:"songs"`
}

type Privilege struct {
	Code               int                `json:"code"`
	Flag               int                `json:"flag"`
	DLLevel            string             `json:"dlLevel"`
	FL                 int                `json:"fl"`
	Subp               int                `json:"subp"`
	Fee                int                `json:"fee"`
	DL                 int                `json:"dl"`
	PlLevel            string             `json:"plLevel"`
	MaxBrLevel         string             `json:"maxBrLevel"`
	RightSource        int                `json:"rightSource"`
	Maxbr              int                `json:"maxbr"`
	ID                 int                `json:"id"`
	SP                 int                `json:"sp"`
	Payed              int                `json:"payed"`
	St                 int                `json:"st"`
	ChargeInfoList     []ChargeInfoList   `json:"chargeInfoList"`
	FreeTrialPrivilege FreeTrialPrivilege `json:"freeTrialPrivilege"`
	DownloadMaxBrLevel string             `json:"downloadMaxBrLevel"`
	DownloadMaxbr      int                `json:"downloadMaxbr"`
	Cp                 int                `json:"cp"`
	PreSell            bool               `json:"preSell"`
	PlayMaxBrLevel     string             `json:"playMaxBrLevel"`
	CS                 bool               `json:"cs"`
	PlayMaxbr          int                `json:"playMaxbr"`
	Toast              bool               `json:"toast"`
	FLLevel            string             `json:"flLevel"`
	Pl                 int                `json:"pl"`
}

type ChargeInfoList struct {
	Rate       int `json:"rate"`
	ChargeType int `json:"chargeType"`
}

type FreeTrialPrivilege struct {
	UserConsumable bool `json:"userConsumable"`
	ResConsumable  bool `json:"resConsumable"`
}

type Song struct {
	Copyright       int           `json:"copyright"`
	Fee             int           `json:"fee"`
	Mst             int           `json:"mst"`
	Dt              int           `json:"dt"`
	ResourceState   bool          `json:"resourceState"`
	ID              int           `json:"id"`
	PublishTime     int           `json:"publishTime"`
	Mv              int           `json:"mv"`
	Al              Al            `json:"al"`
	Version         int           `json:"version"`
	Alia            []interface{} `json:"alia"`
	Ar              []Ar          `json:"ar"`
	MarkTags        []interface{} `json:"markTags"`
	Name            string        `json:"name"`
	No              int           `json:"no"`
	Pop             float64       `json:"pop"`
	Pst             int           `json:"pst"`
	Rtype           int           `json:"rtype"`
	SID             int           `json:"s_id"`
	RtUrls          []interface{} `json:"rtUrls"`
	Sq              H             `json:"sq"`
	CD              string        `json:"cd"`
	St              int           `json:"st"`
	CF              string        `json:"cf"`
	OriginCoverType int           `json:"originCoverType"`
	H               H             `json:"h"`
	L               H             `json:"l"`
	Cp              int           `json:"cp"`
	M               H             `json:"m"`
	DjID            int           `json:"djId"`
	Single          int           `json:"single"`
	Ftype           int           `json:"ftype"`
	T               int           `json:"t"`
	V               int           `json:"v"`
	Mark            int           `json:"mark"`
}

type Al struct {
	PicURL string        `json:"picUrl"`
	Name   string        `json:"name"`
	Tns    []interface{} `json:"tns"`
	PicStr string        `json:"pic_str"`
	ID     int           `json:"id"`
	Pic    float64       `json:"pic"`
}

type Ar struct {
	Name  string        `json:"name"`
	Tns   []interface{} `json:"tns"`
	Alias []interface{} `json:"alias"`
	ID    int           `json:"id"`
}

type H struct {
	Br   int     `json:"br"`
	Fid  int     `json:"fid"`
	Size int     `json:"size"`
	Vd   float64 `json:"vd"`
	Sr   int     `json:"sr"`
}

type NETEASEMusicLyricStruct struct {
	Code int  `json:"code"`
	Qfy  bool `json:"qfy"`
	Sfy  bool `json:"sfy"`
	Lrc  Lrc  `json:"lrc"`
	Sgc  bool `json:"sgc"`
}

type Lrc struct {
	Lyric   string `json:"lyric"`
	Version int    `json:"version"`
}

type NETEASEMusicURLStruct struct {
	Code int     `json:"code"`
	Data []Datum `json:"data"`
}

type Datum struct {
	Code                   int                    `json:"code"`
	Expi                   int                    `json:"expi"`
	Flag                   int                    `json:"flag"`
	Fee                    int                    `json:"fee"`
	BeatType               int                    `json:"beatType"`
	URLSource              int                    `json:"urlSource"`
	CanExtend              bool                   `json:"canExtend"`
	Type                   string                 `json:"type"`
	FreeTimeTrialPrivilege FreeTimeTrialPrivilege `json:"freeTimeTrialPrivilege"`
	Gain                   float64                `json:"gain"`
	Br                     int                    `json:"br"`
	MusicID                string                 `json:"musicId"`
	EncodeType             string                 `json:"encodeType"`
	RightSource            int                    `json:"rightSource"`
	ClosedGain             float64                `json:"closedGain"`
	ID                     int                    `json:"id"`
	Payed                  int                    `json:"payed"`
	Sr                     int                    `json:"sr"`
	Level                  string                 `json:"level"`
	FreeTrialPrivilege     FreeTrialPrivilege     `json:"freeTrialPrivilege"`
	Peak                   float64                `json:"peak"`
	URL                    string                 `json:"url"`
	Size                   int                    `json:"size"`
	Time                   int                    `json:"time"`
	ClosedPeak             float64                `json:"closedPeak"`
	Md5                    string                 `json:"md5"`
}

type FreeTimeTrialPrivilege struct {
	UserConsumable bool `json:"userConsumable"`
	ResConsumable  bool `json:"resConsumable"`
	RemainTime     int  `json:"remainTime"`
	Type           int  `json:"type"`
}
