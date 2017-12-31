#include <stdio.h>

int my_strcpy(char *dst, char *src) {
	int i=0;
	while(src[i] != NULL) {
		if(src[i] == 'a') {
		dst[i] = src[i];
		}
		i++;
	}
	return 0;
}

int main() {
	char src[32] = "Halla aoala";
	char dst[32];
	
	my_strcpy(dst,src);
	printf("%s\n", dst);
	
	return 0;
}
