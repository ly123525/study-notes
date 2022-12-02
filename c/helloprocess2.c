#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <sys/wait.h>

int main(int argc, char const *argv[])
{
	pid_t cid; //child pid

	printf("Before fork Process id:%d\n", getpid());

	cid = fork();

	if (cid == 0){
		printf("Child process id (my parent pid is %d):%d\n", getppid(),getpid());
		for(int i = 0; i < 3; i++){
			printf("hello\n");
		}
	}else{
		printf("Parent Process id: %d\n",getpid());
		for(int i = 0; i < 3; i++){
			printf("world\n");
		}

		wait(NULL);
	}

	return 0;
}