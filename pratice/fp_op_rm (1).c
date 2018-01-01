#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int file_open(char* file_name) {
	FILE* fp =NULL;
	
	fp = fopen(file_name, "w");
	if(fp == NULL) {
		printf("File open Failure!\n");
	}
	else {
		printf("File open Successfill!\n");
	}
	fclose(fp);
	
	return 0;
}

int file_remove(char* file_name){
	if(remove(file_name) == -1) {
		printf("File cannot be removed!\n");
	}
	else {
		printf("File removed!\n");
	}
	return 0;
}

int file_rename(char* old_file_name, char* new_file_name) {
	if(rename(old_file_name, new_file_name) == -1){
		printf("File CANNOT be renamed!\n");
	}
	else {
		printf("FILE renamed!\n");
	}
}

int file_w(char* file_name) {
	int i=0;
	FILE* fp=NULL;
	fp = fopen(file_name, "w");
	if(fp==NULL) {
		printf("file open failure1!\n");
	}
	else {
			fputs("IjjjN",fp);
			fputc('c',fp);
	}
	fclose(fp);
	return 0;
}

int file_r(char* file_name) {
	FILE* fp =NULL;
	char buf[256];
	fp = fopen(file_name, "r");
	if(fp == NULL) {
		printf("File open Failure3!\n");
	}
	else {
		while(fgets(buf, 256, fp)!=NULL){
			printf("%s\n", buf);
		}
	}
	fclose(fp);
	
	return 0;
}

int file_cpy(char* org_file_name, char* dst_file_name) {
	char buf[256];
	FILE* fp_r = NULL;
	FILE* fp_w = NULL;
	
	fp_r = fopen(org_file_name, "r");
	fp_w = fopen(dst_file_name, "w");
	if(fp_r == NULL && fp_w == NULL) {
		printf("File open Failure2!\n");
	}
	else {
		while(fgets(buf, 256, fp_r)!=NULL){
			fputs(buf,fp_w);
		}
	}
	fclose(fp_r);
	fclose(fp_w);
	
	return 0;
	
}

int file_search(char* file_name, char* search_s) {
	FILE* fp=NULL;
	char buf[10000];
	char buf2[10000];
	int i=0,j=0;

	fp =fopen(file_name, "r");
	if(fp == NULL) {
		printf("file open failure!\n");
	}
	else {
		while(fgets(buf,10000, fp)!=NULL) {
			//buf[strlen(buf) - 1] = '\0';
			i=0;
			while(buf[i]!='\0'){
				buf2[j] = buf[i];
				i++;
				j++;
			}
			//printf("buf2 : %s\n\n",buf2);	
			//printf("buf : %s\n\n",buf);
			//printf("%s", strstr(buf, search_s));
		}
		i=0;
		for(i=0;i<1000;i++) {
			if(buf2[i]=='\n') {
				while(buf2[i+j-1]=!'\0'){
					buf2[i+j] = buf2[i+j+1];
					j++;
				}
			
				
			}
		}
		//printf("%s", strstr(buf2, search_s));
		printf("%s\n", buf2);
	}
}

int file_copy_b(char* org_file_name, char* dst_file_name) {
	char buf[1024];
	FILE* fp_r_b;
	FILE* fp_w_b;
	int readCnt;
	
	fp_r_b = fopen(org_file_name, "rb");
	fp_w_b = fopen(dst_file_name, "wb");
	
	if(fp_w_b == NULL || fp_r_b == NULL) {
		puts("파일오픈 실패!");
		return -1; 
	}
	
    while((readCnt = fread(buf, 1, sizeof(buf), fp_r_b))>0){
		int w_count =fwrite(buf, 1, readCnt, fp_w_b);
        if(w_count <0){
        	fprintf(stderr, "파일쓰기오류\n");
			return 1;
		}
		if(w_count < readCnt) {
			fprintf(stderr, "미디어 쓰기 오류 \n");
			return 1;
		}
}

	fclose(fp_r_b); 
	fclose(fp_w_b);
	return 0;
}

int main() {
	//file_w("test.txt");
	//file_cpy("test.txt", "test_cpy.txt");
	//file_r("test_cpy.txt");
	file_search("abcd.txt", "apple");
	//file_copy_b("org_file_name.png", "cop.png");
	//file_open("test3.txt");
	//file_rename("test3.txt", "test4.txt");
	//file_remove("test3.txt");
	
	return 0;
}
