from tCafeteria import *
from datetime import *
import tweepy

print("Parsing cafeteria info...")
cafe = tCafeteria("G100000479", 'DAEJEON', 'MIDDLE')	#탄방중 학교코드, 지역(대전), 학교 종류(중학교)
meal = cafe.parseCafeteria()

# Dictionary+list형식의 return값을 하나의 문자열로 변환
print("Formating data...")
res = str(datetime.today().day) + "일 급식입니다\n\n"

try:
	print(meal['error'][0])
except KeyError:
	for i in range(len(meal['lunch'])):
		res = res + meal['lunch'][i][0] + " " + meal['lunch'][i][1] + '\n'
	
consumerKey = "your consumer key"
consumerSecret = "your consumer secret"
 
#auth.OAuthHandler 객체 반환
print("Setting OAuthHandler...")
auth = tweepy.OAuthHandler(consumerKey, consumerSecret)

accessToken = "your access token"
accessTokenSecret = "your access token secret"
 
#auth.OAuthHandler r객체에 엑세스토큰 지정
print("Setting access token...")
auth.set_access_token(accessToken, accessTokenSecret)

#API 클래스의 인스턴스 반환 - 읽기, 트윗, 리트윗, DM
print("Logging in...")
api = tweepy.API(auth)

print("Uploading... \n")
api.update_status(status=res)
print(res)
print("Done!")