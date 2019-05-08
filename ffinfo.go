package ffinfo

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"strconv"
	"time"
)

// File represents the information about media file returned by FFprobe.
type File struct {
	Format struct {
		Filename       string `json:"filename"`
		NbStreams      int    `json:"nb_streams,omitempty"`
		NbPrograms     int    `json:"nb_programs,omitempty"`
		FormatName     string `json:"format_name,omitempty"`
		FormatLongName string `json:"format_long_name,omitempty"`
		StartTime      string `json:"start_time,omitempty"`
		Duration       string `json:"duration,omitempty"`
		Size           string `json:"size,omitempty"`
		BitRate        string `json:"bit_rate,omitempty"`
		ProbeScore     int    `json:"probe_score,omitempty"`
		Tags           struct {
			MajorBrand                      string    `json:"major_brand,omitempty"`
			MinorVersion                    string    `json:"minor_version,omitempty"`
			CompatibleBrands                string    `json:"compatible_brands,omitempty"`
			CreationTime                    time.Time `json:"creation_time,omitempty"`
			ComAppleFinalcutstudioMediaUUID string    `json:"com.apple.finalcutstudio.media.uuid,omitempty"`
			Encoder                         string    `json:"encoder,omitempty"`
			ProjectName                     string    `json:"project_name,omitempty"`
			UID                             string    `json:"uid,omitempty"`
			GenerationUID                   string    `json:"generation_uid,omitempty"`
			CompanyName                     string    `json:"company_name,omitempty"`
			ProductName                     string    `json:"product_name,omitempty"`
			ProductVersion                  string    `json:"product_version,omitempty"`
			ProductUID                      string    `json:"product_uid,omitempty"`
			ModificationDate                time.Time `json:"modification_date,omitempty"`
			ApplicationPlatform             string    `json:"application_platform,omitempty"`
			MaterialPackageUmid             string    `json:"material_package_umid,omitempty"`
			MaterialPackageName             string    `json:"material_package_name,omitempty"`
			MediaType                       string    `json:"media_type,omitempty"`
			PlaybackRequirements            string    `json:"playback_requirements,omitempty"`
			Timecode                        string    `json:"timecode,omitempty"`
			Title                           string    `json:"title,omitempty"`
			Album                           string    `json:"album,omitempty"`
			Genre                           string    `json:"genre,omitempty"`
			Comment                         string    `json:"comment,omitempty"`
			Track                           string    `json:"track,omitempty"`
			Artist                          string    `json:"artist,omitempty"`
			AlbumArtist                     string    `json:"album_artist,omitempty"`
			Date                            time.Time `json:"date,omitempty"`
			SortName                        string    `json:"sort_name,omitempty"`
			Description                     string    `json:"description,omitempty"`
			Synopsis                        string    `json:"synopsis,omitempty"`
			Copyright                       string    `json:"copyright,omitempty"`
			HdVideo                         string    `json:"hd_video,omitempty"`
			Rating                          string    `json:"rating,omitempty"`
			ITunEXTC                        string    `json:"iTunEXTC,omitempty"`
			ITunMOVI                        string    `json:"iTunMOVI,omitempty"`
		} `json:"tags,omitempty"`
	} `json:"format,omitempty"`
	Streams []struct {
		Index              int    `json:"index"`
		CodecName          string `json:"codec_name,omitempty"`
		CodecLongName      string `json:"codec_long_name,omitempty"`
		Profile            string `json:"profile,omitempty"`
		CodecType          string `json:"codec_type,omitempty"`
		CodecTimeBase      string `json:"codec_time_base,omitempty"`
		CodecTagString     string `json:"codec_tag_string,omitempty"`
		CodecTag           string `json:"codec_tag,omitempty"`
		Width              int    `json:"width,omitempty"`
		Height             int    `json:"height,omitempty"`
		CodedWidth         int    `json:"coded_width,omitempty"`
		CodedHeight        int    `json:"coded_height,omitempty"`
		HasBFrames         int    `json:"has_b_frames,omitempty"`
		SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
		DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
		PixFmt             string `json:"pix_fmt,omitempty"`
		Level              int    `json:"level,omitempty"`
		ColorRange         string `json:"color_range,omitempty"`
		ColorSpace         string `json:"color_space,omitempty"`
		ColorTransfer      string `json:"color_transfer,omitempty"`
		ColorPrimaries     string `json:"color_primaries,omitempty"`
		ChromaLocation     string `json:"chroma_location,omitempty"`
		FieldOrder         string `json:"field_order,omitempty"`
		Refs               int    `json:"refs,omitempty"`
		IsAvc              string `json:"is_avc,omitempty"`
		NalLengthSize      string `json:"nal_length_size,omitempty"`
		ID                 string `json:"id,omitempty"`
		RFrameRate         string `json:"r_frame_rate,omitempty"`
		AvgFrameRate       string `json:"avg_frame_rate,omitempty"`
		TimeBase           string `json:"time_base,omitempty"`
		StartPts           int    `json:"start_pts,omitempty"`
		StartTime          string `json:"start_time,omitempty"`
		DurationTs         int    `json:"duration_ts,omitempty"`
		Duration           string `json:"duration,omitempty"`
		BitRate            string `json:"bit_rate,omitempty"`
		MaxBitRate         string `json:"max_bit_rate,omitempty"`
		BitsPerSample      int    `json:"bits_per_sample,omitempty"`
		BitsPerRawSample   string `json:"bits_per_raw_sample,omitempty"`
		NbFrames           string `json:"nb_frames,omitempty"`
		SampleFmt          string `json:"sample_fmt,omitempty"`
		SampleRate         string `json:"sample_rate,omitempty"`
		Channels           int    `json:"channels,omitempty"`
		ChannelLayout      string `json:"channel_layout,omitempty"`
		DmixMode           string `json:"dmix_mode,omitempty"`
		LtrtCmixlev        string `json:"ltrt_cmixlev,omitempty"`
		LtrtSurmixlev      string `json:"ltrt_surmixlev,omitempty"`
		LoroCmixlev        string `json:"loro_cmixlev,omitempty"`
		LoroSurmixlev      string `json:"loro_surmixlev,omitempty"`
		Disposition        struct {
			Default         int `json:"default,omitempty"`
			Dub             int `json:"dub,omitempty"`
			Original        int `json:"original,omitempty"`
			Comment         int `json:"comment,omitempty"`
			Lyrics          int `json:"lyrics,omitempty"`
			Karaoke         int `json:"karaoke,omitempty"`
			Forced          int `json:"forced,omitempty"`
			HearingImpaired int `json:"hearing_impaired,omitempty"`
			VisualImpaired  int `json:"visual_impaired,omitempty"`
			CleanEffects    int `json:"clean_effects,omitempty"`
			AttachedPic     int `json:"attached_pic,omitempty"`
			TimedThumbnails int `json:"timed_thumbnails,omitempty"`
		} `json:"disposition,omitempty"`
		Tags struct {
			CreationTime                time.Time `json:"creation_time,omitempty"`
			Language                    string    `json:"language,omitempty"`
			Title                       string    `json:"title,omitempty"`
			HandlerName                 string    `json:"handler_name,omitempty"`
			Encoder                     string    `json:"encoder,omitempty"`
			Timecode                    string    `json:"timecode,omitempty"`
			FilePackageUmid             string    `json:"file_package_umid,omitempty"`
			FilePackageName             string    `json:"file_package_name,omitempty"`
			TrackName                   string    `json:"track_name,omitempty"`
			BPSEng                      string    `json:"BPS-eng"`
			DURATIONEng                 string    `json:"DURATION-eng"`
			NUMBEROFFRAMESEng           string    `json:"NUMBER_OF_FRAMES-eng"`
			NUMBEROFBYTESEng            string    `json:"NUMBER_OF_BYTES-eng"`
			STATISTICSWRITINGAPPEng     string    `json:"_STATISTICS_WRITING_APP-eng"`
			STATISTICSWRITINGDATEUTCEng string    `json:"_STATISTICS_WRITING_DATE_UTC-eng"`
			STATISTICSTAGSEng           string    `json:"_STATISTICS_TAGS-eng"`
		} `json:"tags,omitempty"`
		SideDataList []struct {
			SideDataType string `json:"side_data_type,omitempty"`
		} `json:"side_data_list,omitempty"`
	} `json:"streams,omitempty"`
}

