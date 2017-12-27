import tCafeteria

funcSc = tCafeteria.tCafeteria("G100000202","DAEJEON","HIGH")
#funcSc = tCafeteria.tCafeteria("G100000479", 'DAEJEON', 'MIDDLE')

debug = False

response = [
    funcSc.parseCafeteria(return_all=debug),
    funcSc.parseSchedule(return_all=debug)
]

print(response[0])
print(response[1])