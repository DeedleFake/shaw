// Code generated for linux/amd64 by 'ccgo shaw.c', DO NOT EDIT.

//go:build linux && amd64

package main

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var (
	_ reflect.Type
	_ unsafe.Pointer
)

const APOS = "'"
const ARG_MAX = 131072
const BC_BASE_MAX = 99
const BC_DIM_MAX = 2048
const BC_SCALE_MAX = 99
const BC_STRING_MAX = 1000
const BUFSIZ = 1024
const CHARCLASS_NAME_MAX = 14
const CHAR_BIT = 8
const CHAR_MAX = 255
const CHAR_MIN = 0
const COLL_WEIGHTS_MAX = 2
const DELAYTIMER_MAX = 0x7fffffff
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXPR_NEST_MAX = 32
const FILENAME_MAX = 4096
const FILESIZEBITS = 64
const FOPEN_MAX = 1000
const HOST_NAME_MAX = 255
const INT_MAX = 2147483647
const IOV_MAX = 1024
const LINE_MAX = 4096
const LLONG_MAX = 0x7fffffffffffffff
const LOGIN_NAME_MAX = 256
const LONG_BIT = 64
const LONG_MAX = "__LONG_MAX"
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MB_LEN_MAX = 4
const MQ_PRIO_MAX = 32768
const NAME_MAX = 255
const NGROUPS_MAX = 32
const NL_ARGMAX = 9
const NL_LANGMAX = 32
const NL_MSGMAX = 32767
const NL_NMAX = 16
const NL_SETMAX = 255
const NL_TEXTMAX = 2048
const NZERO = 20
const PAGESIZE = 4096
const PAGE_SIZE = "PAGESIZE"
const PATH_MAX = 4096
const PIPE_BUF = 4096
const PTHREAD_DESTRUCTOR_ITERATIONS = 4
const PTHREAD_KEYS_MAX = 128
const PTHREAD_STACK_MIN = 2048
const P_tmpdir = "/tmp"
const RAND_MAX = 0x7fffffff
const RE_DUP_MAX = 255
const SCHAR_MAX = 127
const SEM_NSEMS_MAX = 256
const SEM_VALUE_MAX = 0x7fffffff
const SHRT_MAX = 0x7fff
const SSIZE_MAX = "LONG_MAX"
const SYMLOOP_MAX = 40
const TMP_MAX = 10000
const TTY_NAME_MAX = 32
const TZNAME_MAX = 6
const UCHAR_MAX = 255
const UINT_MAX = 0xffffffff
const USHRT_MAX = 0xffff
const WNOHANG = 1
const WORD_BIT = 32
const WUNTRACED = 2
const _GNU_SOURCE = 1
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const _LP64 = 1
const _POSIX2_BC_BASE_MAX = 99
const _POSIX2_BC_DIM_MAX = 2048
const _POSIX2_BC_SCALE_MAX = 99
const _POSIX2_BC_STRING_MAX = 1000
const _POSIX2_CHARCLASS_NAME_MAX = 14
const _POSIX2_COLL_WEIGHTS_MAX = 2
const _POSIX2_EXPR_NEST_MAX = 32
const _POSIX2_LINE_MAX = 2048
const _POSIX2_RE_DUP_MAX = 255
const _POSIX_AIO_LISTIO_MAX = 2
const _POSIX_AIO_MAX = 1
const _POSIX_ARG_MAX = 4096
const _POSIX_CHILD_MAX = 25
const _POSIX_CLOCKRES_MIN = 20000000
const _POSIX_DELAYTIMER_MAX = 32
const _POSIX_HOST_NAME_MAX = 255
const _POSIX_LINK_MAX = 8
const _POSIX_LOGIN_NAME_MAX = 9
const _POSIX_MAX_CANON = 255
const _POSIX_MAX_INPUT = 255
const _POSIX_MQ_OPEN_MAX = 8
const _POSIX_MQ_PRIO_MAX = 32
const _POSIX_NAME_MAX = 14
const _POSIX_NGROUPS_MAX = 8
const _POSIX_OPEN_MAX = 20
const _POSIX_PATH_MAX = 256
const _POSIX_PIPE_BUF = 512
const _POSIX_RE_DUP_MAX = 255
const _POSIX_RTSIG_MAX = 8
const _POSIX_SEM_NSEMS_MAX = 256
const _POSIX_SEM_VALUE_MAX = 32767
const _POSIX_SIGQUEUE_MAX = 32
const _POSIX_SSIZE_MAX = 32767
const _POSIX_SS_REPL_MAX = 4
const _POSIX_STREAM_MAX = 8
const _POSIX_SYMLINK_MAX = 255
const _POSIX_SYMLOOP_MAX = 8
const _POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const _POSIX_THREAD_KEYS_MAX = 128
const _POSIX_THREAD_THREADS_MAX = 64
const _POSIX_TIMER_MAX = 32
const _POSIX_TRACE_EVENT_NAME_MAX = 30
const _POSIX_TRACE_NAME_MAX = 8
const _POSIX_TRACE_SYS_MAX = 8
const _POSIX_TRACE_USER_EVENT_MAX = 32
const _POSIX_TTY_NAME_MAX = 9
const _POSIX_TZNAME_MAX = 6
const _STDC_PREDEF_H = 1
const _XOPEN_IOV_MAX = 16
const _XOPEN_NAME_MAX = 255
const _XOPEN_PATH_MAX = 1024
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BITINT_MAXWIDTH__ = 65535
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CHAR_BIT__ = 8
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 1
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 14
const __GXX_ABI_VERSION = 1019
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LITTLE_ENDIAN = 1234
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VERSION__ = "14.2.1 20240805"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const alloca1 = "__builtin_alloca"
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type size_t = uint64

