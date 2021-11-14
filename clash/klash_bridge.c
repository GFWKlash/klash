#include "klash_bridge.h"

#include <stddef.h>
#include <stdint.h>
#include <malloc.h>

#include <pthread.h>

pthread_mutex_t info_lock;
pthread_mutex_t error_lock;
pthread_mutex_t warn_lock;
pthread_mutex_t debug_lock;
pthread_mutex_t verbose_lock;

void (*klash_log_info_func)(const char *payload);
void (*klash_log_error_func)(const char *payload);
void (*klash_log_warn_func)(const char *payload);
void (*klash_log_debug_func)(const char *payload);
void (*klash_log_verbose_func)(const char *payload);

void klash_log_info(const char *payload)
{
    pthread_mutex_lock(&info_lock);
    if (klash_log_info_func)
    {
        klash_log_info_func(payload);
    }
    pthread_mutex_unlock(&info_lock);
    free((void *)payload);
}

void klash_log_error(const char *payload)
{
    pthread_mutex_lock(&error_lock);
    if (klash_log_error_func)
    {
        klash_log_error_func(payload);
    }
    pthread_mutex_unlock(&error_lock);
    free((void *)payload);
}

void klash_log_warn(const char *payload)
{
    pthread_mutex_lock(&warn_lock);
    if (klash_log_warn_func)
    {
        klash_log_warn_func(payload);
    }
    pthread_mutex_unlock(&warn_lock);
    free((void *)payload);
}

void klash_log_debug(const char *payload)
{
    pthread_mutex_lock(&debug_lock);
    if (klash_log_debug_func)
    {
        klash_log_debug_func(payload);
    }
    pthread_mutex_unlock(&debug_lock);
    free((void *)payload);
}

void klash_log_verbose(const char *payload)
{
    pthread_mutex_lock(&verbose_lock);
    if (klash_log_verbose_func)
    {
        klash_log_verbose_func(payload);
    }
    pthread_mutex_unlock(&verbose_lock);
    free((void *)payload);
}

void klash_reset_log_info()
{
    pthread_mutex_lock(&info_lock);
    klash_log_info_func = 0;
    pthread_mutex_unlock(&info_lock);
}

void klash_reset_log_error()
{
    pthread_mutex_lock(&error_lock);
    klash_log_error_func = 0;
    pthread_mutex_unlock(&error_lock);
}

void klash_reset_log_warn()
{
    pthread_mutex_lock(&warn_lock);
    klash_log_warn_func = 0;
    pthread_mutex_unlock(&warn_lock);
}

void klash_reset_log_debug()
{
    pthread_mutex_lock(&debug_lock);
    klash_log_debug_func = 0;
    pthread_mutex_unlock(&debug_lock);
}

void klash_reset_log_verbose()
{
    pthread_mutex_lock(&verbose_lock);
    klash_log_verbose_func = 0;
    pthread_mutex_unlock(&verbose_lock);
}

void klash_set_log_info(void (*f)(const char *))
{
    pthread_mutex_lock(&info_lock);
    klash_log_info_func = f;
    pthread_mutex_unlock(&info_lock);
}

void klash_set_log_error(void (*f)(const char *))
{
    pthread_mutex_lock(&error_lock);
    klash_log_error_func = f;
    pthread_mutex_unlock(&error_lock);
}

void klash_set_log_warn(void (*f)(const char *))
{
    pthread_mutex_lock(&warn_lock);
    klash_log_warn_func = f;
    pthread_mutex_unlock(&warn_lock);
}

void klash_set_log_debug(void (*f)(const char *))
{
    pthread_mutex_lock(&debug_lock);
    klash_log_debug_func = f;
    pthread_mutex_unlock(&debug_lock);
}

void klash_set_log_verbose(void (*f)(const char *))
{
    pthread_mutex_lock(&verbose_lock);
    klash_log_verbose_func = f;
    pthread_mutex_unlock(&verbose_lock);
}
