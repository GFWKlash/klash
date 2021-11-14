#ifndef KLASH_BRIDGE_H
#define KLASH_BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stddef.h>
#include <stdint.h>
#include <malloc.h>

/* Wrappers and setters for Klash log handler */
extern void klash_log_info(const char *payload);
extern void klash_log_error(const char *payload);
extern void klash_log_warn(const char *payload);
extern void klash_log_debug(const char *payload);
extern void klash_log_verbose(const char *payload);

extern void klash_set_log_info(void (*f)(const char *));
extern void klash_set_log_error(void (*f)(const char *));
extern void klash_set_log_warn(void (*f)(const char *));
extern void klash_set_log_debug(void (*f)(const char *));
extern void klash_set_log_verbose(void (*f)(const char *));

extern void klash_reset_log_info();
extern void klash_reset_log_error();
extern void klash_reset_log_warn();
extern void klash_reset_log_debug();
extern void klash_reset_log_verbose();
/**********************************/

#ifdef __cplusplus
}
#endif

#endif