type ssize_t = int64

type off_t = int64

type va_list = uintptr

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type wchar_t = int32

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type locale_t = uintptr

func load_custom(tls *libc.TLS, file uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c, line, space, v12, v2, v6 int32
	var fp, prev, wptr, v1, v14, v15, v3, v7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, fp, line, prev, space, wptr, v1, v12, v14, v15, v2, v3, v6, v7
	prev = __ccgo_ts
	line = int32(1)
	fp = libc.Xfopen(tls, file, __ccgo_ts+1)
	if !(fp != 0) {
		libc.Xperror(tls, file)
		libc.Xexit(tls, int32(1))
	}
	dend = dict
	v1 = dend
	dend++
	v2 = libc.Int32FromInt32(0)
	space = v2
	*(*uint8)(unsafe.Pointer(v1)) = uint8(v2)
	wptr = dend
	v3 = libc.UintptrFromInt32(0)
	suffix = v3
	dict = v3
_5:
	;
	v6 = libc.Xfgetc(tls, fp)
	c = v6
	if !(v6 != -int32(1)) {
		goto _4
	}
	if c == int32('\r') {
		goto _5
	}
	if c == int32('\n') {
		v7 = dend
		dend++
		*(*uint8)(unsafe.Pointer(v7)) = uint8(0)
		prev = wptr
		wptr = dend
		goto nxt
	}
	if !(c == int32(0xf0) && libc.Xfgetc(tls, fp) == int32(0x90) && libc.Xfgetc(tls, fp) == int32(0x91)) {
		goto _8
	}
	c = libc.Xfgetc(tls, fp)
	goto _9
_8:
	;
	if !(c == int32(' ')) {
		goto _10
	}
	v12 = space
	space++
	if v12 != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+3, libc.VaList(bp+8, file, line))
		goto del
	}
	goto _11
_10:
	;
	if !(!(libc.BoolInt32(uint32(c)|libc.Uint32FromInt32(32)-libc.Uint32FromUint8('a') < libc.Uint32FromInt32(26)) != 0) && !(libc.Xindex(tls, __ccgo_ts+40, c) != 0)) {
		goto _13
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+47, libc.VaList(bp+8, file, c, line))
	goto del
del:
	;
	for libc.Xfgetc(tls, fp) != int32('\n') {
	}
	dend = wptr
	goto nxt
nxt:
	;
	space = 0
	line++
	goto _5
_13:
	;
_11:
	;
_9:
	;
	if !(suffix != 0) && !(*(*uint8)(unsafe.Pointer(dend + uintptr(-libc.Int32FromInt32(1)))) != 0) && c != int32('^') {
		suffix = dend - uintptr(1)
	}
	if !(dict != 0) && !(*(*uint8)(unsafe.Pointer(dend + uintptr(-libc.Int32FromInt32(1)))) != 0) && c != int32('^') && c != int32('$') {
		dict = dend - uintptr(1)
	}
	v14 = dend
	dend++
	*(*uint8)(unsafe.Pointer(v14)) = uint8(c)
	if c == int32(' ') && !(wordcmp(tls, wptr, prev, 0, 0) != 0) {
		dend = wptr - uintptr(1)
		v15 = dend
		dend++
		*(*uint8)(unsafe.Pointer(v15)) = uint8('@') // heteronym separator
		wptr = prev
	}
	goto _5
_4:
	;
	libc.Xfclose(tls, fp)
	dend--
}

