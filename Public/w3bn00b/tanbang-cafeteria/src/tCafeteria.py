# -*- coding: utf-8 -*- 

# Copyright (c) 2017 w3bn00b, return0927
# See the file LICENSE for copying permission.
# tanbang-cafeteria v1.3
# started by : w3bn00b | re-arranged by : return0927
# description : 중학교의 급식정보와 학사일정을 파싱해옵니다
# Usage : cafe = tCafeteria("학교코드", "관할지역 코드")
# Required library : datetime, bs4, requests

from datetime import *

# 자동 설치 메소드
try:
    from bs4 import BeautifulSoup
    import requests, re
except ImportError:
    try:
        print("Auto Installing requests, bs4")
        import pip
        pip.main(['install','bs4'])
        pip.main(['install','requests'])

        import requests, re
        from bs4 import BeautifulSoup
    except:
        raise ImportError("Error on importing modules")

fetchDate = lambda: [datetime.today().year, datetime.today().month, datetime.today().day]

class tCafeteria:
    locale = None
    schType = None
    region = {
        'SEOUL': 'stu.sen.go.kr',
        'INCHEON': 'stu.ice.go.kr',
        'BUSAN': 'stu.pen.go.kr',
        'GWANGJU': 'stu.gen.go.kr',
        'DAEJEON': 'stu.dje.go.kr',
        'DAEGU': 'stu.dge.go.kr',
        'SEJONG': 'stu.sje.go.kr',
        'ULSAN': 'stu.use.go.kr',
        'GYEONGGI': 'stu.goe.go.kr',
        'KANGWON': 'stu.kwe.go.kr',
        'CHUNGBUK': 'stu.cbe.go.kr',
        'CHUNGNAM': 'stu.cne.go.kr',
        'GYEONGBUK': 'stu.gbe.go.kr',
        'GYEONGNAM': 'stu.gne.go.kr',
        'JEONBUK': 'stu.jbe.go.kr',
        'JEONNAM': 'stu.jne.go.kr',
        'JEJU': 'stu.jje.go.kr'
    }

    Type = {
        'KINDERGARTEN': '1',  # 병설유치원
        'ELEMENTARY': '2',  # 초등학교
        'MIDDLE': '3',  # 중학교
        'HIGH': '4'  # 고등학교
    }

    def __init__(self, schoolcode, locale, schType):
        self.schoolcode = schoolcode  # 학교코드 설정
        self.locale = locale
        self.schType = schType

    def parseAlergic(self, data):
        p = re.compile("[0-9]+[.]") # 알레르기 표기 추출용 정규표현식
        alg = "".join(p.findall(data)) # 정규표현식으로 알레르기 정보 추출
        _temp = data.replace(alg, '')
        return [_temp, alg] # 알레르기 정보가 없으면 alg는 ''를 리턴

    def makeValue(self, tag):
        _TEMP = str(tag)[5:-6].split("<br/>")

        if "[석식]" in _TEMP[1:]: # 석식이 있는지 확인
            _INDEX = _TEMP[1:].index("[석식]")
            _Lunch = _TEMP[1:][:_INDEX]
            _Dinner = _TEMP[1:][_INDEX:]
            self.parseAlergic(_Lunch[1])

            return _TEMP[0], True, [
                [self.parseAlergic(meal) for meal in _Lunch[1:]],
                [self.parseAlergic(meal) for meal in _Dinner[1:]]
            ]

        else:
            return _TEMP[0], False, [
                [self.parseAlergic(meal) for meal in _TEMP[1:][1:]]
            ]

    def parseCafeteria(self, date=fetchDate(), return_all=False ):
        """
        :param date: ['yyyy','m or mm','d or dd']
        :param return_all: whether return all meals or not
        :return: all meals dict or one meal string
        """

        _YEAR, _MONTH, _DAY = date
        try:
            url = "http://%s/sts_sci_md00_001.do?schulCode=%s&schulCrseScCode=%s" + \
                  "&schulKndScCode=0%s&schMmealScCode=1&&ay=%d&mm=%d"
            url = url%\
                  (
                      self.region[self.locale],
                      self.schoolcode,
                      self.Type[self.schType],
                      self.Type[self.schType],
                      _YEAR,
                      _MONTH
                  )
            r = requests.get(url)
        except: # HTTP GET 오류 raise
            raise Exception("Error on getting server information")
        soup = BeautifulSoup(r.text, "html.parser")

        try:
            res = soup.select("#contents > div > table > tbody > tr > td > div")
            retDict = {}
            retDict[_YEAR] = {}; retDict[_YEAR][_MONTH] = {}

            for tag in res:
                if tag.find("br"):
                    _day, _type, _meal = self.makeValue(tag)

                    if _type:
                        retDict[_YEAR][_MONTH][int(_day)] = {
                            "lunch": _meal[0],
                            "dinner": _meal[1]
                        }
                    else:
                        retDict[_YEAR][_MONTH][int(_day)] = {
                            "lunch": _meal[0]
                        }
                else:
                    if tag.text.isdigit():
                        retDict[_YEAR][_MONTH][int(tag.text)] = {
                            "error": ['정보가 없습니다.','']
                        }
                    else:
                        pass

        except: # 파싱 오류 raise
            raise Exception("Error on parsing data")

        if return_all: return retDict
        else:
            if _DAY in retDict[_YEAR][_MONTH].keys():
                return retDict[_YEAR][_MONTH][_DAY]
            else:
                return {"error":['정보가 없습니다','']}

    def parseSchedule(self, date=fetchDate(), return_all=False):
        """
        :param date: ['yyyy','m or mm','d or dd']
        :param return_all: whether return all schedule or not
        :return: schedule dict or one schedule string
        """

        _YEAR, _MONTH, _DAY = date

        url = "http://%s/sts_sci_sf00_001.do?schulCode=%s"\
              "&schulCrseScCode=%s&schulKndScCode=0%s&ay=%d&mm=%d"

        url = url % \
              (
                  self.region[self.locale],
                  self.schoolcode,
                  self.Type[self.schType],
                  self.Type[self.schType],
                  _YEAR,
                  _MONTH
              )

        r = requests.get(url)  # html코드를 불러온다
        soup = BeautifulSoup(r.text, "html.parser")
        del r

        rows = soup.select("tr") # INDEX [0] 월 / [1:] 일정
        schedules = {}

        months = [ tag.text for tag in rows[0].find_all("th", attrs={"colspan":"2"}) ]
        months = [ int(x[:-1]) for x in months ]
        docYear = soup.select_one("select#grade").text.strip() # 받아 온 문서 상의 연도
        docYear = int(docYear)
        schedules[docYear] = {}

        for month in months:
            schedules[docYear][month] = {} # 각 별로 딕셔너리 지정

        for j in range(len(rows[1:])):
            tr = rows[1:][j]
            _legacy = [x.text.strip() for x in tr.find_all("td", attrs={"class": "textL"})] # 임시로 배열 저장
            for n in range(len(months)):
                try:
                    schedules[docYear][months[n]][j+1] = _legacy[n] # 스케쥴 등록
                except:
                    raise Exception(schedules) # 에러 시 지금까지 모은 schedules 리턴

        if return_all: return schedules
        else: return schedules[_YEAR][_MONTH][_DAY]

