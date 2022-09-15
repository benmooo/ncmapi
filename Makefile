testAll:
	go test ./... -v -cover

# package test
testCrypto:
	go test ./crypto -v -cover




# Module Test
testAlbum:
	go test -v -run TestAlbum

testApi:
	go test -v -run TestApi

testArtist:
	go test -v -run TestArtist

testCheckMusic:
	go test -v -run TestCheckMusic

testComment:
	go test -v -run TestComment

testDailySign:
	go test -v -run TestDailySign

testFmTrash:
	go test -v -run TestFmTrash

testLike:
	go test -v -run TestLike

testLogin:
	go test -v -run TestLoginPhone

testLogout:
	go test -v -run TestLogout

testLyric:
	go test -v -run TestLyric

testPersonalFm:
	go test -v -run TestPersonalFm

testPlaylist:
	go test -v -run TestPlaylist

testRecommend:
	go test -v -run TestRecommend

testRelated:
	go test -v -run TestRelated

testScrobble:
	go test -v -run TestScrobble

testSearch:
	go test -v -run TestSearch

testSimi:
	go test -v -run TestSimi

testSong:
	go test -v -run TestSong

testUser:
	go test -v -run TestUser
