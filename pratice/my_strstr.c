#include <stdio.h>
#include <string.h>

char* my_strstr(char* str1, char* str2, int n) {
	char* p=NULL;
	int count=0;
	
	p=str2;
	
	if(*str1!='\0'&& *str2!='\0'){
		while(*str1!='\0'){
			if(*str1==*str2) {
				str2++;
				count++;
			}
			else {
				str2=p;
				count=0;
			}
			str1++;
			if(count==n){
				return str1-n;
			}
		}
	}
	
	return NULL;
}

int main() {
	char str1[8] = "abcdef";
	char str2[4] = "ce";
	char* p=NULL;
	
	p = my_strstr(str1, str2, strlen(str2));
	printf("%s", p);
	
	
	
	return 0;
}
