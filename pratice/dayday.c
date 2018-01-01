#include <stdio.h>

struct data {
	int year;
	int month;
	int day;
};

int d_day(struct data p1, struct data p2) {
	int month[12] = { 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
	int i=0;
	int day=0;
	
	for(i=0;p1.year+i<p2.year;i++) {
		if((((p1.year+i)%4==0)&&((p1.year+i)%100!=0))||((p1.year+i)%400==0)) {
			day = day + 366;
		}
		else {
			day = day + 365;
		}
	}
	
	printf("%d", day);
	
	return 0;
}

int main() {
	struct data d1;
	struct data d2;
	
	d1.day = 29;
	d1.month = 1;
	d1.year = 1999;
	
	d2.day = 7;
	d2.month = 11;
	d2.year = 2017;
	
	d_day(d1, d2);
	
	return 0;
}
