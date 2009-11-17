<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/proc.c</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/proc.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;
#include &#34;malloc.h&#34;

typedef struct Sched Sched;

M	m0;
G	g0;	// idle goroutine for m0

static	int32	debug	= 0;

// Go scheduler
//
// The go scheduler&#39;s job is to match ready-to-run goroutines (`g&#39;s)
// with waiting-for-work schedulers (`m&#39;s).  If there are ready gs
// and no waiting ms, ready() will start a new m running in a new
// OS thread, so that all ready gs can run simultaneously, up to a limit.
// For now, ms never go away.
//
// By default, Go keeps only one kernel thread (m) running user code
// at a single time; other threads may be blocked in the operating system.
// Setting the environment variable $GOMAXPROCS or calling
// runtime.GOMAXPROCS() will change the number of user threads
// allowed to execute simultaneously.  $GOMAXPROCS is thus an
// approximation of the maximum number of cores to use.
//
// Even a program that can run without deadlock in a single process
// might use more ms if given the chance.  For example, the prime
// sieve will use as many ms as there are primes (up to sched.mmax),
// allowing different stages of the pipeline to execute in parallel.
// We could revisit this choice, only kicking off new ms for blocking
// system calls, but that would limit the amount of parallel computation
// that go would try to do.
//
// In general, one could imagine all sorts of refinements to the
// scheduler, but the goal now is just to get something working on
// Linux and OS X.

struct Sched {
	Lock;

	G *gfree;	// available gs (status == Gdead)

	G *ghead;	// gs waiting to run
	G *gtail;
	int32 gwait;	// number of gs waiting to run
	int32 gcount;	// number of gs that are alive

	M *mhead;	// ms waiting for work
	int32 mwait;	// number of ms waiting for work
	int32 mcount;	// number of ms that have been created
	int32 mcpu;	// number of ms executing on cpu
	int32 mcpumax;	// max number of ms allowed on cpu
	int32 gomaxprocs;
	int32 msyscall;	// number of ms in system calls

	int32 predawn;	// running initialization, don&#39;t run new gs.

	Note	stopped;	// one g can wait here for ms to stop
	int32 waitstop;	// after setting this flag
};

Sched sched;

// Scheduling helpers.  Sched must be locked.
static void gput(G*);	// put/get on ghead/gtail
static G* gget(void);
static void mput(M*);	// put/get on mhead
static M* mget(G*);
static void gfput(G*);	// put/get on gfree
static G* gfget(void);
static void matchmg(void);	// match ms to gs
static void readylocked(G*);	// ready, but sched is locked
static void mnextg(M*, G*);

// Scheduler loop.
static void scheduler(void);

