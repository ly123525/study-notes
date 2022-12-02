#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>
#include <sys/wait.h>

int main(int argc, char const *argv[])
{
	pid_t cid; //child pid

	printf("Before fork Process id:%d\n", getpid());

	int value = 100;

	cid = fork();

	if (cid == 0){
		printf("Child process id (my parent pid is %d):%d\n", getppid(),getpid());
		for(int i = 0; i < 3; i++){
			printf("hello(%d)\n",value--);
		}
		sleep(3);
	}else{
		printf("Parent Process id: %d\n",getpid());
		for(int i = 0; i < 3; i++){
			printf("world(%d)\n",value++);
		}

		// wait(NULL); //等待子进程结束了再执行
	}

	// 变成了孤儿进程，交给系统进程托管（Pid = 1）
	return 0;
}