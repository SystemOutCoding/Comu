function dpTime(){ //wtf
  let now = new Date();
  let hours = now.getHours();
  let minutes = now.getMinutes();
  let seconds = now.getSeconds();

  let dates = now.getDate();
  let months = now.getMonth();
  let days = now.getDay();
  let ampm = "오후 ";

  if (hours > 12)
      hours -= 12;
  else
      ampm = "오전 ";

  months++; //0~11 -> 1~12

  days = setDays(days);
  render(months, dates, days, ampm, hours, minutes, seconds);
}

function setDays(days) {
  if (days == 1)
    return "월요일";
  if (days == 2)
    return "화요일";
  if (days == 3)
    return "수요일";
  if (days == 4)
    return "목요일";
  if (days == 5)
    return "금요일";
  if (days == 6)
    return "토요일";
  if (days == 0)
    return "일요일";
}

function zeroFill(number) {
  if (number < 10)
    return "0" + number;
  else
    return number;
}

function render(months, dates, days, ampm, hours, minutes, seconds) {
  var CLOCK = document.getElementById("count").innerHTML = months + "월 "+ dates + "일 "+ days + " "+ ampm+ zeroFill(hours)+ ":" + zeroFill(minutes)+ ":" + zeroFill(seconds);
}

setInterval(() => dpTime(), 1000);
