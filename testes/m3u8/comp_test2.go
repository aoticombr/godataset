package main

import (
	"encoding/json"
	"fmt"
	"testing"

	httpaoti "github.com/aoticombr/golang/http"
)

type Lesson struct {
	ID    string    `json:"id"`
	Type  string    `json:"type"`
	Last  VideoInfo `json:"last"`
	Pivot Pivot     `json:"pivot"`
}

type VideoInfo struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	Slug             string      `json:"slug"`
	Description      string      `json:"description"`
	ChallengeID      string      `json:"challenge_id"`
	Platform         string      `json:"platform"`
	Resource         string      `json:"resource"`
	Duration         int         `json:"duration"`
	Link             string      `json:"link"`
	IsFree           bool        `json:"is_free"`
	JupiterVideoID   string      `json:"jupiter_video_id"`
	SelfEvaluationID string      `json:"self_evaluation_step_id"`
	LessonJourneyID  string      `json:"lesson_journey_id"`
	Downloads        []string    `json:"downloads"`
	ContentTags      []string    `json:"contentTags"`
	Author           Author      `json:"author"`
	Challenge        interface{} `json:"challenge"`
	Sections         []string    `json:"sections"`
	Progress         Progress    `json:"progress"`
}

type Pivot struct {
	LessonID      string `json:"lesson_id"`
	LessonGroupID string `json:"lesson_group_id"`
	Order         int    `json:"order"`
}

type Progress struct {
	ID              string `json:"_id"`
	LessonHistoryID string `json:"lesson_history_id"`
	UserID          string `json:"user_id"`
	UpdatedAt       string `json:"updated_at"`
	Completed       bool   `json:"completed"`
	Percentage      int    `json:"percentage"`
	ProgressTime    int    `json:"progress_time"`
	Meta            Meta   `json:"meta"`
}

type Meta struct {
	JourneyID           string `json:"journey_id"`
	JourneyThumbnailURL string `json:"journey_thumbnail_url"`
	JourneyTitle        string `json:"journey_title"`
	JourneyType         string `json:"journey_type"`
	Type                string `json:"type"`
	IsExpertContent     bool   `json:"is_expert_content"`
	ClusterID           string `json:"cluster_id"`
	ClusterTitle        string `json:"cluster_title"`
	ClusterSlug         string `json:"cluster_slug"`
	ClusterThumbnailURL string `json:"cluster_thumbnail_url"`
	LessonGroupID       string `json:"lesson_group_id"`
	GroupSlug           string `json:"group_slug"`
	GroupTitle          string `json:"group_title"`
	LessonID            string `json:"lesson_id"`
	LessonSlug          string `json:"lesson_slug"`
	LessonTitle         string `json:"lesson_title"`
}

type Author struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	AvatarURL string `json:"avatar_url"`
}

type Cluster struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Slug   string  `json:"slug"`
	Groups []Group `json:"groups"`
}

type Group struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Slug    string   `json:"slug"`
	Lessons []Lesson `json:"lessons"`
	Pivot   Pivot    `json:"pivot"`
}

type Journey struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Slug           string `json:"slug"`
	ContentType    string `json:"content_type"`
	IsCertificable bool   `json:"is_certificable"`
	IsFree         bool   `json:"is_free"`
	Forum          Forum  `json:"forum"`
}

