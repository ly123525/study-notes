#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <semaphore.h>

int ticketAmout = 2;  // 票的数量: 全局变量
sem_t mutex;  // 定义信号量mutex

void* ticketAgent(void* arg){

    sem_wait(&mutex); // 执行P操作

    int t = ticketAmout;

    if (t > 0){
        printf("One ticket sold\n");
        t--;
    }else{
        printf("Ticket sold out\n");
    }

    ticketAmout = t;

    sem_post(&mutex); // 执行V操作

    pthread_exit(0);
}

int main(int argc, char const* agrv[]){

    pthread_t ticketAgent_tid[2];

    sem_init(&mutex, 0, 1);  // 初始化信号量

    for(int i = 0; i < 2; i++){
        pthread_create(ticketAgent_tid+i, NULL, ticketAgent, NULL);
    }

    for (int i = 0; i < 2; i++){
        pthread_join(ticketAgent_tid[i], NULL);
    }

    sleep(1);
    printf("The left ticket is %d\n", ticketAmout);

    sem_destroy(&mutex); // 销毁信号量

    return 0;
}