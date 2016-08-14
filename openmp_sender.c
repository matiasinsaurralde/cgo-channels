// +build openmp

#include <omp.h>
#include <stdio.h>

#include "common.h"
#include "openmp_sender.h"

void doSend() {
  char str[64];
  sprintf(str, "Hello from thread %d, nthreads %d", omp_get_thread_num(), omp_get_num_threads());
  publishString(str);
}
void parallelSend() {
  #pragma omp parallel
  doSend();
}
