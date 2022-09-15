package ncmapi

// This APIRoutes is generated by routes.sh
var APIRoutes = map[string]string{
	"activate_init_profile":     "https://music.163.com/eapi/activate/initProfile",
	"album_detail_dynamic":      "https://music.163.com/api/album/detail/dynamic",
	"album_detail":              "https://music.163.com/weapi/vipmall/albumproduct/detail",
	"album":                     "https://music.163.com/weapi/v1/album/${query.id}",
	"album_list":                "https://music.163.com/weapi/vipmall/albumproduct/list",
	"album_list_style":          "https://music.163.com/weapi/vipmall/appalbum/album/style",
	"album_newest":              "https://music.163.com/api/discovery/newAlbum",
	"album_new":                 "https://music.163.com/weapi/album/new",
	"album_songsaleboard":       "https://music.163.com/api/feealbum/songsaleboard/${type}/type",
	"album_sub":                 "https://music.163.com/api/album/${query.t}",
	"album_sublist":             "https://music.163.com/weapi/album/sublist",
	"artist_album":              "https://music.163.com/weapi/artist/albums/${query.id}",
	"artist_desc":               "https://music.163.com/weapi/artist/introduction",
	"artist_detail":             "https://music.163.com/api/artist/head/info/get",
	"artist_fans":               "https://music.163.com/weapi/artist/fans/get",
	"artist_list":               "https://music.163.com/api/v1/artist/list",
	"artist_mv":                 "https://music.163.com/weapi/artist/mvs",
	"artist_new_mv":             "https://music.163.com/api/sub/artist/new/works/mv/list",
	"artist_new_song":           "https://music.163.com/api/sub/artist/new/works/song/list",
	"artists":                   "https://music.163.com/weapi/v1/artist/${query.id}",
	"artist_songs":              "https://music.163.com/api/v1/artist/songs",
	"artist_sub":                "https://music.163.com/weapi/artist/${query.t}",
	"artist_sublist":            "https://music.163.com/weapi/artist/sublist",
	"artist_top_song":           "https://music.163.com/api/artist/top/song",
	"audio_match":               "https://music.163.com/api/music/audio/match",
	"avatar_upload":             "https://music.163.com/weapi/user/avatar/upload/v1",
	"banner":                    "https://music.163.com/api/v2/banner/get",
	"batch":                     "https://music.163.com/eapi/batch",
	"calendar":                  "https://music.163.com/api/mcalendar/detail",
	"captcha_sent":              "https://music.163.com/api/sms/captcha/sent",
	"captcha_verify":            "https://music.163.com/weapi/sms/captcha/verify",
	"cellphone_existence_check": "https://music.163.com/eapi/cellphone/existence/check",
	"check_music":               "https://music.163.com/weapi/song/enhance/player/url",
	"cloud":                     "https://interface.music.163.com/api/cloud/upload/check",
	// ""https"://music.163.com/weapi/nos/token/alloc",
	// ""https"://music.163.com/api/upload/cloud/info/v2",
	// ""https"://interface.music.163.com/api/cloud/pub/v2",
	"cloud_match":                      "https://music.163.com/api/cloud/user/song/match",
	"cloudsearch":                      "https://music.163.com/api/cloudsearch/pc",
	"comment_album":                    "https://music.163.com/weapi/v1/resource/comments/R_AL_3_${query.id}",
	"comment_dj":                       "https://music.163.com/weapi/v1/resource/comments/A_DJ_1_${query.id}",
	"comment_event":                    "https://music.163.com/weapi/v1/resource/comments/${query.threadId}",
	"comment_floor":                    "https://music.163.com/api/resource/comment/floor/get",
	"comment_hot":                      "https://music.163.com/weapi/v1/resource/hotcomments/${query.type}${query.id}",
	"comment_hug_list":                 "https://music.163.com/api/v2/resource/comments/hug/list",
	"comment":                          "https://music.163.com/weapi/resource/comments/${query.t}",
	"comment_like":                     "https://music.163.com/weapi/v1/comment/${query.t}",
	"comment_music":                    "https://music.163.com/api/v1/resource/comments/R_SO_4_${query.id}",
	"comment_mv":                       "https://music.163.com/weapi/v1/resource/comments/R_MV_5_${query.id}",
	"comment_new":                      "https://music.163.com/api/v2/resource/comments",
	"comment_playlist":                 "https://music.163.com/weapi/v1/resource/comments/A_PL_0_${query.id}",
	"comment_video":                    "https://music.163.com/weapi/v1/resource/comments/R_VI_62_${query.id}",
	"countries_code_list":              "https://interface3.music.163.com/eapi/lbs/countries/v1",
	"daily_signin":                     "https://music.163.com/weapi/point/dailyTask",
	"digitalAlbum_detail":              "https://music.163.com/weapi/vipmall/albumproduct/detail",
	"digitalAlbum_ordering":            "https://music.163.com/api/ordering/web/digital",
	"digitalAlbum_purchased":           "https://music.163.com/api/digitalAlbum/purchased",
	"digitalAlbum_sales":               "https://music.163.com/weapi/vipmall/albumproduct/album/query/sales",
	"dj_banner":                        "https://music.163.com/weapi/djradio/banner/get",
	"dj_category_excludehot":           "https://music.163.com/weapi/djradio/category/excludehot",
	"dj_category_recommend":            "https://music.163.com/weapi/djradio/home/category/recommend",
	"dj_catelist":                      "https://music.163.com/weapi/djradio/category/get",
	"dj_detail":                        "https://music.163.com/api/djradio/v2/get",
	"dj_hot":                           "https://music.163.com/weapi/djradio/hot/v1",
	"dj_paygift":                       "https://music.163.com/weapi/djradio/home/paygift/list?_nmclfl=1",
	"dj_personalize_recommend":         "https://music.163.com/api/djradio/personalize/rcmd",
	"dj_program_detail":                "https://music.163.com/api/dj/program/detail",
	"dj_program":                       "https://music.163.com/weapi/dj/program/byradio",
	"dj_program_toplist_hours":         "https://music.163.com/api/djprogram/toplist/hours",
	"dj_program_toplist":               "https://music.163.com/api/program/toplist/v1",
	"dj_radio_hot":                     "https://music.163.com/api/djradio/hot",
	"dj_recommend":                     "https://music.163.com/weapi/djradio/recommend/v1",
	"dj_recommend_type":                "https://music.163.com/weapi/djradio/recommend",
	"dj_sub":                           "https://music.163.com/weapi/djradio/${query.t}",
	"dj_sublist":                       "https://music.163.com/weapi/djradio/get/subed",
	"dj_subscriber":                    "https://music.163.com/api/djradio/subscriber",
	"dj_today_perfered":                "https://music.163.com/weapi/djradio/home/today/perfered",
	"dj_toplist_hours":                 "https://music.163.com/api/dj/toplist/hours",
	"dj_toplist":                       "https://music.163.com/api/djradio/toplist",
	"dj_toplist_newcomer":              "https://music.163.com/api/dj/toplist/newcomer",
	"dj_toplist_pay":                   "https://music.163.com/api/djradio/toplist/pay",
	"dj_toplist_popular":               "https://music.163.com/api/dj/toplist/popular",
	"event_del":                        "https://music.163.com/eapi/event/delete",
	"event_forward":                    "https://music.163.com/weapi/event/forward",
	"event":                            "https://music.163.com/weapi/v1/event/get",
	"fm_trash":                         "https://music.163.com/weapi/radio/trash/add?alg=RT&songId=%v&time=%v",
	"follow":                           "https://music.163.com/weapi/user/${query.t}/${query.id}",
	"history_recommend_songs_detail":   "https://music.163.com/api/discovery/recommend/songs/history/detail",
	"history_recommend_songs":          "https://music.163.com/api/discovery/recommend/songs/history/recent",
	"hot_topic":                        "https://music.163.com/api/act/hot",
	"hug_comment":                      "https://music.163.com/api/v2/resource/comments/hug/listener",
	"like":                             "https://music.163.com/api/radio/like",
	"likelist":                         "https://music.163.com/weapi/song/like/get",
	"listen_together_status":           "https://music.163.com/api/listen/together/status/get",
	"login_cellphone":                  "https://music.163.com/api/login/cellphone",
	"login":                            "https://music.163.com/weapi/login",
	"login_qr_check":                   "https://music.163.com/weapi/login/qrcode/client/login",
	"login_qr_create":                  "https://music.163.com/login?codekey=${query.key}",
	"login_qr_key":                     "https://music.163.com/weapi/login/qrcode/unikey",
	"login_refresh":                    "https://music.163.com/weapi/login/token/refresh",
	"login_status":                     "https://music.163.com/weapi/w/nuser/account/get",
	"logout":                           "https://music.163.com/weapi/logout",
	"lyric":                            "https://music.163.com/api/song/lyric",
	"mlog_to_video":                    "https://music.163.com/weapi/mlog/video/convert/id",
	"mlog_url":                         "https://music.163.com/weapi/mlog/detail/v1",
	"msg_comments":                     "https://music.163.com/api/v1/user/comments/${query.uid}",
	"msg_forwards":                     "https://music.163.com/api/forwards/get",
	"msg_notices":                      "https://music.163.com/api/msg/notices",
	"msg_private_history":              "https://music.163.com/api/msg/private/history",
	"msg_private":                      "https://music.163.com/api/msg/private/users",
	"msg_recentcontact":                "https://music.163.com/api/msg/recentcontact/get",
	"musician_cloudbean":               "https://music.163.com/weapi/cloudbean/get",
	"musician_cloudbean_obtain":        "https://music.163.com/weapi/nmusician/workbench/mission/reward/obtain/new",
	"musician_data_overview":           "https://music.163.com/weapi/creator/musician/statistic/data/overview/get",
	"musician_play_trend":              "https://music.163.com/weapi/creator/musician/play/count/statistic/data/trend/get",
	"musician_tasks":                   "https://music.163.com/weapi/nmusician/workbench/mission/cycle/list",
	"mv_all":                           "https://interface.music.163.com/api/mv/all",
	"mv_detail_info":                   "https://music.163.com/api/comment/commentthread/info",
	"mv_detail":                        "https://music.163.com/api/v1/mv/detail",
	"mv_exclusive_rcmd":                "https://interface.music.163.com/api/mv/exclusive/rcmd",
	"mv_first":                         "https://interface.music.163.com/weapi/mv/first",
	"mv_sub":                           "https://music.163.com/weapi/mv/${query.t}",
	"mv_sublist":                       "https://music.163.com/weapi/cloudvideo/allvideo/sublist",
	"mv_url":                           "https://music.163.com/weapi/song/enhance/play/mv/url",
	"personal_fm":                      "https://music.163.com/weapi/v1/radio/get",
	"personalized_djprogram":           "https://music.163.com/weapi/personalized/djprogram",
	"personalized":                     "https://music.163.com/weapi/personalized/playlist",
	"personalized_mv":                  "https://music.163.com/weapi/personalized/mv",
	"personalized_newsong":             "https://music.163.com/api/personalized/newsong",
	"personalized_privatecontent":      "https://music.163.com/weapi/personalized/privatecontent",
	"personalized_privatecontent_list": "https://music.163.com/api/v2/privatecontent/list",
	"playlist_catlist":                 "https://music.163.com/weapi/playlist/catalogue",
	"playlist_cover_update":            "https://music.163.com/weapi/playlist/cover/update",
	"playlist_create":                  "https://music.163.com/api/playlist/create",
	"playlist_delete":                  "https://music.163.com/weapi/playlist/remove",
	"playlist_desc_update":             "https://interface3.music.163.com/eapi/playlist/desc/update",
	"playlist_detail_dynamic":          "https://music.163.com/api/playlist/detail/dynamic",
	"playlist_detail":                  "https://music.163.com/api/v6/playlist/detail",
	"playlist_highquality_tags":        "https://music.163.com/api/playlist/highquality/tags",
	"playlist_hot":                     "https://music.163.com/weapi/playlist/hottags",
	"playlist_mylike":                  "https://music.163.com/api/mlog/playlist/mylike/bytime/get",
	"playlist_name_update":             "https://interface3.music.163.com/eapi/playlist/update/name",
	"playlist_order_update":            "https://music.163.com/api/playlist/order/update",
	"playlist_subscribe":               "https://music.163.com/weapi/playlist/${query.t}",
	"playlist_subscribers":             "https://music.163.com/weapi/playlist/subscribers",
	"playlist_tags_update":             "https://interface3.music.163.com/eapi/playlist/tags/update",
	"playlist_track_add":               "https://music.163.com/api/playlist/track/add",
	"playlist_track_delete":            "https://music.163.com/api/playlist/track/delete",
	"playlist_tracks":                  "https://music.163.com/api/playlist/manipulate/tracks",
	// ""https"://music.163.com/api/playlist/manipulate/tracks",
	"playlist_update":            "https://music.163.com/weapi/batch",
	"playlist_video_recent":      "https://music.163.com/api/playlist/video/recent",
	"playmode_intelligence_list": "https://music.163.com/weapi/playmode/intelligence/list",
	"program_recommend":          "https://music.163.com/weapi/program/recommend/v1",
	"rebind":                     "https://music.163.com/api/user/replaceCellphone",
	"recommend_resource":         "https://music.163.com/weapi/v1/discovery/recommend/resource",
	"recommend_songs":            "https://music.163.com/api/v3/discovery/recommend/songs",
	"register_cellphone":         "https://music.163.com/api/register/cellphone",
	"related_allvideo":           "https://music.163.com/weapi/cloudvideo/v1/allvideo/rcmd",
	"related_playlist":           "https://music.163.com/playlist?id=${query.id}",
	"resource_like":              "https://music.163.com/weapi/resource/${query.t}",
	"scrobble":                   "https://music.163.com/weapi/feedback/weblog",
	"search_default":             "https://interface3.music.163.com/eapi/search/defaultkeyword/get",
	"search_hot_detail":          "https://music.163.com/weapi/hotsearchlist/get",
	"search_hot":                 "https://music.163.com/weapi/search/hot",
	"search":                     "https://music.163.com/weapi/search/get",
	"search_multimatch":          "https://music.163.com/weapi/search/suggest/multimatch",
	"search_suggest":             "https://music.163.com/weapi/search/suggest/",
	"send_album":                 "https://music.163.com/api/msg/private/send",
	"send_playlist":              "https://music.163.com/weapi/msg/private/send",
	"send_song":                  "https://music.163.com/api/msg/private/send",
	"send_text":                  "https://music.163.com/weapi/msg/private/send",
	"setting":                    "https://music.163.com/api/user/setting",
	"share_resource":             "https://music.163.com/weapi/share/friends/resource",
	"simi_artist":                "https://music.163.com/weapi/discovery/simiArtist",
	"simi_mv":                    "https://music.163.com/weapi/discovery/simiMV",
	"simi_playlist":              "https://music.163.com/weapi/discovery/simiPlaylist",
	"simi_song":                  "https://music.163.com/weapi/v1/discovery/simiSong",
	"simi_user":                  "https://music.163.com/weapi/discovery/simiUser",
	"song_detail":                "https://music.163.com/api/v3/song/detail",
	// "song_order_update": ",
	"song_purchased":           "https://music.163.com/weapi/single/mybought/song/list",
	"song_url":                 "https://interface3.music.163.com/eapi/song/enhance/player/url",
	"top_album":                "https://music.163.com/api/discovery/new/albums/area",
	"top_artists":              "https://music.163.com/weapi/artist/top",
	"topic_detail_event_hot":   "https://music.163.com/api/act/event/hot",
	"topic_detail":             "https://music.163.com/api/act/detail",
	"topic_sublist":            "https://music.163.com/api/topic/sublist",
	"toplist_artist":           "https://music.163.com/weapi/toplist/artist",
	"toplist_detail":           "https://music.163.com/weapi/toplist/detail",
	"top_list":                 "https://interface3.music.163.com/api/playlist/v4/detail",
	"toplist":                  "https://music.163.com/api/toplist",
	"top_mv":                   "https://music.163.com/weapi/mv/toplist",
	"top_playlist_highquality": "https://music.163.com/api/playlist/highquality/list",
	"top_playlist":             "https://music.163.com/weapi/playlist/list",
	"top_song":                 "https://music.163.com/weapi/v1/discovery/new/songs",
	"user_account":             "https://music.163.com/api/nuser/account/get",
	"user_audio":               "https://music.163.com/weapi/djradio/get/byuser",
	"user_bindingcellphone":    "https://music.163.com/api/user/bindingCellphone",
	"user_binding":             "https://music.163.com/api/v1/user/bindings/${query.uid}",
	"user_cloud_del":           "https://music.163.com/weapi/cloud/del",
	"user_cloud_detail":        "https://music.163.com/weapi/v1/cloud/get/byids",
	"user_cloud":               "https://music.163.com/api/v1/cloud/get",
	"user_comment_history":     "https://music.163.com/api/comment/user/comment/history",
	"user_detail":              "https://music.163.com/weapi/v1/user/detail/${query.uid}",
	"user_dj":                  "https://music.163.com/weapi/dj/program/${query.uid}",
	"user_event":               "https://music.163.com/api/event/get/${query.uid}",
	"user_followeds":           "https://music.163.com/eapi/user/getfolloweds/${query.uid}",
	"user_follows":             "https://music.163.com/weapi/user/getfollows/${query.uid}",
	"user_level":               "https://music.163.com/weapi/user/level",
	"user_playlist":            "https://music.163.com/api/user/playlist",
	"user_record":              "https://music.163.com/weapi/v1/play/record",
	"user_replacephone":        "https://music.163.com/api/user/replaceCellphone",
	"user_subcount":            "https://music.163.com/weapi/subcount",
	"user_update":              "https://music.163.com/weapi/user/profile/update",
	"video_category_list":      "https://music.163.com/api/cloudvideo/category/list",
	"video_detail_info":        "https://music.163.com/api/comment/commentthread/info",
	"video_detail":             "https://music.163.com/weapi/cloudvideo/v1/video/detail",
	"video_group":              "https://music.163.com/api/videotimeline/videogroup/otherclient/get",
	"video_group_list":         "https://music.163.com/api/cloudvideo/group/list",
	"video_sub":                "https://music.163.com/weapi/cloudvideo/video/${query.t}",
	"video_timeline_all":       "https://music.163.com/api/videotimeline/otherclient/get",
	"video_timeline_recommend": "https://music.163.com/api/videotimeline/get",
	"video_url":                "https://music.163.com/weapi/cloudvideo/playurl",
	"vip_growthpoint_details":  "https://music.163.com/weapi/vipnewcenter/app/level/growth/details",
	"vip_growthpoint_get":      "https://music.163.com/weapi/vipnewcenter/app/level/task/reward/get",
	"vip_growthpoint":          "https://music.163.com/weapi/vipnewcenter/app/level/growhpoint/basic",
	"vip_tasks":                "https://music.163.com/weapi/vipnewcenter/app/level/task/list",
	"weblog":                   "https://music.163.com/weapi/feedback/weblog",
	"yunbei_expense":           "https://music.163.com/store/api/point/expense",
	"yunbei_info":              "https://music.163.com/api/v1/user/info",
	"yunbei":                   "https://music.163.com/api/point/signed/get",
	"yunbei_rcmd_song_history": "https://music.163.com/weapi/yunbei/rcmd/song/history/list",
	"yunbei_rcmd_song":         "https://music.163.com/weapi/yunbei/rcmd/song/submit",
	"yunbei_receipt":           "https://music.163.com/store/api/point/receipt",
	"yunbei_sign":              "https://music.163.com/api/point/dailyTask",
	"yunbei_task_finish":       "https://music.163.com/api/usertool/task/point/receive",
	"yunbei_tasks":             "https://music.163.com/api/usertool/task/list/all",
	"yunbei_tasks_todo":        "https://music.163.com/api/usertool/task/todo/query",
	"yunbei_today":             "https://music.163.com/api/point/today/get",
}