func wordcmp(tls *libc.TLS, one uintptr, two uintptr, ic int32, len1 int32) (r int32) {
	var a, b int8
	var v1 int32
	var v2, v3 uintptr
	_, _, _, _, _ = a, b, v1, v2, v3
	if len1 == 0 {
		len1 = int32(INT_MAX)
	}
	ic *= int32(32) // ic = "ignore case"
	for {
		v1 = len1
		len1--
		if !(v1 != 0) {
			break
		}
		v2 = one
		one++
		a = *(*int8)(unsafe.Pointer(v2))
		v3 = two
		two++
		b = *(*int8)(unsafe.Pointer(v3))
		if int32(a) == int32('_') {
			a = int8(' ')
		} // ignore POS tags
		if int32(b) == int32('_') {
			b = int8(' ')
		}
		a = int8(int32(a) | ic)
		b = int8(int32(b) | ic)
		if int32(a)|int32(b) == int32(32) {
			return 0
		} // both hit space or null
		if int32(a) < int32(b) {
			return -int32(1)
		}
		if int32(a) > int32(b) {
			return int32(1)
		}
	}
	return 0
}

func parse_root(tls *libc.TLS, in uintptr, out uintptr) {
	var comp int32
	var cp, high, low, mid, v2, v4 uintptr
	var v1 uint8
	_, _, _, _, _, _, _, _ = comp, cp, high, low, mid, v1, v2, v4
	v1 = libc.Uint8FromInt32(0)
	*(*uint8)(unsafe.Pointer(out + 1)) = v1
	*(*uint8)(unsafe.Pointer(out)) = v1
	low = dict
	high = dend
	for cond := true; cond; cond = comp != 0 {
		mid = low + uintptr((int64(high)-int64(low))/int64(2))
		for *(*int8)(unsafe.Pointer(mid + uintptr(-libc.Int32FromInt32(1)))) != 0 {
			mid--
		}
		if mid == low {
			for {
				v2 = mid
				mid++
				if !(*(*int8)(unsafe.Pointer(v2)) != 0) {
					break
				}
			}
		}
		if mid >= high {
			return
		}
		comp = wordcmp(tls, in, mid, int32(1), 0)
		if comp < 0 {
			high = mid
		}
		if comp > 0 {
			low = mid
		}
	}
	for int32(1) != 0 {
		low = mid - uintptr(1) // look for earlier matches
		for *(*int8)(unsafe.Pointer(low + uintptr(-libc.Int32FromInt32(1)))) != 0 {
			low--
		}
		if wordcmp(tls, in, low, int32(1), 0) != 0 {
			break
		}
		mid = low
	}
	cp = mid
	for {
		if wordcmp(tls, in, cp, int32(1), 0) != 0 {
			break
		}
		if !(wordcmp(tls, in, cp, 0, 0) != 0) {
			goto exact
		}
		for {
			cp++
			v4 = cp
			if !(*(*int8)(unsafe.Pointer(v4)) != 0) {
				break
			}
		}
		goto _3
	_3:
		;
		cp++
	}
	cp = mid
	goto exact
exact:
	;
	if libc.BoolInt32(uint32(*(*int8)(unsafe.Pointer(cp)))-uint32('A') < uint32(26)) != 0 {
		*(*uint8)(unsafe.Pointer(out)) = uint8('*')
	}
	for {
		if !(int32(*(*int8)(unsafe.Pointer(cp))) != int32(' ')) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(cp))) == int32('_') && !(aflag != 0) {
			*(*uint8)(unsafe.Pointer(out)) = uint8(',')
		}
		goto _5
	_5:
		;
		cp++
	}
	libc.Xstrcat(tls, out, cp+uintptr(1))
	if int32(aflag)&int32(1) != 0 && int32(*(*int8)(unsafe.Pointer(cp + 1))) == int32('.') || int32(aflag)&int32(2) != 0 && int32(*(*int8)(unsafe.Pointer(cp + uintptr(libc.Xstrlen(tls, cp)-uint64(1))))) == int32('.') {
		*(*uint8)(unsafe.Pointer(out)) = uint8(0)
	}
}

