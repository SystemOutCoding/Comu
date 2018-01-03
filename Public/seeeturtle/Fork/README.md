:fork_and_knife:Fork
==========
Plate의 서버입니다.

Setting
--------
config-template/ 을 config/ 로 이름을 바꾼후,


config-template.go를 config.go로 이름을 바꾼뒤,


config.go의 패키지 이름을 configTemplate에서 config로 바꾸어주세요.


DB 설정에 알맞게 DBConfig의 필드 값을 바꾸어주세요.

Install
--------
```
$ go get -u github.com/joshua1b/Fork
```
위 명령어로 저장소를 클론하실 수 있습니다.

Run
------
```
$ go build ./
$ ./Fork
```
저장소의 최상위 디렉토리에서 실행하시면 됩니다.

Warning
--------
[Spoon](https://github.com/joshua1b/Spoon)로 미리 데이터베이스를 초기화 해야합니다.

Thanks
-------
[@mingrammer](https://github.com/mingrammer/)님


https://github.com/mingrammer/go-todo-rest-api-example


멋진 예제를 작성해주셔서 감사합니다! 이 프로젝트에 많은 도움을 주셨습니다.

문의
--------
issue에 올려주시기를 바랍니다.
