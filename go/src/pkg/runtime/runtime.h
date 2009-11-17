<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/runtime.h</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/runtime.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * basic types
 */
typedef	signed char		int8;
typedef	unsigned char		uint8;
typedef	signed short		int16;
typedef	unsigned short		uint16;
typedef	signed int		int32;
typedef	unsigned int		uint32;
typedef	signed long long int	int64;
typedef	unsigned long long int	uint64;
typedef	float			float32;
typedef	double			float64;

#ifdef _64BIT
typedef	uint64		uintptr;
#else
typedef	uint32		uintptr;
#endif

/*
 * get rid of C types
 * the / / / forces a syntax error immediately,
 * which will show &#34;last name: XXunsigned&#34;.
 */
#define	unsigned		XXunsigned / / /
#define	signed			XXsigned / / /
#define	char			XXchar / / /
#define	short			XXshort / / /
#define	int			XXint / / /
#define	long			XXlong / / /
#define	float			XXfloat / / /
#define	double			XXdouble / / /

/*
 * defined types
 */
typedef	uint8			bool;
typedef	uint8			byte;
typedef	struct	Alg		Alg;
typedef	struct	Func		Func;
typedef	struct	G		G;
typedef	struct	Gobuf		Gobuf;
typedef	struct	Lock		Lock;
typedef	struct	M		M;
typedef	struct	Mem		Mem;
typedef	union	Note		Note;
typedef	struct	Slice		Slice;
typedef	struct	Stktop		Stktop;
typedef	struct	String		String;
typedef	struct	Usema		Usema;
typedef	struct	SigTab		SigTab;
typedef	struct	MCache		MCache;
typedef	struct	Iface		Iface;
typedef	struct	Itab		Itab;
typedef	struct	Eface	Eface;
typedef	struct	Type		Type;
typedef	struct	Defer		Defer;
typedef	struct	hash		Hmap;
typedef	struct	Hchan		Hchan;

/*
 * per-cpu declaration.
 * &#34;extern register&#34; is a special storage class implemented by 6c, 8c, etc.
 * on machines with lots of registers, it allocates a register that will not be
 * used in generated code.  on the x86, it allocates a slot indexed by a
 * segment register.
 *
 * amd64: allocated downwards from R15
 * x86: allocated upwards from 0(FS)
 * arm: allocated upwards from R9
 *
 * every C file linked into a Go program must include runtime.h
 * so that the C compiler knows to avoid other uses of these registers.
 * the Go compilers know to avoid them.
 */
extern	register	G*	g;
extern	register	M*	m;

/*
 * defined constants
 */
enum
{
	// G status
	Gidle,
	Grunnable,
	Grunning,
	Gsyscall,
	Gwaiting,
	Gmoribund,
	Gdead,
};
enum
{
	true	= 1,
	false	= 0,
};

/*
 * structures
 */
struct	Lock
{
	uint32	key;
	uint32	sema;	// for OS X
};
struct	Usema
{
	uint32	u;
	uint32	k;
};
union	Note
{
	struct {	// Linux
		Lock	lock;
	};
	struct {	// OS X
		int32	wakeup;
		Usema	sema;
	};
};
struct String
{
	byte*	str;
	int32	len;
};
struct Iface
{
	Itab*	tab;
	void*	data;
};
struct Eface
{
	Type*	type;
	void*	data;
};

struct	Slice
{				// must not move anything
	byte*	array;		// actual data
	uint32	len;		// number of elements
	uint32	cap;		// allocated number of elements
};
struct	Gobuf
{
	// The offsets of these fields are known to (hard-coded in) libmach.
	byte*	sp;
	byte*	pc;
	G*	g;
};
struct	G
{
	byte*	stackguard;	// cannot move - also known to linker, libmach, libcgo
	byte*	stackbase;	// cannot move - also known to libmach, libcgo
	Defer*	defer;
	Gobuf	sched;		// cannot move - also known to libmach
	byte*	stack0;
	byte*	entry;		// initial function
	G*	alllink;	// on allg
	void*	param;		// passed parameter on wakeup
	int16	status;
	int32	goid;
	int32	selgen;		// valid sudog pointer
	G*	schedlink;
	bool	readyonstop;
	M*	m;		// for debuggers, but offset not hard-coded
	M*	lockedm;
	void	(*cgofn)(void*);	// for cgo/ffi
	void	*cgoarg;
};
struct	Mem
{
	uint8*	hunk;
	uint32	nhunk;
	uint64	nmmap;
	uint64	nmal;
};
struct	M
{
	// The offsets of these fields are known to (hard-coded in) libmach.
	G*	g0;		// goroutine with scheduling stack
	void	(*morepc)(void);
	void*	morefp;	// frame pointer for more stack
	Gobuf	morebuf;	// gobuf arg to morestack

