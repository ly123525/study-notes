#include <stdio.h>
#include <pthread.h>
#include <unistd.h>
#include <semaphore.h>

sem_t road;

void* cars(void* argc){
    printf("(%u) I INTEND to pass the fork\n",pthread_self());
    sleep(1);

    sem_wait(&road); // 执行P操作

    printf("(%u) I am AT the fork\n",pthread_self());
    sleep(1);

    printf("(%u) I have PASSED the fork\n",pthread_self());
    sleep(1);

    sem_post(&road); // 执行V操作

    pthread_exit(0);
}

int main(int argc, char const* agrv[]){

    pthread_t tid[5];

    sem_init(&road, 0, 2);

    for (int i = 0; i < 5; i++){
        pthread_create(tid+i, NULL, cars, NULL);
    }

    for (int i = 0; i < 5; i++){
        pthread_join(tid[i], NULL);
    }

    sem_destroy(&road);

    return 0;
}