#include <stdio.h>

char* my_strnchr_r(char* src, char c, int n) {
	char* p;
	int counter;
	
	p = src;
	counter = 0;
	
	if(*src!='\0'){
		while(*src!='\0'){
			src++;
		}
		src--;
		
		while(*src!=*p) {
			if(*src == c) {
				counter ++;
				if(counter==n) return src;
				else src--;
			}
			else{
				src--;
			}
		}
	}
	return src;
		
}
 
int main() {
	char str[8] = "adeedee";
	char c= 'e';
	char* p = NULL;
	
	p = my_strnchr_r(str, c, 2);
	printf("%s\n", p);
	
	return 0;
}
