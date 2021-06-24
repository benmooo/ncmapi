testAll:
	go test ./... -v -cover

# package test
testCrypto:
	go test ./crypto -v -cover




# Module
testLogin:
	go test -v -run TestLogin

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

testSearch:
	go test -v -run TestSearch

testLike:
	go test -v -run TestLike

testLyric:
	go test -v -run TestLyric

testPersonalFM:
	go test -v -run TestPersonalFM

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

testSong:
	go test -v -run TestSong

testUser:
	go test -v -run TestUser

testUser:
	go test -v -run TestUser