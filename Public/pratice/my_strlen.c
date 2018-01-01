#include <stdio.h>

int my_strlen(char* str) {
	
	int count=0;
	
	if(*str!=0){
		while(*str!='\0') {
			count++;
			str++;
		}
	}
	
	
	return count;
}



int main() {
	char str[8] = "abcdef";
	
	printf("%d", my_strlen(str));	
	
	
	return 0;
}