// The bootstrap sequence is:
//
//	call osinit
//	call schedinit
//	make &amp; queue new G
//	call mstart
//
// The new G does:
//
//	call main·init_function
//	call initdone
//	call main·main
void
schedinit(void)
{
	int32 n;
	byte *p;

	allm = m;

	mallocinit();
	goargs();

	// Allocate internal symbol table representation now,
	// so that we don&#39;t need to call malloc when we crash.
	findfunc(0);

	sched.gomaxprocs = 1;
	p = getenv(&#34;GOMAXPROCS&#34;);
	if(p != nil &amp;&amp; (n = atoi(p)) != 0)
		sched.gomaxprocs = n;
	sched.mcpumax = sched.gomaxprocs;
	sched.mcount = 1;
	sched.predawn = 1;
}

// Called after main·init_function; main·main will be called on return.
void
initdone(void)
{
	// Let&#39;s go.
	sched.predawn = 0;
	mstats.enablegc = 1;

	// If main·init_function started other goroutines,
	// kick off new ms to handle them, like ready
	// would have, had it not been pre-dawn.
	lock(&amp;sched);
	matchmg();
	unlock(&amp;sched);
}

void
goexit(void)
{
	g-&gt;status = Gmoribund;
	gosched();
}

void
tracebackothers(G *me)
{
	G *g;

	for(g = allg; g != nil; g = g-&gt;alllink) {
		if(g == me || g-&gt;status == Gdead)
			continue;
		printf(&#34;\ngoroutine %d:\n&#34;, g-&gt;goid);
		traceback(g-&gt;sched.pc, g-&gt;sched.sp, g);
	}
}

// Put on `g&#39; queue.  Sched must be locked.
static void
gput(G *g)
{
	M *m;

	// If g is wired, hand it off directly.
	if((m = g-&gt;lockedm) != nil) {
		mnextg(m, g);
		return;
	}

	g-&gt;schedlink = nil;
	if(sched.ghead == nil)
		sched.ghead = g;
	else
		sched.gtail-&gt;schedlink = g;
	sched.gtail = g;
	sched.gwait++;
}

// Get from `g&#39; queue.  Sched must be locked.
static G*
gget(void)
{
	G *g;

	g = sched.ghead;
	if(g){
		sched.ghead = g-&gt;schedlink;
		if(sched.ghead == nil)
			sched.gtail = nil;
		sched.gwait--;
	}
	return g;
}

// Put on `m&#39; list.  Sched must be locked.
static void
mput(M *m)
{
	m-&gt;schedlink = sched.mhead;
	sched.mhead = m;
	sched.mwait++;
}

// Get an `m&#39; to run `g&#39;.  Sched must be locked.
static M*
mget(G *g)
{
	M *m;

	// if g has its own m, use it.
	if((m = g-&gt;lockedm) != nil)
		return m;

	// otherwise use general m pool.
	if((m = sched.mhead) != nil){
		sched.mhead = m-&gt;schedlink;
		sched.mwait--;
	}
	return m;
}

// Put on gfree list.  Sched must be locked.
static void
gfput(G *g)
{
	g-&gt;schedlink = sched.gfree;
	sched.gfree = g;
}

// Get from gfree list.  Sched must be locked.
static G*
gfget(void)
{
	G *g;

	g = sched.gfree;
	if(g)
		sched.gfree = g-&gt;schedlink;
	return g;
}

// Mark g ready to run.
void
ready(G *g)
{
	lock(&amp;sched);
	readylocked(g);
	unlock(&amp;sched);
}

// Mark g ready to run.  Sched is already locked.
// G might be running already and about to stop.
// The sched lock protects g-&gt;status from changing underfoot.
static void
readylocked(G *g)
{
	if(g-&gt;m){
		// Running on another machine.
		// Ready it when it stops.
		g-&gt;readyonstop = 1;
		return;
	}

	// Mark runnable.
	if(g-&gt;status == Grunnable || g-&gt;status == Grunning)
		throw(&#34;bad g-&gt;status in ready&#34;);
	g-&gt;status = Grunnable;

	gput(g);
	if(!sched.predawn)
		matchmg();
}

static void
nop(void)
{
}

// Same as readylocked but a different symbol so that
// debuggers can set a breakpoint here and catch all
// new goroutines.
static void
newprocreadylocked(G *g)
{
	nop();	// avoid inlining in 6l
	readylocked(g);
}

// Pass g to m for running.
static void
mnextg(M *m, G *g)
{
	sched.mcpu++;
	m-&gt;nextg = g;
	if(m-&gt;waitnextg) {
		m-&gt;waitnextg = 0;
		notewakeup(&amp;m-&gt;havenextg);
	}
}

// Get the next goroutine that m should run.
// Sched must be locked on entry, is unlocked on exit.
// Makes sure that at most $GOMAXPROCS gs are
// running on cpus (not in system calls) at any given time.
static G*
nextgandunlock(void)
{
	G *gp;

	if(sched.mcpu &lt; 0)
		throw(&#34;negative sched.mcpu&#34;);

	// If there is a g waiting as m-&gt;nextg,
	// mnextg took care of the sched.mcpu++.
	if(m-&gt;nextg != nil) {
		gp = m-&gt;nextg;
		m-&gt;nextg = nil;
		unlock(&amp;sched);
		return gp;
	}

	if(m-&gt;lockedg != nil) {
		// We can only run one g, and it&#39;s not available.
		// Make sure some other cpu is running to handle
		// the ordinary run queue.
		if(sched.gwait != 0)
			matchmg();
	} else {
		// Look for work on global queue.
		while(sched.mcpu &lt; sched.mcpumax &amp;&amp; (gp=gget()) != nil) {
			if(gp-&gt;lockedm) {
				mnextg(gp-&gt;lockedm, gp);
				continue;
			}
			sched.mcpu++;		// this m will run gp
			unlock(&amp;sched);
			return gp;
		}
		// Otherwise, wait on global m queue.
		mput(m);
	}
	if(sched.mcpu == 0 &amp;&amp; sched.msyscall == 0)
		throw(&#34;all goroutines are asleep - deadlock!&#34;);
	m-&gt;nextg = nil;
	m-&gt;waitnextg = 1;
	noteclear(&amp;m-&gt;havenextg);
	if(sched.waitstop &amp;&amp; sched.mcpu &lt;= sched.mcpumax) {
		sched.waitstop = 0;
		notewakeup(&amp;sched.stopped);
	}
	unlock(&amp;sched);

	notesleep(&amp;m-&gt;havenextg);
	if((gp = m-&gt;nextg) == nil)
		throw(&#34;bad m-&gt;nextg in nextgoroutine&#34;);
	m-&gt;nextg = nil;
	return gp;
}

// TODO(rsc): Remove. This is only temporary,
// for the mark and sweep collector.
void
stoptheworld(void)
{
	lock(&amp;sched);
	sched.mcpumax = 1;
	while(sched.mcpu &gt; 1) {
		noteclear(&amp;sched.stopped);
		sched.waitstop = 1;
		unlock(&amp;sched);
		notesleep(&amp;sched.stopped);
		lock(&amp;sched);
	}
	unlock(&amp;sched);
}

// TODO(rsc): Remove. This is only temporary,
// for the mark and sweep collector.
void
starttheworld(void)
{
	lock(&amp;sched);
	sched.mcpumax = sched.gomaxprocs;
	matchmg();
	unlock(&amp;sched);
}

// Called to start an M.
void
mstart(void)
{
	if(m-&gt;mcache == nil)
		m-&gt;mcache = allocmcache();
	minit();
	scheduler();
}

// When running with cgo, we call libcgo_thread_start
// to start threads for us so that we can play nicely with
// foreign code.
void (*libcgo_thread_start)(void*);

typedef struct CgoThreadStart CgoThreadStart;
struct CgoThreadStart
{
	M *m;
	G *g;
	void (*fn)(void);
};

// Kick off new ms as needed (up to mcpumax).
// There are already `other&#39; other cpus that will
// start looking for goroutines shortly.
// Sched is locked.
static void
matchmg(void)
{
	M *m;
	G *g;

	while(sched.mcpu &lt; sched.mcpumax &amp;&amp; (g = gget()) != nil){
		// Find the m that will run g.
		if((m = mget(g)) == nil){
			m = malloc(sizeof(M));
			// Add to allm so garbage collector doesn&#39;t free m
			// when it is just in a register (R14 on amd64).
			m-&gt;alllink = allm;
			allm = m;
			m-&gt;g0 = malg(8192);
			m-&gt;id = sched.mcount++;

			if(libcgo_thread_start != nil) {
				CgoThreadStart ts;
				// pthread_create will make us a stack,
				// so free the one malg made.
				stackfree(m-&gt;g0-&gt;stack0);
				m-&gt;g0-&gt;stack0 = nil;
				m-&gt;g0-&gt;stackguard = nil;
				m-&gt;g0-&gt;stackbase = nil;
				ts.m = m;
				ts.g = m-&gt;g0;
				ts.fn = mstart;
				runcgo(libcgo_thread_start, &amp;ts);
			} else
				newosproc(m, m-&gt;g0, m-&gt;g0-&gt;stackbase, mstart);
		}
		mnextg(m, g);
	}
}

// Scheduler loop: find g to run, run it, repeat.
static void
scheduler(void)
{
	G* gp;

	lock(&amp;sched);
	if(gosave(&amp;m-&gt;sched) != 0){
		gp = m-&gt;curg;

		// Jumped here via gosave/gogo, so didn&#39;t
		// execute lock(&amp;sched) above.
		lock(&amp;sched);

		if(sched.predawn)
			throw(&#34;init sleeping&#34;);

		// Just finished running gp.
		gp-&gt;m = nil;
		sched.mcpu--;

		if(sched.mcpu &lt; 0)
			throw(&#34;sched.mcpu &lt; 0 in scheduler&#34;);
		switch(gp-&gt;status){
		case Grunnable:
		case Gdead:
			// Shouldn&#39;t have been running!
			throw(&#34;bad gp-&gt;status in sched&#34;);
		case Grunning:
			gp-&gt;status = Grunnable;
			gput(gp);
			break;
		case Gmoribund:
			gp-&gt;status = Gdead;
			if(gp-&gt;lockedm) {
				gp-&gt;lockedm = nil;
				m-&gt;lockedg = nil;
			}
			gfput(gp);
			if(--sched.gcount == 0)
				exit(0);
			break;
		}
		if(gp-&gt;readyonstop){
			gp-&gt;readyonstop = 0;
			readylocked(gp);
		}
	}

	// Find (or wait for) g to run.  Unlocks sched.
	gp = nextgandunlock();
	gp-&gt;readyonstop = 0;
	gp-&gt;status = Grunning;
	m-&gt;curg = gp;
	gp-&gt;m = m;
	if(gp-&gt;sched.pc == (byte*)goexit)	// kickoff
		gogocall(&amp;gp-&gt;sched, (void(*)(void))gp-&gt;entry);
	gogo(&amp;gp-&gt;sched, 1);
}

// Enter scheduler.  If g-&gt;status is Grunning,
// re-queues g and runs everyone else who is waiting
// before running g again.  If g-&gt;status is Gmoribund,
// kills off g.
void
gosched(void)
{
	if(g == m-&gt;g0)
		throw(&#34;gosched of g0&#34;);
	if(gosave(&amp;g-&gt;sched) == 0)
		gogo(&amp;m-&gt;sched, 1);
}

// The goroutine g is about to enter a system call.
// Record that it&#39;s not using the cpu anymore.
// This is called only from the go syscall library, not
// from the low-level system calls used by the runtime.
// The &#34;arguments&#34; are syscall.Syscall&#39;s stack frame
void
runtime·entersyscall(uint64 callerpc, int64 trap)
{
	USED(callerpc, trap);

	lock(&amp;sched);
	if(sched.predawn) {
		unlock(&amp;sched);
		return;
	}
	g-&gt;status = Gsyscall;
	// Leave SP around for gc and traceback.
	// Do before notewakeup so that gc
	// never sees Gsyscall with wrong stack.
	gosave(&amp;g-&gt;sched);
	sched.mcpu--;
	sched.msyscall++;
	if(sched.gwait != 0)
		matchmg();
	if(sched.waitstop &amp;&amp; sched.mcpu &lt;= sched.mcpumax) {
		sched.waitstop = 0;
		notewakeup(&amp;sched.stopped);
	}
	unlock(&amp;sched);
}

// The goroutine g exited its system call.
// Arrange for it to run on a cpu again.
// This is called only from the go syscall library, not
// from the low-level system calls used by the runtime.
void
runtime·exitsyscall(void)
{
	lock(&amp;sched);
	if(sched.predawn) {
		unlock(&amp;sched);
		return;
	}
	g-&gt;status = Grunning;
	sched.msyscall--;
	sched.mcpu++;
	// Fast path - if there&#39;s room for this m, we&#39;re done.
	if(sched.mcpu &lt;= sched.mcpumax) {
		unlock(&amp;sched);
		return;
	}
	unlock(&amp;sched);

	// Slow path - all the cpus are taken.
	// The scheduler will ready g and put this m to sleep.
	// When the scheduler takes g away from m,
	// it will undo the sched.mcpu++ above.
	gosched();
}

/*
 * stack layout parameters.
 * known to linkers.
 *
 * g-&gt;stackguard is set to point StackGuard bytes
 * above the bottom of the stack.  each function
 * compares its stack pointer against g-&gt;stackguard
 * to check for overflow.  to cut one instruction from
 * the check sequence for functions with tiny frames,
 * the stack is allowed to protrude StackSmall bytes
 * below the stack guard.  functions with large frames
 * don&#39;t bother with the check and always call morestack.
 * the sequences are:
 *
 *	guard = g-&gt;stackguard
 *	frame = function&#39;s stack frame size
 *	argsize = size of function arguments (call + return)
 *
 *	stack frame size &lt;= StackSmall:
 *		CMPQ guard, SP
 *		JHI 3(PC)
 *		MOVQ m-&gt;morearg, $(argsize &lt;&lt; 32)
 *		CALL sys.morestack(SB)
 *
 *	stack frame size &gt; StackSmall but &lt; StackBig
 *		LEAQ (frame-StackSmall)(SP), R0
 *		CMPQ guard, R0
 *		JHI 3(PC)
 *		MOVQ m-&gt;morearg, $(argsize &lt;&lt; 32)
 *		CALL sys.morestack(SB)
 *
 *	stack frame size &gt;= StackBig:
 *		MOVQ m-&gt;morearg, $((argsize &lt;&lt; 32) | frame)
 *		CALL sys.morestack(SB)
 *
 * the bottom StackGuard - StackSmall bytes are important:
 * there has to be enough room to execute functions that
 * refuse to check for stack overflow, either because they
 * need to be adjacent to the actual caller&#39;s frame (sys.deferproc)
 * or because they handle the imminent stack overflow (sys.morestack).
 *
 * for example, sys.deferproc might call malloc,
 * which does one of the above checks (without allocating a full frame),
 * which might trigger a call to sys.morestack.
 * this sequence needs to fit in the bottom section of the stack.
 * on amd64, sys.morestack&#39;s frame is 40 bytes, and
 * sys.deferproc&#39;s frame is 56 bytes.  that fits well within
 * the StackGuard - StackSmall = 128 bytes at the bottom.
 * there may be other sequences lurking or yet to be written
 * that require more stack.  sys.morestack checks to make sure
 * the stack has not completely overflowed and should
 * catch such sequences.
 */
enum
{
	// byte offset of stack guard (g-&gt;stackguard) above bottom of stack.
	StackGuard = 256,

	// checked frames are allowed to protrude below the guard by
	// this many bytes.  this saves an instruction in the checking
	// sequence when the stack frame is tiny.
	StackSmall = 128,

	// extra space in the frame (beyond the function for which
	// the frame is allocated) is assumed not to be much bigger
	// than this amount.  it may not be used efficiently if it is.
	StackBig = 4096,
};

void
oldstack(void)
{
	Stktop *top, old;
	uint32 args;
	byte *sp;
	G *g1;

//printf(&#34;oldstack m-&gt;cret=%p\n&#34;, m-&gt;cret);

	g1 = m-&gt;curg;
	top = (Stktop*)g1-&gt;stackbase;
	sp = (byte*)top;
	old = *top;
	args = old.args;
	if(args &gt; 0) {
		sp -= args;
		mcpy(top-&gt;fp, sp, args);
	}

	stackfree((byte*)g1-&gt;stackguard - StackGuard);
	g1-&gt;stackbase = old.stackbase;
	g1-&gt;stackguard = old.stackguard;

	gogo(&amp;old.gobuf, m-&gt;cret);
}

void
newstack(void)
{
	int32 frame, args;
	Stktop *top;
	byte *stk, *sp;
	G *g1;
	Gobuf label;

	frame = m-&gt;moreframe;
	args = m-&gt;moreargs;

	// Round up to align things nicely.
	// This is sufficient for both 32- and 64-bit machines.
	args = (args+7) &amp; ~7;

	if(frame &lt; StackBig)
		frame = StackBig;
	frame += 1024;	// for more functions, Stktop.
	stk = stackalloc(frame);

//printf(&#34;newstack frame=%d args=%d morepc=%p morefp=%p gobuf=%p, %p newstk=%p\n&#34;, frame, args, m-&gt;morepc, m-&gt;morefp, g-&gt;sched.pc, g-&gt;sched.sp, stk);

	g1 = m-&gt;curg;
	top = (Stktop*)(stk+frame-sizeof(*top));
	top-&gt;stackbase = g1-&gt;stackbase;
	top-&gt;stackguard = g1-&gt;stackguard;
	top-&gt;gobuf = m-&gt;morebuf;
	top-&gt;fp = m-&gt;morefp;
	top-&gt;args = args;

	g1-&gt;stackbase = (byte*)top;
	g1-&gt;stackguard = stk + StackGuard;

	sp = (byte*)top;
	if(args &gt; 0) {
		sp -= args;
		mcpy(sp, m-&gt;morefp, args);
	}

	// Continue as if lessstack had just called m-&gt;morepc
	// (the PC that decided to grow the stack).
	label.sp = sp;
	label.pc = (byte*)runtime·lessstack;
	label.g = m-&gt;curg;
	gogocall(&amp;label, m-&gt;morepc);

	*(int32*)345 = 123;	// never return
}

G*
malg(int32 stacksize)
{
	G *g;
	byte *stk;

	g = malloc(sizeof(G));
	stk = stackalloc(stacksize + StackGuard);
	g-&gt;stack0 = stk;
	g-&gt;stackguard = stk + StackGuard;
	g-&gt;stackbase = stk + StackGuard + stacksize;
	return g;
}

/*
 * Newproc and deferproc need to be textflag 7
 * (no possible stack split when nearing overflow)
 * because they assume that the arguments to fn
 * are available sequentially beginning at &amp;arg0.
 * If a stack split happened, only the one word
 * arg0 would be copied.  It&#39;s okay if any functions
 * they call split the stack below the newproc frame.
 */
#pragma textflag 7
void
runtime·newproc(int32 siz, byte* fn, byte* arg0)
{
	byte *stk, *sp;
	G *newg;

//printf(&#34;newproc siz=%d fn=%p&#34;, siz, fn);

	siz = (siz+7) &amp; ~7;
	if(siz &gt; 1024)
		throw(&#34;runtime·newproc: too many args&#34;);

	lock(&amp;sched);

	if((newg = gfget()) != nil){
		newg-&gt;status = Gwaiting;
	} else {
		newg = malg(4096);
		newg-&gt;status = Gwaiting;
		newg-&gt;alllink = allg;
		allg = newg;
	}
	stk = newg-&gt;stack0;

	newg-&gt;stackguard = stk+StackGuard;

	sp = stk + 4096 - 4*8;
	newg-&gt;stackbase = sp;

	sp -= siz;
	mcpy(sp, (byte*)&amp;arg0, siz);

	newg-&gt;sched.sp = sp;
	newg-&gt;sched.pc = (byte*)goexit;
	newg-&gt;sched.g = newg;
	newg-&gt;entry = fn;

	sched.gcount++;
	goidgen++;
	newg-&gt;goid = goidgen;

	newprocreadylocked(newg);
	unlock(&amp;sched);

//printf(&#34; goid=%d\n&#34;, newg-&gt;goid);
}

#pragma textflag 7
void
runtime·deferproc(int32 siz, byte* fn, byte* arg0)
{
	Defer *d;

	d = malloc(sizeof(*d) + siz - sizeof(d-&gt;args));
	d-&gt;fn = fn;
	d-&gt;sp = (byte*)&amp;arg0;
	d-&gt;siz = siz;
	mcpy(d-&gt;args, d-&gt;sp, d-&gt;siz);

	d-&gt;link = g-&gt;defer;
	g-&gt;defer = d;
}

#pragma textflag 7
void
runtime·deferreturn(uintptr arg0)
{
	Defer *d;
	byte *sp, *fn;

	d = g-&gt;defer;
	if(d == nil)
		return;
	sp = (byte*)&amp;arg0;
	if(d-&gt;sp != sp)
		return;
	mcpy(d-&gt;sp, d-&gt;args, d-&gt;siz);
	g-&gt;defer = d-&gt;link;
	fn = d-&gt;fn;
	free(d);
	jmpdefer(fn, sp);
  }

void
runtime·Breakpoint(void)
{
	breakpoint();
}

void
runtime·Goexit(void)
{
	goexit();
}

void
runtime·Gosched(void)
{
	gosched();
}

void
runtime·LockOSThread(void)
{
	if(sched.predawn)
		throw(&#34;cannot wire during init&#34;);
	m-&gt;lockedg = g;
	g-&gt;lockedm = m;
}

// delete when scheduler is stronger
void
runtime·GOMAXPROCS(int32 n)
{
	if(n &lt; 1)
		n = 1;

	lock(&amp;sched);
	sched.gomaxprocs = n;
	sched.mcpumax = n;
	// handle fewer procs
	while(sched.mcpu &gt; sched.mcpumax) {
		noteclear(&amp;sched.stopped);
		sched.waitstop = 1;
		unlock(&amp;sched);
		notesleep(&amp;sched.stopped);
		lock(&amp;sched);
	}
	// handle more procs
	matchmg();
	unlock(&amp;sched);
}

void
runtime·UnlockOSThread(void)
{
	m-&gt;lockedg = nil;
	g-&gt;lockedm = nil;
}

// for testing of wire, unwire
void
runtime·mid(uint32 ret)
{
	ret = m-&gt;id;
	FLUSH(&amp;ret);
}
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
