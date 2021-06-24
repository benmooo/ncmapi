package apitypes

type ResourceType uint8

const (
	ResourceTypeSong ResourceType = iota
	ResourceTypeMV
	ResourceTypeCollection
	ResourceTypeAlbum
	ResourceTypePodcast
	ResourceTypeVideo
	ResourceTypeMoment
)

type ResourceTypeCode string

const (
	ResourceTypeCodeSong    ResourceTypeCode = "R_SO_4_"
	ResourceTypeCodeMV      ResourceTypeCode = "R_MV_5_"
	ResourceTypeCodeCollect ResourceTypeCode = "A_PL_0_"
	ResourceTypeCodeAlbum   ResourceTypeCode = "R_AL_3_"
	ResourceTypeCodePodcast ResourceTypeCode = "A_DJ_1_"
	ResourceTypeCodeVideo   ResourceTypeCode = "R_VI_62_"
	ResourceTypeCodeMoment  ResourceTypeCode = "A_EV_2_"
)