func parse_suffix(tls *libc.TLS, in uintptr, out uintptr, adj int32) (r int32) {
	bp := tls.Alloc(512)
	defer tls.Free(512)
	var best, i, len1, pass, score, split, v10, v11, v17, v8 int32
	var cp, suff, tail, v12, v14, v16, v4, v5, p13, p18, p2 uintptr
	var v7 uint8
	var _ /* try at bp+256 */ [256]uint8
	var _ /* word at bp+0 */ [256]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = best, cp, i, len1, pass, score, split, suff, tail, v10, v11, v12, v14, v16, v17, v4, v5, v7, v8, p13, p18, p2
	score = 0
	best = 0
	parse_root(tls, in, out)
	cp = in
	for {
		if !(*(*uint8)(unsafe.Pointer(cp)) != 0) {
			break
		}
		p2 = cp
		*(*uint8)(unsafe.Pointer(p2)) = uint8(int32(*(*uint8)(unsafe.Pointer(p2))) | libc.Int32FromInt32(32))
		goto _1
	_1:
		;
		cp++
	}
	len1 = int32(int64(cp) - int64(in))
	if *(*uint8)(unsafe.Pointer(out)) != 0 {
		return (len1 + adj) * (len1 + adj)
	}
	suff = suffix + uintptr(1)
	for {
		v4 = suff
		suff++
		if !(int32(*(*uint8)(unsafe.Pointer(v4))) == int32('$')) {
			break
		}
		split = int32(int64(len1) - (int64(libc.Xstrchr(tls, suff, int32(' '))) - int64(suff)))
		if split < int32(2) || wordcmp(tls, in+uintptr(split), suff, int32(1), 0) != 0 {
			goto _3
		}
		if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) {
			if len1-split == int32(1) || libc.Xindex(tls, __ccgo_ts+95, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0 {
				goto _3
			}
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+99) != 0) && !(libc.Xindex(tls, __ccgo_ts+102, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+110) != 0) && libc.Xindex(tls, __ccgo_ts+113, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+120) != 0) && libc.Xindex(tls, __ccgo_ts+123, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+127) != 0) && libc.Xindex(tls, __ccgo_ts+130, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+132) != 0) && libc.Xindex(tls, __ccgo_ts+134, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+143) != 0) && libc.Xindex(tls, __ccgo_ts+145, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+1) != 0) && libc.Xindex(tls, __ccgo_ts+152, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+158) != 0) && libc.Xindex(tls, __ccgo_ts+160, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+164) != 0) && libc.Xindex(tls, __ccgo_ts+166, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+168) != 0) && libc.Xindex(tls, __ccgo_ts+170, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2))))) != int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+174) != 0) && libc.Xindex(tls, __ccgo_ts+178, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2))))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) {
			goto _3
		}
		for {
			v5 = suff
			suff++
			if !(int32(*(*uint8)(unsafe.Pointer(v5))) != int32(' ')) {
				break
			}
		}
		pass = 0
		for {
			if !(pass < int32(2)) {
				break
			}
			libc.Xstrcpy(tls, bp, in)
			v7 = libc.Uint8FromInt32(0)
			(*(*[256]uint8)(unsafe.Pointer(bp)))[split+int32(1)] = v7
			(*(*[256]uint8)(unsafe.Pointer(bp)))[split] = v7
			if !(pass != 0) {
				if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) == int32('i') && !(libc.Xindex(tls, __ccgo_ts+181, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0) {
					(*(*[256]uint8)(unsafe.Pointer(bp)))[split-int32(1)] = uint8('y')
				} else {
					if libc.Xindex(tls, __ccgo_ts+190, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0 && !(libc.Xindex(tls, __ccgo_ts+197, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) {
						if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) == int32('u') && (int32(*(*uint8)(unsafe.Pointer(in + uintptr(split+int32(1))))) == int32('b') || int32(*(*uint8)(unsafe.Pointer(in + uintptr(split+int32(1))))) == int32('p')) {
							goto _6
						}
						if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2))))) && !(libc.Xindex(tls, __ccgo_ts+202, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) {
							(*(*[256]uint8)(unsafe.Pointer(bp)))[split-int32(1)] = uint8(0)
						} else {
							if (libc.Xindex(tls, __ccgo_ts+206, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 || int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) == int32('e') || libc.Xindex(tls, __ccgo_ts+215, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2)))))) != 0) && (!(libc.Xindex(tls, __ccgo_ts+223, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) || !(libc.Xindex(tls, __ccgo_ts+226, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0)) {
								(*(*[256]uint8)(unsafe.Pointer(bp)))[split] = uint8('e')
							} else {
								goto _6
							}
						}
					} else {
						if !(libc.Xstrcmp(tls, bp+uintptr(split)-uintptr(2), __ccgo_ts+230) != 0) {
							(*(*[256]uint8)(unsafe.Pointer(bp)))[split] = uint8('e')
						} else {
							goto _6
						}
					}
				}
			}
			aflag = uint8(int32(aflag) & ^libc.Int32FromInt32(2))
			if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) != int32('\'') {
				aflag = uint8(int32(aflag) | libc.Int32FromInt32(2))
			}
			v8 = parse_suffix(tls, bp, bp+256, int32(uint64(split)-libc.Xstrlen(tls, bp)))
			score = v8
			if v8 != 0 {
				score += (len1 - split + adj) * (len1 - split + adj)
				if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+233) != 0) {
					score = int32(1)
				}
				i = 0
				for {
					if !(uint64(i) < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(8)) {
						break
					}
					if !(libc.Xstrcmp(tls, in+uintptr(split), uintptr(unsafe.Pointer(&penal))+uintptr(i)*8) != 0) {
						score -= int32(9)
					}
					goto _9
				_9:
					;
					i++
				}
				if score < int32(1) {
					score = int32(1)
				}
			}
			if score <= best {
				goto _6
			}
			best = score
			libc.Xstrcpy(tls, out, bp+256)
			i = int32(libc.Xstrlen(tls, out))
			if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) == int32('e') && !(libc.Xindex(tls, __ccgo_ts+238, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2)))))) != 0) && libc.Xindex(tls, __ccgo_ts+226, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0 && libc.Xindex(tls, __ccgo_ts+245, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split+int32(1)))))) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+255) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+259) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+263) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+267) != 0 {
				if libc.Xindex(tls, __ccgo_ts+270, int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1)))))) != 0 {
					i--
				}
				v10 = i
				i++
				*(*uint8)(unsafe.Pointer(out + uintptr(v10))) = uint8(0xa6)
				*(*uint8)(unsafe.Pointer(out + uintptr(i))) = uint8(0)
			}
			cp = suff
			tail = bp + uintptr(libc.Xstrlen(tls, bp))
			if int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(*(*uint8)(unsafe.Pointer(cp))) && libc.Xindex(tls, __ccgo_ts+273, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0 && libc.Xstrlen(tls, cp) < uint64(3) {
				*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1)))) = uint8(0)
			} // with -n, -l, -ly, one is enough
			if int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(0xa6) && libc.Xindex(tls, __ccgo_ts+276, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0 { // link ia and ear
				v12 = cp
				cp++
				if int32(*(*uint8)(unsafe.Pointer(v12))) == int32(0xbc) {
					v11 = int32(0xbd)
				} else {
					v11 = int32(0xbe)
				}
				*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1)))) = uint8(v11)
			}
			if libc.Xindex(tls, __ccgo_ts+279, int32(*(*uint8)(unsafe.Pointer(tail + uintptr(-libc.Int32FromInt32(1)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(0x93) && int32(*(*uint8)(unsafe.Pointer(cp))) > int32(0x97) && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+282) != 0 {
				p13 = out + uintptr(i-int32(1))
				*(*uint8)(unsafe.Pointer(p13)) = uint8(int32(*(*uint8)(unsafe.Pointer(p13))) + libc.Int32FromInt32(10))
			} // change f to v in Slavic surnames
			if !(libc.Xstrcmp(tls, tail-uintptr(2), __ccgo_ts+284) != 0) && !(libc.Xstrcmp(tls, cp, __ccgo_ts+287) != 0) && !(libc.Xstrcmp(tls, out+uintptr(i)-uintptr(2), __ccgo_ts+289) != 0) {
				libc.Xstrcpy(tls, out+uintptr(i)-uintptr(2), __ccgo_ts+292)
			} // drop the schwa when -le becomes -ly
			if int32(*(*uint8)(unsafe.Pointer(cp + 1))) != int32(0x99) && !(libc.Xindex(tls, __ccgo_ts+294, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0) && !(libc.Xstrcmp(tls, out+uintptr(i)-uintptr(3), __ccgo_ts+298) != 0) {
				libc.Xstrcpy(tls, out+uintptr(i)-uintptr(3), __ccgo_ts+302)
			} // drop the schwa in -sm words
			if libc.Xindex(tls, __ccgo_ts+305, int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(2)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(0xa4) && !(libc.Xstrcmp(tls, cp, __ccgo_ts+309) != 0) {
				if libc.Xindex(tls, __ccgo_ts+313, int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(3)))))) != 0 || int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(2))))) == int32(0xbe) {
					v14 = __ccgo_ts + 287
				} else {
					v14 = __ccgo_ts
				}
				libc.Xstrcpy(tls, out+uintptr(i)-uintptr(2), v14)
				libc.Xstrcat(tls, out, __ccgo_ts+318) // strange "-(i)ality" rule
			}
			libc.Xstrcat(tls, out, cp)
			goto _6
		_6:
			;
			pass++
		}
		goto _3
	_3:
		;
		suff += uintptr(libc.Xstrlen(tls, suff) + uint64(1))
	}
	if int32(*(*uint8)(unsafe.Pointer(in + uintptr(len1-int32(1))))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(len1-int32(2))))) && !(libc.Xindex(tls, __ccgo_ts+321, int32(*(*uint8)(unsafe.Pointer(in + uintptr(len1-int32(2)))))) != 0) {
		aflag = uint8(int32(aflag) | libc.Int32FromInt32(2))
		libc.Xstrcpy(tls, bp, in)
		(*(*[256]uint8)(unsafe.Pointer(bp)))[len1-int32(1)] = uint8(0) // drop final double consonant
		score = parse_suffix(tls, bp, bp+256, 0)
		if best < score {
			best = score
			libc.Xstrcpy(tls, out, bp+256)
		}
	}
	cp = out + uintptr(libc.Xstrlen(tls, out))
	if int64(cp)-int64(out) < int64(2) {
		return best
	}
	if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1))))) == int32(':') && int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) > int32(0x99) {
		i = -int32(3)
		for {
			if !(cp+uintptr(i) > out && int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) < int32(0x90)) {
				break
			}
			goto _15
		_15:
			;
			i--
		}
		if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) == int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) || int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) == int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2)))))-int32(10) || int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) == int32(0x9f) && libc.Xindex(tls, __ccgo_ts+313, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i))))) != 0 {
			cp++
			v16 = cp
			*(*uint8)(unsafe.Pointer(v16)) = uint8(0)
			*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1)))) = uint8(':')
			*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2)))) = *(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3))))
			v17 = -libc.Int32FromInt32(3)
			i = v17
			*(*uint8)(unsafe.Pointer(cp + uintptr(v17))) = uint8(0xa9)
		}
		if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) < int32(0x98) {
			p18 = cp + uintptr(-libc.Int32FromInt32(2))
			*(*uint8)(unsafe.Pointer(p18)) = uint8(int32(*(*uint8)(unsafe.Pointer(p18))) - libc.Int32FromInt32(10))
		}
	}
	if int64(cp)-int64(out) < int64(4) {
		return best
	}
	if !(libc.Xstrcmp(tls, cp-uintptr(4), __ccgo_ts+328) != 0) && libc.Xindex(tls, __ccgo_ts+333, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(5)))))) != 0 {
		libc.Xstrcpy(tls, cp-uintptr(4), __ccgo_ts+336)
	} // "ically" is pronounced "icly"
	if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1))))) == int32(0xbe) {
		cp++
	} // ia palatization rules
	if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) == int32(0xbe) && libc.Xindex(tls, __ccgo_ts+340, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3)))))) != 0 && libc.Xindex(tls, __ccgo_ts+344, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1)))))) != 0 {
		i = int32(0x96)
		if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3))))) == int32(0x91) {
			if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(4))))) == int32(0x95) {
				i = int32(0x97)
			}
		} else {
			if libc.Xindex(tls, __ccgo_ts+348, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1)))))) != 0 && libc.Xindex(tls, __ccgo_ts+350, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(4)))))) != 0 {
				i = int32(0xa0)
			}
		}
		*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3)))) = uint8(i)
		*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2)))) = uint8(0xa9)
	}
	return best
}