	// Fields not known to debuggers.
	uint32	moreframe;	// size arguments to morestack
	uint32	moreargs;
	uintptr	cret;		// return value from C
	uint64	procid;		// for debuggers, but offset not hard-coded
	G*	gsignal;	// signal-handling G
	uint32	tls[8];		// thread-local storage (for 386 extern register)
	Gobuf	sched;	// scheduling stack
	G*	curg;		// current running goroutine
	int32	id;
	int32	mallocing;
	int32	gcing;
	int32	locks;
	int32	waitnextg;
	Note	havenextg;
	G*	nextg;
	M*	alllink;	// on allm
	M*	schedlink;
	Mem	mem;
	uint32	machport;	// Return address for Mach IPC (OS X)
	MCache	*mcache;
	G*	lockedg;
};
struct	Stktop
{
	// The offsets of these fields are known to (hard-coded in) libmach.
	uint8*	stackguard;
	uint8*	stackbase;
	Gobuf	gobuf;
	uint32	args;

	// Frame pointer: where args start in old frame.
	// fp == gobuf.sp except in the case of a reflected
	// function call, which uses an off-stack argument frame.
	uint8*	fp;
};
struct	Alg
{
	uintptr	(*hash)(uint32, void*);
	uint32	(*equal)(uint32, void*, void*);
	void	(*print)(uint32, void*);
	void	(*copy)(uint32, void*, void*);
};
struct	SigTab
{
	int32	flags;
	int8	*name;
};
enum
{
	SigCatch = 1&lt;&lt;0,
	SigIgnore = 1&lt;&lt;1,
	SigRestart = 1&lt;&lt;2,
};

// (will be) shared with go; edit ../cmd/6g/sys.go too.
// should move out of sys.go eventually.
// also eventually, the loaded symbol table should
// be closer to this form.
struct	Func
{
	String	name;
	String	type;	// go type string
	String	src;	// src file name
	uint64	entry;	// entry pc
	int64	frame;	// stack frame size
	Slice	pcln;	// pc/ln tab for this func
	int64	pc0;	// starting pc, ln for table
	int32	ln0;
	int32	args;	// number of 32-bit in/out args
	int32	locals;	// number of 32-bit locals
};

/*
 * defined macros
 *    you need super-goru privilege
 *    to add this list.
 */
#define	nelem(x)	(sizeof(x)/sizeof((x)[0]))
#define	nil		((void*)0)

/*
 * known to compiler
 */
enum
{
	AMEM,
	ANOEQ,
	ASTRING,
	AINTER,
	ANILINTER,
	AFAKE,
	Amax
};


enum {
	Structrnd = sizeof(uintptr)
};

/*
 * deferred subroutine calls
 */
struct Defer
{
	int32	siz;
	byte*	sp;
	byte*	fn;
	Defer*	link;
	byte	args[8];	// padded to actual size
};

/*
 * external data
 */
extern	Alg	algarray[Amax];
extern	String	emptystring;
G*	allg;
M*	allm;
int32	goidgen;
extern	int32	gomaxprocs;
extern	int32	panicking;
extern	int32	maxround;
extern	int32	fd;	// usually 1; set to 2 when panicking
int8*	goos;

/*
 * common functions and data
 */
int32	strcmp(byte*, byte*);
int32	findnull(byte*);
void	dump(byte*, int32);
int32	runetochar(byte*, int32);
int32	charntorune(int32*, uint8*, int32);

/*
 * very low level c-called
 */
void	gogo(Gobuf*, uintptr);
void	gogocall(Gobuf*, void(*)(void));
uintptr	gosave(Gobuf*);
void	runtime·lessstack(void);
void	goargs(void);
void	FLUSH(void*);
void*	getu(void);
void	throw(int8*);
uint32	rnd(uint32, uint32);
void	prints(int8*);
void	printf(int8*, ...);
byte*	mchr(byte*, byte, byte*);
void	mcpy(byte*, byte*, uint32);
int32	mcmp(byte*, byte*, uint32);
void	mmov(byte*, byte*, uint32);
void*	mal(uint32);
uint32	cmpstring(String, String);
String	gostring(byte*);
void	initsig(void);
int32	gotraceback(void);
void	traceback(uint8 *pc, uint8 *sp, G* gp);
void	tracebackothers(G*);
int32	open(byte*, int32, ...);
int32	write(int32, void*, int32);
bool	cas(uint32*, uint32, uint32);
void	jmpdefer(byte*, void*);
void	exit1(int32);
void	ready(G*);
byte*	getenv(int8*);
int32	atoi(byte*);
void	newosproc(M *m, G *g, void *stk, void (*fn)(void));
void	signalstack(byte*, int32);
G*	malg(int32);
void	minit(void);
Func*	findfunc(uintptr);
int32	funcline(Func*, uint64);
void*	stackalloc(uint32);
void	stackfree(void*);
MCache*	allocmcache(void);
void	mallocinit(void);
bool	ifaceeq(Iface, Iface);
bool	efaceeq(Eface, Eface);
uintptr	ifacehash(Iface);
uintptr	efacehash(Eface);
uintptr	nohash(uint32, void*);
uint32	noequal(uint32, void*, void*);
void*	malloc(uintptr size);
void*	mallocgc(uintptr size);
void	free(void *v);
void	exit(int32);
void	breakpoint(void);
void	gosched(void);
void	goexit(void);
void	runcgo(void (*fn)(void*), void*);