type Forum struct {
	ID                     string `json:"id"`
	Title                  string `json:"title"`
	Slug                   string `json:"slug"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	GamificationMultiplier int    `json:"gamification_multiplier"`
}

type Course struct {
	ID                   string      `json:"id"`
	JourneyID            string      `json:"journey_id"`
	ClusterID            string      `json:"cluster_id"`
	LessonGroupID        interface{} `json:"lesson_group_id"`
	LessonID             interface{} `json:"lesson_id"`
	IsMilestone          bool        `json:"is_milestone"`
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	Color                string      `json:"color"`
	Thumbnail            string      `json:"thumbnail"`
	Order                int         `json:"order"`
	Conditional          bool        `json:"conditional"`
	Active               bool        `json:"active"`
	CreatedAt            string      `json:"created_at"`
	UpdatedAt            string      `json:"updated_at"`
	ReleaseAt            interface{} `json:"release_at"`
	Slug                 string      `json:"slug"`
	Layout               string      `json:"layout"`
	Author               string      `json:"author"`
	ChallengeID          interface{} `json:"challenge_id"`
	CategoryID           interface{} `json:"category_id"`
	Introductory         bool        `json:"introductory"`
	RewardID             interface{} `json:"reward_id"`
	CheckpointID         interface{} `json:"checkpoint_id"`
	Premiere             bool        `json:"premiere"`
	MilestoneDescription string      `json:"milestone_description"`
	Module               bool        `json:"module"`
	ContentType          interface{} `json:"content_type"`
	Metadata             interface{} `json:"metadata"`
	PartneroID           interface{} `json:"partnero_id"`
	IsBonus              bool        `json:"is_bonus"`
	HasMicroCertificate  bool        `json:"has_micro_certificate"`
	Subtitle             string      `json:"subtitle"`
	IsSingleClassView    bool        `json:"is_single_class_view"`
	Icon                 interface{} `json:"icon"`
	HasAfterAccess       bool        `json:"has_after_access"`
	ThumbnailURL         string      `json:"thumbnail_url"`
	Type                 string      `json:"type"`
	IsNew                bool        `json:"is_new"`
	IconURL              interface{} `json:"icon_url"`
	Journey              Journey     `json:"journey"`
	Cluster              Cluster     `json:"cluster"`
	Group                interface{} `json:"group"`
	Lesson               interface{} `json:"lesson"`
	IsFreemium           bool        `json:"is_freemium"`
	Available            bool        `json:"available"`
	CurrentLesson        string      `json:"current_lesson"`
}

func TestUp_Down2(t *testing.T) {
	fmt.Printf("teste")
	link2 := httpaoti.NewHttp()
	link2.SetUrl("https://skylab-api.rocketseat.com.br/journey-nodes/projeto-01")
	link2.SetMetodo(httpaoti.M_GET)
	link2.Request.Header.Authorization = "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhMmNmOTllOC01OTg0LTQ1ZjEtYTU3Yy01YWZmYzc4ZjE5YjYiLCJkYXRhIjp7fSwiaWF0IjoxNzA1MDE0NjA4LCJleHAiOjE3MDUxODc0MDh9.cz1hktCYBGiHxfGePAROPZSTv1IlYHhkMkPCPyDIMEo8jOb1anHEsDrxLj8JzcR5pWspcD_ZVUrTpBdOK2pIYPHhROs_cY4C1LvNEuBJFTqNrRzbLtyF-bRIYEBm77MutsIciIqg5J7epOhNM6NeF1wxyuTcWPgq1dXGOpLnjRruPis1nwLbHkFfqrtdlrkhguJPDynPhK7e-Q6DPz7l99BSvnhkCSKmQh6gpu4J9Fiel7SGCLuEmY7ffyo-rXH8qC5jHLf77_tuF8Zs3kbCQRpuNzOFdA38aXIR4UEiw0a2jYHGZWWycteXdVwrXm9-LyawB1FolUogEro6hGJda9d6saIczKH4oyDxT71iNt1whVkRKay_TfOQ1njA8bm9pG9KzdQyKxA0dwgo7kMQZcypzplwN5kTcHKt6_GDaFnyJuN41KGkopDs-sOnJOW50I-VvlgbhvVVZPJpaPn2BEqlwhEWayc89lOBTptRbDZwRNCRSqfo-NZektNASW_CL4XenEV-yHxsA3CxSY1MJAQaJvUG5KBBBnEzKqupwu_kNVsvi5vSJpZjXEzIM_0BjomiZMi2a3x9oVxL-P1vv7exibGou30H2DHinG8R0P6tdolR_eAoO1CbB00A5H0clfQL5hRAD_BFtSlTlkzG7khAwoOeIUykC37tIfXp-j4"
	link2.Request.Header.AddField("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36")
	link2.Request.Header.AddField("origin", "https://app.rocketseat.com.br")
	link2.Request.Header.AddField("referer", "https://app.rocketseat.com.br/")
	resp2, err := link2.Send()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(resp2.Body))

	var course Course
	err = json.Unmarshal(resp2.Body, &course)
	if err != nil {
		fmt.Println("Erro ao fazer o unmarshal dos dados JSON:", err)
		return
	}

	// fmt.Println("Teste")
	// down := m3u8.NewM3u8()
	// byt, err := down.GetVideoByte("https://b-vz-762f4670-e04.tv.pandavideo.com.br/dd0f4f59-f80d-4dfe-864f-e53652f4fea7/playlist.m3u8")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// newUUID := uuid.New()
	// timeString := time.Now().Format("2006-01-02-15-04-05")
	// Name := timeString + newUUID.String() + "playlist.mkv"

	// m3u8.SaveByteToFile(Name, byt)
}