var penal = [13][8]int8{
	0:  {'b', 'e', 'd'},
	1:  {'c', 'a', 'n'},
	2:  {'c', 'e', 'n', 't'},
	3:  {'d', 'a', 'n', 'c', 'e'},
	4:  {'d', 'e', 'n'},
	5:  {'i', 'n', 'e'},
	6:  {'k', 'i', 'n'},
	7:  {'o', 'n', 'e'},
	8:  {'p', 'a', 'l'},
	9:  {'p', 'a', 't', 'h'},
	10: {'s', 't', 'e', 'r'},
	11: {'w', 'i', 'n', 'g'},
	12: {'x'},
}

func parse_prefix(tls *libc.TLS, in uintptr, out uintptr, ms int32) (r int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var best, i, len1, score, split, v4, v5 int32
	var pref, v2, v3 uintptr
	var _ /* try at bp+0 */ [256]uint8
	_, _, _, _, _, _, _, _, _, _ = best, i, len1, pref, score, split, v2, v3, v4, v5
	best = parse_suffix(tls, in, out, 0)
	len1 = int32(libc.Xstrlen(tls, in))
	if best == len1*len1 {
		return best
	}
	pref = uintptr(unsafe.Pointer(&prefix)) + uintptr(1)
	for {
		v2 = pref
		pref++
		if !(int32(*(*int8)(unsafe.Pointer(v2))) == int32('^')) {
			break
		}
		split = int32(int64(libc.Xstrchr(tls, pref, int32(' '))) - int64(pref))
		if split <= ms || split > len1-int32(2) || wordcmp(tls, in, pref, int32(1), split) != 0 {
			goto _1
		}
		if int32(*(*int8)(unsafe.Pointer(in + uintptr(split-int32(1)))))|int32(32) == int32('u') && int32(*(*int8)(unsafe.Pointer(in + uintptr(split))))|int32(32) == int32('n') {
			goto _1
		}
		for {
			v3 = pref
			pref++
			if !(int32(*(*int8)(unsafe.Pointer(v3))) != int32(' ')) {
				break
			}
		}
		libc.Xstrcpy(tls, bp, pref)
		i = int32(libc.Xstrlen(tls, bp))
		aflag = libc.BoolUint8(int32(*(*int8)(unsafe.Pointer(in + uintptr(split-int32(1))))) != int32('\''))
		v4 = parse_prefix(tls, in+uintptr(split), bp+uintptr(i), int32(1))
		score = v4
		if v4 != 0 {
			score += split * split
		}
		if split == int32(2) && !(libc.Xstrncmp(tls, in, __ccgo_ts+358, uint64(2)) != 0) {
			score -= int32(4)
		}
		if score <= best {
			goto _1
		}
		best = score
		if int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i-int32(1)]) == int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i]) && int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i-int32(2)]) == int32(0xa6) && libc.Xindex(tls, __ccgo_ts+361, int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i])) != 0 || !(libc.Xstrncmp(tls, bp, __ccgo_ts+366, uint64(4)) != 0) || !(libc.Xstrncmp(tls, bp, __ccgo_ts+371, uint64(5)) != 0) {
			for {
				(*(*[256]uint8)(unsafe.Pointer(bp)))[i-int32(1)] = (*(*[256]uint8)(unsafe.Pointer(bp)))[i]
				goto _6
			_6:
				;
				v5 = i
				i++
				if !((*(*[256]uint8)(unsafe.Pointer(bp)))[v5] != 0) {
					break
				}
			}
		}
		libc.Xstrcpy(tls, out, bp)
		goto _1
	_1:
		;
		pref += uintptr(libc.Xstrlen(tls, pref) + uint64(1))
	}
	return best
}