#pragma	varargck	argpos	printf	1

#pragma	varargck	type	&#34;d&#34;	int32
#pragma	varargck	type	&#34;d&#34;	uint32
#pragma	varargck	type	&#34;D&#34;	int64
#pragma	varargck	type	&#34;D&#34;	uint64
#pragma	varargck	type	&#34;x&#34;	int32
#pragma	varargck	type	&#34;x&#34;	uint32
#pragma	varargck	type	&#34;X&#34;	int64
#pragma	varargck	type	&#34;X&#34;	uint64
#pragma	varargck	type	&#34;p&#34;	void*
#pragma	varargck	type	&#34;p&#34;	uintptr
#pragma	varargck	type	&#34;s&#34;	int8*
#pragma	varargck	type	&#34;s&#34;	uint8*
#pragma	varargck	type	&#34;S&#34;	String

// TODO(rsc): Remove. These are only temporary,
// for the mark and sweep collector.
void	stoptheworld(void);
void	starttheworld(void);

/*
 * mutual exclusion locks.  in the uncontended case,
 * as fast as spin locks (just a few user-level instructions),
 * but on the contention path they sleep in the kernel.
 * a zeroed Lock is unlocked (no need to initialize each lock).
 */
void	lock(Lock*);
void	unlock(Lock*);

/*
 * sleep and wakeup on one-time events.
 * before any calls to notesleep or notewakeup,
 * must call noteclear to initialize the Note.
 * then, any number of threads can call notesleep
 * and exactly one thread can call notewakeup (once).
 * once notewakeup has been called, all the notesleeps
 * will return.  future notesleeps will return immediately.
 */
void	noteclear(Note*);
void	notesleep(Note*);
void	notewakeup(Note*);

/*
 * Redefine methods for the benefit of gcc, which does not support
 * UTF-8 characters in identifiers.
 */
#ifndef __GNUC__
#define runtime_memclr runtime·memclr
#define runtime_getcallerpc runtime·getcallerpc
#define runtime_mmap runtime·mmap
#define runtime_printslice runtime·printslice
#define runtime_printbool runtime·printbool
#define runtime_printfloat runtime·printfloat
#define runtime_printhex runtime·printhex
#define runtime_printint runtime·printint
#define runtime_printiface runtime·printiface
#define runtime_printeface runtime·printeface
#define runtime_printpc runtime·printpc
#define runtime_printpointer runtime·printpointer
#define runtime_printstring runtime·printstring
#define runtime_printuint runtime·printuint
#define runtime_setcallerpc runtime·setcallerpc
#endif

/*
 * low level go-called
 */
uint8*	runtime_mmap(byte*, uint32, int32, int32, int32, uint32);
void	runtime_memclr(byte*, uint32);
void	runtime_setcallerpc(void*, void*);
void*	runtime_getcallerpc(void*);

/*
 * runtime go-called
 */
void	runtime_printbool(bool);
void	runtime_printfloat(float64);
void	runtime_printint(int64);
void	runtime_printiface(Iface);
void	runtime_printeface(Eface);
void	runtime_printstring(String);
void	runtime_printpc(void*);
void	runtime_printpointer(void*);
void	runtime_printuint(uint64);
void	runtime_printhex(uint64);
void	runtime_printslice(Slice);

/*
 * wrapped for go users
 */
float64	Inf(int32 sign);
float64	NaN(void);
float32	float32frombits(uint32 i);
uint32	float32tobits(float32 f);
float64	float64frombits(uint64 i);
uint64	float64tobits(float64 f);
float64	frexp(float64 d, int32 *ep);
bool	isInf(float64 f, int32 sign);
bool	isNaN(float64 f);
float64	ldexp(float64 d, int32 e);
float64	modf(float64 d, float64 *ip);
void	semacquire(uint32*);
void	semrelease(uint32*);

void	mapassign(Hmap*, byte*, byte*);
void	mapaccess(Hmap*, byte*, byte*, bool*);
struct hash_iter*	mapiterinit(Hmap*);
void	mapiternext(struct hash_iter*);
bool	mapiterkey(struct hash_iter*, void*);
void	mapiterkeyvalue(struct hash_iter*, void*, void*);
Hmap*	makemap(Type*, Type*, uint32);

Hchan*	makechan(Type*, uint32);
void	chansend(Hchan*, void*, bool*);
void	chanrecv(Hchan*, void*, bool*);
void	chanclose(Hchan*);
bool	chanclosed(Hchan*);
int32	chanlen(Hchan*);
int32	chancap(Hchan*);

void	ifaceE2I(struct InterfaceType*, Eface, Iface*);
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
