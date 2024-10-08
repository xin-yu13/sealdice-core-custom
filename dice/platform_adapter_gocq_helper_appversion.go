package dice

type ProtocolType int

const (
	AndroidPhone ProtocolType = 1 // Android Phone
	AndroidPad   ProtocolType = 6 // Android Pad
)

type GocqAppVersion struct {
	ApkID           string       `json:"apk_id"`
	AppID           uint32       `json:"app_id"`
	SubAppID        uint32       `json:"sub_app_id"`
	AppKey          string       `json:"app_key"`
	SortVersionName string       `json:"sort_version_name"`
	BuildTime       uint32       `json:"build_time"`
	ApkSign         string       `json:"apk_sign"`
	SdkVersion      string       `json:"sdk_version"`
	SsoVersion      uint32       `json:"sso_version"`
	MiscBitmap      uint32       `json:"misc_bitmap"`
	MainSigMap      uint32       `json:"main_sig_map"`
	SubSigMap       uint32       `json:"sub_sig_map"`
	DumpTime        uint32       `json:"dump_time"`
	Qua             string       `json:"qua"`
	ProtocolType    ProtocolType `json:"protocol_type"`
}

var GocqAppVersionMap = map[string]map[ProtocolType]GocqAppVersion{
	"8.9.63": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537164840,
			SubAppID:        537164840,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.63.11390",
			BuildTime:       1685069178,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2546",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1687796862,
			Qua:             "V1_AND_SQ_8.9.63_4194_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537164888,
			SubAppID:        537164888,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.63.11390",
			BuildTime:       1685069178,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2546",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1687796862,
			Qua:             "V1_AND_SQ_8.9.63_4194_YYB_D",
			ProtocolType:    6,
		},
	},
	"8.9.68": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537168313,
			SubAppID:        537168313,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.68.11565",
			BuildTime:       1688523354,
			ApkSign:         "7772804f3cb4961f57cb764fbe4973e6",
			SdkVersion:      "6.0.0.2549",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1689780543,
			Qua:             "V1_AND_SQ_8.9.68_4264_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537168361,
			SubAppID:        537168361,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.68.11565",
			BuildTime:       1688523354,
			ApkSign:         "7772804f3cb4961f57cb764fbe4973e6",
			SdkVersion:      "6.0.0.2549",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1689780543,
			Qua:             "V1_AND_SQ_8.9.68_4264_YYB_D",
			ProtocolType:    6,
		},
	},
	"8.9.70": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537169928,
			SubAppID:        537169928,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.70.11730",
			BuildTime:       1689956914,
			ApkSign:         "e686fa90d9a33950c46de9cfb4ec7e71",
			SdkVersion:      "6.0.0.2551",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1690350020,
			Qua:             "V1_AND_SQ_8.9.70_4330_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537169976,
			SubAppID:        537169976,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.70.11730",
			BuildTime:       1689956914,
			ApkSign:         "e686fa90d9a33950c46de9cfb4ec7e71",
			SdkVersion:      "6.0.0.2551",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1690350020,
			Qua:             "V1_AND_SQ_8.9.70_4330_YYB_D",
			ProtocolType:    6,
		},
	},
	"8.9.73": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537171007,
			SubAppID:        537171007,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.73.11790",
			BuildTime:       1690515318,
			ApkSign:         "d4dd51c0a4a7a37f7fa9d791cd1c0377",
			SdkVersion:      "6.0.0.2553",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1690715354,
			Qua:             "V1_AND_SQ_8.9.73_4354_HDBM_T",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537171018,
			SubAppID:        537171018,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.73.11790",
			BuildTime:       1690515318,
			ApkSign:         "d4dd51c0a4a7a37f7fa9d791cd1c0377",
			SdkVersion:      "6.0.0.2553",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1690715354,
			Qua:             "V1_AND_SQ_8.9.73_4354_HDBM_T",
			ProtocolType:    6,
		},
	},
	"8.9.76": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537173477,
			SubAppID:        537173477,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.76.12115",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.76_4484_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537173525,
			SubAppID:        537173525,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.76.12115",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.76_4484_YYB_D",
			ProtocolType:    AndroidPad,
		},
	},
	"8.9.78": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537175315,
			SubAppID:        537175315,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.78.12275",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.78_4548_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537175354,
			SubAppID:        537175354,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.78.12275",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.78_4548_YYB_D",
			ProtocolType:    6,
		},
	},

	"8.9.80": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537176863,
			SubAppID:        537176863,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.80.12440",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.80_4614_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537176902,
			SubAppID:        537176902,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.80.12440",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.80_4614_YYB_D",
			ProtocolType:    6,
		},
	},

	"8.9.83": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537178646,
			SubAppID:        537178646,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.83.12605",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.83_4680_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537178685,
			SubAppID:        537178685,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.83.12605",
			BuildTime:       1691565978,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2554",
			SsoVersion:      20,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1691565978,
			Qua:             "V1_AND_SQ_8.9.83_4680_YYB_D",
			ProtocolType:    6,
		},
	},

	"8.9.85": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537180568,
			SubAppID:        537180568,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.85.12820",
			BuildTime:       1697015435,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2556",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      16724722,
			SubSigMap:       66560,
			DumpTime:        1697015435,
			Qua:             "V1_AND_SQ_8.9.85_4766_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537179875,
			SubAppID:        537179875,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.85.12820",
			BuildTime:       1697015435,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2556",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      16724722,
			SubSigMap:       66560,
			DumpTime:        1697015435,
			Qua:             "V1_AND_SQ_8.9.85_4766_YYB_D",
			ProtocolType:    6,
		},
	},
	"8.9.88": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537182769,
			SubAppID:        537182769,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.88.13035",
			BuildTime:       1697015435,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2556",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      16724722,
			SubSigMap:       66560,
			DumpTime:        1697015435,
			Qua:             "V1_AND_SQ_8.9.88_4852_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537182808,
			SubAppID:        537182808,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.88.13035",
			BuildTime:       1697015435,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2556",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      16724722,
			SubSigMap:       66560,
			DumpTime:        1697015435,
			Qua:             "V1_AND_SQ_8.9.88_4852_YYB_D",
			ProtocolType:    6,
		},
	},
	"8.9.90": {
		AndroidPhone: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537185007,
			SubAppID:        537185007,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.90.13250",
			BuildTime:       1697015435,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2556",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      16724722,
			SubSigMap:       66560,
			DumpTime:        1697015435,
			Qua:             "V1_AND_SQ_8.9.90_4938_YYB_D",
			ProtocolType:    AndroidPhone,
		},
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537185046,
			SubAppID:        537185046,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "8.9.90.13250",
			BuildTime:       1697015435,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2556",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      16724722,
			SubSigMap:       66560,
			DumpTime:        1697015435,
			Qua:             "V1_AND_SQ_8.9.90_4938_YYB_D",
			ProtocolType:    AndroidPhone,
		},
	},
	"9.0.56": {
		AndroidPad: {
			ApkID:           "com.tencent.mobileqq",
			AppID:           537220362,
			SubAppID:        537220362,
			AppKey:          "0S200MNJT807V3GE",
			SortVersionName: "9.0.56.16830",
			BuildTime:       1713424357,
			ApkSign:         "a6b745bf24a2c277527716f6f36eb68d",
			SdkVersion:      "6.0.0.2560",
			SsoVersion:      21,
			MiscBitmap:      150470524,
			MainSigMap:      34869472,
			SubSigMap:       66560,
			DumpTime:        1713424357,
			Qua:             "V1_AND_SQ_9.0.56_6372_YYB_D",
			ProtocolType:    6,
		},
	},
}