func shaw_print(tls *libc.TLS, cp uintptr, cap1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var v1 uintptr
	_ = v1
	if (cap1 != 0 || libc.Xindex(tls, cp, int32('*')) != 0) && !(libc.Xindex(tls, cp, int32(',')) != 0) {
		libc.Xprintf(tls, __ccgo_ts+377, 0)
	}
	for {
		if !(libc.Xindex(tls, __ccgo_ts+380, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0) {
			if int32(*(*uint8)(unsafe.Pointer(cp))) == int32('\'') {
				libc.Xprintf(tls, __ccgo_ts+385, libc.VaList(bp+8, __ccgo_ts+388))
			} else {
				if int32(*(*uint8)(unsafe.Pointer(cp))) < int32(0x90) {
					libc.Xputchar(tls, int32(*(*uint8)(unsafe.Pointer(cp))))
				} else {
					libc.Xprintf(tls, __ccgo_ts+390, libc.VaList(bp+8, int32(0xf0), int32(0x90), int32(0x91), int32(*(*uint8)(unsafe.Pointer(cp)))))
				}
			}
		}
		goto _2
	_2:
		;
		cp++
		v1 = cp
		if !(*(*uint8)(unsafe.Pointer(v1)) != 0) {
			break
		}
	}
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(1152)
	defer tls.Free(1152)
	var c, cap1, script, sent, w, v1, v2, v7 uint8
	var fp uintptr
	var i, v8 int32
	var _ /* befto at bp+1024 */ [5][2][10]int8
	var _ /* copy at bp+256 */ [256]int8
	var _ /* out at bp+512 */ [256]int8
	var _ /* pword at bp+768 */ [256]int8
	var _ /* word at bp+0 */ [256]int8
	_, _, _, _, _, _, _, _, _, _, _ = c, cap1, fp, i, script, sent, w, v1, v2, v7, v8
	fp = libc.Xstdin
	w = uint8(0)
	sent = uint8(0)
	script = uint8(0)
	*(*[5][2][10]int8)(unsafe.Pointer(bp + 1024)) = [5][2][10]int8{
		0: {
			0: {'h', 'a', 's'},
			1: {-16, -112, -111, -107},
		},
		1: {
			0: {'h', 'a', 'v', 'e'},
			1: {-16, -112, -111, -109},
		},
		2: {
			0: {'u', 's', 'e', 'd'},
			1: {-16, -112, -111, -107, -16, -112, -111, -111},
		},
		3: {
			0: {'u', 'n', 'u', 's', 'e', 'd'},
			1: {-16, -112, -111, -107, -16, -112, -111, -111},
		},
		4: {
			0: {'s', 'u', 'p', 'p', 'o', 's', 'e', 'd'},
			1: {-16, -112, -111, -107, -16, -112, -111, -111},
		},
	}
	if argc < int32(2) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+399, libc.VaList(bp+1136, *(*uintptr)(unsafe.Pointer(argv))))
		libc.Xexit(tls, int32(1))
	}
	load_custom(tls, *(*uintptr)(unsafe.Pointer(argv + 1*8)))
	for !(libc.Xfeof(tls, fp) != 0) {
		c = uint8(libc.Xfgetc(tls, fp))
		if libc.BoolInt32(uint32(c)|uint32(32)-uint32('a') < uint32(26)) != 0 || w != 0 && int32(c) == int32('\'') {
			if !(w != 0) {
				v1 = sent
				sent++
				cap1 = libc.BoolUint8(v1 != 0 && libc.BoolInt32(uint32(c)-uint32('A') < uint32(26)) != 0)
			}
			v2 = w
			w++
			(*(*[256]int8)(unsafe.Pointer(bp)))[v2] = int8(c)
			continue
		}
		if libc.Xstrchr(tls, __ccgo_ts+426, int32(c)) != 0 {
			sent = uint8(0)
		}
		if !(w != 0) {
			goto _3
		}
		(*(*[256]int8)(unsafe.Pointer(bp)))[w] = 0
		if script != 0 {
			goto pass
		}
		if !(libc.Xstrcasecmp(tls, bp, __ccgo_ts+431) != 0) {
			i = 0
			for {
				if !(i < int32(5)) {
					break
				}
				if !(libc.Xstrcasecmp(tls, bp+768, bp+1024+uintptr(i)*20) != 0) {
					libc.Xprintf(tls, __ccgo_ts+434, libc.VaList(bp+1136, bp+1024+uintptr(i)*20+1*10))
				}
				goto _4
			_4:
				;
				i++
			}
		}
		aflag = uint8(0)
		libc.Xstrcpy(tls, bp+256, bp)
		if !(parse_prefix(tls, bp+256, bp+512, 0) != 0) {
			goto _5
		}
		shaw_print(tls, bp+512, int32(cap1))
		goto _6
	_5:
		;
		goto pass
	pass:
		;
		libc.Xprintf(tls, __ccgo_ts+385, libc.VaList(bp+1136, bp))
	_6:
		;
		libc.Xmemcpy(tls, bp+768, bp, uint64(256))
	_3:
		;
		goto end
	end:
		;
		w = uint8(0)
		if int32(c) != int32(0xff) {
			libc.Xputchar(tls, int32(c))
		}
		if int32(c) == int32('<') { // pass through HTML/XML tags and Javascript
			i = 0
			for cond := true; cond; cond = int32(c) != int32('>') && !(libc.Xfeof(tls, fp) != 0) {
				v7 = uint8(libc.Xfgetc(tls, fp))
				c = v7
				libc.Xputchar(tls, int32(v7))
				if i < int32(256) {
					v8 = i
					i++
					(*(*[256]int8)(unsafe.Pointer(bp)))[v8] = int8(c)
				}
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+439, uint64(3)) != 0) {
				sent = uint8(0)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+443, uint64(6)) != 0) {
				script = uint8(1)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+450, uint64(7)) != 0) {
				script = uint8(0)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+458, uint64(5)) != 0) {
				script = uint8(1)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+464, uint64(6)) != 0) {
				script = uint8(0)
			}
		}
	}
	return r
}