// Print prints out the contents of the media file.
func (f *File) String() string {
	b, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// StreamDuration return streams duration as float64,
// uses MKV tags if the default duration value is not found.
// Returns -1 on errors.
// Returns format duration and error if no stream duration was found.
// returns 0 and error if no stream or fromat duration were found.
func (f *File) StreamDuration(i int) (d float64, err error) {
	if len(f.Streams) < 1 {
		return -1, errors.New("file has no streams")
	}

	if len(f.Streams)-1 < i || i < 0 {
		return -1, errors.New("stream index " + strconv.Itoa(i) + " is out of [0-" + strconv.Itoa(len(f.Streams)-1) + "] range")
	}

	if f.Streams[i].Duration == "" {
		if f.Streams[i].Tags.DURATIONEng != "" {
			return HMSMSToSeconds(f.Streams[i].Tags.DURATIONEng), nil
		}
		if f.Format.Duration != "" {
			d, err = strconv.ParseFloat(f.Format.Duration, 64)
			if err != nil {
				return -1, err
			}
		}
		return d, errors.New("stream " + strconv.Itoa(i) + " has no duration metadata")
	}

	d, err = strconv.ParseFloat(f.Streams[i].Duration, 64)
	if err != nil {
		return -1, err
	}

	return d, nil
}

// Probe returns media file information.
func Probe(filePath string) (f *File, err error) {
	c := exec.Command("ffprobe",
		"-loglevel", "error",
		"-hide_banner",
		"-show_format",
		"-show_streams",
		"-print_format", "json=c=1",
		filePath)
	var o bytes.Buffer
	var e bytes.Buffer
	c.Stdout = &o
	c.Stderr = &e
	err = c.Run()
	if err != nil {
		return f, errors.New(string(e.Bytes()))
	}
	json.Unmarshal(o.Bytes(), &f)
	return f, nil
}

// HMSMSToSeconds - converts timecode (H:M:S.MS) to seconds float64 (S.MS).
func HMSMSToSeconds(hhmmssms string) float64 {
	var hh, mm, ss, ms float64
	var buffer string
	length := len(hhmmssms)
	timecode := []string{}

	for i := length - 1; i >= 0; i-- {
		if hhmmssms[i] == '.' {
			buffer = "." + buffer
			ms, _ = strconv.ParseFloat(buffer, 64)
			buffer = ""
		} else if hhmmssms[i] == ':' {
			timecode = append(timecode, buffer)
			buffer = ""
		} else if i == 0 {
			if buffer != "" {
				timecode = append(timecode, string(hhmmssms[i])+buffer)
			} else {
				timecode = append(timecode, string(hhmmssms[i]))
			}
		} else {
			buffer = string(hhmmssms[i]) + buffer
		}
	}

	length = len(timecode)

	if length == 1 {
		ss, _ = strconv.ParseFloat(timecode[0], 64)
	} else if length == 2 {
		ss, _ = strconv.ParseFloat(timecode[0], 64)
		mm, _ = strconv.ParseFloat(timecode[1], 64)
	} else if length == 3 {
		ss, _ = strconv.ParseFloat(timecode[0], 64)
		mm, _ = strconv.ParseFloat(timecode[1], 64)
		hh, _ = strconv.ParseFloat(timecode[2], 64)
	}

	return hh*3600 + mm*60 + ss + ms
}
