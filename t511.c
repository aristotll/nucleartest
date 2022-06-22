#define _GNU_SOURCE
#include <sys/mount.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#define STACK_SIZE 4096
static char container_stack[STACK_SIZE];
char* const container_args[] = {
  "/bin/sh",
  NULL
};

int container_main(void* arg)
{
  printf("Container - inside the container!\n");
  execv(container_args[0], container_args);
  printf("Something's wrong!\n");
  return 1;
}

int main()
{
  printf("Parent - start a container!\n");
  int container_pid = clone(container_main, container_stack, CLONE_VM|SIGCHLD , NULL);
  int status;
  waitpid(container_pid, &status, 0);
  if (WIFEXITED(status) != 1) {
        printf("error!");
        printf("%d\n", WTERMSIG(status));
  }
  printf("Parent - container stopped!\n");
  return 0;
}