func main() {
	libc.Start(main1)
}

var aflag uint8

var dend uintptr

var dict = uintptr(unsafe.Pointer(&prefix))

/* Support signed or unsigned plain-char */

/* Implementation choices... */

/* Arbitrary numbers... */

/* POSIX/SUS requirements follow. These numbers come directly
 * from SUS and have nothing to do with the host system. */
var prefix [2200000]uint8

var suffix uintptr

var __ccgo_ts = uintptr(unsafe.Pointer(unsafe.StringData(__ccgo_ts1)))

var __ccgo_ts1 = "\x00r\x00%s: Extra space, line %d discarded.\n\x00^$':._\x00%s: Illegal character '%c', line %d discarded.\n\x00eos\x00es\x00hiosuxz\x00ry\x00aeiouf\x00ha\x00cst\x00th\x00e\x00t\x00aeioust'\x00k\x00aceino\x00aeiou\x00m\x00eis\x00z\x00i\x00n\x00eio\x00ess\x00ln\x00cfikmpsv\x00aeiouy\x00aeio\x00hsw\x00cghlsuvz\x00aeiousy\x00cg\x00aou\x00dg\x00call\x00aegiou\x00dlmnprstu\x00arm\x00out\x00und\x00up\x00\xa6\xb0\x00\xa4\xaf\x00\xa9\xbc\x00vw\x00s\x00le\x00\xa6\x00\xa9\xa4\x00\xa4\x00'\x9b\x9f\x00\x9f\xa9\xa5\x00\x9f\xa5\x00\xa9\xad\xbe\x00\xa6\x91\xa6\x00\x96\x97\xa0\xa1\x00\xa8\xa4\x00aeiosu\x00\x92\xa9\xa4\xa6\x00\xa6\xa9\x00\x92\xa4\xa6\x00\x91\x95\x9f\x00\x95\xa4\xaf\x00\xaf\x00\xb0\xb1\xb4\xb5\xb7\xbb\xbf\x00la\x00\xa4\xa5\xae\xaf\x00\xa5\xa9\x92\x92\x00\xa5\xa9\x92*\x92\x00·\x00:.,*\x00%s\x00'\x00%c%c%c%c\x00Usage: %s dictionary-file\n\x00.:?!\x00to\x00@%s \x00div\x00script\x00/script\x00style\x00/style\x00"
