# -*- coding: utf-8 -*- 

# Copyright (c) 2017 w3bn00b
# See the file LICENSE for copying permission.
# example.py v1.3
# author : w3bn00b
# description : 학교 급식 정보, 학사일정을 파싱해와 각각 meal.txt, sche.txt에 저장합니다
# Usage : cafe = tCafeteria("학교코드", '관할지역 코드')
# Required library : datetime, bs4, requests

# 관할지역 정보
# 'SEOUL':'stu.sen.go.kr',
# 'INCHEON':'stu.ice.go.kr',
# 'BUSAN':'stu.pen.go.kr',
# 'GWANGJU':'stu.gen.go.kr',
# 'DAEJEON':'stu.dje.go.kr',
# 'DAEGU':'stu.dge.go.kr',
# 'SEJONG':'stu.sje.go.kr',
# 'ULSAN':'stu.use.go.kr',
# 'GYEONGGI':'stu.goe.go.kr',
# 'KANGWON':'stu.kwe.go.kr',
# 'CHUNGBUK':'stu.cbe.go.kr',
# 'CHUNGNAM':'stu.cne.go.kr',
# 'GYEONGBUK':'stu.gbe.go.kr',
# 'GYEONGNAM':'stu.gne.go.kr',
# 'JEONBUK':'stu.jbe.go.kr',
# 'JEONNAM':'stu.jne.go.kr',
# 'JEJU':'stu.jje.go.kr'

# 학교 종류
# 'KINDERGARTEN':'1'
# 'ELEMENTARY':'2'
# 'MIDDLE':'3'
# 'HIGH':'4'

from tCafeteria import *

cafe = tCafeteria("G100000479", 'DAEJEON', 'MIDDLE')	#탄방중 학교코드, 지역(대전), 학교 종류(중학교)
res = cafe.parseCafeteria()
print (res)

res = cafe.parseSchedule()
print (res)