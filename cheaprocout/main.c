#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/ptrace.h>
#include <sys/syscall.h>
#include <sys/user.h>
#include <unistd.h>

static pid_t pid;

void sig_handler(int signum) {
  long ret;

  printf("\nInside handler function\n");
  printf("try to detach from %d\n", pid);
  ret = ptrace(PTRACE_DETACH, pid, NULL, NULL);
  if (ret < 0) {
    perror("failed to detach");
    exit(1);
  }
  printf("detached from %d (ret: %ld)\n", pid, ret);
  exit(0);
}

int main(int argc, char *argv[]) {
  long ret;
  int status;
  struct user_regs_struct regs;

  if (argc < 2) {
    fprintf(stderr, "specify pid\n");
    exit(1);
  }
  pid = atoi(argv[1]);
  printf("attach to %d\n", pid);

  ret = ptrace(PTRACE_ATTACH, pid, NULL, NULL);
  if (ret < 0) {
    perror("failed to attach");
    exit(1);
  }
  printf("attached to %d (ret: %ld)\n", pid, ret);
  signal(SIGINT, sig_handler);
  ptrace(PTRACE_SETOPTIONS, pid, NULL, PTRACE_O_TRACESYSGOOD);

  int is_enter_stop = 0;
  long prev_orig_rax = -1;
  while (1) {
    waitpid(pid, &status, 0);
    if (WIFEXITED(status)) {
      break;
      // } else if (WIFSTOPPED(status) && WSTOPSIG(status) == (SIGTRAP | 0x80))
      // {
    } else if (WIFSTOPPED(status)) {
      ptrace(PTRACE_GETREGS, pid, NULL, &regs);
      is_enter_stop = prev_orig_rax == regs.orig_rax ? !is_enter_stop : 1;
      prev_orig_rax = regs.orig_rax;
      if (is_enter_stop && regs.orig_rax == SYS_write) {
        peek_and_output(pid, regs.rsi, regs.rdx, (int)regs.rdi);
      }
    }
    ptrace(PTRACE_SYSCALL, pid, NULL, NULL);
  }

  ret = ptrace(PTRACE_DETACH, pid, NULL, NULL);
  if (ret < 0) {
    perror("failed to detach");
    exit(1);
  }
  printf("detached from %d (ret: %ld)\n", pid, ret);

  return 0;
}

void peek_and_output(pid_t pid, long long addr, long long size, int fd) {
  if (fd != 1 && fd != 2) {
    return;
  }
  char *bytes = malloc(size + sizeof(long));
  int i;
  for (i = 0; i < size; i += sizeof(long)) {
    // ptrace fetch memory value as word (32bit)
    long data = ptrace(PTRACE_PEEKDATA, pid, addr + i, NULL);
    if (data == -1) {
      printf("failed to peek data\n");
      free(bytes);
      return;
    }
    memcpy(bytes + i, &data, sizeof(long));
  }
  bytes[size] = '\0';
  write(fd == 2 ? 2 : 1, bytes, size);
  fflush(fd == 2 ? stderr : stdout);
  free(bytes);
}
